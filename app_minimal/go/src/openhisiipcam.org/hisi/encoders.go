package hisi

import (
    "sync"
)

const (
	CodecMjpeg = 1
	CodecH264  = 2
	CodecH265  = 3

	RcCbr   = 1
	RcVbr   = 2
	RcFixqp = 3

	ProfileJpegBaseline     = 10
	ProfileMjpegBaseline    = 20
	ProfileH264Baseline     = 30
	ProfileH264Main         = 31
	ProfileH264Hight        = 32
	ProfileH265Main         = 40
)

var (
    encs encoders
)

type encoder struct {
	lock    sync.Mutex
	id      uint
	enabled uint
}

type encoders struct {
	maxNum uint
	enc    []encoder
}

////////////////////////////////////////////////////////////////////////////////

////export goDataCallback
//func goDataCallback(encId C.uint, encData * C.struct_encoderData) C.int {
//
//    return 0;
//}


////////////////////////////////////////////////////////////////////////////////

/*
func EncodersMaxNum() uint {
	return encs.maxNum
}

func EncodersGetInfo() EncodersInfo {
	var out EncodersInfo
	out.MaxNum = encs.maxNum
	return out
}

func EncoderGetInfo(id uint) (EncoderInfo, int) {
	var out EncoderInfo

	if id >= encs.maxNum {
		return out, ERR_OBJ_NOT_FOUND
	}
	var e *encoder = &encs.enc[id]

	e.lock.Lock()
	defer e.lock.Unlock()

	if e.enabled == 1 {
		var sparams C.struct_encoder_static_params
		var dparams C.struct_encoder_dynamic_params
		if C.hisi_encoder_info(C.uint(e.id), &sparams, &dparams) != 0 {
			panic("1")
		}
		//log.Println("\twidth ", tmp.width, " height ", tmp.height, " fps ", tmp.fps)
		out.Enabled = 1
		switch sparams.codec {
		case C.CODEC_MJPEG:
			out.Codec = CODEC_MJPEG
			switch sparams.profile {
			case C.PROFILE_MJPEG_BASELINE:
				out.Profile = PROFILE_MJPEG_BASELINE
				break
			}
			break
		case C.CODEC_H264:
			out.Codec = CODEC_H264
			switch sparams.profile {
			case C.PROFILE_H264_BASELINE:
				out.Profile = PROFILE_MJPEG_BASELINE
				break
			case C.PROFILE_H264_MAIN:
				out.Profile = PROFILE_H264_MAIN
				break
			case C.PROFILE_H264_HIGH:
				out.Profile = PROFILE_H264_HIGH
				break
			}
			break
		case C.CODEC_H265:
			out.Codec = CODEC_H265
			switch sparams.profile {
			case C.PROFILE_H265_MAIN:
				out.Profile = PROFILE_H265_MAIN
				break
			}
			break
		}
		out.Width = int(sparams.width)
		out.Height = int(sparams.height)
		out.Fps = int(sparams.fps)
		switch sparams.rc {
		case C.RC_CBR:
			out.Codec = RC_CBR
			out.Cbr.Bitrate = int(dparams.cbr.bitrate)
			out.Cbr.Gop = int(dparams.cbr.gop)
			out.Cbr.StatTime = int(dparams.cbr.stattime)
			out.Cbr.Fluctuate = int(dparams.cbr.fluctuate)
			break
		case C.RC_VBR:
			out.Codec = RC_VBR
			out.Vbr.MaxBitrate = int(dparams.vbr.maxbitrate)
			out.Vbr.Gop = int(dparams.vbr.gop)
			out.Vbr.MaxQp = int(dparams.vbr.maxqp)
			out.Vbr.MinQp = int(dparams.vbr.minqp)
			out.Vbr.StatTime = int(dparams.vbr.stattime)
			break
		}
		out.ChannelId = int(sparams.channel_id)
	}

	return out, ERR_NONE
}

func EncoderEnable(id uint, codec, rc, w, h, fps, bitrate int) int {
	if id >= encs.maxNum {
		return ERR_OBJ_NOT_FOUND
	}
	var e *encoder = &encs.enc[id]

	e.lock.Lock()
	defer e.lock.Unlock()

	if e.enabled == 1 {
		return ERR_NOT_ALLOWED
	}

	return ERR_NONE
}

func EncoderUpdate(id uint, bitrate int) int {
	if id >= encs.maxNum {
		return ERR_OBJ_NOT_FOUND
	}
	var e *encoder = &encs.enc[id]

	e.lock.Lock()
	defer e.lock.Unlock()

	return ERR_NONE
}

func EncoderDisable(id uint) int {
	if id >= encs.maxNum {
		return ERR_OBJ_NOT_FOUND
	}
	var e *encoder = &encs.enc[id]

	e.lock.Lock()
	defer e.lock.Unlock()

	if e.enabled == 0 {
		return ERR_NOT_ALLOWED
	}

	//if(e.hisi_encoder_disable(C.uint(e.id)) != 0) {
	//    panic("1")
	//}

	e.enabled = 0

	return ERR_NONE
}
*/
