#!/bin/bash

BOARD=$1
BR_HISICAM_OUTDIR=$(readlink -f $2)
BR_HISICAM_ROOT=$(readlink -f $3)



function get_br_vars() {
    eval $(make -s --no-print-directory -C $BR_HISICAM_ROOT VARS="$*" br-printvars)
}

get_br_vars GO_GOARCH GO_GOARM GO_VERSION


cat - << EOF
BOARD := $BOARD
STAGING_DIR := $BR_HISICAM_OUTDIR/staging

# toolchain common variables are defined here
include $BR_HISICAM_OUTDIR/toolchain-params.mk

# go specific variables
GOARCH := $GO_GOARCH
GOARM := $GO_GOARM
GOVERSION := $GO_VERSION

# board parameters
EOF
make -C $BR_HISICAM_ROOT -s --no-print-directory BOARD="$BOARD" board-info
