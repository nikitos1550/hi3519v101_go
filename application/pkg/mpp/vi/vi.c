#include "vi.h"

#include <stdint.h>

static void mpp_videv_sync(VI_SYNC_CFG_S *stSynCfg, mpp_vi_init_in * in) {
    stSynCfg->enVsync                       = in->dc_sync_attrs.v_sync;
    stSynCfg->enVsyncNeg                    = in->dc_sync_attrs.v_sync_neg;
    stSynCfg->enHsync                       = in->dc_sync_attrs.h_sync; 
    stSynCfg->enHsyncNeg                    = in->dc_sync_attrs.h_sync_neg; 
    stSynCfg->enVsyncValid                  = in->dc_sync_attrs.v_sync_valid;
    stSynCfg->enVsyncValidNeg               = in->dc_sync_attrs.v_sync_valid_neg;

    stSynCfg->stTimingBlank.u32HsyncHfb     = in->dc_sync_attrs.timing_hfb;
    stSynCfg->stTimingBlank.u32HsyncAct     = in->dc_sync_attrs.timing_act;
    stSynCfg->stTimingBlank.u32HsyncHbb     = in->dc_sync_attrs.timing_hbb;
    stSynCfg->stTimingBlank.u32VsyncVfb     = in->dc_sync_attrs.timing_vfb;
    stSynCfg->stTimingBlank.u32VsyncVact    = in->dc_sync_attrs.timing_vact;
    stSynCfg->stTimingBlank.u32VsyncVbb     = in->dc_sync_attrs.timing_vbb;
    stSynCfg->stTimingBlank.u32VsyncVbfb    = in->dc_sync_attrs.timing_vbfb;
    stSynCfg->stTimingBlank.u32VsyncVbact   = in->dc_sync_attrs.timing_vbact;
    stSynCfg->stTimingBlank.u32VsyncVbbb    = in->dc_sync_attrs.timing_vbbb;
}

static uint32_t mpp_videv_mask(unsigned int bitness, unsigned int offset) {
    uint32_t mask = 0;

    if (bitness == 8)       mask = 0xFF000000;
    else if (bitness == 10) mask = 0xFFC00000;
    else if (bitness == 12) mask = 0xFFF00000;
    else if (bitness == 14) mask = 0xFFFF0000;
    else {
                            GO_LOG_VI(LOGGER_ERROR, "VI unsupported cmos pixel bitness!");   
                            return 0x00000000;
    }
    return mask >> offset;
}

static void mpp_videv_set_attrs(VI_DEV_ATTR_S  *stViDevAttr, mpp_vi_init_in * in) {
    memset(stViDevAttr, 0, sizeof(stViDevAttr));

	#if HI_MPP == 1
        stViDevAttr->enIntfMode                                     = in->data_type;
        stViDevAttr->enWorkMode                                     = VI_WORK_MODE_1Multiplex;
        
        stViDevAttr->au32CompMask[0]                                = mpp_videv_mask(in->pixel_bitness, in->dc_zero_bit_offset);
        stViDevAttr->au32CompMask[1]                                = 0;
        
        stViDevAttr->enScanMode                                     = VI_SCAN_PROGRESSIVE;
        
        stViDevAttr->s32AdChnId[0]                                  = -1;
        stViDevAttr->s32AdChnId[1]                                  = -1;
        stViDevAttr->s32AdChnId[2]                                  = -1;
        stViDevAttr->s32AdChnId[3]                                  = -1;
        
        stViDevAttr->enDataSeq                                      = VI_INPUT_DATA_YUYV;       

        mpp_videv_sync(&stViDevAttr->stSynCfg, in);
        
        stViDevAttr->enDataPath                                     = VI_PATH_ISP;
        stViDevAttr->enInputDataType                                = VI_DATA_TYPE_RGB;
        stViDevAttr->bDataRev                                       = HI_FALSE;
	#elif HI_MPP == 2 || HI_MPP == 3
        stViDevAttr->enIntfMode                                     = in->data_type;
        stViDevAttr->enWorkMode                                     = VI_WORK_MODE_1Multiplex;
    
        stViDevAttr->au32CompMask[0]                                = mpp_videv_mask(in->pixel_bitness, in->dc_zero_bit_offset);
        stViDevAttr->au32CompMask[1]                                = 0;
    
        stViDevAttr->enScanMode                                     = VI_SCAN_PROGRESSIVE;
    
        stViDevAttr->s32AdChnId[0]                                  = -1;
        stViDevAttr->s32AdChnId[1]                                  = -1;
        stViDevAttr->s32AdChnId[2]                                  = -1;
        stViDevAttr->s32AdChnId[3]                                  = -1;
    
        stViDevAttr->enDataSeq                                      = VI_INPUT_DATA_YUYV;

        mpp_videv_sync(&stViDevAttr->stSynCfg, in);
    
        stViDevAttr->enDataPath                                     = VI_PATH_ISP;
        stViDevAttr->enInputDataType                                = VI_DATA_TYPE_RGB;
        stViDevAttr->bDataRev                                       = HI_FALSE;

        stViDevAttr->stDevRect.s32X                                 = in->vi_crop_x0;
        stViDevAttr->stDevRect.s32Y                                 = in->vi_crop_y0;
        stViDevAttr->stDevRect.u32Width                             = in->vi_crop_width;
        stViDevAttr->stDevRect.u32Height                            = in->vi_crop_height;
        #if defined(HI3516AV200)
            stViDevAttr->stBasAttr.stSacleAttr.stBasSize.u32Width   = in->vi_crop_width;
            stViDevAttr->stBasAttr.stSacleAttr.stBasSize.u32Height  = in->vi_crop_height;
            stViDevAttr->stBasAttr.stSacleAttr.bCompress            = HI_FALSE;
            stViDevAttr->stBasAttr.stRephaseAttr.enHRephaseMode     = VI_REPHASE_MODE_NONE;
            stViDevAttr->stBasAttr.stRephaseAttr.enVRephaseMode     = VI_REPHASE_MODE_NONE;
        #endif
	#elif HI_MPP == 4
        stViDevAttr->enIntfMode                                     = in->data_type;
        stViDevAttr->enWorkMode                                     = VI_WORK_MODE_1Multiplex;

        stViDevAttr->au32ComponentMask[0]                           = mpp_videv_mask(in->pixel_bitness, in->dc_zero_bit_offset);
        stViDevAttr->au32ComponentMask[1]                           = 0;

        stViDevAttr->enScanMode                                     = VI_SCAN_PROGRESSIVE;
    
        stViDevAttr->as32AdChnId[0]                                 = -1;
        stViDevAttr->as32AdChnId[1]                                 = -1;
        stViDevAttr->as32AdChnId[2]                                 = -1;
        stViDevAttr->as32AdChnId[3]                                 = -1;
    
        stViDevAttr->enDataSeq                                      = VI_DATA_SEQ_YUYV;

        mpp_videv_sync(&stViDevAttr->stSynCfg, in);

        stViDevAttr->enInputDataType                                = VI_DATA_TYPE_RGB;
        stViDevAttr->bDataReverse                                   = HI_FALSE;

        stViDevAttr->stSize.u32Width                                = in->vi_crop_width;
        stViDevAttr->stSize.u32Height                               = in->vi_crop_height;

        stViDevAttr->stBasAttr.stSacleAttr.stBasSize.u32Width       = in->vi_crop_width;
        stViDevAttr->stBasAttr.stSacleAttr.stBasSize.u32Height      = in->vi_crop_height;
        stViDevAttr->stBasAttr.stRephaseAttr.enHRephaseMode         = VI_REPHASE_MODE_NONE;
        stViDevAttr->stBasAttr.stRephaseAttr.enVRephaseMode         = VI_REPHASE_MODE_NONE;
    
        stViDevAttr->stWDRAttr.enWDRMode                            = in->wdr;
        stViDevAttr->stWDRAttr.u32CacheLine                         = in->vi_crop_height;
    
        stViDevAttr->enDataRate                                     = DATA_RATE_X1;
	#endif
}

static void mpp_vichan_set_atts(VI_CHN_ATTR_S *stChnAttr, mpp_vi_init_in * in) {
    memset(stChnAttr, 0, sizeof(stChnAttr));

    #if HI_MPP == 1
        stChnAttr->stCapRect.s32X               = 0;
        stChnAttr->stCapRect.s32Y               = 0;
        stChnAttr->stCapRect.u32Width           = in->width;
        stChnAttr->stCapRect.u32Height          = in->height;

        stChnAttr->stDestSize.u32Width          = in->width;
        stChnAttr->stDestSize.u32Height         = in->height;
        
        stChnAttr->enCapSel                     = VI_CAPSEL_BOTH;
        stChnAttr->enPixFormat                  = PIXEL_FORMAT_YUV_SEMIPLANAR_420; //TODO
        stChnAttr->bChromaResample              = HI_FALSE;
        stChnAttr->s32SrcFrameRate              = in->cmos_fps;
        stChnAttr->s32FrameRate                 = in->fps;
    #elif HI_MPP == 2 || HI_MPP == 3
        stChnAttr->stCapRect.s32X               = 0;
        stChnAttr->stCapRect.s32Y               = 0;
        stChnAttr->stCapRect.u32Width           = in->width;
        stChnAttr->stCapRect.u32Height          = in->height;

        stChnAttr->stDestSize.u32Width          = in->width;
        stChnAttr->stDestSize.u32Height         = in->height;

        stChnAttr->enCapSel                     = VI_CAPSEL_BOTH;
        stChnAttr->enPixFormat                  = PIXEL_FORMAT_YUV_SEMIPLANAR_420; //TODO
        stChnAttr->enCompressMode               = COMPRESS_MODE_NONE; //TODO
        stChnAttr->s32SrcFrameRate              = in->cmos_fps;
        stChnAttr->s32DstFrameRate              = in->fps;
    #elif HI_MPP == 4
        stChnAttr->stSize.u32Width              = in->width;
        stChnAttr->stSize.u32Height             = in->height;

        stChnAttr->enPixelFormat                = PIXEL_FORMAT_YVU_SEMIPLANAR_420; //TODO
        stChnAttr->enDynamicRange               = DYNAMIC_RANGE_SDR8; //TODO
        stChnAttr->enVideoFormat                = VIDEO_FORMAT_LINEAR; //TODO
        stChnAttr->enCompressMode               = COMPRESS_MODE_NONE; //TODO
        stChnAttr->u32Depth                     = 0; //TODO
        stChnAttr->stFrameRate.s32SrcFrameRate  = in->cmos_fps;
        stChnAttr->stFrameRate.s32DstFrameRate  = in->fps;
    #endif

    if (in->mirror == 1) {
        GO_LOG_VI(LOGGER_TRACE, "VI mirror on");
        stChnAttr->bMirror                      = HI_TRUE;
    } else {
        stChnAttr->bMirror                      = HI_FALSE;
    }

    if (in->flip == 1) {
        GO_LOG_VI(LOGGER_TRACE, "VI flip on");
        stChnAttr->bFlip                        = HI_TRUE;
    } else {
        stChnAttr->bFlip                        = HI_FALSE;
    }

}

int mpp_vi_init(error_in *err, mpp_vi_init_in * in) {
    /////DEV

    VI_DEV_ATTR_S  stViDevAttr;
    
    mpp_videv_set_attrs(&stViDevAttr, in);

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VI_SetDevAttr, 0, &stViDevAttr);

    #if HI_MPP == 2 || HI_MPP == 3
        VI_WDR_ATTR_S stWdrAttr;
    
        stWdrAttr.enWDRMode = in->wdr;
        stWdrAttr.bCompress = HI_FALSE;//WTF?
    
        DO_OR_RETURN_ERR_MPP(err, HI_MPI_VI_SetWDRAttr, 0, &stWdrAttr);
    #endif

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VI_EnableDev, 0);

    #if HI_MPP == 4
    /////PIPE

        VI_DEV_BIND_PIPE_S  stDevBindPipe; // = {0};
        memset(&stDevBindPipe, 0, sizeof(stDevBindPipe));

        if (in->wdr != WDR_MODE_NONE) {
            GO_LOG_VI(LOGGER_TRACE, "VI two pipes");
            stDevBindPipe.u32Num    = 2;
            stDevBindPipe.PipeId[0] = 0;
            stDevBindPipe.PipeId[1] = 1; 
        } else {
            GO_LOG_VI(LOGGER_TRACE, "VI one pipe");
            stDevBindPipe.u32Num    = 1;
            stDevBindPipe.PipeId[0] = 0;
        }

        DO_OR_RETURN_ERR_MPP(err, HI_MPI_VI_SetDevBindPipe, 0, &stDevBindPipe);
      
        VI_PIPE_ATTR_S  stPipeAttr;

        //typedef struct hiVI_PIPE_ATTR_S {
        //    VI_PIPE_BYPASS_MODE_E enPipeBypassMode;
        //    HI_BOOL               bYuvSkip;               /* RW;YUV skip enable */
        //    HI_BOOL               bIspBypass;             /* RW;Range:[0, 1];ISP bypass enable */
        //    HI_U32                u32MaxW;                /* RW;Range:[0, 1];Range[VI_PIPE_MIN_WIDTH, VI_PIPE_MAX_WIDTH]; Maximum width */
        //    HI_U32                u32MaxH;                /* RW;Range[VI_PIPE_MIN_HEIGHT, VI_PIPE_MAX_HEIGHT];Maximum height */
        //    PIXEL_FORMAT_E        enPixFmt;               /* RW;Pixel format */
        //    COMPRESS_MODE_E       enCompressMode;         /* RW;Range:[0, 4];Compress mode. */
        //    DATA_BITWIDTH_E       enBitWidth;             /* RW;Range:[0, 4];Bit width */
        //    HI_BOOL               bNrEn;                  /* RW;Range:[0, 1];3DNR enable */
        //    VI_NR_ATTR_S          stNrAttr;               /* RW;Attribute of 3DNR */
        //        typedef struct hiVI_NR_ATTR_S {
        //            PIXEL_FORMAT_E      enPixFmt;                       /* RW;Pixel format of reference frame */
        //            DATA_BITWIDTH_E     enBitWidth;                     /* RW;Bit Width of reference frame */
        //            VI_NR_REF_SOURCE_E  enNrRefSource;                  /* RW;Source of 3DNR reference frame */
        //            COMPRESS_MODE_E     enCompressMode;                 /* RW;Reference frame compress mode */
        //        } VI_NR_ATTR_S;
        //    HI_BOOL               bSharpenEn;             /* RW;Range:[0, 1];Sharpen enable */
        //    FRAME_RATE_CTRL_S     stFrameRate;            /* RW;Frame rate */
        //    HI_BOOL               bDiscardProPic;         /* RW;Range:[0, 1];when professional mode snap, whether to discard long exposure picture in the video pipe. */
        //} VI_PIPE_ATTR_S;

        stPipeAttr.enPipeBypassMode             = VI_PIPE_BYPASS_NONE; //TODO
        stPipeAttr.bYuvSkip                     = HI_FALSE;
        stPipeAttr.bIspBypass                   = HI_FALSE;
        stPipeAttr.u32MaxW                      = in->width;
        stPipeAttr.u32MaxH                      = in->height;
        stPipeAttr.enCompressMode               = COMPRESS_MODE_NONE;

        if (in->pixel_bitness == 12) {
            GO_LOG_VI(LOGGER_TRACE, "VI PIPE bitness 12bit");
            stPipeAttr.enBitWidth               = DATA_BITWIDTH_12;
            stPipeAttr.enPixFmt                 = PIXEL_FORMAT_RGB_BAYER_12BPP;
        } else if (in->pixel_bitness == 10) {
            stPipeAttr.enBitWidth               = DATA_BITWIDTH_10;
            stPipeAttr.enPixFmt                 = PIXEL_FORMAT_RGB_BAYER_10BPP;
        } else {
            ;;; //TODO error
        }

        stPipeAttr.bNrEn                        = HI_TRUE; //HI_FALSE;
        stPipeAttr.stNrAttr.enPixFmt            = PIXEL_FORMAT_YVU_SEMIPLANAR_420,
        stPipeAttr.stNrAttr.enBitWidth          = DATA_BITWIDTH_8;
        stPipeAttr.stNrAttr.enNrRefSource       = VI_NR_REF_FROM_RFR;
        stPipeAttr.stNrAttr.enCompressMode      = COMPRESS_MODE_NONE;

        stPipeAttr.bSharpenEn                   = HI_FALSE;
        stPipeAttr.stFrameRate.s32SrcFrameRate  = in->cmos_fps;
        stPipeAttr.stFrameRate.s32DstFrameRate  = in->cmos_fps;
        stPipeAttr.bDiscardProPic               = HI_FALSE;

        DO_OR_RETURN_ERR_MPP(err, HI_MPI_VI_CreatePipe, 0, &stPipeAttr);
        DO_OR_RETURN_ERR_MPP(err, HI_MPI_VI_StartPipe, 0);

        if (in->wdr != WDR_MODE_NONE) {
            DO_OR_RETURN_ERR_MPP(err, HI_MPI_VI_CreatePipe, 1, &stPipeAttr);
            DO_OR_RETURN_ERR_MPP(err, HI_MPI_VI_StartPipe, 1);
        }

    #endif

    /////CHANNEL

    VI_CHN_ATTR_S stChnAttr;

    mpp_vichan_set_atts(&stChnAttr, in);

    #if HI_MPP <= 3
        DO_OR_RETURN_ERR_MPP(err, HI_MPI_VI_SetChnAttr, 0, &stChnAttr);
    #elif HI_MPP == 4
        DO_OR_RETURN_ERR_MPP(err, HI_MPI_VI_SetChnAttr, 0, 0, &stChnAttr);
    #endif

    #if HI_MPP == 2 || HI_MPP == 3
        if (in->ldc == 1) {
            VI_LDC_ATTR_S stLDCAttr;

            stLDCAttr.bEnable                   = HI_TRUE;
            stLDCAttr.stAttr.enViewType         = LDC_VIEW_TYPE_ALL;         //LDC_VIEW_TYPE_CROP;
            stLDCAttr.stAttr.s32CenterXOffset   = in->ldc_offset_x;
            stLDCAttr.stAttr.s32CenterYOffset   = in->ldc_offset_y;
            stLDCAttr.stAttr.s32Ratio           = in->ldc_k;

            #if HI_MPP == 3
                stLDCAttr.stAttr.s32MinRatio = 0; //should be 0 for LDC_VIEW_TYPE_ALL
            #endif

            DO_OR_RETURN_ERR_MPP(err, HI_MPI_VI_SetLDCAttr, 0, &stLDCAttr);
        }
    #endif

    #if HI_MPP <= 3
        DO_OR_RETURN_ERR_MPP(err, HI_MPI_VI_EnableChn, 0);
    #elif HI_MPP == 4
        DO_OR_RETURN_ERR_MPP(err, HI_MPI_VI_EnableChn, 0, 0);
    #endif

    return ERR_NONE;
}
