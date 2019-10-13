#!/usr/bin/env python3
from device import Device
import argparse


PORT = "/dev/ttyACM0"
BAUDRATE = 115200
PROMPT = b"READY"
COLOR_GREEN = '\033[92m'
COLOR_DEFAULT = '\033[0m'


parser = argparse.ArgumentParser(description="Manage devices' power")
parser.add_argument("mode", help="What to do with power (on|off|reset)")
parser.add_argument("num", type=int, help="Target device's number")

args = parser.parse_args()
if args.mode not in ("on", "off", "reset"):
    print("Allowed modes: 'on', 'off', or 'reset'")
    exit(1)

command = args.mode.encode("ascii") + b" " + bytes([0x30 + args.num]) + b"\n"

pm = Device(port=PORT, baudrate=BAUDRATE)
while PROMPT not in pm.read_line():
    pass
pm.write_data(command)

print(COLOR_GREEN + "Power is {}ed on {}".format(args.mode.upper(), args.num) + COLOR_DEFAULT)
