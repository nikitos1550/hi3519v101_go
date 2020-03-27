//+build streamerRtsp

package rtsp

import (
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/aler9/gortsplib"
)

var Version string = "v0.0.0"

const (
	_READ_TIMEOUT  = 5 * time.Second
	_WRITE_TIMEOUT = 5 * time.Second
)

type trackFlow int

const (
	_TRACK_FLOW_RTP trackFlow = iota
	_TRACK_FLOW_RTCP
)

type track struct {
	rtpPort  int
	rtcpPort int
}

type streamProtocol int

const (
	_STREAM_PROTOCOL_UDP streamProtocol = iota
	_STREAM_PROTOCOL_TCP
)

func (s streamProtocol) String() string {
	if s == _STREAM_PROTOCOL_UDP {
		return "udp"
	}
	return "tcp"
}

type program struct {
	protocols   map[streamProtocol]struct{}
	rtspPort    int
	rtpPort     int
	rtcpPort    int
	publishUser string
	publishPass string
	preScript   string
	postScript  string
	mutex       sync.RWMutex
	rtspl       *serverTcpListener
	rtpl        *serverUdpListener
	rtcpl       *serverUdpListener
	clients     map[*client]struct{}
	publishers  map[string]*client
}

func newProgram() (*program, error) {
	version := false
	protocolsStr := "udp,tcp"
	rtspPort := 8554
	rtpPort := 8000
	rtcpPort := 8001
	publishUser := ""
	publishPass := ""
	preScript := ""
	postScript := ""

	if version == true {
		fmt.Println("rtsp-simple-server " + Version)
		os.Exit(0)
	}

	if rtspPort == 0 {
		return nil, fmt.Errorf("rtsp port not provided")
	}

	if rtpPort == 0 {
		return nil, fmt.Errorf("rtp port not provided")
	}

	if rtcpPort == 0 {
		return nil, fmt.Errorf("rtcp port not provided")
	}

	if (rtpPort % 2) != 0 {
		return nil, fmt.Errorf("rtp port must be even")
	}

	if rtcpPort != (rtpPort + 1) {
		return nil, fmt.Errorf("rtcp port must be rtp port plus 1")
	}

	protocols := make(map[streamProtocol]struct{})
	for _, proto := range strings.Split(protocolsStr, ",") {
		switch proto {
		case "udp":
			protocols[_STREAM_PROTOCOL_UDP] = struct{}{}

		case "tcp":
			protocols[_STREAM_PROTOCOL_TCP] = struct{}{}

		default:
			return nil, fmt.Errorf("unsupported protocol: %s", proto)
		}
	}
	if len(protocols) == 0 {
		return nil, fmt.Errorf("no protocols provided")
	}

	if publishUser != "" {
		if !regexp.MustCompile("^[a-zA-Z0-9]+$").MatchString(publishUser) {
			return nil, fmt.Errorf("publish username must be alphanumeric")
		}
	}

	if publishPass != "" {
		if !regexp.MustCompile("^[a-zA-Z0-9]+$").MatchString(publishPass) {
			return nil, fmt.Errorf("publish password must be alphanumeric")
		}
	}

	if publishUser != "" && publishPass == "" || publishUser == "" && publishPass != "" {
		return nil, fmt.Errorf("publish username and password must be both filled")
	}

	log.Printf("rtsp-simple-server %s", Version)

	p := &program{
		protocols:   protocols,
		rtspPort:    rtspPort,
		rtpPort:     rtpPort,
		rtcpPort:    rtcpPort,
		publishUser: publishUser,
		publishPass: publishPass,
		preScript:   preScript,
		postScript:  postScript,
		clients:     make(map[*client]struct{}),
		publishers:  make(map[string]*client),
	}

	var err error

	p.rtpl, err = newServerUdpListener(p, rtpPort, _TRACK_FLOW_RTP)
	if err != nil {
		return nil, err
	}

	p.rtcpl, err = newServerUdpListener(p, rtcpPort, _TRACK_FLOW_RTCP)
	if err != nil {
		return nil, err
	}

	p.rtspl, err = newServerTcpListener(p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (p *program) run() {
	go p.rtpl.run()
	go p.rtcpl.run()
	go p.rtspl.run()
}

func (p *program) forwardTrack(path string, id int, flow trackFlow, frame []byte) {
	for c := range p.clients {
		if c.path == path && c.state == _CLIENT_STATE_PLAY {
			if c.streamProtocol == _STREAM_PROTOCOL_UDP {
				if flow == _TRACK_FLOW_RTP {
					p.rtpl.chanWrite <- &udpWrite{
						addr: &net.UDPAddr{
							IP:   c.ip,
							Port: c.streamTracks[id].rtpPort,
						},
						buf: frame,
					}
				} else {
					p.rtcpl.chanWrite <- &udpWrite{
						addr: &net.UDPAddr{
							IP:   c.ip,
							Port: c.streamTracks[id].rtcpPort,
						},
						buf: frame,
					}
				}

			} else {
				c.chanWrite <- &gortsplib.InterleavedFrame{
					Channel: trackToInterleavedChannel(id, flow),
					Content: frame,
				}
			}
		}
	}
}

func (p *program) AddPublisher(sdp string, streamPath string, ch chan gortsplib.InterleavedFrame) bool {
	sdpParsed, err := sdpParse([]byte(sdp))
	if err != nil {
		log.Println(fmt.Errorf("invalid SDP: %s", err))
		return false
	}

	sdpParsed, _ = sdpFilter(sdpParsed, []byte(sdp))

	c := &client{
		p: p,
		path: streamPath,
		streamSdpText: []byte(sdp),
		streamSdpParsed: sdpParsed,
		chanWrite: make(chan *gortsplib.InterleavedFrame),
		cameraPackets: ch,
		clientsCount: 0,
		started: true,
	}

	p.publishers[streamPath] = c
	return true
}

func (p *program) DeletePublisher(streamPath string) bool {
	stream, streamExists := p.publishers[streamPath]
	if (!streamExists) {
		return false
	}

	stream.started = false
	delete(p.publishers, streamPath)
	return true
}

func (p *program) HasClients(streamPath string) bool {
	stream, streamExists := p.publishers[streamPath]
	if (!streamExists) {
		return false
	}

	return stream.clientsCount > 0
}

func CreateRtspServer() *program {
	p, err := newProgram()
	if err != nil {
		log.Fatal("ERR: ", err)
	}
	p.run()
	return p
}

