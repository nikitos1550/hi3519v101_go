package webrtc

import (
    "log"
    "fmt"
    "math/rand"
    "time"

    "net/http"
    "application/pkg/openapi"

    "github.com/pion/webrtc/v2"
    "github.com/pion/webrtc/v2/pkg/media"

    "io/ioutil"
    //"reflect"
)

func init() {
    openapi.AddApiRoute("connectWebrt",   "/webrtc/connect",   "POST",      connectWebrtc)

    Init()
    log.Println("WebRTC inited")
}

func Init() {
    loadTestVideo()
    parseTestVideo()

    go func() {
        // Create a new RTCPeerConnection, to evaluate our sdp in advance
        log.Println("Webrtc: stunning in advance")
        api := webrtc.NewAPI()
        _, err := api.NewPeerConnection(webrtc.Configuration{
            ICEServers: []webrtc.ICEServer{
                {
                    URLs: []string{"stun:stun.l.google.com:19302"},
                },
            },
        })
        if err != nil {
            log.Println("Webrtc: ", err)
            return
        }
        log.Println("Webrtc: stunning in advance DONE")
    }()
}

func connectWebrtc(w http.ResponseWriter, r *http.Request) {
    log.Println("connectWebrtc")

    w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    bodyt, _ := ioutil.ReadAll(r.Body)
    //log.Println("type of body is ", reflect.TypeOf(bodyt))
    //log.Println(bodyt)
    //return

    offer := webrtc.SessionDescription{}
    //Decode(MustReadStdin(), &offer)
    Decode(string(bodyt), &offer)
    //log.Println(offer)

    //return

    // We make our own mediaEngine so we can place the sender's codecs in it.  This because we must use the
    // dynamic media type from the sender in our answer. This is not required if we are the offerer
    mediaEngine := webrtc.MediaEngine{}
    err := mediaEngine.PopulateFromSDP(offer)
    if err != nil {
        panic(err)
    }

    // Search for VP8 Payload type. If the offer doesn't support VP8 exit since
    // since they won't be able to decode anything we send them
    var payloadType uint8
    for _, videoCodec := range mediaEngine.GetCodecsByKind(webrtc.RTPCodecTypeVideo) {
        if videoCodec.Name == "H264" {
            payloadType = videoCodec.PayloadType
            break
        }
    }
    if payloadType == 0 {
        log.Println("Webrtc: Remote peer does not support H264")
        return 
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
        //panic(err)
        log.Println("Webrtc: ", err)
        return
    }

    // Create a video track
    videoTrack, err := peerConnection.NewTrack(payloadType, rand.Uint32(), "video", "pion")
    if err != nil {
        //panic(err)
        log.Println("Webrtc: ", err)
        return
    }
    if _, err = peerConnection.AddTrack(videoTrack); err != nil {
        //panic(err)
        log.Println("Webrtc: ", err)
        return
    }

    /*
    go func() {
        sleepTime := time.Millisecond * 40

        var i int
        i = 1
        for {
            time.Sleep(sleepTime)
            var h264Err error
            if h264Err = videoTrack.WriteSample(media.Sample{Data: getFrameTestVideo(i), Samples: 90000}); h264Err != nil {
                panic(h264Err)
            }
            i++
            if (i>=frames) {
                i=1
            }
        }
    }()
    */

    // Set the handler for ICE connection state
    // This will notify you when the peer has connected/disconnected
    peerConnection.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
        log.Println("Webrtc: connection state has changed ", connectionState.String(), connectionState)
        if (connectionState.String() == "connected") {
            log.Println("Webrtc: Time to start data push (with small 1s delay)!")
            time.Sleep(time.Millisecond * 1000)
            go func() {
                sleepTime := time.Millisecond * 40

                var i int
                i = 1
                for {
                    time.Sleep(sleepTime)
                    var h264Err error
                    if h264Err = videoTrack.WriteSample(media.Sample{Data: getFrameTestVideo(i), Samples: 90000}); h264Err != nil {
                        //panic(h264Err)
                        log.Println("Webrtc: ", h264Err)
                    }
                    i++
                    if (i>=frames) {
                        i=1
                    }
                }
            }()

        }
    })

    // Set the remote SessionDescription
    if err = peerConnection.SetRemoteDescription(offer); err != nil {
        //panic(err)
        log.Println("Webrtc: ", err)
        return
    }

    // Create answer
    answer, err := peerConnection.CreateAnswer(nil)
    if err != nil {
        //panic(err)
        log.Println("Webrtc: ", err)
        return
    }

    // Sets the LocalDescription, and starts our UDP listeners
    if err = peerConnection.SetLocalDescription(answer); err != nil {
        //panic(err)
        log.Println("Webrtc: ", err)
        return
    }

    // Output the answer in base64 so we can paste it in browser
    //fmt.Println(Encode(answer))
    fmt.Fprintf(w, "%s", Encode(answer))
}
