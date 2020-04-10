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
	fuASize       = 1387   //ipPacketSize >= RtpHeaderSize + fuASize
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

func (p *packetizer) FrameMJPEGToRtp(nal []byte) [][]byte {
	return nil
}
/*
func (p *packetizer) NalH265ToRtp(nal []byte) [][]byte {
	var nalType byte = (nal[0] & 126) >> 1
	var payloadH1 byte = nal[0];
        var payloadH2 byte = nal[1];

	nal2 := nal[2:len(nal)]

	var count uint = uint(len(nal2)/fuASize) + 1

	var out [][]byte = make([][]byte, count)

	if (count > 1) { //FUA
		var FuAStart byte = 1
        	var FuAEnd byte = 0

		for i:= 0; i < count; i++ {

			packetSize := 12 + 3 + fuASize //rtpheader fragmented mode + paylaod 
			if i == (count-1) {
				packetSize = 12 + 3 + uint(len(nal2)) - (i)*uint(fuASize)
			}
			out[i] = make([]byte, packetSize)

	                out[i][0] = 2 << 6;
	                out[i][1] = (FuAEnd << 7) | 98; // 98 is our hardcoded h265 payload type
	                out[i][2] = byte((p.sequence >> 8) & 0xFF)
	                out[i][3] = byte((p.sequence) & 0xFF)
	
	                out[i][4] = byte((p.timestamp >> 24) & 0xFF)
	                out[i][5] = byte((p.timestamp >> 16) & 0xFF)
	                out[i][6] = byte((p.timestamp >> 8) & 0xFF)
	                out[i][7] = byte(p.timestamp & 0xFF)
	                out[i][8] = 0
	                out[i][9] = 0
	                out[i][10] = 0
	                out[i][11] = 2;

                	out[i][12] = 98
        	        out[i][13] = 1
			out[i][14] = fuaStart << 7 | fuaEnd << 6 | nalType;
	
        	        copy(out[0][15:packetSize], nal2[(i*fuASize):(i*fuASize)+packetSize-15])
	                p.sequence = p.sequence + 1

		}
	} else { //single packet
		var packetSize uint = 12 + 2 + len(nal2) // rtpheader non fragmented mode + payload
		
		out[i] = make([]byte, packetSize)

		out[0][0] = 2 << 6;
                out[0][1] = 98; // 98 is our hardcoded h265 payload type
		out[0][2] = byte((p.sequence >> 8) & 0xFF)
		out[0][3] = byte((p.sequence) & 0xFF)

                out[0][4] = byte((p.timestamp >> 24) & 0xFF)
                out[0][5] = byte((p.timestamp >> 16) & 0xFF)
                out[0][6] = byte((p.timestamp >> 8) & 0xFF)
                out[0][7] = byte(p.timestamp & 0xFF)
                out[0][8] = 0
		out[0][9] = 0
		out[0][10] = 0
                out[0][11] = 2;

		out[0][12] = payloadH1
		out[0][13] = payloadH2

		copy(out[0][14:packetSize], nal2[0:len(nal2)])

		FuAStart = 0 //after first iter it will be 0 forever
		p.sequence = p.sequence + 1

	}

	p.timestamp = p.timestamp + (90000 / 25)

	return out
}
*/
// This is always FuA packatizer, even if NAL length is less than
// one packet size we will pack it as FuA with corresponding start and end marks
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

