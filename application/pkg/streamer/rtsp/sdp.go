//+build streamerRtsp

package rtsp

import (
    "encoding/base64"
    "fmt"
    "math/rand"
)

var sdpPattern =
`v=0
o=- %d 1 IN IP4 127.0.0.1
s=Session streamed by "hisi cam"
i=%s
t=0 0
a=tool:hisi cam streaming server
a=type:broadcast
a=control:*
a=range:npt=0-
a=x-qt-text-nam:Session streamed by "hisi cam"
a=x-qt-text-inf:%s
m=video 0 RTP/AVP %d
c=IN IP4 127.0.0.1
a=rtpmap:%d %s/%d
a=fmtp:%d packetization-mode=1;profile-level-id=%s;sprop-parameter-sets=%s,%s
a=control:track1`

func getSessionId() int{
    return rand.Int()
}

func getPayloadType(encoder string) int{
    if (encoder == "h265"){
        return 96
    }
    return 96
}

func getPayloadFormat(encoder string) string {
    if (encoder == "h265"){
        return "H265"
    }
    return "H264"
}

func getTimestampFrequency() int{
    return 90000
}

func getProfileLevelId(sps []byte) string {
    return fmt.Sprintf("%x%x%x",
        sps[1],
        sps[2],
        sps[3])
}

func CreateSdp(encoder string, streamName string, sps []byte, pps []byte) string {
    return fmt.Sprintf(sdpPattern,
        getSessionId(),
        streamName,
        streamName,
        getPayloadType(encoder),
        getPayloadType(encoder),
        getPayloadFormat(encoder),
        getTimestampFrequency(),
        getPayloadType(encoder),
        getProfileLevelId(sps),
        base64.StdEncoding.EncodeToString(sps),
        base64.StdEncoding.EncodeToString(pps))
}
