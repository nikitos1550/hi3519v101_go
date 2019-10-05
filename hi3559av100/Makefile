THIS_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

BR = buildroot-2019.08

guard:
	echo "USAGE: make toolchain or make rootfs"

toolchain:
	test -e toolchain || mkdir toolchain
	make -C ../$(BR) \
          O=toolchain \
            defconfig BR2_DEFCONFIG=toolchain.buildroot
	#cd toolchain;make toolchain

rootfs: toolchain
	test -e rootfs || mkdir rootfs
	make -C ../$(BR) \
          O=rootfs \
            defconfig BR2_DEFCONFIG=rootfs.buildroot
	#cd rootfs;make
