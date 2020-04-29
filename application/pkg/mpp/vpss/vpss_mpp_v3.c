#include "vpss.h"

#if defined(HI_MPP_V3) 
VIDEO_FRAME_INFO_S channelFrames[MAX_CHANNELS];

int mpp_vpss_init(error_in *err, mpp_vpss_init_in *in) {
    unsigned int mpp_error_code = 0;

    VPSS_GRP_ATTR_S stVpssGrpAttr;

    memset(&stVpssGrpAttr, 0, sizeof(stVpssGrpAttr));

    stVpssGrpAttr.u32MaxW           = in->width;
    stVpssGrpAttr.u32MaxH           = in->height;
    stVpssGrpAttr.bIeEn             = HI_FALSE;                         //reserved
    stVpssGrpAttr.bHistEn           = HI_FALSE;                         //reserved

    if (in->nr == 1) {
        GO_LOG_VPSS(LOGGER_TRACE, "VPSS NR on");
        stVpssGrpAttr.bNrEn = HI_TRUE;
    } else {
        GO_LOG_VPSS(LOGGER_TRACE, "VPSS NR off");
        stVpssGrpAttr.bNrEn = HI_FALSE;
    }

    #if defined(HI3516AV200)
        stVpssGrpAttr.bDciEn            = HI_FALSE;                         //reserved

        stVpssGrpAttr.stNrAttr.enNrType                         = VPSS_NR_TYPE_VIDEO;       //video or snapshot, we use video (i don`t know anything about snapshot mode)

        //VPSS_NR_REF_FROM_RFR Reconstruction frame as the reference frame
        //VPSS_NR_REF_FROM_CHN0 Channel 0 output as the reference frame
        //VPSS_NR_REF_FROM_SRC Input source picture as the reference frame
        stVpssGrpAttr.stNrAttr.stNrVideoAttr.enNrRefSource      = VPSS_NR_REF_FROM_RFR;

        //VPSS_NR_OUTPUT_NORMAL Normal output mode. There is no delay.
        //VPSS_NR_OUTPUT_DELAY Delay output mode. The output is one-frame delayed.
        //Only Hi3519 V101 supports this data structure. TODO what about hi3516av200 chip?
        //In delay output mode, the number of reference frames cannot be set to 1, the single component is not supported, and the large stream cannot be used as the reference frame.

        stVpssGrpAttr.stNrAttr.stNrVideoAttr.enNrOutputMode     = VPSS_NR_OUTPUT_NORMAL;
    
        //1 or 2 for video mode
        stVpssGrpAttr.stNrAttr.u32RefFrameNum                   = in->nr_frames; //2;
    
        stVpssGrpAttr.bStitchBlendEn    = HI_FALSE;
    #endif

    stVpssGrpAttr.enDieMode         = VPSS_DIE_MODE_NODIE;              //reserved
    stVpssGrpAttr.enPixFmt          = PIXEL_FORMAT_YUV_SEMIPLANAR_420;  //yuv420 or yuv422

    DO_OR_RETURN_MPP(HI_MPI_VPSS_CreateGrp, 0, &stVpssGrpAttr);

    DO_OR_RETURN_MPP(HI_MPI_VPSS_StartGrp, 0);

    MPP_CHN_S stSrcChn;
    MPP_CHN_S stDestChn;

    stSrcChn.enModId  = HI_ID_VIU;
    stSrcChn.s32DevId = 0;
    stSrcChn.s32ChnId = 0;

    stDestChn.enModId  = HI_ID_VPSS;
    stDestChn.s32DevId = 0;
    stDestChn.s32ChnId = 0;

    DO_OR_RETURN_MPP(HI_MPI_SYS_Bind, &stSrcChn, &stDestChn);

    return ERR_NONE;
}

int mpp_vpss_create_channel(error_in *err, mpp_vpss_create_channel_in * in) {
    unsigned int mpp_error_code = 0;

    VPSS_CHN_ATTR_S stVpssChnAttr;

    memset(&stVpssChnAttr, 0, sizeof(stVpssChnAttr));

    //typedef struct hiVPSS_CHN_ATTR_S
    //{
    //    HI_BOOL bSpEn;            //SP enable. It must be set to HI_FALSE.
    //    HI_BOOL bBorderEn;        //Border enable. It must be set to HI_FALSE.
    //    HI_BOOL bMirror;
    //    HI_BOOL bFlip;
    //    HI_S32 s32SrcFrameRate;
    //    HI_S32 s32DstFrameRate;
    //    BORDER_S
    //    stBorder;
    //}VPSS_CHN_ATTR_S;

    stVpssChnAttr.s32SrcFrameRate = in->vi_fps;
    stVpssChnAttr.s32DstFrameRate = in->fps;

    DO_OR_RETURN_MPP(HI_MPI_VPSS_SetChnAttr, 0, in->channel_id, &stVpssChnAttr);

    VPSS_CHN_MODE_S stVpssChnMode;

    stVpssChnMode.enChnMode      = VPSS_CHN_MODE_USER;
    stVpssChnMode.bDouble        = HI_FALSE;
    stVpssChnMode.enPixelFormat  = PIXEL_FORMAT_YUV_SEMIPLANAR_420;
    stVpssChnMode.u32Width       = in->width;
    stVpssChnMode.u32Height      = in->height;
    stVpssChnMode.enCompressMode = COMPRESS_MODE_NONE; //COMPRESS_MODE_SEG;

    DO_OR_RETURN_MPP(HI_MPI_VPSS_SetChnMode, 0, in->channel_id, &stVpssChnMode);

    HI_U32 u32Depth = 1; //TODO

    DO_OR_RETURN_MPP(HI_MPI_VPSS_SetDepth, 0, in->channel_id, u32Depth);

    DO_OR_RETURN_MPP(HI_MPI_VPSS_EnableChn, 0, in->channel_id);

    return ERR_NONE;
}

int mpp_vpss_destroy_channel(error_in * err, mpp_vpss_destroy_channel_in *in) {
    unsigned int mpp_error_code = 0;

    DO_OR_RETURN_MPP(HI_MPI_VPSS_DisableChn, 0, in->channel_id);

    return ERR_NONE;
}

int mpp_receive_frame(error_in *err, unsigned int channel_id, void** frame) {
    unsigned int mpp_error_code;

    DO_OR_RETURN_MPP(HI_MPI_VPSS_GetChnFrame, 0, channel_id, &channelFrames[channel_id], -1); //blocking mode call

    *frame = &channelFrames[channel_id];
    return ERR_NONE;
}

int mpp_release_frame(error_in *err, unsigned int channel_id) {
    unsigned int mpp_error_code;

    DO_OR_RETURN_MPP(HI_MPI_VPSS_ReleaseChnFrame, 0, channel_id, &channelFrames[channel_id]);

    return ERR_NONE;
}

#endif

