//+build arm
//+build hi3516av200

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
    fd = open("/dev/hi_mipi", O_RDWR);
    if (fd < 0) return ERR_GENERAL;

    combo_dev_attr_t stcomboDevAttr;

    memcpy(&stcomboDevAttr, mipi, sizeof(combo_dev_attr_t));

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
	    logger.Log.Debug().
		    Msg("C.mpp3_mipi_init() ok")
    default:
	    logger.Log.Fatal().
		    Int("error", int(err)).
		    Msg("Unexpected return of C.mpp3_mipi_init()")
    }
}

