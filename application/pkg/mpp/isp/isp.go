package isp

//#include "isp.h"
import "C"

import (
    //"errors"
    "application/pkg/mpp/errmpp"
    "application/pkg/logger"  
    "application/pkg/mpp/cmos"
)

func Init() {

    cmos.Register() // TODO check return

    var inErr C.error_in
    var in C.mpp_isp_init_in

    in.width = C.uint(cmos.S.Width())
    in.height = C.uint(cmos.S.Height())
    in.fps = C.uint(cmos.S.Fps())

    switch cmos.S.Wdr() {
        case cmos.WDRNone:
            in.wdr = C.WDR_MODE_NONE
        case cmos.WDR2TO1:
            in.wdr = C.WDR_MODE_2To1_LINE
        default:
            logger.Log.Fatal().
                Msg("Unknown WDR mode")
    }

    switch cmos.S.Bayer() {
        case cmos.RGGB:
            in.bayer = C.BAYER_RGGB
        case cmos.GRBG:
            in.bayer = C.BAYER_GRBG
        case cmos.GBRG:
            in.bayer = C.BAYER_GBRG
        case cmos.BGGR:
            in.bayer = C.BAYER_BGGR
        default:
            logger.Log.Fatal().
                Msg("Unknown CMOS bayer")
    }

    logger.Log.Trace().
        Uint("width", uint(in.width)).
        Uint("height", uint(in.height)).
        Uint("fps", uint(in.fps)).
        Uint("bayer", uint(in.bayer)).
        Uint("wdr", uint(in.wdr)).
        Msg("ISP params")

    err := C.mpp_isp_init(&inErr, &in)
    switch err {
        case C.ERR_MPP:
            //return errmpp.New(uint(inErr.f), uint(inErr.mpp))
            logger.Log.Fatal().
                Str("error", errmpp.New(C.GoString(inErr.name), uint(inErr.code)).Error()).
                Msg("ISP")
        case C.ERR_GENERAL:
            //return errors.New("ISP error TODO")
            logger.Log.Fatal().
                Str("error", "ISP error TODO").
                Msg("ISP")
        default:
            break
    }

    go func() {
        logger.Log.Trace().
            Msg("ISP task started")
        C.mpp_isp_thread(nil)
        logger.Log.Error().
            Msg("ISP task failed")
    }()

    /*
    err := ispInit()
    if err != nil {
        logger.Log.Fatal().
            Str("error", err.Error()).
            Msg("ISP")
    }
    */
    logger.Log.Debug().
        Msg("ISP inited")

}
/*
func ispInit() error {
    err := C.mpp_isp_init(&inErr, &in)
    switch err {
        case C.ERR_MPP:
            return errmpp.New(uint(inErr.f), uint(inErr.mpp))
        case C.ERR_GENERAL:
            return errors.New("ISP error TODO")
        default:
            return nil
    }
}
*/

//export go_logger_isp
func go_logger_isp(level C.int, msgC *C.char) {
        logger.CLogger("ISP", int(level), C.GoString(msgC))
}

