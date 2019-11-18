SHELL := /bin/bash

ifndef BOARD
    $(error 'BOARD' variable must be defined)
endif

# TODO: It may be needed some magic here when the config becomes JSON
-include ./$(BOARD)/config

# source directories
BUILDROOT_DIR  ?= $(abspath ../buildroot-2019.08)
BOARD_DIR      := $(abspath ./$(BOARD))
FAMILY_DIR     := $(abspath ../$(FAMILY))

# destination directories
FAMILY_OUTDIR  ?= $(abspath ../output/$(FAMILY))
BOARD_OUTDIR   ?= $(abspath ../output/boards/$(BOARD))

# artifacts
TOOLCHAIN_DIR       := $(BOARD_OUTDIR)/toolchain
KERNEL_SOURCE_DIR   := $(BOARD_OUTDIR)/kernel/linux
KERNEL_UIMAGE       := $(BOARD_OUTDIR)/kernel/uImage
ROOTFS_DIR          := $(BOARD_OUTDIR)/rootfs
MAKEFILE_PARAMS     := $(BOARD_OUTDIR)/Makefile.params


.PHONY: help info toolchain rootfs kernel clean all


help:
	@echo -en "Help: \n" \
		" - make help           - print this help \n" \
		" - make info           - print info of target board and artifacts \n" \
		" - make toolchain      - build toolchain \n" \
		" - make rootfs         - build RootFS directory \n" \
		" - make kernel         - build Linux Kernel's uImage \n" \
		" - make makfile.params - build aggregated Makefile.params \n" \
		" - make clean          - completely remove board's outdir \n"


info:
	@echo -en ""\
		"-- Board information:\n"\
		"  BOARD   = $(BOARD)\n"\
		"  FAMILY  = $(FAMILY)\n"\
		"  CHIP    = $(CHIP)\n"\
		"-- Sources:\n"\
		"  Buildroot:        $(BUILDROOT_DIR)\n"\
		"-- Destinations:\n"\
		"  Toolchain:        $(TOOLCHAIN_DIR)\n"\
		"  RootFS:           $(ROOTFS_DIR)\n"\
		"  Kernel sources:   $(KERNEL_SOURCE_DIR)\n"\
		"  Kernel uImage:    $(KERNEL_UIMAGE)\n"\
		"  Makefile.params:  $(MAKEFILE_PARAMS)\n"


# ====================================================================================================================
all: toolchain rootfs kernel makefile.params

toolchain: $(TOOLCHAIN_DIR)
	@echo "-- Toolchain for $(BOARD) is ready: $<"

rootfs: $(ROOTFS_DIR)
	@echo "-- RootFS for $(BOARD) is ready: $<"

kernel: $(KERNEL_UIMAGE)
	@echo "-- Linux Kernel for $(BOARD) is ready: $<"

makefile.params: $(MAKEFILE_PARAMS)
	@echo "-- Makefiles.params for $(BOARD) is ready: $<"

clean:
	rm -rf $(BOARD_OUTDIR)

# ====================================================================================================================
# BOARD
# --------------------------------------------------------------------------------------------------------------------

$(MAKEFILE_PARAMS): $(BOARD_DIR)/config $(FAMILY_DIR)/Makefile.params
	rm -rf $@; mkdir -p $(BOARD_OUTDIR)
	echo -e "# THIS FILE WAS GENERATED, PLEASE DO NOT MODIFY IT\n\
	BOARD         := $(BOARD)\n\
	TOOLCHAIN_DIR := $(TOOLCHAIN_DIR)" >> $@
	echo -e "\n# Parameters of the board" >> $@; cat $(BOARD_DIR)/config >> $@;
	echo -e "\n# Parameters of the family" >> $@; cat $(FAMILY_DIR)/Makefile.params >> $@

$(TOOLCHAIN_DIR): $(FAMILY_OUTDIR)/toolchain
	rm -rf $@; mkdir -p $(BOARD_OUTDIR)
	ln -f --symbolic $< $@

$(ROOTFS_DIR): $(FAMILY_OUTDIR)/rootfs $(BOARD_DIR)/putonrootfs $(BOARD_DIR)/config $(APP_DIR)
	if [ -d $@ ]; then rm -rf $@; fi
	mkdir -p $@
	cp -r $(FAMILY_OUTDIR)/rootfs/* $@/
	if [ -d $(BOARD_DIR)/putonrootfs ]; then cp -r $(BOARD_DIR)/putonrootfs/* $@/; fi
	cat $(BOARD_DIR)/config | tr -d "[:blank:]" > $@/etc/board.config

# TODO: For now Kernel is built by family's makefile that should be fixed later
$(KERNEL_UIMAGE): $(KERNEL_SOURCE_DIR) $(TOOLCHAIN_DIR)
	make -C $(FAMILY_DIR)/kernel LINUX_DIR=$(KERNEL_SOURCE_DIR) TOOLCHAIN_DIR=$(TOOLCHAIN_DIR) build
	cp -f $(FAMILY_DIR)/kernel/uImage $@

$(KERNEL_SOURCE_DIR): $(FAMILY_DIR)/kernel/linux
	if [ -d $@ ]; then rm -rf $@; fi
	mkdir -p $@
# Copy family's pristine kernel source tree 
	make -C $(FAMILY_DIR)/kernel clean
	cp -r $(FAMILY_DIR)/kernel/linux/* $@/
# Apply the board's patch
	if [ -e $(BOARD_DIR)/kernel/patch ]; then cp -r $(BOARD_DIR)/kernel/patch/* $@; fi
	if [ -e $(BOARD_DIR)/kernel/kernel.config ]; then \
		cp $(BOARD_DIR)/kernel/kernel.config $@/.config; \
	elif [ -e $(FAMILY_DIR)/kernel/$(CHIP).generic.config ]; then \
		cp $(FAMILY_DIR)/kernel/$(CHIP).generic.config $@/.config; \
	fi


# ====================================================================================================================
# FAMILY

# --- TOOLCHAIN
FAMILY_TOOLCHAIN_BR_CONFIG := $(FAMILY_DIR)/toolchain.buildroot  

$(FAMILY_OUTDIR)/toolchain: $(BUILDROOT_DIR) $(FAMILY_TOOLCHAIN_BR_CONFIG)
	@echo "-- Build toolchain for $(FAMILY) family"
	mkdir -p $(FAMILY_OUTDIR)
	make -C $(BUILDROOT_DIR) O=$(FAMILY_OUTDIR)/.br-toolchain defconfig BR2_DEFCONFIG=$(FAMILY_TOOLCHAIN_BR_CONFIG)
	make -C $(FAMILY_OUTDIR)/.br-toolchain toolchain
	ln -f --symbolic $(FAMILY_OUTDIR)/.br-toolchain/host $@

# --- ROOTFS
FAMILY_ROOTFS_BR_CONFIG := $(FAMILY_OUTDIR)/rootfs.buidlroot.config

# hack for compatibility
$(FAMILY_ROOTFS_BR_CONFIG): $(FAMILY_DIR)/rootfs.buildroot
	mkdir -p $(FAMILY_OUTDIR)
	sed 's|BR2_TOOLCHAIN_EXTERNAL_PATH=.*|BR2_TOOLCHAIN_EXTERNAL_PATH="$(FAMILY_OUTDIR)/toolchain"|' $< > $@

$(FAMILY_OUTDIR)/rootfs: $(BUILDROOT_DIR) $(FAMILY_ROOTFS_BR_CONFIG) $(FAMILY_OUTDIR)/toolchain
	@echo "-- Build RootFS for $(FAMILY) family"
	mkdir -p $(FAMILY_OUTDIR)
	make -C $(BUILDROOT_DIR) O=$(FAMILY_OUTDIR)/.br-rootfs defconfig BR2_DEFCONFIG=$(FAMILY_ROOTFS_BR_CONFIG)
	make -C $(FAMILY_OUTDIR)/.br-rootfs
	ln -f --symbolic $(FAMILY_OUTDIR)/.br-rootfs/target $@

# --- KERNEL SOURCE TREE
$(FAMILY_DIR)/kernel/linux:
	make -C $(FAMILY_DIR)/kernel linux
