#!/usr/bin/env python3
# -*- coding: utf-8 -*-
import logging
import os
import shutil
import tempfile
import utils
from uboot_console import UBootConsole


# =====================================================================================================================
class UBootContext:
    @classmethod
    def add_args(cls, parser):
        parser.add_argument("--reset-power", required=False, help="Use given command to reset target device", metavar="CMD")
        parser.add_argument("--port", "-p", required=True, help="Serial port device")
        parser.add_argument("--baudrate", type=int, default=Defaults.dev_baudrate,
            help="Serial port baudrate (default: {})".format(Defaults.dev_baudrate))

    @classmethod
    def create(cls, args):
        uboot_logger = utils.get_device_logger("uboot")

        if args.reset_power is not None:
            import subprocess

            logging.info("Execute '{}' to reset power on device".format(args.reset_power))
            subprocess.check_call(args.reset_power, shell=True)

            return UBootConsole.catch_console(
                port=args.port,
                baudrate=args.baudrate,
                logger=uboot_logger
            )
        else:
            return UBootConsole(
                port=args.port,
                baudrate=args.baudrate,
                logger=uboot_logger
            )


# =====================================================================================================================
class Defaults:
    dev_baudrate = 115200
    initrd_mem_size = "8M"
    linux_mem_size = "64"
    uboot_mem_size = "512K"
    linux_console = "ttyAMA0,115200"


# =====================================================================================================================
def upload_files_via_tftp(uboot, listen_ip, listen_port, files_and_addrs):
    with tempfile.TemporaryDirectory() as tmpdir:
        with utils.TftpContext(tmpdir, listen_ip=listen_ip, listen_port=listen_port):
            for filename, addr in files_and_addrs:
                tmp_filename = utils.copy_to_dir(filename, tmpdir)
                logging.info("Upload '{}' via TFTP at {:#x} address".format(filename, addr))
                uboot.tftp(addr, tmp_filename)



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
        parser.add_argument('--lconsole', default=Defaults.linux_console,
            help="Linux load console (default: {})".format(Defaults.linux_console))

        parser.set_defaults(action=cls.run)
    
    @staticmethod
    def run(uboot: UBootConsole, args):
        BLOCK_SIZE = 64 * 1024  # 64Kb
        BASE_ADDR = 0x82000000

        network = NetworkContext(args)
        uboot.setenv(ipaddr=network.target_ip, netmask=network.mask, serverip=network.host_ip)

        uimage_addr = BASE_ADDR
        rootfs_addr = utils.aligned_address(BLOCK_SIZE, uimage_addr + os.path.getsize(args.uimage))
        upload_files_via_tftp(uboot, network.host_ip, 69, files_and_addrs=[(args.uimage, uimage_addr), (args.rootfs, rootfs_addr)])

        bootargs = ""
        bootargs += "mem={} ".format(utils.to_hsize(args.memory_size))
        bootargs += "console={} ".format(args.lconsole)
        bootargs += "ip={}:{}:{}:{}:camera1::off; ".format(
            network.target_ip, network.host_ip, network.host_ip, network.mask)
        bootargs += "mtdparts=hi_sfc:512k(boot) root=/dev/ram0 ro initrd=" \
            + hex(rootfs_addr)+"," + utils.to_hsize(args.initrd_size)
        
        logging.info("Load kernel with bootargs: {}".format(bootargs))
        uboot.setenv(bootargs=bootargs)
        uboot.bootm(uimage_addr)
        logging.info("OS seems successfully started")


# =====================================================================================================================
def main():
    import argparse

    parser = argparse.ArgumentParser(description="Interact with devices via serial port")
    parser.set_defaults(action=lambda *_: parser.print_help())
    
    # common arguments
    parser.add_argument("--log-level", "-l", default="INFO", help="Logging level (default: INFO)", metavar="LVL")
    UBootContext.add_args(parser)

    # each action may add its' own arguments
    action_parsers = parser.add_subparsers(title="Action")
    MacAction.register(action_parsers)
    LoadAction.register(action_parsers)

    args = parser.parse_args()
    
    # configure logging
    log_level = getattr(logging, args.log_level.upper())
    logging.basicConfig(level=log_level)

    uboot = UBootContext.create(args)
    args.action(uboot, args)


if __name__ == "__main__":
    main()


