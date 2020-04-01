//+build streamerRtsp

package rtsp

import (
	"bytes"
	_ "math"
	"math/rand"
	"time"
)

const (
	RtpHeaderSize = 12 + 2 //4 bytes interleaved prefix + 12 bytes RTP header + 2 bytes FUA header
	fuASize       = 1387   //ipPacketSize - RtpHeaderSize
)

type Packetizer interface {
	NalH264ToRtp(nal []byte) [][]byte
	H264ToRtp(nal []byte) [][]byte
}

type packetizer struct {
	timestamp uint32
	sequence  uint16
	ssrc      uint32
}

func CreatePacketizer() Packetizer {
	rs := rand.NewSource(time.Now().UnixNano())
	r := rand.New(rs)

	return &packetizer{
		timestamp: r.Uint32(),
		sequence:  0, //uint16(r.Uint32() % math.MaxUint16), //tmp for debug
		ssrc:      rand.Uint32(),
	}
}

func (p *packetizer) NalH264ToRtp(nal []byte) [][]byte {
	//var count uint = uint(len(nal) / fuASize) + 1

	var nalType byte = nal[0] & 31
	var nri byte = (nal[0] >> 5) & 3

	nal2 := nal[1:len(nal)]
	var count uint = uint(len(nal2)/fuASize) + 1

	var out [][]byte = make([][]byte, count)

	var FuAStart byte = 1
	var FuAEnd byte = 0

	for i := uint(0); i < count; i++ {
		packetSize := uint(fuASize)

		if i == (count - 1) {
			FuAEnd = 1
			packetSize = uint(len(nal2)) - (i)*uint(fuASize)
		}

		out[i] = make([]byte, RtpHeaderSize+packetSize)

		out[i][0] = 2 << 6                         //WTF? r[0] = 2 << 6;
		out[i][1] = (FuAEnd << 7) | 96             //r[1] = (fua_end ? 1 : 0) << 7 | 96; // 96 is our hardcoded h264 payload type
		out[i][2] = byte((p.sequence >> 8) & 0xFF) //TODO r[2] = server->seq[frame->stream_id] >> 8;
		out[i][3] = byte((p.sequence) & 0xFF)      //TODO r[3] = server->seq[frame->stream_id] >> 0;

		out[i][4] = byte((p.timestamp >> 24) & 0xFF) //TODO r[4] = tc >> 24;
		out[i][5] = byte((p.timestamp >> 16) & 0xFF) //TODO r[5] = tc >> 16;
		out[i][6] = byte((p.timestamp >> 8) & 0xFF)  //TODO r[6] = tc >> 8;
		out[i][7] = byte(p.timestamp & 0xFF)         //TODO r[7] = tc >> 0;
		out[i][8] = 0                                //OK  r[8] = r[9] = r[10] = 0;
		out[i][9] = 0                                //OK
		out[i][10] = 0                               //OK
		out[i][11] = 1                               //  r[11] = frame->stream_id+1;

		out[i][12] = nri<<5 | 28                       //r[0] = nri << 5 | 28; // 28 = H264 FUA
		out[i][13] = FuAStart<<7 | FuAEnd<<6 | nalType //r[1] = fua_start << 7 | fua_end << 6 | nal_type;

		copy(out[i][RtpHeaderSize:RtpHeaderSize+packetSize], nal2[(i*fuASize):(i*fuASize)+packetSize])
		FuAStart = 0 //after first iter it will be 0 forever
		p.sequence = p.sequence + 1
	}

	p.timestamp = p.timestamp + (90000 / 25)
	return out
}

func (p *packetizer) H264ToRtp(nal []byte) [][]byte {
	var out [][]byte
	payloads := bytes.Split(nal, KeyData)
	for _, payload := range payloads {
		if (len(payload) <= 0){
			continue
		}

		packets := p.NalH264ToRtp(payload)
		for _, p := range packets {
			out = append(out, p)
		}
	}
	return out
}

