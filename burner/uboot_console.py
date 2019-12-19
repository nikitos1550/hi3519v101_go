from device import Device, SerialConn, TelnetConn
from utils import get_device_logger


class UBootConsoleParams:
    def __init__(self, params_dict=None):
        if params_dict is None:
            params_dict = {}

        # First (or so) line of U-Boot loading process
        self.GREETING = params_dict.get("GREETING", "System startup").encode("ascii")

        # U-Boot console's prompt
        self.PROMPT = params_dict.get("PROMPT", "hisilicon #").encode("ascii")

        # Key to interrupt autoboot
        self.AUTOBOOT_STOP_KEY = params_dict.get("AUTOBOOT_STOP_KEY", "\x03").encode("ascii")


class UBootConsole:
    READ_TIMEOUT    = 0.1
    ENCODING        = "utf-8"
    CTRL_C          = b"\x03"
    LF              = b"\n"

    def dlog(self, *a, **kw):
        self.device.dlog(*a, **kw)

    def __init__(self, conn, logger=None, params=UBootConsoleParams()):
        conn.set_read_timeout(self.READ_TIMEOUT)
        self.device = Device(conn=conn, logger=logger)
        self.params = params
        self.dlog("UBoot console constructed")

    def fetch_console(self):
        self.device.clear_input_buff()

        self.dlog("Wait for greeting line: {} ...", self.params.GREETING)
        while self.params.GREETING not in self.device.read_line():
            pass

        self.dlog("Greeting line received, send '{}' key", self.params.AUTOBOOT_STOP_KEY)
        self.device.write_data(self.params.AUTOBOOT_STOP_KEY)

        self.dlog("Wait for prompt: {} ...", self.params.PROMPT)
        self.wait_for(self.params.PROMPT)

        self.dlog("Prompt received")

    def write_and_check(self, data):
        while True:
            self.device.write_data(data)
            echoed = self.device.read_line()[-len(data):]
            if echoed == data:
                self.device.write_data(self.LF)
                self.device.read_line()
                return
            self.device.dlog("Failure in echo - got {} instead of {}; retype the command", echoed, data)
            self.device.write_data(self.CTRL_C)
            self.device.read_line()

    def wait_for(self, line):
        while self.device.read_line().strip() != line:
            pass

    def command(self, cmd, wait=True):
        self.write_and_check(cmd.encode(self.ENCODING))
        if not wait:
            return

        response = []
        while True:
            line = self.device.read_line().strip()
            if line == self.params.PROMPT:
                break
            response.append(line.decode(self.ENCODING))
        return response

    def setenv(self, **kwargs):
        for k, v in kwargs.items():
            v = v.replace(";", "\;")
            self.command("setenv {} {}".format(k, v))

    def tftp(self, offset, file_name):
        self.command("tftp {:#x} {}".format(offset, file_name))

    def bootm(self, uimage_addr):
        self.command("bootm {:#x}".format(uimage_addr), wait=False)
