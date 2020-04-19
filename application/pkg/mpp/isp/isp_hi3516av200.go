//+build arm
//+build hi3516av200

package isp

/*
#include "../include/mpp.h"
#include "../errmpp/error.h"
#include "../../logger/logger.h"

#include <string.h>
#include <pthread.h>

static pthread_t hi3516av200_isp_thread_pid;

static void* hi3516av200_isp_thread(HI_VOID *param){
    int error_code = 0;
    GO_LOG_ISP(LOGGER_TRACE, "HI_MPI_ISP_Run")
    error_code = HI_MPI_ISP_Run(0);
    GO_LOG_ISP(LOGGER_ERROR, "HI_MPI_ISP_Run failed")
}

typedef struct hi3516av200_isp_init_in_struct {
    unsigned int width;
    unsigned int height;
    unsigned int fps;
    unsigned int bayer;
    unsigned int wdr;
} hi3516av200_isp_init_in;

static int hi3516av200_isp_init(error_in *err, hi3516av200_isp_init_in *in) { 
    unsigned int mpp_error_code = 0;
    int general_error_code = 0;

    ALG_LIB_S stLib;

    ALG_LIB_S stAeLib;

    stAeLib.s32Id = 0;
    strncpy(stAeLib.acLibName,  HI_AE_LIB_NAME,     sizeof(HI_AE_LIB_NAME));
    stLib.s32Id = 0;
    strcpy(stLib.acLibName, HI_AE_LIB_NAME);

    mpp_error_code = HI_MPI_AE_Register(0, &stLib);
    if (mpp_error_code != HI_SUCCESS) {
        GO_LOG_ISP(LOGGER_ERROR, "HI_MPI_AE_Register")        
        err->mpp = mpp_error_code;
        return ERR_MPP;
    }

    ALG_LIB_S stAwbLib;

    strncpy(stAwbLib.acLibName, HI_AWB_LIB_NAME,    sizeof(HI_AWB_LIB_NAME));
    stLib.s32Id = 0;
    strcpy(stLib.acLibName, HI_AWB_LIB_NAME);

    mpp_error_code = HI_MPI_AWB_Register(0, &stLib);
    if (mpp_error_code != HI_SUCCESS) {
        GO_LOG_ISP(LOGGER_ERROR, "HI_MPI_AWB_Register")
        err->mpp = mpp_error_code;
        return ERR_MPP;
    }

    ALG_LIB_S stAfLib;

    stAfLib.s32Id = 0;
    strncpy(stAfLib.acLibName,  HI_AF_LIB_NAME,     sizeof(HI_AF_LIB_NAME));
    stLib.s32Id = 0;
    strcpy(stLib.acLibName, HI_AF_LIB_NAME);

    mpp_error_code = HI_MPI_AF_Register(0, &stLib);
    if (mpp_error_code != HI_SUCCESS) {
        GO_LOG_ISP(LOGGER_ERROR, "HI_MPI_AF_Register")
        err->mpp = mpp_error_code;
        return ERR_MPP;
    }

    mpp_error_code = HI_MPI_ISP_MemInit(0);
    if (mpp_error_code != HI_SUCCESS) {
        GO_LOG_ISP(LOGGER_ERROR, "HI_MPI_ISP_MemInit")
        err->mpp = mpp_error_code;
        return ERR_MPP;
    }

    ISP_WDR_MODE_S stWdrMode;

    stWdrMode.enWDRMode  = in->wdr;
    //stWdrMode.enWDRMode  = WDR_MODE_2To1_LINE;
    //stWdrMode.enWDRMode  = WDR_MODE_NONE;


    mpp_error_code = HI_MPI_ISP_SetWDRMode(0, &stWdrMode);
    if (mpp_error_code != HI_SUCCESS) {
        GO_LOG_ISP(LOGGER_ERROR, "HI_MPI_ISP_SetWDRMode")
        err->mpp = mpp_error_code;
        return ERR_MPP;
    }

    ISP_PUB_ATTR_S stPubAttr;

    stPubAttr.enBayer               = in->bayer;
    stPubAttr.f32FrameRate          = in->fps;
    stPubAttr.stSnsSize.u32Width    = in->width; 
    stPubAttr.stSnsSize.u32Height   = in->height; 
    //Start position of the cropping window, image width, and image height
    stPubAttr.stWndRect.s32X        = 0;
    stPubAttr.stWndRect.s32Y        = 0;
    stPubAttr.stWndRect.u32Width    = in->width;
    stPubAttr.stWndRect.u32Height   = in->height;

    mpp_error_code = HI_MPI_ISP_SetPubAttr(0, &stPubAttr);
    if (mpp_error_code != HI_SUCCESS) {
        GO_LOG_ISP(LOGGER_ERROR, "HI_MPI_ISP_SetPubAttr")
        err->mpp = mpp_error_code;
        return ERR_MPP;
    }

    mpp_error_code = HI_MPI_ISP_Init(0);
    if (mpp_error_code != HI_SUCCESS) {
        GO_LOG_ISP(LOGGER_ERROR, "HI_MPI_ISP_Init")
        err->mpp = mpp_error_code;
        return ERR_MPP;
    }

    general_error_code = pthread_create(&hi3516av200_isp_thread_pid, 0, (void* (*)(void*))hi3516av200_isp_thread, NULL);
    if (general_error_code != 0) {
        GO_LOG_ISP(LOGGER_ERROR, "pthread_create")
        err->general = general_error_code;
        return ERR_GENERAL;
    }

    return ERR_NONE;
}
*/
import "C"

import (
    "errors"
    "application/pkg/mpp/errmpp"
    "application/pkg/logger"
    "application/pkg/mpp/cmos"
)

func initFamily() error {
    var inErr C.error_in

    cmos.Register()

    var in C.hi3516av200_isp_init_in
    in.width = C.uint(cmos.S.Width())
    in.height = C.uint(cmos.S.Height())
    in.fps = C.uint(cmos.S.Fps())

    switch cmos.S.Wdr() {
        case cmos.WDRNone:
            in.wdr = C.WDR_MODE_NONE
        case cmos.WDR2TO1:
            in.wdr = C.WDR_MODE_2To1_LINE
        default:
            logger.Log.Fatal().
                Msg("Unknown WDR mode")
    }

    switch cmos.S.Bayer() {
        case cmos.RGGB:
            in.bayer = C.BAYER_RGGB
        case cmos.GRBG:
            in.bayer = C.BAYER_GRBG
        case cmos.GBRG:
            in.bayer = C.BAYER_GBRG
        case cmos.BGGR:
            in.bayer = C.BAYER_BGGR
        default:
            logger.Log.Fatal().
                Msg("Unknown CMOS bayer")
    }

    logger.Log.Trace().
        Uint("width", uint(in.width)).
        Uint("height", uint(in.height)).
        Uint("fps", uint(in.fps)).
        Uint("bayer", uint(in.bayer)).
        Uint("wdr", uint(in.wdr)).
        Msg("ISP params")

    err := C.hi3516av200_isp_init(&inErr, &in)
    switch err {
        case C.ERR_MPP:
            return errmpp.New("funcname", uint(inErr.mpp))
        case C.ERR_GENERAL:
            return errors.New("ISP error TODO")
        default:
            return nil
    }
}
