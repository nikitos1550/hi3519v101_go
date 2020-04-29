#pragma once

#include "../include/mpp.h"
#include "../errmpp/errmpp.h"
#include "../../logger/logger.h"

#include <string.h>
#include <pthread.h>

typedef struct mpp_isp_init_in_struct {
    unsigned int width;
    unsigned int height;
    unsigned int fps;
    unsigned int bayer;
    unsigned int wdr;
} mpp_isp_init_in;

int mpp_isp_init(error_in *err, mpp_isp_init_in *in);
