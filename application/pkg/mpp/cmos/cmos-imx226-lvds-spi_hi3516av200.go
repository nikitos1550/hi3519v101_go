//+build arm
//+build hi3516av200
//+build imx226,cmos_data_lvds,cmos_control_spi,cmos_bus_0

package cmos

/*
#include "../include/mpp.h"
#include "cmos.h"

#include <string.h>

int mpp_cmos_init(int *error_code) {
	*error_code = 0;

    ALG_LIB_S stAeLib;
    ALG_LIB_S stAwbLib;

    stAeLib.s32Id = 0;
    stAwbLib.s32Id = 0;
    strncpy(stAeLib.acLibName,  HI_AE_LIB_NAME,     sizeof(HI_AE_LIB_NAME));
    strncpy(stAwbLib.acLibName, HI_AWB_LIB_NAME,    sizeof(HI_AWB_LIB_NAME));

    ISP_SNS_OBJ_S *cmos = &stSnsImx226Obj;
    if (cmos->pfnRegisterCallback != HI_NULL) {
        *error_code = cmos->pfnRegisterCallback(0, &stAeLib, &stAwbLib);
        if (*error_code != HI_SUCCESS) return ERR_GENERAL;
    } else {
        return ERR_GENERAL;
    }

    return ERR_NONE;
}

combo_dev_attr_t LVDS_4lane_SENSOR_IMX226_12BIT_8M_NOWDR_ATTR =
{
    .devno = 0,
    // input mode
    .input_mode = INPUT_MODE_LVDS,
    .phy_clk_share = PHY_CLK_SHARE_PHY0,
    .img_rect = {0, 0, 4248, 2182},

    {
        .lvds_attr = {
            .raw_data_type    = RAW_DATA_12BIT,
            .wdr_mode         = HI_WDR_MODE_NONE,
            .sync_mode        = LVDS_SYNC_MODE_SAV,
            .vsync_type       = {LVDS_VSYNC_NORMAL, 0, 0},
            .fid_type         = {LVDS_FID_NONE, HI_TRUE},
            .data_endian      = LVDS_ENDIAN_BIG,
            .sync_code_endian = LVDS_ENDIAN_BIG,
            .lane_id = { -1, -1, 0, -1, 1, 2, -1, 3, -1, -1, -1, -1},
            .sync_code = {
                {
                    {0xab0, 0xb60, 0x800, 0x9d0},      // lane 0
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0}
                },

                {
                    {0xab0, 0xb60, 0x800, 0x9d0},      // lane 1
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0}
                },

                {
                    {0xab0, 0xb60, 0x800, 0x9d0},      // lane2
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0}
                },

                {
                    {0xab0, 0xb60, 0x800, 0x9d0},      // lane3
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0}
                },
            }
        }
    }
};

// 8lane 30fps
combo_dev_attr_t LVDS_8lane_SENSOR_IMX226_12BIT_8M_NOWDR_ATTR =
{
    .devno = 0,
    // input mode
    .input_mode = INPUT_MODE_LVDS,
    .phy_clk_share = PHY_CLK_SHARE_PHY0,
    .img_rect = {252, 18, 3840, 2160},

    .lvds_attr =
    {
        .raw_data_type    = RAW_DATA_12BIT,
        .wdr_mode         = HI_WDR_MODE_NONE,
        .sync_mode        = LVDS_SYNC_MODE_SAV,
        .vsync_type       = {LVDS_VSYNC_NORMAL, 0, 0},
        .fid_type         = {LVDS_FID_NONE, HI_TRUE},
        .data_endian      = LVDS_ENDIAN_BIG,
        .sync_code_endian = LVDS_ENDIAN_BIG,
        .lane_id = {0, 1, 2, -1, 3, 4, -1, 5, 6, 7, -1, -1},
        .sync_code =
        {
            {   {0xab0, 0xb60, 0x800, 0x9d0},      // lane 0
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0}
            },

            {   {0xab0, 0xb60, 0x800, 0x9d0},      // lane 1
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0}
            },

            {   {0xab0, 0xb60, 0x800, 0x9d0},      // lane2
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0}
            },

            {   {0xab0, 0xb60, 0x800, 0x9d0},      // lane3
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0}
            },

            {   {0xab0, 0xb60, 0x800, 0x9d0},      // lane4
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0}
            },

            {   {0xab0, 0xb60, 0x800, 0x9d0},      // lane5
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0}
            },

            {   {0xab0, 0xb60, 0x800, 0x9d0},      // lane6
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0}
            },

            {   {0xab0, 0xb60, 0x800, 0x9d0},      // lane7
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0}
            },
        }
    }
};

combo_dev_attr_t LVDS_10lane_SENSOR_IMX226_10BIT_8M_NOWDR_ATTR =
{
    .devno = 0,
    .input_mode = INPUT_MODE_LVDS,
    .phy_clk_share = PHY_CLK_SHARE_PHY0,
    .img_rect = {128, 0, 4120, 2182},
    .lvds_attr =
    {
        .raw_data_type    = RAW_DATA_10BIT,
        .wdr_mode         = HI_WDR_MODE_NONE,
        .sync_mode        = LVDS_SYNC_MODE_SAV,
        .vsync_type       = {LVDS_VSYNC_NORMAL, 0, 0},
        .fid_type         = {LVDS_FID_NONE, HI_TRUE},
        .data_endian      = LVDS_ENDIAN_BIG,
        .sync_code_endian = LVDS_ENDIAN_BIG,
        .lane_id = {0, 1, 2, 3, 4, 5, 6, 7, 8, 9, -1, -1},
        .sync_code =
        {
            {{0x2ac, 0x2d8, 0x200, 0x274},      // lane 0
                {0x2ac, 0x2d8, 0x200, 0x274},
                {0x2ac, 0x2d8, 0x200, 0x274},
                {0x2ac, 0x2d8, 0x200, 0x274}},
            {{0x2ac, 0x2d8, 0x200, 0x274},      // lane 1
                {0x2ac, 0x2d8, 0x200, 0x274},
                {0x2ac, 0x2d8, 0x200, 0x274},
                {0x2ac, 0x2d8, 0x200, 0x274}},
            {{0x2ac, 0x2d8, 0x200, 0x274},      // lane 2
                {0x2ac, 0x2d8, 0x200, 0x274},
                {0x2ac, 0x2d8, 0x200, 0x274},
                {0x2ac, 0x2d8, 0x200, 0x274}},
            {{0x2ac, 0x2d8, 0x200, 0x274},      // lane 3
                {0x2ac, 0x2d8, 0x200, 0x274},
                {0x2ac, 0x2d8, 0x200, 0x274},
                {0x2ac, 0x2d8, 0x200, 0x274}},
            {{0x2ac, 0x2d8, 0x200, 0x274},      // lane 4
                {0x2ac, 0x2d8, 0x200, 0x274},
                {0x2ac, 0x2d8, 0x200, 0x274},
                {0x2ac, 0x2d8, 0x200, 0x274}},
            {{0x2ac, 0x2d8, 0x200, 0x274},      // lane 5
                {0x2ac, 0x2d8, 0x200, 0x274},
                {0x2ac, 0x2d8, 0x200, 0x274},
                {0x2ac, 0x2d8, 0x200, 0x274}},
            {{0x2ac, 0x2d8, 0x200, 0x274},      // lane 6
                {0x2ac, 0x2d8, 0x200, 0x274},
                {0x2ac, 0x2d8, 0x200, 0x274},
                {0x2ac, 0x2d8, 0x200, 0x274}},
            {{0x2ac, 0x2d8, 0x200, 0x274},      // lane 7
                {0x2ac, 0x2d8, 0x200, 0x274},
                {0x2ac, 0x2d8, 0x200, 0x274},
                {0x2ac, 0x2d8, 0x200, 0x274}},
            {{0x2ac, 0x2d8, 0x200, 0x274},      // lane 8
                {0x2ac, 0x2d8, 0x200, 0x274},
                {0x2ac, 0x2d8, 0x200, 0x274},
                {0x2ac, 0x2d8, 0x200, 0x274}},
            {{0x2ac, 0x2d8, 0x200, 0x274},      // lane 9
                {0x2ac, 0x2d8, 0x200, 0x274},
                {0x2ac, 0x2d8, 0x200, 0x274},
                {0x2ac, 0x2d8, 0x200, 0x274}},
        }
    }
};
combo_dev_attr_t LVDS_10lane_SENSOR_IMX226_12BIT_12M_NOWDR_ATTR =
{
    .devno = 0,
    .input_mode = INPUT_MODE_LVDS,
    .phy_clk_share = PHY_CLK_SHARE_PHY0,
    .img_rect = {132, 40, 4000, 3000},
    .lvds_attr =
    {
        .raw_data_type    = RAW_DATA_12BIT,
        .wdr_mode         = HI_WDR_MODE_NONE,
        .sync_mode        = LVDS_SYNC_MODE_SAV,
        .vsync_type       = {LVDS_VSYNC_NORMAL, 0, 0},
        .fid_type         = {LVDS_FID_NONE, HI_TRUE},
        .data_endian      = LVDS_ENDIAN_BIG,
        .sync_code_endian = LVDS_ENDIAN_BIG,
        .lane_id = {0, 1, 2, 3, 4, 5, 6, 7, 8, 9, -1, -1},
        .sync_code =
        {
            {{0xab0, 0xb60, 0x800, 0x9d0},      // lane 0
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0}},
            {{0xab0, 0xb60, 0x800, 0x9d0},      // lane 1
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0}},
            {{0xab0, 0xb60, 0x800, 0x9d0},      // lane2
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0}},
            {{0xab0, 0xb60, 0x800, 0x9d0},      // lane3
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0}},
            {{0xab0, 0xb60, 0x800, 0x9d0},      // lane4
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0}},
            {{0xab0, 0xb60, 0x800, 0x9d0},      // lane5
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0}},
            {{0xab0, 0xb60, 0x800, 0x9d0},      // lane6
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0}},
            {{0xab0, 0xb60, 0x800, 0x9d0},      // lane7
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0}},
			  {{0xab0, 0xb60, 0x800, 0x9d0},      // lane8
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0}},
			  {{0xab0, 0xb60, 0x800, 0x9d0},      // lane9
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0}},
        }
    }
};

// 4Kx3K crop to 3kx3k
combo_dev_attr_t LVDS_10lane_SENSOR_IMX226_12BIT_9M_NOWDR_ATTR =
{
    .devno = 0,
    .input_mode = INPUT_MODE_LVDS,
    .phy_clk_share = PHY_CLK_SHARE_PHY0,
    .img_rect = {632, 40, 3000, 3000},
    .lvds_attr =
    {
        .raw_data_type    = RAW_DATA_12BIT,
        .wdr_mode         = HI_WDR_MODE_NONE,
        .sync_mode        = LVDS_SYNC_MODE_SAV,
        .vsync_type       = {LVDS_VSYNC_NORMAL, 0, 0},
        .fid_type         = {LVDS_FID_NONE, HI_TRUE},
        .data_endian      = LVDS_ENDIAN_BIG,
        .sync_code_endian = LVDS_ENDIAN_BIG,
        .lane_id = {0, 1, 2, 3, 4, 5, 6, 7, 8, 9, -1, -1},
        .sync_code =
        {
            {{0xab0, 0xb60, 0x800, 0x9d0},      // lane 0
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0}},
            {{0xab0, 0xb60, 0x800, 0x9d0},      // lane 1
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0}},
            {{0xab0, 0xb60, 0x800, 0x9d0},      // lane2
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0}},
            {{0xab0, 0xb60, 0x800, 0x9d0},      // lane3
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0}},
            {{0xab0, 0xb60, 0x800, 0x9d0},      // lane4
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0}},
            {{0xab0, 0xb60, 0x800, 0x9d0},      // lane5
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0}},
            {{0xab0, 0xb60, 0x800, 0x9d0},      // lane6
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0}},
            {{0xab0, 0xb60, 0x800, 0x9d0},      // lane7
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0}},
			  {{0xab0, 0xb60, 0x800, 0x9d0},      // lane8
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0}},
			  {{0xab0, 0xb60, 0x800, 0x9d0},      // lane9
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0},
                {0xab0, 0xb60, 0x800, 0x9d0}},
        }
    }
};

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


*/
import "C"

import (
	"unsafe"
)

var (
	cmosItem = cmos{
		vendor: "Sony",
		model: "imx226",
		modes: []cmosMode {
			cmosMode {
				width: 3840,
				height: 2160,
				fps: 30,
				mipi: unsafe.Pointer(&C.LVDS_8lane_SENSOR_IMX226_12BIT_8M_NOWDR_ATTR),
                viDev: unsafe.Pointer(&C.DEV_ATTR_LVDS_BASE),
                clock: 72,
                wdr: WDRNone,
                description: "normal",
			},
		},
        control: cmosControl {
            bus: SPI,
            busNum: 0,
        },
        data: LVDS,
        bayer: RGGB,
	}
)
