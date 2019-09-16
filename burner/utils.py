import random, logging, sys, time
import serial


# ----------------------------------------------------------------------------------------------------------------------
# Constants

CTRL_C = "\x03"

DUPLEX_FULL = 1
DUPLEX_HALF = 2

PROMPTS = ("xmtech #", "hisilicon #")

def line_has_prompt(line):
    return any(map(lambda p: p in line, PROMPTS))

# ----------------------------------------------------------------------------------------------------------------------


def __init_logger(
    name,
    level=logging.INFO,
    fmt="[%(name)s:%(levelname)s, %(asctime)s.%(msecs)d] %(message)s",
    datefmt="%H:%M:%S",
    stream=sys.stdout
):
    logger = logging.getLogger(name)
    if logger.propagate:
        # means the logger hasn't been initialized yet
        handler = logging.StreamHandler(stream)
        handler.setFormatter(logging.Formatter(fmt=fmt, datefmt=datefmt))
        logger.propagate = False
        logger.setLevel(level)
        logger.addHandler(handler)
    return logger


def get_device_logger(name, level=logging.DEBUG):
    return __init_logger(name, level, fmt="[%(levelname)s at %(asctime)s.%(msecs)d] %(name)s %(message)s")


__conn_logger = __init_logger("conn", fmt="[%(name)s %(asctime)s.%(msecs)d] %(message)s")


# Developer's logger
__dev_logger = __init_logger(
    name="develop",
    level=logging.DEBUG,
    fmt="[%(name)s:%(levelname)s][%(filename)s:%(lineno)d] %(message)s"
)

def DLOG_WARN(msg):
    __dev_logger.warning(msg)

def DLOG_INFO(msg):
    __dev_logger.info(msg)

def DLOG(msg):
    __dev_logger.debug(msg)


# Logging of interaction with device
def CLOG_INCOMING(msg):
    __conn_logger.info("<- " + msg)

def CLOG_OUTGOING(msg):
    __conn_logger.info("-> " + msg)


# ----------------------------------------------------------------------------------------------------------------------


class Device:
    def __init__(self, port, baudrate, timeout, duplex = DUPLEX_FULL):
        self._serial_port = serial.Serial(
            port=port,
            baudrate=baudrate,
            timeout=timeout
        )
        self._duplex = duplex
    
    def close():

        self._serial_port.close()

    def write_data(self, data):
        time.sleep(0.1)
        for item in send:
            self._serial_port.write(item)
            time.sleep(0.1)

    def write_cmd(self, cmd):
        cmd = (cmd.replace(";", "\;") + "\n").encode("ascii")
        if self._duplex == DUPLEX_HALF:
            self.write_data(cmd)
        else:
            self._serial_port.write(cmd)
        CLOG_OUTGOING(cmd)

    def write_ctrlc(self):
        """ Send Ctrl+C to device
        """
        self.write_data(CTRL_C)

    def wait_prompt(self, clear=True):
        """If 'clear' then we wait for a line contains only prompt, otherwise any
        line with prompt is fine.
        """
        while True:
            line = self.read_line()
            if clear:
                if (line in PROMPTS):
                    break
            elif line_has_prompt(line):
                break
            
    def read_line(self):
        line = self._serial_port.readline().strip()
        CLOG_INCOMING(line)
        return line


# ----------------------------------------------------------------------------------------------------------------------


def random_mac():
    mac = [ 0x00, 0x00, 0x23,  # first line is defined for specified vendor
            random.randint(0x01, 0xfe),
            random.randint(0x01, 0xfe),
            random.randint(0x01, 0xfe) ]
    return ':'.join(map(lambda x: "%02x" % x, mac))


def get_iface_ip_and_mask(if_name, ipv6=False):
    try:
        import netifaces

        addrs = netifaces.ifaddresses(if_name).get(netifaces.AF_INET6 if ipv6 else netifaces.AF_INET)
        if (addrs is None) or (len(addrs) == 0):
            raise ValueError("Network interface '{}' has no addresses".format(if_name))

        addr, netmask = addrs[0]["addr"], addrs[0]["netmask"]
        DLOG("Interface '{}' has {} addresses, use the first: addr={} netmask={}".format(
            if_name, len(addrs), addr, netmask))
        return addr, netmask
    except ValueError:
        raise ValueError("Network interface '{}' not found, available: {}".format(
            if_name, ", ".join(netifaces.interfaces())
        ))
    except ImportError:
        DLOG_WARN("'netifaces' module not found")
        raise
