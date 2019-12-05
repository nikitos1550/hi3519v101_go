from server import register, routine
import os
import logging
import asyncio


@register
def sysinfo():
    """Print server process info
    """
    return "PID: {}".format(os.getpid())
