import random, logging, sys


def __init_logger(
    name,
    level=logging.INFO,
    fmt="[%(name)s:%(levelname)s, %(asctime)s.%(msecs)d] %(message)s",
    datefmt="%H:%M:%S",
    stream=sys.stdout
):
    handler = logging.StreamHandler(stream)
    handler.setFormatter(logging.Formatter(fmt=fmt, datefmt=datefmt))
    logger = logging.getLogger(name)
    logger.propagate = False
    logger.setLevel(level)
    logger.addHandler(handler)

    return logger


__conn_logger = __init_logger("conn", fmt="[%(name)s %(asctime)s.%(msecs)d] %(message)s")
__dev_logger = __init_logger("dev", level=logging.DEBUG, fmt="[%(name)s:%(levelname)s][%(filename)s:%(lineno)d] - %(message)s")


def DEV_WARN(msg):
    __dev_logger.warning(msg)

def DEV_INFO(msg):
    __dev_logger.info(msg)

def DEV_DBG(msg):
    __dev_logger.debug(msg)


def CONN_INCOMING(msg):
    __conn_logger.info("<- " + msg)

def CONN_OUTGOING(msg):
    __conn_logger.info("-> " + msg)


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
        DEV_DBG("Interface '{}' has {} addresses, use the first: addr={} netmask={}".format(
            if_name, len(addrs), addr, netmask))
        return addr, netmask
    except ValueError:
        raise ValueError("Network interface '{}' not found, available: {}".format(
            if_name, ", ".join(netifaces.interfaces())
        ))
    except ImportError:
        DEV_WARN("'netifaces' module not found")
        raise
