//+build arm
//+build hi3516cv300
//+build imx323,cmos_data_dc,cmos_control_spi,cmos_bus_0

package cmos


/*
#include "../include/mpp_v3.h"

#include "cmos.h"

#include <string.h>

//int mpp_cmos_init(int *error_code) {
//    *error_code = 0;
//
//    return ERR_NONE;
//}

combo_dev_attr_t MIPI_CMOS323_ATTR = 
{
    // input mode
    .input_mode = INPUT_MODE_CMOS,
    {
        
    }
};

VI_DEV_ATTR_S DEV_ATTR_DC_IMX323 = 
{
    // interface mode 
    VI_MODE_DIGITAL_CAMERA,
    // multiplex mode
    VI_WORK_MODE_1Multiplex,
    // r_mask    g_mask    b_mask
    {0xFFF0000,    0x0},
    // progessive or interleaving
    VI_SCAN_PROGRESSIVE,
    //AdChnId
    { -1, -1, -1, -1},
    //enDataSeq, only support yuv
    VI_INPUT_DATA_YUYV,

    // synchronization information
    { 
        //port_vsync   port_vsync_neg     port_hsync        port_hsync_neg
        VI_VSYNC_PULSE, VI_VSYNC_NEG_HIGH, VI_HSYNC_VALID_SINGNAL, VI_HSYNC_NEG_HIGH, VI_VSYNC_VALID_SINGAL, VI_VSYNC_VALID_NEG_HIGH,
 
        //hsync_hfb    hsync_act    hsync_hhb
        {
            0,            1920,        0,
            //vsync0_vhb vsync0_act vsync0_hhb
            0,            1080,        0,
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
    {200, 20, 1920, 1080}
};

*/
import "C"

import (
    "unsafe"
)

var (
    cmosItem = cmos{
        vendor: "Sony",
        model: "imx323",
        modes: []cmosMode {
            cmosMode {
                width: 1920,
                height: 1080,
                fps: 30,
                mipi: unsafe.Pointer(&C.MIPI_CMOS323_ATTR),
                viDev: unsafe.Pointer(&C.DEV_ATTR_DC_IMX323),
                clock: 37.125,
                wdr: WDRNone,
                description: "normal",                
            },
        },
        control: cmosControl {
            bus: SPI,
            busNum: 0,
        },
        data: DC,
        bayer: RGGB,
    }
)

