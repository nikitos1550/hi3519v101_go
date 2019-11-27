//+build hi3516cv300 hi3516av200

package mipi

/*
#include "../include/hi3516av200_mpp.h"
#include <string.h>
#include <fcntl.h>
#include <sys/ioctl.h>
#include <unistd.h>


#define ERR_NONE    0
#define ERR_GENERAL 1

int mpp3_mipi_init(int *error_code) {
    *error_code = 0;
    
    int fd;
    fd = open("/dev/hi_mipi", O_RDWR);
    if (fd < 0) return ERR_GENERAL;

    combo_dev_attr_t stcomboDevAttr;

    //TODO
    //memcpy(&stcomboDevAttr, c->mipidev, sizeof(combo_dev_attr_t));
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

func Init() {
    var errorCode C.int
    C.mpp3_mipi_init(&errorCode)
}

