//+build arm
//+build hi3516av100

package mipi

/*
#include "../include/mpp_v2.h"

#include <string.h>
#include <fcntl.h>
#include <sys/ioctl.h>
#include <unistd.h>


#define ERR_NONE    0
#define ERR_GENERAL 1

int mpp2_mipi_init(int *error_code, void *mipi) {
    *error_code = 0;


    int fd;
    combo_dev_attr_t *pstcomboDevAttr;

    // mipi reset unrest
    fd = open("/dev/hi_mipi", O_RDWR);
    if (fd < 0)
    {
        //printf("warning: open hi_mipi dev failed\n");
        return ERR_GENERAL;
    }

    pstcomboDevAttr = mipi; //&LVDS_4lane_SENSOR_IMX178_12BIT_5M_NOWDR_ATTR;

    if (ioctl(fd, HI_MIPI_SET_DEV_ATTR, pstcomboDevAttr))
    {
        //printf("set mipi attr failed\n");
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

    switch err := C.mpp2_mipi_init(&errorCode, cmos.Mipi()); err {
    case C.ERR_NONE:
        logger.Log.Debug().
                Msg("C.mpp2_mipi_init() ok")
    default:
        logger.Log.Fatal().
                Int("error", int(err)).
                Msg("Unexpected return of C.mpp2_mipi_init()")
    }

}
