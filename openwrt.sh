#!/bin/sh

BASEPATH=/home/nikitos1550/openwrt/wrt-hisicam/openwrt/build_dir/target-arm_cortex-a7_glibc_eabi/linux-hi3516av200_hi3519v101

make KERNEL=$BASEPATH/uImage ROOTFS=$BASEPATH/root.squashfs deploy-external-control-uart
