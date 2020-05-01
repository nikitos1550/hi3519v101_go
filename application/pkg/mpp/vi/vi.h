#pragma once

#include "../include/mpp.h"
#include "../errmpp/errmpp.h"
#include "../../logger/logger.h"

#include <string.h>

typedef struct mpp_vi_init_in_struct {
    void *videv;

    unsigned int cmos_width;
    unsigned int cmos_height;

    unsigned int x0;
    unsigned int y0;
    unsigned int width;
    unsigned int height;
    unsigned int cmos_fps;
    unsigned int fps;

    unsigned int wdr;

    unsigned char mirror;
    unsigned char flip;

    unsigned char ldc;
    int ldc_offset_x;
    int ldc_offset_y;
    int ldc_k;
} mpp_vi_init_in;

int mpp_vi_init(error_in *err, mpp_vi_init_in * in);

