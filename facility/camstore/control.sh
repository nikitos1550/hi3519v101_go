#!/bin/bash

DIR=`dirname $0`
MAIN_PY=$DIR/__main__.py
PID_FILE=$DIR/.camstore.pid
DAEMON_PORT=43500
DAEMON_HOME=/var/lib/buildbot/camstore

# -------------------------------------------------------------------------------------------------

function deploy() {
    echo "Stop running daemon..."
    $DAEMON_HOME/control.sh stop_daemon

    echo "Install itself..."
    install

    echo "Run installed daemon..."
    $DAEMON_HOME/control.sh start_daemon

    sleep 1
    echo "Check client..."
    $DAEMON_HOME/client.py sysinfo
    $DAEMON_HOME/client.py help
}


# Copy local camstore (both daemon and client) to its' home directory
function install() {
    rm -rf $DAEMON_HOME/*
    mkdir -p $DAEMON_HOME; cp $DIR/*.{py,sh} $DAEMON_HOME/
    mkdir -p $DAEMON_HOME/lib; cp $DIR/lib/*.py $DAEMON_HOME/lib
    chmod 755 $DAEMON_HOME/*.{py,sh}
}


function start_daemon() {
    $DIR/daemon.py --port $DAEMON_PORT --pidf $PID_FILE --detach
}


function stop_daemon() {
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


function client() {
    $DIR/client.py --port $DAEMON_PORT "$@"
}

CMD=$1
if [[ $CMD == "" ]]; then
    echo "Usage: $0 {deploy|start_daemon|stop_daemon|client}" >&2
    exit 1
fi

shift
$CMD "$@"
