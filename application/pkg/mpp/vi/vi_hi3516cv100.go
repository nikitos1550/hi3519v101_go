//+build arm
//+build hi3516cv100

package vi

/*

#include "../include/mpp_v1.h"
#include <string.h>

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

#define ERR_NONE                    0
#define ERR_HI_MPI_VI_SetDevAttr        2
#define ERR_HI_MPI_VI_EnableDev     3
#define ERR_HI_MPI_VI_SetChnAttr    4
#define ERR_HI_MPI_VI_EnableChn     5

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
    //	if (*error_code != HI_SUCCESS) return ERR_
    //}
    
    *error_code = HI_MPI_VI_EnableChn(ViChn);
    if (*error_code != HI_SUCCESS) return ERR_HI_MPI_VI_EnableChn;

    return ERR_NONE;
}


*/
import "C"

import (
        "application/pkg/logger"

    "application/pkg/mpp/error"
)

func Init() {
    var errorCode C.uint

    switch err := C.mpp1_vi_init(&errorCode); err {
    case C.ERR_NONE:
        logger.Log.Debug().
                Msg("C.mpp1_vi_init() ok")
    case C.ERR_HI_MPI_VI_SetDevAttr:
        logger.Log.Fatal().
                Str("func", "ERR_HI_MPI_VI_SetDevAttr()").
                Int("error", int(errorCode)).
                Str("error_desc", error.Resolve(int64(errorCode))).
                Msg("C.mpp1_vi_init() error")
    case C.ERR_HI_MPI_VI_EnableDev:
        logger.Log.Fatal().
                Str("func", "ERR_HI_MPI_VI_EnableDev()").
                Int("error", int(errorCode)).
                Str("error_desc", error.Resolve(int64(errorCode))).
                Msg("C.mpp1_vi_init() error")

    case C.ERR_HI_MPI_VI_SetChnAttr:
        logger.Log.Fatal().
                Str("func", "ERR_HI_MPI_VI_SetChnAttr()"). 
                Int("error", int(errorCode)).
                Str("error_desc", error.Resolve(int64(errorCode))).
                Msg("C.mpp1_vi_init() error")
    case C.ERR_HI_MPI_VI_EnableChn:
        logger.Log.Fatal().
                Str("func", "ERR_HI_MPI_VI_EnableChn()"). 
                Int("error", int(errorCode)).
                Str("error_desc", error.Resolve(int64(errorCode))).
                Msg("C.mpp1_vi_init() error")
    default:
        logger.Log.Fatal().
                Int("error", int(err)).
                Msg("C.mpp1_vi_init() Unexpected return")

    }
}

