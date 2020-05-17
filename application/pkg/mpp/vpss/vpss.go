package vpss

//#include "vpss.h"
import "C"

import (
    "flag"
    "unsafe"
    "application/pkg/mpp/vi"
    "application/pkg/mpp/errmpp"
    "application/pkg/logger"
    "application/pkg/buildinfo"

    "time"
)

var (
    nr bool
    nrFrmNum uint
)
func init() {
    flag.BoolVar(&nr, "vpss-nr", true, "Noise remove enable")

    if buildinfo.Family == "hi3516av200" {
        //flag.BoolVar(&nr, "vpss-nr", true, "Noise remove enable") //moved outside as common param
        flag.UintVar(&nrFrmNum, "vpss-nr-frames", 2, "Noise remove reference frames number [1;2]")
    }
}

func maxChannels() uint {
    return uint(C.VPSS_MAX_PHY_CHN_NUM)
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
            //in.nr = 1
        }   //else {
            //in.nr = 0
            //}
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

func createChannel(channel Channel) { //TODO return error
    var inErr C.error_in
    var in C.mpp_vpss_create_channel_in

    in.channel_id = C.uint(channel.ChannelId)
    in.width = C.uint(channel.Width)
    in.height = C.uint(channel.Height)
    in.vi_fps = C.uint(vi.Fps())
    in.fps = C.uint(channel.Fps)

    logger.Log.Trace().
        Int("channelId", channel.ChannelId).
        Uint("width", uint(in.width)).
        Uint("height", uint(in.height)).
        Uint("vi_fps", uint(in.vi_fps)).
        Uint("fps", uint(in.fps)).
        Msg("VPSS channel params")

    err := C.mpp_vpss_create_channel(&inErr, &in)
    
    if err != 0 {
        logger.Log.Fatal(). //log temporary, should generate and return error
            Str("error", errmpp.New(C.GoString(inErr.name), uint(inErr.code)).Error()).
            Msg("VPSS")
    }

    go func() {
        sendDataToClients(channel)
    }()

    //return nil
}

func destroyChannel(channel Channel) { //TODO return error
    var inErr C.error_in
    var in C.mpp_vpss_destroy_channel_in

    in.channel_id = C.uint(channel.ChannelId)

    err := C.mpp_vpss_destroy_channel(&inErr, &in)

    if err != 0 {
        logger.Log.Fatal(). //log temporary, should generate and return error
            Str("error", errmpp.New(C.GoString(inErr.name), uint(inErr.code)).Error()).
            Msg("VPSS")
    }

    //return nil
}

func sendDataToClients(channel Channel) {
    logger.Log.Trace().
        Int("channelId", channel.ChannelId).
        Str("name", "sendDataToClients").
        Msg("VPSS rutine started")

    for {
        if (!channel.Started){
            break
        }

        var err C.int
        var inErr C.error_in
        var frame unsafe.Pointer

        //hi3516cv100 family doesn`t provide blocking getFrame call
        if buildinfo.Family == "hi3516cv100" {  //TODO
            time.Sleep(1 * time.Second)         //now we will just sleep here
        }

        err = C.mpp_receive_frame(&inErr, C.uint(channel.ChannelId), &frame);
        if err != C.ERR_NONE {
            logger.Log.Warn().
                Int("channelId", channel.ChannelId).
                Str("error", errmpp.New(C.GoString(inErr.name), uint(inErr.code)).Error()).
                Msg("VPSS failed receive frame")
            continue
        } else {
            //logger.Log.Trace().
            //    Int("channelId", channel.ChannelId).
            //    Msg("VPSS received frame")
        }

        
        for processing, _ := range channel.Clients {
            processing.Callback(frame)
        }
        
        err = C.mpp_release_frame(&inErr, C.uint(channel.ChannelId));
        if err != C.ERR_NONE {
            logger.Log.Error().
                Int("channelId", channel.ChannelId).
                Str("error", errmpp.New(C.GoString(inErr.name), uint(inErr.code)).Error()).
                Msg("VPSS failed release frame")
        } else {
            //logger.Log.Trace().
            //    Int("channelId", channel.ChannelId).
            //    Msg("VPSS released frame")
        }

    }

    logger.Log.Trace().        
        Int("channelId", channel.ChannelId).    
        Str("name", "sendDataToClients").
        Msg("VPSS rutine stopped")
}

//export go_logger_vpss
func go_logger_vpss(level C.int, msgC *C.char) {
        logger.CLogger("VPSS", int(level), C.GoString(msgC))
}
