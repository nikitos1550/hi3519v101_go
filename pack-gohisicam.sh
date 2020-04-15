#!/bin/sh

# exit when any command fails
set -e

#DATE=`date --utc "+%Y.%m.%d"`

rm -rf ./output/gohisicam
mkdir -p output/gohisicam
#for U in hi3516av100 hi3516av200 hi3516cv100 hi3516cv200 hi3516cv300 hi3516cv500 hi3516ev200
for U in hi3516av200
do
       	echo "================${U}================"
	make BOARD=unknown_unknown_${U}_unknown APP_TARGET=cam build-app
	cp ./application/distrib/${U}/opt/gohisicam ./output/gohisicam/gohisicam-${U} 
	#-${DATE}
done

echo "================COMPLETE================"
