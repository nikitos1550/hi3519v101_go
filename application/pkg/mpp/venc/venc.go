package venc

//#include "venc.h"
import "C"

import (
    "errors"

    "application/pkg/mpp/vi"

    "application/pkg/logger"
    "application/pkg/mpp/errmpp"
    "application/pkg/buildinfo"
)

var (
    channelsAmount  int = C.VENC_MAX_CHN_NUM
    minWidth        int
    minHeight       int
    //maxWidth        int   //already limeted by VI
    //maxHeight       int   //already limited by VI
    maxBitrate      int = 1024*15 //TODO
)

func Init() {
    loopInit()
	readEncoders()
}

func mppCreateEncoder(id int, params Parameters) error {
    var inErr C.error_in
    var in C.mpp_venc_create_encoder_in

    in.id               = C.uint(id)

    switch params.Codec {
        case MJPEG:
            in.codec			= C.uint(C.PT_MJPEG)
        case H264:
            in.codec            = C.uint(C.PT_H264)
        case H265:
            if buildinfo.Family == "hi3516cv100" {
                return errors.New("Codec is not supported")
            }
            in.codec            = C.uint(C.PT_H265)
        default:
            return errors.New("Unknown codec")
    }

    switch params.Profile {
        case Baseline:
            in.profile			= C.uint(0)
        case Main:
            if params.Codec == MJPEG {
                return errors.New("MJPEG supports only baseline profile")
            }
            in.profile          = C.uint(1)
        //case Main10:
        //    in.profile          = C.uint(?)
        case High:
            in.profile          = C.uint(3)
        default:
            return errors.New("Unknown profile")
    }

    switch params.BitControl {
        case Cbr:
            switch params.Codec {
                case MJPEG:
                    in.bitrate_control  = C.uint(C.VENC_RC_MODE_MJPEGCBR)
                case H264:
                    in.bitrate_control  = C.uint(C.VENC_RC_MODE_H264CBR)
                case H265:
                    in.bitrate_control  = C.uint(C.VENC_RC_MODE_H265CBR)
            }
        case Vbr:
            switch params.Codec {
                case MJPEG:
                    in.bitrate_control  = C.uint(C.VENC_RC_MODE_MJPEGVBR)
                case H264:
                    in.bitrate_control  = C.uint(C.VENC_RC_MODE_H264VBR)
                case H265:
                    in.bitrate_control  = C.uint(C.VENC_RC_MODE_H265VBR)
            }
        case FixQp:
            switch params.Codec {
                case MJPEG:
                    in.bitrate_control  = C.uint(C.VENC_RC_MODE_MJPEGFIXQP)
                case H264:
                    in.bitrate_control  = C.uint(C.VENC_RC_MODE_H264FIXQP)
                case H265:
                    in.bitrate_control  = C.uint(C.VENC_RC_MODE_H265FIXQP)
            }
        case CVbr:
            if  buildinfo.Family == "hi3516cv100" ||
                buildinfo.Family == "hi3516cv200" ||
                buildinfo.Family == "hi3516cv300" ||
                buildinfo.Family == "hi3516av100" ||
                buildinfo.Family == "hi3516av200" {
                return errors.New("Chip doesn`t support cvbr")
            }

            switch params.Codec {
                case MJPEG:
                    return errors.New("MJPEG doesn`t support cvbr")
                case H264:
                    in.bitrate_control  = C.uint(C.VENC_RC_MODE_H264CVBR)
                case H265:
                    in.bitrate_control  = C.uint(C.VENC_RC_MODE_H265CVBR)
            }
        case AVbr:
            if buildinfo.Family == "hi3516cv100" {
                return errors.New("Chip doesn`t support avbr")
            }
            switch params.Codec {
                case MJPEG:
                    return errors.New("MJPEG doesn`t support avbr")
                case H264:
                    in.bitrate_control  = C.uint(C.VENC_RC_MODE_H264AVBR)
                case H265:
                    in.bitrate_control  = C.uint(C.VENC_RC_MODE_H265AVBR)
            }
        case QVbr:
            if  buildinfo.Family == "hi3516cv100" ||
                buildinfo.Family == "hi3516cv200" ||
                buildinfo.Family == "hi3516cv300" ||
                buildinfo.Family == "hi3516av100" ||
                buildinfo.Family == "hi3516av200" {
                return errors.New("Chip doesn`t support qvbr")
            }

            switch params.Codec {
                case MJPEG:
                    return errors.New("MJPEG doesn`t support qvbr")
                case H264:
                    in.bitrate_control  = C.uint(C.VENC_RC_MODE_H264QVBR)
                case H265:
                    in.bitrate_control  = C.uint(C.VENC_RC_MODE_H265QVBR)
            }
        default:
            return errors.New("Unknown bitrate control strategy")
    }

    if params.Width > uint(vi.Width()) {
        return errors.New("Width can`t be too much") //TODO
    }
    if params.Width < uint(minWidth) {
        return errors.New("Width can`t be so small")
    }
    in.width			= C.uint(params.Width)

    if params.Height > uint(vi.Height()) {
        return errors.New("Height can`t be too much") //TODO
    }
    if params.Height < uint(minHeight) {
        return errors.New("Height can`t be so small")
    }
    in.height			= C.uint(params.Height)

    //TODO check with channel
    in.in_fps			= C.int(-1)

    if params.Fps > uint(vi.Fps()) {
        return errors.New("Fps can`t be too large")
    }
    //TODO check with channel
    in.out_fps			= C.int(params.Fps)

    if params.Gop > 65536 {
        return errors.New("Gop should be <= than 65536")
    }
    in.gop				= C.uint(params.Gop)

    //in.gop_mode			= C.uint(0) //TODO

    if params.BitControlParams.Bitrate > uint(maxBitrate) {
        return errors.New("Bitrate is too large")
    }
	in.bitrate			= C.uint(params.BitControlParams.Bitrate)

    if params.BitControlParams.StatTime > 5 { //TODO
        return errors.New("Stattime is too large")
    }
    in.stat_time		= C.uint(params.BitControlParams.StatTime)

    in.fluctuate_level	= C.uint(params.BitControlParams.Fluctuate)
    in.q_factor			= C.uint(params.BitControlParams.QFactor)
    in.min_q_factor		= C.uint(params.BitControlParams.MinQFactor)
    in.max_q_factor		= C.uint(params.BitControlParams.MaxQFactor)
    in.i_qp				= C.uint(params.BitControlParams.IQp)
    in.p_qp				= C.uint(params.BitControlParams.PQp)
    in.b_qp				= C.uint(params.BitControlParams.BQp)
    in.min_qp			= C.uint(params.BitControlParams.MinQp)
    in.max_qp			= C.uint(params.BitControlParams.MaxQp)
    in.min_i_qp			= C.uint(params.BitControlParams.MinIQp)

    logger.Log.Trace().
    	Uint("codec", 			uint(in.codec)).
        Uint("profile", 		uint(in.profile)).
        Uint("width", 			uint(in.width)).
        Uint("height", 			uint(in.height)).
        Int("in_fps", 			int(in.in_fps)).
        Int("out_fps", 		    int(in.out_fps)).
        Uint("bitrate_control", uint(in.bitrate_control)).
        Uint("gop", 			uint(in.gop)).
        Uint("gop_mode", 		uint(in.gop_mode)).
        Uint("bitrate", 		uint(in.bitrate)).
        Uint("stat_time", 		uint(in.stat_time)).
        Uint("fluctuate_level",	uint(in.fluctuate_level)).
        Uint("q_factor", 		uint(in.q_factor)).
        Uint("min_q_factor", 	uint(in.min_q_factor)).
        Uint("max_q_factor", 	uint(in.max_q_factor)).
        Uint("i_qp", 			uint(in.i_qp)).
        Uint("p_qp", 			uint(in.p_qp)).
        Uint("b_qp", 			uint(in.b_qp)).
        Uint("min_qp", 			uint(in.min_qp)).
        Uint("max_qp", 			uint(in.max_qp)).
        Uint("min_i_qp", 		uint(in.min_i_qp)).
        Msg("VENC encoder params")

    err := C.mpp_venc_create_encoder(&inErr, &in)

    if err != 0 {
        logger.Log.Fatal().
            Str("error", errmpp.New(C.GoString(inErr.name), uint(inErr.code)).Error()).
            Msg("VENC")
        return errmpp.New(C.GoString(inErr.name), uint(inErr.code))
    }

    return nil
}

//func mppUpdateEncoder(id int, params Parameters) error {
//    var inErr C.error_in
//    var in C.mpp_venc_create_encoder_in
//
//    err := C.mpp_venc_update_encoder(&inErr, &in)
//
//    if err != 0 {
//    	logger.Log.Fatal().
//        	Str("error", errmpp.New(C.GoString(inErr.name), uint(inErr.code)).Error()).
//            Msg("VENC")
//      	return errmpp.New(C.GoString(inErr.name), uint(inErr.code))
//  	}
//
//	return nil
//}

func mppDestroyEncoder(id int) error {
    var inErr C.error_in
    var in C.mpp_venc_destroy_encoder_in

    in.id = C.uint(id)

    err := C.mpp_venc_destroy_encoder(&inErr, &in)

    if err != 0 {
        logger.Log.Fatal().
            Str("error", errmpp.New(C.GoString(inErr.name), uint(inErr.code)).Error()).
            Msg("VENC")
        return errmpp.New(C.GoString(inErr.name), uint(inErr.code))
    }

    return nil
}

//export go_logger_venc
func go_logger_venc(level C.int, msgC *C.char) {
        logger.CLogger("VENC", int(level), C.GoString(msgC))
}
/////////////////////////////////////////////////////////////////////////

//func SubsribeEncoder(encoderId int, ch chan []byte) {
func SubsribeEncoder(encoderId int, ch chan ChannelEncoder) {
    encoder, encoderExists := ActiveEncoders[encoderId]
    if !encoderExists {
		logger.Log.Error().
			Int("encoderId", encoderId).
			Msg("Failed to find encoder")
        return
    }
		
    _, exists := encoder.Channels[ch]
    if (exists) {
		logger.Log.Error().
			Int("encoderId", encoderId).
			Msg("Already subscribed")
        return
    }

    encoder.Channels[ch] = true
    ActiveEncoders[encoderId] = encoder
}

//func RemoveSubscription(encoderId int, ch chan []byte) {
func RemoveSubscription(encoderId int, ch chan ChannelEncoder) {
    encoder, encoderExists := ActiveEncoders[encoderId]
    if !encoderExists {
		logger.Log.Error().
			Int("encoderId", encoderId).
			Msg("Failed to find encoder")
        return
    }
		
    _, exists := encoder.Channels[ch]
    if (!exists) {
		logger.Log.Error().
			Int("encoderId", encoderId).
			Msg("Not subscribed")
        return
    }


    delete(encoder.Channels, ch)
    ActiveEncoders[encoderId] = encoder
}
