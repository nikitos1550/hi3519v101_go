//+build arm
//+build hi3516cv500

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

typedef struct hi3516cv500_mipi_init_in_struct {
    void *mipi;
} hi3516cv500_mipi_init_in;

static int hi3516cv500_mipi_init(int *error_code, hi3516cv500_mipi_init_in *in) {
    *error_code = 0;

    int fd;
    
    fd = open( "/dev/hi_mipi", O_RDWR);
    if (fd < 0) {
         GO_LOG_MIPI(LOGGER_ERROR, "open /dev/hi_mipi")     
        *error_code = fd;
        return ERR_GENERAL;
    }

    lane_divide_mode_t lane_divide_mode = LANE_DIVIDE_MODE_0;

    *error_code = ioctl(fd, HI_MIPI_SET_HS_MODE, &lane_divide_mode);
    if (*error_code != HI_SUCCESS) {
        GO_LOG_MIPI(LOGGER_ERROR, "ioctl HI_MIPI_SET_HS_MODE")
        close(fd);  
        return ERR_GENERAL; 
    }

    combo_dev_t devno = 0;

    *error_code = ioctl(fd, HI_MIPI_ENABLE_MIPI_CLOCK, &devno);
    if (*error_code != HI_SUCCESS) {
        GO_LOG_MIPI(LOGGER_ERROR, "ioctl HI_MIPI_ENABLE_MIPI_CLOCK")
        close(fd);
        return ERR_GENERAL; 
    }

    *error_code = ioctl(fd, HI_MIPI_RESET_MIPI, &devno);
    if (*error_code != HI_SUCCESS) {
        GO_LOG_MIPI(LOGGER_ERROR, "ioctl HI_MIPI_RESET_MIPI")
        close(fd);
        return ERR_GENERAL; 
    }

    sns_clk_source_t SnsDev = 0;

    *error_code = ioctl(fd, HI_MIPI_ENABLE_SENSOR_CLOCK, &SnsDev);
    if (*error_code != HI_SUCCESS) {
        GO_LOG_MIPI(LOGGER_ERROR, "ioctl HI_MIPI_ENABLE_SENSOR_CLOCK")
        close(fd);
        return ERR_GENERAL; 
    }

    *error_code = ioctl(fd, HI_MIPI_RESET_SENSOR, &SnsDev);
    if (*error_code != HI_SUCCESS) {
        GO_LOG_MIPI(LOGGER_ERROR, "ioctl HI_MIPI_RESET_SENSOR")
		close(fd);
		return ERR_GENERAL; 
    }

    *error_code = ioctl(fd, HI_MIPI_SET_DEV_ATTR, in->mipi);
	if (*error_code != HI_SUCCESS) {
        GO_LOG_MIPI(LOGGER_ERROR, "ioctl HI_MIPI_SET_DEV_ATTR")
		close(fd);
		return ERR_GENERAL; 
	}

    *error_code = ioctl(fd, HI_MIPI_UNRESET_MIPI, &devno);
	if (*error_code != HI_SUCCESS) {
        GO_LOG_MIPI(LOGGER_ERROR, "ioctl HI_MIPI_UNRESET_MIPI")
		close(fd);
		return ERR_GENERAL; 
	}

    *error_code = ioctl(fd, HI_MIPI_UNRESET_SENSOR, &SnsDev);
	if (*error_code != HI_SUCCESS) {
        GO_LOG_MIPI(LOGGER_ERROR, "ioctl HI_MIPI_UNRESET_SENSOR")
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
    var in C.hi3516cv500_mipi_init_in

    in.mipi = mipi

    err := C.hi3516cv500_mipi_init(&errorCode, &in)
    if err != C.ERR_NONE {
        return errors.New("MIPI error TODO")
    }

    return nil
}

