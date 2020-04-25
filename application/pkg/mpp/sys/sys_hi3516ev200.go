//+build arm
//+build hi3516ev200

package sys

/*
#include "../include/mpp.h"
#include "../errmpp/errmpp.h"

#include <string.h>

typedef struct hi3516ev200_sys_init_in_struct {
    unsigned int width;
    unsigned int height;
    unsigned int cnt;
} hi3516ev200_sys_init_in;

static int hi3516ev200_sys_init(error_in *err, hi3516ev200_sys_init_in *in) {
    unsigned int mpp_error_code = 0;

    return ERR_GENERAL;
}
*/
import "C"

import (
    "application/pkg/mpp/errmpp"   
    "application/pkg/logger"
)

func initFamily() error {
    var inErr C.error_in
    var in C.hi3516ev200_sys_init_int

    in.width = C.uint(width)
    in.height = C.uint(height)
    in.cnt = C.uint(cnt)

    logger.Log.Trace().
        Uint("width", uint(in.width)).
        Uint("height", uint(in.height)).
        Uint("cnt", uint(in.cnt)).
        Msg("SYS params")
  
    err := C.hi3516ev200_sys_init(&errorCode, &in)
    if err != C.ERR_NONE {
        return errmpp.New(uint(inErr.f), uint(inErr.mpp))
    }

    return nil
}


