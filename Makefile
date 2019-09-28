THIS_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

include Makefile.params

BR = buildroot-2019.08

CAMERA_IP = 192.168.10.1$(shell printf '%02d' $(CAMERA))

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

	#cp $(FAMILY)/kernel/uImage boards/$(BOARD)/build/$(APP)/uImage
	cp boards/$(BOARD)/kernel/uImage boards/$(BOARD)/build/$(APP)/uImage

	rm -f boards/$(BOARD)/build/$(APP)/rootfs.squashfs
	rm -rf boards/$(BOARD)/build/$(APP)/rootfs.tmp; mkdir boards/$(BOARD)/build/$(APP)/rootfs.tmp
	cp -r $(FAMILY)/rootfs/target/* boards/$(BOARD)/build/$(APP)/rootfs.tmp/
	test ! -e boards/$(BOARD)/putonrootfs || cp -r boards/$(BOARD)/putonrootfs/* boards/$(BOARD)/build/$(APP)/rootfs.tmp/
	cp -r $(APP)/distrib/* boards/$(BOARD)/build/$(APP)/rootfs.tmp

	mksquashfs  boards/$(BOARD)/build/$(APP)/rootfs.tmp \
                boards/$(BOARD)/build/$(APP)/rootfs.squashfs \
                -all-root

build-kernel:
	test -e ./boards/$(BOARD)/kernel || mkdir ./boards/$(BOARD)/kernel
	test ! -e ./boards/$(BOARD)/kernel/uImage || ( echo "Kernel already built"; exit 1 )
	cd ./$(FAMILY)/kernel; make clean
	test ! -e ./boards/$(BOARD)/kernel/patch || (echo "USING BOARD KERNEL PATCH!";  cp -r ./boards/$(BOARD)/kernel/patch/* ./$(FAMILY)/kernel/linux)
	test ! -e ./boards/$(BOARD)/kernel/kernel.config || (echo "USING BOARD KERNEL CONFIG!"; cp ./boards/$(BOARD)/kernel/kernel.config ./$(FAMILY)/kernel/linux/.config)
	test -e ./$(FAMILY)/kernel/linux/.config || (echo "USING DEFAULT KERNEL CONFIG!"; cp ./$(FAMILY)/kernel/$(CHIP).generic.config ./$(FAMILY)/kernel/linux/.config)
	cd ./$(FAMILY)/kernel; make build
	cp ./$(FAMILY)/kernel/uImage ./boards/$(BOARD)/kernel

########################################################################

deploy-app: pack-app
	cp boards/$(BOARD)/build/$(APP)/uImage burner/images/uImage
	cp boards/$(BOARD)/build/$(APP)/rootfs.squashfs burner/images/rootfs.squashfs
	cd burner; \
        authbind --deep ./burner.py \
            load \
            --port /dev/ttyCAM$(CAMERA) \
            --uimage ./images/uImage \
            --rootfs ./images/rootfs.squashfs \
            --ip $(CAMERA_IP) \
            --skip $(UBOOT_SIZE) \
            --initrd $(INITRD_TMP) \
            --memory $(RAM_LINUX) \
            --servercamera $(CAMERA)
	screen -L /dev/ttyCAM$(CAMERA) 115200

deploy-empty:
	cp boards/$(BOARD)/kernel/uImage burner/images/uImage
	cp $(FAMILY)/rootfs/images/rootfs.squashfs burner/images/rootfs.squashfs
	cd burner; \
        authbind --deep ./burner.py \
            load \
            --port /dev/ttyCAM$(CAMERA) \
            --uimage ./images/uImage \
            --rootfs ./images/rootfs.squashfs \
            --ip $(CAMERA_IP) \
            --skip $(UBOOT_SIZE) \
            --initrd $(INITRD_TMP) \
            --memory $(RAM_LINUX) \
            --servercamera $(CAMERA)
	screen -L /dev/ttyCAM$(CAMERA) 115200

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

$(BR):
	tar -xzf $(BR).tar.gz -C $(THIS_DIR)
	cp -r ./$(BR)-patch/* ./$(BR)
	@echo "RUN prepare-env.sh to build all toolchains and base rootfses!"

enviroiment-setup: $(BR)
	make FAMILY=hi3516cv100 toolchain
	make FAMILY=hi3516cv100 rootfs
	cd hi3516cv100/kernel; make linux
	make FAMILY=hi3516cv200 toolchain
	make FAMILY=hi3516cv200 rootfs
	cd hi3516cv200/kernel; make linux
	make FAMILY=hi3516cv300 toolchain
	make FAMILY=hi3516cv300 rootfs
	cd hi3516cv300/kernel; make linux
	make FAMILY=hi3516cv500 toolchain
	make FAMILY=hi3516cv500 rootfs
	cd hi3516cv500/kernel; make linux
	make FAMILY=hi3516av100 toolchain
	make FAMILY=hi3516av100 rootfs
	cd hi3516av100/kernel; make linux
	make FAMILY=hi3516av200 toolchain
	make FAMILY=hi3516av200 rootfs
	cd hi3516av200/kernel; make linux

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

########################################################################
