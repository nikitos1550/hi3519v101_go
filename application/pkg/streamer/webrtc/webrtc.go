//+build streamerWebrtc

package webrtc

import (
    "log"
    "math/rand"
	"strings"
	"strconv"

    "application/pkg/mpp/ai"
    "application/pkg/mpp/venc"
	"application/pkg/streamer/rtsp"

    "github.com/pion/webrtc/v2"
    "github.com/pion/webrtc/v2/pkg/media"

	"github.com/google/uuid"
)

type WebrtcSession struct {
	SessionId string
	EncoderId int
	EncoderType string
	Started bool
	VideoTrack *webrtc.Track
	AudioTrack *webrtc.Track
    //VideoPayload chan []byte
    VideoPayload chan venc.ChannelEncoder
    AudioPayload chan []byte
}

var (
	WebrtcSessions map[string] WebrtcSession
)

func WebrtcInit() {
	WebrtcSessions = make(map[string] WebrtcSession)
    // Create a new RTCPeerConnection, to evaluate our sdp in advance
    log.Println("Webrtc: stunning in advance")
    //api := webrtc.NewAPI()
    /*
    _, err := api.NewPeerConnection(webrtc.Configuration{
        ICEServers: []webrtc.ICEServer{
            {
                URLs: []string{}, //"stun:stun.l.google.com:19302"},
            },
        },
    })
    if err != nil {
        log.Println("Webrtc: ", err)
        return
    }
    */
    log.Println("Webrtc: stunning in advance DONE")
}

func WebrtcConnect(browserSdp string, encoderId int) (int, string, string) {
	var webrtcSession WebrtcSession
	webrtcSession.SessionId = uuid.New().String()
	webrtcSession.EncoderId = encoderId
	//webrtcSession.VideoPayload = make(chan []byte, 1)
    webrtcSession.VideoPayload = make(chan venc.ChannelEncoder, 1)
	webrtcSession.AudioPayload = make(chan []byte, 1)

	encoder, encoderExists := venc.ActiveEncoders[encoderId]
	if (!encoderExists) {
		return -1, "", "Failed to find encoder  " + strconv.Itoa(encoderId)
	}

	webrtcSession.EncoderType = encoder.Format
	webrtcSession.Started = true
	

    offer := webrtc.SessionDescription{}
    Decode(string(browserSdp), &offer)

    // We make our own mediaEngine so we can place the sender's codecs in it.  This because we must use the
    // dynamic media type from the sender in our answer. This is not required if we are the offerer
    mediaEngine := webrtc.MediaEngine{}
    err := mediaEngine.PopulateFromSDP(offer)
    if err != nil {
		return -1, "", "Failed to populate media server from browser SDP"
    }

    // Search for VP8 Payload type. If the offer doesn't support VP8 exit since
    // since they won't be able to decode anything we send them
    var videoPayloadType uint8
    for _, videoCodec := range mediaEngine.GetCodecsByKind(webrtc.RTPCodecTypeVideo) {
        if strings.ToLower(videoCodec.Name) == strings.ToLower(webrtcSession.EncoderType) {
            videoPayloadType = videoCodec.PayloadType
        }
    }
    if videoPayloadType == 0 {
		return -1, "", "Remote peer does not support " + webrtcSession.EncoderType
    }

    var audioPayloadType uint8
	for _, audioCodec := range mediaEngine.GetCodecsByKind(webrtc.RTPCodecTypeAudio) {
        if audioCodec.Name == "opus" {
            audioPayloadType = audioCodec.PayloadType
        }
    }
    if audioPayloadType == 0 {
		return -1, "", "Remote peer does not support opus"
    }

    // Create a new RTCPeerConnection
    api := webrtc.NewAPI(webrtc.WithMediaEngine(mediaEngine))
    peerConnection, err := api.NewPeerConnection(webrtc.Configuration{
        ICEServers: []webrtc.ICEServer{
            {
                URLs: []string{"stun:stun.l.google.com:19302"},
            },
        },
    })

    if err != nil {
		return -1, "", "Failed to create peer connection"
    }

    // Create a video track
    webrtcSession.VideoTrack, err = peerConnection.NewTrack(videoPayloadType, rand.Uint32(), "video", "pion")
    if err != nil {
		return -1, "", "Failed to create video track"
    }

    if _, err = peerConnection.AddTrack(webrtcSession.VideoTrack); err != nil {
		return -1, "", "Failed to add video track"
    }

    webrtcSession.AudioTrack, err = peerConnection.NewTrack(audioPayloadType, rand.Uint32(), "audio", "pion")
    if err != nil {
		return -1, "", "Failed to create audio track"
    }

    if _, err = peerConnection.AddTrack(webrtcSession.AudioTrack); err != nil {
		return -1, "", "Failed to add audio track"
    }

    // Set the handler for ICE connection state
    // This will notify you when the peer has connected/disconnected
    peerConnection.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
        log.Println("Webrtc: connection state has changed ", connectionState.String(), connectionState)
        if (connectionState.String() == "connected") {
			go func() {
				sendVideoData(webrtcSession.SessionId)
			}()

			go func() {
				sendAudioData(webrtcSession.SessionId)
			}()
        } else if (connectionState.String() == "failed" || connectionState.String() == "disconnected") {
			WebrtcDisconnect(webrtcSession.SessionId)
        } 
    })

    // Set the remote SessionDescription
    if err = peerConnection.SetRemoteDescription(offer); err != nil {
		return -1, "", "Failed to set remote description"
    }

    // Create answer
    answer, err := peerConnection.CreateAnswer(nil)
    if err != nil {
		return -1, "", "Failed to create answer"
    }

    // Sets the LocalDescription, and starts our UDP listeners
    if err = peerConnection.SetLocalDescription(answer); err != nil {
		return -1, "", "Failed to set local description"
    }

	venc.SubsribeEncoder(webrtcSession.EncoderId, webrtcSession.VideoPayload)
	ai.SubsribeOpus(webrtcSession.AudioPayload)

	WebrtcSessions[webrtcSession.SessionId] = webrtcSession
	return 0, webrtcSession.SessionId, Encode(answer)
}

func sendVideoData(sessionId string) {
    spsSended := false
    for {
		session, exists := WebrtcSessions[sessionId]
		if (!exists) {
			log.Println("Webrtc session not found")
			break
		}

		if (!session.Started) {
			venc.RemoveSubscription(session.EncoderId, session.VideoPayload)
			delete(WebrtcSessions, sessionId)
			break
		}

		buf := <- session.VideoPayload

		if (!spsSended){
			sps := rtsp.ExtractSps(session.EncoderType, buf.Data);
			if (len(sps) == 0){
				continue
			}
			spsSended = true
		}

        var h264Err error
        if h264Err = session.VideoTrack.WriteSample(media.Sample{Data: buf.Data, Samples: 90000}); h264Err != nil {
            log.Println("Webrtc video: ", h264Err)
        }
    }
}

func sendAudioData(sessionId string) {
    for {
		session, exists := WebrtcSessions[sessionId]
		if (!exists) {
			log.Println("Webrtc session not found")
			break
		}

		if (!session.Started) {
			ai.RemoveOpus(session.AudioPayload)
			break
		}

		buf := <- session.AudioPayload

        var err error
        if err = session.AudioTrack.WriteSample(media.Sample{Data: buf, Samples: 960}); err != nil {
            log.Println("Webrtc audio: ", err)
        }
    }
}

func WebrtcDisconnect(sessionId string) (int, string) {
	session, exists := WebrtcSessions[sessionId]
	if (!exists) {
		return -1, "Webrtc session not found " + sessionId
    }

	session.Started = false
	WebrtcSessions[session.SessionId] = session
	return 0,""
}
