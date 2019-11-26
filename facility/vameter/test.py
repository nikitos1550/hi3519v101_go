#!/usr/bin/env python

import serial
import binascii

vameter = serial.Serial("/dev/ttyVAMETER", 2400, timeout = 2)

packet="\x0C\x03\x00\x2B\x00\x03\x74\xDE"
#packet = bytearray()
#packet.append(0x01)
#packet.append(0x03)
#packet.append(0x43)
print packet #.decode("hex")

vameter.write(packet)

answer = vameter.readline()
print len(answer)



