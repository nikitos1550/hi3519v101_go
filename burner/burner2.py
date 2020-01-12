#!/usr/bin/env python3
# -*- coding: utf-8 -*-
import logging
import os, sys
import shutil
import tempfile
import utils
from kv_read import kv_read
from uboot_console import UBootConsole, UBootConsoleParams
from device import SerialConn, TelnetConn


def create_camstore_client():
    from utils import FACILITY_DIR
    sys.path.insert(0, FACILITY_DIR)

    from camstore.lib.client import Client

    return Client("localhost", 43500)


# =====================================================================================================================
class Defaults:
    dev_baudrate = 115200
    initrd_mem_size = "8M"
    linux_mem_size = "64"
    uboot_mem_size = "512K"
    linux_console = "ttyAMA0,115200"


# =====================================================================================================================
class UBootContext:
    @classmethod
    def add_args(cls, parser):
        parser.add_argument("--mode", choices=["raw", "camstore"], default="raw")
        parser.add_argument("--reset-power", required=False, help="Use given command to reset target device", metavar="CMD")
        parser.add_argument("--port", "-p", required=True, help="Serial port device")
        parser.add_argument("--baudrate", type=int, default=Defaults.dev_baudrate,
            help="Serial port baudrate for raw mode (default: {})".format(Defaults.dev_baudrate))
        parser.add_argument("--uboot-params", type=kv_read, help="U-Boot console's parameters")

    def __init__(self, args):
        logging.info("U-Boot params: {}".format(args.uboot_params))

        if args.mode == "camstore":
            self._client = create_camstore_client()
            host, port = self._client.forward_serial(args.port)
            conn = TelnetConn(port=port)
        else:
            conn = SerialConn(port=args.port, baudrate=args.baudrate)

        self.uboot = UBootConsole(
            conn=conn,
            logger=utils.get_device_logger("uboot", args.log_level),
            params=UBootConsoleParams(args.uboot_params if args.uboot_params else None)
        )

        if args.reset_power is not None:
            import subprocess
            
            logging.info("Execute '{}' to reset power of device".format(args.reset_power))
            subprocess.check_call(args.reset_power, shell=True)
            self.uboot.fetch_console()


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
def upload_files_via_tftp(uboot, listen_ip, listen_port, files_and_addrs):
    with tempfile.TemporaryDirectory() as tmpdir:
        with utils.TftpContext(tmpdir, listen_ip=listen_ip, listen_port=listen_port):
            for filename, addr in files_and_addrs:
                logging.info("Upload '{}' via TFTP at {:#x} address".format(filename, addr))
                tmp_filename = utils.copy_to_dir(filename, tmpdir)
                uboot.tftp(addr, tmp_filename)


# =====================================================================================================================
class MemProbeAction:
    @classmethod
    def register(cls, aps):
        parser = aps.add_parser("mprobe", help="Probe memroy")
        parser.set_defaults(action=cls.run)
    
    @staticmethod
    def run(uboot, args):
        BASE_ADDR = 0x80000000
        STEP = 1024 * 1024

        uboot.command("mw {:#x} babacaca".format(BASE_ADDR))

        offset = STEP * 400
        while True:
            addr = BASE_ADDR + offset
            uboot.command("mw {:#x} deadbeef".format(addr))

            res1 = uboot.command("md {:#x}".format(addr))
            res2 = uboot.command("md {:#x}".format(BASE_ADDR))
            print("Mem at {} is: {}".format(utils.to_hsize(offset), res1[0]))
            print("Mem at 0 is:  {}\n----------".format(" ".join(res2[0:2])))

            uboot.command("mw {:#x} babacaca".format(addr))
            offset += STEP * 8


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

        print("Sorry, don't work for now")


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
class PrintEnvAction:
    @classmethod
    def register(cls, aps):
        parser = aps.add_parser("printenv", help="Print U-Boot's environment variables")
        parser.set_defaults(action=cls.run)
    
    @staticmethod
    def run(uboot: UBootConsole, args):
        env = uboot.command("printenv")
        print("U-Boot environment:\n" + "\n".join(" - " + l for l in env if l))


# =====================================================================================================================
def main():
    import argparse

    def log_level(val):
        return getattr(logging, val.upper())

    parser = argparse.ArgumentParser(description="Interact with devices via serial port")
    
    # common arguments
    parser.add_argument("--log-level", "-l", type=log_level, default="INFO",
        help="Logging level (default: INFO)", metavar="LVL")
    UBootContext.add_args(parser)

    # each action may add its' own arguments
    action_parsers = parser.add_subparsers(title="Action")
    PrintEnvAction.register(action_parsers)
    MacAction.register(action_parsers)
    LoadAction.register(action_parsers)
    MemProbeAction.register(action_parsers)

    args = parser.parse_args()

    if getattr(args, "action", None) is None:
        print("Action must be specified")
        return

    logging.basicConfig(level=args.log_level)

    uboot = UBootContext(args).uboot
    args.action(uboot, args)


if __name__ == "__main__":
    main()
