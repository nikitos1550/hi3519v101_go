//+build 386 amd64
//+build host
package venc

/*
#include "loop.h"
*/
//import "C"

/*
import (
	"application/pkg/logger"
	"application/pkg/mpp/error"
)

func DeleteEncoder(encoder ActiveEncoder) {
	var errorCode C.uint
	var err C.int

	delVenc(encoder.VencId) //first we remove fd from loop

	err = C.mpp3_venc_delete_encoder(&errorCode, C.int(encoder.VencId))
	switch err {
	case C.ERR_NONE:
		//log.Println("Encoder deleted ", encoder.VencId)
		logger.Log.Debug().
			Int("vencId", encoder.VencId).
			Msg("Encoder deleted")
	case C.ERR_MPP:
		//log.Fatal("Failed to delete encoder ", encoder.VencId, " error ", error.Resolve(int64(errorCode)))
		logger.Log.Fatal().
			Int("vencId", encoder.VencId).
			Int("error", int(errorCode)).
			Str("error_code", error.Resolve(int64(errorCode))).
			Msg("Failed to delete encoder")
	default:
		//log.Fatal("Failed to delete encoder ", encoder.VencId, "Unexpected return ", err)
		logger.Log.Fatal().
			Int("error", int(err)).
			Msg("Failed to delete encoder, unexpected return")

	}

}

func CreateEncoder(encoder ActiveEncoder) {
	var errorCode C.uint
	var err C.int
	switch encoder.Format {
	case "h264":
		err = C.mpp3_venc_sample_h264(&errorCode, C.int(encoder.Width), C.int(encoder.Height), C.int(encoder.Bitrate), C.int(encoder.VencId))
	case "h265":
		err = C.mpp3_venc_sample_h265(&errorCode, C.int(encoder.Width), C.int(encoder.Height), C.int(encoder.Bitrate), C.int(encoder.VencId))
	case "mjpeg":
		err = C.mpp3_venc_sample_mjpeg(&errorCode, C.int(encoder.Width), C.int(encoder.Height), C.int(encoder.Bitrate), C.int(encoder.VencId))
	default:
		//log.Println("Unknown encoder format ", encoder.Format)
		logger.Log.Warn().
			Str("codec", encoder.Format).
			Msg("Unknown encoder format")
	}

	switch err {
	case C.ERR_NONE:
		//log.Println("Encoder created ", encoder.Format)
		logger.Log.Debug(). //TODO encoderId
			Str("codec", encoder.Format).
			Msg("Encoder created")

	case C.ERR_MPP:
		//log.Fatal("Failed to create encoder ", encoder.Format, " error ", error.Resolve(int64(errorCode)))
		logger.Log.Fatal().
			Str("codec", encoder.Format).
			Int("error", int(errorCode)).
			Str("error-dec", error.Resolve(int64(errorCode))).
			Msg("Failed to create encoder")
	default:
		//log.Fatal("Failed to create encoder ", encoder.Format, "Unexpected return ", err)
		logger.Log.Fatal().
			Str("codec", encoder.Format).
			Int("error", int(err)).
			Msg("Failed to create encoder, unexpected return")
	}

	addVenc(encoder.VencId)
}
*/
