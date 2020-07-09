#!/bin/sh

CAMERAIP=192.168.10.105
#CAMERAIP=192.168.10.131

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

create_encoder()
{
    echo "Creating encoder with params $1" 
    curl  --header "Content-Type: application/json" \
          --request POST \
          --data "@$1" \
          "http://$CAMERAIP/api/mpp/encoders"
    echo ""
}

create_encoder_id()
{
  echo "Creating encoder $1 with params $2"
  curl  --header "Content-Type: application/json" \
        --request POST \
        --data "@$2" \
        "http://$CAMERAIP/api/mpp/encoders/$1"
    echo ""
}

delete_encoder()
{
  echo "Deleting channel $1"
  curl  --header "Content-Type: application/json" \
        --request DELETE \
        "http://$CAMERAIP/api/mpp/encoders/$1"
  echo ""
}

bind_encoder_channel()
{
  echo "Bind encoder $1 to channel $2"
  curl  --header "Content-Type: application/json" \
        --request GET \
        "http://$CAMERAIP/api/mpp/channels/$2/bind/encoder/$1"
  echo ""
}

unbind_encoder_channel()
{
  echo "UnBind encoder $1 from channel $2"
  curl  --header "Content-Type: application/json" \
        --request GET \
        "http://$CAMERAIP/api/mpp/channels/$2/unbind/encoder/$1"
  echo ""
}

encoder_start()
{
  echo "Start encoder $1"
  curl  --header "Content-Type: application/json" \
        --request GET \
        "http://$CAMERAIP/api/mpp/encoders/$1/start"
  echo ""
}

encoder_stop()
{
  echo "Stop encoder $1" 
  curl  --header "Content-Type: application/json" \
        --request GET \
        "http://$CAMERAIP/api/mpp/encoders/$1/stop" 
  echo ""
} 

create_jpeg()
{
    echo    "Creating jpeg streamer named $1"
    curl    --header "Content-Type: application/json" \
            --request POST \
            "http://$CAMERAIP/api/jpeg/$1" 
    echo    ""
}

delete_jpeg()
{
    echo    "Delete jpeg streamer $1"
    curl    --header "Content-Type: application/json" \
            --request DELETE \
            "http://$CAMERAIP/api/jpeg/$1"
    echo    ""
}

bind_encoder_jpeg()
{
  echo "Bind encoder $1 to jpeg $2"
  curl  --header "Content-Type: application/json" \
        --request GET \
        "http://$CAMERAIP/api/mpp/encoders/$1/bind/jpeg/$2"
  echo ""
}

unbind_encoder_jpeg()                               
{
  echo "Unbind encoder $1 from jpeg $2"
  curl  --header "Content-Type: application/json" \
        --request GET \
        "http://$CAMERAIP/api/mpp/encoders/$1/unbind/jpeg/$2"
  echo ""
}


delete_all_encoders()
{
    delete_encoder 0
    delete_encoder 1
    delete_encoder 2
    delete_encoder 3
    delete_encoder 4
    delete_encoder 5
    delete_encoder 6
    delete_encoder 7
    delete_encoder 8
    delete_encoder 9
    delete_encoder 10
    delete_encoder 11
    delete_encoder 12
    delete_encoder 13
    delete_encoder 14
    delete_encoder 15
    delete_encoder 16
    delete_encoder 17
    delete_encoder 18
    delete_encoder 19
    delete_encoder 20
}

#delete_all_encoders
#delete_all_channels

#create_channel 0 c_3840x2160.json
#create_channel 1 c_1920x1080.json
#create_channel 2 c_1280x720.json
#create_channel 3 c_640x360.json

#create_encoder_id 0 e_3840x2160_h264_cbr.json
#create_encoder_id 1 e_1920x1080_h264_cbr.json
#create_encoder_id 2 e_1280x720_h264_cbr.json
#create_encoder_id 3 e_640x360_h264_cbr.json

#bind_encoder_channel 0 0
#bind_encoder_channel 1 1
#bind_encoder_channel 2 2
#bind_encoder_channel 3 3

#encoder_start 0
#encoder_start 1
#encoder_start 2
#encoder_start 3

#create_encoder_id 4 e_3840x2160_mjpeg_cbr.json 
#create_encoder_id 5 e_1920x1080_mjpeg_cbr.json 
#create_encoder_id 6 e_1280x720_mjpeg_cbr.json 
#create_encoder_id 7 e_640x360_mjpeg_cbr.json 

#bind_encoder_channel 4 0
#bind_encoder_channel 5 1
#bind_encoder_channel 6 2
#bind_encoder_channel 7 3

#unbind_encoder 0

#delete_all_encoders
#delete_all_channels

test_0()
{
    create_channel 1 c_1920x1080.json
    create_encoder_id 0 e_1920x1080_h264_cbr.json

    bind_encoder_channel 0 1

    encoder_start 0

    encoder_stop 0

    unbind_encoder 0 1
}

test_1()
{
    create_channel 1 c_640x360.json
    
    create_encoder_id 0 e_640x360_h264_cbr.json
    create_encoder_id 1 e_640x360_h264_cbr.json
    create_encoder_id 2 e_640x360_h264_cbr.json
    create_encoder_id 3 e_640x360_h264_cbr.json

    bind_encoder_channel 0 1
    bind_encoder_channel 1 1
    bind_encoder_channel 2 1
    bind_encoder_channel 3 1

    encoder_start 0
    encoder_start 1
    encoder_start 2
    encoder_start 3
}

test_2()
{
    create_channel 1 c_1920x1080.json
    
    create_encoder_id 0 e_1920x1080_h264_cbr.json
    create_encoder_id 1 e_1920x1080_h264_cbr.json
    #create_encoder_id 2 e_1920x1080_h264_cbr.json
    #create_encoder_id 3 e_1920x1080_h264_cbr.json
    #create_encoder_id 4 e_1920x1080_h264_cbr.json
    #create_encoder_id 5 e_1920x1080_h264_cbr.json
    #create_encoder_id 6 e_1920x1080_h264_cbr.json
    #create_encoder_id 7 e_1920x1080_h264_cbr.json

    bind_encoder_channel 0 1
    bind_encoder_channel 1 1
    #bind_encoder_channel 2 1
    #bind_encoder_channel 3 1
    #bind_encoder_channel 4 1
    #bind_encoder_channel 5 1
    #bind_encoder_channel 6 1
    #bind_encoder_channel 7 1

    encoder_start 0
    encoder_start 1
    #encoder_start 2
    #encoder_start 3
    #encoder_start 4
    #encoder_start 5
    #encoder_start 6
    #encoder_start 7
}

test_3()
{
    create_channel 0 c_3840x2160.json
    create_channel 1 c_1920x1080.json

    create_encoder_id 0 e_3840x2160_h264_cbr.json
    create_encoder_id 1 e_3840x2160_mjpeg_cbr.json
    create_encoder_id 2 e_1920x1080_h264_cbr.json

    bind_encoder_channel 0 0
    bind_encoder_channel 1 0
    bind_encoder_channel 2 1

    encoder_start 0
    encoder_start 1
    encoder_start 2
}

test_4()
{
    create_channel 0 c_3840x2160.json
    create_channel 1 c_1920x1080.json
    create_channel 2 c_1280x720.json
    create_channel 3 c_640x360.json
    
    create_encoder_id 0 e_3840x2160_h264_cbr.json
    create_encoder_id 1 e_1920x1080_h264_cbr.json
    create_encoder_id 2 e_1280x720_h264_cbr.json
    create_encoder_id 3 e_640x360_h264_cbr.json
    
    bind_encoder_channel 0 0   
    bind_encoder_channel 1 1
    bind_encoder_channel 2 2
    bind_encoder_channel 3 3
    
    encoder_start 0   
    encoder_start 1   
    encoder_start 2
    encoder_start 3
    
    #sleep 1

    encoder_stop 0    
    encoder_stop 1    
    encoder_stop 2 
    encoder_stop 3 

    unbind_encoder 0
    unbind_encoder 1
    unbind_encoder 2
    unbind_encoder 3

    delete_all_encoders
    delete_all_channels
}

test_5()
{
    create_channel 1 c_1280x720.json
    
    create_encoder_id 0 e_1280x720_h264_cbr.json
    create_encoder_id 1 e_1280x720_h264_cbr.json
    create_encoder_id 2 e_1280x720_h264_cbr.json
    create_encoder_id 3 e_1280x720_h264_cbr.json
    create_encoder_id 4 e_1280x720_h264_cbr.json
    create_encoder_id 5 e_1280x720_h264_cbr.json
    create_encoder_id 6 e_1280x720_h264_cbr.json 
    create_encoder_id 7 e_1280x720_h264_cbr.json 
    create_encoder_id 8 e_1280x720_h264_cbr.json 
    create_encoder_id 9 e_1280x720_h264_cbr.json 
    create_encoder_id 10 e_1280x720_h264_cbr.json 
    create_encoder_id 11 e_1280x720_h264_cbr.json
    create_encoder_id 12 e_1280x720_h264_cbr.json 
    create_encoder_id 13 e_1280x720_h264_cbr.json 
    create_encoder_id 14 e_1280x720_h264_cbr.json 
    create_encoder_id 15 e_1280x720_h264_cbr.json 

    bind_encoder_channel 0 1
    bind_encoder_channel 1 1
    bind_encoder_channel 2 1
    bind_encoder_channel 3 1
    bind_encoder_channel 4 1
    bind_encoder_channel 5 1
    bind_encoder_channel 6 1
    bind_encoder_channel 7 1
    bind_encoder_channel 8 1
    bind_encoder_channel 9 1
    bind_encoder_channel 10 1
    bind_encoder_channel 11 1
    bind_encoder_channel 12 1
    bind_encoder_channel 13 1
    bind_encoder_channel 14 1 
    bind_encoder_channel 15 1

    encoder_start 0
    encoder_start 1
    encoder_start 2
    encoder_start 3
    encoder_start 4
    encoder_start 5
    encoder_start 6
    encoder_start 7
    encoder_start 8
    encoder_start 9
    encoder_start 10
    encoder_start 11
    encoder_start 12
    encoder_start 13
    encoder_start 14
    encoder_start 15
}

test_6()
{
    create_channel 0 c_2592x1944.json
    create_encoder_id 0 e_2592x1944_h264_cbr.json
    bind_encoder_channel 0 0
    encoder_start 0
}

test_7()
{
    create_channel 1 c_1920x1080.json
    create_encoder_id 0 e_1920x1080_mjpeg_cbr.json

    bind_encoder_channel 0 1

    encoder_start 0

    create_jpeg sample1
    create_jpeg test2

    bind_encoder_jpeg 0 0
    bind_encoder_jpeg 0 1

    #sleep 10

    #unbind_encoder_jpeg 0 0
    #unbind_encoder_jpeg 0 1

    #delete_jpeg 0
    #delete_jpeg 1

    #unbind_encoder_channel 0 1
    

    #encoder_stop 0
   
    #delete_encoder 0

    #delete_channel 1
}

delete_all_encoders
delete_all_channels

#test_0

#test_1
#test_2
#test_3

#while :
#do
#	test_4
#done

#test_5
#test_6
test_7
