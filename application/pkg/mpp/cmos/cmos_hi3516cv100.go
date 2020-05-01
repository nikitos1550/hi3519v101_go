//+build arm
//+build hi3516cv100

package cmos

/*
#include "../include/mpp.h"   
#include "cmos.h"

int mpp_cmos_init(int *error_code) {
    *error_code = 0;

    *error_code = sensor_register_callback();
    if (*error_code != HI_SUCCESS) {
        if (*error_code != HI_SUCCESS) return ERR_GENERAL;
    }

    return ERR_NONE;
}
*/
import "C"
