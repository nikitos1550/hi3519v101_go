THIS_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

TOOLCHAIN=$(THIS_DIR)../toolchain/host/bin/arm-buildroot-linux-uclibcgnueabi-

linux:
	mkdir linux; tar -xzf ./linux-3.0.y.b0.patched2.tgz -C ./linux
	#cp -r ./patch/* ./linux

build:
	cd linux; make clean; make -j8 ARCH=arm CROSS_COMPILE=$(TOOLCHAIN) uImage
	cp linux/arch/arm/boot/uImage ./uImage

clean:
	rm -f ./uImage
	cd linux; make distclean
	#rm -rf ./linux

