package main

import (
    "net/http"

    "application/core/mpp/vpss"
    "application/core/mpp/venc"
    "application/core/mpp/connection"
)

func presetPreview(w http.ResponseWriter, r *http.Request) {
    channelPreview, err := channelGroup.CreateChannel("preview", vpss.Parameters{
        Width: 1280,
        Height: 720,
        Fps: 1,
    })
    if err != nil {

    }

    encoderPreview, err := encoderGroup.CreateEncoder("preview", venc.Parameters{
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

    }

    jpegPreview, err := jpegGroup.CreateJpeg("preview")
    if err != nil {

    }

    err = connection.ConnectBind(channelPreview, encoderPreview)
    if err != nil {

    }

    err = connection.ConnectEncodedData(encoderPreview, jpegPreview)
    if err != nil {

    }

}
