//+build arm
//+build hi3516cv100

package isp

/*
#include "../include/mpp_v1.h"

#include <stdio.h>
#include <string.h>
#include <pthread.h>

#define ERR_NONE    0
#define ERR_GENERAL 1
#define ERR_MPP     2

static pthread_t mpp1_isp_thread_pid;

HI_VOID* mpp1_isp_thread(HI_VOID *param){
    int error_code = 0;
    printf("C DEBUG: starting HI_MPI_ISP_Run...\n");
    error_code = HI_MPI_ISP_Run();
    printf("C DEBUG: HI_MPI_ISP_Run %d\n", error_code);
    //return error_code;
}

int mpp1_isp_init(int *error_code) {
    *error_code = 0;

       // 2. sensor register callback
    *error_code = sensor_register_callback();
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    ISP_IMAGE_ATTR_S stImageAttr;
    ISP_INPUT_TIMING_S stInputTiming;
    
    ALG_LIB_S stLib;

    // 1. register ae lib
    stLib.s32Id = 0;
    strcpy(stLib.acLibName, HI_AE_LIB_NAME);
    *error_code = HI_MPI_AE_Register(&stLib);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    // 2. register awb lib
    stLib.s32Id = 0;
    strcpy(stLib.acLibName, HI_AWB_LIB_NAME);
    *error_code = HI_MPI_AWB_Register(&stLib);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    // 3. register af lib
    stLib.s32Id = 0;
    strcpy(stLib.acLibName, HI_AF_LIB_NAME);
    *error_code = HI_MPI_AF_Register(&stLib);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    // 4. isp init
    *error_code = HI_MPI_ISP_Init();
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    // 5. isp set image attributes
    // note : different sensor, different ISP_IMAGE_ATTR_S define.
    //          if the sensor you used is different, you can change
    //          ISP_IMAGE_ATTR_S definition


	//case SONY_IMX122_DC_1080P_30FPS:
        //case SONY_IMX236_DC_1080P_30FPS:
            stImageAttr.enBayer      = BAYER_RGGB;
            stImageAttr.u16FrameRate = 30;
            stImageAttr.u16Width     = 1920;
            stImageAttr.u16Height    = 1080;

	    *error_code = HI_MPI_ISP_SetImageAttr(&stImageAttr);
	    if (*error_code != HI_SUCCESS) return ERR_MPP;

    // 6. isp set timing
    //    case SONY_IMX122_DC_1080P_30FPS:
            stInputTiming.enWndMode = ISP_WIND_ALL;
            stInputTiming.u16HorWndStart = 200;
            stInputTiming.u16HorWndLength = 1920;
            stInputTiming.u16VerWndStart = 18;
            stInputTiming.u16VerWndLength = 1080;

     *error_code = HI_MPI_ISP_SetInputTiming(&stInputTiming);
     if (*error_code != HI_SUCCESS) return ERR_MPP;

    if (pthread_create(&mpp1_isp_thread_pid, 0, (void* (*)(void*))mpp1_isp_thread, NULL) != 0) {
        return ERR_GENERAL;
    }



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

        switch err := C.mpp1_isp_init(&errorCode); err {
    case C.ERR_NONE:
        logger.Log.Debug().
                Msg("C.mpp1_isp_init() ok")
    case C.ERR_MPP:
        logger.Log.Fatal().
                Int("error", int(errorCode)).
                Str("error_desc", error.Resolve(int64(errorCode))).
                Msg("C.mpp1_isp_init() mpp error ")
    default:
            logger.Log.Fatal().
                Int("error", int(err)).
                Msg("Unexpected return of C.mpp1_isp_init()")
        }

}
