#!/bin/sh

CAMERAIP=192.168.10.101

start()
{
  echo "Start"
  curl  --header "Content-Type: application/json" \
        --request GET \
        "http://$CAMERAIP/api/recorder/start"
  echo ""
}

stop()
{ 
  echo "Stop"                                
  curl  --header "Content-Type: application/json" \
        --request GET \
        "http://$CAMERAIP/api/recorder/stop"
  echo ""
}

start
