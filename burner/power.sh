#!/bin/bash

if [ $# != 2 ]; then
    echo "Control power on a device
Usage: $0 {reset|on|off} <device ID>" >&2
    exit 1
fi

BURNER_DIR=$(realpath $(dirname $0))
$BURNER_DIR/device.py --verbose --port /dev/ttyACM0 --br 115200 write --wait-for "READY" --data "$1 $2"
