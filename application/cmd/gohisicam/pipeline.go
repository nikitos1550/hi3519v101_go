package main

import (
    "application/core/mpp/vpss"
    "application/core/mpp/venc"

    "application/core/streamer/jpeg"
    "application/core/streamer/mjpeg"

    //"application/pkg/streamer/pipe"
    //"application/pkg/streamer/rtsp"
    //"application/pkg/streamer/yuv"
    //"application/pkg/streamer/webrtc"

    "application/core/processing/forward"
    "application/core/processing/quirc"
)

var (
    channelGroup    *vpss.ChannelGroup
    encoderGroup    *venc.EncoderGroup
    jpegGroup       *jpeg.JpegGroup
    mjpegGroup      *mjpeg.MjpegGroup

    forwardGroup    *forward.ForwardGroup
    quircGroup      *quirc.ForwardGroup
)

func pipelineInit() {
    channelGroup    = vpss.NewGroup(4)
    encoderGroup    = venc.NewGroup(16)

    jpegGroup       = jpeg.NewGroup(16)
    mjpegGroup      = mjpeg.NewGroup(16)

    //webrtc.Init()

    forwardGroup    = forward.NewGroup(8)
    quircGroup      = quirc.NewGroup(1)
}
