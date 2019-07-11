THIS_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

unpack:
	test -e buildroot-2019.05.1 || tar -xzf buildroot-2019.05.1.tar.gz -C $(THIS_DIR)
	cd buildroot-2019.05.1; make defconfig BR2_DEFCONFIG=$(THIS_DIR)/defconfig.buildroot
	#cp -r ./buildroot-2019.05.1-patch/* ./buildroot-2019.05.1

pack:
	cd src; make; cp ./hello $(THIS_DIR)putonrootfs/opt
	cd $(THIS_DIR)buildroot-2019.02.3; make
	cp $(THIS_DIR)buildroot-2019.02.3/output/images/rootfs.romfs $(THIS_DIR)burner/images
