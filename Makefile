THIS_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

include Makefile.user.params

BR             := buildroot-2019.08
BUILDROOT_DIR  := $(abspath ./buildroot-2019.08)
BOARD_OUTDIR   := $(abspath ./output/boards/$(BOARD))
CAMERA_IP      := 192.168.10.1$(shell printf '%02d' $(CAMERA))

########################################################################

.PHONY: $(APP)/distrib/$(FAMILY) help prepare cleanall 

-include ./boards/$(BOARD)/config


# -----------------------------------------------------------------------------------------------------------

help:
	@echo -e "Help:\n"\
		" - make prepare          - prepare; MUST be done before anything\n"\
		" - make rootfs.squashfs  - build application and pack it within RootFS image\n"\
		" - make deploy-app       - build&deploy application onto prticular board\n"\
		" - make cleanall         - remove all artifacts"

prepare: $(BUILDROOT_DIR)
	@echo "All prepared"

$(BUILDROOT_DIR):
	tar -xzf $(BR).tar.gz -C $(THIS_DIR)
	cp -r ./$(BR)-patch/* $(BUILDROOT_DIR)/

cleanall:
	if [ -d ./output ]; then chmod --recursive 777 ./output; fi
	rm -rf ./output $(BUILDROOT_DIR)


# -----------------------------------------------------------------------------------------------------------

rootfs.squashfs: $(BOARD_OUTDIR)/rootfs+app.squashfs
	@echo "--- RootFS image is ready: $^"

kernel: $(BOARD_OUTDIR)/kernel/uImage
	@echo "--- Kernel uImage is ready: $^"

$(BOARD_OUTDIR)/rootfs+app.squashfs: $(BOARD_OUTDIR)/rootfs+app
	rm -f $@
	mksquashfs $< $@ -all-root

$(BOARD_OUTDIR)/rootfs+app: $(BOARD_OUTDIR)/rootfs $(APP)/distrib/$(FAMILY)
	rm -rf $@; mkdir -p $@
	cp -r $(BOARD_OUTDIR)/rootfs/* $@/
	cp -r $(APP)/distrib/$(FAMILY)/* $@/

$(APP)/distrib/$(FAMILY): $(BOARD_OUTDIR)/Makefile.params
	rm -rf $@
	make -C $(APP) PARAMS_FILE=$< build-tester

# -----------------------------------------------------------------------------------------------------------
# Board's artifacts

$(BOARD_OUTDIR)/rootfs:
	make -C ./boards BOARD_OUTDIR=$(BOARD_OUTDIR) BOARD=$(BOARD) rootfs

$(BOARD_OUTDIR)/Makefile.params:
	make -C ./boards BOARD_OUTDIR=$(BOARD_OUTDIR) BOARD=$(BOARD) toolchain

$(BOARD_OUTDIR)/kernel/uImage:
	make -C ./boards BOARD_OUTDIR=$(BOARD_OUTDIR) BOARD=$(BOARD) kernel

# ====================================================================================================================


deprecated-pack-app:
	@[ "$(BOARD)" ] && echo "all good" || ( echo "var is not set"; exit 1 )
	cd $(APP); make FAMILY=$(FAMILY) build

	test -e boards/$(BOARD)/build/ || mkdir boards/$(BOARD)/build
	test -e boards/$(BOARD)/build/$(APP) || mkdir boards/$(BOARD)/build/$(APP)

	cp boards/$(BOARD)/kernel/uImage boards/$(BOARD)/build/$(APP)/uImage

	
	rm -rf boards/$(BOARD)/build/$(APP)/rootfs.tmp; mkdir boards/$(BOARD)/build/$(APP)/rootfs.tmp
	cp -r $(FAMILY)/rootfs/target/* boards/$(BOARD)/build/$(APP)/rootfs.tmp/
	test ! -e boards/$(BOARD)/putonrootfs || cp -r boards/$(BOARD)/putonrootfs/* boards/$(BOARD)/build/$(APP)/rootfs.tmp/
	#cp boards/$(BOARD)/config boards/$(BOARD)/build/$(APP)/rootfs.tmp/etc/board.config
	cat boards/$(BOARD)/config | tr -d "[:blank:]" > boards/$(BOARD)/build/$(APP)/rootfs.tmp/etc/board.config
	cp -r $(APP)/distrib/$(FAMILY)/* boards/$(BOARD)/build/$(APP)/rootfs.tmp

	mksquashfs  boards/$(BOARD)/build/$(APP)/rootfs.tmp \
                boards/$(BOARD)/build/$(APP)/rootfs.squashfs \
                -all-root

deprecated-build-kernel:
	test -e ./boards/$(BOARD)/kernel || mkdir ./boards/$(BOARD)/kernel
	test ! -e ./boards/$(BOARD)/kernel/uImage || ( echo "Kernel already built"; exit 1 )
	cd ./$(FAMILY)/kernel; make clean
	test ! -e ./boards/$(BOARD)/kernel/patch || (echo "USING BOARD KERNEL PATCH!";  cp -r ./boards/$(BOARD)/kernel/patch/* ./$(FAMILY)/kernel/linux)
	test ! -e ./boards/$(BOARD)/kernel/kernel.config || (echo "USING BOARD KERNEL CONFIG!"; cp ./boards/$(BOARD)/kernel/kernel.config ./$(FAMILY)/kernel/linux/.config)
	test -e ./$(FAMILY)/kernel/linux/.config || (echo "USING DEFAULT KERNEL CONFIG!"; cp ./$(FAMILY)/kernel/$(CHIP).generic.config ./$(FAMILY)/kernel/linux/.config)
	cd ./$(FAMILY)/kernel; make build
	cp ./$(FAMILY)/kernel/uImage ./boards/$(BOARD)/kernel

########################################################################

build-app: $(APP)/distrib/$(FAMILY)

pack-app: $(BOARD_OUTDIR)/rootfs+app.squashfs $(BOARD_OUTDIR)/kernel/uImage

deploy-app: pack-app
	cd burner; authbind --deep ./burner2.py \
		--port /dev/ttyCAM$(CAMERA) \
		--reset-power "./power.py reset $(CAMERA)" \
		load \
		--target-ip $(CAMERA_IP) --iface enp2s0 \
		--uimage $(BOARD_OUTDIR)/kernel/uImage \
		--rootfs $(BOARD_OUTDIR)/rootfs+app.squashfs \
		--initrd-size 16M --memory-size 256M

deploy-app-control: deploy-app
	screen -L /dev/ttyCAM$(CAMERA) 115200

deprecated-deploy-empty:
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

deprecated-toolchain:
	test -e $(THIS_DIR)/$(FAMILY)/toolchain || mkdir $(THIS_DIR)/$(FAMILY)/toolchain
	make -C $(THIS_DIR)/$(BR) \
          O=$(THIS_DIR)/$(FAMILY)/toolchain \
            defconfig BR2_DEFCONFIG=$(THIS_DIR)/$(FAMILY)/toolchain.buildroot
	cd $(THIS_DIR)/$(FAMILY)/toolchain; make toolchain

deprecated-rootfs: toolchain
	test -e $(THIS_DIR)/$(FAMILY)/rootfs || mkdir $(THIS_DIR)/$(FAMILY)/rootfs
	make -C $(THIS_DIR)/$(BR) \
          O=$(THIS_DIR)/$(FAMILY)/rootfs \
            defconfig BR2_DEFCONFIG=$(THIS_DIR)/$(FAMILY)/rootfs.buildroot
	cd $(THIS_DIR)/$(FAMILY)/rootfs; make

########################################################################

$(BR):
	tar -xzf $(BR).tar.gz -C $(THIS_DIR)
	cp -r ./$(BR)-patch/* ./$(BR)
	#@echo "RUN prepare-env.sh to build all toolchains and base rootfses!"

enviroiment-setup: $(BR)
    #deprecated
	#make FAMILY=hi3516cv100 toolchain
	#make FAMILY=hi3516cv100 rootfs
	#cd hi3516cv100/kernel; make linux
	#make FAMILY=hi3516cv200 toolchain
	#make FAMILY=hi3516cv200 rootfs
	#cd hi3516cv200/kernel; make linux
	#make FAMILY=hi3516cv300 toolchain
	#make FAMILY=hi3516cv300 rootfs
	#cd hi3516cv300/kernel; make linux
	#make FAMILY=hi3516cv500 toolchain
	#make FAMILY=hi3516cv500 rootfs
	#cd hi3516cv500/kernel; make linux
	#make FAMILY=hi3516av100 toolchain
	#make FAMILY=hi3516av100 rootfs
	#cd hi3516av100/kernel; make linux
	#make FAMILY=hi3516av200 toolchain
	#make FAMILY=hi3516av200 rootfs
	#cd hi3516av200/kernel; make linux

########################################################################

#deploy-burner:
#	cd burner; \
#		authbind --deep ./burner.py \
#			load \
#			--uimage ./images/uImage \
#			--rootfs ./images/rootfs.squashfs \
#			--ip $(CAMERA_IP) \
#           --initrd 16 \
#			--memory $(RAM_LINUX) \
#			--servercamera $(CAMERA)
#	screen /dev/ttyCAM$(CAMERA) 115200

########################################################################

control:
	screen -L /dev/ttyCAM$(CAMERA) 115200

control-%:
	screen -L /dev/ttyCAM$(subst control-,,$@) 115200

########################################################################
