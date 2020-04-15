//+build arm
//+build hi3516cv100
//+build imx122,cmos_data_dc,cmos_control_spi,cmos_bus_0

package cmos

/*
#include "../include/mpp_v1.h"

#include "cmos.h"

int mpp_cmos_init(int *error_code) {
    *error_code = 0;

    return ERR_NONE;
}

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


*/
import "C"

import (
    "unsafe"
)

var (
    cmosItem = cmos{
        vendor: "Sony",
        model: "imx122",
        modes: []cmosMode {
            cmosMode {
                width: 1920,
                height: 1080,
                fps: 30,
                mipi: nil,
                viDev: unsafe.Pointer(&C.DEV_ATTR_IMX122_DC_1080P),
                clock: 37.125,
            },
        },
        control: cmosControl {
            bus: SPI,
            busNum: 0,
        },
        data: DC,
    }
)

