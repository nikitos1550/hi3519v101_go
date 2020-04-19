//+build arm
//+build hi3516ev200

package sys

/*
#include "../include/mpp_v4.h"

#include <string.h>

#define ERR_NONE                0
#define ERR_MPP                 1
#define ERR_GENERAL             2

typedef struct hi3516ev200_sys_init_in_struct {
    unsigned int width;
    unsigned int height;
    unsigned int cnt;
} hi3516ev200_sys_init_in;

static int hi3516ev200_sys_init(unsigned int *error_code, hi3516ev200_sys_init_in *in) {
    *error_code = 0;

    return ERR_GENERAL;
}
*/
import "C"

import (
    "application/pkg/mpp/errmpp"   
    "application/pkg/logger"
)

func initFamily() error {
    var errorCode C.uint
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
        return errors.New("SYS error TOD")
    }

    return nil
}


