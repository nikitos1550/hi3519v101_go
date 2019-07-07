#!/usr/bin/env python

import serial
import time
import argparse
import os
import tftpy
import thread
import random

import netifaces
#from netifaces import interfaces, ifaddresses, AF_INET

###############################################################################

DATA_PORT 		= "/dev/ttyUSB0"
#DATA_PORT 		= "/dev/ttyACM0"

SPEED			= 115200

promt 			= ["xmtech #", "hisilicon #"]
duplex 			= "full"

iface           = "enp3s0"

###############################################################################
def datawrite(send):
    time.sleep(.1)
    for item in send:
        #print item
        data.write(item)
        time.sleep(.1)

def checkcmd(str):
	#print str
	if str in promt: return 1
        return 0

def waitcmd():
        while True:
                answer = data.readline().strip()
                if checkcmd(answer) == 1: break
                print "DATA: " + answer	

def datacmd(cmd):
	print "-->" + cmd
	cmd = cmd.replace(";", "\;") + "\n"
	if duplex == "half":
		datawrite(cmd)
	else:
		data.write(cmd)

def setvar(key, value):
	cmd = "setenv " + key + " " + value
	datacmd(cmd)
	waitcmd()

def printenv():
	env = {}
	datacmd("printenv")
	while True:
        	answer = data.readline().strip()
		if checkcmd(answer) == 1: break
        	print "DATA: " + answer
		parsed = answer.split("=")
		if len(parsed) == 2:
			env[parsed[0]] = parsed[1]
		elif len(parsed) > 2:
			tmp = ""
			for i in range(1, len(parsed)):
				tmp = tmp + "=" + parsed[i]
			env[parsed[0]] = tmp[1:]
	print env

def cleanenv():
        env = {}
        datacmd("printenv")
        while True:
                answer = data.readline().strip()
                if checkcmd(answer) == 1: break
                #print "DATA: " + answer
                parsed = answer.split("=")
                if len(parsed) == 2:
                        env[parsed[0]] = parsed[1]
                elif len(parsed) > 2:
                        tmp = ""
                        for i in range(1, len(parsed)):
                                tmp = tmp + "=" + parsed[i]
                        env[parsed[0]] = tmp[1:]
        #print env
        keep_params = ["mdio_intf", "baudrate", "ethaddr", "netmask", "ipaddr", "serverip", "bootcmd", "bootargs", "stdin", "stdout", "stderr"]
        for env_one in env:
                #print env_one + "->" + env[env_one]
                if env_one not in keep_params:
                        #print "setenv " + str(env_one)
                        datacmd("setenv " + str(env_one))
			answer = data.readline().strip()
			print "DATA: " + answer
			answer = data.readline().strip()
                        print "DATA: " + answer

def tftpupload(file, size, pointer):
	if pointer == "kernel":
		datacmd("tftp 0x82000000 " + file)
	if pointer == "rootfs":
		datacmd("tftp 0x82400000 " + file)
	waitcmd()
	return 1

def tftpuploadspi(file, size, offset):
	print "-->burning " + file + " file"
	blocks = size/(64*1024)
	print "-->full 64K blocks " + str(blocks)
	if size - (blocks*64*1024) > 0:
		print "-->align difference " + str(size - (blocks*64*1024)) + " bytes"
		size_e = (blocks+1)*64*1024
		print "-->erase size " + hex(size_e) + " " + str(size_e/(64*1024)) + " blocks"
	else:
		size_e = size

	datacmd("tftp 0x82000000 " + file)
	waitcmd()
	datacmd("sf probe 0")
	waitcmd()
	datacmd("sf erase " + hex(offset) + " " + hex(size_e))
	waitcmd()
	datacmd("sf write 0x82000000 " + hex(offset) + " " + hex(size))
	waitcmd()
	return size_e

def saveenv():
	cmd = "saveenv"
        datacmd(cmd)
        waitcmd()

def filepath(file):
	answer = ["", ""]
	file_path = file.split("/")
	answer[0] = file_path[len(file_path) - 1]
	if len(file_path) == 1:
        	file_path = "."
	else:
        	if file_path[0] == ".":
                	tmp = ""
        	else:
                	tmp = "./"
        	for i in range(0, len(file_path)-1):
                	tmp = tmp + file_path[i] + "/"
        	file_path = tmp
	answer[1] = file_path
	return answer

def ipFormatChk(ip_str):
    if len(ip_str.split()) == 1:
        ipList = ip_str.split('.')
        if len(ipList) == 4:
            for i, item in enumerate(ipList):
                try:
                    ipList[i] = int(item)
                except:
                    return False
                if not isinstance(ipList[i], int):
                    return False
            if max(ipList) < 256:
                return True
            else:
                return False
        else:
            return False
    else:
        return False

def threadtftd():
	server.listen(server_ip, 69)


# The first line is defined for specified vendor
def randommac():
    #00:00:23:34:45:66
    mac = [ 0x00, 0x00, 0x23,
            random.randint(0x01, 0xfe),
            random.randint(0x01, 0xfe),
            random.randint(0x01, 0xfe) ]
    return ':'.join(map(lambda x: "%02x" % x, mac))

###############################################################################


parser = argparse.ArgumentParser(description='Load/burn firmware to hisilicon ip device.')

parser.add_argument('action',   type=str,   help='Required action load|burn|uboot|mac')

parser.add_argument('--ip', 	    type=str,	help='Ip address for device', required=True)
#parser.add_argument('--mask',     type=str,   help='Owerrite network mask for device')
#parser.add_argument('--server',   type=str,   help='Server ip')

parser.add_argument('--skip',   type=int,   help='Uboot size to skip in kb (default 512kb)',    required=False)
parser.add_argument('--uboot', 	type=str, 	help='Uboot image file name',                       required=False)
parser.add_argument('--uimage', type=str, 	help='Kernel uImage file name',                     required=False)
parser.add_argument('--rootfs', type=str, 	help='RootFS file name',                            required=False)
parser.add_argument('--memory', type=int,   help='Amount of RAM for Linux in MB',               required=False)

parser.add_argument('--initrd', type=int,   help='Amount of RAM for Linux Initrd (only for load action)', required=False)
parser.add_argument('--lserial', type=int,   help='Linux load serial config', required=False)

parser.add_argument('--port', type=int,   help='Serial port dev name',               required=False)
parser.add_argument('--speed', type=int,   help='Serial port speed',               required=False)

parser.add_argument('--iface', type=int,   help='Network interface name',               required=False)

parser.add_argument('--duplex', type=int,   help='Full (default) or Half duplex mode',               required=False)
parser.add_argument('--mc21', type=int,   help='To remove',               required=False)

parser.add_argument('--force', action='store_true', help='Skip usefull checks',               required=False)

args = parser.parse_args()

print args

if args.iface != None:
    iface = args.iface
    print "Iface is not set, default iface " + iface
else:
    print "Iface " + iface

for ifaceName in netifaces.interfaces():
    if ifaceName == iface:
        #addresses = [i['addr'] for i in ifaddresses(ifaceName).setdefault(AF_INET, [{'addr':'No IP addr'}] )]
        #print '%s: %s' % (ifaceName, ', '.join(addresses))
        addrs = netifaces.ifaddresses(ifaceName).get(netifaces.AF_INET)
        server_ips = [addr.get("addr") for addr in addrs]
        server_mask = [addr.get("netmask") for addr in addrs]
        server_ip = str(server_ips[0])
        mask = str(server_mask[0])
        print "Server ip " + server_ip + " mask " + mask

if ipFormatChk(args.ip) 	    != True: 
    exit("Ip is not valid")
else:
    ip = args.ip
    
#if ipFormatChk(args.mask) 	    != True: exit("Network mask is not valid")
#if ipFormatChk(args.server) 	!= True: exit("Server ip is not valid")

###

print "Target device ip " + ip

###

if args.port != None:
	DATA_PORT = args.port
if args.speed != None:
	SPEED = args.speed

print "Serial port " + str(DATA_PORT) + " speed " + str(SPEED)

###

if args.lserial != None:
    lserial = args.lserial
else:
    lserial = "ttyAMA0,115200"

print "Linux console \"" + lserial + "\""

###

if args.initrd != None:
    initrd = str(int(args.initrd)) + "M"
else:
    initrd = "32M"
    
if args.action in ["load"]:
    print "Initrd size " + initrd

###

if args.action not in ["load", "burn", "uboot", "mac"]: exit("Action should be load|burn|uboot")

if args.action in ["mac"]:
    mac = randommac()
    print str(mac) + " mac address will be setuped"

if args.action in ["load", "burn"]:
    if args.skip:
	    uboot_size = args.skip*1024
    else:
        #uboot_size = 512*1024 #default uboot skip size is 512kb
        exit("Please specify uboot skip size!")

    print "Uboot size is " + str(uboot_size/1024) + " Kbytes, this will be skipped"

    if args.uimage == None: exit("Please specify uImage file")
    if os.path.isfile(args.uimage) == False: exit("uImage file doesn`t exist")
    uimage_size = os.path.getsize(args.uimage)
    uimage_filepath = filepath(args.uimage)
    print "uImage file (" + uimage_filepath[1] + uimage_filepath[0] + ") size is " + str(uimage_size) + " bytes"

    if args.rootfs == None: exit("Please specify rootfs file")   
    if os.path.isfile(args.rootfs) == False: exit("RootFS file doesn`t exist")
    rootfs_size = os.path.getsize(args.rootfs)
    rootfs_filepath = filepath(args.rootfs)
    print "rootfs file (" + rootfs_filepath[1] + rootfs_filepath[0] + ") size is " + str(rootfs_size) + " bytes"
        
if args.action in ["uboot"]:
    if args.uboot == None: exit("Please specify uboot file")
    if os.path.isfile(args.uboot) == False: exit("uboot file doesn`t exist")
    uboot_size = os.path.getsize(args.uboot)
    uboot_filepath = filepath(args.uboot)
    print "uImage file (" + uboot_filepath[1] + uboot_filepath[0] + ") size is " + str(uboot_size) + " bytes"

    if uboot_size > 512*1024 and args.force != True:
        exit("Seems your uboot image is too big, correct burner.py if you know what you are doing")

data = serial.Serial(DATA_PORT,      SPEED, timeout = 1)
time.sleep(1)

if args.mc21 != None:
    datawrite("%%%camterm\r\n") #data.write("%%%camterm\r\n")
    while True:
        answer  = data.readline().strip()
        print "DATA: " + answer
        if answer.find(">") != -1: break
    #print answer
    #exit("TEST")
    datawrite("hisiuboot\r\n") #data.write("hisiuboot\r\n")
    while True:
        answer = data.readline().strip()
        got = 0
        for promt_item in promt:
            if answer.find(promt_item) != -1:
                got = 1
                break
        if got: break
        print "DATA: " + answer
    datawrite("\x03") #data.write("\x03") #just another ctrl+c 
    waitcmd()
    print "-->got cmd"

else:
    print "Please plug power to module"

    while True:
        answer = data.readline().strip()
        got = 0
        for promt_item in promt:
            #print promt_item
            if answer.find(promt_item) != -1:
                got = 1
                break
        if got: break
        print "DATA: " + answer
        datawrite("\x03") #data.write("\x03")
    
    #if answer.find("Err:   serial") != -1:#if answer.find("U-Boot 2010.06 (May 11 2018 - 15:06:27)") != -1:
    #    print "-->Pressing Ctrl+C"
    #    data.write("\x03")
    #    answer = data.readline().strip()
    #    print "DATA: " + answer
    #    break

    datawrite("\x03") #data.write("\x03") #just another ctrl+c 
    waitcmd()
    print "-->got cmd"

printenv()

if args.action in ["mac"]:
    setvar("ethaddr", mac)
    printenv()
    saveenv()
    exit("Mac " + mac + " setuped")

setvar("ipaddr", 	ip)
setvar("netmask", 	mask)
setvar("serverip", 	server_ip)

#datacmd("mm 0x120100cc 0x4a")

if args.action in ["uboot"]:
    server = tftpy.TftpServer(uboot_filepath[1])
    thread.start_new_thread( threadtftd, () )
    print "-->tftpd started"
    size_wroute = tftpuploadspi(uboot_filepath[0], uboot_size, 0)
    server.stop(now=True)
    time.sleep(5)
    #setvar("bootdelay", "3")
    #saveenv()
    #datacmd("reset")
    data.close()
    exit("Uboot burn completed")

if args.action in ["load", "burn"]:
    server = tftpy.TftpServer(uimage_filepath[1])
    thread.start_new_thread( threadtftd, () )
    print "-->tftpd started"
    if args.action == "burn":
	    size_wroute = tftpuploadspi(uimage_filepath[0], uimage_size, uboot_size)
    if args.action == "load":
	    tftpupload(uimage_filepath[0], uimage_size, "kernel")
    server.stop(now=True)
    time.sleep(10)

    server = tftpy.TftpServer(rootfs_filepath[1])
    thread.start_new_thread( threadtftd, () )
    print "-->tftpd started"
    if args.action == "burn":
        tftpuploadspi(rootfs_filepath[0], rootfs_size, uboot_size + size_wroute)
    if args.action == "load":
        tftpupload(rootfs_filepath[0], rootfs_size, "rootfs")
    server.stop(True)
    time.sleep(10)
	#print "-->tftpd stoped"


    #cleanenv()

    bootcmd = ""

    if args.action == "burn":
        bootcmd = bootcmd + "sf probe 0;"
        bootcmd = bootcmd + "sf read 0x82000000 " + hex(uboot_size) + " " + hex(uimage_size) 
        bootcmd = bootcmd + "; sf read 0x82400000 " + hex(uboot_size + size_wroute) + " " + hex(rootfs_size)
        bootcmd = bootcmd + "; bootm 0x82000000"
        setvar("bootcmd", bootcmd)

    bootargs = ""
    bootargs = "mem=" + str(args.memory) + "M "

    bootargs = bootargs + "console=" + lserial + " "
    #if rootfs == "file":
    bootargs = bootargs + "root=/dev/ram rw initrd=0x82400000," + initrd

    #if rootfs == "nfs":
	#    bootargs = bootargs + "root=/dev/nfs rw nfsrootdebug nfsroot=" + args.nfs + ",v3"
	#    bootargs = bootargs + " ip="+args.ip+":"+args.server+":"+args.server+":"+args.mask+":"+"hi"+args.hisi+":eth0:off"

#nfsroot=192.168.250.1:/opt/ipcam/rootfs-ipcam2-5mp,proto=tcp ip=192.168.250.20:192.168.250.1:192.168.250.1:255.255.255.0:hi3516a:eth0:off


    setvar("bootargs", bootargs)

#if args.version != None:
#	setvar("version", args.version)

    printenv()

    if args.action == "burn":
	    saveenv()
	    datacmd("reset")
    if args.action == "load":
	    datacmd("bootm 0x82000000")

    while True:
        answer = data.readline().strip()
        print "DATA: " + answer

    data.close()

    exit()

################################################################################
