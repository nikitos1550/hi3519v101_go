//+build arm
//+build hi3516av200

package venc

/*
#include "../include/mpp.h"
#include "../errmpp/error.h"
#include "../../logger/logger.h"

#include <string.h>

typedef struct hi3516av200_venc_create_mjpeg_in_struct {
	unsigned int venc_id;
	unsigned int width;
	unsigned int height;
	unsigned int bitrate;
	
	unsigned int vpss_fps;
	unsigned int fps;
} hi3516av200_venc_create_mjpeg_in;

static int hi3516av200_venc_sample_mjpeg(error_in *err, hi3516av200_venc_create_mjpeg_in *in) {
    unsigned int mpp_error_code = 0;

	VENC_ATTR_MJPEG_S stMjpegAttr;

    stMjpegAttr.u32MaxPicWidth      = in->width;
    stMjpegAttr.u32MaxPicHeight     = in->height;
    stMjpegAttr.u32PicWidth         = in->width;
    stMjpegAttr.u32PicHeight        = in->height;
    stMjpegAttr.u32BufSize          = in->width * in->height * 3; //3840*2160*3;
    stMjpegAttr.bByFrame            = HI_TRUE;

	VENC_ATTR_MJPEG_CBR_S stMjpegCbr;

    stMjpegCbr.u32SrcFrmRate        = 30;   //in->vpss_fps; TODO
    stMjpegCbr.fr32DstFrmRate       = 1;    //in->fps; TODO
    stMjpegCbr.u32BitRate           = in->bitrate;
	stMjpegCbr.u32StatTime          = 1;
    stMjpegCbr.u32FluctuateLevel    = 1;

    VENC_CHN_ATTR_S stVencChnAttr;  

    stVencChnAttr.stVeAttr.enType   = PT_MJPEG;
    stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_MJPEGCBR;
	memcpy(&stVencChnAttr.stVeAttr.stAttrH264e, &stMjpegAttr, sizeof(VENC_ATTR_MJPEG_S));
    memcpy(&stVencChnAttr.stRcAttr.stAttrMjpegeCbr, &stMjpegCbr, sizeof(VENC_ATTR_MJPEG_CBR_S));
    stVencChnAttr.stGopAttr.enGopMode  = VENC_GOPMODE_NORMALP;
    stVencChnAttr.stGopAttr.stNormalP.s32IPQpDelta = 0;

    mpp_error_code = HI_MPI_VENC_CreateChn(in->venc_id, &stVencChnAttr);
    if (mpp_error_code != HI_SUCCESS) {
        //GO_LOG_VENC(LOGGER_ERROR, "HI_MPI_VENC_CreateChn")
        //err->mpp = mpp_error_code;
		//return ERR_MPP;
        RETURN_ERR_MPP(ERR_F_HI_MPI_VENC_CreateChn, mpp_error_code);
	}

    mpp_error_code = HI_MPI_VENC_StartRecvPic(in->venc_id);
    if (mpp_error_code != HI_SUCCESS) {
        //GO_LOG_VENC(LOGGER_ERROR, "HI_MPI_VENC_StartRecvPic")
        //err->mpp = mpp_error_code;
		//return ERR_MPP;
        RETURN_ERR_MPP(ERR_F_HI_MPI_VENC_StartRecvPic, mpp_error_code);
	}

    return ERR_NONE;
}

typedef struct hi3516av200_venc_create_h264_in_struct {
    unsigned int venc_id; 
    unsigned int width;
    unsigned int height;
    unsigned int bitrate;
    
    unsigned int vpss_fps;
    unsigned int fps;
} hi3516av200_venc_create_h264_in;

static int hi3516av200_venc_create_h264(error_in *err, hi3516av200_venc_create_h264_in *in) {
    unsigned int mpp_error_code = 0;

    VENC_ATTR_H264_S stH264Attr;

    stH264Attr.u32MaxPicWidth   = in->width;
    stH264Attr.u32MaxPicHeight  = in->height;
    stH264Attr.u32PicWidth      = in->width; 
    stH264Attr.u32PicHeight     = in->height;
    stH264Attr.u32BufSize       = in->width * in->height * 1.5;
    stH264Attr.u32Profile       = 1;            //0: baseline; 1:MP; 2:HP;  3:svc_t
    stH264Attr.bByFrame         = HI_TRUE;      //get stream mode is slice mode or frame mode?
    //stH264Attr.u32BFrameNum   = 0;            //0: not support B frame; >=1: number of B frames
    //stH264Attr.u32RefNum      = 1;            //0: default; number of refrence frame

    VENC_ATTR_H264_CBR_S    stH264Cbr;

    stH264Cbr.u32Gop            = 30;   //in->fps; TODO
    stH264Cbr.u32SrcFrmRate     = 30;   //in->vpss_fps; TODO
    stH264Cbr.fr32DstFrmRate    = 30;   //in->fps; TODO
    stH264Cbr.u32BitRate        = in->bitrate;
    stH264Cbr.u32StatTime       = 1;
    stH264Cbr.u32FluctuateLevel = 1; 

    VENC_CHN_ATTR_S stVencChnAttr;

    stVencChnAttr.stVeAttr.enType   = PT_H264;
    stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_H264CBR;
    memcpy(&stVencChnAttr.stVeAttr.stAttrH264e, &stH264Attr, sizeof(VENC_ATTR_H264_S));
    memcpy(&stVencChnAttr.stRcAttr.stAttrH264Cbr, &stH264Cbr, sizeof(VENC_ATTR_H264_CBR_S));
    stVencChnAttr.stGopAttr.enGopMode               = VENC_GOPMODE_NORMALP;
    stVencChnAttr.stGopAttr.stNormalP.s32IPQpDelta  = 2;

    mpp_error_code = HI_MPI_VENC_CreateChn(in->venc_id, &stVencChnAttr);
    if (mpp_error_code != HI_SUCCESS) {
        //GO_LOG_VENC(LOGGER_ERROR, "HI_MPI_VENC_CreateChn")
        //err->mpp = mpp_error_code;    
        //return ERR_MPP;
        RETURN_ERR_MPP(ERR_F_HI_MPI_VENC_CreateChn, mpp_error_code);
    }

    mpp_error_code = HI_MPI_VENC_StartRecvPic(in->venc_id);
    if (mpp_error_code != HI_SUCCESS) {
        //GO_LOG_VENC(LOGGER_ERROR, "HI_MPI_VENC_StartRecvPic")
        //err->mpp = mpp_error_code;
        //return ERR_MPP;
        RETURN_ERR_MPP(ERR_F_HI_MPI_VENC_StartRecvPic, mpp_error_code);
    }

    return ERR_NONE;
}

typedef struct hi3516av200_venc_create_h265_in_struct {
    unsigned int venc_id;
    unsigned int width;
    unsigned int height;
    unsigned int bitrate;
    
    unsigned int vpss_fps;
    unsigned int fps;
} hi3516av200_venc_create_h265_in;

static int hi3516av200_venc_create_h265(error_in *err, hi3516av200_venc_create_h265_in *in) {
    unsigned int mpp_error_code = 0;

    VENC_ATTR_H265_S stH265Attr;

    stH265Attr.u32MaxPicWidth   = in->width;
    stH265Attr.u32MaxPicHeight  = in->height;
    stH265Attr.u32PicWidth      = in->width; 
    stH265Attr.u32PicHeight     = in->height;
    stH265Attr.u32BufSize       = in->width * in->height * 1.5;
    stH265Attr.u32Profile       = 0;            //0: MP
    stH265Attr.bByFrame         = HI_TRUE;

    VENC_ATTR_H265_CBR_S    stH265Cbr;

    stH265Cbr.u32Gop            = 30;   //in->fps; TODO
    stH265Cbr.u32SrcFrmRate     = 30;   //in->vpss_fps; TODO
    stH265Cbr.fr32DstFrmRate    = 30;   //in->fps; TODO
    stH265Cbr.u32BitRate        = in->bitrate;
    stH265Cbr.u32StatTime       = 1;
    stH265Cbr.u32FluctuateLevel = 1;

    VENC_CHN_ATTR_S stVencChnAttr;
    stVencChnAttr.stVeAttr.enType = PT_H265;
    stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_H265CBR;
    memcpy(&stVencChnAttr.stVeAttr.stAttrH265e, &stH265Attr, sizeof(VENC_ATTR_H265_S));
    memcpy(&stVencChnAttr.stRcAttr.stAttrH265Cbr, &stH265Cbr, sizeof(VENC_ATTR_H265_CBR_S));

    stVencChnAttr.stGopAttr.enGopMode               = VENC_GOPMODE_NORMALP;
    stVencChnAttr.stGopAttr.stNormalP.s32IPQpDelta  = 2;

    mpp_error_code = HI_MPI_VENC_CreateChn(in->venc_id, &stVencChnAttr);
    if (mpp_error_code != HI_SUCCESS) {
        //GO_LOG_VENC(LOGGER_ERROR, "HI_MPI_VENC_CreateChn")
        //err->mpp = mpp_error_code;        
        //return ERR_MPP;
        RETURN_ERR_MPP(ERR_F_HI_MPI_VENC_CreateChn, mpp_error_code);
    }

    mpp_error_code = HI_MPI_VENC_StartRecvPic(in->venc_id);
    if (mpp_error_code != HI_SUCCESS) {
        //GO_LOG_VENC(LOGGER_ERROR, "HI_MPI_VENC_StartRecvPic")
        //err->mpp = mpp_error_code;
        //return ERR_MPP;
        RETURN_ERR_MPP(ERR_F_HI_MPI_VENC_StartRecvPic, mpp_error_code);
    }

    return ERR_NONE;
}

static int hi3516av200_venc_delete_encoder(error_in *err, unsigned int venc_id) {
	unsigned int mpp_error_code = 0;

    mpp_error_code = HI_MPI_VENC_StopRecvPic(venc_id);
    if (mpp_error_code != HI_SUCCESS) {
        //GO_LOG_VENC(LOGGER_ERROR, "HI_MPI_VENC_StopRecvPic")
        //err->mpp = mpp_error_code;
        //return ERR_MPP;
        RETURN_ERR_MPP(ERR_F_HI_MPI_VENC_StopRecvPic, mpp_error_code);
    }

	mpp_error_code = HI_MPI_VENC_DestroyChn(venc_id);
	if (mpp_error_code != HI_SUCCESS) {
        //GO_LOG_VENC(LOGGER_ERROR, "HI_MPI_VENC_DestroyChn")
        //err->mpp = mpp_error_code;
        //return ERR_MPP;
        RETURN_ERR_MPP(ERR_F_HI_MPI_VENC_DestroyChn, mpp_error_code);
    }

	return ERR_NONE;
}
*/
import "C"

import (
    "errors"
    "application/pkg/logger"
    "application/pkg/mpp/errmpp"
)

func createVencEncoder(encoder ActiveEncoder) error {
    var inErr C.error_in
    var err C.int

    switch encoder.Format {
    case "h264":
        var in C.hi3516av200_venc_create_h264_in

        in.venc_id = C.uint(encoder.VencId)
        in.width = C.uint(encoder.Width)
        in.height = C.uint(encoder.Height)
        in.bitrate = C.uint(encoder.Bitrate)

        err = C.hi3516av200_venc_create_h264(&inErr, &in)
    case "h265":
        var in C.hi3516av200_venc_create_h265_in
        
        in.venc_id = C.uint(encoder.VencId)
        in.width = C.uint(encoder.Width)
        in.height = C.uint(encoder.Height)
        in.bitrate = C.uint(encoder.Bitrate)

        err = C.hi3516av200_venc_create_h265(&inErr, &in)
    case "mjpeg":
        var in C.hi3516av200_venc_create_mjpeg_in

        in.venc_id = C.uint(encoder.VencId)
        in.width = C.uint(encoder.Width)
        in.height = C.uint(encoder.Height)
        in.bitrate = C.uint(encoder.Bitrate)

        err = C.hi3516av200_venc_sample_mjpeg(&inErr, &in)
    default:
        logger.Log.Fatal().
            Str("codec", encoder.Format).
            Msg("VENC unknown codec")
        return errors.New("VENC unknown codec")
    }

    if err != C.ERR_NONE {
        return errmpp.New(uint(inErr.f), uint(inErr.mpp))
        //logger.Log.Fatal(). //log temporary, should generate and return error
        //    Str("error", errmpp.New("funcname", uint(inErr.mpp)).Error()).
        //    Msg("VENC")
    }

    addVenc(encoder.VencId)

    return nil
}

func deleteVencEncoder(encoder ActiveEncoder) error {
    var inErr C.error_in
    var err C.int

    delVenc(encoder.VencId) //first we remove fd from loop

    err = C.hi3516av200_venc_delete_encoder(&inErr, C.uint(encoder.VencId))

    if err != C.ERR_NONE {
        return errmpp.New(uint(inErr.f), uint(inErr.mpp))
        //logger.Log.Fatal(). //log temporary, should generate and return error
        //    Str("error", errmpp.New("funcname", uint(inErr.mpp)).Error()).
        //    Msg("VENC")
    }

    return nil
}

