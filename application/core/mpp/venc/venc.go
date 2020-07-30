package venc

//#include "venc.h"
import "C"

import (
    "errors"
    "unsafe"

    "application/core/mpp/vi"
    "application/core/mpp/connection"
    "application/core/logger"
    "application/core/mpp/errmpp"
    "application/core/compiletime"
)

var (
    EncodersAmount  int = C.VENC_MAX_CHN_NUM
    minWidth        int //TODO
    minHeight       int //TODO
    maxBitrate      int = 1024*30 //TODO
)

//const invalidValue2 int = int(C.INVALID_VALUE)

func Init() {
    err := loopInit()
    if err != nil {
        logger.Log.Fatal().
            Msg(err.Error())
    }
}

func mppCreateEncoder(id int, params Parameters) error {
    var inErr C.error_in
    var in C.mpp_venc_create_encoder_in

    C.invalidate_mpp_venc_create_encoder_in(&in)

    var err error

    //setParamsNotValidIfZero(&params)

    in.id = C.int(id)

    switch params.Codec {
        case MJPEG:
            in.codec			= C.int(C.PT_MJPEG)
        case H264:
            in.codec            = C.int(C.PT_H264)
        case H265:
            if compiletime.Family == "hi3516cv100" { return errors.New("Codec is not supported") }
            in.codec            = C.int(C.PT_H265)
        default:
            return errors.New("Unknown codec")
    }

    switch params.Profile {
        case Baseline:
            in.profile			= C.int(0)
        case Main:
            if params.Codec == MJPEG { return errors.New("MJPEG supports only baseline profile") }
            in.profile          = C.int(1)
        case High:
            if params.Codec == MJPEG { return errors.New("MJPEG supports only baseline profile") }
            in.profile          = C.int(2)
        default:
            return errors.New("Unknown profile")
    }

    switch params.BitControl {
        case Cbr:
            switch params.Codec {
                case MJPEG:
                    in.bitrate_control  = C.int(C.VENC_RC_MODE_MJPEGCBR)
                case H264:
                    in.bitrate_control  = C.int(C.VENC_RC_MODE_H264CBR)
                case H265:
                    in.bitrate_control  = C.int(C.VENC_RC_MODE_H265CBR)
            }

            if err = checkParamBitrate(&params); err != nil { return err }
            in.cbr.bitrate          = C.int(params.BitControlParams.Bitrate)
            if err = checkParamStatTime(&params); err != nil { return err }
            in.cbr.stat_time        = C.int(params.BitControlParams.StatTime)

            if  compiletime.Family != "hi3516cv500" &&
                compiletime.Family != "hi3516ev200" &&
                compiletime.Family != "hi3519av100" {
                if err = checkParamFluctuate(&params); err != nil { return err }
                in.cbr.fluctuate_level  = C.int(params.BitControlParams.Fluctuate)
            }

        case Vbr:
            switch params.Codec {
                case MJPEG:
                    in.bitrate_control  = C.int(C.VENC_RC_MODE_MJPEGVBR)
                case H264:
                    in.bitrate_control  = C.int(C.VENC_RC_MODE_H264VBR)
                case H265:
                    in.bitrate_control  = C.int(C.VENC_RC_MODE_H265VBR)
            }

			if err = checkParamBitrate(&params); err != nil { return err }
            in.vbr.maxbitrate          = C.int(params.BitControlParams.Bitrate)
            if err = checkParamStatTime(&params); err != nil { return err }
            in.vbr.stat_time        = C.int(params.BitControlParams.StatTime)

			if  compiletime.Family != "hi3516cv500" &&
				compiletime.Family != "hi3516ev200" &&
                compiletime.Family != "hi3519av100" {
				switch params.Codec {
					case MJPEG:
						if err = checkParamMinQFactor(&params); err != nil { return err }
                        in.vbr.min_q_factor     = C.int(params.BitControlParams.MinQFactor)
						if err = checkParamMaxQFactor(&params); err != nil { return err }
                        in.vbr.max_q_factor     = C.int(params.BitControlParams.MaxQFactor)

					case H264, H265:
						if err = checkParamMinQp(&params); err != nil { return err }
                        in.vbr.min_qp           = C.int(params.BitControlParams.MinQp)
						if err = checkParamMaxQp(&params); err != nil { return err }
                        in.vbr.max_qp           = C.int(params.BitControlParams.MaxQp)

						if	compiletime.Family == "hi3516cv300" ||
							compiletime.Family == "hi3516av200" {
							if err = checkParamMinIQp(&params); err != nil { return err }
                            in.vbr.min_i_qp         = C.int(params.BitControlParams.MinIQp)
						}
				}
           }

        case FixQp:
            switch params.Codec {
                case MJPEG:
                    in.bitrate_control  = C.int(C.VENC_RC_MODE_MJPEGFIXQP)
                case H264:
                    in.bitrate_control  = C.int(C.VENC_RC_MODE_H264FIXQP)
                case H265:
                    in.bitrate_control  = C.int(C.VENC_RC_MODE_H265FIXQP)
            }

            if err = checkParamQFactor(&params); err != nil { return err }
            in.fixqp.q_factor         = C.int(params.BitControlParams.QFactor)

        case CVbr:
            if  compiletime.Family == "hi3516cv100" ||
                compiletime.Family == "hi3516cv200" ||
                compiletime.Family == "hi3516cv300" ||
                compiletime.Family == "hi3516av100" ||
                compiletime.Family == "hi3516av200" {
                return errors.New("Chip doesn`t support cvbr")
            }

            switch params.Codec {
                case MJPEG:
                    return errors.New("MJPEG doesn`t support cvbr")
                case H264:
                    in.bitrate_control  = C.int(C.VENC_RC_MODE_H264CVBR)
                case H265:
                    in.bitrate_control  = C.int(C.VENC_RC_MODE_H265CVBR)
            }

            return errors.New("Not implemented, TODO")  //TODO

        case AVbr:
            if compiletime.Family == "hi3516cv100" {
                return errors.New("Chip doesn`t support avbr")
            }

            switch params.Codec {
                case MJPEG:
                    return errors.New("MJPEG doesn`t support avbr")
                case H264:
                    in.bitrate_control  = C.int(C.VENC_RC_MODE_H264AVBR)
                case H265:
                    in.bitrate_control  = C.int(C.VENC_RC_MODE_H265AVBR)
            }

            if err = checkParamStatTime(&params); err != nil { return err }
            in.avbr.stat_time        = C.int(params.BitControlParams.StatTime)
            in.avbr.maxbitrate          = C.int(params.BitControlParams.Bitrate)
        case QVbr:
            if  compiletime.Family == "hi3516cv100" ||
                compiletime.Family == "hi3516cv200" ||
                compiletime.Family == "hi3516cv300" ||
                compiletime.Family == "hi3516av100" ||
                compiletime.Family == "hi3516av200" {
                return errors.New("Chip doesn`t support qvbr")
            }

            switch params.Codec {
                case MJPEG:
                    return errors.New("MJPEG doesn`t support qvbr")
                case H264:
                    in.bitrate_control  = C.int(C.VENC_RC_MODE_H264QVBR)
                case H265:
                    in.bitrate_control  = C.int(C.VENC_RC_MODE_H265QVBR)
            }

            if err = checkParamStatTime(&params); err != nil { return err }
            in.qvbr.stat_time        = C.int(params.BitControlParams.StatTime)

        default:
            return errors.New("Unknown bitrate control strategy")
    }

    if params.Width > vi.Width() {
        return errors.New("Width can`t be too much") //TODO
    }
    if params.Width < minWidth {
        return errors.New("Width can`t be so small")
    }
    in.width			= C.int(params.Width)

    if params.Height > vi.Height() {
        return errors.New("Height can`t be too much") //TODO
    }
    if params.Height < minHeight {
        return errors.New("Height can`t be so small")
    }
    in.height			= C.int(params.Height)

    //TODO set maximum, it will be update on connection to source
    in.in_fps			= C.int(vi.Fps())

    if params.Fps > vi.Fps() {
        return errors.New("Fps can`t be too large")
    }
    //TODO check with channel
    in.out_fps			= C.int(params.Fps)

    if params.Codec != MJPEG {

        if  params.GopParams.Gop < 1 ||
            params.GopParams.Gop > 65536 {
            return errors.New("Gop should be [1; 65536]")
        }
        in.gop				= C.int(params.GopParams.Gop)

        switch params.GopType {
            case NormalP:
                in.gop_mode           = C.int(C.VENC_GOPMODE_NORMALP)

                if  compiletime.Family != "hi3516cv100" &&
                    compiletime.Family != "hi3516cv200" &&
                    compiletime.Family != "hi3516av100" {
                    if err = checkParamIPQpDelta(&params); err != nil { return err }
                    in.normalp.i_pq_delta       = C.int(params.GopParams.IPQpDelta)
                }

            case DualP:
			    if  compiletime.Family == "hi3516cv100" ||
				    compiletime.Family == "hi3516cv200" ||
				    compiletime.Family == "hi3516av100" {
				    return errors.New("Chip doesn`t support dualp gop type")
			    }
                in.gop_mode           = C.int(C.VENC_GOPMODE_DUALP)

                if err = checkParamIPQpDelta(&params); err != nil { return err }
                in.dualp.i_pq_delta       = C.int(params.GopParams.IPQpDelta)
                if err = checkParamSPInterval(&params); err != nil { return err }
                in.dualp.s_p_interval     = C.int(params.GopParams.SPInterval)
                if err = checkParamSPQpDelta(&params); err != nil { return err}
                in.dualp.s_pq_delta       = C.int(params.GopParams.SPQpDelta)

            case SmartP:
                if  compiletime.Family == "hi3516cv100" ||
				    compiletime.Family == "hi3516cv200" ||
				    compiletime.Family == "hi3516av100" {
                    return errors.New("Chip doesn`t support smartp gop type")
			    }
                in.gop_mode           = C.int(C.VENC_GOPMODE_SMARTP)

                if err = checkParamBgInterval(&params); err != nil { return err }
                in.smartp.bg_interval      = C.int(params.GopParams.BgInterval)
                if err = checkParamBgQpDelta(&params); err != nil { return err }
                in.smartp.bg_qp_delta      = C.int(params.GopParams.BgQpDelta)
                if err = checkParamViQpDelta(&params); err != nil { return err }
                in.smartp.vi_qp_delta      = C.int(params.GopParams.ViQpDelta)

            case AdvSmartP:
                if  compiletime.Family == "hi3516cv100" ||
				    compiletime.Family == "hi3516cv200" ||
				    compiletime.Family == "hi3516cv300" ||
				    compiletime.Family == "hi3516av100" ||
				    compiletime.Family == "hi3516av200" {
				    return errors.New("Chip doesn`t support advsmartp gop type")
			    }
                in.gop_mode           = C.int(C.VENC_GOPMODE_ADVSMARTP)

                if err = checkParamBgInterval(&params); err != nil { return err }
                in.advsmartp.bg_interval      = C.int(params.GopParams.BgInterval)
                if err = checkParamBgQpDelta(&params); err != nil { return err }
                in.advsmartp.bg_qp_delta      = C.int(params.GopParams.BgQpDelta)
                if err = checkParamViQpDelta(&params); err != nil { return err }
                in.advsmartp.vi_qp_delta      = C.int(params.GopParams.ViQpDelta)

            case BipredB:
                if  compiletime.Family == "hi3516cv100" ||
				    compiletime.Family == "hi3516cv200" ||
                    compiletime.Family == "hi3516cv300" ||
				    compiletime.Family == "hi3516av100" {
				    return errors.New("Chip doesn`t support bipredb gop type")
			    }
                if params.Profile == Baseline {
                    return errors.New("Baseline doesn`t support gop type bipredb")
                }
                in.gop_mode           = C.int(C.VENC_GOPMODE_BIPREDB)

                if err = checkParamBFrmNum(&params); err != nil { return err }
                in.bipredb.b_frm_num        = C.int(params.GopParams.BFrmNum)
                if err = checkParamBQpDelta(&params); err != nil { return err }
                in.bipredb.b_qp_delta       = C.int(params.GopParams.BQpDelta)
                if err = checkParamIPQpDelta(&params); err != nil { return err }
                in.bipredb.i_pq_delta       = C.int(params.GopParams.IPQpDelta)

            case IntraR:
                in.gop_mode           = C.int(C.VENC_GOPMODE_INTRAREFRESH)

                if  compiletime.Family != "hi3516cv100" &&
                    compiletime.Family != "hi3516cv200" &&
                    compiletime.Family != "hi3516av100" {
                    if err = checkParamIPQpDelta(&params); err != nil { return err }
                    in.intrar.i_pq_delta       = C.int(params.GopParams.IPQpDelta)
                }
	    }

        //TODO gop params
    }


    logger.Log.Trace().
        Int("id",               int(in.id)).
        Int("codec",		    int(in.codec)).
        Int("profile",          int(in.profile)).
        Int("width",            int(in.width)).
        Int("height",           int(in.height)).
        Int("in_fps",           int(in.in_fps)).
        Int("out_fps",          int(in.out_fps)).
        Int("bitrate_control",  int(in.bitrate_control)).
        Int("gop",              int(in.gop)).
        Int("gop_mode",         int(in.gop_mode)).
        //Int("i_pq_delta",       int(in.i_pq_delta)).
        //Int("s_p_interval",     int(in.s_p_interval)).
        //Int("s_pq_delta",       int(in.s_pq_delta)).
        //Int("bg_interval",      int(in.bg_interval)).
        //Int("bg_qp_delta",      int(in.bg_qp_delta)).
        //Int("vi_qp_delta",      int(in.vi_qp_delta)).
        //Int("b_frm_num",        int(in.b_frm_num)).
        //Int("b_qp_delta",       int(in.b_qp_delta)).
        //Int("bitrate",          int(in.bitrate)).
        //Int("stat_time",        int(in.stat_time)).
        //Int("fluctuate_level",  int(in.fluctuate_level)).
        //Int("q_factor",         int(in.q_factor)).
        //Int("min_q_factor",     int(in.min_q_factor)).
        //Int("max_q_factor",     int(in.max_q_factor)).
        //Int("i_qp",             int(in.i_qp)).
        //Int("p_qp",             int(in.p_qp)).
        //Int("b_qp",             int(in.b_qp)).
        //Int("min_qp",           int(in.min_qp)).
        //Int("max_qp",           int(in.max_qp)).
        //Int("min_i_qp",         int(in.min_i_qp)).
        Msg("VENC encoder params")

    err2 := C.mpp_venc_create_encoder(&inErr, &in)//TODO err2 rename

    if err2 != 0 {
        logger.Log.Error().
            Str("error", errmpp.New(C.GoString(inErr.name), uint(inErr.code)).Error()).
            Msg("VENC")
        return errmpp.New(C.GoString(inErr.name), uint(inErr.code))
    }

    return nil
}

func mppSetScene(id int, scene int) error {
    var inErr C.error_in

    err := C.mpp_venc_scene(&inErr, C.int(id), C.int(scene))

    if err != 0 {
        logger.Log.Error().
            Str("error", errmpp.New(C.GoString(inErr.name), uint(inErr.code)).Error()).
            Msg("VENC")
        return errmpp.New(C.GoString(inErr.name), uint(inErr.code))
    }

    return nil
}

func mppUpdateEncoderFps(id int, inFps int, outFps int) error {
    var inErr C.error_in
    var in C.mpp_venc_create_encoder_in

    C.invalidate_mpp_venc_create_encoder_in(&in)

    if inFps > vi.Fps() {
        return errors.New("In FPS can`t be more than VI fps")
    }

    if outFps > vi.Fps() {
        return errors.New("Out FPS can`t be more than VI fps")
    }

    if inFps > outFps {
        return errors.New("Out FPS can`t be more than In fps")
    }


    err2 := C.mpp_venc_update_fps2(&inErr, C.int(id), C.int(inFps), C.int(outFps))
    if err2 != 0 {
        logger.Log.Error().
            Str("error", errmpp.New(C.GoString(inErr.name), uint(inErr.code)).Error()).
            Msg("VENC")
        return errmpp.New(C.GoString(inErr.name), uint(inErr.code))
    }


    in.id       = C.int(id)
    in.in_fps   = C.int(inFps)
    in.out_fps  = C.int(outFps)

    err := C.mpp_venc_update_fps(&inErr, &in)

    if err != 0 {
        logger.Log.Error().
            Str("error", errmpp.New(C.GoString(inErr.name), uint(inErr.code)).Error()).
            Msg("VENC")
        return errmpp.New(C.GoString(inErr.name), uint(inErr.code))
    }

    return nil
}

func mppDestroyEncoder(id int) error {
    var inErr C.error_in
    var in C.mpp_venc_destroy_encoder_in

    in.id = C.uint(id)

    err := C.mpp_venc_destroy_encoder(&inErr, &in)

    if err != 0 {
        logger.Log.Error().
            Str("error", errmpp.New(C.GoString(inErr.name), uint(inErr.code)).Error()).
            Msg("VENC")
        return errmpp.New(C.GoString(inErr.name), uint(inErr.code))
    }

    return nil
}

func mppStartEncoder(id int) error {
    var inErr C.error_in
    var in C.mpp_venc_start_encoder_in

    in.id = C.uint(id)

    err := C.mpp_venc_start_encoder(&inErr, &in)

    if err != 0 {
        logger.Log.Error().
            Str("error", errmpp.New(C.GoString(inErr.name), uint(inErr.code)).Error()).
            Msg("VENC")
        return errmpp.New(C.GoString(inErr.name), uint(inErr.code))
    }

    return nil
}

func mppStopEncoder(id int) error {
    var inErr C.error_in
    var in C.mpp_venc_stop_encoder_in

    in.id = C.uint(id)

    err := C.mpp_venc_stop_encoder(&inErr, &in)

    if err != 0 {
        logger.Log.Error().
            Str("error", errmpp.New(C.GoString(inErr.name), uint(inErr.code)).Error()).
            Msg("VENC")
        return errmpp.New(C.GoString(inErr.name), uint(inErr.code))
    }

    return nil
}

func mppRequestIdr(id int) error {
    var inErr C.error_in
    var in C.mpp_venc_request_idr_in

    in.id = C.int(id)

    err := C.mpp_venc_request_idr(&inErr, &in)

    if err != 0 {
        logger.Log.Error().
            Str("error", errmpp.New(C.GoString(inErr.name), uint(inErr.code)).Error()).
            Msg("VENC")
        return errmpp.New(C.GoString(inErr.name), uint(inErr.code))
    }

    return nil
}

func mppReset(id int) error {
    var inErr C.error_in
    var in C.mpp_venc_reset_in

    in.id = C.int(id)

    err := C.mpp_venc_reset(&inErr, &in)

    if err != 0 {
        logger.Log.Error().
            Str("error", errmpp.New(C.GoString(inErr.name), uint(inErr.code)).Error()).
            Msg("VENC")
        return errmpp.New(C.GoString(inErr.name), uint(inErr.code))
    }

    return nil
}

func mppSendFrameToEncoder(id int, f connection.Frame) error {
    var inErr C.error_in
    var in C.mpp_send_frame_to_encoder_in

    //var tmp unsafe.Pointer = unsafe.Pointer(&f.FrameMPP)

    in.id = C.int(id)
    in.frame = f.Frame //TODO!

    err := C.mpp_send_frame_to_encoder(&inErr, &in, unsafe.Pointer(&f.FrameMPP))
    if err != 0 {
        logger.Log.Error().
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
