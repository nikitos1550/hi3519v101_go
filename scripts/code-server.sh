#!/bin/sh

# Script for code-server deploy automation

URL=https://github.com/cdr/code-server/releases/download/2.1692-vsc1.39.2/code-server2.1692-vsc1.39.2-linux-x86_64.tar.gz
FILE=code-server2.1692-vsc1.39.2-linux-x86_64.tar.gz
#NAME=`basename $FILE .tar.gz`

start () {
    if [ -S ~/code-server/code-server.sock ]
    then
        echo "Seems code-server is already running. Try restart."
    else
        nohup ~/code-server/code-server \
            --log off \
            --auth password \
            --base-path /$USER/vs/ \
            --disable-telemetry \
            --socket ~/code-server/code-server.sock > ~/code-server/nohup.out 2>&1 &
        sleep 2
        chmod 777 ~/code-server/code-server.sock
        echo "Code-server started. You can access it via http://213.141.129.12:8080/$USER/vs/"
        grep "Password" ~/code-server/nohup.out
    fi
}

stop () {
    killall -u $USER code-server
    rm -f ~/code-server/code-server.sock
}

case "$1" in
  start)
    start
    ;;
  stop)
    stop
    ;;
  restart)
   stop
   start
   ;;
  status)
    echo "TODO"
    grep "Password" ~/code-server/nohup.out
    #PIDs=`ps aux | grep code-server | grep if | cut -b 1-5`
    PIDs=`ps -u $USER | grep code-server | cut -b 1-5`

    echo $PIDs
    ;;
  install)
    if [ -d ~/code-server ]
    then
        echo "Seems already installed."
    else
        #echo "dir NOT found"
        if [ -f ~/$FILE ]
        then
            echo "$FILE already downloaded."
        else
            #echo "$FILE not found"
            wget -O ~/$FILE $URL
        fi
        mkdir ~/code-server
        tar -zxvf ~/$FILE --strip 1 -C ~/code-server
    fi
    ;;
  *)
    echo "Usage: $0 {start|stop|restart|status|install}"
    exit 1
esac

exit $?

