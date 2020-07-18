package main

import (
    "application/core/mpp/vpss"
    "application/core/mpp/venc"

    "application/core/streamer/jpeg"

    "application/archive/recorder"

    "application/core/mpp/connection"

    "application/core/logger"
)

var (
    channelMain     *vpss.Channel
    channelSmall    *vpss.Channel

    encoderH264Main *venc.Encoder
    encoderMjpeg    *venc.Encoder

    jpegSmall       *jpeg.Jpeg

    recorderObj     *recorder.Recorder
)

func pipelineInit() {
    var err error

    channelMain, err        = vpss.New(0, "main", vpss.Parameters{
        Width: 3840,
        Height: 2160,
        Fps: 30,
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

    encoderH264Main, err    = venc.New(0, "h264Main", venc.Parameters{
        Codec: venc.H264,
        Profile: venc.High,
        Width: 3840,
        Height: 2160,
        Fps: 30,
        GopType: venc.NormalP,
        GopParams: venc.GopParameters{
            Gop: 60,
        },
        BitControl: venc.Cbr,
        BitControlParams: venc.BitrateControlParameters{
            Bitrate: 1024,
            StatTime: 60,
            Fluctuate: 1,
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
            Gop: 60,
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

    err = connection.ConnectBind(channelMain, encoderH264Main)
    if err != nil {
        logger.Log.Fatal().
            Str("reason", err.Error()).
            Msg("Connect main channel to main encoder failed")

    }

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

    err = encoderH264Main.Start()
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

    recorderObj, _ = recorder.New("testrecorder")

    err = connection.ConnectEncodedData(encoderH264Main, recorderObj)
    if err != nil {
        logger.Log.Fatal().
            Str("reason", err.Error()).
            Msg("Connect main encoder to recorder failed")
    }
}
