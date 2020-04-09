#!/bin/bash

PROJECT_ROOT=`cd $(dirname $0)/.. && pwd`


function get_telnet_endpoint() {
    PORT=$1

    RESP=`/var/lib/buildbot/camstore/client.py --quite --no-exec forward_serial $PORT`
    if [ $? != 0 ]; then
        echo "Couldn't forward serial $PORT" >&2
        exit -1
    fi
    TELNET_ENDPOINT=`echo $RESP | awk '{print $2 ":" $3}'`
}


NUM=$1
shift

if [[ $NUM == "" ]]; then
    echo "Wrap around hiburn to use camstore daemon" >&2
    echo "Usage: $0 <CAM_NUMBER> [hiburn args...]" >&2
    exit
fi


POWER2_PATH="$PROJECT_ROOT/burner/power2.py"
HIBURN_APP_PATH="$PROJECT_ROOT/hiburn/hiburn_app.py"

get_telnet_endpoint /dev/ttyCAM$NUM
echo "Serial-over-telnet endpoint $TELNET_ENDPOINT" >&2


authbind --deep $HIBURN_APP_PATH \
    --reset-cmd "$POWER2_PATH --num $NUM reset" \
    --serial-over-telnet $TELNET_ENDPOINT \
    "$@"
