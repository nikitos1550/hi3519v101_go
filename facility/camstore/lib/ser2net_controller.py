import logging
import os
import asyncio
import time

from .common import success, failure, InvalidArgument
from .ser2net import Ser2Net
from .server import register, routine
from .devices import Devices, Device


class DeviceState:
    def __init__(self, devname):
        self.devname = devname
        self.owner = None
        self.active_ts = None
        self.endpoint = None

    def acquire(self, user):
        if self.owner not in (user, None):
            raise InvalidArgument(f"device '{self.devname}' is already acquired by '{self.owner}'")
        self.owner = user
        self.active_ts = time.time()
        logging.info(f"device '{self.devname}' is acquired by '{user}'")
    
    def release(self):
        self.owner = None
        self.active_ts = None


class Ser2NetWrap:
    @property
    def control_port(self):
        return self.s2n.control_port

    @property
    def config_file(self):
        return self.s2n.config_file

    @property
    def is_running(self):
        return self.s2n.is_running

    def __init__(self, config={}):
        config.setdefault("control_port", 45300)
        config.setdefault("config_file", os.path.join(os.path.dirname(__file__), "ser2net.cfg"))
        config.setdefault("ttl", 300)  # in seconds

        self.config = config
        self.s2n = Ser2Net(config)
        self._devname_to_state = {}

    async def start(self):
        await self.s2n.start()

    async def stop(self):
        await self.s2n.stop()

    def update_config(self, devices):
        config_lines = []

        port = self.control_port + 1
        for dev in devices:
            endpoint = "localhost,{}".format(port)
            config_lines.append("{}:off:300:{}:115200\n".format(endpoint, dev.devname))
            devstate = self._devname_to_state.setdefault(dev.devname, DeviceState(dev.devname))
            devstate.endpoint = endpoint
            port += 1

        with open(self.config_file, "w") as f:
            f.writelines(config_lines)

        self.s2n.reload_config()

    async def update_states(self):
        now = time.time()
        logging.debug(f"Update devices' states at {now} ts")
        for devname, devstate in self._devname_to_state.items():
            if devstate.owner is None:
                continue  # device is free

            acquire_duration = now - devstate.active_ts
            logging.debug(f"Device '{devname}' has been acquired for {acquire_duration:.0f} seconds by {devstate.owner}")

            if acquire_duration < self.config["ttl"]:
                continue  # too late to kick
            portstate = self.s2n.showport(devstate.endpoint)
            if portstate["remote_addr"] != "unconnected":
                continue  # owner is actve (accordingly ser2net)

            logging.info(f"Release device '{devname}'")
            try:
                self.s2n.disconnect(devstate.endpoint)
            except: pass
            devstate.release()
            await asyncio.sleep(0) # to be asynchronous

    def _get_state(self, devname):
        state = self._devname_to_state.get(devname)
        if state is None:
            raise InvalidArgument("device '{}' does not exist".format(devname))
        return state

    def disconnect(self, port):
        return self.s2n.disconnect("localhost,{}".format(port))

    def showport(self, port):
        return self.s2n.showport("localhost,{}".format(port))
    
    def acquire_device(self, devname, user):
        devstate = self._get_state(devname)
        devstate.acquire(user)
        logging.info(f"Device '{devname}' acquired by '{user}'")

    def release_device(self, devname, user):
        devstate = self._get_state(devname)
        if devstate.owner != user:
            raise InvalidArgument(f"device '{devname}' isn't owned by '{user}'")
        try:
            self.s2n.disconnect(devstate.endpoint)
        except: pass
        devstate.release()
        logging.info(f"Device '{devname}' released by '{user}'")

    def forward_device(self, devname, user, mode="telnet"):
        devstate = self._get_state(devname)
        devstate.acquire(user)

        self.s2n.setportenable(devstate.endpoint, mode)
        logging.info(f"forward device '{devname}' to {devstate.endpoint} in '{mode}' mode")

        telnet_dest = devstate.endpoint.replace(",", " ")
        return success(
            message=f"{devname} is forwarded to {telnet_dest}",
            execute=f"telnet {telnet_dest}"
        )

    def port_state(self, devname):
        devstate = self._get_state(devname)
        return success(
            message=str(self.s2n.showport(devstate.endpoint))
        )


# -------------------------------------------------------------------------------------------------


__devices = None
__s2n_wrap = None
__device_owners = {}


@routine
async def main_routine():
    global __devices, __s2n_wrap

    logging.info("Main routine started")

    __devices = Devices()
    __s2n_wrap = Ser2NetWrap()

    try:
        await __s2n_wrap.start()

        while True:
            if __devices.update():
                logging.info("Devices were changed, update ser2net config")
                __s2n_wrap.update_config(__devices.devs.values())
            await __s2n_wrap.update_states()
            await asyncio.sleep(10)
    except:
        logging.debug("exceptions occured in main_routine")
        await __s2n_wrap.stop()
        raise


@register
def list_devices():
    """ Print list of available devices
    """
    return success("\n".join(
        f"{d.devname} {d.model}" for d in __devices.devs.values()
    ))


@register
def forward_serial(devname, user, mode="telnet"):
    """ Acquire device and forward its' serial port to TCP
    args: devname [mode=telnet]
    """
    return __s2n_wrap.forward_device(devname, user, mode)


@register
def acquire_device(devname, user):
    """ Acquire device for 5 minutes
    args: devname
    """
    return __s2n_wrap.acquire_device(devname, user)


@register
def release_device(devname, user):
    """ Release device acquired by user
    args: devname
    """
    return __s2n_wrap.release_device(devname, user)


@register
def device_state(devname):
    """ Print a device's state
    args: devname
    """
    devstate = __s2n_wrap._get_state(devname)
    dev = __devices.devs.get(devname)
    return success(f"device={devname} model={dev.model} owner={devstate.owner}")

@register
def power(devname, user, mode):
    """ Control a device's power (ON/OFF)
    args: devname user {ON|OFF|RESET}
    """
    mode = mode.lower()
    if mode not in ("on", "off", "reset"):
        raise InvalidArgument("`mode` argument MUST be ON|OFF|RESET")
    return failure("Not implemented, sorry")

@register
def s2n_port_state(devname):
    """ System: print ser2net port state
    args: devname
    """
    return str(__s2n_wrap.port_state(devname))


# eloop = asyncio.get_event_loop()
#
# eloop.run_until_complete(s2n.start())
# print(s2n.showport())
# print(s2n.add_device("/dev/x", 43521))
# eloop.run_until_complete(asyncio.sleep(2))
# print(s2n.showport())
# eloop.run_until_complete(s2n.stop())
