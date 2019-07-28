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

DATA_PORT 		= "/dev/ttyCAM1"
POWER_PORT 		= "/dev/ttyACM0"

SPEED			= 115200

promt 			= ["xmtech #", "hisilicon #"]
duplex 			= "full"

iface           = "enx503eaa7b65cb"

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
		data.write(cmd.encode())

def setvar(key, value):
	cmd = "setenv " + key + " " + value
	datacmd(cmd)
	waitcmd()

def setvararray(kvs):
	cmd = ""
	for key in kvs:
        	cmd = cmd + "setenv " + key + " " + kvs[key]+"; "
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


def validate_ip_address(ip_str):  # throw on error
    import socket
    try:
        socket.inet_aton(ip_str)
    except socket.error:
        raise ValueError("Invalid IP address: {}".format(ip_str))


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


def get_iface_ip_and_mask(iface):
    try:
        addrs = netifaces.ifaddresses(iface).get(netifaces.AF_INET)
        if (addrs is None) or (len(addrs) == 0):
            raise ValueError("iface has no addresses")

        addr = addrs[0]["addr"]
        netmask = addrs[0]["netmask"]
        return addr, netmask
    except StandardError as err:
        raise ValueError(
            "Network interface '{}' address and mask wasn't defined: {}\nAvailable interfaces: {}".format(
                iface, err.message, ", ".join(netifaces.interfaces())
            )
        )


###############################################################################


parser = argparse.ArgumentParser(description='Load/burn firmware to hisilicon ip device.')

parser.add_argument('action',   type=str,   help='Required action load|burn|uboot|mac')

parser.add_argument('--ip', 	    type=str,	help='Ip address for device', required=True)
#parser.add_argument('--mask',     type=str,   help='Owerrite network mask for device')
#parser.add_argument('--server',   type=str,   help='Server ip')

parser.add_argument('--skip', type=int, default=512, help='Uboot size to skip in kb (default 512kb)', required=False)
parser.add_argument('--uboot', 	type=str, 	help='Uboot image file name',                       required=False)
parser.add_argument('--uimage', type=str, 	help='Kernel uImage file name',                     required=False)
parser.add_argument('--rootfs', type=str, 	help='RootFS file name',                            required=False)
parser.add_argument('--memory', type=int,   help='Amount of RAM for Linux in MB',               required=False)

parser.add_argument('--initrd', type=int,   help='Amount of RAM for Linux Initrd (only for load action)', required=False)
parser.add_argument('--lserial', type=int,   help='Linux load serial config', required=False)

parser.add_argument('--port', help='Serial port dev name', required=False)
parser.add_argument('--speed', type=int,   help='Serial port speed',               required=False)

parser.add_argument('--iface', default=iface, help='Network interface name (default: {})'.format(iface), required=False)

parser.add_argument('--duplex', type=int,   help='Full (default) or Half duplex mode',               required=False)
parser.add_argument('--mc21', type=int,   help='To remove',               required=False)

parser.add_argument('--force', action='store_true', help='Skip usefull checks',               required=False)

parser.add_argument('--servercamera', type=int, help='Camera number',               required=False)


args = parser.parse_args()

print args

validate_ip_address(args.ip)
ip = args.ip
server_ip, mask = get_iface_ip_and_mask(args.iface)

print("Iface " + args.iface)
print("- Server IP: {}\n- Mask: {}\n- Target device IP: {}".format(server_ip, mask, ip))

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

##PREPARE cat uImage+rootfs
uimage_blocks = uimage_size/(64*1024)
print "uImage full 64K blocks " + str(uimage_blocks)
if uimage_size - (uimage_blocks*64*1024) > 0:
	print "uImage align difference " + str( ( (uimage_blocks+1)*64*1024 ) - uimage_size) + " bytes"
        uimage_size_e = ((uimage_blocks+1)*64*1024)
os.system("rm ./images/difference; rm ./images/tmp.tftp")
os.system("dd if=/dev/zero of=./images/difference conv=sync bs=1 count="+str(( (uimage_blocks+1)*64*1024 ) - uimage_size ))
os.system("cat "+args.uimage+" ./images/difference "+args.rootfs+" > ./images/tmp.tftp")
os.system("chmod 777 ./images/difference; chmod 777 ./images/tmp.tftp;")

uimage_rootfs_size = uimage_size_e + rootfs_size

###########################

data = serial.Serial(DATA_PORT,      SPEED, timeout = 0.5)
#data.flush()
#data.flushInput()
#data.flushOutput()
#power = serial.Serial(POWER_PORT,      SPEED, timeout = 1)
#time.sleep(3)

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

    if args.servercamera != None:
	power = serial.Serial(POWER_PORT,      SPEED, timeout = 0.1)
	time.sleep(3)
	print "Server camera "+str(args.servercamera)+" setted, auto power reset"
    	power.write("reset "+str(args.servercamera)+"\n")
	#power.close()
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
    
    if args.servercamera != None:
	power.close()

    #if answer.find("Err:   serial") != -1:#if answer.find("U-Boot 2010.06 (May 11 2018 - 15:06:27)") != -1:
    #    print "-->Pressing Ctrl+C"
    #    data.write("\x03")
    #    answer = data.readline().strip()
    #    print "DATA: " + answer
    #    break

    datawrite("\x03") #data.write("\x03") #just another ctrl+c 
    waitcmd()
    print "-->got cmd"

#printenv()

if args.action in ["mac"]:
    setvar("ethaddr", mac)
    printenv()
    saveenv()
    exit("Mac " + mac + " setuped")

setvar("ipaddr", 	ip)
setvar("netmask", 	mask)
setvar("serverip", 	server_ip)

#netdata = {}
#netdata["ipaddr"] = ip
#netdata["netmask"] = mask
#netdata["serverip"] = server_ip
#setvararray(netdata)

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
    #server = tftpy.TftpServer(".")
    thread.start_new_thread( threadtftd, () )
    print "-->tftpd started"
    if args.action == "burn":
	    size_wroute = tftpuploadspi(uimage_filepath[0], uimage_size, uboot_size)
    if args.action == "load":
	    #tftpupload(uimage_filepath[0], uimage_size, "kernel")
	    tftpupload("tmp.tftp", uimage_rootfs_size, "kernel")
    server.stop(now=True)
    #time.sleep(10)

    #server = tftpy.TftpServer(rootfs_filepath[1])
    #thread.start_new_thread( threadtftd, () )
    #print "-->tftpd started"
    #if args.action == "burn":
    #    tftpuploadspi(rootfs_filepath[0], rootfs_size, uboot_size + size_wroute)
    #if args.action == "load":
    #    tftpupload(rootfs_filepath[0], rootfs_size, "rootfs")
    #server.stop(True)
    #time.sleep(10)
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
    bootargs = bootargs + "ip="+str(ip)+":"+str(server_ip)+":"+str(server_ip)+":"+str(mask)+":camera"+str(args.servercamera)+"::off; "
    bootargs = bootargs + "mtdparts=hi_sfc:512k(boot) root=/dev/ram rw initrd="+hex(0x82000000 + uimage_size_e)+"," + initrd

    #if rootfs == "nfs":
	#    bootargs = bootargs + "root=/dev/nfs rw nfsrootdebug nfsroot=" + args.nfs + ",v3"
	#    bootargs = bootargs + " ip="+args.ip+":"+args.server+":"+args.server+":"+args.mask+":"+"hi"+args.hisi+":eth0:off"

#nfsroot=192.168.250.1:/opt/ipcam/rootfs-ipcam2-5mp,proto=tcp ip=192.168.250.20:192.168.250.1:192.168.250.1:255.255.255.0:hi3516a:eth0:off


    setvar("bootargs", bootargs)

#if args.version != None:
#	setvar("version", args.version)

    #printenv()

    if args.action == "burn":
	    saveenv()
	    datacmd("reset")
    if args.action == "load":
	    datacmd("bootm 0x82000000")

    #while True:
    #    answer = data.readline().strip()
    #    print "DATA: " + answer

    data.close()

    exit()

################################################################################
