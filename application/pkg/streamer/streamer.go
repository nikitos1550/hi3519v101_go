package streamer

import (
	//TODO avoid implicit init
	"application/pkg/streamer/jpeg"
    "application/pkg/streamer/webrtc"
)

func Init() {
    jpeg.Init()
    webrtc.Init()
}
