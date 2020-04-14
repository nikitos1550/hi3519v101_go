//+build arm
//+build hi3516cv300

package mipi
/*
#include "../include/mpp_v3.h"

#include <string.h>
#include <fcntl.h>
#include <sys/ioctl.h>
#include <unistd.h>

#define ERR_NONE    0
#define ERR_GENERAL 1

int mpp3_mipi_init(int *error_code, void *mipi) {
    *error_code = 0;

    int fd;
    combo_dev_attr_t *pstcomboDevAttr, stcomboDevAttr;

    // mipi reset unrest
    fd = open("/dev/hi_mipi", O_RDWR);
    if (fd < 0) {
        return ERR_GENERAL;
    }
	pstcomboDevAttr = mipi;

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

        "application/pkg/mpp/cmos"
)

func Init() {
    var errorCode C.int

    switch err := C.mpp3_mipi_init(&errorCode, cmos.Mipi() ); err {
    case C.ERR_NONE:
        logger.Log.Debug().
                Msg("C.mpp3_mipi_init() ok")
    default:
        logger.Log.Fatal().
                Int("error", int(err)).
                Msg("Unexpected return of C.mpp3_mipi_init()")
    }

}

