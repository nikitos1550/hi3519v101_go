//+build streamerRtsp

package rtsp

import (
	"application/pkg/mpp/venc"
	"application/pkg/openapi"
	"log"
	"net/http"
	"strconv"    

	"github.com/aler9/gortsplib"
)

type responseRecord struct {
	Message string
}

type rtspStream struct {
	Name string
	EncoderId int
	EncoderType string
	Started bool
	Published bool
	Sps []byte
	Pps []byte
	//CameraIn chan []byte
    CameraIn chan venc.ChannelEncoder
	RtspOut chan gortsplib.InterleavedFrame
}

type rtspInfo struct {
	Name string
	EncoderId int
}

var (
	server *program
	rtspStreams map[string] rtspStream
)

func init() {
//	rtspStreams = make(map[string] rtspStream)

	//server = CreateRtspServer()

    openapi.AddApiRoute("rtspApiDescription", "/rtsp", "GET", rtspApiDescription)
    openapi.AddApiRoute("startRtspStream", "/rtsp/start", "GET", startRtspStream)
    openapi.AddApiRoute("stopRtspStream", "/rtsp/stop", "GET", stopRtspStream)
    openapi.AddApiRoute("listRtspStreams", "/rtsp/list", "GET", listRtspStreams)
}

func Init() {
        rtspStreams = make(map[string] rtspStream)

        server = CreateRtspServer()

}

func rtspApiDescription(w http.ResponseWriter, r *http.Request)  {
	openapi.ApiDescription(w, r, "Rtsp api:\n\n", "/rtsp")
}

func startRtspStream(w http.ResponseWriter, r *http.Request)  {
	ok, encoderId := openapi.GetIntParameter(w, r, "encoderId")
	if !ok {
		return
	}

	ok, streamName := openapi.GetStringParameter(w, r, "streamName")
	if !ok {
		return
	}

	encoder, encoderExists := venc.ActiveEncoders[encoderId]
	if (!encoderExists) {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{Message: "Failed to find encoder  " + strconv.Itoa(encoderId)})
		return
	}

	_, streamExists := rtspStreams[streamName]
	if (streamExists) {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{Message: "Stream with name " + streamName + " already exists"})
		return
	}

	rtspStreams[streamName] = rtspStream{
		Name: streamName,
		EncoderId: encoderId,
		EncoderType: encoder.Format,
		Started: true,
		Published: false,
		Sps: []byte{},
		Pps: []byte{},
		//CameraIn: make(chan []byte, 100),
        CameraIn: make(chan venc.ChannelEncoder, 100),
		RtspOut: make(chan gortsplib.InterleavedFrame, 1000),
	}

	venc.SubsribeEncoder(encoderId, rtspStreams[streamName].CameraIn)

    go writeVideoData(streamName)

	openapi.ResponseSuccessWithDetails(w, responseRecord{Message: "Rtsp was started"})
}

func stopRtspStream(w http.ResponseWriter, r *http.Request)  {
	ok, streamName := openapi.GetStringParameter(w, r, "streamName")
	if !ok {
		return
	}

	stream, exists := rtspStreams[streamName]
	if (!exists) {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{Message: "Stream not found"})
		return
	}

	stream.Started = false
	rtspStreams[streamName] = stream
	openapi.ResponseSuccessWithDetails(w, responseRecord{Message: "Stream was stopped"})
}

func listRtspStreams(w http.ResponseWriter, r *http.Request)  {
	var infos []rtspInfo
	for name, stream := range rtspStreams {
		info := rtspInfo{
			Name: name,
			EncoderId: stream.EncoderId,
		}

		infos = append(infos, info)
	}
	openapi.ResponseSuccessWithDetails(w, infos)
}

func writeVideoData(streamName string) {
	stream := rtspStreams[streamName]
	packetizer := CreatePacketizer()
	for {
		data := <-stream.CameraIn

		if (len(stream.Sps) == 0) {
			stream.Sps = ExtractSps(stream.EncoderType, data.Data)
		}

		if (len(stream.Pps) == 0) {
			stream.Pps = ExtractPps(stream.EncoderType, data.Data)
		}

		if (!stream.Published && len(stream.Sps) > 0 && len(stream.Pps) > 0) {
			sdp := CreateSdp(stream.EncoderType, stream.Name, stream.Sps, stream.Pps)
			server.AddPublisher(sdp, stream.Name, stream.RtspOut)
			stream.Published = true
		}

		if (!rtspStreams[streamName].Started) {
			break
		}

		if (!server.HasClients(streamName)) {
			for len(stream.RtspOut) > 0 {
				<-stream.RtspOut
			}

			continue
		}

		packets := packetizer.H264ToRtp(data.Data)
		for _, p := range packets {
			if (cap(stream.RtspOut) <= len(stream.RtspOut)) {
				log.Println("Rtsp channel is full. Capacity ", cap(stream.RtspOut), " Length ", len(stream.RtspOut), "Skip element")
				<-stream.RtspOut
			}
			stream.RtspOut <- gortsplib.InterleavedFrame{
				Channel: 0,
				Content: p,
			}
		}
	}

	server.DeletePublisher(stream.Name)
	venc.RemoveSubscription(stream.EncoderId, stream.CameraIn)
	delete(rtspStreams, streamName)
}