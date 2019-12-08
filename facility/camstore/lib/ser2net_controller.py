import logging
import signal
import os
import asyncio
import asyncio.subprocess
from telnetlib import Telnet

from .common import success, failure, InvalidArgument
from .ser2net import Ser2Net
from .server import register, routine
from .devices import Devices, Device


class DeviceState:
    def __init__(self):
        pass


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

        self.s2n = Ser2Net(config)

        self._devname_to_port = {}
        self._devname_to_owner = {}
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
            self._devname_to_state.setdefault(dev.devname, {})["endpoint"] = endpoint
            config_lines.append("{}:off:300:{}:115200\n".format(endpoint, dev.devname))
            port += 1

        with open(self.config_file, "w") as f:
            f.writelines(config_lines)

        self.s2n.reload_config()

    def _get_state(self, devname):
        state = self._devname_to_state.get(devname)
        if state is None:
            raise InvalidArgument("device '{}' does not exist".format(devname))
        return state

    def _acquire_device(self, devstate, user):
        owner = devstate.get("owner")
        if (owner is not None) and (owner != user):
            raise InvalidArgument("device is already acquired by '{}'".format(owner))
        devstate["owner"] = user
        logging.info("device is acquired by '{}'".format(user))

    def disconnect(self, port):
        return self.s2n.disconnect("localhost,{}".format(port))

    def showport(self, port):
        return self.s2n.showport("localhost,{}".format(port))

    def forward_device(self, devname, user, mode="tenet"):
        devstate = self._get_state(devname)

        self._acquire_device(devstate, user)

        endpoint = devstate["endpoint"]
        self.s2n.setportenable(endpoint, mode)
        logging.info("forward device '{}' to {} in '{}' mode".format(devname, endpoint, mode))

        telnet_dest = endpoint.replace(",", " ")
        return success(
            message="{} is forwarded to {}".format(devname, telnet_dest),
            execute="telnet {}".format(telnet_dest)
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
            await asyncio.sleep(10)
    except:
        logging.debug("exceptions occured in main_routine")
        await __s2n_wrap.stop()
        raise


@register
def list_devices():
    """ Print list of available camera devices
    """
    return success("\n".join(devname for devname in __devices.devs.keys()))


@register
def forward_serial(devname, user, mode="telnet"):
    """ Forward a device's serial port to TCP
    args: devname user [mode=telnet]
    """
    return __s2n_wrap.forward_device(devname, user, mode)


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
def sys_s2n_port_state(devname):
    """ System: print ser2net port state
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
