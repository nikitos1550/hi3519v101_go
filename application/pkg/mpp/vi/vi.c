#include "vi.h"

#if defined(HI3516CV100)
int mpp_vi_init(error_in *err, mpp_vi_init_in * in) {

    VI_DEV_ATTR_S    stViDevAttr;

    memset(&stViDevAttr, 0, sizeof(stViDevAttr));

    memcpy(&stViDevAttr, in->videv, sizeof(stViDevAttr));

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VI_SetDevAttr, 0, &stViDevAttr);

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VI_EnableDev, 0);

    VI_CHN_ATTR_S stChnAttr;
    //ROTATE_E enRotate = ROTATE_NONE;

    //memcpy(&stChnAttr.stCapRect, pstCapRect, sizeof(RECT_S));
    stChnAttr.stCapRect.s32X        = in->x0;
    stChnAttr.stCapRect.s32Y        = in->y0;
    stChnAttr.stCapRect.u32Width    = in->width;
    stChnAttr.stCapRect.u32Height   = in->height;


    stChnAttr.enCapSel = VI_CAPSEL_BOTH;
    stChnAttr.stDestSize.u32Width   = in->width; 
    stChnAttr.stDestSize.u32Height  = in->height;
    stChnAttr.enPixFormat           = PIXEL_FORMAT_YUV_SEMIPLANAR_420;   // sp420 or sp422

    stChnAttr.bMirror = HI_FALSE;
    stChnAttr.bFlip = HI_FALSE;

    stChnAttr.bChromaResample   = HI_FALSE;
    stChnAttr.s32SrcFrameRate   = 30;
    stChnAttr.s32FrameRate      = 30;

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VI_SetChnAttr, 0, &stChnAttr);
    
    //if(ROTATE_NONE != enRotate)
    //{
    //    *error_code = HI_MPI_VI_SetRotate(ViChn, enRotate);
    //  if (*error_code != HI_SUCCESS) return ERR_
    //}
    
    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VI_EnableChn, 0);

    return ERR_NONE;
}

#endif // defined(HI3516CV100)

//#if defined(HI3516AV100)
//int mpp_vi_init(error_in *err, mpp_vi_init_in * in) {
//    //unsigned int mpp_error_code = 0;
//
//    VI_DEV_ATTR_S  stViDevAttr;
//
//    memset(&stViDevAttr, 0, sizeof(stViDevAttr));
//    memcpy(&stViDevAttr, in->videv, sizeof(stViDevAttr));
//
//    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VI_SetDevAttr, 0, &stViDevAttr);
//
//    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VI_EnableDev, 0);
//
//    RECT_S stCapRect;
//
//    stCapRect.s32X = in->x0;
//    stCapRect.s32Y = in->y0;
//    stCapRect.u32Width  = in->width;
//    stCapRect.u32Height = in->height;
//
//    VI_CHN_ATTR_S stChnAttr;
//
//    memcpy(&stChnAttr.stCapRect, &stCapRect, sizeof(RECT_S));
//    stChnAttr.enCapSel              = VI_CAPSEL_BOTH;
//    stChnAttr.stDestSize.u32Width   = in->width;
//    stChnAttr.stDestSize.u32Height  = in->height;
//    stChnAttr.enPixFormat           = PIXEL_FORMAT_YUV_SEMIPLANAR_420;   // sp420 or sp422
//
//    if (in->mirror == 1) {
//        stChnAttr.bMirror = HI_TRUE;
//    } else {
//        stChnAttr.bMirror = HI_FALSE;
//    }
//
//    if (in->flip == 1) {
//        stChnAttr.bFlip = HI_TRUE;
//    } else {
//        stChnAttr.bFlip = HI_FALSE;
//    }
//
//    stChnAttr.s32SrcFrameRate = in->cmos_fps;
//    stChnAttr.s32DstFrameRate = in->fps;
//
//    stChnAttr.enCompressMode = COMPRESS_MODE_NONE;
//    //TODO check family mpp datasheet
//    //if (in->ldc == 1) {
//    //    stChnAttr.enCompressMode        = COMPRESS_MODE_NONE;
//    //} else {
//    //    stChnAttr.enCompressMode        = COMPRESS_MODE_SEG;
//    //}
//
//    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VI_SetChnAttr, 0, &stChnAttr);
//
//    if (in->ldc == 1) {
//        VI_LDC_ATTR_S stLDCAttr;
//        
//        stLDCAttr.bEnable = HI_TRUE;
//        stLDCAttr.stAttr.enViewType =   LDC_VIEW_TYPE_ALL;
//                                        //LDC_VIEW_TYPE_CROP;
//        stLDCAttr.stAttr.s32CenterXOffset = in->ldc_offset_x;
//        stLDCAttr.stAttr.s32CenterYOffset = in->ldc_offset_y;
//        stLDCAttr.stAttr.s32Ratio = in->ldc_k;
//
//        DO_OR_RETURN_ERR_MPP(err, HI_MPI_VI_SetLDCAttr, 0, &stLDCAttr);
//
//        //s32Ret = HI_MPI_VI_GetLDCAttr (ViChn, &stLDCAttr);
//        //if (HI_SUCCESS != s32ret)
//        //{
//        //printf("Get vi LDC attr err:0x%x\n", s32ret);
//        //return s32Ret;
//        //}
//    }
//
//    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VI_EnableChn, 0)
//
//    return ERR_NONE;
//}
//#endif // defined(HI3516AV100)

#if defined(HI3516AV200) \
    || defined(HI3516CV300) \
    || defined(HI3516AV100)
int mpp_vi_init(error_in *err, mpp_vi_init_in * in) {
    unsigned int mpp_error_code = 0;

    VI_DEV_ATTR_S  stViDevAttr;

    memset(&stViDevAttr, 0, sizeof(stViDevAttr));

    memcpy(&stViDevAttr, in->videv, sizeof(stViDevAttr));

    //ATTENTION, videv struct should be constructed here!!!!
    //Now this info should be in cmos source file
    //stViDevAttr.stDevRect.s32X                              = 0;
    //stViDevAttr.stDevRect.s32Y                              = 0;
    //stViDevAttr.stDevRect.u32Width                          = in->cmos_width; 
    //stViDevAttr.stDevRect.u32Height                         = in->cmos_height;

    #if defined(HI3516AV200)
        //For Hi3519 V101, Dev0 does not support scaling and phase adjustment of the Bayer
        //domain. Therefore, for Dev0, the width and height of stBasAttr must be the same as
        //those of stDevRect. Dev1 supports scaling and phase adjustment of the Bayer domain.
        //Note that the width and height can be scaled down only by the multiple of 1, 1/2, or 1/3.
        //Otherwise, calling the interface fails.
    
        stViDevAttr.stBasAttr.stSacleAttr.stBasSize.u32Width    = in->cmos_width; 
        stViDevAttr.stBasAttr.stSacleAttr.stBasSize.u32Height   = in->cmos_height;
        stViDevAttr.stBasAttr.stSacleAttr.bCompress             = HI_FALSE;
    #endif

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VI_SetDevAttr, 0, &stViDevAttr);
  
    //TODO when we use 274 on 19v101 we set wdr in isp, but don`t set here in vi, but I am not sure
    VI_WDR_ATTR_S stWdrAttr;

    stWdrAttr.enWDRMode = in->wdr;
    stWdrAttr.bCompress = HI_FALSE;//TODO WTF?
    
    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VI_SetWDRAttr, 0, &stWdrAttr);

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VI_EnableDev, 0);

    RECT_S stCapRect;

    //Vi channel capture region coordinates are relative to Vi device
    stCapRect.s32X          = in->x0;
    stCapRect.s32Y          = in->y0;
    stCapRect.u32Width      = in->width;
    stCapRect.u32Height     = in->height;

    VI_CHN_ATTR_S stChnAttr;

    memcpy(&stChnAttr.stCapRect, &stCapRect, sizeof(RECT_S));

    stChnAttr.enCapSel              = VI_CAPSEL_BOTH;
    stChnAttr.stDestSize.u32Width   = in->width;
    stChnAttr.stDestSize.u32Height  = in->height;
    stChnAttr.enPixFormat           = PIXEL_FORMAT_YUV_SEMIPLANAR_420;   // sp420 or sp422
    
    if (in->mirror == 1) {
        GO_LOG_VI(LOGGER_TRACE, "VI mirror on");
        stChnAttr.bMirror               = HI_TRUE;
    } else {
        stChnAttr.bMirror               = HI_FALSE;
    }

    if (in->flip == 1) {
        GO_LOG_VI(LOGGER_TRACE, "VI flip on");
        stChnAttr.bFlip                 = HI_TRUE;
    } else {
        stChnAttr.bFlip                 = HI_FALSE;
    }

    stChnAttr.s32SrcFrameRate       = in->cmos_fps;
    stChnAttr.s32DstFrameRate       = in->fps;

    //When Hi3519 V100/Hi3519 V101 enables the LDC function, the compression mode in
    //the physical channel attributes must be set to non-compression mode.
    //COMPRESS_MODE_NONE        = 0x0,  //no compress 
    //COMPRESS_MODE_SEG         = 0x1,  //compress unit is 256 bytes as a segment, default seg mode
    //COMPRESS_MODE_SEG128      = 0x2,  //compress unit is 128 bytes as a segment
    //COMPRESS_MODE_LINE        = 0x3,  //compress unit is the whole line
    //COMPRESS_MODE_FRAME       = 0x4,  //compress unit is the whole frame

    if (in->ldc == 1) {
        stChnAttr.enCompressMode        = COMPRESS_MODE_NONE;
    } else {
        stChnAttr.enCompressMode        = COMPRESS_MODE_SEG;//TODO
        //stChnAttr.enCompressMode        = COMPRESS_MODE_NONE;
    }

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VI_SetChnAttr, 0, &stChnAttr);

    if (in->ldc == 1) {
        VI_LDC_ATTR_S stLDCAttr;

        stLDCAttr.bEnable = HI_TRUE;
        stLDCAttr.stAttr.enViewType =   LDC_VIEW_TYPE_ALL;
                                        //LDC_VIEW_TYPE_CROP;
        stLDCAttr.stAttr.s32CenterXOffset = in->ldc_offset_x;
        stLDCAttr.stAttr.s32CenterYOffset = in->ldc_offset_y;
        stLDCAttr.stAttr.s32Ratio = in->ldc_k;

        #if HI_MPP == 3
            stLDCAttr.stAttr.s32MinRatio = 0; //should be 0 for LDC_VIEW_TYPE_ALL
        #endif

        DO_OR_RETURN_ERR_MPP(err, HI_MPI_VI_SetLDCAttr, 0, &stLDCAttr);

        //Obtain LDC attributes.
        //s32Ret = HI_MPI_VI_GetLDCAttr (0, &stLDCAttr);
        //if (HI_SUCCESS != s32Ret) {
        //    printf("Get vi LDC attr err:0x%x\n", s32Ret);
        //    return HI_FAILURE;
        //}
    }

    //hi3516av200
    //// when VI-VPSS online, VI Rotate is not support, HI_MPI_VI_SetRotate will failed
    //if (ROTATE_NONE != enRotate && !SAMPLE_COMM_IsViVpssOnline()) {
    //    s32Ret = HI_MPI_VI_SetRotate(ViChn, enRotate);
    //    if (s32Ret != HI_SUCCESS)
    //    {
    //        SAMPLE_PRT("HI_MPI_VI_SetRotate failed with %#x!\n", s32Ret);
    //        return HI_FAILURE;
    //    }
    //}


    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VI_EnableChn, 0);

    return ERR_NONE;
}
#endif

//#if defined(HI3516CV300)
//int mpp_vi_init(error_in *err, mpp_vi_init_in * in) {
//    unsigned int mpp_error_code = 0;
//
//    ISP_WDR_MODE_S stWdrMode;
//    VI_DEV_ATTR_S  stViDevAttr;
//    
//    memset(&stViDevAttr,0,sizeof(stViDevAttr));
//
//    memcpy(&stViDevAttr, in->videv, sizeof(stViDevAttr));
//
//    //stViDevAttr.stDevRect.s32X = 0;
//    //stViDevAttr.stDevRect.s32Y = 0;
//    //stViDevAttr.stDevRect.u32Width  = in->cmos_width;
//    //stViDevAttr.stDevRect.u32Height = in->cmos_height;
//
//    mpp_error_code = HI_MPI_VI_SetDevAttr(0, &stViDevAttr);
//    if (mpp_error_code != HI_SUCCESS) {
//        RETURN_ERR_MPP(ERR_F_HI_MPI_VI_SetDevAttr, mpp_error_code);
//    }
//
//    VI_WDR_ATTR_S stWdrAttr;
//
//    // WDR_MODE_NONE or WDR_MODE_2To1_LINE TODO
//    stWdrAttr.enWDRMode = WDR_MODE_NONE;
//    stWdrAttr.bCompress = HI_FALSE;
//
//    mpp_error_code = HI_MPI_VI_SetWDRAttr(0, &stWdrAttr);
//    if (mpp_error_code != HI_SUCCESS) {
//        RETURN_ERR_MPP(ERR_F_HI_MPI_VI_SetWDRAttr, mpp_error_code);
//    }
//    
//    mpp_error_code = HI_MPI_VI_EnableDev(0);
//    if (mpp_error_code != HI_SUCCESS) {
//        RETURN_ERR_MPP(ERR_F_HI_MPI_VI_EnableDev, mpp_error_code);
//    }
//
//    RECT_S stCapRect;
//
//    stCapRect.s32X = in->x0;
//    stCapRect.s32Y = in->y0;
//    stCapRect.u32Width  = in->width;
//    stCapRect.u32Height = in->height;
//
//
//    VI_CHN_ATTR_S stChnAttr;
//
//    memcpy(&stChnAttr.stCapRect, &stCapRect, sizeof(RECT_S));
//
//    stChnAttr.enCapSel = VI_CAPSEL_BOTH;
//    stChnAttr.stDestSize.u32Width = in->width;
//    stChnAttr.stDestSize.u32Height =  in->height;
//    stChnAttr.enPixFormat = PIXEL_FORMAT_YUV_SEMIPLANAR_420;   // sp420 or sp422
//
//    if (in->mirror == 1) {
//        GO_LOG_VI(LOGGER_TRACE, "VI mirror on");
//        stChnAttr.bMirror               = HI_TRUE;
//    } else {
//        stChnAttr.bMirror               = HI_FALSE;
//    }
//
//    if (in->flip == 1) {
//        GO_LOG_VI(LOGGER_TRACE, "VI flip on");
//        stChnAttr.bFlip                 = HI_TRUE;
//    } else {
//        stChnAttr.bFlip                 = HI_FALSE;
//    }
//
//
//    stChnAttr.s32SrcFrameRate = in->cmos_fps;
//    stChnAttr.s32DstFrameRate = in->fps;
//    //stChnAttr.enCompressMode = COMPRESS_MODE_SEG;
//    stChnAttr.enCompressMode = COMPRESS_MODE_NONE;
//
//    mpp_error_code = HI_MPI_VI_SetChnAttr(0, &stChnAttr);
//    if (mpp_error_code != HI_SUCCESS) {
//        RETURN_ERR_MPP(ERR_F_HI_MPI_VI_SetChnAttr, mpp_error_code);
//    }
//
//    mpp_error_code = HI_MPI_VI_EnableChn(0);
//    if (mpp_error_code != HI_SUCCESS) {
//        RETURN_ERR_MPP(ERR_F_HI_MPI_VI_EnableChn, mpp_error_code);
//    }
//
//    return ERR_NONE;
//}
//#endif // defined(HI3516CV300)


#if defined(HI3516CV500) \
    || defined(HI3516EV200)
static VI_PIPE_ATTR_S PIPE_ATTR_2592x1944_RAW12_420_3DNR_RFR =
{
    VI_PIPE_BYPASS_NONE, HI_FALSE, HI_FALSE,
    2592, 1944,
    PIXEL_FORMAT_RGB_BAYER_12BPP,  
    COMPRESS_MODE_LINE,
    DATA_BITWIDTH_12,
    HI_FALSE,
    {
        PIXEL_FORMAT_YVU_SEMIPLANAR_420,
        DATA_BITWIDTH_8,
        VI_NR_REF_FROM_RFR,
        COMPRESS_MODE_NONE
    },
    HI_FALSE,                        
    {-1, -1}                
};                                           

VI_PIPE_ATTR_S PIPE_ATTR_1920x1080_RAW12_420_3DNR_RFR =
{
    VI_PIPE_BYPASS_NONE, HI_FALSE, HI_FALSE,
    1920, 1080,
    PIXEL_FORMAT_RGB_BAYER_12BPP,
    COMPRESS_MODE_NONE,
    DATA_BITWIDTH_12,
    HI_FALSE,
    {
        PIXEL_FORMAT_YVU_SEMIPLANAR_420,
        DATA_BITWIDTH_8,   
        VI_NR_REF_FROM_RFR,
        COMPRESS_MODE_NONE
    },
    HI_FALSE,
    { -1, -1}
};

static VI_CHN_ATTR_S CHN_ATTR_2592x1944_420_SDR8_LINEAR =
{
    {2592, 1944},
    PIXEL_FORMAT_YVU_SEMIPLANAR_420,  
    DYNAMIC_RANGE_SDR8,
    VIDEO_FORMAT_LINEAR,
    COMPRESS_MODE_NONE,
    0,      0,
    0,
    {-1, -1}
};

VI_CHN_ATTR_S CHN_ATTR_1920x1080_420_SDR8_LINEAR =
{
    {1920, 1080},
    PIXEL_FORMAT_YVU_SEMIPLANAR_420,
    DYNAMIC_RANGE_SDR8, 
    VIDEO_FORMAT_LINEAR,
    COMPRESS_MODE_NONE,
    0,      0,
    0,
    { -1, -1}
};


int mpp_vi_init(error_in *err, mpp_vi_init_in * in) { 
    //*error_code = 0;
   
  	//VI_StartDev
    VI_DEV_ATTR_S       stViDevAttr;
    hi_memcpy(&stViDevAttr, sizeof(VI_DEV_ATTR_S), in->videv, sizeof(VI_DEV_ATTR_S));

	stViDevAttr.stWDRAttr.enWDRMode = WDR_MODE_NONE;
    //stViDevAttr.enDataRate = DATA_RATE_X2; //???????????????

	DO_OR_RETURN_ERR_MPP(err, HI_MPI_VI_SetDevAttr, 0, &stViDevAttr);

	DO_OR_RETURN_ERR_MPP(err, HI_MPI_VI_EnableDev, 0);

    VI_DEV_BIND_PIPE_S  stDevBindPipe = {0};
    
   	stDevBindPipe.u32Num = 1;

	DO_OR_RETURN_ERR_MPP(err, HI_MPI_VI_SetDevBindPipe, 0, &stDevBindPipe);
      
    VI_PIPE_ATTR_S  stPipeAttr;

    #if defined(HI3516CV500)
    hi_memcpy(&stPipeAttr, sizeof(VI_PIPE_ATTR_S), &PIPE_ATTR_2592x1944_RAW12_420_3DNR_RFR, sizeof(VI_PIPE_ATTR_S));
    #elif defined(HI3516EV200)
    hi_memcpy(&stPipeAttr, sizeof(VI_PIPE_ATTR_S), &PIPE_ATTR_1920x1080_RAW12_420_3DNR_RFR, sizeof(VI_PIPE_ATTR_S));
    #endif

	DO_OR_RETURN_ERR_MPP(err, HI_MPI_VI_CreatePipe, 0, &stPipeAttr);

	DO_OR_RETURN_ERR_MPP(err, HI_MPI_VI_StartPipe, 0);
    
    VI_CHN_ATTR_S       stChnAttr;

    #if defined(HI3516CV500)
    hi_memcpy(&stChnAttr, sizeof(VI_CHN_ATTR_S), &CHN_ATTR_2592x1944_420_SDR8_LINEAR, sizeof(VI_CHN_ATTR_S));
    #elif defined(HI3516EV200)
    hi_memcpy(&stChnAttr, sizeof(VI_CHN_ATTR_S), &CHN_ATTR_1920x1080_420_SDR8_LINEAR, sizeof(VI_CHN_ATTR_S));
    #endif

    //stChnAttr.enDynamicRange = DYNAMIC_RANGE_SDR8;
    //stChnAttr.enVideoFormat  = VIDEO_FORMAT_LINEAR;
    //stChnAttr.enPixelFormat  = PIXEL_FORMAT_YVU_SEMIPLANAR_420;
    //stChnAttr.enCompressMode = COMPRESS_MODE_SEG;

	DO_OR_RETURN_ERR_MPP(err, HI_MPI_VI_SetChnAttr, 0, 0, &stChnAttr);

	DO_OR_RETURN_ERR_MPP(err, HI_MPI_VI_EnableChn, 0, 0);

    return ERR_NONE;
}

#endif // defined(HI3516CV500)
