//+build arm
//+build hi3516av100

package mipi

/*
#include "../include/mpp.h"
#include "../../logger/logger.h"

#include <string.h>
#include <fcntl.h>
#include <sys/ioctl.h>
#include <unistd.h>

#define ERR_NONE    0
#define ERR_GENERAL 1

typedef struct hi3516av100_mipi_init_in_struct {
    void *mipi;
} hi3516av100_mipi_init_in;

static int hi3516av100_mipi_init(int *error_code, hi3516av100_mipi_init_in *in) {
    *error_code = 0;

    int fd;

    fd = open("/dev/hi_mipi", O_RDWR);
    if (fd < 0) {
        GO_LOG_MIPI(LOGGER_ERROR, "open /dev/hi_mipi")
        *error_code = fd;
        return ERR_GENERAL;
    }

    combo_dev_attr_t stcomboDevAttr;

    memcpy(&stcomboDevAttr, in->mipi, sizeof(combo_dev_attr_t));

    *error_code = ioctl(fd, HI_MIPI_SET_DEV_ATTR, &stcomboDevAttr);
    if (*error_code != 0) {
        GO_LOG_MIPI(LOGGER_ERROR, "ioctl HI_MIPI_SET_DEV_ATTR")        
        close(fd);
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
    var errorCode C.int
    var in C.hi3516av100_mipi_init_in

    in.mipi = mipi

    err := C.hi3516av100_mipi_init(&errorCode, &in)
    if err != C.ERR_NONE {
        return errors.New("MIPI error TODO")
    }

    return nil
}

