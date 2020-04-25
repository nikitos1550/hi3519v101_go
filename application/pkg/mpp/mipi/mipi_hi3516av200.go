//+build arm
//+build hi3516av200

package mipi

/*
#include "../include/mpp.h"
#include "../errmpp/errmpp.h"
#include "../../logger/logger.h"

#include <string.h>
#include <fcntl.h>
#include <sys/ioctl.h>
#include <unistd.h>

typedef struct hi3516av200_mipi_init_in_struct {
    void *mipi;
} hi3516av200_mipi_init_in;

static int hi3516av200_mipi_init(error_in *err, hi3516av200_mipi_init_in * in) {
    int general_error_code = 0;
    
    int fd;

    fd = open("/dev/hi_mipi", O_RDWR);
    if (fd < 0) {
        GO_LOG_MIPI(LOGGER_ERROR, "open /dev/hi_mipi");
        err->general = fd;    
        return ERR_GENERAL;
    }

    combo_dev_attr_t stcomboDevAttr;

    memcpy(&stcomboDevAttr, in->mipi, sizeof(combo_dev_attr_t));
    stcomboDevAttr.devno = 0;

    general_error_code = ioctl(fd, HI_MIPI_RESET_MIPI, &stcomboDevAttr.devno);
    if (general_error_code != 0){
        GO_LOG_MIPI(LOGGER_ERROR, "ioctl HI_MIPI_RESET_MIPI");
        close(fd);
        err->general = general_error_code;
        return ERR_GENERAL;
    }

    general_error_code = ioctl(fd, HI_MIPI_RESET_SENSOR, &stcomboDevAttr.devno); 
    if (general_error_code != 0) {
        GO_LOG_MIPI(LOGGER_ERROR, "ioctl HI_MIPI_RESET_SENSOR");
        close(fd);
        err->general = general_error_code;
        return ERR_GENERAL;
    }

    general_error_code = ioctl(fd, HI_MIPI_SET_DEV_ATTR, &stcomboDevAttr);
    if (general_error_code != 0) {
		GO_LOG_MIPI(LOGGER_ERROR, "ioctl HI_MIPI_SET_DEV_ATTR");
        close(fd);
        err->general = general_error_code;
        return ERR_GENERAL;
    }

    general_error_code = ioctl(fd, HI_MIPI_UNRESET_MIPI, &stcomboDevAttr.devno);
    if (general_error_code != 0) {
		GO_LOG_MIPI(LOGGER_ERROR, "ioctl HI_MIPI_UNRESET_MIPI");
        close(fd);
        err->general = general_error_code;
        return ERR_GENERAL;
    }

    general_error_code = ioctl(fd, HI_MIPI_UNRESET_SENSOR, &stcomboDevAttr.devno); 
    if (general_error_code != 0) {
		GO_LOG_MIPI(LOGGER_ERROR, "ioctl HI_MIPI_UNRESET_SENSOR");
        close(fd);
        err->general = general_error_code;
        return ERR_GENERAL;
    }

    close(fd);

    return ERR_NONE;
}
*/
import "C"

import (
    "errors"
)

func initFamily() error {
    var inErr C.error_in
    var in C.hi3516av200_mipi_init_in

    in.mipi = mipi

    err := C.hi3516av200_mipi_init(&inErr, &in)
    if err != C.ERR_NONE {
        return errors.New("MIPI error TODO")
    }

    return nil
}

