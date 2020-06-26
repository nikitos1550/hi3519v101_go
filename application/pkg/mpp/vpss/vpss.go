package vpss

//#include "vpss.h"
import "C"

import (
    "flag"
    "time"
    "unsafe"

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

    in.channel_id   = C.uint(id)            //C.uint(channel.ChannelId)
    in.width        = C.uint(params.Width)  //C.uint(channel.Width)
    in.height       = C.uint(params.Height) //C.uint(channel.Height)
    in.vi_fps       = C.uint(vi.Fps())
    in.fps          = C.uint(params.Fps)    //C.uint(channel.Fps)

    logger.Log.Trace().
        Uint("channel", uint(in.channel_id)).
        Uint("width", uint(in.width)).
        Uint("height", uint(in.height)).
        Uint("vi_fps", uint(in.vi_fps)).
        Uint("fps", uint(in.fps)).
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

    in.channel_id = C.uint(id)          //C.uint(channel.ChannelId)

    err := C.mpp_vpss_destroy_channel(&inErr, &in)

    if err != 0 {
        logger.Log.Fatal().
            Str("error", errmpp.New(C.GoString(inErr.name), uint(inErr.code)).Error()).
            Msg("VPSS")
        return errmpp.New(C.GoString(inErr.name), uint(inErr.code))
    }

    return nil
}

func sendDataToClients(c *channel) {
    periodTreshhold := uint64(1500000 / c.params.Fps)
    periodWait      := uint64(5000 / c.params.Fps)

    logger.Log.Trace().
        Int("channel", c.id).
        Str("name", "sendDataToClients").
        Uint64("treshhold", periodTreshhold).
        Uint64("wait", periodWait).
        Msg("VPSS rutine started")

    for {
        if (!c.started){
            break
        }

        var err C.int
        var inErr C.error_in
        var frame unsafe.Pointer
        var pts C.ulonglong

        //hi3516cv100 family doesn`t provide blocking getFrame call
        if buildinfo.Family == "hi3516cv100" {  //TODO
            time.Sleep(1 * time.Second)         //now we will just sleep here
        }

        err = C.mpp_receive_frame(&inErr, C.uint(c.id), &frame, &pts, C.uint(periodWait));

        if err != C.ERR_NONE {
            logger.Log.Warn().
                Int("channel", c.id).
                Str("error", errmpp.New(C.GoString(inErr.name), uint(inErr.code)).Error()).
                Msg("VPSS failed receive frame")
            continue
        }

        c.stat.TsPrev = c.stat.TsLast
        c.stat.TsLast = uint64(pts)
        c.stat.Count++

        var period uint64 = c.stat.TsLast - c.stat.TsPrev

        if period > periodTreshhold {
            if (c.stat.Count > 1) {
                c.stat.Drops++
            }
        } else {
            c.stat.PeriodAvg += ((float64(period) - c.stat.PeriodAvg) / float64(c.stat.Count))
        }

        //logger.Log.Trace().
        //    Float64("period", channel.Stat.PeriodAvg).
        //    //Uint("period", channel.Stat.PeriodAvg).
        //    Uint64("count", channel.Stat.Count).
        //    Uint64("drops", channel.Stat.Drops).
        //    Float64("delta", float64(period)).
        //    //Float32("test",  ((1 / float32(channel.Fps))*1.5)).
        //    Msg("STAT")

        for processing, _ := range c.clients {
            processing.Callback(frame)
        }

        err = C.mpp_release_frame(&inErr, C.uint(c.id));

        if err != C.ERR_NONE {
            logger.Log.Error().
                Int("channel", c.id).
                Str("error", errmpp.New(C.GoString(inErr.name), uint(inErr.code)).Error()).
                Msg("VPSS failed release frame")
        }
    }

    c.rutineStop <- true //TODO improve channel stop

    logger.Log.Trace().
        Int("channel", c.id).
        Str("name", "sendDataToClients").
        Msg("VPSS rutine stopped")
}

//export go_logger_vpss
func go_logger_vpss(level C.int, msgC *C.char) {
        logger.CLogger("VPSS", int(level), C.GoString(msgC))
}
