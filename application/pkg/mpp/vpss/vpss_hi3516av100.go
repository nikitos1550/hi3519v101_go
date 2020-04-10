//+build arm
//+build hi3516av100

package vpss

/*
#include "../include/mpp_v2.h"

#include <string.h>

#define ERR_NONE                    0
#define ERR_MPP                     2
#define ERR_HI_MPI_VPSS_CreateGrp   3
#define ERR_HI_MPI_VPSS_StartGrp    4
#define ERR_HI_MPI_SYS_Bind         5

int mpp2_vpss_init(unsigned int *error_code) {
    *error_code = 0;
    return ERR_NONE;
}

int mpp2_vpss_sample_channel0(unsigned int *error_code) {
    *error_code = 0;
    return ERR_NONE;
}
*/
import "C"

import (
	"application/pkg/mpp/error"
	"log"
)

func Init() {
	var errorCode C.uint

	switch err := C.mpp2_vpss_init(&errorCode); err {
	case C.ERR_NONE:
		log.Println("C.mpp2_vpss_init() ok")
	case C.ERR_HI_MPI_VPSS_CreateGrp:
		log.Fatal("C.mpp2_vpss_init() HI_MPI_VPSS_CreateGrp() error ", error.Resolve(int64(errorCode)))
	case C.ERR_HI_MPI_VPSS_StartGrp:
		log.Fatal("C.mpp2_vpss_init() HI_MPI_VPSS_StartGrp() error ", error.Resolve(int64(errorCode)))
	case C.ERR_HI_MPI_SYS_Bind:
		log.Fatal("C.mpp2_vpss_init() HI_MPI_SYS_Bind() error ", error.Resolve(int64(errorCode)))
	default:
		log.Fatal("Unexpected return ", err, " of C.mpp2_vpss_init()")
	}
}

func SampleChannel0() {
	var errorCode C.uint

	switch err := C.mpp2_vpss_sample_channel0(&errorCode); err {
	case C.ERR_NONE:
		log.Println("C.mpp2_vpss_sample_channel0() ok")
	case C.ERR_MPP:
		log.Fatal("C.mpp2_vpss_sample_channel0() MPP error ", error.Resolve(int64(errorCode)))
	default:
		log.Fatal("Unexpected return ", err, " of C.mpp2_vpss_sample_channel0()")
	}

}

func CreateChannel(channel Channel) {}

func DestroyChannel(channel Channel) {}
