//+build arm
//+build hi3516av100

package venc

/*
#include "../include/mpp_v2.h"

#include <string.h>

#define ERR_NONE                0
#define ERR_MPP                 2

int mpp3_venc_sample_mjpeg(unsigned int *error_code) {
    *error_code = 0;

    return ERR_NONE;
}

int mpp3_venc_sample_h264(unsigned int *error_code) {
    *error_code = 0;

    return ERR_NONE;
}
*/
import "C"

import (
	"application/pkg/mpp/mpperror"
	"log"
)


func createEncoder(encoder Encoder) {}

func deleteEncoder(encoder Encoder) {}
