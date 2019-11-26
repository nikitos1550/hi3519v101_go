THIS_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

include Makefile.user.params

BR             := buildroot-2019.08
BUILDROOT_DIR  := $(abspath ./buildroot-2019.08)
BOARD_OUTDIR   := $(abspath ./output/boards/$(BOARD))
CAMERA_IP      := 192.168.10.1$(shell printf '%02d' $(CAMERA))

########################################################################

-include ./boards/$(BOARD)/config

.PHONY: $(APP)/distrib/$(FAMILY) help prepare cleanall

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

########################################################################

$(BR):
	tar -xzf $(BR).tar.gz -C $(THIS_DIR)
	cp -r ./$(BR)-patch/* ./$(BR)

enviroiment-setup: $(BR)

########################################################################

control:
	screen -L /dev/ttyCAM$(CAMERA) 115200

control-%:
	screen -L /dev/ttyCAM$(subst control-,,$@) 115200

########################################################################
