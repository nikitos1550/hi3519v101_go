#include "vpss.h"

int mpp_vpss_init(error_in *err, mpp_vpss_init_in *in) {

    VPSS_GRP_ATTR_S stVpssGrpAttr;

    memset(&stVpssGrpAttr, 0, sizeof(stVpssGrpAttr));

    #if HI_MPP == 4
        stVpssGrpAttr.enDynamicRange                = DYNAMIC_RANGE_SDR8;
        stVpssGrpAttr.enPixelFormat                 = PIXEL_FORMAT_YVU_SEMIPLANAR_420;
        stVpssGrpAttr.stFrameRate.s32SrcFrameRate   = 30;//TODO
        stVpssGrpAttr.stFrameRate.s32DstFrameRate   = 30;//TODO
    #endif

    #if HI_MPP == 1 \
        || HI_MPP == 2 \
        || HI_MPP == 3
        stVpssGrpAttr.enDieMode         = VPSS_DIE_MODE_NODIE;              //reserved
        stVpssGrpAttr.enPixFmt          = PIXEL_FORMAT_YUV_SEMIPLANAR_420;  //yuv420 or yuv422
        stVpssGrpAttr.bHistEn           = HI_FALSE;                         //reserved
    #endif

    #if HI_MPP == 1
        stVpssGrpAttr.bDrEn     = HI_FALSE;
        stVpssGrpAttr.bDbEn     = HI_FALSE;
    #endif

    #if HI_MPP == 1
        stVpssGrpAttr.bIeEn             = HI_TRUE;
    #elif HI_MPP == 2 \
        || HI_MPP == 3
        stVpssGrpAttr.bIeEn             = HI_FALSE;                         //reserved
    #endif

    stVpssGrpAttr.u32MaxW           = in->width;
    stVpssGrpAttr.u32MaxH           = in->height;

    #if defined(HI3516AV100) \
        || defined(HI3516AV200)
        stVpssGrpAttr.bDciEn            = HI_FALSE;                         //reserved
    #endif

    if (in->nr == 1) {
        GO_LOG_VPSS(LOGGER_TRACE, "VPSS NR on");
        stVpssGrpAttr.bNrEn = HI_TRUE;
    } else {
        GO_LOG_VPSS(LOGGER_TRACE, "VPSS NR off");
        stVpssGrpAttr.bNrEn = HI_FALSE;
    }

    #if HI_MPP == 4
        stVpssGrpAttr.stNrAttr.enNrType             = VPSS_NR_TYPE_VIDEO;
        stVpssGrpAttr.stNrAttr.enNrMotionMode       = NR_MOTION_MODE_NORMAL;
        stVpssGrpAttr.stNrAttr.enCompressMode       = COMPRESS_MODE_FRAME;
    #endif

    #if defined(HI3516AV200)
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
    
        stVpssGrpAttr.bStitchBlendEn                            = HI_FALSE;
    #endif

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VPSS_CreateGrp, 0, &stVpssGrpAttr);

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VPSS_StartGrp, 0);

    MPP_CHN_S stSrcChn;
    MPP_CHN_S stDestChn;

    #if HI_MPP == 1 \
        || HI_MPP == 2 \
        || HI_MPP == 3
        stSrcChn.enModId  = HI_ID_VIU;
    #elif HI_MPP == 4
        stSrcChn.enModId   = HI_ID_VI;
    #endif
    stSrcChn.s32DevId = 0;
    stSrcChn.s32ChnId = 0;

    stDestChn.enModId  = HI_ID_VPSS;
    stDestChn.s32DevId = 0;
    stDestChn.s32ChnId = 0;

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_SYS_Bind, &stSrcChn, &stDestChn);

    return ERR_NONE;
}

int mpp_vpss_create_channel(error_in *err, mpp_vpss_create_channel_in * in) {
    VPSS_CHN_ATTR_S stVpssChnAttr;

    memset(&stVpssChnAttr, 0, sizeof(stVpssChnAttr));

    #if HI_MPP == 1
        stVpssChnAttr.bSpEn             = HI_FALSE;
    #endif

    #if HI_MPP == 2 \
        || HI_MPP == 3
        stVpssChnAttr.s32SrcFrameRate = in->vi_fps;
        stVpssChnAttr.s32DstFrameRate = in->fps;
    #endif

    #if HI_MPP == 4
        stVpssChnAttr.u32Width                     = in->width;
        stVpssChnAttr.u32Height                    = in->height;
        stVpssChnAttr.enChnMode                    = VPSS_CHN_MODE_USER;
        stVpssChnAttr.enCompressMode               = COMPRESS_MODE_NONE;//COMPRESS_MODE_SEG;
        stVpssChnAttr.enDynamicRange               = DYNAMIC_RANGE_SDR8;
        stVpssChnAttr.enPixelFormat                = PIXEL_FORMAT_YVU_SEMIPLANAR_420;
        stVpssChnAttr.stFrameRate.s32SrcFrameRate  = in->vi_fps;
        stVpssChnAttr.stFrameRate.s32DstFrameRate  = in->fps;

        stVpssChnAttr.u32Depth                     = in->depth;
        stVpssChnAttr.bMirror                      = HI_FALSE;
        stVpssChnAttr.bFlip                        = HI_FALSE;

        stVpssChnAttr.enVideoFormat                = VIDEO_FORMAT_LINEAR;
        stVpssChnAttr.stAspectRatio.enMode         = ASPECT_RATIO_NONE;
    #endif

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VPSS_SetChnAttr, 0, in->id, &stVpssChnAttr);

    #if HI_MPP == 1 \
        || HI_MPP == 2 \
        || HI_MPP == 3
        VPSS_CHN_MODE_S stVpssChnMode;

        memset(&stVpssChnMode, 0, sizeof(stVpssChnMode));

        stVpssChnMode.enChnMode      = VPSS_CHN_MODE_USER;
        stVpssChnMode.bDouble        = HI_FALSE;
        stVpssChnMode.enPixelFormat  = PIXEL_FORMAT_YUV_SEMIPLANAR_420;
        stVpssChnMode.u32Width       = in->width;
        stVpssChnMode.u32Height      = in->height;
        #if HI_MPP == 2 \
            || HI_MPP == 3
            stVpssChnMode.enCompressMode = COMPRESS_MODE_NONE; //COMPRESS_MODE_SEG;
        #endif

        DO_OR_RETURN_ERR_MPP(err, HI_MPI_VPSS_SetChnMode, 0, in->id, &stVpssChnMode);

        HI_U32 u32Depth = in->depth;

        DO_OR_RETURN_ERR_MPP(err, HI_MPI_VPSS_SetDepth, 0, in->id, u32Depth);
    #endif

    if (in->crop_width > 0 && in->crop_height > 0) {
        #if HI_MPP == 1
            //TODO crop per channel isn`t supported
        #elif HI_MPP == 2 || \
            HI_MPP == 3 || \
            HI_MPP == 4
            VPSS_CROP_INFO_S stCropInfo;

            stCropInfo.bEnable = HI_TRUE;
            stCropInfo.enCropCoordinate = VPSS_CROP_RATIO_COOR;
            stCropInfo.stCropRect.s32X = in->crop_x;
            stCropInfo.stCropRect.s32Y = in->crop_y;
            stCropInfo.stCropRect.u32Width = in->crop_width;
            stCropInfo.stCropRect.u32Height = in->crop_height;

            DO_OR_RETURN_ERR_MPP(err, HI_MPI_VPSS_SetChnCrop, 0, in->id, &stCropInfo);
        #endif
    }

    /*
    #if HI_MPP == 4
        //TEST!!! TODO
        VPSS_LOW_DELAY_INFO_S stLowDelayInfo;

        stLowDelayInfo.bEnable = HI_TRUE;
        stLowDelayInfo.u32LineCnt = 16;

        DO_OR_RETURN_ERR_MPP(err, HI_MPI_VPSS_SetLowDelayAttr, 0, in->id, &stLowDelayInfo);
        //HI_MPI_VPSS_SetLowDelayAttr(VPSS_GRP VpssGrp, VPSS_CHN VpssChn, const VPSS_LOW_DELAY_INFO_S *pstLowDelayInfo);
    #endif
    */

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VPSS_EnableChn, 0, in->id);

    return ERR_NONE;
}

int mpp_vpss_destroy_channel(error_in * err, mpp_vpss_destroy_channel_in *in) {
    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VPSS_DisableChn, 0, in->id);

    return ERR_NONE;
}

int mpp_vpss_change_channel_depth(error_in * err, mpp_vpss_change_channel_depth_in *in) {
    #if HI_MPP == 1 \
        || HI_MPP == 2 \
        || HI_MPP == 3
        HI_U32 u32Depth = in->depth;
        DO_OR_RETURN_ERR_MPP(err, HI_MPI_VPSS_SetDepth, 0, in->id, u32Depth);
    #elif HI_MPP == 4
        VPSS_CHN_ATTR_S stVpssChnAttr;
        DO_OR_RETURN_ERR_MPP(err, HI_MPI_VPSS_GetChnAttr, 0, in->id, &stVpssChnAttr);
        stVpssChnAttr.u32Depth = in->depth;
        DO_OR_RETURN_ERR_MPP(err, HI_MPI_VPSS_SetChnAttr, 0, in->id, &stVpssChnAttr);
    #endif

    return ERR_NONE;
}

VIDEO_FRAME_INFO_S channelFrames[VPSS_MAX_PHY_CHN_NUM];

int mpp_receive_frame(error_in *err, unsigned int id, void **frame, unsigned long long *pts, unsigned int wait) {
//int mpp_receive_frame(error_in *err, unsigned int id, VIDEO_FRAME_INFO_S *frame, unsigned long long *pts, unsigned int wait) {
//int mpp_receive_frame(error_in *err, unsigned int id, void *frame, unsigned long long *pts, unsigned int wait) {
    #if HI_MPP == 1
        DO_OR_RETURN_ERR_MPP(err, HI_MPI_VPSS_UserGetFrame, 0, id, &channelFrames[id])      //don`t have block mode
    #elif HI_MPP == 2 \
        || HI_MPP == 3 \
        || HI_MPP == 4
        DO_OR_RETURN_ERR_MPP(err, HI_MPI_VPSS_GetChnFrame, 0, id, &channelFrames[id], wait);  //blocking mode call
    #endif

    *frame = &channelFrames[id];//TODO

    #if HI_MPP == 1 \
        || HI_MPP == 2 \
        || HI_MPP == 3
        *pts = channelFrames[id].stVFrame.u64pts;
    #elif HI_MPP == 4
        *pts = channelFrames[id].stVFrame.u64PTS;
    #endif

    return ERR_NONE;
}

int mpp_release_frame(error_in *err, unsigned int id) {
    #if HI_MPP == 1
        DO_OR_RETURN_ERR_MPP(err, HI_MPI_VPSS_UserReleaseFrame, 0, id, &channelFrames[id]);
    #elif HI_MPP == 2 \
        || HI_MPP == 3 \
        || HI_MPP == 4
        DO_OR_RETURN_ERR_MPP(err, HI_MPI_VPSS_ReleaseChnFrame, 0, id, &channelFrames[id]);
    #endif

    return ERR_NONE;
}
