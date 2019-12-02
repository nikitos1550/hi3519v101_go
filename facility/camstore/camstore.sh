#!/bin/bash

DIR=`dirname $0`
MAIN_PY=$DIR/__main__.py
PID_FILE=$DIR/.camstore.pid
DAEMON_PORT=43500


function start() {
    $MAIN_PY --port $DAEMON_PORT --detach --pidf $PID_FILE
}

function stop() {
    if [ ! -f $PID_FILE ]; then
        echo "PID file is absent" >&2
        return 0
    fi

    PID=`cat $PID_FILE`
    echo "process PID=$PID"

    if kill $PID 2>/dev/null; then
        sleep 3
    else
        echo "process with PID=$PID is not running" >&2
        rm -f $PID_FILE
        return 0
    fi

    if kill -0 $PID 2>/dev/null; then
        echo "kill $PID process with SIGKILL" >&2
        kill -9 $PID
    else
        echo "process with PID=$PID terminated" >&2
    fi

    rm -f $PID_FILE
}

CMD=$1
shift

$CMD
