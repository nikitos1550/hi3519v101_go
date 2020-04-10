//+build arm
//+build hi3516av200

package isp

/*
#include "../include/mpp_v3.h"

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

int mpp3_isp_init(int *error_code) {
    *error_code = 0;

    ISP_PUB_ATTR_S stPubAttr;
    ALG_LIB_S stLib;

    *error_code = HI_MPI_ISP_Exit(0);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    ALG_LIB_S stAeLib;
    ALG_LIB_S stAwbLib;
    ALG_LIB_S stAfLib;

    stAeLib.s32Id = 0;
    stAwbLib.s32Id = 0;
    stAfLib.s32Id = 0;
    strncpy(stAeLib.acLibName,  HI_AE_LIB_NAME,     sizeof(HI_AE_LIB_NAME));
    strncpy(stAwbLib.acLibName, HI_AWB_LIB_NAME,    sizeof(HI_AWB_LIB_NAME));
    strncpy(stAfLib.acLibName,  HI_AF_LIB_NAME,     sizeof(HI_AF_LIB_NAME));


    //TODO
    #ifdef HI3516AV200
    ISP_SNS_OBJ_S *cmos = &stSnsImx274Obj;
    if (cmos->pfnRegisterCallback != HI_NULL) {
        *error_code = cmos->pfnRegisterCallback(0, &stAeLib, &stAwbLib);
        if (*error_code != HI_SUCCESS) return ERR_GENERAL;
    } else {
        return ERR_GENERAL;
    }
    #endif
    #ifdef HI3516CV300
    //*error_code = sensor_register_callback();
    //if (*error_code != HI_SUCCESS) return ERR_MPP;
    #endif

    stLib.s32Id = 0;
    strcpy(stLib.acLibName, HI_AE_LIB_NAME);

    *error_code = HI_MPI_AE_Register(0, &stLib);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    stLib.s32Id = 0;
    strcpy(stLib.acLibName, HI_AWB_LIB_NAME);

    *error_code = HI_MPI_AWB_Register(0, &stLib);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    stLib.s32Id = 0;
    strcpy(stLib.acLibName, HI_AF_LIB_NAME);

    *error_code = HI_MPI_AF_Register(0, &stLib);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    *error_code = HI_MPI_ISP_MemInit(0);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    ISP_WDR_MODE_S stWdrMode;
    stWdrMode.enWDRMode  = WDR_MODE_NONE;

    *error_code = HI_MPI_ISP_SetWDRMode(0, &stWdrMode);
    if (*error_code != HI_SUCCESS) return ERR_MPP;
    //TODO WDR modes support

    stPubAttr.enBayer               = BAYER_RGGB;
    stPubAttr.f32FrameRate          = 30;
    stPubAttr.stWndRect.s32X        = 0;
    stPubAttr.stWndRect.s32Y        = 0;
    stPubAttr.stWndRect.u32Width    = 3840;     //TODO What is WND rect?
    stPubAttr.stWndRect.u32Height   = 2160;    //TODO
    #ifdef HI3516AV200
    stPubAttr.stSnsSize.u32Width    = 3840;
    stPubAttr.stSnsSize.u32Height   = 2160;
    #endif

    *error_code = HI_MPI_ISP_SetPubAttr(0, &stPubAttr);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    *error_code = HI_MPI_ISP_Init(0);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    if (pthread_create(&mpp3_isp_thread_pid, 0, (void* (*)(void*))mpp3_isp_thread, NULL) != 0) {
        return ERR_GENERAL;
    }

    return ERR_NONE;
}
*/
import "C"

import (
    //"log"
    "application/pkg/mpp/error"
    "application/pkg/logger"
)

func Init() {
    var errorCode C.int

	switch err := C.mpp3_isp_init(&errorCode); err {
    case C.ERR_NONE:
        //log.Println("C.mpp3_isp_init() ok")
	logger.Log.Debug().
		Msg("C.mpp3_isp_init() ok")
    case C.ERR_MPP:
        //log.Fatal("C.mpp3_isp_init() mpp error ", error.Resolve(int64(errorCode)))
	logger.Log.Fatal().
                Int("error", int(errorCode)).
                Str("error_desc", error.Resolve(int64(errorCode))).
		Msg("C.mpp3_isp_init() mpp error ")
    default:
	    //log.Fatal("Unexpected return ", err , " of C.mpp3_isp_init()")
	    logger.Log.Fatal().
	    	Int("error", int(err)).
		Msg("Unexpected return of C.mpp3_isp_init()")
	}
}
