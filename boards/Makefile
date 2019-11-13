ifndef BOARD
    $(error 'BOARD' variable must be defined)
endif
-include ./$(BOARD)/config

# source directories
BUILDROOT_DIR  ?= $(abspath ../buildroot-2019.08)
BOARD_DIR      := $(abspath ./$(BOARD))
FAMILY_DIR     := $(abspath ../$(FAMILY))
ifdef APP  # TODO
    APP_DIR    := $(abspath ../api)
endif


# destination directories
OUTDIR         ?= $(abspath ../build)
BOARD_OUTDIR   := $(OUTDIR)/boards/$(BOARD)
FAMILY_OUTDIR  := $(OUTDIR)/families/$(FAMILY)

# artifacts
TOOLCHAIN_DIR  := $(BOARD_OUTDIR)/toolchain
KERNEL_IMAGE   := $(BOARD_OUTDIR)/uImage
ifdef APP
    ROOTFS_DIR := $(BOARD_OUTDIR)/rootfs+$(APP)
else
    ROOTFS_DIR := $(BOARD_OUTDIR)/rootfs
endif
ROOTFS_IMAGE   := $(ROOTFS_DIR).squashfs

.PHONY: help info toolchain rootfs kernel rootfs.squashfs


help:
	@echo "Help: \n" \
		" - make help        - print this help \n" \
		" - make info        - just print info about target board \n" \
		" - make toolchain   - build toolchain \n" \
		" - make rootfs      - build RootFS directory \n" \
		" - make kernel      - build Linux Kernel's uImage \n"


info:
	@echo ""\
		"-- Board information:\n"\
		"  BOARD   = $(BOARD)\n"\
		"  FAMILY  = $(FAMILY)\n"\
		"  CHIP    = $(CHIP)\n"\
		"-- Sources:\n"\
		"  Buildroot:           '$(BUILDROOT_DIR)'\n"\
		"  Application FS dir:  '$(APP_DIR)'\n"\
		"-- Destinations:\n"\
		"  Toolchain:        $(TOOLCHAIN_DIR)\n"\
		"  RootFS:           $(ROOTFS_DIR)\n"\
		"  RootFS image:     $(ROOTFS_IMAGE)\n"\
		"  Kernel's uImage:  $(KERNEL_IMAGE)\n"


# ====================================================================================================================
toolchain: $(TOOLCHAIN_DIR)
	@echo "-- Toolchain for $(BOARD) is ready: $<"

rootfs: $(ROOTFS_DIR)
	@echo "-- RootFS for $(BOARD) is ready: $<"

rootfs.squashfs: $(ROOTFS_IMAGE)
	@echo "-- RootFS image for $(BOARD) is ready: $<"

kernel: $(KERNEL_IMAGE)
	@echo "-- Linux Kernel for $(BOARD) is ready: $<"


# ====================================================================================================================
# BOARD
# --------------------------------------------------------------------------------------------------------------------

$(TOOLCHAIN_DIR): $(FAMILY_OUTDIR)/toolchain
	mkdir -p $(BOARD_OUTDIR)
	ln -f --symbolic $< $@

$(ROOTFS_DIR): $(FAMILY_OUTDIR)/rootfs $(BOARD_DIR)/putonrootfs $(BOARD_DIR)/config $(APP_DIR)
	if [ -d $@ ]; then rm -rf $@; fi
	mkdir -p $@
	cp -r $(FAMILY_OUTDIR)/rootfs/* $@/
	if [ -d $(BOARD_DIR)/putonrootfs ]; then cp -r $(BOARD_DIR)/putonrootfs/* $@/; fi
	cat $(BOARD_DIR)/config | tr -d "[:blank:]" > $@/etc/board.config
ifdef APP_DIR
	if [ -d $(APP_DIR) ]; then cp -r $(APP_DIR)/* $@/; fi
endif


$(ROOTFS_IMAGE): $(ROOTFS_DIR)
	mksquashfs $< $@ -all-root

$(KERNEL_IMAGE): $(FAMILY_DIR)/kernel/linux $(TOOLCHAIN_DIR)
	make -C $(FAMILY_DIR)/kernel clean
	if [ -e $(BOARD_DIR)/kernel/patch ]; then cp -r $(BOARD_DIR)/kernel/patch/* $(FAMILY_DIR)/kernel/linux; fi
	if [ -e $(BOARD_DIR)/kernel/kernel.config ]; then \
		cp $(BOARD_DIR)/kernel/kernel.config $(FAMILY_DIR)/kernel/linux/.config; \
	elif [ -e $(FAMILY_DIR)/kernel/$(CHIP).generic.config ]; then \
		cp $(FAMILY_DIR)/kernel/$(CHIP).generic.config $(FAMILY_DIR)/kernel/linux/.config; \
	fi
	make -C $(FAMILY_DIR)/kernel TOOLCHAIN_DIR=$(TOOLCHAIN_DIR) build 
	cp -f $(FAMILY_DIR)/kernel/uImage $@


# ====================================================================================================================
# FAMILY
# --------------------------------------------------------------------------------------------------------------------

# buildroot configs
FAMILY_ROOTFS_BR_CONFIG      := $(FAMILY_OUTDIR)/.br-rootfs.config
FAMILY_TOOLCHAIN_BR_CONFIG   := $(FAMILY_DIR)/toolchain.buildroot  


$(FAMILY_OUTDIR)/toolchain: $(BUILDROOT_DIR) $(FAMILY_TOOLCHAIN_BR_CONFIG)
	@echo "-- Build toolchain for $(FAMILY) family"
	mkdir -p $(FAMILY_OUTDIR)
	make -C $(BUILDROOT_DIR) O=$(FAMILY_OUTDIR)/.br-toolchain defconfig BR2_DEFCONFIG=$(FAMILY_TOOLCHAIN_BR_CONFIG)
	make -C $(FAMILY_OUTDIR)/.br-toolchain toolchain
	ln -f --symbolic $(FAMILY_OUTDIR)/.br-toolchain/host $@


$(FAMILY_OUTDIR)/rootfs: $(BUILDROOT_DIR) $(FAMILY_ROOTFS_BR_CONFIG) $(FAMILY_OUTDIR)/toolchain
	@echo "-- Build RootFS for $(FAMILY) family"
	mkdir -p $(FAMILY_OUTDIR)
	make -C $(BUILDROOT_DIR) O=$(FAMILY_OUTDIR)/.br-rootfs defconfig BR2_DEFCONFIG=$(FAMILY_ROOTFS_BR_CONFIG)
	make -C $(FAMILY_OUTDIR)/.br-rootfs
	ln -f --symbolic $(FAMILY_OUTDIR)/.br-rootfs/target $@


# Prepare family's kernel sources
$(FAMILY_DIR)/kernel/linux:
	make -C $(FAMILY_DIR)/kernel linux


# hack for compatibility
$(FAMILY_ROOTFS_BR_CONFIG): $(FAMILY_DIR)/rootfs.buildroot
	mkdir -p $(FAMILY_OUTDIR)
	sed 's|BR2_TOOLCHAIN_EXTERNAL_PATH=.*|BR2_TOOLCHAIN_EXTERNAL_PATH="$(FAMILY_OUTDIR)/toolchain"|' $< > $@
