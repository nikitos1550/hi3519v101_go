#!/usr/bin/env python

import serial
import binascii

relay0 = serial.Serial("/dev/ttyRELAY0", 115200, timeout = 2)

#packet="\x01\x03\x00\x26\x00\x01\x65\xc1"

#packet="\x01\x10\x00\x00\x00\x02\x04\x00\x82\x00\xFF\x13\xc7"

packet="\x01\x10\x00\x00\x00\x02\x41\xc8"

print packet

relay0.write(packet)

answer = relay0.readline()
print len(answer)
for i in answer:
    print hex(ord(i))


