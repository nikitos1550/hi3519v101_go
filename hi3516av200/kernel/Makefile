THIS_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

include ../Makefile.params

linux:
	tar -xzf ./linux-3.18.20.tar.gz
	mv ./linux-3.18.20 ./linux
	cd ./linux; patch -p1 < ../hi3519v101_for_linux_v3.18.y.patch
	cp -r ./patch/* ./linux

build:
	cd linux; make clean; make -j8 ARCH=arm CROSS_COMPILE=$(TOOLCHAIN) uImage
	cp linux/arch/arm/boot/uImage ./uImage

clean:
	rm -f ./uImage
	cd linux; make distclean
	#rm -rf ./linux
