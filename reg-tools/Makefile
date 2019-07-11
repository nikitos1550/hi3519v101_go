THIS_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))


all:
	cd reg-tools-1.0.0; make OSDRV_CROSS=$(THIS_DIR)../buildroot-2018.11.2/output/host/bin/arm-buildroot-linux-uclibcgnueabi

clean:
	cd reg-tools-1.0.0; make clean
