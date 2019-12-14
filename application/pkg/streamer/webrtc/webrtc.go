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
)

func init() {
    openapi.AddRoute("connectWebrt",   "/webrtc/connect",   "POST",      connectWebrtc)
}

func Init() {
    loadTestVideo()
    parseTestVideo()
}

func connectWebrtc(w http.ResponseWriter, r *http.Request) {
    log.Println("connectWebrtc")

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusNotImplemented)

    offer := webrtc.SessionDescription{}
    Decode(MustReadStdin(), &offer)

    log.Println(offer)

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
        panic("Remote peer does not support H264")
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
        panic(err)
    }

    // Create a video track
    videoTrack, err := peerConnection.NewTrack(payloadType, rand.Uint32(), "video", "pion")
    if err != nil {
        panic(err)
    }
    if _, err = peerConnection.AddTrack(videoTrack); err != nil {
        panic(err)
    }

   go func() {
        // Open a IVF file and start reading using our IVFReader
        //file, ivfErr := os.Open("output.ivf")
        //if ivfErr != nil {
        //    panic(ivfErr)
        //}

        //ivf, header, ivfErr := ivfreader.NewWith(file)
        //if ivfErr != nil {
        //    panic(ivfErr)
        //}

        // Send our video file frame at a time. Pace our sending so we send it at the same speed it should be played back as.
        // This isn't required since the video is timestamped, but we will such much higher loss if we send all at once.
        //sleepTime := time.Millisecond * time.Duration((float32(header.TimebaseNumerator)/float32(header.TimebaseDenominator))*1000)
        sleepTime := time.Millisecond * 40


        for {
            log.Printf("%d \n", peerConnection.ConnectionState())
            if (peerConnection.ConnectionState() == 3) {
                log.Println("Starting data send loop")
                break
            }
            time.Sleep(time.Millisecond * 1000)
        }


        var i int
        i = 1
        for {
            //frame, _, ivfErr := ivf.ParseNextFrame()
            //if ivfErr != nil {
            //    panic(ivfErr)
            //}

            time.Sleep(sleepTime)
            //if ivfErr = videoTrack.WriteSample(media.Sample{Data: frame, Samples: 90000}); ivfErr != nil {
            //    panic(ivfErr)
            //}
            var h264Err error
            if h264Err = videoTrack.WriteSample(media.Sample{Data: getFrameTestVideo(i), Samples: 90000}); h264Err != nil {
                panic(h264Err)
            }
            log.Println("frame ", i, " sent")
            i++
            if (i>=frames) {
                i=1
            }
        }
    }()

    // Set the handler for ICE connection state
    // This will notify you when the peer has connected/disconnected
    peerConnection.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
        fmt.Printf("Connection State has changed %s \n", connectionState.String())
    })

    // Set the remote SessionDescription
    if err = peerConnection.SetRemoteDescription(offer); err != nil {
        panic(err)
    }

    // Create answer
    answer, err := peerConnection.CreateAnswer(nil)
    if err != nil {
        panic(err)
    }

    // Sets the LocalDescription, and starts our UDP listeners
    if err = peerConnection.SetLocalDescription(answer); err != nil {
        panic(err)
    }

    // Output the answer in base64 so we can paste it in browser
    fmt.Println(Encode(answer))


}

