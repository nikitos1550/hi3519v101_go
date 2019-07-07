THIS_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

unpack:
	test -e buildroot-2019.02.3 || tar -xzf buildroot-2019.02.3.tar.gz -C $(THIS_DIR)
	cd buildroot-2019.02.3; make defconfig BR2_DEFCONFIG=$(THIS_DIR)/defconfig.buildroot

pack:
	cd src; make; cp ./hello $(THIS_DIR)putonrootfs/opt
	cd $(THIS_DIR)buildroot-2019.02.3; make
	cp $(THIS_DIR)buildroot-2019.02.3/output/images/rootfs.romfs $(THIS_DIR)burner/images
