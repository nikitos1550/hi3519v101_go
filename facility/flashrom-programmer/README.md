# Arduino based commutable spi flash memory programmer

## Hardware
Device is based on Arduino Uno R3 and is just a bit modified serprog for flashrom.
Shield design are files are located in ./hwdesign dir.

## Uniq id for Arduino and usb serial bridge
It is only possible to buy Arduino Uno R2 with AVR 16u2 usb-serial bridge,
ftdi based chines clones are not in production and version with ch340g doesn`t have eeprom.

We changed 16u2 chip firmware for firmware that located in ./fast-usbserial dir
Originally it is https://github.com/urjaman/fast-usbserial.
This version of usb-serial bridge has several improvements that allow bridge to 
work on 2Mbps speed with flashrom, comparing to stock firmware.

Also we added uniq serial for each device on the facility. 
Grep "ATTENTION" in firmware sources dir to find exact place for uniq serial number.

## Serprog
Sources located in ./frser-duino dir.
Originally https://github.com/urjaman/frser-duino. We added arduino`s digital pin 9
control to switch flash memory from camera to arduino during flashrom sessions.
Grep "RELAY" in sources dir to look how it is exactly working.

Makefile was corrected to set 2Mbit speed of serprog, but upload firmware at 115200bps speed.

## Notes on programming device
### 16u2
* make clean
* make
* make dfu

### serprog
* make clean
* make u2
* make flash-u2

### Host PC
Sometimes there are some brltty and ModemManagers working on your GNU/Linux PC, they will
connect to any serial device to check is it modem or not, this can cause a lot of headache
as the device (arduino) will behave ambiguously (like you can`t upload firmware right after 
plugging to pc and have to wait some time)
