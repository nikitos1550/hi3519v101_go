#!/bin/sh

# exit when any command fails
set -e

for U in `ls -d boards/unknown_unknown_*_unknown` 
do
	BOARD=`echo $U | cut -d/ -f2`
       	echo "================$BOARD================"
	make BOARD=$BOARD kernel
	make BOARD=$BOARD build-rootfs
done

echo "================COMPLETE================"
