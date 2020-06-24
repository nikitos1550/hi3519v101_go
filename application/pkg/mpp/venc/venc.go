package venc

//#include "venc.h"
import "C"

import (
    "application/pkg/logger"
   // "application/pkg/mpp/errmpp"
)


func Init() {
    loopInit()
	readEncoders()

}

//func SubsribeEncoder(encoderId int, ch chan []byte) {
func SubsribeEncoder(encoderId int, ch chan ChannelEncoder) {
    encoder, encoderExists := ActiveEncoders[encoderId]
    if !encoderExists {
		logger.Log.Error().
			Int("encoderId", encoderId).
			Msg("Failed to find encoder")
        return
    }
		
    _, exists := encoder.Channels[ch]
    if (exists) {
		logger.Log.Error().
			Int("encoderId", encoderId).
			Msg("Already subscribed")
        return
    }

    encoder.Channels[ch] = true
    ActiveEncoders[encoderId] = encoder
}

//func RemoveSubscription(encoderId int, ch chan []byte) {
func RemoveSubscription(encoderId int, ch chan ChannelEncoder) {
    encoder, encoderExists := ActiveEncoders[encoderId]
    if !encoderExists {
		logger.Log.Error().
			Int("encoderId", encoderId).
			Msg("Failed to find encoder")
        return
    }
		
    _, exists := encoder.Channels[ch]
    if (!exists) {
		logger.Log.Error().
			Int("encoderId", encoderId).
			Msg("Not subscribed")
        return
    }


    delete(encoder.Channels, ch)
    ActiveEncoders[encoderId] = encoder
}

/*
func CreateVencEncoder(encoder ActiveEncoder) {
	var inErr C.error_in
	var err C.int

	switch encoder.Format {
	case "h264":
        var in C.hi3516av200_venc_create_h264_in

        in.venc_id = C.uint(encoder.VencId)
        in.width = C.uint(encoder.Width)
        in.height = C.uint(encoder.Height)
        in.bitrate = C.uint(encoder.Bitrate)

		err = C.hi3516av200_venc_create_h264(&inErr, &in)
	case "h265":
        var in C.hi3516av200_venc_create_h265_in
		
        in.venc_id = C.uint(encoder.VencId)
        in.width = C.uint(encoder.Width)
        in.height = C.uint(encoder.Height)
        in.bitrate = C.uint(encoder.Bitrate)

		err = C.hi3516av200_venc_create_h265(&inErr, &in)
	case "mjpeg":
        var in C.hi3516av200_venc_create_mjpeg_in

        in.venc_id = C.uint(encoder.VencId)
        in.width = C.uint(encoder.Width)
        in.height = C.uint(encoder.Height)
        in.bitrate = C.uint(encoder.Bitrate)

		err = C.mpp3_venc_sample_mjpeg(&inErr, &in)
	default:
		logger.Log.Error().
			Str("codec", encoder.Format).
			Msg("VENC unknown codec")
        return
	}

    if err != C.ERR_NONE {
        logger.Log.Fatal(). //log temporary, should generate and return error
            Str("error", errmpp.New("funcname", uint(inErr.mpp)).Error()).
            Msg("VENC")
    }

	addVenc(encoder.VencId)
}
*/
/*
func DeleteVencEncoder(encoder ActiveEncoder) {
	var inErr C.error_in
	var err C.int

	delVenc(encoder.VencId) //first we remove fd from loop

	err = C.hi3516av200_venc_delete_encoder(&inErr, C.uint(encoder.VencId))

	if err != C.ERR_NONE {
    	logger.Log.Fatal(). //log temporary, should generate and return error
        	Str("error", errmpp.New("funcname", uint(inErr.mpp)).Error()).
            Msg("VENC")
    }

}
*/
