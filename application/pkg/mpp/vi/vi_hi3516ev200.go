//+build arm
//+build hi3516ev200

package vi

/*
#include "../include/mpp.h"
#include "../../logger/logger.h"

#include <string.h>

#define ERR_NONE                    0
#define ERR_HI_MPI_VI_SetDevAttr        2
#define ERR_HI_MPI_VI_EnableDev     3
#define ERR_HI_MPI_VI_SetChnAttr    4
#define ERR_HI_MPI_VI_EnableChn     5

int mpp4_vi_init(unsigned int *error_code) {
    *error_code = 0;

    return ERR_NONE;
}


*/
import "C"

import (
        "application/pkg/logger"

    "application/pkg/mpp/error"
)

func Init() {
    var errorCode C.uint

    switch err := C.mpp4_vi_init(&errorCode); err {
    case C.ERR_NONE:
        logger.Log.Debug().
                Msg("C.mpp4_vi_init() ok")
    case C.ERR_HI_MPI_VI_SetDevAttr:
        logger.Log.Fatal().
                Str("func", "ERR_HI_MPI_VI_SetDevAttr()").
                Int("error", int(errorCode)).
                Str("error_desc", error.Resolve(int64(errorCode))).
                Msg("C.mpp4_vi_init() error")
    case C.ERR_HI_MPI_VI_EnableDev:
        logger.Log.Fatal().
                Str("func", "ERR_HI_MPI_VI_EnableDev()").
                Int("error", int(errorCode)).
                Str("error_desc", error.Resolve(int64(errorCode))).
                Msg("C.mpp4_vi_init() error")

    case C.ERR_HI_MPI_VI_SetChnAttr:
        logger.Log.Fatal().
                Str("func", "ERR_HI_MPI_VI_SetChnAttr()"). 
                Int("error", int(errorCode)).
                Str("error_desc", error.Resolve(int64(errorCode))).
                Msg("C.mpp4_vi_init() error")
    case C.ERR_HI_MPI_VI_EnableChn:
        logger.Log.Fatal().
                Str("func", "ERR_HI_MPI_VI_EnableChn()"). 
                Int("error", int(errorCode)).
                Str("error_desc", error.Resolve(int64(errorCode))).
                Msg("C.mpp4_vi_init() error")
    default:
        logger.Log.Fatal().
                Int("error", int(err)).
                Msg("C.mpp4_vi_init() Unexpected return")

    }
}

