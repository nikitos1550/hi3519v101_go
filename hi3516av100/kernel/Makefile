THIS_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

TOOLCHAIN=$(THIS_DIR)../toolchain/host/bin/arm-buildroot-linux-uclibcgnueabi-

linux:
	tar -xzf ./linux-3.4.y.tgz
	mv ./linux-3.4.y ./linux
	cp -r ./patchfiles/* ./linux

build:
	cd linux; make clean; make -j4 ARCH=arm CROSS_COMPILE=$(TOOLCHAIN) uImage
	cp linux/arch/arm/boot/uImage ./uImage

clean:
	rm -f ./uImage
	cd linux; make distclean
	#rm -rf ./linux

