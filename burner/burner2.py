#!/usr/bin/env python3
# -*- coding: utf-8 -*-
import utils
import logging
from uboot_console import UBootConsole


class Defaults:
    dev_baudrate = 115200
    initrd_mem_size = "32M"
    uboot_mem_size = "512K"




# =====================================================================================================================
class NetworkContext:
    @classmethod
    def add_args(cls, parser):
        parser.add_argument("--target-ip", help="Ip address for device")
        parser.add_argument("--iface",  help="Network interface name")

    def __init__(self, args):
        utils.validate_ip_address(args.target_ip)
        self.target_ip = args.target_ip
        self.host_ip, self.mask = utils.get_iface_ip_and_mask(args.iface)

        utils.dlog("Network: iface={} mask={} host_addr={} target_addr={}",
            args.iface, self.mask, self.host_ip, self.target_ip)


# =====================================================================================================================
class MacAction:
    @classmethod
    def register(cls, aps):
        parser = aps.add_parser("mac", help="Change MAC address")

        parser.add_argument("--val", help="MAC address to be set (random by default)")
        parser.set_defaults(action=cls.run)
    
    @staticmethod
    def run(uboot, args):
        mac = args.val or utils.random_mac()
        uboot.dlog("set MAC={}", mac)

        uboot.command("setenv ethaddr {}".format(mac))

        env = uboot.command("printenv")
        print("U-Boot environment:\n" + "\n".join(" - " + l for l in env if l))

        print("Wanna change {} MAC on {} you bastard?".format(args.port, mac))


# =====================================================================================================================
class LoadAction:
    @classmethod
    def register(cls, aps):
        parser = aps.add_parser("load", help="Load image onto device and boot it (without burning)")

        NetworkContext.add_args(parser)
        parser.add_argument("--uimage", required=True, help="Kernel's uImage file")
        parser.add_argument("--rootfs", required=True, help="RootFS file")
        parser.add_argument("--skip", type=utils.from_hsize, default=Defaults.uboot_mem_size,
            help="U-Boot size to skip (default: {})".format(Defaults.uboot_mem_size), metavar="SIZE")
        parser.add_argument("--initrd-size", type=utils.from_hsize, default=Defaults.initrd_mem_size,
            help="Amount of RAM for initrd (default: {})".format(Defaults.initrd_mem_size), metavar="SIZE")

        parser.set_defaults(action=cls.run)
    
    @staticmethod
    def run(uboot, args):
        network = NetworkContext(args)

        logging.info("U-Boot size is {}; it'll be skipped".format(utils.to_hsize(args.skip)))

        print("Wanna boot {} from {} you bastard?".format(args.port, args.uimage))


# =====================================================================================================================
def main():
    import argparse


    parser = argparse.ArgumentParser(description="Interact with devices via serial port")
    parser.set_defaults(action=lambda *_: parser.print_help())
    
    # common arguments
    parser.add_argument("--log-level", "-l", default="INFO", help="Logging level (default: INFO)", metavar="LVL")
    parser.add_argument("--reset-power", required=False, help="Use given command to reset target device", metavar="CMD")
    parser.add_argument("--port", "-p", required=True, help="Serial port device")
    parser.add_argument("--baudrate", type=int, default=Defaults.dev_baudrate,
        help="Serial port baudrate (default: {})".format(Defaults.dev_baudrate))

    # each action may add its' own arguments
    action_parsers = parser.add_subparsers(title="Action")
    MacAction.register(action_parsers)
    LoadAction.register(action_parsers)
    
    args = parser.parse_args()
    
    # configure logging
    log_level = getattr(logging, args.log_level.upper())
    logging.basicConfig(level=log_level)
    logger = utils.get_device_logger("uboot", level=log_level)

    # construct U-Boot's console wrap
    if args.reset_power is not None:
        import subprocess
        logging.info("Execute '{}' to reset power on device".format(args.reset_power))
        subprocess.check_call(args.reset_power, shell=True)
        uboot = UBootConsole.catch_console(
            port=args.port,
            baudrate=args.baudrate,
            logger=utils.get_device_logger("uboot"))
    else:
        uboot = UBootConsole(
            port=args.port,
            baudrate=args.baudrate,
            logger=utils.get_device_logger("uboot"))

    # run action
    args.action(uboot, args)


if __name__ == "__main__":
    main()


