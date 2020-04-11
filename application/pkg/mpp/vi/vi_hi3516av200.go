//+build arm
//+build hi3516av200

package vi

/*
#include "../include/mpp_v3.h"
#include <string.h>

VI_DEV_ATTR_S DEV_ATTR_LVDS_BASE =
{
    // interface mode
    VI_MODE_LVDS,
    // multiplex mode
    VI_WORK_MODE_1Multiplex,
    // r_mask    g_mask    b_mask
    {0xFFF00000,    0x0},
    // progessive or interleaving
    VI_SCAN_PROGRESSIVE,
    //AdChnId
    { -1, -1, -1, -1},
    //enDataSeq, only support yuv
    VI_INPUT_DATA_YUYV,

    // synchronization information
    {
        //port_vsync   port_vsync_neg     port_hsync        port_hsync_neg
        VI_VSYNC_PULSE, VI_VSYNC_NEG_LOW, VI_HSYNC_VALID_SINGNAL, VI_HSYNC_NEG_HIGH, VI_VSYNC_VALID_SINGAL, VI_VSYNC_VALID_NEG_HIGH,

        //hsync_hfb    hsync_act    hsync_hhb
        {
            0,            1280,        0,
            //vsync0_vhb vsync0_act vsync0_hhb
            0,            720,        0,
            //vsync1_vhb vsync1_act vsync1_hhb
            0,            0,            0
        }
    },
    // use interior ISP
    VI_PATH_ISP,
    // input data type
    VI_DATA_TYPE_RGB,
    // bRever
    HI_FALSE,
    // DEV CROP
    {0, 0, 1920, 1080},
    {
        {
            {1920, 1080},
            HI_FALSE,

        },
        {
            VI_REPHASE_MODE_NONE,
            VI_REPHASE_MODE_NONE
        }
    }
};

#define ERR_NONE                    0
#define ERR_HI_MPI_VI_SetDevAttr        2
#define ERR_HI_MPI_VI_EnableDev     3
#define ERR_HI_MPI_VI_SetChnAttr    4
#define ERR_HI_MPI_VI_EnableChn     5

int mpp3_vi_init(unsigned int *error_code, void *videv, unsigned int width, unsigned int height, unsigned int fps) {
    *error_code = 0;

    VI_DEV_ATTR_S  stViDevAttr;

    memset(&stViDevAttr, 0, sizeof(stViDevAttr));
    
    //memcpy(&stViDevAttr, &DEV_ATTR_LVDS_BASE, sizeof(stViDevAttr));
    memcpy(&stViDevAttr, videv, sizeof(stViDevAttr));

    stViDevAttr.stDevRect.s32X                              = 0;
    stViDevAttr.stDevRect.s32Y                              = 0;
    stViDevAttr.stDevRect.u32Width                          = width; //3840;
    stViDevAttr.stDevRect.u32Height                         = height; //2160;
    stViDevAttr.stBasAttr.stSacleAttr.stBasSize.u32Width    = width; //3840;
    stViDevAttr.stBasAttr.stSacleAttr.stBasSize.u32Height   = height ;//2160;
    stViDevAttr.stBasAttr.stSacleAttr.bCompress             = HI_FALSE;

    *error_code = HI_MPI_VI_SetDevAttr(0, &stViDevAttr);
    if (*error_code != HI_SUCCESS) return ERR_HI_MPI_VI_SetDevAttr;

    *error_code = HI_MPI_VI_EnableDev(0);
    if (*error_code != HI_SUCCESS) return ERR_HI_MPI_VI_EnableDev;

    RECT_S stCapRect;
    SIZE_S stTargetSize;

    stCapRect.s32X          = 0;
    stCapRect.s32Y          = 0;
    stCapRect.u32Width      = width; //3840;
    stCapRect.u32Height     = height; //2160;
    stTargetSize.u32Width   = stCapRect.u32Width;
    stTargetSize.u32Height  = stCapRect.u32Height;

    VI_CHN_ATTR_S stChnAttr;

    memcpy(&stChnAttr.stCapRect, &stCapRect, sizeof(RECT_S));

    stChnAttr.enCapSel              = VI_CAPSEL_BOTH;
    stChnAttr.stDestSize.u32Width   = stTargetSize.u32Width ;
    stChnAttr.stDestSize.u32Height  = stTargetSize.u32Height ;
    stChnAttr.enPixFormat           = PIXEL_FORMAT_YUV_SEMIPLANAR_420;   // sp420 or sp422

    stChnAttr.bMirror               = HI_FALSE;
    stChnAttr.bFlip                 = HI_FALSE;

    stChnAttr.s32SrcFrameRate       = fps; //30;
    stChnAttr.s32DstFrameRate       = fps; //30;
    stChnAttr.enCompressMode        = COMPRESS_MODE_NONE;

    *error_code = HI_MPI_VI_SetChnAttr(0, &stChnAttr);
    if (*error_code != HI_SUCCESS) return ERR_HI_MPI_VI_SetChnAttr;

    *error_code = HI_MPI_VI_EnableChn(0);
    if (*error_code != HI_SUCCESS) return ERR_HI_MPI_VI_EnableChn;

    return ERR_NONE;
}
*/
import "C"

import (
    //"log"
	"application/pkg/logger"

	"application/pkg/mpp/cmos"

    "application/pkg/mpp/error"
)

func Init() {
    var errorCode C.uint

    switch err := C.mpp3_vi_init(&errorCode, cmos.ViDev(), C.uint(cmos.Width()), C.uint(cmos.Height()), C.uint(cmos.Fps())); err {
    case C.ERR_NONE:
        //log.Println("C.mpp3_vi_init() ok")
	logger.Log.Debug().
		Msg("C.mpp3_vi_init() ok")
    case C.ERR_HI_MPI_VI_SetDevAttr:
        // log.Fatal("C.mpp3_vi_init() ERR_HI_MPI_VI_SetDevAttr() error ", error.Resolve(int64(errorCode)))
	logger.Log.Fatal().
		Str("func", "ERR_HI_MPI_VI_SetDevAttr()").
		Int("error", int(errorCode)).
		Str("error_desc", error.Resolve(int64(errorCode))).
		Msg("C.mpp3_vi_init() error")
    case C.ERR_HI_MPI_VI_EnableDev:
        //log.Fatal("C.mpp3_vi_init() ERR_HI_MPI_VI_EnableDev() error ", error.Resolve(int64(errorCode)))
	logger.Log.Fatal().
                Str("func", "ERR_HI_MPI_VI_EnableDev()").
                Int("error", int(errorCode)).
                Str("error_desc", error.Resolve(int64(errorCode))).
                Msg("C.mpp3_vi_init() error")

    case C.ERR_HI_MPI_VI_SetChnAttr:
        //log.Fatal("C.mpp3_vi_init() ERR_HI_MPI_VI_SetChnAttr() error ", error.Resolve(int64(errorCode)))
	logger.Log.Fatal().
                Str("func", "ERR_HI_MPI_VI_SetChnAttr()"). 
                Int("error", int(errorCode)).
                Str("error_desc", error.Resolve(int64(errorCode))).
                Msg("C.mpp3_vi_init() error")
    case C.ERR_HI_MPI_VI_EnableChn:
        //log.Fatal("C.mpp3_vi_init() ERR_HI_MPI_VI_EnableChn() error ", error.Resolve(int64(errorCode)))        
	logger.Log.Fatal().
                Str("func", "ERR_HI_MPI_VI_EnableChn()"). 
                Int("error", int(errorCode)).
                Str("error_desc", error.Resolve(int64(errorCode))).
                Msg("C.mpp3_vi_init() error")
    default:
        //log.Fatal("Unexpected return ", err , " of C.mpp3_vi_init()")
	logger.Log.Fatal().
		Int("error", int(err)).
		Msg("C.mpp3_vi_init() Unexpected return")

    }
}
