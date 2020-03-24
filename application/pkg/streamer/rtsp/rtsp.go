//+build streamerRtsp

package rtsp

import (
	"net/http"
	"application/pkg/mpp/venc"
	"application/pkg/openapi"
	
	"github.com/aler9/gortsplib" 
)

type responseRecord struct {
	Message string
}

type rtspStream struct {
	Name string
	EncoderId string
	Started bool
	CameraIn chan []byte
	RtspOut chan gortsplib.InterleavedFrame
}

var (
	server *program
	sdpPattern string
	rtspStreams map[string] rtspStream
)

func init() {
	rtspStreams = make(map[string] rtspStream)

	server = CreateRtspServer()

    openapi.AddApiRoute("rtspApiDescription", "/rtsp", "GET", rtspApiDescription)
    openapi.AddApiRoute("startRtspStream", "/rtsp/start", "GET", startRtspStream)
    openapi.AddApiRoute("stopRtspStream", "/rtsp/stop", "GET", stopRtspStream)
    openapi.AddApiRoute("listRtspStreams", "/rtsp/list", "GET", listRtspStreams)

	sdpPattern =
`v=0
o=- 123 1 IN IP4 10.104.44.103
s=Session streamed by "testOnDemandRTSPServer"
i=stream
t=0 0
a=tool:LIVE555 Streaming Media v2020.03.06
a=type:broadcast
a=control:*
a=range:npt=0-
a=x-qt-text-nam:Session streamed by "testOnDemandRTSPServer"
a=x-qt-text-inf:stream
m=video 0 RTP/AVP 96
c=IN IP4 0.0.0.0
b=AS:500
a=rtpmap:96 H264/90000
a=fmtp:96 packetization-mode=1;profile-level-id=64001F;sprop-parameter-sets=Z2QAH6wsaoFAFum4CAgIEA==,aO48sA==
a=control:track1`

}

func Init() {
}

func rtspApiDescription(w http.ResponseWriter, r *http.Request)  {
	openapi.ApiDescription(w, r, "Rtsp api:\n\n", "/rtsp")
}

func startRtspStream(w http.ResponseWriter, r *http.Request)  {
	ok, encoderId := openapi.GetStringParameter(w, r, "encoderId")
	if !ok {
		return
	}

	ok, streamName := openapi.GetStringParameter(w, r, "streamName")
	if !ok {
		return
	}

	_, encoderExists := venc.Encoders[encoderId]
	if (!encoderExists) {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{Message: "Failed to find encoder  " + encoderId})
		return
	}

	_, streamExists := rtspStreams[streamName]
	if (streamExists) {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{Message: "Stream with name " + streamName + " already exists"})
		return
	}

	stream := rtspStream{
		Name: streamName,
		EncoderId: encoderId,
		Started: true,
		CameraIn: make(chan []byte, 100),
		RtspOut: make(chan gortsplib.InterleavedFrame, 10),
	}

	server.AddPublisher(sdpPattern, stream.Name, stream.RtspOut)

	venc.SubsribeEncoder(encoderId, stream.CameraIn)
	
    go func() {
		packetizer := CreatePacketizer()
		for {
			if (!stream.Started){
				break
			}
			data := <- stream.CameraIn
			packets := packetizer.NalH264ToRtp(data)
			for _, p := range packets {
				stream.RtspOut <- gortsplib.InterleavedFrame{
					Channel: 0,
					Content: p,
				}
			}
		}
    }()

	openapi.ResponseSuccessWithDetails(w, responseRecord{Message: "Rtsp was started"})
}

func stopRtspStream(w http.ResponseWriter, r *http.Request)  {
}

func listRtspStreams(w http.ResponseWriter, r *http.Request)  {
}

