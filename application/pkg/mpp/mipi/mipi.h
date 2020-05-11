#pragma once

#include "../include/mpp.h"
#include "../errmpp/errmpp.h"
#include "../../logger/logger.h"

#include <string.h>
#include <fcntl.h>
#include <sys/ioctl.h>
#include <unistd.h>

typedef struct mpp_mipi_init_in_struct {
    //void *mipi;

    unsigned int mipi_crop_x0;
    unsigned int mipi_crop_y0;
    unsigned int mipi_crop_width;
    unsigned int mipi_crop_height;

    unsigned int width;
    unsigned int height;

    unsigned int data_type;
    unsigned int pixel_bitness;

    void *mipi_lvds_attr;
    void *mipi_mipi_attr;

} mpp_mipi_init_in;

int mpp_mipi_init(error_in *err, mpp_mipi_init_in *in);
