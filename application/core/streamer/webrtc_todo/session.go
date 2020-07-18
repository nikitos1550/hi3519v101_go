package webrtc

import (
    "errors"
    "math/rand"
	"strings"
    "sync"

    "github.com/pion/webrtc/v2"
    "github.com/google/uuid"

    "application/core/mpp/frames"
    "application/core/logger"
)

type WebrtcSession struct {
    sync.RWMutex

	SessionId string

	Connected bool

    peerConnection *webrtc.PeerConnection

	videoTrack *webrtc.Track
	//audioTrack *webrtc.Track

    notify  chan frames.FrameItem

    keyframe bool
}

////////////////////////////////////////////////////////////////////////////////

func (s *WebrtcSession) getState() (string, error) {
    if s == nil {
        return "", errors.New("Null pointer")
    }

    s.RLock()
    defer s.RUnlock()

    return s.peerConnection.ConnectionState().String(), nil
}

func (s *WebrtcSession) getConnected() (bool, error) {
    if s == nil {
        return false, errors.New("Null pointer")
    }

    s.RLock()
    defer s.RUnlock()

    return s.Connected, nil
}

func (s *WebrtcSession) setConnected(c bool) error {
    if s == nil {
        return errors.New("Null pointer")
    }

    s.Lock()
    defer s.Unlock()

    s.Connected = c

    return nil
}

////////////////////////////////////////////////////////////////////////////////
func (w *webrtcServer) Connect(sdp string) (*WebrtcSession, string, error) {
    if w == nil {
        return nil, "", errors.New("Null pointer")
    }

    w.RLock()
    defer w.RUnlock()

	var session WebrtcSession
    var err error

	session.SessionId = uuid.New().String()
    //session.Notify = make(chan frames.FrameItem, 5)

    offer := webrtc.SessionDescription{}
    decode(string(sdp), &offer)

    logger.Log.Trace().
        Msg("WebRTC browser sdp decoded")

    // We make our own mediaEngine so we can place the sender's codecs in it.  This because we must use the
    // dynamic media type from the sender in our answer. This is not required if we are the offerer
    mediaEngine := webrtc.MediaEngine{}
    err = mediaEngine.PopulateFromSDP(offer)
    if err != nil {
        return nil, "", errors.New("Failed to populate media server from browser SDP")
    }

    logger.Log.Trace().
        Msg("WebRTC populated from browser sdp")

    // Search for VP8 Payload type. If the offer doesn't support VP8 exit since
    // since they won't be able to decode anything we send them
    var videoPayloadType uint8
    for _, videoCodec := range mediaEngine.GetCodecsByKind(webrtc.RTPCodecTypeVideo) {
        logger.Log.Trace().
            Str("codec", videoCodec.Name).
            Msg("WebRTC check browser SDP")

        if strings.ToLower(videoCodec.Name) == strings.ToLower("h264") {
            videoPayloadType = videoCodec.PayloadType
        }
    }
    if videoPayloadType == 0 {
        return nil, "", errors.New("Remote peer does not support h264")
    }

    logger.Log.Trace().
        Uint("type", uint(videoPayloadType)).
        Msg("WebRTC payload type ok")

    /*
    var audioPayloadType uint8
	for _, audioCodec := range mediaEngine.GetCodecsByKind(webrtc.RTPCodecTypeAudio) {
        if audioCodec.Name == "opus" {
            audioPayloadType = audioCodec.PayloadType
        }
    }
    if audioPayloadType == 0 {
		return -1, "", "Remote peer does not support opus"
    }
    */

    // Create a new RTCPeerConnection
    api := webrtc.NewAPI(webrtc.WithMediaEngine(mediaEngine))

    if false {
        session.peerConnection, err = api.NewPeerConnection(webrtc.Configuration{})
    } else {
        session.peerConnection, err = api.NewPeerConnection(webrtc.Configuration{
            ICEServers: []webrtc.ICEServer{
                {
                    URLs: []string{"stun:stun.l.google.com:19302"},
                },
            },
        })
    }

    if err != nil {
        return nil, "", err //errors.New("Failed to create peer connection")
    }

    logger.Log.Trace().
        Msg("WebRTC peer connection created")

    // Create a video track
    session.videoTrack, err = session.peerConnection.NewTrack(videoPayloadType, rand.Uint32(), "video", "pion")
    if err != nil {
        return nil, "", errors.New("Failed to create video track")
    }

    if _, err = session.peerConnection.AddTrack(session.videoTrack); err != nil {
        return nil, "", errors.New("Failed to add video track")
    }

    /*
    webrtcSession.AudioTrack, err = peerConnection.NewTrack(audioPayloadType, rand.Uint32(), "audio", "pion")
    if err != nil {
		return -1, "", "Failed to create audio track"
    }

    if _, err = peerConnection.AddTrack(webrtcSession.AudioTrack); err != nil {
		return -1, "", "Failed to add audio track"
    }
    */

    session.peerConnection.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
        logger.Log.Trace().
            Str("state", connectionState.String()).
            Str("uuid", session.SessionId).
            Msg("WebRTC state changed")

        if (connectionState == webrtc.ICEConnectionStateConnected) {
		    session.setConnected(true)
        } else if   connectionState == webrtc.ICEConnectionStateFailed ||
                    connectionState == webrtc.ICEConnectionStateDisconnected ||
                    connectionState == webrtc.ICEConnectionStateClosed {

            session.setConnected(false)

			w.Disconnect(&session)

            logger.Log.Trace().
                Str("state", connectionState.String()).
                Str("uuid", session.SessionId).
                Msg("WebRTC client droped")
        }
    })

    if err = session.peerConnection.SetRemoteDescription(offer); err != nil {
        return nil, "", errors.New("Failed to set remote description")
    }

    answer, err := session.peerConnection.CreateAnswer(nil)
    if err != nil {
        return nil, "", errors.New("Failed to create answer")
    }

    if err = session.peerConnection.SetLocalDescription(answer); err != nil {
        return nil, "", errors.New("Failed to set local description")
    }

    w.clientsMutex.Lock()
    w.clients[session.SessionId] = &session
    w.clientsMutex.Unlock()

    return &session, encode(answer), nil
}

func (w *webrtcServer) Disconnect(s *WebrtcSession) error {
    if w == nil {
        return errors.New("Null pointer")
    }

    w.RLock()
    defer w.RUnlock()

    w.clientsMutex.Lock()
    defer w.clientsMutex.Unlock()

    guid := ""
    for key, value := range(w.clients) {
        if s == value {
            guid = key
            break
        }
    }

    if guid == "" {
        return errors.New("Client is not in list")
    }

    err := s.peerConnection.Close()
    if err != nil {
        logger.Log.Warn().
            Str("reson", err.Error()).
            Msg("WebRTC disconnect")
    }

    delete(w.clients, guid)

    return nil
}

func (w *webrtcServer) GetClientByUUID(guid string) (*WebrtcSession, error) {
    if w == nil {
        return nil, errors.New("Null pointer")
    }

    w.RLock()
    defer w.RUnlock()

    w.clientsMutex.Lock()
    defer w.clientsMutex.Unlock()

    var item *WebrtcSession
    for key, value := range(w.clients) {
        if guid == key {
            item = value
            break
        }
    }

    if item == nil {
        return nil, errors.New("Client is not in list")
    }

    return item, nil
}
