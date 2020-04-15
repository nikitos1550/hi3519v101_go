package streamer

import (
	"application/pkg/streamer/file"
    "application/pkg/streamer/jpeg"
	"application/pkg/streamer/pipe"
    "application/pkg/streamer/raw"
    "application/pkg/streamer/rtsp"
    "application/pkg/streamer/webrtc"
    "application/pkg/streamer/ws"

    "application/pkg/logger"
)

func Init() {
    file.Init()
    jpeg.Init()
    pipe.Init()
    raw.Init()
    rtsp.Init()
    webrtc.Init()
    ws.Init()

    logger.Log.Debug().
    	Msg("Streamers inited")
}
