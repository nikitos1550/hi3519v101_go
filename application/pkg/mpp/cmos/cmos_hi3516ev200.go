//+build arm
//+build hi3516ev200

package cmos

/*
#include "cmos.h"
#include "cmos_hi3516ev200.h"

int mpp_cmos_init(int *error_code, unsigned char cmos) {
    *error_code = 0;

    ALG_LIB_S stAeLib;
    ALG_LIB_S stAwbLib;

    stAeLib.s32Id = 0;
    stAwbLib.s32Id = 0;
    strncpy(stAeLib.acLibName,  HI_AE_LIB_NAME,     sizeof(HI_AE_LIB_NAME));
    strncpy(stAwbLib.acLibName, HI_AWB_LIB_NAME,    sizeof(HI_AWB_LIB_NAME));

//    if (cmos == 0) {
//        ISP_SNS_OBJ_S *cmos = &stSnsImx274Obj;
//
//        if (cmos->pfnRegisterCallback != HI_NULL) {
//            *error_code = cmos->pfnRegisterCallback(0, &stAeLib, &stAwbLib);
//            if (*error_code != HI_SUCCESS) return ERR_GENERAL;
//        } else {
//            return ERR_GENERAL;
//        }
//
//        return ERR_NONE;
//    }

    *error_code = 999;
    return ERR_GENERAL;
}
*/
import "C"

import (
    _"unsafe"
)

var (
    cmosItems = []cmos {}
)
