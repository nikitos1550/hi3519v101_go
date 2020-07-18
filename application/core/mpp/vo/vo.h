#pragma once

#include "../include/mpp.h"
#include "../errmpp/errmpp.h"
#include "../../logger/logger.h"

#include <string.h>

int mpp_vo_init(error_in *err);

int mpp_vo_bind_vpss_test(error_in *err);
