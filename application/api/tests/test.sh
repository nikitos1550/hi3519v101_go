#!/bin/sh

CAMERAIP=192.168.10.105
#CAMERAIP=192.168.10.131

link()
{
  echo "Link $1 instance $2 to $3 instance $4"
  curl  --header "Content-Type: application/json" \
        --request POST \
        "http://$CAMERAIP/api/link/$1/$2/$3/$4"
  echo ""
}

unlink()
{
  echo "Unlink $1 instance $2 to $3 instance $4"
  curl  --header "Content-Type: application/json" \
        --request DELETE \
        "http://$CAMERAIP/api/link/$1/$2/$3/$4"
  echo ""
}

linkraw()
{
  echo "Link $1 instance $2 to $3 instance $4"
  curl  --header "Content-Type: application/json" \
        --request POST \
        "http://$CAMERAIP/api/link/$1/$2/$3/$4/raw"
  echo ""
}

unlinkraw()
{
  echo "Unlink $1 instance $2 to $3 instance $4"
  curl  --header "Content-Type: application/json" \
        --request DELETE \
        "http://$CAMERAIP/api/link/$1/$2/$3/$4/raw"
  echo ""
}

delete()
{
  echo "Delete $1 instance $2"
  curl  --header "Content-Type: application/json" \
        --request DELETE \
        "http://$CAMERAIP/api/$1/$2"
  echo ""
}

create_channel()
{
  echo "Create channel named $1 with params $2"
  curl  --header "Content-Type: application/json" \
        --request POST \
        --data "@$2" \
        "http://$CAMERAIP/api/channel/$1"
    echo ""
}

create_encoder()
{
    echo "Create encoder named $1 with params $2" 
    curl  --header "Content-Type: application/json" \
          --request POST \
          --data "@$2" \
          "http://$CAMERAIP/api/encoder/$1"
    echo ""
}

encoder_start()
{
  echo "Start encoder $1"
  curl  --header "Content-Type: application/json" \
        --request GET \
        "http://$CAMERAIP/api/encoder/$1/start"
  echo ""
}

encoder_stop()
{
    echo "Stop encoder $1" 
    curl  --header "Content-Type: application/json" \
          --request GET \
          "http://$CAMERAIP/api/encoder/$1/stop" 
    echo ""
} 

create_jpeg()
{
    echo    "Create jpeg streamer named $1"
    curl    --header "Content-Type: application/json" \
            --request POST \
            "http://$CAMERAIP/api/jpeg/$1" 
    echo    ""
}

create_mjpeg()
{
    echo    "Create mjpeg streamer named $1"
    curl    --header "Content-Type: application/json" \
            --request POST \
            "http://$CAMERAIP/api/mjpeg/$1"
    echo    ""
}

create_webrtc()
{
    echo    "Creating webrtc streamer"
    curl    --header "Content-Type: application/json" \
            --request POST \
            "http://$CAMERAIP/api/webrtc"
    echo    ""
}

test()
{
    create_channel main c_3840x2160.json
    create_channel fullhd c_1920x1080.json

    create_encoder h264_1 e_1920x1080_h264_cbr.json
    create_encoder mjpeg_1 e_1920x1080_mjpeg_cbr.json

    link channel fullhd encoder h264_1
    link channel fullhd encoder mjpeg_1

    encoder_start h264_1
    encoder_start mjpeg_1

    create_jpeg fullhd
    create_mjpeg fullhd

    link encoder mjpeg_1 jpeg fullhd
    link encoder mjpeg_1 mjpeg fullhd

}

test
