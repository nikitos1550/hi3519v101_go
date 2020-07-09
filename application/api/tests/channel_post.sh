#!/bin/sh

CAMERAIP=192.168.10.105

create_channel()
{
  echo "Creating channel $1 with params $2"
  curl  --header "Content-Type: application/json" \
        --request POST \
        --data "@$2" \
        "http://$CAMERAIP/api/mpp/channels/$1"
    echo ""
}

delete_channel()
{
  echo "Deleting channel $1"
  curl  --header "Content-Type: application/json" \
        --request DELETE \
        "http://$CAMERAIP/api/mpp/channels/$1"
  echo ""
}

delete_all_channels()
{
    delete_channel 0
    delete_channel 1
    delete_channel 2
    delete_channel 3
    delete_channel 4
}

create_channel 0 c_3840x2160.json
create_channel 1 c_1920x1080.json
create_channel 2 c_1280x720.json
create_channel 3 c_640x360.json

delete_all_channels

create_channel 0 c_640x360.json
create_channel 1 c_640x360.json
create_channel 2 c_640x360.json
create_channel 3 c_640x360.json

delete_all_channels

create_channel 0 c_3840x2160_d1.json
create_channel 1 c_3840x2160_d1.json
create_channel 2 c_3840x2160_d1.json
create_channel 3 c_3840x2160_d1.json
