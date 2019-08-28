#include "../hisi_external.h"
#include "../hisi_utils.h"

#include "hi3516av200_ko.h"
#include "hi3516av200_mpp.h"

#include "hi3516av200_cmos.h"

#include <string.h>

int hi3516av200_sys_init();
int hi3516av200_mipi_init();
int hi3516av200_isp_init();
int hi3516av200_vi_init();
int hi3516av200_vpss_init();

int hisi_init(struct cmos_params * cp) {
    //int error_code = 0;

    if(hi3516av200_sys_init()   != ERR_NONE) return ERR_INTERNAL;
    if(hi3516av200_mipi_init()  != ERR_NONE) return ERR_INTERNAL;
    if(hi3516av200_isp_init()   != ERR_NONE) return ERR_INTERNAL;
    if(hi3516av200_vi_init()    != ERR_NONE) return ERR_INTERNAL;
    if(hi3516av200_vpss_init()  != ERR_NONE) return ERR_INTERNAL;

    return ERR_NONE;
}

int hi3516av200_sys_init() {
    return ERR_NONE;
}

int hi3516av200_mipi_init() {
    return ERR_NONE;
}

int hi3516av200_isp_init() {
    return ERR_NONE;
}

int hi3516av200_vi_init() {
    int error_code;

    VI_DEV_ATTR_S  stViDevAttr;

    memset(&stViDevAttr,0,sizeof(stViDevAttr));
    memcpy(&stViDevAttr, &DEV_ATTR_LVDS_BASE, sizeof(stViDevAttr));

    stViDevAttr.stDevRect.s32Y                              = 0;
    stViDevAttr.stDevRect.u32Width                          = 3840;
    stViDevAttr.stDevRect.u32Height                         = 2160;
    stViDevAttr.stBasAttr.stSacleAttr.stBasSize.u32Width    = 3840;
    stViDevAttr.stBasAttr.stSacleAttr.stBasSize.u32Height   = 2160;
    stViDevAttr.stBasAttr.stSacleAttr.bCompress             = HI_FALSE;

    error_code = HI_MPI_VI_SetDevAttr(0, &stViDevAttr);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: HI_MPI_VI_SetDevAttr failed with %#x!\n", error_code);
        return ERR_MPP;
    }

    error_code = HI_MPI_VI_EnableDev(0);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: HI_MPI_VI_EnableDev failed with %#x!\n", error_code);
        return ERR_MPP;
    }

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
    stChnAttr.enPixFormat           = PIXEL_FORMAT_YUV_SEMIPLANAR_420;   /* sp420 or sp422 */

    stChnAttr.bMirror       = HI_FALSE;
    stChnAttr.bFlip         = HI_FALSE;

    stChnAttr.s32SrcFrameRate       = 30;
    stChnAttr.s32DstFrameRate       = 30;
    stChnAttr.enCompressMode        = COMPRESS_MODE_NONE;

    error_code = HI_MPI_VI_SetChnAttr(0, &stChnAttr);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: HI_MPI_VI_SetChnAttr failed with %#x!\n", error_code);
        return ERR_MPP;
    }

    error_code = HI_MPI_VI_EnableChn(0);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: HI_MPI_VI_EnableChn failed with %#x!\n", error_code);
        return ERR_MPP;
    }



    return ERR_NONE;
}

////////////////////////////////////////////////////////////////////////////////

int hi3516av200_vpss_init() {
    int error_code;

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

    error_code = HI_MPI_VPSS_CreateGrp(0, &stVpssGrpAttr);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: HI_MPI_VPSS_CreateGrp failed with %#x!\n", error_code);
        return ERR_MPP;
    }

    error_code = HI_MPI_VPSS_StartGrp(0);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: HI_MPI_VPSS_StartGrp failed with %#x\n", error_code);
        return -1;
    }

    MPP_CHN_S stSrcChn;
    MPP_CHN_S stDestChn;

    stSrcChn.enModId  = HI_ID_VIU;
    stSrcChn.s32DevId = 0;
    stSrcChn.s32ChnId = 0;

    stDestChn.enModId  = HI_ID_VPSS;
    stDestChn.s32DevId = 0;
    stDestChn.s32ChnId = 0;

    error_code = HI_MPI_SYS_Bind(&stSrcChn, &stDestChn);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: HI_MPI_SYS_Bind failed with %#x!\n", error_code);
        return ERR_MPP;
    }

    return ERR_NONE;
}

