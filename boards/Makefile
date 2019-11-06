ifndef BOARD
    $(error 'BOARD' variable must be defined)
endif
-include ./$(BOARD)/config

# source directories
BUILDROOT_DIR  ?= $(abspath ../buildroot-2019.08)
BOARD_DIR      := $(abspath ./$(BOARD))
FAMILY_DIR     := $(abspath ../$(FAMILY))

# destination directories
OUTDIR         ?= $(abspath ../build)
BOARD_OUTDIR   := $(OUTDIR)/boards/$(BOARD)
FAMILY_OUTDIR  := $(OUTDIR)/families/$(FAMILY)
TOOLCHAIN_DIR  := $(FAMILY_OUTDIR)/toolchain
ROOTFS_DIR     := $(OUTDIR)/rootfs

# buildroot configs
FAMILY_ROOTFS_BR_CONFIG      := $(FAMILY_OUTDIR)/.br-rootfs.config
FAMILY_TOOLCHAIN_BR_CONFIG   := $(FAMILY_DIR)/toolchain.buildroot  


.PHONY: help info toolchain rootfs


help:
	@echo "Help: \n" \
		" - make info        - just print info about target board \n" \
		" - make toolchain   - build toolchain \n" \
		" - make rootfs      - build RootFS directory \n" \


info:
	@echo ""\
		"-- Board information:\n"\
		"  BOARD   = $(BOARD)\n"\
		"  FAMILY  = $(FAMILY)\n"\
		"  CHIP    = $(CHIP)\n"\
		"-- Destinations:\n"\

		"  Toolchain: $(BOARD_OUTDIR)/toolchain\n"\
		"  RootFS:    $(BOARD_OUTDIR)/rootfs\n"


# ====================================================================================================================


toolchain: $(BOARD_OUTDIR)/toolchain
	@echo "-- Toolchain for $(BOARD) is ready"


rootfs: $(BOARD_OUTDIR)/rootfs
	@echo "-- RootFS for $(BOARD) is ready"


# ====================================================================================================================
$(BOARD_OUTDIR)/toolchain: $(FAMILY_OUTDIR)/toolchain
	mkdir -p $(BOARD_OUTDIR)
	ln -f --symbolic $(FAMILY_OUTDIR)/toolchain $@


$(BOARD_OUTDIR)/rootfs: $(FAMILY_OUTDIR)/rootfs $(BOARD_DIR)/putonrootfs $(BOARD_DIR)/config
	if [ -d $@ ]; then rm -rf $@; fi
	mkdir -p $@
	cp -r $(FAMILY_OUTDIR)/rootfs/* $@/
	if [ -d $(BOARD_DIR)/putonrootfs ]; then cp -r $(BOARD_DIR)/putonrootfs/* $@/; fi
	cat $(BOARD_DIR)/config | tr -d "[:blank:]" > $@/etc/board.config


# ====================================================================================================================
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


# ====================================================================================================================
# hack for compatibility
$(FAMILY_ROOTFS_BR_CONFIG): $(FAMILY_DIR)/rootfs.buildroot
	sed 's|BR2_TOOLCHAIN_EXTERNAL_PATH=.*|BR2_TOOLCHAIN_EXTERNAL_PATH="$(FAMILY_OUTDIR)/toolchain"|' $< > $@
