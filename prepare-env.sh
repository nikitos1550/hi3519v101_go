#!/bin/sh


make enviroiment-setup

make FAMILY=hi3516cv100 toolchain
make FAMILY=hi3516cv100 rootfs

make FAMILY=hi3516cv200 toolchain
make FAMILY=hi3516cv200 rootfs

make FAMILY=hi3516cv300 toolchain
make FAMILY=hi3516cv300 rootfs

make FAMILY=hi3516cv500 toolchain
make FAMILY=hi3516cv500 rootfs

make FAMILY=hi3516av100 toolchain
make FAMILY=hi3516av100 rootfs

make FAMILY=hi3516av200 toolchain
make FAMILY=hi3516av200 rootfs

