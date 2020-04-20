THIS_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))


build-hi3516cv300:
	cd hi3516cv300; make OSDRV_CROSS=$(THIS_DIR)/../../output/hisilicon/hi3516cv300/toolchain/bin/arm-buildroot-linux-uclibcgnueabi

clean:
	cd hi3516cv300; make clean
