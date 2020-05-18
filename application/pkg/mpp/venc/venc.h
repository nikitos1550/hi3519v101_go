#pragma once

#include "../include/mpp.h"
#include "../errmpp/errmpp.h"
#include "../../logger/logger.h"

typedef struct mpp_venc_create_in_struct {              
    unsigned int venc_id;
    unsigned int width;
    unsigned int height;
    unsigned int bitrate;

    unsigned int vpss_fps;
    unsigned int fps;
} mpp_venc_create_in;              


