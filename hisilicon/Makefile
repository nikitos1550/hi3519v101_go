SHELL := /bin/bash
ifndef FAMILY
    $(error 'FAMILY' variable must be defined)
endif

ROOTFS_CONFIG_DIR := $(abspath ../rootfs)
BUILDROOT_DIR  	  ?= $(abspath ../buildroot-2020.02.3)
FAMILY_DIR        := $(abspath $(FAMILY))
OUTDIR            ?= $(abspath ../output/hisilicon/$(FAMILY))

TOOLCHAIN_DIR       := $(OUTDIR)/toolchain
MAKEFILE_PARAMS     := $(OUTDIR)/Makefile.params
KERNEL_SOURCE_DIR   := $(OUTDIR)/kernel/linux
KERNEL_MAKEFILE     := $(OUTDIR)/kernel/Makefile
ROOTFS_DIR          := $(OUTDIR)/rootfs


.PHONY: help info toolchain rootfs kernel clean all

# -------------------------------------------------------------------------------------------------
help:
	@echo -e "Help:\n"\
		" - make help           - print this help\n"\
		" - make info           - print info\n"\
		" - make all            - toolchain & rootfs & kernel\n"\
		" - make toolchain      - build toolchain & Makefile.params\n"\
		" - make rootfs         - build RootFS directory\n"\
		" - make kernel         - prepare Kernel's source tree & Makefile to build uImage\n"\
		" - make clean          - clean output directory"

info:
	@echo -e ""\
		"-- Sources:\n"\
		"  Buildroot:        $(BUILDROOT_DIR)\n"\
		"-- Destinations:\n"\
		"  Toolchain:        $(TOOLCHAIN_DIR)\n"\
		"  Makefile.params:  $(MAKEFILE_PARAMS)\n"\
		"  RootFS:           $(ROOTFS_DIR)\n"\
		"  Kernel sources:   $(KERNEL_SOURCE_DIR)\n"\
		"  Kernel Makefile:  $(KERNEL_MAKEFILE)"

all: toolchain rootfs kernel

toolchain: $(TOOLCHAIN_DIR) $(MAKEFILE_PARAMS)
	@echo "-- Toolchain for $(FAMILY) is ready: $^"

rootfs: $(ROOTFS_DIR)
	@echo "-- RootFS for $(FAMILY) is ready: $^"

kernel: $(KERNEL_SOURCE_DIR) $(KERNEL_MAKEFILE)
	@echo "-- Linux Kernel for $(FAMILY) is prepared: $^"

clean:
	rm -rf $(TOOLCHAIN_DIR) $(MAKEFILE_PARAMS) $(ROOTFS_DIR) $(KERNEL_SOURCE_DIR) $(KERNEL_MAKEFILE)


# --- TOOLCHAIN -----------------------------------------------------------------------------------
TOOLCHAIN_BR_CONFIG := $(FAMILY_DIR)/toolchain.buildroot  

$(TOOLCHAIN_DIR): $(BUILDROOT_DIR) $(TOOLCHAIN_BR_CONFIG)
	mkdir -p $(OUTDIR)
	make -C $(BUILDROOT_DIR) O=$(OUTDIR)/.br-toolchain defconfig BR2_DEFCONFIG=$(TOOLCHAIN_BR_CONFIG)
	make -C $(OUTDIR)/.br-toolchain toolchain
	ln -f --symbolic $(OUTDIR)/.br-toolchain/host $@

$(MAKEFILE_PARAMS): $(FAMILY_DIR)/Makefile.params
	mkdir -p $(OUTDIR); rm -f $@
	echo -e "TOOLCHAIN_DIR := $(TOOLCHAIN_DIR)\n" >> $@
	cat $< >> $@


# --- ROOTFS --------------------------------------------------------------------------------------
ROOTFS_BR_CONFIG := $(OUTDIR)/.br-rootfs.config

# TODO: it's dirty here, fix somehow
$(ROOTFS_BR_CONFIG): $(FAMILY_DIR)/rootfs-toolchain.buildroot
	mkdir -p $(OUTDIR)
	sed 's|BR2_TOOLCHAIN_EXTERNAL_PATH=.*|BR2_TOOLCHAIN_EXTERNAL_PATH="$(TOOLCHAIN_DIR)"|' $< > $@
	cat $(ROOTFS_CONFIG_DIR)/rootfs-common.buildroot >> $@

$(ROOTFS_DIR): $(BUILDROOT_DIR) $(ROOTFS_BR_CONFIG) $(TOOLCHAIN_DIR)
	rm -rf $@; mkdir -p $(OUTDIR)
	make -C $(BUILDROOT_DIR) O=$(OUTDIR)/.br-rootfs defconfig BR2_DEFCONFIG=$(ROOTFS_BR_CONFIG)
	#shopt -s nullglob; for P in $(ROOTFS_CONFIG_DIR)/overlay-common/*; do cp -r $$P $(OUTDIR)/.br-rootfs ; done
	make -C $(OUTDIR)/.br-rootfs
	ln -f --symbolic $(OUTDIR)/.br-rootfs/target $@


# --- KERNEL SOURCES ------------------------------------------------------------------------------
$(KERNEL_SOURCE_DIR): $(FAMILY_DIR)/kernel/linux.tar.gz
	rm -rf $@; mkdir -p $@
	tar -C $@ --strip-components=1 -xf $<
	cd $@; shopt -s nullglob; for P in $(FAMILY_DIR)/kernel/patches/*.patch; do patch -p1 < $$P; done
	shopt -s nullglob; for P in $(FAMILY_DIR)/kernel/overlay/*; do cp -r $$P $(@) ; done
	#below actual only for hi3559av100
	#test -f $(FAMILY_DIR)/kernel/arm-trusted-firmware.tgz && mkdir -p $(OUTDIR)/kernel/arm-trusted-firmware
	#test -f $(FAMILY_DIR)/kernel/arm-trusted-firmware.tgz && tar -C  $(OUTDIR)/kernel/arm-trusted-firmware --strip-components=1 -xf $(FAMILY_DIR)/kernel/arm-trusted-firmware.tgz

$(KERNEL_MAKEFILE): ./Makefile.kernel
	mkdir -p `dirname $@`
	ln --symbolic --relative --force $< $@
