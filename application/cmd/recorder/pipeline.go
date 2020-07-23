package main

import (
    "sync"

    "application/core/mpp/vpss"
    "application/core/mpp/venc"

    "application/core/processing/schedule"

    "application/core/streamer/jpeg"
    "application/core/streamer/mjpeg"

    "application/core/mpp/connection"

    "application/core/logger"
)

var (
    channelMain     *vpss.Channel
    channelSmall    *vpss.Channel

    scheduleObj     *schedule.Schedule

    encoderH26XMain *venc.Encoder
    encoderMjpeg    *venc.Encoder

    jpegSmall       *jpeg.Jpeg
    mjpegSmall      *mjpeg.Mjpeg

    pipelineLock    sync.RWMutex
)

func initPipeline() {
    var err error

    channelMain, err        = vpss.New(0, "main", vpss.Parameters{
        Width: 3840,
        Height: 2160,
        Fps: 25,
    })
    if err != nil {
        logger.Log.Fatal().
            Str("reason", err.Error()).
            Msg("Main channel failed")
    }

    channelSmall, err       = vpss.New(1, "small", vpss.Parameters{
        Width: 1280,
        Height: 720,
        Fps: 1,
    })
    if err != nil {
        logger.Log.Fatal().
            Str("reason", err.Error()).
            Msg("Small channel failed")
    }

    scheduleObj, _          = schedule.New("scheduler", true)

    encoderH26XMain, err    = venc.New(0, "h26XMain", venc.Parameters{
        Codec: venc.H264,
        Profile: venc.High,
        //Codec: venc.H265,
        //Profile: venc.Baseline,
        Width: 3840,
        Height: 2160,
        Fps: 25,
        GopType: venc.BipredB,
        GopParams: venc.GopParameters{
            Gop: 100,
            BFrmNum: 3,
            BQpDelta:10,
            IPQpDelta:10,
        },
        BitControl: venc.Vbr,
        BitControlParams: venc.BitrateControlParameters{
            Bitrate: 1024*16,
            StatTime: 60,
            Fluctuate: 1,
            MinQp: 30,
            MaxQp: 50,
            MinIQp: 30,
        },
    })
    if err != nil {
        logger.Log.Fatal().
            Str("reason", err.Error()).
            Msg("Main encoder failed")
    }


    encoderMjpeg, err       = venc.New(1, "mjpegSmall", venc.Parameters{
        Codec: venc.MJPEG,
        Profile: venc.Baseline,
        Width: 1280,
        Height: 720,
        Fps: 1,
        GopType: venc.NormalP,
        GopParams: venc.GopParameters{
            Gop: 30,
        },
        BitControl: venc.Cbr,
        BitControlParams: venc.BitrateControlParameters{
            Bitrate: 512,
            StatTime: 60,
            Fluctuate: 1,
        },
    })
    if err != nil {
        logger.Log.Fatal().
            Str("reason", err.Error()).
            Msg("Small encoder failed")
    }



    jpegSmall, err          = jpeg.New("small")
    if err != nil {
        logger.Log.Fatal().
            Str("reason", err.Error()).
            Msg("Jpeg streamer failed")
    }

    mjpegSmall, err          = mjpeg.New("small")
    if err != nil {
        logger.Log.Fatal().
            Str("reason", err.Error()).
            Msg("Mjpeg streamer failed")
    }

    err = connection.ConnectRawFrame(channelMain, scheduleObj)
    if err != nil {
        logger.Log.Fatal().
            Str("reason", err.Error()).
            Msg("Connect main channel to schedule processing failed")

    }

    err = connection.ConnectRawFrame(scheduleObj, encoderH26XMain)
    if err != nil {
        logger.Log.Fatal().
            Str("reason", err.Error()).
            Msg("Connect schedule processing to main encoder failed")

    }
    /*
    err = connection.ConnectBind(channelMain, encoderH264Main)
    if err != nil {
        logger.Log.Fatal().
            Str("reason", err.Error()).
            Msg("Connect small channel to jpeg encoder failed")

    }
    */

    //err = connection.ConnectRawFrame(channelSmall, encoderMjpeg)
    err = connection.ConnectBind(channelSmall, encoderMjpeg)
    if err != nil {
        logger.Log.Fatal().
            Str("reason", err.Error()).
            Msg("Connect small channel to jpeg encoder failed")

    }

    err = connection.ConnectEncodedData(encoderMjpeg, jpegSmall)
    if err != nil {
        logger.Log.Fatal().
            Str("reason", err.Error()).
            Msg("Connect small channel to jpeg encoder failed")
    }

    err = connection.ConnectEncodedData(encoderMjpeg, mjpegSmall)
    if err != nil {
        logger.Log.Fatal().
            Str("reason", err.Error()).
            Msg("Connect small channel to mjpeg encoder failed")
    }

    //err = encoderH26XMain.SetScene(1) //experimental
    //if err != nil {
    //    logger.Log.Fatal().
    //        Str("reason", err.Error()).
    //        Msg("Can`t set scene for main encoder")
    //}

    err = encoderH26XMain.Start()
    if err != nil {
        logger.Log.Fatal().
            Str("reason", err.Error()).
            Msg("Can`t start main encoder")
    }

    err = encoderMjpeg.Start()
    if err != nil {
        logger.Log.Fatal().
            Str("reason", err.Error()).
            Msg("Can`t start small encoder")
    }
}
