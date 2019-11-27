#!/bin/bash
LOCKS_DIR=/tmp/.hisi/camstore
mkdir -p $LOCKS_DIR

function err() {
    if [ -t 2 ]; then
        echo -e "\033[0;31m$@\033[0m" >&2
    else
        echo -e "$@" >&2
    fi
}

function ok() {
    if [ -t 2 ]; then
        echo -e "\033[0;32m$@\033[0m" >&2
    else
        echo -e "$@" >&2
    fi
}

function state() {
    LOCKFILE=$LOCKS_DIR/${1//'/'/'.'}
    if [ -f $LOCKFILE ]; then
        echo -n "L(`stat --format=%U $LOCKFILE`)"
    else
        echo -n "F"
    fi
}

function list() {
    for F in /dev/ttyCAM*; do
        state $F
        echo " $F"
    done
}

function acquire() {
    if [ ! $1 ]; then
        err "failed: argument required"
        return -1
    elif [ -f $1 ]; then
        err "failed: device $1 doesn't exist"
        return -1
    fi

    TMPFILE=`mktemp`
    echo $TMPFILE > $TMPFILE

    LOCKFILE=$LOCKS_DIR/${1//'/'/'.'}
    if link $TMPFILE $LOCKFILE 2>/dev/null; then
        ok "device $1 locked"
    else
        err "failed: device $1 is already locked"
        rm $TMPFILE
        return -1
    fi
}

function release() {
    if [ ! $1 ]; then
        err "failed: argument required"
        return -1
    fi

    if [ $1 == "ALL" ]; then
        for DEV in $(locks); do release $DEV; done
    else
        LOCKFILE=$LOCKS_DIR/${1//'/'/'.'}
        if [ ! -f $LOCKFILE ]; then
            err "failed: device $1 wasn't locked"
            return -1
        else
            rm -f `cat $LOCKFILE` $LOCKFILE
            ok "device $1 released"
        fi
    fi
}

function locks() {
    find $LOCKS_DIR -type f -user `whoami` -printf '%f\n' | tr '.' '/'
}

function help() {
    echo -e "" \
        "Usage:  $0 <command>\n" \
        "Available commands:\n" \
        "    help                     - print this help\n" \
        "    list                     - list all available devices (F - free, L - locked)\n" \
        "    locks                    - list devices locked by you (`whoami`)\n" \
        "    acquire <cam-dev>        - lock device\n" \
        "    release {<cam-dev>|ALL}  - release device(s)"
}

# -------------------------------------------------------------------------------------------------

CMD=$1
shift

${CMD:-help} "$@"
