//+build arm
//+build hi3516av100

package isp

/*
#include "../include/mpp_v2.h"

#include <string.h>
#include <pthread.h>

#define ERR_NONE    0
#define ERR_GENERAL 1
#define ERR_MPP     2

static pthread_t mpp2_isp_thread_pid;

HI_VOID* mpp2_isp_thread(HI_VOID *param){
    int error_code = 0;
    printf("C DEBUG: starting HI_MPI_ISP_Run...\n");
    error_code = HI_MPI_ISP_Run(0);
    printf("C DEBUG: HI_MPI_ISP_Run %d\n", error_code);
    //return error_code;
}

int mpp2_isp_init(int *error_code,
            unsigned int width,
            unsigned int height,
            unsigned int fps) {
    *error_code = 0;

    //*error_code = HI_MPI_ISP_Exit(0);
    //if (*error_code != HI_SUCCESS) return ERR_MPP;

    ISP_DEV IspDev = 0;
    ISP_PUB_ATTR_S stPubAttr;  
    ALG_LIB_S stLib;
                
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
    *error_code = HI_MPI_ISP_SetWDRMode(0, &stWdrMode);    
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    // 7. isp set pub attributes
    	// note : different sensor, different ISP_PUB_ATTR_S define.
      //        if the sensor you used is different, you can change
      //        ISP_PUB_ATTR_S definition 
      // case SONY_IMX178_LVDS_5M_30FPS:
            stPubAttr.enBayer               = BAYER_BGGR; //BAYER_GBRG;
            stPubAttr.f32FrameRate          = fps;
            stPubAttr.stWndRect.s32X        = 0;
            stPubAttr.stWndRect.s32Y        = 0;
            stPubAttr.stWndRect.u32Width    = width;
            stPubAttr.stWndRect.u32Height   = height;

    *error_code = HI_MPI_ISP_SetPubAttr(IspDev, &stPubAttr);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    //* 8. isp init 
    *error_code = HI_MPI_ISP_Init(IspDev);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    if (pthread_create(&mpp2_isp_thread_pid, 0, (void* (*)(void*))mpp2_isp_thread, NULL) != 0) {
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

        //switch err := C.mpp2_isp_init(&errorCode); err {

    switch err := C.mpp2_isp_init(  &errorCode, 
                    C.uint(cmos.Width()), 
                    C.uint(cmos.Height()), 
                    C.uint(cmos.Fps()) ); err {

    case C.ERR_NONE:
        logger.Log.Debug().
                Msg("C.mpp2_isp_init() ok")
    case C.ERR_MPP:
        logger.Log.Fatal().
                Int("error", int(errorCode)).
                Str("error_desc", error.Resolve(int64(errorCode))).
                Msg("C.mpp2_isp_init() mpp error ")
    default:
            logger.Log.Fatal().
                Int("error", int(err)).
                Msg("Unexpected return of C.mpp2_isp_init()")
        }

}
