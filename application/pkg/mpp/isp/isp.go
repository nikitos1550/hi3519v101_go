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

    //cmos.Register() // moved to main mpp init func

    var inErr C.error_in
    var in C.mpp_isp_init_in

    in.width = C.uint(cmos.S.Width())
    in.height = C.uint(cmos.S.Height())

    ispCrop := cmos.S.IspCrop()

    in.isp_crop_x0 = C.uint(ispCrop.X0)
    in.isp_crop_y0 = C.uint(ispCrop.Y0)
    in.isp_crop_width = C.uint(ispCrop.Width)
    in.isp_crop_height = C.uint(ispCrop.Height)
    
    in.fps = C.uint(cmos.S.Fps())

    switch cmos.S.Wdr() {
        case cmos.WDRNone:
            in.wdr = C.WDR_MODE_NONE
        case cmos.WDR2TO1:
            in.wdr = C.WDR_MODE_2To1_LINE
        case cmos.WDR2TO1F:
            in.wdr = C.WDR_MODE_2To1_FRAME
        case cmos.WDR2TO1FFR:
            in.wdr = C.WDR_MODE_2To1_FRAME_FULL_RATE
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
        Uint("crop_x0", uint(in.isp_crop_x0)).
        Uint("crop_y0", uint(in.isp_crop_y0)).
        Uint("crop_width", uint(in.isp_crop_width)).
        Uint("crop_height", uint(in.isp_crop_height)).
        Uint("width", uint(in.width)).
        Uint("height", uint(in.height)).
        Uint("fps", uint(in.fps)).
        Uint("bayer", uint(in.bayer)).
        Uint("wdr", uint(in.wdr)).
        Msg("ISP params")

    err := C.mpp_isp_init(&inErr, &in)
    switch err {
        case C.ERR_MPP:
            logger.Log.Fatal().
                Str("error", errmpp.New(C.GoString(inErr.name), uint(inErr.code)).Error()).
                Msg("ISP")
        case C.ERR_GENERAL:
            logger.Log.Fatal().
                Uint("error", uint(inErr.code)).
                Str("name", C.GoString(inErr.name)).
                Msg("ISP")
        default:
            break
    }

    //go func() {   //ISP thread started in C space
    //    logger.Log.Trace(). 
    //        Msg("ISP task started")
    //    C.mpp_isp_thread(nil)
    //    logger.Log.Error().
    //        Msg("ISP task failed")
    //}()

    logger.Log.Debug().
        Msg("ISP inited")

}

//export go_logger_isp
func go_logger_isp(level C.int, msgC *C.char) {
        logger.CLogger("ISP", int(level), C.GoString(msgC))
}

