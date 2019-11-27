//+build hi3516av200 hi3516cv300

package vi

/*
#include "../include/hi3516av200_mpp.h"
#include <string.h>

#define ERR_NONE                    0
#define ERR_HI_MPI_VI_SetDevAttr        2
#define ERR_HI_MPI_VI_EnableDev     3
#define ERR_HI_MPI_VI_SetChnAttr    4
#define ERR_HI_MPI_VI_EnableChn     5

int mpp3_vi_init(int *error_code) {
    *error_code = 0;

    VI_DEV_ATTR_S  stViDevAttr;

    memset(&stViDevAttr, 0, sizeof(stViDevAttr));
    //TODO videv description struct
    //memcpy(&stViDevAttr, c->videv, sizeof(stViDevAttr));

    //stViDevAttr.stDevRect.s32X                              = 0;
    //stViDevAttr.stDevRect.s32Y                              = 0;
    stViDevAttr.stDevRect.u32Width                          = 3840;
    stViDevAttr.stDevRect.u32Height                         = 2160;
    stViDevAttr.stBasAttr.stSacleAttr.stBasSize.u32Width    = 3840;
    stViDevAttr.stBasAttr.stSacleAttr.stBasSize.u32Height   = 2160;
    stViDevAttr.stBasAttr.stSacleAttr.bCompress             = HI_FALSE;

    *error_code = HI_MPI_VI_SetDevAttr(0, &stViDevAttr);
    if (*error_code != HI_SUCCESS) return ERR_HI_MPI_VI_SetDevAttr;

    *error_code = HI_MPI_VI_EnableDev(0);
    if (*error_code != HI_SUCCESS) return ERR_HI_MPI_VI_EnableDev;

    RECT_S stCapRect;
    SIZE_S stTargetSize;

    stCapRect.s32X          = 0;
    stCapRect.s32Y          = 0;
    stCapRect.u32Width      = 3840;
    stCapRect.u32Height     = 2160;
    stTargetSize.u32Width   = stCapRect.u32Width;
    stTargetSize.u32Height  = stCapRect.u32Height;

    VI_CHN_ATTR_S stChnAttr;

    memcpy(&stChnAttr.stCapRect, &stCapRect, sizeof(RECT_S));

    stChnAttr.enCapSel              = VI_CAPSEL_BOTH;
    stChnAttr.stDestSize.u32Width   = stTargetSize.u32Width ;
    stChnAttr.stDestSize.u32Height  = stTargetSize.u32Height ;
    stChnAttr.enPixFormat           = PIXEL_FORMAT_YUV_SEMIPLANAR_420;   // sp420 or sp422

    stChnAttr.bMirror               = HI_FALSE;
    stChnAttr.bFlip                 = HI_FALSE;

    stChnAttr.s32SrcFrameRate       = 30;
    stChnAttr.s32DstFrameRate       = 30;
    stChnAttr.enCompressMode        = COMPRESS_MODE_NONE;

    *error_code = HI_MPI_VI_SetChnAttr(0, &stChnAttr);
    if (*error_code != HI_SUCCESS) return ERR_HI_MPI_VI_SetChnAttr;

    *error_code = HI_MPI_VI_EnableChn(0);
    if (*error_code != HI_SUCCESS) return ERR_HI_MPI_VI_EnableChn;

    return ERR_NONE;
}
*/
import "C"

import (
    "log"
)

func Init() {
    var errorCode C.int

    switch err := C.mpp3_vi_init(&errorCode); err {
    case C.ERR_NONE:
        log.Println("C.mpp3_vi_init() ok")
    case C.ERR_HI_MPI_VI_SetDevAttr:
        log.Println("C.mpp3_vi_init() ERR_HI_MPI_VI_SetDevAttr() error ", errorCode)
    case C.ERR_HI_MPI_VI_EnableDev:
        log.Println("C.mpp3_vi_init() ERR_HI_MPI_VI_EnableDev() error ", errorCode)
    case C.ERR_HI_MPI_VI_SetChnAttr:
        log.Println("C.mpp3_vi_init() ERR_HI_MPI_VI_SetChnAttr() error ", errorCode)
    case C.ERR_HI_MPI_VI_EnableChn:
        log.Println("C.mpp3_vi_init() ERR_HI_MPI_VI_EnableChn() error ", errorCode)        
    default:
        panic("Unexpected return of C.mpp3_vi_init()")
    }
}