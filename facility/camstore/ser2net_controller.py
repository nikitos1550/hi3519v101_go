import logging
import signal
import os
import asyncio
import asyncio.subprocess

import common
from telnetlib import Telnet
from server import register, routine
from devices import Devices, Device


class DeviceState:
    def __init__(self):
        pass


class Ser2NetWrap:
    @property
    def control_port(self):
        return int(self.config["control_port"])

    @property
    def config_file(self):
        return self.config["config_file"]

    @property
    def is_running(self):
        return (self.s2n_proc is not None) and (self.s2n_proc.returncode is None)

    def __init__(self, config={}):
        self.config = config
        self.config.setdefault("control_port", 45300)
        self.config.setdefault("config_file", "./ser2net.cfg")

        self.s2n_proc = None
        self.control = None

        self._devname_to_port = {}
        self._devname_to_owner = {}

    async def start(self):
        open(self.config_file, "w").close()  # ensure that config file exists and it's clear

        args = [
            "-d",
            "-p", "localhost,{}".format(self.control_port),
            "-c", self.config_file
        ]
        logging.debug("run subprocess: ser2net {}".format(" ".join(args)))
        self.s2n_proc = await asyncio.subprocess.create_subprocess_exec("ser2net", *args, stderr=asyncio.subprocess.DEVNULL)
        logging.info("ser2net started with PID {}".format(self.s2n_proc.pid))

        await asyncio.sleep(3)
        if not self.is_running:
            raise RuntimeError("ser2net process died very fast")

        logging.debug("connect to ser2net control port...")
        await asyncio.sleep(1)
        self.control = Telnet(host="localhost", port=self.control_port)
        logging.debug("connection with ser2net control port established")
        self.control.read_until(b"-> ")
        logging.info("ser2net control is ready")

    async def stop(self):
        if not self.is_running:
            return

        logging.debug("ser2net process is stopping...")
        self.s2n_proc.terminate()
        try:
            await asyncio.wait_for(self.s2n_proc.wait(), timeout=5)
            logging.info("ser2net process successfully terminated")
        except asyncio.TimeoutError:
            self.s2n_proc.kill()
            await self.s2n_proc.wait()
            logging.warn("ser2net process killed")

    def update_config(self, devices):
        config_lines = []

        port = self.control_port + 1
        for dev in devices:
            self._devname_to_port[dev.devname] = port
            config_lines.append("{}:off:300:{}:115200\n".format(port, dev.devname))
            port += 1

        with open(self.config_file, "w") as f:
            f.writelines(config_lines)

        self.s2n_proc.send_signal(signal.SIGHUP)

    def cmd(self, command):
        data = command.encode("ascii") + b"\n\r"
        self.control.write(data)
        res = self.control.read_until(b"-> ")
        return res[len(data):-5].decode("ascii")

    def _acquire_device(self, devname, user):
        owner = self._devname_to_owner.get(devname)
        if (owner is not None) and (owner != user):
            raise common.InvalidArgument("device '{}' is already acquired by '{}'".format(devname, owner))
        self._devname_to_owner[devname] = user
        logging.info("Device '{}' is acquired by '{}'".format(devname, user))

    def disconnect(self, port):
        return self.cmd("disconnect localhost,{}".format(port))

    def showport(self, port=""):
        return self.cmd("showshortport {}".format(port))

    def port_state(self, devname):
        """
        Port-name Type Timeout Remote-address Device        TCP-to-device  Device-to-TCP TCP-in TCP-out Dev-in Dev-out State
        45330     off  300     unconnected    /dev/ttyCAM30 unconnected    unconnected   0      0       0      0       115200 1STOPBIT 8DATABITS NONE
        """
        port = self._devname_to_port.get(devname)
        if port is None:
            raise common.InvalidArgument("device '{}' does not exist".format(devname))

        res = self.cmd("showshortport {}".format(port)).split("\n")[1].strip().split(" ")
        res = filter(None, res)

        result = {}
        fields = ("port", "type", "timeout", "remote_addr", "device")
        for field in fields:
            result[field] = res.__next__()

        return result

    def forward_device(self, devname, user, mode="tenet"):
        port = self._devname_to_port.get(devname)
        if port is None:
            raise common.InvalidArgument("device '{}' does not exist".format(devname))

        self._acquire_device(devname, user)

        self.cmd("setportenable {} {}".format(port, mode))
        logging.info("Forward device '{}' to {} TCP port in '{}' mode".format(devname, port, mode))

        return "ok/exec: telnet localhost {}".format(port)


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
    return "\n".join(devname for devname in __devices.devs.keys())


@register
def forward_serial(devname, user, mode="telnet"):
    """ Forward a device's serial port to TCP
    args: devname user [mode=telnet]
    """
    return __s2n_wrap.forward_device(devname, user, mode)


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
