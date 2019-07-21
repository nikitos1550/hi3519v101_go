THIS_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))


TOOLCHAIN=$(THIS_DIR)../buildroot-2019.05.1-debug/output/host/bin/arm-buildroot-linux-uclibcgnueabi-
#TOOLCHAIN=/opt/hisi-linux/x86-arm/arm-hisiv500-linux/bin/arm-hisiv500-linux-uclibcgnueabi-

unpack:
	test -e linux-3.18.y || tar -xzf ./linux-3.18.20.tar.gz
	mv ./linux-3.18.20 ./linux-3.18.y
	cd ./linux-3.18.y; patch -p1 < ../hi3519v101_for_linux_v3.18.y.patch
	cp -r ./patch/* ./linux-3.18.y
	mv ./linux-3.18.y/linux.config ./linux-3.18.y/.config

build:
	cd linux-3.18.y; make clean; make -j4 ARCH=arm CROSS_COMPILE=$(TOOLCHAIN) uImage
	cp linux-3.18.y/arch/arm/boot/uImage ./uImage



