THIS_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

BR = buildroot-2019.08

guard:
	echo "USAGE: make toolchain or make rootfs"

toolchain:
	test -e toolchain || mkdir toolchain
	make -C $(THIS_DIR)../$(BR) \
          O=$(THIS_DIR)toolchain \
            defconfig BR2_DEFCONFIG=$(THIS_DIR)toolchain.buildroot
	#cd toolchain;make toolchain

rootfs: toolchain
	test -e rootfs || mkdir rootfs
	make -C $(THIS_DIR)../$(BR) \
          O=$(THIS_DIR)rootfs \
            defconfig BR2_DEFCONFIG=$(THIS_DIR)rootfs.buildroot
	#cd rootfs;make
