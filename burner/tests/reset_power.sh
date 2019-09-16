#!/bin/bash

BURNER_DIR=$(realpath $(dirname $0)/..)
DEVICE_PY=$BURNER_DIR/device.py
POWER_PORT=/dev/ttyACM0
POWER_BR=115200

if [ $# != 1 ]; then
    echo -e \
"Reset power on a given camera
Usage:    $0 <camera-id>"
    exit 1
fi

CAM=$1
$DEVICE_PY --port $POWER_PORT --br $POWER_BR write --wait-for "READY" --data "reset $CAM"
