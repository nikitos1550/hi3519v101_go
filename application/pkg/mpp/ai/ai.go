package ai

//#include "ai.h"
import "C"

import (
    "application/pkg/logger"
    "application/pkg/mpp/errmpp"
)

func Init() {
    logger.Log.Debug().
        Msg("AI init")

    var inErr C.error_in

    err := C.mpp_ai_config_inner(&inErr)       

    if err != C.ERR_NONE {
        logger.Log.Fatal().
            Str("error", C.GoString(inErr.name)).
            Int("code", int(inErr.code)).
            Msg("mpp_ai_config_inner")
    }

    err = C.mpp_ao_test(&inErr)

    if err != C.ERR_NONE {
        if err == C.ERR_GENERAL {
        logger.Log.Fatal().
            Str("error", C.GoString(inErr.name)).
            Int("code", int(inErr.code)).
            Msg("mpp_ao_test")
        }
        if err == C.ERR_MPP {
            logger.Log.Fatal().
                Str("error", errmpp.New(C.GoString(inErr.name), uint(inErr.code)).Error()).
                Msg("mpp_ao_test")
        }
    }


    err = C.mpp_ai_test(&inErr)

    if err != C.ERR_NONE {
        if err == C.ERR_GENERAL {
        logger.Log.Fatal().
            Str("error", C.GoString(inErr.name)).
            Int("code", int(inErr.code)).
            Msg("mpp_ai_test")
        }
        if err == C.ERR_MPP {
            logger.Log.Fatal().
                Str("error", errmpp.New(C.GoString(inErr.name), uint(inErr.code)).Error()).
                Msg("mpp_ai_test")
        }
    }

}
