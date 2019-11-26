#!/usr/bin/env python

import serial
import argparse
import time

parser = argparse.ArgumentParser(description='Load/burn firmware to hisilicon ip device.')

parser.add_argument('action',   type=str,   help='Required action on|off|reset')
parser.add_argument('cam',   type=int,   help='Required camera number')

args = parser.parse_args()

print args

if (args.action not in ["on", "off", "reset"]):
	exit("action unknown")

if (args.cam < 0 or args.cam > 9):
	exit("camera device number of range")

data = serial.Serial("/dev/ttyACM0", 115200, timeout = 1)
time.sleep(3)
print(data.readline())

if args.action == "reset":
	data.write("reset "+str(args.cam)+"\n")
