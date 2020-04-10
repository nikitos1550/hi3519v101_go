//+build arm
//+build hi3516av100

package mipi

/*
#include "../include/mpp_v2.h"

#include <string.h>
#include <fcntl.h>
#include <sys/ioctl.h>
#include <unistd.h>


combo_dev_attr_t LVDS_4lane_SENSOR_IMX178_12BIT_5M_NOWDR_ATTR =
{
    // input mode
    .input_mode = INPUT_MODE_LVDS,
    {
        .lvds_attr = {
            .img_size = {2592, 1944},
            HI_WDR_MODE_NONE,
            LVDS_SYNC_MODE_SAV,
            RAW_DATA_12BIT,
            LVDS_ENDIAN_BIG,
            LVDS_ENDIAN_BIG,
            .lane_id = {0, 1, 2, 3, -1, -1, -1, -1},
            .sync_code = {
                    {{0xab0, 0xb60, 0x800, 0x9d0},  //1
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0}},

                    {{0xab0, 0xb60, 0x800, 0x9d0},   //2
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0}},

                    {{0xab0, 0xb60, 0x800, 0x9d0},   //3
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0}},

                    {{0xab0, 0xb60, 0x800, 0x9d0},  //4
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0}},

                    {{0xab0, 0xb60, 0x800, 0x9d0},   //5
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0}},

                    {{0xab0, 0xb60, 0x800, 0x9d0},   //6
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0}},

                    {{0xab0, 0xb60, 0x800, 0x9d0},   //7
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0}},

                    {{0xab0, 0xb60, 0x800, 0x9d0},   //8
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0},
                    {0xab0, 0xb60, 0x800, 0x9d0}}
		    }
        }
    }
};


#define ERR_NONE    0
#define ERR_GENERAL 1

int mpp2_mipi_init(int *error_code) {
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

    pstcomboDevAttr = &LVDS_4lane_SENSOR_IMX178_12BIT_5M_NOWDR_ATTR;

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
)

func Init() {
    var errorCode C.int

    switch err := C.mpp2_mipi_init(&errorCode); err {
    case C.ERR_NONE:
        logger.Log.Debug().
                Msg("C.mpp2_mipi_init() ok")
    default:
        logger.Log.Fatal().
                Int("error", int(err)).
                Msg("Unexpected return of C.mpp2_mipi_init()")
    }

}
