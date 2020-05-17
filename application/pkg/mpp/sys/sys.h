#pragma once

#include "../include/mpp.h"
#include "../errmpp/errmpp.h"
#include "../../logger/logger.h"

#include <string.h>

typedef struct mpp_sys_init_in_struct {
    unsigned int width; 
    unsigned int height;
    unsigned int cnt;    
} mpp_sys_init_in;

int mpp_sys_init(error_in *err, mpp_sys_init_in *in);
