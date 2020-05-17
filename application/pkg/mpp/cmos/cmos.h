#pragma once

#include "../include/mpp.h"
#include <string.h>

#define ERR_NONE    0
#define ERR_GENERAL 1

int mpp_cmos_init(int *error_code, unsigned char cmos); //now only for hi3516cv100
//int mpp_cmos_init(int *error_code);
