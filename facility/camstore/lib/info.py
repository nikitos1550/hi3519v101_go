from .server import register
from .common import success
import os


@register
def sysinfo():
    """Print server process info
    """
    return success("PID={}".format(os.getpid()))
