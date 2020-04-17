THIS_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

ifeq ("$(wildcard Makefile.user.params)","")
 $(error cp Makefile.user.params.example to Makefile.user.params) 
endif

ifndef NO_USER_MAKEFILE
    include $(THIS_DIR)Makefile.user.params
endif
ifndef BOARD
    $(error BOARD variable MUST be defined)
endif
ifndef CAMERA
    $(warning CAMERA variable isn't defined, only build targets are allowed)
endif


BR             := buildroot-2020.02
BUILDROOT_DIR  := $(abspath ./$(BR))
BOARD_OUTDIR   := $(abspath ./output/boards/$(BOARD))
CAMERA_TTY     := /dev/ttyCAM$(CAMERA)
CAMERA_IP      := 192.168.10.1$(shell printf '%02d' $(CAMERA))
TELNET_PORT    := 453$(shell printf '%02d' $(CAMERA))

########################################################################

CAMSTORE       := $(THIS_DIR)/facility/camstore/control.sh client

########################################################################

GREETING	?= System startup
PROMPT		?= hisilicon

########################################################################

APP             := application
APP_TARGET      ?= probe   #default target will be tester, daemon build on request durin it`s early dev stage

-include ./boards/boards/$(strip $(BOARD))/config
-include ./hisilicon/$(strip $(FAMILY))/Makefile.params

.PHONY: $(APP)/distrib/$(FAMILY) help prepare cleanall

# -----------------------------------------------------------------------------------------------------------

help:
	@echo -e "Help:\n" \
		" - make prepare                          - prepare; MUST be done once before anything\n"\
		" - make deploy-app                       - build&deploy application onto particular board\n"\
		" - make deploy-app-control-[uart|telnet] - build&deploy application, then attach control console onto particular board\n"\
		" - make control-[uart|telnet]            - attach control console onto particular board\n"\
		" - make rootfs.squashfs                  - build application and pack it within RootFS image\n"\
		" - make kernel                           - build board kernel\n"\
		" - make cleanall                         - remove all artifacts"

submodules:
	git submodule update --init --recursive
	#	git submodule init
	#	git submodule update

br-hihisim-prepare:
	make -C br-hisicam prepare

boards/boards: submodules
	ln -s ../br-hisicam/br-ext-hisicam/board boards/boards

prepare: $(BUILDROOT_DIR) submodules br-hihisim-prepare boards/boards/
	@echo "All prepared"

$(BUILDROOT_DIR):
	tar -xzf $(BR).tar.gz -C $(THIS_DIR)
	cp -r ./$(BR)-patch/* $(BUILDROOT_DIR)/

cleanall:
	if [ -d ./output ]; then chmod --recursive 777 ./output; fi
	rm -rf ./output $(BUILDROOT_DIR)
	make -C $(APP) clean
	rm -f ./boards/boards
	rm -rf ./.buildroot-ccache

info:
	@echo "BOARD=$(BOARD)"
	@echo "FAMILY=$(FAMILY)"
	@echo "APP_OVERLAY=$(APP)/distrib/$(FAMILY)"

# -----------------------------------------------------------------------------------------------------------

rootfs-only.squashfs: $(BOARD_OUTDIR)/rootfs.squashfs
	@echo "--- RootFS only image is ready: $^"

$(BOARD_OUTDIR)/rootfs.squashfs: $(BOARD_OUTDIR)/rootfs
	rm -f $@
	mksquashfs $< $@ -all-root -comp xz -b 64K -Xdict-size 100%

rootfs.squashfs: $(BOARD_OUTDIR)/rootfs+app.squashfs
	@echo "--- RootFS image is ready: $^"

kernel: $(BOARD_OUTDIR)/kernel/uImage
	@echo "--- Kernel uImage is ready: $^"

$(BOARD_OUTDIR)/rootfs+app.squashfs: $(BOARD_OUTDIR)/rootfs+app
	rm -f $@
	mksquashfs $< $@ -all-root -comp xz -b 64K -Xdict-size 100%
	#rm -f $(BOARD_OUTDIR)/rootfs+app.squashfs.xz
	#mksquashfs $< $(BOARD_OUTDIR)/rootfs+app.squashfs.xz -all-root -comp xz -b 64K -Xdict-size 100%
	#rm -f $(BOARD_OUTDIR)/rootfs+app.squashfs.lz4
	#mksquashfs $< $(BOARD_OUTDIR)/rootfs+app.squashfs.lz4 -all-root -comp lz4 -b 64K -Xhc
	#rm -f $(BOARD_OUTDIR)/rootfs+app.squashfs.lzo
	#mksquashfs $< $(BOARD_OUTDIR)/rootfs+app.squashfs.lzo -all-root -comp lzo -b 64K -Xcompression-level 9
	#rm -f $(BOARD_OUTDIR)/rootfs+app.squashfs.gzip
	#mksquashfs $< $(BOARD_OUTDIR)/rootfs+app.squashfs.gzip -all-root -comp gzip -b 64K -Xcompression-level 9

$(BOARD_OUTDIR)/rootfs+app: $(BOARD_OUTDIR)/rootfs $(APP)/distrib/$(FAMILY)
	rm -rf $@; mkdir -p $@
	cp -r $(BOARD_OUTDIR)/rootfs/* $@/
	cp -r $(APP)/distrib/$(FAMILY)/* $@/

$(APP)/distrib/$(FAMILY): $(BOARD_OUTDIR)/Makefile.params
	rm -rf $@
	make -C $(APP) PARAMS_FILE=$< APP=$(APP_TARGET) build

# -----------------------------------------------------------------------------------------------------------
# Board's artifacts

$(BOARD_OUTDIR)/rootfs:
	make -C ./boards BOARD_OUTDIR=$(BOARD_OUTDIR) BOARD=$(BOARD) rootfs

$(BOARD_OUTDIR)/Makefile.params:
	make -C ./boards BOARD_OUTDIR=$(BOARD_OUTDIR) BOARD=$(BOARD) toolchain

$(BOARD_OUTDIR)/kernel/uImage:
	make -C ./boards BOARD_OUTDIR=$(BOARD_OUTDIR) BOARD=$(BOARD) kernel

# ====================================================================================================================

build-rootfs: $(BOARD_OUTDIR)/rootfs

build-app: $(APP)/distrib/$(FAMILY)

pack-app: $(BOARD_OUTDIR)/rootfs+app.squashfs $(BOARD_OUTDIR)/kernel/uImage

pack: $(BOARD_OUTDIR)/rootfs.squashfs $(BOARD_OUTDIR)/kernel/uImage

deploy: pack
	authbind --deep scripts/hiburn.sh $(CAMERA) --verbose \
        --net-device_ip $(CAMERA_IP) \
        --net-host_ip 192.168.10.2/24 \
        --mem-linux_size $(RAM_LINUX_SIZE) \
        --linux_console "ttyAMA0,115200" \
        boot \
	--upload-addr $(KERNEL_UPLOAD_ADDR) \
        --uimage $(BOARD_OUTDIR)/kernel/uImage \
        --rootfs $(BOARD_OUTDIR)/rootfs.squashfs \
        --no-wait

#--mem-start_addr $(MEM_START_ADDR) \

deploy-app: pack-app
	authbind --deep scripts/hiburn.sh $(CAMERA) --verbose \
		--net-device_ip $(CAMERA_IP) \
        	--net-host_ip 192.168.10.2/24 \
        	--mem-linux_size $(RAM_LINUX_SIZE) \
        	--linux_console "ttyAMA0,115200" \
		boot \
		        --upload-addr $(KERNEL_UPLOAD_ADDR) \
		        --uimage $(BOARD_OUTDIR)/kernel/uImage \
        		--rootfs $(BOARD_OUTDIR)/rootfs+app.squashfs \
			--no-wait

#		--target-ip $(CAMERA_IP) --iface enp3s0 \
#		--uimage $(BOARD_OUTDIR)/kernel/uImage \
#		--rootfs $(BOARD_OUTDIR)/rootfs+app.squashfs \
#		--initrd-size $(shell ls -s --block-size=1048576 $(BOARD_OUTDIR)/rootfs+app.squashfs | cut -d' ' -f1)M --memory-size $(RAM_LINUX_SIZE) \
#		--lconsole "ttyAMA0,115200"


deploy-app-control-uart: deploy-app control-uart

deploy-control-uart: deploy control-uart

deploy-app-control-telnet: deploy-app
	@echo "waiting for 10s"
	@sleep 3
	@echo "7s more..."
	@sleep 5
	@echo "be patient, 2s more"
	telnet $(CAMERA_IP)

deploy-control-telnet: deploy
	@echo "waiting for 10s"
	@sleep 3
	@echo "7s more..."
	@sleep 5
	@echo "be patient, 2s more"
	telnet $(CAMERA_IP)


########################################################################

deprecated-control-uart:
	miniterm $(CAMERA_TTY) 115200

catch-uboot:
	scripts/hiburn.sh $(CAMERA) -v printenv
	$(CAMSTORE) forward_serial $(CAMERA_TTY)

control-uart:
	#telnet localhost $(TELNET_PORT)
	$(CAMSTORE) forward_serial $(CAMERA_TTY)

deprecated-control-uart-%:
	miniterm /dev/ttyCAM$(subst control-uart-,,$@) 115200

control-uart-%:
	#telnet localhost 453$(shell printf '%02d' $(subst control-uart-,,$@))
	$(CAMSTORE) forward_serial /dev/ttyCAM$(subst control-uart-,,$@)

control-telnet:
	telnet $(CAMERA_IP)

control-telnet-%:
	telnet 192.168.10.1$(shell printf '%02d' $(subst control-telnet-,,$@))

########################################################################

pack-archive: pack-app
	mkdir -p $(BOARD_OUTDIR)/$(BOARD)
	cp $(BOARD_OUTDIR)/kernel/uImage $(BOARD_OUTDIR)/$(BOARD)/uImage
	cp $(BOARD_OUTDIR)/rootfs+app.squashfs $(BOARD_OUTDIR)/$(BOARD)/rootfs+app.squashfs
	cd $(BOARD_OUTDIR); tar -cvzf ./$(BOARD).tar.gz ./$(BOARD)
	rm -rf $(BOARD_OUTDIR)/$(BOARD)
