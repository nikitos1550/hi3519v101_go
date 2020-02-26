SHELL := /bin/bash

ifndef BOARD
    $(error 'BOARD' variable must be defined)
endif

# TODO: It may be needed some magic here when the config becomes JSON
-include ./$(BOARD)/config

# source directories
BUILDROOT_DIR  ?= $(abspath ../buildroot-2019.08)
BOARD_DIR      := $(abspath ./$(BOARD))
FAMILY_ROOT    := $(abspath ../hisilicon)
FAMILY_DIR     := $(FAMILY_ROOT)/$(FAMILY)

# destination directories
OUTDIR         ?= $(abspath ../output/boards/$(BOARD))
FAMILY_OUTDIR  ?= $(abspath ../output/hisilicon/$(FAMILY))

# workspaces
GOPATH         ?= $(abspath ../output/go_workspace)

# artifacts
TOOLCHAIN_DIR       := $(OUTDIR)/toolchain
MAKEFILE_PARAMS     := $(OUTDIR)/Makefile.params
KERNEL_DIR          := $(OUTDIR)/kernel
KERNEL_UIMAGE       := $(OUTDIR)/kernel/uImage
ROOTFS_DIR          := $(OUTDIR)/rootfs


.PHONY: help info all toolchain rootfs kernel clean

help:
	@echo -en "Help: \n" \
		" - make help           - print this help \n" \
		" - make info           - print info of target board and artifacts \n" \
		" - make toolchain      - build toolchain & Makefile.params \n" \
		" - make rootfs         - build RootFS directory \n" \
		" - make kernel         - build Linux Kernel's uImage \n" \
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
		"  Kernel dir:       $(KERNEL_DIR)\n"\
		"  Kernel uImage:    $(KERNEL_UIMAGE)\n"\
		"  Makefile.params:  $(MAKEFILE_PARAMS)\n"


# ====================================================================================================================
all: toolchain rootfs kernel

toolchain: $(TOOLCHAIN_DIR) $(MAKEFILE_PARAMS)
	@echo "-- Toolchain for $(BOARD) is ready: $^"

rootfs: $(ROOTFS_DIR)
	@echo "-- RootFS for $(BOARD) is ready: $^"

kernel: $(KERNEL_UIMAGE)
	@echo "-- Linux Kernel for $(BOARD) is ready: $^"

clean:
	rm -rf $(TOOLCHAIN_DIR) $(MAKEFILE_PARAMS) $(KERNEL_DIR) $(ROOTFS_DIR)

# --- TOOLCHAIN ---------------------------------------------------------------------------------------------

# All sibling boards use their family's toolchain
$(TOOLCHAIN_DIR): $(FAMILY_OUTDIR)/toolchain
	mkdir -p $(OUTDIR); rm -rf $@
	ln -f --symbolic $< $@

# This params file must be used to build artifacts for chosen board
$(MAKEFILE_PARAMS): $(BOARD_DIR)/config $(FAMILY_DIR)/Makefile.params
	mkdir -p $(OUTDIR); rm -f $@
	echo -e "# THIS FILE WAS GENERATED, PLEASE DO NOT MODIFY IT\n\
	BOARD         := $(BOARD)\n\
	TOOLCHAIN_DIR := $(TOOLCHAIN_DIR)\n\
	GOPATH        := $(GOPATH)" >> $@
	echo -e "\n# Parameters of the board" >> $@; cat $(BOARD_DIR)/config >> $@;
	echo -e "\n# Parameters of the family" >> $@; cat $(FAMILY_DIR)/Makefile.params >> $@

# --- ROOTFS ------------------------------------------------------------------------------------------------

$(ROOTFS_DIR): $(FAMILY_OUTDIR)/rootfs $(BOARD_DIR)/putonrootfs $(BOARD_DIR)/config
	rm -rf $@; mkdir -p $@; cp -r $</* $@/
	#if [ -d $(BOARD_DIR)/putonrootfs ]; then cp -r $(BOARD_DIR)/putonrootfs/* $@/; fi
	if [ -d $(BOARD_DIR)/putonrootfs ]; then shopt -s nullglob; for P in $(BOARD_DIR)/putonrootfs/*; do cp -r $$P $(@) ; done; fi
	cat $(BOARD_DIR)/config | tr -d "[:blank:]" > $@/etc/board.config

# --- KERNEL ------------------------------------------------------------------------------------------------

$(KERNEL_UIMAGE): $(KERNEL_DIR) $(MAKEFILE_PARAMS) $(TOOLCHAIN_DIR)
	make -C $(KERNEL_DIR) MAKEFILE_PARAMS=$(MAKEFILE_PARAMS) build

$(KERNEL_DIR): $(FAMILY_OUTDIR)/kernel
	rm -rf $@; mkdir -p $@; cp -r $</* $@/
	if [ -e $(BOARD_DIR)/kernel/patch ]; then \
		cp -r $(BOARD_DIR)/kernel/patch/* $@/linux; \
	fi
	if [ -e $(BOARD_DIR)/kernel/kernel.config ]; then \
		cp $(BOARD_DIR)/kernel/kernel.config $@/linux/.config; \
	elif [ -e $(FAMILY_DIR)/kernel/$(CHIP).generic.config ]; then \
		cp $(FAMILY_DIR)/kernel/$(CHIP).generic.config $@/linux/.config; \
	fi

# --- FAMILY ------------------------------------------------------------------------------------------------

$(FAMILY_OUTDIR)/toolchain:
	make -C $(FAMILY_ROOT) FAMILY=$(FAMILY) toolchain

$(FAMILY_OUTDIR)/rootfs:
	make -C $(FAMILY_ROOT) FAMILY=$(FAMILY) rootfs

$(FAMILY_OUTDIR)/kernel:
	make -C $(FAMILY_ROOT) FAMILY=$(FAMILY) kernel
