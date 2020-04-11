//+build arm
//+build hi3516av200

package mipi

/*
#include "../include/mpp_v3.h"

#include <string.h>
#include <fcntl.h>
#include <sys/ioctl.h>
#include <unistd.h>

#ifdef HI3516AV200
combo_dev_attr_t LVDS_6lane_SENSOR_IMX274_12BIT_8M_NOWDR_ATTR =
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
#endif
#ifdef HI3516CV300
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

combo_dev_attr_t MIPI_CMOS323_ATTR = 
{
    // input mode
    .input_mode = INPUT_MODE_CMOS,
    {
        
    }
};
#endif

#define ERR_NONE    0
#define ERR_GENERAL 1

int mpp3_mipi_init(int *error_code, void *mipi) {
    *error_code = 0;
    
    int fd;
    fd = open("/dev/hi_mipi", O_RDWR);
    if (fd < 0) return ERR_GENERAL;

    combo_dev_attr_t stcomboDevAttr;

    //TODO
    #ifdef HI3516AV200
    //memcpy(&stcomboDevAttr, &LVDS_6lane_SENSOR_IMX274_12BIT_8M_NOWDR_ATTR, sizeof(combo_dev_attr_t));
    memcpy(&stcomboDevAttr, mipi, sizeof(combo_dev_attr_t));
    #endif
    #ifdef HI3516CV300
    memcpy(&stcomboDevAttr, &MIPI_CMOS323_ATTR, sizeof(combo_dev_attr_t));
    #endif

    stcomboDevAttr.devno = 0;

    if(ioctl(fd, HI_MIPI_RESET_MIPI, &stcomboDevAttr.devno)) {
        close(fd);
        return ERR_GENERAL;
    }

    if(ioctl(fd, HI_MIPI_RESET_SENSOR, &stcomboDevAttr.devno)) {
        close(fd);
        return ERR_GENERAL;
    }

    if (ioctl(fd, HI_MIPI_SET_DEV_ATTR, &stcomboDevAttr)) {
        close(fd);
        return ERR_GENERAL;
    }

    if(ioctl(fd, HI_MIPI_UNRESET_MIPI, &stcomboDevAttr.devno)) {
        close(fd);
        return ERR_GENERAL;
    }

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
    //"log"
	"application/pkg/logger"

	"application/pkg/mpp/cmos"
)

func Init() {
    var errorCode C.int

    switch err := C.mpp3_mipi_init(&errorCode, cmos.Mipi() ); err {
    case C.ERR_NONE:
        //log.Println("C.mpp3_mipi_init() ok")
	logger.Log.Debug().
		Msg("C.mpp3_mipi_init() ok")
    default:
        //log.Fatal("Unexpected return ", err , " of C.mpp3_mipi_init()")
	logger.Log.Fatal().
		Int("error", int(err)).
		Msg("Unexpected return of C.mpp3_mipi_init()")
    }
}

