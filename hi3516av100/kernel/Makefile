THIS_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

TOOLCHAIN=$(THIS_DIR)../toolchain/host/bin/arm-buildroot-linux-uclibcgnueabi-

unpack:
	test -e linux || tar -xzf ./linux-3.4.y.tgz
	mv ./linux-3.4.y ./linux
	#cp -r ./patch/* ./linux-3.18.y

build:
	cd linux; make clean; make -j4 ARCH=arm CROSS_COMPILE=$(TOOLCHAIN) uImage
	cp linux/arch/arm/boot/uImage ./uImage

clean:
	rm -rf ./linux

