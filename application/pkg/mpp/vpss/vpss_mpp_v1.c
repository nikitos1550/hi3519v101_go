#include "vpss.h"

#if defined(HI_MPP_V1)
VIDEO_FRAME_INFO_S channelFrames[MAX_CHANNELS];

int mpp_vpss_init(error_in *err, mpp_vpss_init_in *in) {

    VPSS_GRP_ATTR_S stGrpAttr;

    stGrpAttr.u32MaxW   = in->width;
    stGrpAttr.u32MaxH   = in->height;
    stGrpAttr.bDrEn     = HI_FALSE;
    stGrpAttr.bDbEn     = HI_FALSE;
    stGrpAttr.bIeEn     = HI_TRUE;
    stGrpAttr.bNrEn     = HI_TRUE;
    stGrpAttr.bHistEn   = HI_FALSE;
    stGrpAttr.enDieMode = VPSS_DIE_MODE_NODIE; //VPSS_DIE_MODE_AUTO;
    stGrpAttr.enPixFmt  = PIXEL_FORMAT_YUV_SEMIPLANAR_420;

    //stGrpVpssAttr.u32MaxW = 720;
    //stGrpVpssAttr.u32MaxH = 576;
    //stGrpVpssAttr.bDrEn = HI_FALSE;
    //stGrpVpssAttr.bDbEn = HI_FALSE;
    //stGrpVpssAttr.bIeEn = HI_FALSE;
    //stGrpVpssAttr.bNrEn = HI_FALSE;
    //stGrpVpssAttr.bHistEn = HI_FALSE;
    //stGrpVpssAttr.enDieMode = VPSS_DIE_MODE_NODIE;
    //stGrpVpssAttr.enPixFmt = PIXEL_FORMAT_YUV_SEMIPLANAR_422;

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VPSS_CreateGrp, 0, &stGrpAttr);

    //VPSS_GRP_PARAM_S stVpssParam;
    //
    //DO_OR_RETURN_ERR_MPP(err, HI_MPI_VPSS_GetGrpParam, 0, &stVpssParam);
    //    
    //stVpssParam.u32MotionThresh = 0;
    //
    //DO_OR_RETURN_ERR_MPP(err, HI_MPI_VPSS_SetGrpParam, 0, &stVpssParam);

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VPSS_StartGrp, 0);

    MPP_CHN_S stSrcChn;
    MPP_CHN_S stDestChn;

    stSrcChn.enModId  = HI_ID_VIU;
    stSrcChn.s32DevId = 0;
    stSrcChn.s32ChnId = 0;
 
    stDestChn.enModId  = HI_ID_VPSS;
    stDestChn.s32DevId = 0;
    stDestChn.s32ChnId = 0;

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_SYS_Bind, &stSrcChn, &stDestChn);

    return ERR_NONE;
}

int mpp_vpss_create_channel(error_in *err, mpp_vpss_create_channel_in * in) {
    
    VPSS_CHN_ATTR_S stChnAttr;

    memset(&stChnAttr, 0, sizeof(VPSS_CHN_ATTR_S));

    stChnAttr.bSpEn             = HI_FALSE;
    //stChnAttr.bBorderEn         = HI_FALSE;
    //stChnAttr.bMirror           = HI_FALSE;
    //stChnAttr.bFlip             = HI_FALSE;
    //stChnAttr.s32SrcFrameRate   = in->vi_fps;
    //stChnAttr.s32DstFrameRate   = in->fps;

    //stChnAttr.stBorder.u32TopWidth = 0;
    //stChnAttr.stBorder.u32BottomWidth = 0;
    //stChnAttr.stBorder.u32LeftWidth = 0;
    //stChnAttr.stBorder.u32RightWidth = 0;
    //stChnAttr.stBorder.u32Color = 0;

    //stChnAttr.bFrameEn = HI_TRUE;
    //stChnAttr.stFrame.u32Color[VPSS_FRAME_WORK_LEFT] = 0xff00;
    //stChnAttr.stFrame.u32Color[VPSS_FRAME_WORK_RIGHT] = 0xff00;
    //stChnAttr.stFrame.u32Color[VPSS_FRAME_WORK_BOTTOM] = 0xff00;
    //stChnAttr.stFrame.u32Color[VPSS_FRAME_WORK_TOP] = 0xff00;
    //stChnAttr.stFrame.u32Width[VPSS_FRAME_WORK_LEFT] = 2;
    //stChnAttr.stFrame.u32Width[VPSS_FRAME_WORK_RIGHT] = 2;
    //stChnAttr.stFrame.u32Width[VPSS_FRAME_WORK_TOP] = 2;
    //stChnAttr.stFrame.u32Width[VPSS_FRAME_WORK_BOTTOM] = 2;
    
    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VPSS_SetChnAttr, 0, in->channel_id, &stChnAttr);

    VPSS_CHN_MODE_S stVpssChnMode;

    memset(&stVpssChnMode, 0, sizeof(VPSS_CHN_MODE_S));

    stVpssChnMode.enChnMode         = VPSS_CHN_MODE_USER;
    stVpssChnMode.u32Width          = in->width;
    stVpssChnMode.u32Height         = in->height;
    stVpssChnMode.bDouble           = HI_FALSE;
    stVpssChnMode.enPixelFormat     = PIXEL_FORMAT_YUV_SEMIPLANAR_420;
    //stVpssChnMode.enCompressMode    = COMPRESS_MODE_NONE;

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VPSS_SetChnMode, 0, in->channel_id, &stVpssChnMode);

    HI_U32 u32Depth = 2; //TODO

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VPSS_SetDepth, 0, in->channel_id, u32Depth);

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VPSS_EnableChn, 0, in->channel_id);

    return ERR_NONE;
}

int mpp_vpss_destroy_channel(error_in * err, mpp_vpss_destroy_channel_in *in) {

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VPSS_DisableChn, 0, in->channel_id);

    return ERR_NONE;
}

int mpp_receive_frame(error_in *err, unsigned int channel_id, void** frame) {

    //HI_S32 HI_MPI_VPSS_UserGetFrame(VPSS_GRP VpssGrp, VPSS_CHN VpssChn, VIDEO_FRAME_INFO_S *pstVideoFrame)

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VPSS_UserGetFrame, 0, channel_id, &channelFrames[channel_id])

    //DO_OR_RETURN_ERR_MPP(err, HI_MPI_VPSS_GetChnFrame, 0, channel_id, &channelFrames[channel_id], -1); //blocking mode call

    *frame = &channelFrames[channel_id];

    return ERR_NONE;
}

int mpp_release_frame(error_in *err, unsigned int channel_id) {

    //HI_S32 HI_MPI_VPSS_UserReleaseFrame (VPSS_GRP VpssGrp, VPSS_CHN VpssChn, VIDEO_FRAME_INFO_S *pstVideoFrame)

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VPSS_UserReleaseFrame, 0, channel_id, &channelFrames[channel_id]);

    //DO_OR_RETURN_ERR_MPP(err, HI_MPI_VPSS_ReleaseChnFrame, 0, channel_id, &channelFrames[channel_id]);

    return ERR_NONE;
}
#endif
