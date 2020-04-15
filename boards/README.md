## TODO!!! Env build for custom board

Naming: {VENDOR}_{MODEL}_{CHIP}_{CMOS}

1. Build kernel with custom config and custom dts
2. Use custom toolchain

## Config includes

* kernel patch
* dts
* uboot env config

## config params example
```
FAMILY      =hi3516cv300
CHIP        =hi3516ev100
VENDOR      =XM
MODEL       =53H20S
RAM_SIZE    =32
RAM_LINUX   =32
RAM_MPP     =0
CMOS        =unknown
ROM_SIZE    =16
```

**TODO** gpio, audio, pinout, usb, etc

## build structure

* boards->{board_name}->config
* boards->{board_name}->putontrootfs
* **TODO** kernel config patch
* **TODO** kernel dts 


* boards->{board_name}->build->{build_name}->uImage
* boards->{board_name}->build->{build_name}->rootfs
* boards->{board_name}->build->{build_name}->rom_layout
