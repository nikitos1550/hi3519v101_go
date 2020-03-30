//+build streamerRtsp

package rtsp

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/url"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/aler9/gortsplib"
	"gortc.io/sdp"
)

func sdpParse(in []byte) (*sdp.Message, error) {
	s, err := sdp.DecodeSession(in, nil)
	if err != nil {
		return nil, err
	}

	m := &sdp.Message{}
	d := sdp.NewDecoder(s)
	err = d.Decode(m)
	if err != nil {
		return nil, err
	}

	if len(m.Medias) == 0 {
		return nil, fmt.Errorf("no tracks defined in SDP")
	}

	return m, nil
}

// remove everything from SDP except the bare minimum
func sdpFilter(msgIn *sdp.Message, byteIn []byte) (*sdp.Message, []byte) {
	msgOut := &sdp.Message{}

	msgOut.Name = "Stream"
	msgOut.Origin = sdp.Origin{
		Username:    "-",
		NetworkType: "IN",
		AddressType: "IP4",
		Address:     "127.0.0.1",
	}

	for i, m := range msgIn.Medias {
		var attributes []sdp.Attribute
		for _, attr := range m.Attributes {
			if attr.Key == "rtpmap" || attr.Key == "fmtp" {
				attributes = append(attributes, attr)
			}
		}
		// control attribute is needed by gstreamer
		attributes = append(attributes, sdp.Attribute{
			Key:   "control",
			Value: "streamid=" + strconv.FormatInt(int64(i), 10),
		})

		msgOut.Medias = append(msgOut.Medias, sdp.Media{
			Bandwidths: m.Bandwidths,
			Description: sdp.MediaDescription{
				Type:     m.Description.Type,
				Protocol: m.Description.Protocol,
				Formats:  m.Description.Formats,
			},
			Attributes: attributes,
		})
	}

	sdps := sdp.Session{}
	sdps = msgOut.Append(sdps)
	byteOut := sdps.AppendTo(nil)

	return msgOut, byteOut
}

func interleavedChannelToTrack(channel uint8) (int, trackFlow) {
	if (channel % 2) == 0 {
		return int(channel / 2), _TRACK_FLOW_RTP
	}
	return int((channel - 1) / 2), _TRACK_FLOW_RTCP
}

func trackToInterleavedChannel(id int, flow trackFlow) uint8 {
	if flow == _TRACK_FLOW_RTP {
		return uint8(id * 2)
	}
	return uint8((id * 2) + 1)
}

type clientState int

const (
	_CLIENT_STATE_STARTING clientState = iota
	_CLIENT_STATE_ANNOUNCE
	_CLIENT_STATE_PRE_PLAY
	_CLIENT_STATE_PLAY
	_CLIENT_STATE_PRE_RECORD
	_CLIENT_STATE_RECORD
)

type client struct {
	p               *program
	conn            *gortsplib.ConnServer
	state           clientState
	ip              net.IP
	path            string
	as              *gortsplib.AuthServer
	streamSdpText   []byte       // filled only if publisher
	streamSdpParsed *sdp.Message // filled only if publisher
	streamProtocol  streamProtocol
	streamTracks    []*track
	chanWrite       chan *gortsplib.InterleavedFrame
	cameraPackets   chan gortsplib.InterleavedFrame
	clientsCount    int
	started         bool
}

func newClient(p *program, nconn net.Conn) *client {
	c := &client{
		p:         p,
		conn:      gortsplib.NewConnServer(nconn),
		state:     _CLIENT_STATE_STARTING,
		chanWrite: make(chan *gortsplib.InterleavedFrame),
		cameraPackets: make(chan gortsplib.InterleavedFrame),
		started: true,
	}

	c.p.mutex.Lock()
	c.p.clients[c] = struct{}{}
	c.p.mutex.Unlock()

	return c
}

func (c *client) close() error {
	// already deleted
	if _, ok := c.p.clients[c]; !ok {
		return nil
	}

	delete(c.p.clients, c)
	c.conn.NetConn().Close()
	close(c.chanWrite)

	if c.path != "" {
		if pub, ok := c.p.publishers[c.path]; ok && pub == c {
			delete(c.p.publishers, c.path)

			// if the publisher has disconnected
			// close all other connections that share the same path
			for oc := range c.p.clients {
				if oc.path == c.path {
					oc.close()
				}
			}
		}
	}
	return nil
}

func (c *client) log(format string, args ...interface{}) {
	format = "[RTSP client " + c.conn.NetConn().RemoteAddr().String() + "] " + format
	log.Printf(format, args...)
}

func (c *client) run() {
	defer func() {
		if c.p.postScript != "" {
			postScript := exec.Command(c.p.postScript)
			err := postScript.Run()
			if err != nil {
				c.log("ERR: %s", err)
			}
		}
	}()

	defer c.log("disconnected")

	defer func() {
		c.p.mutex.Lock()
		defer c.p.mutex.Unlock()
		c.close()
	}()

	ipstr, _, _ := net.SplitHostPort(c.conn.NetConn().RemoteAddr().String())
	c.ip = net.ParseIP(ipstr)

	c.log("connected")

	if c.p.preScript != "" {
		preScript := exec.Command(c.p.preScript)
		err := preScript.Run()
		if err != nil {
			c.log("ERR: %s", err)
		}
	}

	for {
		req, err := c.conn.ReadRequest()
		if err != nil {
			if err != io.EOF {
				c.log("ERR: %s", err)
			}
			return
		}

		ok := c.handleRequest(req)
		if !ok {
			return
		}
	}
}

func (c *client) writeResDeadline(res *gortsplib.Response) {
	c.conn.NetConn().SetWriteDeadline(time.Now().Add(_WRITE_TIMEOUT))
	c.conn.WriteResponse(res)
}

func (c *client) writeResError(req *gortsplib.Request, err error) {
	c.log("ERR: %s", err)

	if cseq, ok := req.Header["CSeq"]; ok && len(cseq) == 1 {
		c.writeResDeadline(&gortsplib.Response{
			StatusCode: 400,
			Status:     "Bad Request",
			Header: gortsplib.Header{
				"CSeq": []string{cseq[0]},
			},
		})
	} else {
		c.writeResDeadline(&gortsplib.Response{
			StatusCode: 400,
			Status:     "Bad Request",
		})
	}
}

func (c *client) handleRequest(req *gortsplib.Request) bool {
	c.log(req.Method)

	cseq, ok := req.Header["CSeq"]
	if !ok || len(cseq) != 1 {
		c.writeResError(req, fmt.Errorf("cseq missing"))
		return false
	}

	ur, err := url.Parse(req.Url)
	if err != nil {
		c.writeResError(req, fmt.Errorf("unable to parse path '%s'", req.Url))
		return false
	}

	path := func() string {
		ret := ur.Path

		// remove leading slash
		if len(ret) > 1 {
			ret = ret[1:]
		}

		// strip any subpath
		if n := strings.Index(ret, "/"); n >= 0 {
			ret = ret[:n]
		}

		return ret
	}()

	switch req.Method {
	case "OPTIONS":
		// do not check state, since OPTIONS can be requested
		// in any state

		c.writeResDeadline(&gortsplib.Response{
			StatusCode: 200,
			Status:     "OK",
			Header: gortsplib.Header{
				"CSeq": []string{cseq[0]},
				"Public": []string{strings.Join([]string{
					"DESCRIBE",
					"ANNOUNCE",
					"SETUP",
					"PLAY",
					"PAUSE",
					"RECORD",
					"TEARDOWN",
				}, ", ")},
			},
		})
		return true

	case "DESCRIBE":
		if c.state != _CLIENT_STATE_STARTING {
			c.writeResError(req, fmt.Errorf("client is in state '%d'", c.state))
			return false
		}

		sdp, err := func() ([]byte, error) {
			c.p.mutex.RLock()
			defer c.p.mutex.RUnlock()

			pub, ok := c.p.publishers[path]
			if !ok {
				return nil, fmt.Errorf("no one is streaming on path '%s'", path)
			}

			return pub.streamSdpText, nil
		}()
		if err != nil {
			c.writeResError(req, err)
			return false
		}

		c.writeResDeadline(&gortsplib.Response{
			StatusCode: 200,
			Status:     "OK",
			Header: gortsplib.Header{
				"CSeq":         []string{cseq[0]},
				"Content-Base": []string{req.Url},
				"Content-Type": []string{"application/sdp"},
			},
			Content: sdp,
		})
		return true

	case "ANNOUNCE":
		if c.state != _CLIENT_STATE_STARTING {
			c.writeResError(req, fmt.Errorf("client is in state '%d'", c.state))
			return false
		}

		if c.p.publishUser != "" {
			initialRequest := false
			if c.as == nil {
				initialRequest = true
				c.as = gortsplib.NewAuthServer(c.p.publishUser, c.p.publishPass)
			}

			err := c.as.ValidateHeader(req.Header["Authorization"], "ANNOUNCE", req.Url)
			if err != nil {
				if !initialRequest {
					c.log("ERR: Unauthorized: %s", err)
				}

				c.writeResDeadline(&gortsplib.Response{
					StatusCode: 401,
					Status:     "Unauthorized",
					Header: gortsplib.Header{
						"CSeq":             []string{cseq[0]},
						"WWW-Authenticate": c.as.GenerateHeader(),
					},
				})

				if !initialRequest {
					return false
				}

				return true
			}
		}

		ct, ok := req.Header["Content-Type"]
		if !ok || len(ct) != 1 {
			c.writeResError(req, fmt.Errorf("Content-Type header missing"))
			return false
		}

		if ct[0] != "application/sdp" {
			c.writeResError(req, fmt.Errorf("unsupported Content-Type '%s'", ct))
			return false
		}

		sdpParsed, err := sdpParse(req.Content)
		if err != nil {
			c.writeResError(req, fmt.Errorf("invalid SDP: %s", err))
			return false
		}

		sdpParsed, req.Content = sdpFilter(sdpParsed, req.Content)

		err = func() error {
			c.p.mutex.Lock()
			defer c.p.mutex.Unlock()

			_, ok := c.p.publishers[path]
			if ok {
				return fmt.Errorf("another client is already publishing on path '%s'", path)
			}

			c.path = path
			c.p.publishers[path] = c
			c.streamSdpText = req.Content
			c.streamSdpParsed = sdpParsed
			c.state = _CLIENT_STATE_ANNOUNCE
			return nil
		}()
		if err != nil {
			c.writeResError(req, err)
			return false
		}

		c.writeResDeadline(&gortsplib.Response{
			StatusCode: 200,
			Status:     "OK",
			Header: gortsplib.Header{
				"CSeq": []string{cseq[0]},
			},
		})
		return true

	case "SETUP":
		tsRaw, ok := req.Header["Transport"]
		if !ok || len(tsRaw) != 1 {
			c.writeResError(req, fmt.Errorf("transport header missing"))
			return false
		}

		th := gortsplib.ReadHeaderTransport(tsRaw[0])

		if _, ok := th["unicast"]; !ok {
			c.writeResError(req, fmt.Errorf("transport header does not contain unicast"))
			return false
		}

		switch c.state {
		// play
		case _CLIENT_STATE_STARTING, _CLIENT_STATE_PRE_PLAY:
			// play via UDP
			if func() bool {
				_, ok := th["RTP/AVP"]
				if ok {
					return true
				}
				_, ok = th["RTP/AVP/UDP"]
				if ok {
					return true
				}
				return false
			}() {
				if _, ok := c.p.protocols[_STREAM_PROTOCOL_UDP]; !ok {
					c.log("ERR: udp streaming is disabled")
					c.writeResDeadline(&gortsplib.Response{
						StatusCode: 461,
						Status:     "Unsupported Transport",
						Header: gortsplib.Header{
							"CSeq": []string{cseq[0]},
						},
					})
					return false
				}

				rtpPort, rtcpPort := th.GetPorts("client_port")
				if rtpPort == 0 || rtcpPort == 0 {
					c.writeResError(req, fmt.Errorf("transport header does not have valid client ports (%s)", tsRaw[0]))
					return false
				}

				if c.path != "" && path != c.path {
					c.writeResError(req, fmt.Errorf("path has changed"))
					return false
				}

				err = func() error {
					c.p.mutex.Lock()
					defer c.p.mutex.Unlock()

					pub, ok := c.p.publishers[path]
					if !ok {
						return fmt.Errorf("no one is streaming on path '%s'", path)
					}

					if len(c.streamTracks) > 0 && c.streamProtocol != _STREAM_PROTOCOL_UDP {
						return fmt.Errorf("client want to send tracks with different protocols")
					}

					if len(c.streamTracks) >= len(pub.streamSdpParsed.Medias) {
						return fmt.Errorf("all the tracks have already been setup")
					}

					c.path = path
					c.streamProtocol = _STREAM_PROTOCOL_UDP
					c.streamTracks = append(c.streamTracks, &track{
						rtpPort:  rtpPort,
						rtcpPort: rtcpPort,
					})

					c.state = _CLIENT_STATE_PRE_PLAY
					return nil
				}()
				if err != nil {
					c.writeResError(req, err)
					return false
				}

				c.writeResDeadline(&gortsplib.Response{
					StatusCode: 200,
					Status:     "OK",
					Header: gortsplib.Header{
						"CSeq": []string{cseq[0]},
						"Transport": []string{strings.Join([]string{
							"RTP/AVP/UDP",
							"unicast",
							fmt.Sprintf("client_port=%d-%d", rtpPort, rtcpPort),
							fmt.Sprintf("server_port=%d-%d", c.p.rtpPort, c.p.rtcpPort),
						}, ";")},
						"Session": []string{"12345678"},
					},
				})
				return true

				// play via TCP
			} else if _, ok := th["RTP/AVP/TCP"]; ok {
				if _, ok := c.p.protocols[_STREAM_PROTOCOL_TCP]; !ok {
					c.log("ERR: tcp streaming is disabled")
					c.writeResDeadline(&gortsplib.Response{
						StatusCode: 461,
						Status:     "Unsupported Transport",
						Header: gortsplib.Header{
							"CSeq": []string{cseq[0]},
						},
					})
					return false
				}

				if c.path != "" && path != c.path {
					c.writeResError(req, fmt.Errorf("path has changed"))
					return false
				}

				err = func() error {
					c.p.mutex.Lock()
					defer c.p.mutex.Unlock()

					pub, ok := c.p.publishers[path]
					if !ok {
						return fmt.Errorf("no one is streaming on path '%s'", path)
					}

					if len(c.streamTracks) > 0 && c.streamProtocol != _STREAM_PROTOCOL_TCP {
						return fmt.Errorf("client want to send tracks with different protocols")
					}

					if len(c.streamTracks) >= len(pub.streamSdpParsed.Medias) {
						return fmt.Errorf("all the tracks have already been setup")
					}

					c.path = path
					c.streamProtocol = _STREAM_PROTOCOL_TCP
					c.streamTracks = append(c.streamTracks, &track{
						rtpPort:  0,
						rtcpPort: 0,
					})

					c.state = _CLIENT_STATE_PRE_PLAY
					return nil
				}()
				if err != nil {
					c.writeResError(req, err)
					return false
				}

				interleaved := fmt.Sprintf("%d-%d", ((len(c.streamTracks) - 1) * 2), ((len(c.streamTracks)-1)*2)+1)

				c.writeResDeadline(&gortsplib.Response{
					StatusCode: 200,
					Status:     "OK",
					Header: gortsplib.Header{
						"CSeq": []string{cseq[0]},
						"Transport": []string{strings.Join([]string{
							"RTP/AVP/TCP",
							"unicast",
							fmt.Sprintf("interleaved=%s", interleaved),
						}, ";")},
						"Session": []string{"12345678"},
					},
				})
				return true

			} else {
				c.writeResError(req, fmt.Errorf("transport header does not contain a valid protocol (RTP/AVP, RTP/AVP/UDP or RTP/AVP/TCP) (%s)", tsRaw[0]))
				return false
			}

		// record
		case _CLIENT_STATE_ANNOUNCE, _CLIENT_STATE_PRE_RECORD:
			if _, ok := th["mode=record"]; !ok {
				c.writeResError(req, fmt.Errorf("transport header does not contain mode=record"))
				return false
			}

			if path != c.path {
				c.writeResError(req, fmt.Errorf("path has changed"))
				return false
			}

			// record via UDP
			if func() bool {
				_, ok := th["RTP/AVP"]
				if ok {
					return true
				}
				_, ok = th["RTP/AVP/UDP"]
				if ok {
					return true
				}
				return false
			}() {
				if _, ok := c.p.protocols[_STREAM_PROTOCOL_UDP]; !ok {
					c.log("ERR: udp streaming is disabled")
					c.writeResDeadline(&gortsplib.Response{
						StatusCode: 461,
						Status:     "Unsupported Transport",
						Header: gortsplib.Header{
							"CSeq": []string{cseq[0]},
						},
					})
					return false
				}

				rtpPort, rtcpPort := th.GetPorts("client_port")
				if rtpPort == 0 || rtcpPort == 0 {
					c.writeResError(req, fmt.Errorf("transport header does not have valid client ports (%s)", tsRaw[0]))
					return false
				}

				err = func() error {
					c.p.mutex.Lock()
					defer c.p.mutex.Unlock()

					if len(c.streamTracks) > 0 && c.streamProtocol != _STREAM_PROTOCOL_UDP {
						return fmt.Errorf("client want to send tracks with different protocols")
					}

					if len(c.streamTracks) >= len(c.streamSdpParsed.Medias) {
						return fmt.Errorf("all the tracks have already been setup")
					}

					c.streamProtocol = _STREAM_PROTOCOL_UDP
					c.streamTracks = append(c.streamTracks, &track{
						rtpPort:  rtpPort,
						rtcpPort: rtcpPort,
					})

					c.state = _CLIENT_STATE_PRE_RECORD
					return nil
				}()
				if err != nil {
					c.writeResError(req, err)
					return false
				}

				c.writeResDeadline(&gortsplib.Response{
					StatusCode: 200,
					Status:     "OK",
					Header: gortsplib.Header{
						"CSeq": []string{cseq[0]},
						"Transport": []string{strings.Join([]string{
							"RTP/AVP/UDP",
							"unicast",
							fmt.Sprintf("client_port=%d-%d", rtpPort, rtcpPort),
							fmt.Sprintf("server_port=%d-%d", c.p.rtpPort, c.p.rtcpPort),
						}, ";")},
						"Session": []string{"12345678"},
					},
				})
				return true

				// record via TCP
			} else if _, ok := th["RTP/AVP/TCP"]; ok {
				if _, ok := c.p.protocols[_STREAM_PROTOCOL_TCP]; !ok {
					c.log("ERR: tcp streaming is disabled")
					c.writeResDeadline(&gortsplib.Response{
						StatusCode: 461,
						Status:     "Unsupported Transport",
						Header: gortsplib.Header{
							"CSeq": []string{cseq[0]},
						},
					})
					return false
				}

				var interleaved string
				err = func() error {
					c.p.mutex.Lock()
					defer c.p.mutex.Unlock()

					if len(c.streamTracks) > 0 && c.streamProtocol != _STREAM_PROTOCOL_TCP {
						return fmt.Errorf("client want to send tracks with different protocols")
					}

					if len(c.streamTracks) >= len(c.streamSdpParsed.Medias) {
						return fmt.Errorf("all the tracks have already been setup")
					}

					interleaved = th.GetValue("interleaved")
					if interleaved == "" {
						return fmt.Errorf("transport header does not contain interleaved field")
					}

					expInterleaved := fmt.Sprintf("%d-%d", 0+len(c.streamTracks)*2, 1+len(c.streamTracks)*2)
					if interleaved != expInterleaved {
						return fmt.Errorf("wrong interleaved value, expected '%s', got '%s'", expInterleaved, interleaved)
					}

					c.streamProtocol = _STREAM_PROTOCOL_TCP
					c.streamTracks = append(c.streamTracks, &track{
						rtpPort:  0,
						rtcpPort: 0,
					})

					c.state = _CLIENT_STATE_PRE_RECORD
					return nil
				}()
				if err != nil {
					c.writeResError(req, err)
					return false
				}

				c.writeResDeadline(&gortsplib.Response{
					StatusCode: 200,
					Status:     "OK",
					Header: gortsplib.Header{
						"CSeq": []string{cseq[0]},
						"Transport": []string{strings.Join([]string{
							"RTP/AVP/TCP",
							"unicast",
							fmt.Sprintf("interleaved=%s", interleaved),
						}, ";")},
						"Session": []string{"12345678"},
					},
				})
				return true

			} else {
				c.writeResError(req, fmt.Errorf("transport header does not contain a valid protocol (RTP/AVP, RTP/AVP/UDP or RTP/AVP/TCP) (%s)", tsRaw[0]))
				return false
			}

		default:
			c.writeResError(req, fmt.Errorf("client is in state '%d'", c.state))
			return false
		}

	case "PLAY":
		if c.state != _CLIENT_STATE_PRE_PLAY {
			c.writeResError(req, fmt.Errorf("client is in state '%d'", c.state))
			return false
		}

		pub, ok := c.p.publishers[c.path]
		if !ok {
			c.writeResError(req, fmt.Errorf("no one is streaming on path '%s'", c.path))
			return false
		}

		err := func() error {
			c.p.mutex.Lock()
			defer c.p.mutex.Unlock()

			if len(c.streamTracks) != len(pub.streamSdpParsed.Medias) {
				return fmt.Errorf("not all tracks have been setup")
			}

			return nil
		}()
		if err != nil {
			c.writeResError(req, err)
			return false
		}

		// first write response, then set state
		// otherwise, in case of TCP connections, RTP packets could be written
		// before the response
		c.writeResDeadline(&gortsplib.Response{
			StatusCode: 200,
			Status:     "OK",
			Header: gortsplib.Header{
				"CSeq":    []string{cseq[0]},
				"Session": []string{"12345678"},
			},
		})

		c.log("is receiving on path '%s', %d %s via %s", c.path, len(c.streamTracks), func() string {
			if len(c.streamTracks) == 1 {
				return "track"
			}
			return "tracks"
		}(), c.streamProtocol)

		c.p.mutex.Lock()
		c.state = _CLIENT_STATE_PLAY
		c.p.mutex.Unlock()

		c.p.mutex.Lock()
		pub.clientsCount++
		log.Println("//////////////client count", pub.clientsCount)
		if (pub.clientsCount == 1) {
			log.Println("//////////////run stream ")
			go func () {
				spsSended := false
				for {
					if (!pub.started || pub.clientsCount == 0){
						break
					}

					packet := <- pub.cameraPackets

					if (!spsSended) {
						if (!IsSpsPacket("h264", packet.Content)) {
							continue
						}
					}
					spsSended = true

					c.writePacket(&packet)
				}
			}()
		}
		c.p.mutex.Unlock()

		return true

	case "PAUSE":
		if c.state != _CLIENT_STATE_PLAY {
			c.writeResError(req, fmt.Errorf("client is in state '%d'", c.state))
			return false
		}

		if path != c.path {
			c.writeResError(req, fmt.Errorf("path has changed"))
			return false
		}

		c.log("paused")

		c.p.mutex.Lock()
		c.state = _CLIENT_STATE_PRE_PLAY
		c.p.mutex.Unlock()

		c.writeResDeadline(&gortsplib.Response{
			StatusCode: 200,
			Status:     "OK",
			Header: gortsplib.Header{
				"CSeq":    []string{cseq[0]},
				"Session": []string{"12345678"},
			},
		})
		return true

	case "RECORD":
		if c.state != _CLIENT_STATE_PRE_RECORD {
			c.writeResError(req, fmt.Errorf("client is in state '%d'", c.state))
			return false
		}

		if path != c.path {
			c.writeResError(req, fmt.Errorf("path has changed"))
			return false
		}

		err := func() error {
			c.p.mutex.Lock()
			defer c.p.mutex.Unlock()

			if len(c.streamTracks) != len(c.streamSdpParsed.Medias) {
				return fmt.Errorf("not all tracks have been setup")
			}

			return nil
		}()
		if err != nil {
			c.writeResError(req, err)
			return false
		}

		c.writeResDeadline(&gortsplib.Response{
			StatusCode: 200,
			Status:     "OK",
			Header: gortsplib.Header{
				"CSeq":    []string{cseq[0]},
				"Session": []string{"12345678"},
			},
		})

		c.p.mutex.Lock()
		c.state = _CLIENT_STATE_RECORD
		c.p.mutex.Unlock()

		c.log("is publishing on path '%s', %d %s via %s", c.path, len(c.streamTracks), func() string {
			if len(c.streamTracks) == 1 {
				return "track"
			}
			return "tracks"
		}(), c.streamProtocol)

		// when protocol is TCP, the RTSP connection becomes a RTP connection
		// receive RTP data and parse it
		if c.streamProtocol == _STREAM_PROTOCOL_TCP {
			for {
				c.conn.NetConn().SetReadDeadline(time.Now().Add(_READ_TIMEOUT))
				frame, err := c.conn.ReadInterleavedFrame()
				if err != nil {
					c.log("ERR: %s", err)
					return false
				}

				trackId, trackFlow := interleavedChannelToTrack(frame.Channel)

				if trackId >= len(c.streamTracks) {
					c.log("ERR: invalid track id '%d'", trackId)
					return false
				}

				c.p.mutex.RLock()
				c.p.forwardTrack(c.path, trackId, trackFlow, frame.Content)
				c.p.mutex.RUnlock()
			}
		}

		return true

	case "TEARDOWN":
		c.p.mutex.RLock()
		pub, ok := c.p.publishers[c.path]
		if !ok {
			c.writeResError(req, fmt.Errorf("no one is streaming on path '%s'", c.path))
			return false
		}
		pub.clientsCount--
		c.p.mutex.RUnlock()
		// close connection silently
		return false

	default:
		c.writeResError(req, fmt.Errorf("unhandled method '%s'", req.Method))
		return false
	}
}


func (c *client) writePacket(packet *gortsplib.InterleavedFrame){
	// when protocol is TCP, the RTSP connection becomes a RTP connection
	if c.streamProtocol == _STREAM_PROTOCOL_TCP {
		// write RTP frames sequentially
		c.conn.NetConn().SetWriteDeadline(time.Now().Add(_WRITE_TIMEOUT))
		c.conn.WriteInterleavedFrame(packet)

		// receive RTP feedback, do not parse it, wait until connection closes
		buf := make([]byte, 2048)
		for {
			_, err := c.conn.NetConn().Read(buf)
			if err != nil {
				if err != io.EOF {
					c.log("ERR: %s", err)
				}
				return
			}
		}
	} else {
		trackId, trackFlow := interleavedChannelToTrack(packet.Channel)
		if len(c.streamTracks) <= trackId{
			return
		}
		
		if trackFlow == _TRACK_FLOW_RTP {
			c.p.rtpl.chanWrite <- &udpWrite{
				addr: &net.UDPAddr{
					IP:   c.ip,
					Port: c.streamTracks[trackId].rtpPort,
				},
				buf: packet.Content,
			}
		} else {
			c.p.rtcpl.chanWrite <- &udpWrite{
				addr: &net.UDPAddr{
					IP:   c.ip,
					Port: c.streamTracks[trackId].rtcpPort,
				},
				buf: packet.Content,
			}
		}
	}
}
