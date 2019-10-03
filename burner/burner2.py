#!/usr/bin/env python3
# -*- coding: utf-8 -*-
import logging
import os
import tempfile
import threading
import utils
import tftpy
from uboot_console import UBootConsole


class Defaults:
    dev_baudrate = 115200
    initrd_mem_size = "32M"
    linux_mem_size = "32M"
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
        parser.add_argument("--initrd-size", type=utils.from_hsize, default=Defaults.initrd_mem_size,
            help="Amount of RAM for initrd (default: {})".format(Defaults.initrd_mem_size), metavar="SIZE")
        parser.add_argument("--memory-size", "-m", type=utils.from_hsize, default=Defaults.linux_mem_size,
            help="Amount of RAM for Linux (default: {})".format(Defaults.linux_mem_size), metavar="SIZE")

        parser.set_defaults(action=cls.run)
    
    @staticmethod
    def run(uboot: UBootConsole, args):
        BLOCK_SIZE = 64 * 1024  # 64Kb

        network = NetworkContext(args)
        uboot.setenv(ipaddr=network.target_ip, netmask=network.mask, serverip=network.host_ip)

        with tempfile.TemporaryDirectory() as tempdir:
            logging.info("Pack {} & {} with {} alignment".format(
                args.uimage, args.rootfs, utils.to_hsize(BLOCK_SIZE)
            ))
            pack_file_name = os.path.join(tempdir, "uimage-n-rootfs.pack")
            uimage_offset, rootfs_offset = utils.aligned_pack(BLOCK_SIZE, pack_file_name,
                args.uimage, args.rootfs)
            logging.info("RootFS has {} offset in the pack".format(rootfs_offset))

            with utils.TftpContext(tempdir, listen_ip=network.host_ip, listen_port=69):
                uboot.tftp("0x82000000", pack_file_name)

            bootargs = ""
            bootargs += "mem={} ".format(utils.to_hsize(args.memory_size))
            bootargs += "console=ttyAMA0,115200 "
            bootargs += "ip={}:{}:{}:{}:camera1::off; ".format(
                network.target_ip, network.host_ip, network.host_ip, network.mask)
            bootargs += "mtdparts=hi_sfc:512k(boot) root=/dev/ram0 ro initrd=" \
                + hex(0x82000000 + rootfs_offset)+"," + utils.to_hsize(args.initrd_size)
            
            uboot.setenv(bootargs=bootargs)

            
            print("Wanna boot {} from {} you bastard?".format(args.port, args.uimage))
            uboot.command("bootm 0x82000000")


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


