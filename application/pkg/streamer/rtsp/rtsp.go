//+build streamerRtsp

package rtsp

func Init() {
	server := CreateRtspServer()
	sdp :=
`v=0
o=- 1584621403523799 1 IN IP4 10.104.44.103
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

	cameraPackets := make(chan gortsplib.InterleavedFrame)
	server.AddPublisher(sdp, "stream", cameraPackets)
}
