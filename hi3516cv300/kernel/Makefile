THIS_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

TOOLCHAIN=$(THIS_DIR)../toolchain/host/bin/arm-buildroot-linux-uclibcgnueabi-

linux:
	tar -xzf ./linux-3.18.20.tar.gz
	mv ./linux-3.18.20 ./linux
	cd ./linux; patch -p1 < ../hi3516cv300_for_linux_v3.18.y.patch
	#cp -r ./patch/* ./linux

build:
	cd linux; make clean; make -j8 ARCH=arm CROSS_COMPILE=$(TOOLCHAIN) uImage
	cp linux/arch/arm/boot/uImage ./uImage

clean:
	rm -f ./uImage
	cd linux; make distclean
	#rm -rf ./linux


