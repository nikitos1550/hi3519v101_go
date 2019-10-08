# if __name__ == "__main__":
from device import Device
from utils import get_device_logger
# else:
#     from .device import Device
#     from .utils import get_device_logger


class UBootConsole:
    READ_TIMEOUT    = 0.1
    ENCODING        = "utf-8"
    CTRL_C          = b"\x03"
    LF              = b"\n"
    
    # Parameters below are supposed to be configured for particular boards
    GREETING            = b"System startup"
    PROMPT              = b"hisilicon #"
    AUTOBOOT_STOP_KEY   = b"\x03"

    @classmethod
    def catch_console(cls, **kw):
        uboot = cls(**kw)

        uboot.dlog("Wait for greeting line: {} ...", cls.GREETING)
        while cls.GREETING not in uboot.device.read_line():
            pass
        uboot.dlog("Greeting line received, send '{}' key", cls.AUTOBOOT_STOP_KEY)
        uboot.device.write_data(cls.AUTOBOOT_STOP_KEY)

        uboot.dlog("Wait for prompt: {} ...", cls.PROMPT)
        while not uboot.device.read_line().startswith(cls.PROMPT):
            pass
        uboot.dlog("Prompt received")

        return uboot

    def dlog(self, *a, **kw):
        self.device.dlog(*a, **kw)

    def __init__(self, device=None, port=None, baudrate=None, logger=None):
        if device is not None:
            if (port, baudrate) != (None, None):
                raise Exception("device, port&baudrate mustn't be used simultaneously")
            device.serial.timeout = self.READ_TIMEOUT
            device.logger = logger
            self.device = device
        elif None not in (port, baudrate):
            self.device = Device(port=port, baudrate=baudrate, read_timeout=self.READ_TIMEOUT, logger=logger)
        else:
            raise Exception("Either device or port&baudrate must be defines")
        
        self.dlog("UBoot console constructed")

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
            if line == self.PROMPT:
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


def get_uboot_console_type(uboot_params):
    class UBootConsoleParametrized(UBootConsole):
        pass

    def set_if_needed(key):
        if key in uboot_params:
            setattr(UBootConsoleParametrized, key, uboot_params[key])

    set_if_needed("GREETING")
    set_if_needed("PROMPT")
    set_if_needed("AUTOBOOT_STOP_KEY")

    return UBootConsoleParametrized


if __name__ == "__main__":
    import sys

    logger = get_device_logger("uboot")
    uboot = UBootConsole.catch_console(port=sys.argv[1], baudrate=115200, logger=logger)
    
    while True:
        cmd = sys.stdin.readline().strip()
        for l in uboot.command(cmd):
            print(l)
