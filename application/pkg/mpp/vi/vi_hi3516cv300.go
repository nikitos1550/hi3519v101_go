//+build arm
//+build hi3516cv300

package vi

/*

#include "../include/mpp_v3.h"
#include <string.h>

#define ERR_NONE                    0
#define ERR_HI_MPI_VI_SetDevAttr        2
#define ERR_HI_MPI_VI_EnableDev     3
#define ERR_HI_MPI_VI_SetChnAttr    4
#define ERR_HI_MPI_VI_EnableChn     5
#define ERR_TMP   100

int mpp3111_vi_init(unsigned int *error_code,  void *videv, unsigned int width, unsigned int height, unsigned int fps) {
    *error_code = 0;

    printf("VI width=%d height=%d fps=%d\n", width, height, fps);

    //HI_S32 s32IspDev = 0;
    ISP_WDR_MODE_S stWdrMode;
    VI_DEV_ATTR_S  stViDevAttr;
    
    memset(&stViDevAttr,0,sizeof(stViDevAttr));
    memcpy(&stViDevAttr, videv, sizeof(stViDevAttr));

	    //stViDevAttr.stDevRect.s32X = 0;//200;
        //    stViDevAttr.stDevRect.s32Y = 0;//20;
            stViDevAttr.stDevRect.u32Width  = width;
            stViDevAttr.stDevRect.u32Height = height;


    *error_code = HI_MPI_VI_SetDevAttr(0, &stViDevAttr);
    if (*error_code != HI_SUCCESS) return ERR_HI_MPI_VI_SetDevAttr;

      VI_WDR_ATTR_S stWdrAttr;

        stWdrAttr.enWDRMode = WDR_MODE_NONE;//WDR_MODE_2To1_LINE;
        stWdrAttr.bCompress = HI_FALSE;

    *error_code = HI_MPI_VI_SetWDRAttr(0, &stWdrAttr);
    if (*error_code != HI_SUCCESS) return ERR_TMP;
    
    *error_code = HI_MPI_VI_EnableDev(0);
    if (*error_code != HI_SUCCESS) return ERR_HI_MPI_VI_EnableDev;

       RECT_S stCapRect;
    SIZE_S stTargetSize;

     stCapRect.s32X = 0;
        stCapRect.s32Y = 0;
                stCapRect.u32Width  = width;
                stCapRect.u32Height = height;
       stTargetSize.u32Width = stCapRect.u32Width;
        stTargetSize.u32Height = stCapRect.u32Height;


    VI_CHN_ATTR_S stChnAttr;

	memcpy(&stChnAttr.stCapRect, &stCapRect, sizeof(RECT_S));

    stChnAttr.enCapSel = VI_CAPSEL_BOTH;
    stChnAttr.stDestSize.u32Width = stTargetSize.u32Width ;
    stChnAttr.stDestSize.u32Height =  stTargetSize.u32Height ;
    stChnAttr.enPixFormat = PIXEL_FORMAT_YUV_SEMIPLANAR_420;   // sp420 or sp422

    stChnAttr.bMirror = HI_FALSE;
    stChnAttr.bFlip = HI_FALSE;

    stChnAttr.s32SrcFrameRate = fps;
    stChnAttr.s32DstFrameRate = fps;
    stChnAttr.enCompressMode = COMPRESS_MODE_NONE;

    *error_code = HI_MPI_VI_SetChnAttr(0, &stChnAttr);
    if (*error_code != HI_SUCCESS) return ERR_HI_MPI_VI_SetChnAttr;

    *error_code = HI_MPI_VI_EnableChn(0);
    if (*error_code != HI_SUCCESS) return ERR_HI_MPI_VI_EnableChn;

    return ERR_NONE;
}


*/
import "C"

import (
        "application/pkg/logger"

    "application/pkg/mpp/error"

    "application/pkg/mpp/cmos"
)

func Init() {
    var errorCode C.uint

    //switch err := C.mpp3_vi_init(&errorCode); err {
   switch err := C.mpp3111_vi_init(&errorCode, cmos.ViDev(), C.uint(cmos.Width()), C.uint(cmos.Height()), C.uint(cmos.Fps())); err {

    case C.ERR_NONE:
        logger.Log.Debug().
                Msg("C.mpp3_vi_init() ok")
    case C.ERR_HI_MPI_VI_SetDevAttr:
        logger.Log.Fatal().
                Str("func", "ERR_HI_MPI_VI_SetDevAttr()").
                Int("error", int(errorCode)).
                Str("error_desc", error.Resolve(int64(errorCode))).
                Msg("C.mpp3_vi_init() error")
    case C.ERR_HI_MPI_VI_EnableDev:
        logger.Log.Fatal().
                Str("func", "ERR_HI_MPI_VI_EnableDev()").
                Int("error", int(errorCode)).
                Str("error_desc", error.Resolve(int64(errorCode))).
                Msg("C.mpp3_vi_init() error")

    case C.ERR_HI_MPI_VI_SetChnAttr:
        logger.Log.Fatal().
                Str("func", "ERR_HI_MPI_VI_SetChnAttr()"). 
                Int("error", int(errorCode)).
                Str("error_desc", error.Resolve(int64(errorCode))).
                Msg("C.mpp3_vi_init() error")
    case C.ERR_HI_MPI_VI_EnableChn:
        logger.Log.Fatal().
                Str("func", "ERR_HI_MPI_VI_EnableChn()"). 
                Int("error", int(errorCode)).
                Str("error_desc", error.Resolve(int64(errorCode))).
                Msg("C.mpp3_vi_init() error")
    default:
        logger.Log.Fatal().
                Int("error", int(err)).
                Msg("C.mpp3_vi_init() Unexpected return")

    }
}

