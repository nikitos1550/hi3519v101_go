#!/bin/sh

. /opt/board.config
dmesg -n 1

#/opt/gohisiprobe \
#    -mem-total=$RAM_SIZE \
#    -mem-linux=$RAM_LINUX_SIZE \
#    -mem-mpp=$RAM_MPP_SIZE \
#    &

echo $CHIP

/opt/gohisicam -mem-total=$RAM_SIZE -mem-linux=$RAM_LINUX_SIZE -mem-mpp=$RAM_MPP_SIZE -chip=$CHIP &
