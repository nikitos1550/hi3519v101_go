#!/bin/bash

TEST_DIR=`cd $(dirname $0) && pwd`
$TEST_DIR/core/checks.py "$@"
