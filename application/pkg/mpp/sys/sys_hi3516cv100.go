//+build arm
//+build hi3516cv100

package sys

/*
#include "../include/mpp_v1.h"

#include <string.h>

#define ERR_NONE                0
#define ERR_HI_MPI_SYS_Exit     2
#define ERR_HI_MPI_VB_Exit      3
#define ERR_HI_MPI_VB_SetConf   4
#define ERR_HI_MPI_VB_Init      5
#define ERR_HI_MPI_SYS_SetConf  6
#define ERR_HI_MPI_SYS_Init     7

int mpp1_sys_init(unsigned int *error_code) {
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

	switch err := C.mpp1_sys_init(&errorCode); err {
	case C.ERR_NONE:
		log.Println("C.mpp1_sys_init ok")
	case C.ERR_HI_MPI_SYS_Exit:
		log.Fatal("C.mpp1_sys_init() HI_MPI_SYS_Exit() error ", error.Resolve(int64(errorCode)))
	case C.ERR_HI_MPI_VB_Exit:
		log.Fatal("C.mpp1_sys_init() HI_MPI_VB_Exit() error ", error.Resolve(int64(errorCode)))
	case C.ERR_HI_MPI_VB_SetConf:
		log.Fatal("C.mpp1_sys_init() HI_MPI_VB_SetConf() error ", error.Resolve(int64(errorCode)))
	case C.ERR_HI_MPI_VB_Init:
		log.Fatal("C.mpp1_sys_init() HI_MPI_VB_Init() error ", error.Resolve(int64(errorCode)))
	case C.ERR_HI_MPI_SYS_SetConf:
		log.Fatal("C.mpp1_sys_init() HI_MPI_SYS_SetConf() error ", error.Resolve(int64(errorCode)))
	case C.ERR_HI_MPI_SYS_Init:
		log.Fatal("C.mpp1_sys_init() HI_MPI_SYS_Init() error ", error.Resolve(int64(errorCode)))
	default:
		log.Fatal("Unexpected return ", err, " of C.mpp1_sys_init()")
	}
}
