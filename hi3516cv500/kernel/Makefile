THIS_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

TOOLCHAIN=$(THIS_DIR)../toolchain/host/bin/arm-buildroot-linux-uclibcgnueabi-

unpack:
	test -e linux-4.9.y || tar -xzf ./linux-4.9.37.tar.gz
	mv ./linux-4.9.37 ./linux-4.9.y
	cd ./linux-4.9.y; patch -p1 < ../linux-4.9.37.patch
	#cp -r ./patch/* ./linux-3.18.y
	#mv ./linux-3.18.y/linux.config ./linux-3.18.y/.config

build:
	cd linux-4.9.y; make clean; make -j4 ARCH=arm CROSS_COMPILE=$(TOOLCHAIN) uImage
	cp linux-4.9.y/arch/arm/boot/uImage ./uImage



