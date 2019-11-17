LINUX_DIR ?= $(abspath ./linux)
UIMAGE_DIR := $(abspath $(LINUX_DIR)/..)

include ../Makefile.params

linux:
	tar -xzf ./linux-3.18.20.tar.gz
	mv ./linux-3.18.20 ./linux
	cd ./linux; patch -p1 < ../hi3519v101_for_linux_v3.18.y.patch
	cp -r ./patchfiles/* ./linux

build:
	cd $(LINUX_DIR); make clean; make -j8 ARCH=arm CROSS_COMPILE=$(TOOLCHAIN) uImage
	cp -f $(LINUX_DIR)/arch/arm/boot/uImage ./uImage

clean:
	rm -f ./uImage
	cd linux; make distclean
	#rm -rf ./linux
