#pragma once

combo_dev_attr_t imx290_lvds_mode_0_1 = 
{
    /* input mode */
    .input_mode = INPUT_MODE_LVDS,
    {
        .lvds_attr = {
        .img_size = {1920, 1080},
        HI_WDR_MODE_NONE,
        LVDS_SYNC_MODE_SAV,
        RAW_DATA_12BIT,
        LVDS_ENDIAN_BIG,
        LVDS_ENDIAN_BIG,
        //.lane_id = {0, 1, 2, 3, -1, -1, -1, -1},
        .lane_id = {1, 0, 3, 2, -1, -1, -1, -1},
        .sync_code = { 
                {{0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}},
                
                {{0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}},

                {{0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}},
                
                {{0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}},
                
                {{0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}},
                    
                {{0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}},

                {{0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}},
                
                {{0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}, 
                {0xab0, 0xb60, 0x800, 0x9d0}} 
            }
        }
    }
};

combo_dev_attr_t imx290_lvds_mode_2 = 
{
    /* input mode */
    .input_mode = INPUT_MODE_LVDS,
    {
        .lvds_attr = {
        .img_size = {1920, 1080},
        HI_WDR_MODE_DOL_2F,
        LVDS_SYNC_MODE_SAV,
        RAW_DATA_10BIT, 
        LVDS_ENDIAN_BIG,
        LVDS_ENDIAN_BIG,
        //.lane_id = {0, 1, 2, 3, -1, -1, -1, -1},
        .lane_id = {1, 0, 3, 2, -1, -1, -1, -1},
        .sync_code = {                         
                {{0x0AC, 0x0D8, 0x001, 0x075},
                {0x0AC, 0x0D8, 0x002, 0x076}, 
                {0x0AC, 0x0D8, 0x100, 0x174}, 
                {0x0AC, 0x0D8, 0x100, 0x174}},
                                               
                {{0x0AC, 0x0D8, 0x001, 0x075},
                {0x0AC, 0x0D8, 0x002, 0x076}, 
                {0x0AC, 0x0D8, 0x100, 0x174}, 
                {0x0AC, 0x0D8, 0x100, 0x174}},
                                               
                {{0x0AC, 0x0D8, 0x001, 0x075},
                {0x0AC, 0x0D8, 0x002, 0x076}, 
                {0x0AC, 0x0D8, 0x100, 0x174}, 
                {0x0AC, 0x0D8, 0x100, 0x174}},
                                               
                {{0x0AC, 0x0D8, 0x001, 0x075},
                {0x0AC, 0x0D8, 0x002, 0x076}, 
                {0x0AC, 0x0D8, 0x100, 0x174}, 
                {0x0AC, 0x0D8, 0x100, 0x174}},     
              /*only 4 ch ,below may be not use*/
                {{0x0AC, 0x0D8, 0x001, 0x075},
                {0x0AC, 0x0D8, 0x002, 0x076}, 
                {0x0AC, 0x0D8, 0x001, 0x075}, 
                {0x0AC, 0x0D8, 0x002, 0x076}},
                                               
                {{0x0AC, 0x0D8, 0x001, 0x075},
                {0x0AC, 0x0D8, 0x002, 0x076}, 
                {0x0AC, 0x0D8, 0x001, 0x075}, 
                {0x0AC, 0x0D8, 0x002, 0x076}},
                                               
                {{0x0AC, 0x0D8, 0x001, 0x075},
                {0x0AC, 0x0D8, 0x002, 0x076}, 
                {0x0AC, 0x0D8, 0x001, 0x075}, 
                {0x0AC, 0x0D8, 0x002, 0x076}},
                                               
                {{0x0AC, 0x0D8, 0x001, 0x075},
                {0x0AC, 0x0D8, 0x002, 0x076}, 
                {0x0AC, 0x0D8, 0x001, 0x075},  
                {0x0AC, 0x0D8, 0x002, 0x076}},
            }
        }
    }
};

combo_dev_attr_t imx178_lvds_mode_0 =
{
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

combo_dev_attr_t imx178_lvds_mode_1 =
{
    /* input mode */
    .input_mode = INPUT_MODE_LVDS,
    {
        .lvds_attr = {
            .img_size = {1920, 1080},
            HI_WDR_MODE_NONE,  
            LVDS_SYNC_MODE_SAV,
            RAW_DATA_12BIT, 
            LVDS_ENDIAN_BIG,
            LVDS_ENDIAN_BIG,
            .lane_id = {0, 1, 2, 3, -1, -1, -1, -1},
            .sync_code = {
                {   {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0} 
                },
                    
                {   {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0} 
                },
                    
                {   {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0} 
                },
                    
                {   {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0} 
                },
                    
                {   {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0} 
                },
                  
                {   {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0} 
                },
                  
                {   {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0} 
                },
                  
                {   {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0} 
                }
            }    
        }    
    }    
}; 

combo_dev_attr_t ov4689_mipi_mode_0 =
{
    .input_mode = INPUT_MODE_MIPI,
    {

        .mipi_attr =
        {
            RAW_DATA_12BIT,
            {0, 1, 2, 3, -1, -1, -1, -1}
        }
    }
};

