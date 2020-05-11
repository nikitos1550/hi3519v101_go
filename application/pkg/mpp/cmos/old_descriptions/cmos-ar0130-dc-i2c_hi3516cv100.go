//+build nobuild

//+build arm
//+build hi3516cv100
//+build ar0130,cmos_data_dc,cmos_control_i2c,cmos_bus_0

package cmos

/*
#include "cmos.h"

int sensor_register_callback_ar0130(void);

int mpp_cmos_init(int *error_code) {
    *error_code = 0;

    *error_code = sensor_register_callback_ar0130();
    if (*error_code != HI_SUCCESS) {
        if (*error_code != HI_SUCCESS) return ERR_GENERAL;
    }

    return ERR_NONE;
}


//AR0130 DC 12bitÊäÈë720P@30fps
//VI_DEV_ATTR_S DEV_ATTR_AR0130_DC_720P =
//{
//    //½Ó¿ÚÄ£Ê½
//    VI_MODE_DIGITAL_CAMERA,
//    //1¡¢2¡¢4Â·¹¤×÷Ä£Ê½
//    VI_WORK_MODE_1Multiplex,
//    // r_mask    g_mask    b_mask
//    {0xFFF00000,    0x0}, 
//    //ÖðÐÐor¸ôÐÐÊäÈë
//    VI_SCAN_PROGRESSIVE,
//    //AdChnId
//    {-1, -1, -1, -1},
//    //enDataSeq, ½öÖ§³ÖYUV¸ñÊ½
//    VI_INPUT_DATA_YUYV,
//     
//    //Í¬²½ÐÅÏ¢£¬¶ÔÓ¦regÊÖ²áµÄÈçÏÂÅäÖÃ, --bt1120Ê±ÐòÎÞÐ§
//    {
//    //port_vsync   port_vsync_neg     port_hsync        port_hsync_neg      
//    VI_VSYNC_PULSE, VI_VSYNC_NEG_HIGH, VI_HSYNC_VALID_SINGNAL,VI_HSYNC_NEG_HIGH,VI_VSYNC_VALID_SINGAL,VI_VSYNC_VALID_NEG_HIGH,
//    
//    //timingÐÅÏ¢£¬¶ÔÓ¦regÊÖ²áµÄÈçÏÂÅäÖÃ
//    //hsync_hfb    hsync_act    hsync_hhb
//    {0,            1280,        0,
//    //vsync0_vhb vsync0_act vsync0_hhb
//     0,            720,        0,
//    //vsync1_vhb vsync1_act vsync1_hhb
//     0,            0,            0}
//    },    
//    //Ê¹ÓÃÄÚ²¿ISP
//    VI_PATH_ISP,
//    //ÊäÈëÊý¾ÝÀàÐÍ
//    VI_DATA_TYPE_RGB
//};


*/
import "C"

//import (
//    "unsafe"
//)

var (
    cmosItem = cmos{
        vendor: "Aptina",
        model: "ar0130",
        modes: []cmosMode {
            cmosMode {
                mipiCrop:   crop{X0: 0, Y0: 0, Width: 0, Height: 0,}, //not applicable, keep 0
                viCrop:     crop{X0: 0, Y0: 0, Width: 0, Height: 0,}, //not applicable, keep 0
                ispCrop:    crop{X0: 0, Y0: 0, Width: 1280, Height: 720,},
                width: 1280,
                height: 720,
                fps: 30,
                bitness: 12,
                data: DC,
                dcSync: dcSyncAttr {
                    VSync:          DCVSyncPulse,
                    VSyncNeg:       DCVSyncNegHigh,
                    HSync:          DCHSyncValid,
                    HSyncNeg:       DCHSyncNegHigh,
                    VSyncValid:     DCVSyncValidValid,
                    VSyncValidNeg:  DCVSyncValidNegHigh,
                    TimingHfb:      0,
                    TimingAct:      1280,
                    TimingHbb:      0,
                    TimingVfb:      0,
                    TimingVact:     720,
                    TimingVbb:      0,
                    TimingVbfb:     0,
                    TimingVbact:    0,
                    TimingVbbb:     0,
                },
                mipi: nil,
                //viDev: unsafe.Pointer(&C.DEV_ATTR_AR0130_DC_720P),
                clock: 27,
                wdr: WDRNone,
                description: "normal",
            },
        },
        control: cmosControl {
            bus: I2C,
            busNum: 0,
        },
        data: DC,
        bayer: RGGB,
    }
)

