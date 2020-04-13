//+build arm
//+build hi3516cv300

package mipi
/*
#include "../include/mpp_v3.h"

#include <string.h>
#include <fcntl.h>
#include <sys/ioctl.h>
#include <unistd.h>

combo_dev_attr_t MIPI_CMOS323_ATTR = 
{
    // input mode
    .input_mode = INPUT_MODE_CMOS,
    {
        
    }
};

combo_dev_attr_t MIPI_4lane_SENSOR_IMX290_12BIT_1080_NOWDR_ATTR = 
{
        .devno = 0,
        .input_mode = INPUT_MODE_MIPI, 
        {
                .mipi_attr = 
                {
                        RAW_DATA_12BIT,
                        HI_MIPI_WDR_MODE_NONE,
			{0, 1, 2, 3}
			//{0, 1, 3, 2}
                }
        }
};

combo_dev_attr_t MIPI_4lane_SENSOR_IMX290_10BIT_1080_2DOL1_ATTR = 
{
    .devno = 0,
    .input_mode = INPUT_MODE_MIPI,    

    .mipi_attr =    
    {
        .raw_data_type = RAW_DATA_10BIT,
        .wdr_mode = HI_MIPI_WDR_MODE_DOL,
        .lane_id = {0, 1, 2, 3}
    }
};

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
            .lane_id = {0, 1, 2, 3},
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

#define ERR_NONE    0
#define ERR_GENERAL 1


int mpp3_mipi_init(int *error_code) {
    *error_code = 0;

    int fd;
    combo_dev_attr_t *pstcomboDevAttr, stcomboDevAttr;

    // mipi reset unrest
    fd = open("/dev/hi_mipi", O_RDWR);
    if (fd < 0) {
        return ERR_GENERAL;
    }
		pstcomboDevAttr = &MIPI_CMOS323_ATTR;
    //pstcomboDevAttr = &MIPI_4lane_SENSOR_IMX290_12BIT_1080_NOWDR_ATTR;
	//pstcomboDevAttr = &LVDS_4lane_SENSOR_IMX290_12BIT_1080_NOWDR_ATTR;

    memcpy(&stcomboDevAttr, pstcomboDevAttr, sizeof(combo_dev_attr_t));

  // 1.reset mipi
    if(ioctl(fd, HI_MIPI_RESET_MIPI, &stcomboDevAttr.devno)) {
        close(fd);
        return ERR_GENERAL;
    }

    // 2.reset sensor
    if(ioctl(fd, HI_MIPI_RESET_SENSOR, &stcomboDevAttr.devno)) {
        close(fd);
        return ERR_GENERAL;    
	}

    if (ioctl(fd, HI_MIPI_SET_DEV_ATTR, pstcomboDevAttr)) {
        close(fd);
        return ERR_GENERAL;
    }

usleep(10000);

   // 4.unreset mipi 
    if(ioctl(fd, HI_MIPI_UNRESET_MIPI, &stcomboDevAttr.devno)) {
        close(fd);
        return ERR_GENERAL;
    }

    // 5.unreset sensor
    if(ioctl(fd, HI_MIPI_UNRESET_SENSOR, &stcomboDevAttr.devno)) {
        close(fd);
        return ERR_GENERAL;
    }


    close(fd);


    return ERR_NONE;
}

*/
import "C"

import (
        "application/pkg/logger"
)

func Init() {
    var errorCode C.int

    switch err := C.mpp3_mipi_init(&errorCode); err {
    case C.ERR_NONE:
        logger.Log.Debug().
                Msg("C.mpp3_mipi_init() ok")
    default:
        logger.Log.Fatal().
                Int("error", int(err)).
                Msg("Unexpected return of C.mpp3_mipi_init()")
    }

}

