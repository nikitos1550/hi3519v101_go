#!/bin/sh

. /opt/board.config
#dmesg -n 1
#/opt/application
/opt/gohisicam -mem-total=$RAM_SIZE -mem-linux=$RAM_LINUX_SIZE -mem-mpp=$RAM_MPP_SIZE &

