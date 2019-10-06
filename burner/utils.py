import random, logging, sys, os
import tempfile, shutil


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


# =====================================================================================================================
def dlog(message, *args, **kwargs):
    if True:  # TODO: should be configurable
        logging.debug(message.format(*args, **kwargs))


# =====================================================================================================================
def random_mac():
    mac = [ 0x00, 0x00, 0x23,  # first line is defined for specified vendor
            random.randint(0x01, 0xfe),
            random.randint(0x01, 0xfe),
            random.randint(0x01, 0xfe) ]
    return ':'.join(map(lambda x: "%02x" % x, mac))


# =====================================================================================================================
def get_iface_ip_and_mask(if_name, ipv6=False):
    try:
        import netifaces

        addrs = netifaces.ifaddresses(if_name).get(netifaces.AF_INET6 if ipv6 else netifaces.AF_INET)
        if (addrs is None) or (len(addrs) == 0):
            raise ValueError("Network interface '{}' has no addresses".format(if_name))

        addr, netmask = addrs[0]["addr"], addrs[0]["netmask"]
        dlog("Interface '{}' has {} addresses, use the first: addr={} netmask={}".format(
            if_name, len(addrs), addr, netmask))
        return addr, netmask
    except ValueError:
        raise ValueError("Network interface '{}' not found, available: {}".format(
            if_name, ", ".join(netifaces.interfaces())
        ))
    except ImportError:
        dlog("'netifaces' module not found")
        raise


# =====================================================================================================================
def validate_ip_address(ip_str):  # throw on error
    import socket
    try:
        socket.inet_aton(ip_str)
    except socket.error:
        raise ValueError("Invalid IP address: {}".format(ip_str))


# =====================================================================================================================
def from_hsize(val):  # throw on error
    suffixes = {"B": 1, "K": 1 << 10, "M": 1 << 20, "G": 1 << 30}
    if val[-1].isalpha():
        mul = suffixes.get(val[-1].upper())
        if mul is None:
            raise ValueError("Couldn't parse {}".format(val))
        return mul * int(val[:-1])
    else:
        return int(val)


# =====================================================================================================================
def to_hsize(val):
    for suf in ("", "K", "M", "G"):
        if val == 0 or val % 1024:
            break
        val //= 1024
    return str(val) + suf


# =====================================================================================================================
def aligned_address(alignment, addr):
    blocks = addr // alignment + (1 if addr % alignment else 0)
    return blocks * alignment


# =====================================================================================================================
def copy_to_dir(src_file, dst_dir):
    dst_file = os.path.join(dst_dir, os.path.split(src_file)[1])
    shutil.copyfile(src_file, dst_file)
    return dst_file


# =====================================================================================================================
class TftpContext:
    """ Context manager for TFTP server
    """
    def __enter__(self):
        self.thread.start()
        return self

    def __exit__(self, *args, **kwargs):
        self.server.stop()
        self.thread.join()

    def __init__(self, root_dir, listen_ip, listen_port=69):
        import tftpy
        import threading

        self.server =  tftpy.TftpServer(root_dir)

        def run():
            self.server.listen(listen_ip, listen_port)

        self.thread = threading.Thread(target=run, name="Thread-TftpServer")