//+build arm
//+build hi3516cv500

package venc

/*
#include "../include/mpp_v4.h"

#include <string.h>

#define ERR_NONE                0
#define ERR_MPP                 2

int mpp4_venc_sample_mjpeg(unsigned int *error_code) {
    *error_code = 0;

    return ERR_NONE;
}

int mpp4_venc_sample_h264(unsigned int *error_code) {
    *error_code = 0;

    return ERR_NONE;
}
*/
import "C"

import (
	"application/pkg/mpp/error"
	"log"
)

var (
	SampleMjpegFrames *frames
	SampleH264Frames  *frames
	SampleH264Notify  chan int
	SampleH264Start   chan int
)

func SampleMjpeg() {
	var errorCode C.uint

	switch err := C.mpp4_venc_sample_mjpeg(&errorCode); err {
	case C.ERR_NONE:
		log.Println("C.mpp4_venc_sample_mjpeg() ok")
	case C.ERR_MPP:
		log.Fatal("C.mpp4_venc_sample_mjpeg() MPP error ", error.Resolve(int64(errorCode)))
	default:
		log.Fatal("Unexpected return ", err, " of C.mpp4_venc_sample_mjpeg()")
	}

	//TODO //create corresponding encoder object
	SampleMjpegFrames = CreateFrames(3)
	addVenc(1)
}

func SampleH264() {
	var errorCode C.uint

	switch err := C.mpp4_venc_sample_h264(&errorCode); err {
	case C.ERR_NONE:
		log.Println("C.mpp4_venc_sample_h264() ok")
	case C.ERR_MPP:
		log.Fatal("C.mpp4_venc_sample_h264() MPP error ", error.Resolve(int64(errorCode)))
	default:
		log.Fatal("Unexpected return ", err, " of C.mpp4_venc_sample_h264()")
	}

	//TODO //create corresponding encoder object
	SampleH264Frames = CreateFrames(10)
	SampleH264Notify = make(chan int, 10)
	SampleH264Start = make(chan int, 1)
	addVenc(0) //add venc to get loop
}

func createEncoder(encoder Encoder) {}

func deleteEncoder(encoder Encoder) {}

