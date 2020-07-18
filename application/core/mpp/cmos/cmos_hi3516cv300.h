#pragma once

combo_dev_attr_t imx290_lvds_mode_0 =
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
            .lane_id = {0, 1, 2, 3},  //TODO
            //.lane_id = {1, 0, 3, 2},
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

