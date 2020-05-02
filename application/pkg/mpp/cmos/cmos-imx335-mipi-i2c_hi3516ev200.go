//+build arm
//+build hi3516ev200
//+build imx335,cmos_data_mipi,cmos_control_i2c,cmos_bus_0

package cmos

/*
#include "../include/mpp.h"

#include "cmos.h"

#include <string.h>

int mpp_cmos_init(int *error_code) {
    *error_code = 0;

    ALG_LIB_S stAeLib;
    ALG_LIB_S stAwbLib;
    const ISP_SNS_OBJ_S* pstSnsObj;

    pstSnsObj = &stSnsImx335Obj;

    stAeLib.s32Id = 0;
    stAwbLib.s32Id = 0;
    strncpy(stAeLib.acLibName, HI_AE_LIB_NAME, sizeof(HI_AE_LIB_NAME));
    strncpy(stAwbLib.acLibName, HI_AWB_LIB_NAME, sizeof(HI_AWB_LIB_NAME));
    
    *error_code = pstSnsObj->pfnRegisterCallback(0, &stAeLib, &stAwbLib);
    if (*error_code != HI_SUCCESS) {
        printf("sensor_register_callback failed with %#x!\n", *error_code);
        return ERR_GENERAL;
    }

    ISP_SNS_COMMBUS_U uSnsBusInfo;
    ISP_SNS_TYPE_E enBusType;

    enBusType = ISP_SNS_I2C_TYPE;
    uSnsBusInfo.s8I2cDev = 0;

    *error_code = pstSnsObj->pfnSetBusInfo(0, uSnsBusInfo);
    if (*error_code != HI_SUCCESS) {
        printf("set sensor bus info failed with %#x!\n", *error_code);
        return ERR_GENERAL;
    }

    return ERR_NONE;
}

combo_dev_attr_t MIPI_4lane_CHN0_SENSOR_IMX335_12BIT_5M_NOWDR_ATTR =
{
    .devno = 0,
    .input_mode = INPUT_MODE_MIPI,
    .data_rate = MIPI_DATA_RATE_X1,
    .img_rect = {0, 0, 2592, 1944},

    {
        .mipi_attr =
        {
            DATA_TYPE_RAW_12BIT,
            HI_MIPI_WDR_MODE_NONE,
            {0, 1, 2, 3}
        }
    }
};

combo_dev_attr_t MIPI_4lane_CHN0_SENSOR_IMX335_10BIT_5M_WDR2TO1_ATTR =
{
    .devno = 0,
    .input_mode = INPUT_MODE_MIPI,
    .data_rate = MIPI_DATA_RATE_X1,
    .img_rect = {0, 0, 2592, 1944},

    {
        .mipi_attr =
        {
            DATA_TYPE_RAW_10BIT,
            HI_MIPI_WDR_MODE_VC,
            {0, 1, 2, 3}
        }
    }
};

combo_dev_attr_t MIPI_4lane_CHN0_SENSOR_IMX335_12BIT_4M_NOWDR_ATTR =
{
    .devno = 0,
    .input_mode = INPUT_MODE_MIPI,
    .data_rate = MIPI_DATA_RATE_X1,
    .img_rect = {0, 204, 2592, 1536},

    {
        .mipi_attr =
        {
            DATA_TYPE_RAW_12BIT,
            HI_MIPI_WDR_MODE_NONE,
            {0, 1, 2, 3}
        }
    }
};

combo_dev_attr_t MIPI_4lane_CHN0_SENSOR_IMX335_10BIT_4M_WDR2TO1_ATTR =
{
    .devno = 0,
    .input_mode = INPUT_MODE_MIPI,
    .data_rate = MIPI_DATA_RATE_X1,
    .img_rect = {0, 204, 2592, 1536},

    {
        .mipi_attr =
        {
            DATA_TYPE_RAW_10BIT,
            HI_MIPI_WDR_MODE_VC,
            {0, 1, 2, 3}
        }
    }
};

VI_DEV_ATTR_S DEV_ATTR_IMX335_5M_BASE =
{
    VI_MODE_MIPI,
    VI_WORK_MODE_1Multiplex,
    {0xFFF00000,    0x0},
    VI_SCAN_PROGRESSIVE,
    {-1, -1, -1, -1},
    VI_DATA_SEQ_YUYV,

    {
    //port_vsync   port_vsync_neg     port_hsync        port_hsync_neg        
    VI_VSYNC_PULSE, VI_VSYNC_NEG_LOW, VI_HSYNC_VALID_SINGNAL,VI_HSYNC_NEG_HIGH,VI_VSYNC_VALID_SINGAL,VI_VSYNC_VALID_NEG_HIGH,

    //hsync_hfb    hsync_act    hsync_hhb
    {0,            1280,        0,
    //vsync0_vhb vsync0_act vsync0_hhb
     0,            720,        0,
    //vsync1_vhb vsync1_act vsync1_hhb
     0,            0,            0}
    },
    VI_DATA_TYPE_RGB,
    HI_FALSE,
    {2592 , 1944},
    {
        {
            {2592 , 1944},
        },
        {
            VI_REPHASE_MODE_NONE,
            VI_REPHASE_MODE_NONE
        }
    },
    {
        WDR_MODE_NONE,
        1944
    },
    DATA_RATE_X1
};

VI_DEV_ATTR_S DEV_ATTR_IMX335_4M_BASE =
{
    VI_MODE_MIPI,
    VI_WORK_MODE_1Multiplex,
    {0xFFF00000,    0x0},
    VI_SCAN_PROGRESSIVE,
    {-1, -1, -1, -1},
    VI_DATA_SEQ_YUYV,

    {
    //port_vsync   port_vsync_neg     port_hsync        port_hsync_neg      
    VI_VSYNC_PULSE, VI_VSYNC_NEG_LOW, VI_HSYNC_VALID_SINGNAL,VI_HSYNC_NEG_HIGH,VI_VSYNC_VALID_SINGAL,VI_VSYNC_VALID_NEG_HIGH,

    //hsync_hfb    hsync_act    hsync_hhb
    {0,            1280,        0,
    //vsync0_vhb vsync0_act vsync0_hhb
     0,            720,        0,
    //vsync1_vhb vsync1_act vsync1_hhb
     0,            0,            0}
    },
    VI_DATA_TYPE_RGB,
    HI_FALSE,
    {2592 , 1536},
    {
        {
            {2592 , 1536},
        },
        {
            VI_REPHASE_MODE_NONE,
            VI_REPHASE_MODE_NONE
        }
    },
    {
        WDR_MODE_NONE,
        1536
    },
    DATA_RATE_X1
};

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
     
static VI_PIPE_ATTR_S PIPE_ATTR_2592x1944_RAW10_420_3DNR_RFR =
{
    VI_PIPE_BYPASS_NONE, HI_FALSE, HI_FALSE,
    2592, 1944,
    PIXEL_FORMAT_RGB_BAYER_10BPP,
    COMPRESS_MODE_LINE,
    DATA_BITWIDTH_10,
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

static VI_PIPE_ATTR_S PIPE_ATTR_2592x1536_RAW12_420_3DNR_RFR =
{
    VI_PIPE_BYPASS_NONE, HI_FALSE, HI_FALSE,
    2592, 1536,
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

static VI_PIPE_ATTR_S PIPE_ATTR_2592x1536_RAW10_420_3DNR_RFR =
{
    VI_PIPE_BYPASS_NONE, HI_FALSE, HI_FALSE,
    2592, 1536,
    PIXEL_FORMAT_RGB_BAYER_10BPP,
    COMPRESS_MODE_LINE,
    DATA_BITWIDTH_10,
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



*/
import "C"

import (    
    "unsafe"
)

var (
    cmosItem = cmos{   
        vendor: "Sony", 
        model: "imx335",   
        modes: []cmosMode {
            cmosMode {
                width: 2592, 
                height: 1944,
                fps: 30,
                mipi: unsafe.Pointer(&C.MIPI_4lane_CHN0_SENSOR_IMX335_12BIT_5M_NOWDR_ATTR),
                viDev: unsafe.Pointer(&C.DEV_ATTR_IMX335_5M_BASE),
                clock: 0,
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

