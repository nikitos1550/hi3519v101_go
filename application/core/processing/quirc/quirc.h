#pragma once

#include "../../mpp/include/mpp.h"
#include "../../mpp/errmpp/errmpp.h"
#include "../../logger/logger.h"

#include "../../../vendors/quirc/quirc/lib/quirc.h"

#include <string.h>

int quirc_quirc_init(int w, int y);
int quirc_quirc_deinit();
int quirc_process(error_in *err, void *frame);
