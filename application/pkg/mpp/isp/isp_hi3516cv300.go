//+build arm
//+build hi3516cv300

package isp

/*
#include "../include/mpp_v3.h"

#include <stdio.h>
#include <string.h>
#include <pthread.h>

#define ERR_NONE    0
#define ERR_GENERAL 1
#define ERR_MPP     2

static pthread_t mpp3_isp_thread_pid;

HI_VOID* mpp3_isp_thread(HI_VOID *param){
    int error_code = 0;
    printf("C DEBUG: starting HI_MPI_ISP_Run...\n");
    error_code = HI_MPI_ISP_Run(0);
    printf("C DEBUG: HI_MPI_ISP_Run %d\n", error_code);
    //return error_code;
}

int mpp3_isp_init(int *error_code,
            unsigned int width,
            unsigned int height,
            unsigned int fps) {
    *error_code = 0;

    ISP_DEV IspDev = 0;
    ISP_PUB_ATTR_S stPubAttr;
    ALG_LIB_S stLib;

    ALG_LIB_S stAeLib;
    ALG_LIB_S stAwbLib;
    ALG_LIB_S stAfLib;

    stAeLib.s32Id = 0;
    stAwbLib.s32Id = 0;
    stAfLib.s32Id = 0;
    strncpy(stAeLib.acLibName, HI_AE_LIB_NAME, sizeof(HI_AE_LIB_NAME));
    strncpy(stAwbLib.acLibName, HI_AWB_LIB_NAME, sizeof(HI_AWB_LIB_NAME));
    strncpy(stAfLib.acLibName, HI_AF_LIB_NAME, sizeof(HI_AF_LIB_NAME)); 

    // 1. sensor register callback
    *error_code = sensor_register_callback();
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    // 2. register hisi ae lib
    stLib.s32Id = 0;
    strcpy(stLib.acLibName, HI_AE_LIB_NAME);
    *error_code = HI_MPI_AE_Register(IspDev, &stLib);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    // 3. register hisi awb lib
    stLib.s32Id = 0;
    strcpy(stLib.acLibName, HI_AWB_LIB_NAME);
    *error_code = HI_MPI_AWB_Register(IspDev, &stLib);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    // 4. register hisi af lib
    stLib.s32Id = 0;
    strcpy(stLib.acLibName, HI_AF_LIB_NAME);
    *error_code = HI_MPI_AF_Register(IspDev, &stLib);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    // 5. isp mem init 
    *error_code = HI_MPI_ISP_MemInit(IspDev);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    // 6. isp set WDR mode 
    ISP_WDR_MODE_S stWdrMode;
    stWdrMode.enWDRMode  = WDR_MODE_NONE;//enWDRMode;
    //stWdrMode.enWDRMode  = WDR_MODE_2To1_LINE ;//WDR_MODE_2To1_FRAME_FULL_RATE;//WDR_MODE_NONE;//enWDRMode;
    *error_code = HI_MPI_ISP_SetWDRMode(0, &stWdrMode);    
    if (*error_code != HI_SUCCESS) return ERR_MPP;


//./hi_comm_video.h:490:    BAYER_RGGB    = 0,
//./hi_comm_video.h:491:    BAYER_GRBG    = 1,
//./hi_comm_video.h:492:    BAYER_GBRG    = 2,
//./hi_comm_video.h:493:    BAYER_BGGR    = 3,

            stPubAttr.enBayer               = BAYER_RGGB; //bad
            //stPubAttr.enBayer               = BAYER_GRBG;   //better
            //stPubAttr.enBayer               = BAYER_GBRG; //bad
            //stPubAttr.enBayer               = BAYER_BGGR; //bad
            stPubAttr.stWndRect.s32X        = 0;//30;
            stPubAttr.stWndRect.s32Y        = 0;
            stPubAttr.stWndRect.u32Width    = width;
            stPubAttr.stWndRect.u32Height   = height;
            stPubAttr.f32FrameRate          = fps;

    *error_code = HI_MPI_ISP_SetPubAttr(IspDev, &stPubAttr);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    // 8. isp init 
    *error_code = HI_MPI_ISP_Init(IspDev);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    if (pthread_create(&mpp3_isp_thread_pid, 0, (void* (*)(void*))mpp3_isp_thread, NULL) != 0) {
        return ERR_GENERAL;
    }

	return ERR_NONE;
}



*/
import "C"

import (
         "application/pkg/mpp/error"
        
        "application/pkg/logger"

        "application/pkg/mpp/cmos"
)

func Init() {
    var errorCode C.int

     //   switch err := C.mpp3_isp_init(&errorCode); err {


    switch err := C.mpp3_isp_init(  &errorCode, 
                    C.uint(cmos.Width()), 
                    C.uint(cmos.Height()), 
                    C.uint(cmos.Fps()) ); err {

    case C.ERR_NONE:
        logger.Log.Debug().
                Msg("C.mpp3_isp_init() ok")
    case C.ERR_MPP:
        logger.Log.Fatal().
                Int("error", int(errorCode)).
                Str("error_desc", error.Resolve(int64(errorCode))).
                Msg("C.mpp3_isp_init() mpp error ")
    default:
            logger.Log.Fatal().
                Int("error", int(err)).
                Msg("Unexpected return of C.mpp3_isp_init()")
        }

}
