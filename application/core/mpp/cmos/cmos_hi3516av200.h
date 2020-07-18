#pragma once


combo_dev_attr_t imx274_mode_0 =
{
    .devno = 0,  
    // input mode
    .input_mode = INPUT_MODE_LVDS,
    .phy_clk_share = PHY_CLK_SHARE_PHY0,
    .img_rect = {12, 40, 3840, 2160},

    .lvds_attr = 
    {
        .raw_data_type    = RAW_DATA_12BIT,    
        .wdr_mode         = HI_WDR_MODE_NONE,  
        .sync_mode        = LVDS_SYNC_MODE_SAV,
        .vsync_type       = {LVDS_VSYNC_NORMAL, 0, 0},
        .fid_type         = {LVDS_FID_NONE, HI_FALSE},
        .data_endian      = LVDS_ENDIAN_BIG,
        .sync_code_endian = LVDS_ENDIAN_BIG,
        .lane_id = {-1, 0, 1, -1, 2, 3, -1, 4, 5, -1, -1, -1},
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
                {0xab0, 0xb60, 0x800, 0x9d0}}
        }
    }
};

combo_dev_attr_t imx274_mode_1 =
{
    .devno = 0,
    // input mode
    .input_mode = INPUT_MODE_LVDS,
    .phy_clk_share = PHY_CLK_SHARE_PHY0,
    .img_rect = {12, 40, 3840, 2160},

    .lvds_attr = 
    {
        .raw_data_type    = RAW_DATA_10BIT,
        .wdr_mode         = HI_WDR_MODE_DOL_2F,
        .sync_mode        = LVDS_SYNC_MODE_SAV,
        .vsync_type       = {LVDS_VSYNC_NORMAL, 0, 0},
        .fid_type         = {LVDS_FID_IN_SAV, HI_TRUE},
        .data_endian      = LVDS_ENDIAN_BIG, 
        .sync_code_endian = LVDS_ENDIAN_BIG,   
        .lane_id = {0, 1, 2, 3, 4, 5, 6, 7, 8, 9, -1, -1},
        .sync_code = 
        {
            {{0x2ac,0x2d8,0x201,0x275},      // lane 0
                {0x2ac,0x2d8,0x202,0x276},  
                {0x2ac,0x2d8,0x201,0x275}, 
                {0x2ac,0x2d8,0x202,0x276}},
         
            {{0x2ac,0x2d8,0x201,0x275},      // lane 1   
                {0x2ac,0x2d8,0x202,0x276},    
                {0x2ac,0x2d8,0x201,0x275},    
                {0x2ac,0x2d8,0x202,0x276}},   

            {{0x2ac,0x2d8,0x201,0x275},      // lane 2   
                {0x2ac,0x2d8,0x202,0x276},    
                {0x2ac,0x2d8,0x201,0x275},    
                {0x2ac,0x2d8,0x202,0x276}},   
                                                          
            {{0x2ac,0x2d8,0x201,0x275},      // lane 3  
                {0x2ac,0x2d8,0x202,0x276},    
                {0x2ac,0x2d8,0x201,0x275},    
                {0x2ac,0x2d8,0x202,0x276}},   

            {{0x2ac,0x2d8,0x201,0x275},      // lane 4  
                {0x2ac,0x2d8,0x202,0x276},    
                {0x2ac,0x2d8,0x201,0x275},    
                {0x2ac,0x2d8,0x202,0x276}},   
                                                        
            {{0x2ac,0x2d8,0x201,0x275},      // lane 5
                {0x2ac,0x2d8,0x202,0x276},    
                {0x2ac,0x2d8,0x201,0x275},    
                {0x2ac,0x2d8,0x202,0x276}},
                                                        
             {{0x2ac,0x2d8,0x201,0x275},      // lane 6
                {0x2ac,0x2d8,0x202,0x276},    
                {0x2ac,0x2d8,0x201,0x275},   
                {0x2ac,0x2d8,0x202,0x276}},

             {{0x2ac,0x2d8,0x201,0x275},      // lane 7
                {0x2ac,0x2d8,0x202,0x276}, 
                {0x2ac,0x2d8,0x201,0x275}, 
                {0x2ac,0x2d8,0x202,0x276}},

             {{0x2ac,0x2d8,0x201,0x275},      // lane 8
                {0x2ac,0x2d8,0x202,0x276}, 
                {0x2ac,0x2d8,0x201,0x275}, 
                {0x2ac,0x2d8,0x202,0x276}},

             {{0x2ac,0x2d8,0x201,0x275},      // lane 9
                {0x2ac,0x2d8,0x202,0x276}, 
                {0x2ac,0x2d8,0x201,0x275}, 
                {0x2ac,0x2d8,0x202,0x276}},
        }
    }
};

combo_dev_attr_t imx226_mode_0 =
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

combo_dev_attr_t os08a10_mode_0 =
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

combo_dev_attr_t os08a10_mode_1 =
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

