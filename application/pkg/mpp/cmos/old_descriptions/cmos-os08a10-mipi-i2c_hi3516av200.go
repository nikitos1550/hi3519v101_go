//+build nobuild

//+build arm
//+build hi3516av200
//+build os08a10,cmos_data_mipi,cmos_control_i2c,cmos_bus_0

package cmos

/*
#include "cmos.h"

int mpp_cmos_init(int *error_code) {
	*error_code = 0;

    ALG_LIB_S stAeLib;
    ALG_LIB_S stAwbLib;

    stAeLib.s32Id = 0;
    stAwbLib.s32Id = 0;
    strncpy(stAeLib.acLibName,  HI_AE_LIB_NAME,     sizeof(HI_AE_LIB_NAME));
    strncpy(stAwbLib.acLibName, HI_AWB_LIB_NAME,    sizeof(HI_AWB_LIB_NAME));

    ISP_SNS_OBJ_S *cmos = &stSnsOs08a10Obj;
    if (cmos->pfnRegisterCallback != HI_NULL) {
        *error_code = cmos->pfnRegisterCallback(0, &stAeLib, &stAwbLib);
        if (*error_code != HI_SUCCESS) return ERR_GENERAL;
    } else {
        return ERR_GENERAL;
    }

    return ERR_NONE;
}

combo_dev_attr_t MIPI_4lane_SENSOR_OS08A_12BIT_4K_NOWDR_ATTR =
{
    .devno = 0,
    .input_mode = INPUT_MODE_MIPI,
    .phy_clk_share = PHY_CLK_SHARE_NONE,
    .img_rect = {0, 0, 3840, 2160},

    .mipi_attr =
    {
        .raw_data_type = RAW_DATA_12BIT,
        .wdr_mode = HI_MIPI_WDR_MODE_NONE,
        .lane_id = {0, 1, 2, 3, -1, -1, -1, -1}
    }
};

combo_dev_attr_t MIPI_4lane_SENSOR_OS08A_10BIT_4K_2WDR1_ATTR =
{
    .devno = 0,
    .input_mode = INPUT_MODE_MIPI,
    .phy_clk_share = PHY_CLK_SHARE_NONE,
    .img_rect = {0, 0, 3840, 2160},

    .mipi_attr =
    {
        .raw_data_type = RAW_DATA_10BIT,
        .wdr_mode = HI_MIPI_WDR_MODE_NONE,
        .lane_id = {0, 1, 2, 3, -1, -1, -1, -1}
    }
};

//VI_DEV_ATTR_S DEV_ATTR_MIPI_BASE =
//{
//    // interface mode 
//    VI_MODE_MIPI,
//    // multiplex mode 
//    VI_WORK_MODE_1Multiplex,
//    // r_mask    g_mask    b_mask
//    {0xFFF00000,    0x0},
//    // progessive or interleaving 
//    VI_SCAN_PROGRESSIVE,
//    //AdChnId
//    { -1, -1, -1, -1},
//    //enDataSeq, only support yuv
//    VI_INPUT_DATA_YUYV,
//
//    // synchronization information 
//    {
//        //port_vsync   port_vsync_neg     port_hsync        port_hsync_neg       
//        VI_VSYNC_PULSE, VI_VSYNC_NEG_LOW, VI_HSYNC_VALID_SINGNAL, VI_HSYNC_NEG_HIGH, VI_VSYNC_VALID_SINGAL, VI_VSYNC_VALID_NEG_HIGH,
//
//        //hsync_hfb    hsync_act    hsync_hhb
//        {
//            0,            1280,        0,
//            //vsync0_vhb vsync0_act vsync0_hhb
//            0,            720,        0,
//            //vsync1_vhb vsync1_act vsync1_hhb
//            0,            0,            0
//        }
//    },
//    // use interior ISP
//    VI_PATH_ISP,
//    // input data type
//    VI_DATA_TYPE_RGB,
//    // bRever
//    HI_FALSE,
//    // DEV CROP
//    {0, 0, 3840, 2160},
//    {
//        {
//            {1920, 1080},
//            HI_FALSE,
//
//        },
//        {
//            VI_REPHASE_MODE_NONE,
//            VI_REPHASE_MODE_NONE
//        }
//    }
//};
*/
import "C"

import (
	"unsafe"
)

var (
	cmosItem = cmos{
		vendor: "OmniVision",
		model: "os08a10",
		modes: []cmosMode {
			cmosMode {
				width: 3840,
				height: 2160,
				fps: 30,
                bitness: 12,
				mipi: unsafe.Pointer(&C.MIPI_4lane_SENSOR_OS08A_12BIT_4K_NOWDR_ATTR),
                //viDev: unsafe.Pointer(&C.DEV_ATTR_MIPI_BASE),
                clock: 24,
                wdr: WDRNone,
                description: "normal",
			},
		},
		control: cmosControl {
			bus: I2C,
			busNum: 0,
		},
        data: MIPI,
        bayer: RGGB,
	}
)