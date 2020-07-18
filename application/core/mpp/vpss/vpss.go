package vpss

//#include "vpss.h"
import "C"

import (
    "flag"
    "errors"
    "fmt"

    "application/core/mpp/vi"
    "application/core/mpp/errmpp"
    "application/core/logger"
    "application/core/compiletime"
)

var (
    nr bool
    nrFrmNum uint

    Amount int = C.VPSS_MAX_PHY_CHN_NUM

    channels []Channel
)

func init() {
    flag.BoolVar(&nr, "vpss-nr", true, "Noise remove enable")

    if compiletime.Family == "hi3516av200" {
        flag.UintVar(&nrFrmNum, "vpss-nr-frames", 2, "Noise remove reference frames number [1;2]")
    }

    channels = make([]Channel, Amount)

    for i := 0; i < Amount; i++ {
        channels[i].Id = i
        //logger.Log.Trace().Int("id", i).Int("channelid", channels[i].Id).Msg("VPSS init")
        fmt.Println("id ",i," channelid ", channels[i].Id)
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

    if compiletime.Family == "hi3516av200" {
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

    in.id           = C.uint(id)

    if params.Width < 0 || params.Width > vi.Width() {
        return errors.New("Width is not correct")
    }
    if params.Height < 0 || params.Height > vi.Height() {
        return errors.New("Height is not correct")
    }

    if id == 0 {
        if  params.Width != vi.Width() ||
            params.Height != vi.Height() {
            return errors.New("Channel 0 width and height should be same as cmos")
        }
    }

    in.width        = C.uint(params.Width)
    in.height       = C.uint(params.Height)

    logger.Log.Trace().
        Int("fps", params.Fps).
        Msg("TMP VPSS FPS")

    if params.Fps < 1 || params.Fps > vi.Fps() {
        return errors.New("Fps is not correct")
    }

    in.vi_fps       = C.uint(vi.Fps())
    in.fps          = C.uint(params.Fps)

    //TODO depth valkue check
    in.depth        = 0 //C.uint(params.Depth)

    //TODO crop params check
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

func mppChangeDepth(id int, depth int) error {
    var inErr C.error_in
    var in C.mpp_vpss_change_channel_depth_in

    in.id = C.uint(id)
    in.depth = C.uint(depth)

    err := C.mpp_vpss_change_channel_depth(&inErr, &in)

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
