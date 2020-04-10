//+build arm
//+build hi3516ev200

package isp

/*
#include "../include/mpp_v4.h"

#include <stdio.h>
#include <string.h>
#include <pthread.h>

#define ERR_NONE    0
#define ERR_GENERAL 1
#define ERR_MPP     2

static pthread_t mpp1_isp_thread_pid;

HI_VOID* mpp4_isp_thread(HI_VOID *param){
    int error_code = 0;
    printf("C DEBUG: starting HI_MPI_ISP_Run...\n");
    error_code = HI_MPI_ISP_Run(0);
    printf("C DEBUG: HI_MPI_ISP_Run %d\n", error_code);
    //return error_code;
}

int mpp4_isp_init(int *error_code) {
    *error_code = 0;

	return ERR_NONE;
}



*/
import "C"

import (
         "application/pkg/mpp/error"
        
        "application/pkg/logger"
)

func Init() {
    var errorCode C.int

        switch err := C.mpp4_isp_init(&errorCode); err {
    case C.ERR_NONE:
        logger.Log.Debug().
                Msg("C.mpp4_isp_init() ok")
    case C.ERR_MPP:
        logger.Log.Fatal().
                Int("error", int(errorCode)).
                Str("error_desc", error.Resolve(int64(errorCode))).
                Msg("C.mpp4_isp_init() mpp error ")
    default:
            logger.Log.Fatal().
                Int("error", int(err)).
                Msg("Unexpected return of C.mpp4_isp_init()")
        }

}
