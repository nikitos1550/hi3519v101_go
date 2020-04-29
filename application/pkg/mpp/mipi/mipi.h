#pragma once

#include "../include/mpp.h"
#include "../errmpp/errmpp.h"
#include "../../logger/logger.h"

#include <string.h>
#include <fcntl.h>
#include <sys/ioctl.h>
#include <unistd.h>

typedef struct mpp_mipi_init_in_struct {
    void *mipi;
} mpp_mipi_init_in;

int mpp_mipi_init(error_in *err, mpp_mipi_init_in *in);

