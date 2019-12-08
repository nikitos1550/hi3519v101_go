import os
import logging
import asyncio
from .server import routine


class Device:
    def __init__(self, devname):
        self.devname = devname
        # it will also keep some important information


class Devices:
    def __init__(self):
        self.devs = {}

    def _get_current_device_list(self):
        ls = set()
        for f in os.listdir("/dev"):
            if not f.startswith("ttyCAM"):
                continue
            ls.add(os.path.join("/dev", f))
        return ls

    def update(self):
        logging.debug("Update devices...")
        current = self._get_current_device_list()

        changed = False

        # remove
        for devname in self.devs.keys():
            if devname not in current:
                self.devs.pop(devname)
                logging.info("Device {} is removed".format(devname))
                changed = True

        # add
        for devname in current:
            if devname not in self.devs:
                self.devs[devname] = Device(devname)
                logging.info("Device {} is added".format(devname))
                changed = True

        return changed
