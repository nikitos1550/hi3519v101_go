#!/bin/sh

CAMERAIP=192.168.10.105

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

create_encoder e_640x360_mjpeg_cbr.json
create_encoder e_640x360_mjpeg_cbr.json
create_encoder e_640x360_mjpeg_cbr.json
create_encoder e_640x360_mjpeg_cbr.json
create_encoder e_640x360_mjpeg_cbr.json
create_encoder e_640x360_mjpeg_cbr.json
create_encoder e_640x360_mjpeg_cbr.json

delete_all_encoders

create_encoder_id 10 e_640x360_mjpeg_cbr.json
create_encoder_id 9 e_640x360_mjpeg_cbr.json
create_encoder_id 8 e_640x360_mjpeg_cbr.json
create_encoder_id 22 e_640x360_mjpeg_cbr.json
create_encoder_id 0 e_640x360_mjpeg_cbr.json
create_encoder_id 1 e_640x360_mjpeg_cbr.json
create_encoder_id 2 e_640x360_mjpeg_cbr.json

delete_all_encoders

create_encoder e_3840x2160_mjpeg_cbr.json

