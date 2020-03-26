//+build streamerRtsp

package rtsp

import (
	"application/pkg/mpp/venc"
	"application/pkg/openapi"
	"log"
	"net/http"

	"github.com/aler9/gortsplib"
)

type responseRecord struct {
	Message string
}

type rtspStream struct {
	Name string
	EncoderId string
	EncoderType string
	Started bool
	SendDataStarted bool
	Published bool
	Sps []byte
	Pps []byte
	CameraIn chan []byte
	RtspOut chan gortsplib.InterleavedFrame
}

type rtspInfo struct {
	Name string
	EncoderId string
}

var (
	server *program
	rtspStreams map[string] rtspStream
)

func init() {
	rtspStreams = make(map[string] rtspStream)

	server = CreateRtspServer()

    openapi.AddApiRoute("rtspApiDescription", "/rtsp", "GET", rtspApiDescription)
    openapi.AddApiRoute("startRtspStream", "/rtsp/start", "GET", startRtspStream)
    openapi.AddApiRoute("stopRtspStream", "/rtsp/stop", "GET", stopRtspStream)
    openapi.AddApiRoute("listRtspStreams", "/rtsp/list", "GET", listRtspStreams)
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

	encoder, encoderExists := venc.Encoders[encoderId]
	if (!encoderExists) {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{Message: "Failed to find encoder  " + encoderId})
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
		SendDataStarted: false,
		Published: false,
		Sps: []byte{},
		Pps: []byte{},
		CameraIn: make(chan []byte, 100),
		RtspOut: make(chan gortsplib.InterleavedFrame, 10),
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
			stream.Sps = ExtractSps(stream.EncoderType, data)
		}

		if (len(stream.Pps) == 0) {
			stream.Pps = ExtractPps(stream.EncoderType, data)
		}

		if (!stream.Published && len(stream.Sps) > 0 && len(stream.Pps) > 0) {
			sdp := CreateSdp(stream.EncoderType, stream.Name, stream.Sps, stream.Pps)
			server.AddPublisher(sdp, stream.Name, stream.RtspOut)
			stream.Published = true
		}

		if (!stream.Started) {
			log.Println("//////////////Stopped stream")
			break
		}

		if (!stream.SendDataStarted) {
			if (len(ExtractSps(stream.EncoderType, data))) <= 0 {
				continue
			}
		}

		stream.SendDataStarted = true

		packets := packetizer.NalH264ToRtp(data)
		for _, p := range packets {
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