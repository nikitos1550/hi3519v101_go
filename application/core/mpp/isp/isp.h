#pragma once

#include "../include/mpp.h"
#include "../errmpp/errmpp.h"
#include "../../logger/logger.h"

#include <string.h>
#include <pthread.h>

int pthread_setname_np(pthread_t thread, const char *name);

typedef struct mpp_isp_init_in_struct {
    unsigned int isp_crop_x0;
    unsigned int isp_crop_y0;
    unsigned int isp_crop_width;
    unsigned int isp_crop_height;

    unsigned int width;
    unsigned int height;

    unsigned int fps;
    unsigned int bayer;
    unsigned int wdr;
} mpp_isp_init_in;

int mpp_isp_init(error_in *err, mpp_isp_init_in *in);

void* mpp_isp_thread(HI_VOID *param);
