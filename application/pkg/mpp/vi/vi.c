#include "vi.h"

#if defined(HI3516CV100)
//imx122 DC 12bitÊäÈë
VI_DEV_ATTR_S DEV_ATTR_IMX122_DC_1080P =
{
    //½Ó¿ÚÄ£Ê½
    VI_MODE_DIGITAL_CAMERA,
    //1¡¢2¡¢4Â·¹¤×÷Ä£Ê½
    VI_WORK_MODE_1Multiplex,
    // r_mask    g_mask    b_mask
    {0xFFF00000,    0x0},
    //ÖðÐÐor¸ôÐÐÊäÈë
    VI_SCAN_PROGRESSIVE,
    //AdChnId
    {-1, -1, -1, -1},
    //enDataSeq, ½öÖ§³ÖYUV¸ñÊ½
    VI_INPUT_DATA_YUYV,

    //Í¬²½ÐÅÏ¢£¬¶ÔÓ¦regÊÖ²áµÄÈçÏÂÅäÖÃ, --bt1120Ê±ÐòÎÞÐ§
    {
    //port_vsync   port_vsync_neg     port_hsync        port_hsync_neg      
    VI_VSYNC_PULSE, VI_VSYNC_NEG_HIGH, VI_HSYNC_VALID_SINGNAL,VI_HSYNC_NEG_HIGH,VI_VSYNC_NORM_PULSE,VI_VSYNC_VALID_NEG_HIGH,
    
    //timingÐÅÏ¢£¬¶ÔÓ¦regÊÖ²áµÄÈçÏÂÅäÖÃ
    //hsync_hfb    hsync_act    hsync_hhb
    {0,            1920,        0,
    //vsync0_vhb vsync0_act vsync0_hhb
     0,            1080,        0,
    //vsync1_vhb vsync1_act vsync1_hhb
     0,            0,            0}
    },
    //Ê¹ÓÃÄÚ²¿ISP
    VI_PATH_ISP,
    //ÊäÈëÊý¾ÝÀàÐÍ
    VI_DATA_TYPE_RGB
};


int mpp1_vi_init(unsigned int *error_code) {
    *error_code = 0;

    VI_DEV ViDev = 0;

    VI_DEV_ATTR_S    stViDevAttr;
    memset(&stViDevAttr,0,sizeof(stViDevAttr));

    //case SONY_IMX122_DC_1080P_30FPS:                                        
    memcpy(&stViDevAttr,&DEV_ATTR_IMX122_DC_1080P,sizeof(stViDevAttr));    

    *error_code = HI_MPI_VI_SetDevAttr(ViDev, &stViDevAttr);
    if (*error_code != HI_SUCCESS) return ERR_HI_MPI_VI_SetDevAttr;

    *error_code = HI_MPI_VI_EnableDev(ViDev);
    if (*error_code != HI_SUCCESS) return ERR_HI_MPI_VI_EnableDev;

    VI_CHN ViChn = 0;

    VI_CHN_ATTR_S stChnAttr;
    ROTATE_E enRotate = ROTATE_NONE;

    // step  5: config & start vicap dev
    //memcpy(&stChnAttr.stCapRect, pstCapRect, sizeof(RECT_S));
    stChnAttr.stCapRect.s32X = 0;
    stChnAttr.stCapRect.s32Y = 0;
    stChnAttr.stCapRect.u32Width = 1920;
    stChnAttr.stCapRect.u32Height = 1080;


    stChnAttr.enCapSel = VI_CAPSEL_BOTH;
    // to show scale. this is a sample only, we want to show dist_size = D1 only 
    stChnAttr.stDestSize.u32Width = 1920; //pstTarSize->u32Width;
    stChnAttr.stDestSize.u32Height = 1080; //pstTarSize->u32Height;
    stChnAttr.enPixFormat = PIXEL_FORMAT_YUV_SEMIPLANAR_420;   // sp420 or sp422

    stChnAttr.bMirror = HI_FALSE;
    stChnAttr.bFlip = HI_FALSE;

    stChnAttr.bChromaResample = HI_FALSE;
    stChnAttr.s32SrcFrameRate = 30;
    stChnAttr.s32FrameRate = 30;

    *error_code = HI_MPI_VI_SetChnAttr(ViChn, &stChnAttr);
    if (*error_code != HI_SUCCESS) return ERR_HI_MPI_VI_SetChnAttr;
    
    //if(ROTATE_NONE != enRotate)
    //{
    //    *error_code = HI_MPI_VI_SetRotate(ViChn, enRotate);
    //  if (*error_code != HI_SUCCESS) return ERR_
    //}
    
    *error_code = HI_MPI_VI_EnableChn(ViChn);
    if (*error_code != HI_SUCCESS) return ERR_HI_MPI_VI_EnableChn;

    return ERR_NONE;
}

#endif // defined(HI3516CV100)

#if defined(HI3516AV100)

int mpp_vi_init(int *error_code, mpp_vi_init_in *in) {
    unsigned int mpp_error_code = 0;

    VI_DEV_ATTR_S  stViDevAttr;

    memset(&stViDevAttr, 0, sizeof(stViDevAttr));
    memcpy(&stViDevAttr, in->videv, sizeof(stViDevAttr));
    stViDevAttr.stDevRect.s32X = in->x0;
    stViDevAttr.stDevRect.s32Y = in->y0;
    stViDevAttr.stDevRect.u32Width  = in->width;
    stViDevAttr.stDevRect.u32Height = in->height;

    mpp_error_code = HI_MPI_VI_SetDevAttr(0, &stViDevAttr);
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_VI_SetDevAttr, mpp_error_code);
    }

    mpp_error_code = HI_MPI_VI_EnableDev(0);
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_VI_EnableDev, mpp_error_code);
    }

    RECT_S stCapRect;

    stCapRect.s32X = in->x0;
    stCapRect.s32Y = in->y0;
    stCapRect.u32Width  = in->width;
    stCapRect.u32Height = in->height;

    VI_CHN_ATTR_S stChnAttr;

    memcpy(&stChnAttr.stCapRect, &stCapRect, sizeof(RECT_S));
    stChnAttr.enCapSel              = VI_CAPSEL_BOTH;
    stChnAttr.stDestSize.u32Width   = in->width;
    stChnAttr.stDestSize.u32Height  = in->height;
    stChnAttr.enPixFormat           = PIXEL_FORMAT_YUV_SEMIPLANAR_420;   // sp420 or sp422

    if (in->mirror == 1) {
        stChnAttr.bMirror = HI_TRUE;
    } else {
        stChnAttr.bMirror = HI_FALSE;
    }

    if (in->flip == 1) {
        stChnAttr.bFlip = HI_TRUE;
    } else {
        stChnAttr.bFlip = HI_FALSE;
    }

    stChnAttr.s32SrcFrameRate = in->fps;
    stChnAttr.s32DstFrameRate = in->fps;

    stChnAttr.enCompressMode = COMPRESS_MODE_NONE;
    //TODO check family mpp datasheet
    //if (in->ldc == 1) {
    //    stChnAttr.enCompressMode        = COMPRESS_MODE_NONE;
    //} else {
    //    stChnAttr.enCompressMode        = COMPRESS_MODE_SEG;
    //}

    mpp_error_code = HI_MPI_VI_SetChnAttr(0, &stChnAttr);
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_VI_SetChnAttr, mpp_error_code);
    }

    if (in.ldc == 1) {
        VI_LDC_ATTR_S stLDCAttr;
        
        stLDCAttr.bEnable = HI_TRUE;
        stLDCAttr.stAttr.enViewType =   LDC_VIEW_TYPE_ALL;
                                        //LDC_VIEW_TYPE_CROP;
        stLDCAttr.stAttr.s32CenterXOffset = ldc_offset_x;
        stLDCAttr.stAttr.s32CenterYOffset = ldc_offset_y;
        stLDCAttr.stAttr.s32Ratio = ldc_k;

        mpp_error_code = HI_MPI_VI_SetLDCAttr(0, &stLDCAttr);
        if (mpp_error_code != HI_SUCCESS) {
            RETURN_ERR_MPP(ERR_F_HI_MPI_VI_SetLDCAttr, mpp_error_code);
        }

        //s32Ret = HI_MPI_VI_GetLDCAttr (ViChn, &stLDCAttr);
        //if (HI_SUCCESS != s32ret)
        //{
        //printf("Get vi LDC attr err:0x%x\n", s32ret);
        //return s32Ret;
        //}
    }

    mpp_error_code = HI_MPI_VI_EnableChn(0);
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_VI_EnableChn, mpp_error_code);
    }

    return ERR_NONE;
}
#endif // defined(HI3516AV100)

#if defined(HI3516AV200)
int mpp_vi_init(error_in *err, mpp_vi_init_in * in) {
    unsigned int mpp_error_code = 0;

    VI_DEV_ATTR_S  stViDevAttr;

    memset(&stViDevAttr, 0, sizeof(stViDevAttr));

    memcpy(&stViDevAttr, in->videv, sizeof(stViDevAttr));

    stViDevAttr.stDevRect.s32X                              = 0;
    stViDevAttr.stDevRect.s32Y                              = 0;
    stViDevAttr.stDevRect.u32Width                          = in->cmos_width; 
    stViDevAttr.stDevRect.u32Height                         = in->cmos_height;

    //For Hi3519 V101, Dev0 does not support scaling and phase adjustment of the Bayer
    //domain. Therefore, for Dev0, the width and height of stBasAttr must be the same as
    //those of stDevRect. Dev1 supports scaling and phase adjustment of the Bayer domain.
    //Note that the width and height can be scaled down only by the multiple of 1, 1/2, or 1/3.
    //Otherwise, calling the interface fails.
    
    stViDevAttr.stBasAttr.stSacleAttr.stBasSize.u32Width    = in->cmos_width; 
    stViDevAttr.stBasAttr.stSacleAttr.stBasSize.u32Height   = in->cmos_height;
    stViDevAttr.stBasAttr.stSacleAttr.bCompress             = HI_FALSE;

    mpp_error_code = HI_MPI_VI_SetDevAttr(0, &stViDevAttr);
    if (mpp_error_code != HI_SUCCESS) {
        //GO_LOG_VI(LOGGER_ERROR, "HI_MPI_VI_SetDevAttr");
        //err->mpp = mpp_error_code;
        //return ERR_MPP;
        RETURN_ERR_MPP(ERR_F_HI_MPI_VI_SetDevAttr, mpp_error_code);
    }

    mpp_error_code = HI_MPI_VI_EnableDev(0);
    if (mpp_error_code != HI_SUCCESS) {
        //GO_LOG_VI(LOGGER_ERROR, "HI_MPI_VI_EnableDev")
        //err->mpp = mpp_error_code;
        //return ERR_MPP;
        RETURN_ERR_MPP(ERR_F_HI_MPI_VI_EnableDev, mpp_error_code);
    }

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
        stChnAttr.enCompressMode        = COMPRESS_MODE_SEG;
    }

    mpp_error_code = HI_MPI_VI_SetChnAttr(0, &stChnAttr);
    if (mpp_error_code != HI_SUCCESS) {
        //GO_LOG_VI(LOGGER_ERROR, "HI_MPI_VI_SetChnAttr")   
        //err->mpp = mpp_error_code;
        //return ERR_MPP;
        RETURN_ERR_MPP(ERR_F_HI_MPI_VI_SetChnAttr, mpp_error_code);
    }

    if (in->ldc == 1) {
        VI_LDC_ATTR_S stLDCAttr;

        stLDCAttr.bEnable = HI_TRUE;
        stLDCAttr.stAttr.enViewType =   LDC_VIEW_TYPE_ALL;
                                        //LDC_VIEW_TYPE_CROP;
        stLDCAttr.stAttr.s32CenterXOffset = in->ldc_offset_x;
        stLDCAttr.stAttr.s32CenterYOffset = in->ldc_offset_y;
        stLDCAttr.stAttr.s32Ratio = in->ldc_k;
        stLDCAttr.stAttr.s32MinRatio = 0; //should be 0 for LDC_VIEW_TYPE_ALL
    
        mpp_error_code = HI_MPI_VI_SetLDCAttr(0, &stLDCAttr);
        if (mpp_error_code != HI_SUCCESS) {
            //GO_LOG_VI(LOGGER_ERROR, "HI_MPI_VI_SetLDCAttr")
            //err->mpp = mpp_error_code;
            //return ERR_MPP;
            RETURN_ERR_MPP(ERR_F_HI_MPI_VI_SetLDCAttr, mpp_error_code);
        }
        //Obtain LDC attributes.
        //s32Ret = HI_MPI_VI_GetLDCAttr (0, &stLDCAttr);
        //if (HI_SUCCESS != s32Ret) {
        //    printf("Get vi LDC attr err:0x%x\n", s32Ret);
        //    return HI_FAILURE;
        //}
    }

    mpp_error_code = HI_MPI_VI_EnableChn(0);
    if (mpp_error_code != HI_SUCCESS) {
        //GO_LOG_VI(LOGGER_ERROR, "HI_MPI_VI_EnableChn")
        //err->mpp = mpp_error_code;
        //return ERR_MPP;
        RETURN_ERR_MPP(ERR_F_HI_MPI_VI_EnableChn, mpp_error_code);
    }

    return ERR_NONE;
}
#endif // defined(HI3516AV200)

#if defined(HI3516CV300)
int mpp_vi_init(error_in *err, mpp_vi_init_in * in) {
    unsigned int mpp_error_code = 0;

    ISP_WDR_MODE_S stWdrMode;
    VI_DEV_ATTR_S  stViDevAttr;
    
    memset(&stViDevAttr,0,sizeof(stViDevAttr));

    memcpy(&stViDevAttr, in->videv, sizeof(stViDevAttr));

    //stViDevAttr.stDevRect.s32X = 0;
    //stViDevAttr.stDevRect.s32Y = 0;
    //stViDevAttr.stDevRect.u32Width  = in->cmos_width;
    //stViDevAttr.stDevRect.u32Height = in->cmos_height;

    mpp_error_code = HI_MPI_VI_SetDevAttr(0, &stViDevAttr);
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_VI_SetDevAttr, mpp_error_code);
    }

    VI_WDR_ATTR_S stWdrAttr;

    // WDR_MODE_NONE or WDR_MODE_2To1_LINE TODO
    stWdrAttr.enWDRMode = WDR_MODE_NONE;
    stWdrAttr.bCompress = HI_FALSE;

    mpp_error_code = HI_MPI_VI_SetWDRAttr(0, &stWdrAttr);
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_VI_SetWDRAttr, mpp_error_code);
    }
    
    mpp_error_code = HI_MPI_VI_EnableDev(0);
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_VI_EnableDev, mpp_error_code);
    }

    RECT_S stCapRect;

    stCapRect.s32X = in->x0;
    stCapRect.s32Y = in->y0;
    stCapRect.u32Width  = in->width;
    stCapRect.u32Height = in->height;


    VI_CHN_ATTR_S stChnAttr;

    memcpy(&stChnAttr.stCapRect, &stCapRect, sizeof(RECT_S));

    stChnAttr.enCapSel = VI_CAPSEL_BOTH;
    stChnAttr.stDestSize.u32Width = in->width;
    stChnAttr.stDestSize.u32Height =  in->height;
    stChnAttr.enPixFormat = PIXEL_FORMAT_YUV_SEMIPLANAR_420;   // sp420 or sp422

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


    stChnAttr.s32SrcFrameRate = in->cmos_fps;
    stChnAttr.s32DstFrameRate = in->fps;
    //stChnAttr.enCompressMode = COMPRESS_MODE_SEG;
    stChnAttr.enCompressMode = COMPRESS_MODE_NONE;

    mpp_error_code = HI_MPI_VI_SetChnAttr(0, &stChnAttr);
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_VI_SetChnAttr, mpp_error_code);
    }

    mpp_error_code = HI_MPI_VI_EnableChn(0);
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_VI_EnableChn, mpp_error_code);
    }

    return ERR_NONE;
}
#endif // defined(HI3516CV300)


#if defined(HI3516CV500)
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

int mpp_vi_init(unsigned int *error_code, void *videv, unsigned int width, unsigned int height, unsigned int fps) {
    *error_code = 0;


      //VI_StartDev
        VI_DEV_ATTR_S       stViDevAttr;
        hi_memcpy(&stViDevAttr, sizeof(VI_DEV_ATTR_S), videv, sizeof(VI_DEV_ATTR_S));
        stViDevAttr.stWDRAttr.enWDRMode = WDR_MODE_NONE;
        //stViDevAttr.enDataRate = DATA_RATE_X2; //???????????????

        *error_code = HI_MPI_VI_SetDevAttr(0, &stViDevAttr);
        if (*error_code != HI_SUCCESS) {
            printf("HI_MPI_VI_SetDevAttr failed with %#x!\n", *error_code);
            return -1;
        }

        *error_code = HI_MPI_VI_EnableDev(0);
        if (*error_code != HI_SUCCESS) {
            printf("HI_MPI_VI_EnableDev failed with %#x!\n", *error_code);
            return -1;
        }        

        //VI_BindPipeDev
        VI_DEV_BIND_PIPE_S  stDevBindPipe = {0};
        stDevBindPipe.u32Num = 1;

        *error_code = HI_MPI_VI_SetDevBindPipe(0, &stDevBindPipe);
        if (*error_code != HI_SUCCESS) {
            printf("HI_MPI_VI_SetDevBindPipe failed with %#x!\n", *error_code);
            return -1;
        }
        //VI_StartViPipe
        VI_PIPE_ATTR_S  stPipeAttr;

        hi_memcpy(&stPipeAttr, sizeof(VI_PIPE_ATTR_S), &PIPE_ATTR_2592x1944_RAW12_420_3DNR_RFR, sizeof(VI_PIPE_ATTR_S));

        *error_code = HI_MPI_VI_CreatePipe(0, &stPipeAttr);
        if (*error_code != HI_SUCCESS) {
            printf("HI_MPI_VI_CreatePipe failed with %#x!\n", *error_code);
            return -1;
        }

        *error_code = HI_MPI_VI_StartPipe(0);
        if (*error_code != HI_SUCCESS) {
            //HI_MPI_VI_DestroyPipe(ViPipe);
            printf("HI_MPI_VI_StartPipe failed with %#x!\n", *error_code);
            return -1;
        }
        //VI_StartViChn
        VI_CHN_ATTR_S       stChnAttr;

        hi_memcpy(&stChnAttr, sizeof(VI_CHN_ATTR_S), &CHN_ATTR_2592x1944_420_SDR8_LINEAR, sizeof(VI_CHN_ATTR_S));

        stChnAttr.enDynamicRange = DYNAMIC_RANGE_SDR8;
        stChnAttr.enVideoFormat  = VIDEO_FORMAT_LINEAR;
        stChnAttr.enPixelFormat  = PIXEL_FORMAT_YVU_SEMIPLANAR_420;
        stChnAttr.enCompressMode = COMPRESS_MODE_SEG;

        *error_code = HI_MPI_VI_SetChnAttr(0, 0, &stChnAttr);
        if (*error_code != HI_SUCCESS) {
            printf("HI_MPI_VI_SetChnAttr failed with %#x!\n", *error_code);
            return -1;
        }

        *error_code = HI_MPI_VI_EnableChn(0, 0);
        if (*error_code != HI_SUCCESS) {
            printf("HI_MPI_VI_EnableChn failed with %#x!\n", -1);
            return -1;
        }



    return ERR_NONE;
}

#endif // defined(HI3516CV500)
