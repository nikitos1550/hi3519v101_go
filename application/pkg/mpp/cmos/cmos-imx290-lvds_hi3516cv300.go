//+build arm
//+build hi3516cv300
//+build imx290,cmos_data_lvds,cmos_control_i2c,cmos_bus_0

package cmos


/*
#include "../include/mpp_v3.h"

#include "cmos.h"

#include <string.h>

int mpp_cmos_init(int *error_code) {
    *error_code = 0;

    return ERR_NONE;
}

combo_dev_attr_t LVDS_4lane_SENSOR_IMX290_12BIT_1080_NOWDR_ATTR =
{
    .devno         = 0,
    .input_mode    = INPUT_MODE_LVDS,      // input mode 
        .lvds_attr = {
            .img_size         = {1920, 1080},   // width x height
            .raw_data_type    = RAW_DATA_12BIT,
            .wdr_mode         = HI_WDR_MODE_NONE,
            .sync_mode        = LVDS_SYNC_MODE_SAV,
            .vsync_type       = {LVDS_VSYNC_NORMAL, 0, 0},
            .fid_type         = {LVDS_FID_NONE, HI_TRUE},
            .data_endian      = LVDS_ENDIAN_BIG,
            .sync_code_endian = LVDS_ENDIAN_BIG,
            .lane_id = {0, 1, 2, 3}, //rggb
            //.lane_id = {1, 0, 3, 2}, 
            //.lane_id = {1, 0, 3, 2}, //grbg
            //.lane_id = {1, 0, 2, 3},
            // .lane_id = {3, 0, 2, 1},
            // .lane_id = {3, 2, 0, 1},
            // .lane_id = {2, 3, 1, 0},
            .sync_code = {
                {
                    {0xab0, 0xb60, 0x800, 0x9d0},      // lane 0
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0}},
                {
                    {0xab0, 0xb60, 0x800, 0x9d0},      // lane 1
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0}},
                {
                    {0xab0, 0xb60, 0x800, 0x9d0},      // lane2
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0}},
                {
                    {0xab0, 0xb60, 0x800, 0x9d0},      // lane3
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0}
                }
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
    {0, 0, 1920, 1080}
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
                mipi: unsafe.Pointer(&C.LVDS_4lane_SENSOR_IMX290_12BIT_1080_NOWDR_ATTR),
                viDev: unsafe.Pointer(&C.DEV_ATTR_LVDS_BASE),
                clock: 37.125,
            },
        },
        control: cmosControl {
            bus: I2C,
            busNum: 0,
        },
        data: LVDS,
    }
)

