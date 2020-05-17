#pragma once


combo_dev_attr_t imx307_mode_0 =
{
    .devno = 0,
    .input_mode = INPUT_MODE_MIPI, 
    .data_rate = MIPI_DATA_RATE_X1,  
    .img_rect = {0, 4, 1920, 1080},  
    {
        .mipi_attr =
        {
            DATA_TYPE_RAW_12BIT,  
            HI_MIPI_WDR_MODE_NONE,
            {0, 2, -1, -1}
        }
    }
};

combo_dev_attr_t imx307_mode_1 =
{
    .devno = 0,
    .input_mode = INPUT_MODE_MIPI,
    .data_rate = MIPI_DATA_RATE_X1,
    .img_rect = {0, 4, 1920, 1080},
    {
        .mipi_attr =
        {
            DATA_TYPE_RAW_10BIT,
            HI_MIPI_WDR_MODE_DOL, //HI_MIPI_WDR_MODE_VC,
            {0, 2, -1, -1}
        }
    }
};

combo_dev_attr_t imx335_mode_0 =
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

combo_dev_attr_t imx335_mode_1 =
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

combo_dev_attr_t imx335_mode_2 =
{
    .devno = 0,
    .input_mode = INPUT_MODE_MIPI, 
    .data_rate = MIPI_DATA_RATE_X1,  
    .img_rect = {0, 0, 2592, 1520},
    {
        .mipi_attr =
        {
            DATA_TYPE_RAW_12BIT,  
            HI_MIPI_WDR_MODE_NONE,
            {0, 1, 2, 3}
        }
    }
};

combo_dev_attr_t imx335_mode_3 =
{
    .devno = 0,
    .input_mode = INPUT_MODE_MIPI, 
    .data_rate = MIPI_DATA_RATE_X1,  
    .img_rect = {12, 14, 2592, 1520},

    {
        .mipi_attr =
        {
            DATA_TYPE_RAW_10BIT,
            HI_MIPI_WDR_MODE_VC,
            {0, 1, 2, 3}
        }
    }
};

