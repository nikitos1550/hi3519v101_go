//+build nobuild

//+build arm
//+build hi3516cv300

package venc

/*
#include "../include/mpp.h"
#include "../errmpp/errmpp.h"
#include "../../logger/logger.h"

#include <string.h>

typedef struct hi3516cv300_venc_create_mjpeg_in_struct {
    unsigned int venc_id;
    unsigned int width;
    unsigned int height;
    unsigned int bitrate;
    
    unsigned int vpss_fps;
    unsigned int fps;
} hi3516cv300_venc_create_mjpeg_in;

static int hi3516cv300_venc_sample_mjpeg(error_in *err, hi3516cv300_venc_create_mjpeg_in *in) {
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
        RETURN_ERR_MPP(ERR_F_HI_MPI_VENC_CreateChn, mpp_error_code);
    }

    mpp_error_code = HI_MPI_VENC_StartRecvPic(in->venc_id);
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_VENC_StartRecvPic, mpp_error_code);
    }

    return ERR_NONE;
}

static int hi3516cv300_venc_delete_encoder(error_in *err, unsigned int venc_id) {
    unsigned int mpp_error_code = 0;

    mpp_error_code = HI_MPI_VENC_StopRecvPic(venc_id);
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_VENC_StopRecvPic, mpp_error_code);
    }

    mpp_error_code = HI_MPI_VENC_DestroyChn(venc_id);
    if (mpp_error_code != HI_SUCCESS) {
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
        /*
    case "h264":
        var in C.hi3516cv300_venc_create_h264_in

        in.venc_id = C.uint(encoder.VencId)
        in.width = C.uint(encoder.Width)
        in.height = C.uint(encoder.Height)
        in.bitrate = C.uint(encoder.Bitrate)

        err = C.hi3516cv300_venc_create_h264(&inErr, &in)
    case "h265":
        var in C.hi3516cv300_venc_create_h265_in
        
        in.venc_id = C.uint(encoder.VencId)
        in.width = C.uint(encoder.Width)
        in.height = C.uint(encoder.Height)
        in.bitrate = C.uint(encoder.Bitrate)

        err = C.hi3516cv300_venc_create_h265(&inErr, &in)
        */
    case "mjpeg":
        var in C.hi3516cv300_venc_create_mjpeg_in

        in.venc_id = C.uint(encoder.VencId)
        in.width = C.uint(encoder.Width)
        in.height = C.uint(encoder.Height)
        in.bitrate = C.uint(encoder.Bitrate)

        err = C.hi3516cv300_venc_sample_mjpeg(&inErr, &in)
    default:
        logger.Log.Fatal().
            Str("codec", encoder.Format).
            Msg("VENC unknown codec")
        return errors.New("VENC unknown codec")
    }

    if err != C.ERR_NONE {
        return errmpp.New(uint(inErr.f), uint(inErr.mpp))
    }

    addVenc(encoder.VencId)

    return nil
}

func deleteVencEncoder(encoder ActiveEncoder) error {
    var inErr C.error_in
    var err C.int

    delVenc(encoder.VencId) //first we remove fd from loop

    err = C.hi3516cv300_venc_delete_encoder(&inErr, C.uint(encoder.VencId))

    if err != C.ERR_NONE {
        return errmpp.New(uint(inErr.f), uint(inErr.mpp))
    }

    return nil
}


