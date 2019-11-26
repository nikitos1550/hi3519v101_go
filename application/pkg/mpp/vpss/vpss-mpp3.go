//+build hi3516cv300 hi3516av200

package vpss

/*
int hi3516av200_vpss_init(struct capture_params * cp) {
    int error_code;

    VPSS_GRP_ATTR_S stVpssGrpAttr;

    stVpssGrpAttr.u32MaxW           = cp->width;
    stVpssGrpAttr.u32MaxH           = cp->width;
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
*/
import "C"

