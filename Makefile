THIS_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

include Makefile.params

BR = buildroot-2019.08

CAMERA_IP = 192.169.0.1$(shell printf '%02d' $(CAMERA))

########################################################################

-include ./boards/$(BOARD)/config

guard:
	@echo "USAGE:"
	@echo "prepare env:"
	@echo "make enviroiment-setup, make FAMILY=XXX toolchain, make FAMILY=XXX rootfs"
	@echo "build and deploy:"
	@echo "make BOARD=XXX pack-app"
	#@echo "family $(FAMILY)"
	#@echo $(CAMERA_IP)

pack-app:
	@[ "$(BOARD)" ] && echo "all good" || ( echo "var is not set"; exit 1 )
	cd $(APP); make FAMILY=$(FAMILY) build
	#@echo "FAMILY = $(FAMILY)"
	#@echo "CAMERA FACILITY ID = $(CAMERA)"
	test -e boards/$(BOARD)/build/ || mkdir boards/$(BOARD)/build
	test -e boards/$(BOARD)/build/$(APP) || mkdir boards/$(BOARD)/build/$(APP)
	cp $(FAMILY)/kernel/uImage boards/$(BOARD)/build/$(APP)/uImage
	rm -f boards/$(BOARD)/build/$(APP)/rootfs.squashfs
	rm -rf boards/$(BOARD)/build/$(APP)/rootfs.tmp; mkdir boards/$(BOARD)/build/$(APP)/rootfs.tmp
	cp -r $(FAMILY)/rootfs/target/* boards/$(BOARD)/build/$(APP)/rootfs.tmp/
	test ! -e boards/$(BOARD)/putonrootfs || cp -r boards/$(BOARD)/putonrootfs/* boards/$(BOARD)/build/$(APP)/rootfs.tmp/
	mksquashfs  boards/$(BOARD)/build/$(APP)/rootfs.tmp \
                boards/$(BOARD)/build/$(APP)/rootfs.squashfs \
                -all-root


	#cp boards/$(BOARD)/build/uImage burner/images/uImage
	#cp boards/$(BOARD)/build/rootfs.squashfs burner/images/rootfs.squashfs
	#cp $(FAMILY)/rootfs/images/rootfs.squashfs burner/images/rootfs.squashfs
	#cd burner; \
    #    authbind --deep ./burner.py \
    #        load \
    #        --port /dev/ttyCAM$(CAMERA) \
    #        --uimage ./images/uImage \
    #        --rootfs ./images/rootfs.squashfs \
    #        --ip $(CAMERA_IP) \
    #        --skip 1024 \
    #        --initrd 4 \
    #        --memory $(RAM_LINUX) \
    #        --servercamera $(CAMERA)
	#screen -L /dev/ttyCAM$(CAMERA) 115200

board-kernel:
	@echo "todo"

deploy-app: pack-app deploy-burner
	@echo "deploy-app"

########################################################################

toolchain:
	test -e $(THIS_DIR)/$(FAMILY)/toolchain || mkdir $(THIS_DIR)/$(FAMILY)/toolchain
	make -C $(THIS_DIR)/$(BR) \
          O=$(THIS_DIR)/$(FAMILY)/toolchain \
            defconfig BR2_DEFCONFIG=$(THIS_DIR)/$(FAMILY)/toolchain.buildroot
	cd $(THIS_DIR)/$(FAMILY)/toolchain; make toolchain

rootfs: toolchain
	test -e $(THIS_DIR)/$(FAMILY)/rootfs || mkdir $(THIS_DIR)/$(FAMILY)/rootfs
	make -C $(THIS_DIR)/$(BR) \
          O=$(THIS_DIR)/$(FAMILY)/rootfs \
            defconfig BR2_DEFCONFIG=$(THIS_DIR)/$(FAMILY)/rootfs.buildroot
	cd $(THIS_DIR)/$(FAMILY)/rootfs; make

########################################################################
enviroiment-setup:
	tar -xzf $(BR).tar.gz -C $(THIS_DIR)
	cp -r ./$(BR)-patch/* ./$(BR)

########################################################################

deploy-burner:
	cd burner; \
		authbind --deep ./burner.py \
			load \
			--uimage ./images/uImage \
			--rootfs ./images/rootfs.squashfs \
			--ip $(CAMERA_IP) \
            --initrd 16 \
			--memory $(RAM_LINUX) \
			--servercamera $(CAMERA)
	screen /dev/ttyCAM$(CAMERA) 115200

########################################################################
camera-serial:
	screen -L /dev/ttyCAM$(CAMERA) 115200

camera-serial-%:
	screen -L /dev/ttyCAM$(subst camera-serial-,,$@) 115200
