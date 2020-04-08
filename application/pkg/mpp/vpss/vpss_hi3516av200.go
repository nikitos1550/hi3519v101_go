//+build arm
//+build hi3516av200

package vpss

/*
#include "../include/mpp_v3.h"

#include <string.h>

#define ERR_NONE                    0
#define ERR_MPP                     2
#define ERR_HI_MPI_VPSS_CreateGrp   3
#define ERR_HI_MPI_VPSS_StartGrp    4
#define ERR_HI_MPI_SYS_Bind         5

int mpp3_vpss_init(unsigned int *error_code) {
    *error_code = 0;

    VPSS_GRP_ATTR_S stVpssGrpAttr;

    stVpssGrpAttr.u32MaxW           = 3840;
    stVpssGrpAttr.u32MaxH           = 2160;
    stVpssGrpAttr.bIeEn             = HI_FALSE;
    stVpssGrpAttr.bNrEn             = HI_TRUE;
    stVpssGrpAttr.bHistEn           = HI_FALSE;
    stVpssGrpAttr.bDciEn            = HI_FALSE;
    stVpssGrpAttr.enDieMode         = VPSS_DIE_MODE_NODIE;
    stVpssGrpAttr.enPixFmt          = PIXEL_FORMAT_YUV_SEMIPLANAR_420;//SAMPLE_PIXEL_FORMAT;
    #ifdef HI3516AV200
    stVpssGrpAttr.bStitchBlendEn    = HI_FALSE;
    #endif

    #ifdef HI3516AV200
    stVpssGrpAttr.stNrAttr.enNrType                         = VPSS_NR_TYPE_VIDEO;
    stVpssGrpAttr.stNrAttr.stNrVideoAttr.enNrRefSource      = VPSS_NR_REF_FROM_RFR;//VPSS_NR_REF_FROM_CHN0, VPSS_NR_REF_FROM_SRC
    stVpssGrpAttr.stNrAttr.stNrVideoAttr.enNrOutputMode     = VPSS_NR_OUTPUT_NORMAL;//VPSS_NR_OUTPUT_DELAY NORMAL
    stVpssGrpAttr.stNrAttr.u32RefFrameNum                   = 2;
    #endif
    
    //    stVpssGrpAttr.u32MaxW = global_width;
    //    stVpssGrpAttr.u32MaxH = global_height;
    //    stVpssGrpAttr.bIeEn = HI_FALSE;
    //    stVpssGrpAttr.bNrEn = HI_TRUE;//HI_FALSE;//HI_TRUE;
    //    stVpssGrpAttr.bHistEn = HI_FALSE;
    //    stVpssGrpAttr.bSharpenEn = HI_FALSE;//HI_TRUE;
    //    stVpssGrpAttr.enDieMode = VPSS_DIE_MODE_NODIE;
    //    stVpssGrpAttr.enPixFmt = PIXEL_FORMAT_YUV_SEMIPLANAR_420;
    

    *error_code = HI_MPI_VPSS_CreateGrp(0, &stVpssGrpAttr);
    if (*error_code != HI_SUCCESS) return ERR_HI_MPI_VPSS_CreateGrp;

    *error_code = HI_MPI_VPSS_StartGrp(0);
    if (*error_code != HI_SUCCESS) return ERR_HI_MPI_VPSS_StartGrp;

    MPP_CHN_S stSrcChn;
    MPP_CHN_S stDestChn;

    stSrcChn.enModId  = HI_ID_VIU;
    stSrcChn.s32DevId = 0;
    stSrcChn.s32ChnId = 0;

    stDestChn.enModId  = HI_ID_VPSS;
    stDestChn.s32DevId = 0;
    stDestChn.s32ChnId = 0;

    *error_code = HI_MPI_SYS_Bind(&stSrcChn, &stDestChn);
    if (*error_code != HI_SUCCESS) return ERR_HI_MPI_SYS_Bind;

    return ERR_NONE;
}

int mpp3_vpss_sample_channel0(unsigned int *error_code) {
    *error_code = 0;

    VPSS_CHN_ATTR_S stVpssChnAttr;
    VPSS_CHN_MODE_S stVpssChnMode;

    stVpssChnMode.enChnMode      = VPSS_CHN_MODE_USER;
    stVpssChnMode.bDouble        = HI_FALSE;
    stVpssChnMode.enPixelFormat  = PIXEL_FORMAT_YUV_SEMIPLANAR_420;
    stVpssChnMode.u32Width       = 3840;
    stVpssChnMode.u32Height      = 2160;
    stVpssChnMode.enCompressMode = COMPRESS_MODE_NONE; //COMPRESS_MODE_SEG;

    memset(&stVpssChnAttr, 0, sizeof(stVpssChnAttr));

    stVpssChnAttr.s32SrcFrameRate = 30;
    stVpssChnAttr.s32DstFrameRate = 30;

    *error_code = HI_MPI_VPSS_SetChnAttr(0, 0, &stVpssChnAttr);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    *error_code = HI_MPI_VPSS_SetChnMode(0, 0, &stVpssChnMode);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    *error_code = HI_MPI_VPSS_EnableChn(0, 0);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    return ERR_NONE;
}
*/
import "C"

import (
    "log"
    "application/pkg/mpp/error"
)

func Init() {
    var errorCode C.uint

    switch err := C.mpp3_vpss_init(&errorCode); err {
    case C.ERR_NONE:
        log.Println("C.mpp3_vpss_init() ok")
    case C.ERR_HI_MPI_VPSS_CreateGrp:
        log.Fatal("C.mpp3_vpss_init() HI_MPI_VPSS_CreateGrp() error ", error.Resolve(int64(errorCode)))
    case C.ERR_HI_MPI_VPSS_StartGrp:
        log.Fatal("C.mpp3_vpss_init() HI_MPI_VPSS_StartGrp() error ", error.Resolve(int64(errorCode)))
    case C.ERR_HI_MPI_SYS_Bind:
        log.Fatal("C.mpp3_vpss_init() HI_MPI_SYS_Bind() error ", error.Resolve(int64(errorCode)))
    default:
        log.Fatal("Unexpected return ", err , " of C.mpp3_vpss_init()")
    }
}

func SampleChannel0() {
    var errorCode C.uint

    switch err := C.mpp3_vpss_sample_channel0(&errorCode); err {
    case C.ERR_NONE:
        log.Println("C.mpp3_vpss_sample_channel0() ok")
    case C.ERR_MPP:
        log.Fatal("C.mpp3_vpss_sample_channel0() MPP error ", error.Resolve(int64(errorCode)))
    default:
        log.Fatal("Unexpected return ", err , " of C.mpp3_vpss_sample_channel0()")
    }

}
