//+build hi3516cv300 hi3516av200

package vpss

/*
#include "../include/hi3516av200_mpp.h"

#define ERR_NONE                    0
#define ERR_HI_MPI_VPSS_CreateGrp   2
#define ERR_HI_MPI_VPSS_StartGrp    3
#define ERR_HI_MPI_SYS_Bind         4

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
    stVpssGrpAttr.bStitchBlendEn    = HI_FALSE;

    stVpssGrpAttr.stNrAttr.enNrType                         = VPSS_NR_TYPE_VIDEO;
    stVpssGrpAttr.stNrAttr.stNrVideoAttr.enNrRefSource      = VPSS_NR_REF_FROM_RFR;//VPSS_NR_REF_FROM_CHN0, VPSS_NR_REF_FROM_SRC
    stVpssGrpAttr.stNrAttr.stNrVideoAttr.enNrOutputMode     = VPSS_NR_OUTPUT_NORMAL;//VPSS_NR_OUTPUT_DELAY NORMAL
    stVpssGrpAttr.stNrAttr.u32RefFrameNum                   = 2;

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
