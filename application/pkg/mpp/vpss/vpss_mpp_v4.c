#include "vpss.h"

#if defined(HI_MPP_V4) 
VIDEO_FRAME_INFO_S channelFrames[VPSS_MAX_PHY_CHN_NUM];

int mpp_vpss_init(error_in *err, mpp_vpss_init_in *in) {

    VPSS_GRP_ATTR_S stVpssGrpAttr;

    memset(&stVpssGrpAttr, 0, sizeof(stVpssGrpAttr));

    stVpssGrpAttr.enDynamicRange                = DYNAMIC_RANGE_SDR8;
    stVpssGrpAttr.enPixelFormat                 = PIXEL_FORMAT_YVU_SEMIPLANAR_420;
    stVpssGrpAttr.u32MaxW                       = in->width;
    stVpssGrpAttr.u32MaxH                       = in->height;
    stVpssGrpAttr.stFrameRate.s32SrcFrameRate   = 30;//TODO
    stVpssGrpAttr.stFrameRate.s32DstFrameRate   = 30;//TODO


    if (in->nr == 1) {
        GO_LOG_VPSS(LOGGER_TRACE, "VPSS NR on");
        stVpssGrpAttr.bNrEn = HI_TRUE;
    } else {
        GO_LOG_VPSS(LOGGER_TRACE, "VPSS NR off");
        stVpssGrpAttr.bNrEn = HI_FALSE;
    }

    //stVpssGrpAttr.bNrEn = HI_FALSE; //TESTING

    stVpssGrpAttr.stNrAttr.enNrType             = VPSS_NR_TYPE_VIDEO;
    stVpssGrpAttr.stNrAttr.enNrMotionMode       = NR_MOTION_MODE_NORMAL;
    stVpssGrpAttr.stNrAttr.enCompressMode       = COMPRESS_MODE_FRAME;

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VPSS_CreateGrp, 0, &stVpssGrpAttr);

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VPSS_StartGrp, 0);
    
    MPP_CHN_S stSrcChn;
    MPP_CHN_S stDestChn;

    stSrcChn.enModId   = HI_ID_VI;
    stSrcChn.s32DevId  = 0;
    stSrcChn.s32ChnId  = 0;

    stDestChn.enModId  = HI_ID_VPSS;
    stDestChn.s32DevId = 0;
    stDestChn.s32ChnId = 0;

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_SYS_Bind, &stSrcChn, &stDestChn);

    return ERR_NONE;
}

int mpp_vpss_create_channel(error_in *err, mpp_vpss_create_channel_in * in) {

	VPSS_CHN_ATTR_S stVpssChnAttr;

    stVpssChnAttr.u32Width                     = in->width;
    stVpssChnAttr.u32Height                    = in->height;
    stVpssChnAttr.enChnMode                    = VPSS_CHN_MODE_USER;
    stVpssChnAttr.enCompressMode               = COMPRESS_MODE_NONE;//COMPRESS_MODE_SEG;
    stVpssChnAttr.enDynamicRange               = DYNAMIC_RANGE_SDR8;
    stVpssChnAttr.enPixelFormat                = PIXEL_FORMAT_YVU_SEMIPLANAR_420;
    stVpssChnAttr.stFrameRate.s32SrcFrameRate  = in->vi_fps;
    stVpssChnAttr.stFrameRate.s32DstFrameRate  = in->fps;

    stVpssChnAttr.u32Depth                     = 1;
    stVpssChnAttr.bMirror                      = HI_FALSE;
    stVpssChnAttr.bFlip                        = HI_FALSE;

    stVpssChnAttr.enVideoFormat                = VIDEO_FORMAT_LINEAR;
    stVpssChnAttr.stAspectRatio.enMode         = ASPECT_RATIO_NONE;

	DO_OR_RETURN_ERR_MPP(err, HI_MPI_VPSS_SetChnAttr, 0, in->channel_id, &stVpssChnAttr)

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VPSS_EnableChn, 0, in->channel_id);

    return ERR_NONE;
}

int mpp_vpss_destroy_channel(error_in * err, mpp_vpss_destroy_channel_in *in) {
    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VPSS_DisableChn, 0, in->channel_id);

    return ERR_NONE;
}

int mpp_receive_frame(error_in *err, unsigned int channel_id, void** frame) {
    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VPSS_GetChnFrame, 0, channel_id, &channelFrames[channel_id], -1); //blocking mode call

    *frame = &channelFrames[channel_id];
    return ERR_NONE;
}

int mpp_release_frame(error_in *err, unsigned int channel_id) {
    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VPSS_ReleaseChnFrame, 0, channel_id, &channelFrames[channel_id]);

    return ERR_NONE;
}

#endif
