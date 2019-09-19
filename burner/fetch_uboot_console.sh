#!/bin/bash

BURNER_DIR=$(realpath $(dirname $0)/..)
DEVICE_PY=$BURNER_DIR/device.py
UBOOT_CONSOLE_PY=$BURNER_DIR/uboot_console.py
POWER_PORT=/dev/ttyACM0
POWER_BR=115200

if [ $# != 1 ]; then
    echo -e \
"Reset power and fetch UBoot console on a given camera
Usage:    $0 <camera-id>"
    exit 1
fi

CAM_ID=$1
CAM_PORT=/dev/ttyCAM$CAM_ID

# reset power
$DEVICE_PY --port $POWER_PORT --br $POWER_BR write --wait-for "READY" --data "reset $CAM_ID"

# stop kernel loading
$DEVICE_PY --port $CAM_PORT write --wait-for "System startup" --data "X"

# open terminal =)
$DEVICE_PY --port $CAM_PORT
