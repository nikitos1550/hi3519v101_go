package vo

//#include "vo.h"
import "C"

import (
    "application/pkg/logger"
    "application/pkg/mpp/errmpp"
)

func Init() {
    var inErr C.error_in

    err := C.mpp_vo_init(&inErr)

    if err != C.ERR_NONE {
        logger.Log.Fatal().
            Str("error", errmpp.New(C.GoString(inErr.name), uint(inErr.code)).Error()).
            Msg("VO")
    }
    logger.Log.Debug().
        Msg("VO inited")


    err2 := C.mpp_vo_bind_vpss_test(&inErr)

    if err2 != C.ERR_NONE {
        logger.Log.Fatal().
            Str("error", errmpp.New(C.GoString(inErr.name), uint(inErr.code)).Error()).
            Msg("VO")
    }
    logger.Log.Debug().
        Msg("VO binded")

}
