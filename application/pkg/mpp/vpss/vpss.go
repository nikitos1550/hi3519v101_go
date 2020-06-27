package vpss

//#include "vpss.h"
import "C"

import (
    "flag"

    "application/pkg/mpp/vi"
    "application/pkg/mpp/errmpp"
    "application/pkg/logger"
    "application/pkg/buildinfo"
)

var (
    nr bool
    nrFrmNum uint

    channelsAmount int = C.VPSS_MAX_PHY_CHN_NUM
)

func init() {
    flag.BoolVar(&nr, "vpss-nr", true, "Noise remove enable")

    if buildinfo.Family == "hi3516av200" {
        flag.UintVar(&nrFrmNum, "vpss-nr-frames", 2, "Noise remove reference frames number [1;2]")
    }
}

func Init() {
    var inErr C.error_in
    var in C.mpp_vpss_init_in

    in.width = C.uint(vi.Width())
    in.height = C.uint(vi.Height())

    if nr == true {
        in.nr = 1
    } else {
        in.nr = 0
    }

    if buildinfo.Family == "hi3516av200" {
        if nr == true {

            if nrFrmNum < 1 || nrFrmNum > 2 {
                logger.Log.Fatal().
                    Uint("vpss-nr-frames", nrFrmNum).
                    Msg("vpss-nr-frames shoud be 1 or 2")
            }
            in.nr_frames = C.uchar(nrFrmNum)
        }
    }

    logger.Log.Trace().
        Uint("width", uint(in.width)).
        Uint("height", uint(in.height)).
        Uint("nr", uint(in.nr)).
        Uint("nr_frames", uint(in.nr_frames)).
        Msg("VPSS params")

    err := C.mpp_vpss_init(&inErr, &in)

    if err != 0 {
        logger.Log.Fatal().
            Str("error", errmpp.New(C.GoString(inErr.name), uint(inErr.code)).Error()).
            Msg("VPSS")
    }

    logger.Log.Debug().
        Msg("VPSS inited")
}

func mppCreateChannel(id int, params Parameters) error {
    var inErr C.error_in
    var in C.mpp_vpss_create_channel_in

    in.id   = C.uint(id)
    in.width        = C.uint(params.Width)
    in.height       = C.uint(params.Height)
    in.vi_fps       = C.uint(vi.Fps())
    in.fps          = C.uint(params.Fps)
    in.depth        = C.uint(params.Depth)
    in.crop_x       = C.uint(params.CropX)
    in.crop_y       = C.uint(params.CropY)
    in.crop_width   = C.uint(params.CropWidth)
    in.crop_height  = C.uint(params.CropHeight)

    logger.Log.Trace().
        Uint("channel",     uint(in.id)).
        Uint("width",       uint(in.width)).
        Uint("height",      uint(in.height)).
        Uint("vi_fps",      uint(in.vi_fps)).
        Uint("fps",         uint(in.fps)).
        Uint("depth",       uint(in.depth)).
        Uint("crop_x",      uint(in.crop_x)).
        Uint("crop_y",      uint(in.crop_y)).
        Uint("crop_width",  uint(in.crop_width)).
        Uint("crop_height", uint(in.crop_height)).
        Msg("VPSS channel params")

    err := C.mpp_vpss_create_channel(&inErr, &in)

    if err != 0 {
        logger.Log.Fatal().
            Str("error", errmpp.New(C.GoString(inErr.name), uint(inErr.code)).Error()).
            Msg("VPSS")
        return errmpp.New(C.GoString(inErr.name), uint(inErr.code))
    }

    return nil
}

func mppDestroyChannel(id int) error {
    var inErr C.error_in
    var in C.mpp_vpss_destroy_channel_in

    in.id = C.uint(id)

    err := C.mpp_vpss_destroy_channel(&inErr, &in)

    if err != 0 {
        logger.Log.Fatal().
            Str("error", errmpp.New(C.GoString(inErr.name), uint(inErr.code)).Error()).
            Msg("VPSS")
        return errmpp.New(C.GoString(inErr.name), uint(inErr.code))
    }

    return nil
}

//export go_logger_vpss
func go_logger_vpss(level C.int, msgC *C.char) {
        logger.CLogger("VPSS", int(level), C.GoString(msgC))
}
