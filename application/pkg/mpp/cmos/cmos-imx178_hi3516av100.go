//+build arm
//+build hi3516av100
//+build imx178

package cmos

/*
#include "../include/mpp_v2.h"

#include "cmos.h"

#include <string.h>

int mpp_cmos_init(int *error_code) {
	*error_code = 0;

    return ERR_NONE;
}

combo_dev_attr_t LVDS_4lane_SENSOR_IMX178_12BIT_5M_NOWDR_ATTR___ =
{
    // input mode
    .input_mode = INPUT_MODE_LVDS,
    {
        .lvds_attr = {
            .img_size = {2592, 1944},
            HI_WDR_MODE_NONE,  
            LVDS_SYNC_MODE_SAV,
            RAW_DATA_12BIT, 
            LVDS_ENDIAN_BIG,
            LVDS_ENDIAN_BIG,
            .lane_id = {0, 1, 2, 3, -1, -1, -1, -1},
            .sync_code = {
                    {{0xab0, 0xb60, 0x800, 0x9d0},  //1
                    {0xab0, 0xb60, 0x800, 0x9d0}, 
                    {0xab0, 0xb60, 0x800, 0x9d0}, 
                    {0xab0, 0xb60, 0x800, 0x9d0}},

                    {{0xab0, 0xb60, 0x800, 0x9d0},   //2
                    {0xab0, 0xb60, 0x800, 0x9d0}, 
                    {0xab0, 0xb60, 0x800, 0x9d0}, 
                    {0xab0, 0xb60, 0x800, 0x9d0}},

                    {{0xab0, 0xb60, 0x800, 0x9d0},   //3
                    {0xab0, 0xb60, 0x800, 0x9d0}, 
                    {0xab0, 0xb60, 0x800, 0x9d0}, 
                    {0xab0, 0xb60, 0x800, 0x9d0}},

                    {{0xab0, 0xb60, 0x800, 0x9d0},  //4
                    {0xab0, 0xb60, 0x800, 0x9d0}, 
                    {0xab0, 0xb60, 0x800, 0x9d0}, 
                    {0xab0, 0xb60, 0x800, 0x9d0}},

                    {{0xab0, 0xb60, 0x800, 0x9d0},   //5
                    {0xab0, 0xb60, 0x800, 0x9d0}, 
                    {0xab0, 0xb60, 0x800, 0x9d0}, 
                    {0xab0, 0xb60, 0x800, 0x9d0}},

                    {{0xab0, 0xb60, 0x800, 0x9d0},   //6
                    {0xab0, 0xb60, 0x800, 0x9d0}, 
                    {0xab0, 0xb60, 0x800, 0x9d0}, 
                    {0xab0, 0xb60, 0x800, 0x9d0}},

                    {{0xab0, 0xb60, 0x800, 0x9d0},   //7
                    {0xab0, 0xb60, 0x800, 0x9d0}, 
                    {0xab0, 0xb60, 0x800, 0x9d0}, 
                    {0xab0, 0xb60, 0x800, 0x9d0}},

                    {{0xab0, 0xb60, 0x800, 0x9d0},   //8
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0}}
                    }
        }
    }
};

VI_DEV_ATTR_S DEV_ATTR_LVDS_BASE__ =
{
    // interface mode
    VI_MODE_LVDS,
    // multiplex mode
    VI_WORK_MODE_1Multiplex,
    //r_mask    g_mask    b_mask
    {0xFFF00000,    0x0},
    //progessive or interleaving
    VI_SCAN_PROGRESSIVE,
    //AdChnId
    {-1, -1, -1, -1},
    //enDataSeq, only support yuv
    VI_INPUT_DATA_YUYV,

    // synchronization information
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
		model: "imx178",
		modes: []cmosMode {
			cmosMode {
				width: 2592,
				height: 1944,
				fps: 30,
				mipi: unsafe.Pointer(&C.LVDS_4lane_SENSOR_IMX178_12BIT_5M_NOWDR_ATTR___),
			},
		},
		viDev: unsafe.Pointer(&C.DEV_ATTR_LVDS_BASE__),
		clock: 25,
                control: cmosControl {
                        bus: I2C,
                        busNum: 0,
                },		
	}
)
