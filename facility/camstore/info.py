from server import register, routine
import os
import logging
import asyncio


@register
def sysinfo():
    """Print server process info
    """
    return "PID: {}".format(os.getpid())


@register
def list_cameras():
    """List available camera devices
    """
    return "it'll be camera list here"


@register
def abrakadabra(a, b, c):
    """Just abrakadabra command

    :param a:
    :param b:
    :param c:
    :return:
    """
    return "Abrakadabra: {} {} {}".format(a, b, c)
