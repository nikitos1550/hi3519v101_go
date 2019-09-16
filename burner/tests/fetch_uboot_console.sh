#!/bin/bash

BURNER_DIR=$(realpath $(dirname $0)/..)
DEVICE_PY=$BURNER_DIR/device.py
POWER_PORT=/dev/ttyACM0
POWER_BR=115200

if [ $# != 2 ]; then
    echo -e \
"Fetch UBoot console on a given camera
Usage:    $0 <camera-id> <camera-port>"
    exit 1
fi

CAM_ID=$1
CAM_PORT=$2

# reset power
$DEVICE_PY --port $POWER_PORT --br $POWER_BR write --wait-for "READY" --data "reset $CAM_ID"

# stop kernel loading
$DEVICE_PY --port $CAM_PORT write --wait-for "System startup" --data "A"

# open terminal =)
$DEVICE_PY --port $CAM_PORT
