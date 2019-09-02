#!/bin/bash

CAM=$1
DEVICE_PY=./device.py
POWER_PORT=/dev/ttyACM0
POWER_BR=115200

$DEVICE_PY --port $POWER_PORT --br $POWER_BR write --wait-for "READY" --data "reset $CAM"