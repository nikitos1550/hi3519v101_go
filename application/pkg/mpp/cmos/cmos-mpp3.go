//+build hi3516cv300 hi3516av200

package cmos

/*
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



struct hi3516av200_cmos hi3516av200_cmoses[] = {
    {1, "imx274", "sony imx274 lvds 6lanes",
        3840, 2160, 30,
        &LVDS_6lane_SENSOR_IMX274_12BIT_8M_NOWDR_ATTR,
        &DEV_ATTR_LVDS_BASE,
        BAYER_RGGB,
        &stSnsImx274Obj},
    {0, NULL, NULL, 0, 0, 0, NULL, NULL, BAYER_BUTT}
};

*/
//import "C"
