#!/bin/sh

. /opt/board.config
#dmesg -n 1
#/opt/application
/opt/application -mem-total=$RAM_SIZE -mem-linux=$RAM_LINUX_SIZE -mem-mpp=$RAM_MPP_SIZE &

