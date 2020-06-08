//+build hi3516cv200 hi3516av100 hi3516av200 hi3516cv300

package ai

//#include "ai.h"
import "C"

import (
    "unsafe"

    "application/pkg/logger"
    "application/pkg/mpp/errmpp"
    "application/pkg/buildinfo"
)

func IsAudioExistTmp() bool { //temporary function to check if there any audio avalible in the system
    if buildinfo.Family == "hi3516av200" {
        return true
    } else {
        return false
    }
}

//export go_callback_raw_tmp
func go_callback_raw_tmp(info_pointer *C.audio_info_from_c, data_pointer *C.audio_data_from_c) {

    dataFromC := (*[1 << 10]C.audio_data_from_c)(unsafe.Pointer(data_pointer))[:1:1]
    length := int(dataFromC[0].length)

    //infoFromC := (*[1 << 10]C.audio_info_from_c)(unsafe.Pointer(info_pointer))[:1:1]

    //var timestamp uint64
    //timestamp = uint64(infoFromC[0].timestamp)
    //logger.Log.Trace().
    //    Uint64("ts", timestamp).
    //    Msg("raw audio")

    data := make([]byte, length)

    p := (*[1 << 28]byte)(unsafe.Pointer(dataFromC[0].data))[:dataFromC[0].length:dataFromC[0].length]
    n := copy(data[0:], p)

    if n != length {
        logger.Log.Warn().
            Int("length", length).
            Int("copied", n).
            Msg("go_callback_raw_tmp")
    }
    //data is your target
}

//export go_callback_opus_tmp
func go_callback_opus_tmp(info_pointer *C.audio_info_from_c, data_pointer *C.audio_data_from_c) {

    dataFromC := (*[1 << 10]C.audio_data_from_c)(unsafe.Pointer(data_pointer))[:1:1]
    length := int(dataFromC[0].length)

    //infoFromC := (*[1 << 10]C.audio_info_from_c)(unsafe.Pointer(info_pointer))[:1:1]

    //var timestamp uint64
    //timestamp = uint64(infoFromC[0].timestamp)
    //logger.Log.Trace().
    //    Uint64("ts", timestamp).
    //    Msg("raw audio")

    data := make([]byte, length)

    p := (*[1 << 28]byte)(unsafe.Pointer(dataFromC[0].data))[:dataFromC[0].length:dataFromC[0].length]
    n := copy(data[0:], p)

    if n != length {
        logger.Log.Warn().
            Int("length", length).
            Int("copied", n).
            Msg("go_callback_raw_tmp")
    }    
    //data is your target
}

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
