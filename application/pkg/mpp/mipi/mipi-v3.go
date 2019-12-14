//+build hi3516cv300 hi3516av200

package mipi

/*
#include "../include/hi3516av200_mpp.h"
#include <string.h>
#include <fcntl.h>
#include <sys/ioctl.h>
#include <unistd.h>

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


#define ERR_NONE    0
#define ERR_GENERAL 1

int mpp3_mipi_init(int *error_code) {
    *error_code = 0;
    
    int fd;
    fd = open("/dev/hi_mipi", O_RDWR);
    if (fd < 0) return ERR_GENERAL;

    combo_dev_attr_t stcomboDevAttr;

    //TODO
    memcpy(&stcomboDevAttr, &LVDS_6lane_SENSOR_IMX274_12BIT_8M_NOWDR_ATTR, sizeof(combo_dev_attr_t));
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
    "log"
)

func Init() {
    var errorCode C.int

    switch err := C.mpp3_mipi_init(&errorCode); err {
    case C.ERR_NONE:
        log.Println("C.mpp3_mipi_init() ok")
    default:
        log.Fatal("Unexpected return ", err , " of C.mpp3_mipi_init()")
    }
}

