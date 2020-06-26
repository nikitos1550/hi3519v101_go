//+build arm
//+build hi3516cv100

package venc

/*
#include "../include/mpp.h"
#include "../errmpp/errmpp.h"
#include "../../logger/logger.h"

#include <string.h>


typedef struct hi3516cv100_venc_create_mjpeg_in_struct {
    unsigned int venc_id;
    unsigned int width;
    unsigned int height;
    unsigned int bitrate;
    
    unsigned int vpss_fps;
    unsigned int fps;
} hi3516cv100_venc_create_mjpeg_in;

int hi3516cv100_venc_sample_mjpeg(error_in *err, hi3516cv100_venc_create_mjpeg_in *in) {
    //HI_S32 s32Ret;
    VENC_CHN_ATTR_S stVencChnAttr;
    VENC_ATTR_MJPEG_S stMjpegAttr;
                       
    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VENC_CreateGroup, in->venc_id);

    stVencChnAttr.stVeAttr.enType = PT_MJPEG;

    stMjpegAttr.u32MaxPicWidth = in->width; //stPicSize.u32Width;
    stMjpegAttr.u32MaxPicHeight = in->height; //stPicSize.u32Height;
    stMjpegAttr.u32PicWidth = in->width; //stPicSize.u32Width;
    stMjpegAttr.u32PicHeight = in->height; //stPicSize.u32Height;
    stMjpegAttr.u32BufSize = in->width * in->height * 2;
    stMjpegAttr.bByFrame = HI_TRUE;     //get stream mode is field mode  or frame mode
    stMjpegAttr.bMainStream = HI_TRUE;  //main stream or minor stream types?
    stMjpegAttr.bVIField = HI_FALSE;    //the sign of the VI picture is field or frame?
    stMjpegAttr.u32Priority = 0;        //channels precedence level
    
    memcpy(&stVencChnAttr.stVeAttr.stAttrMjpeg, &stMjpegAttr, sizeof(VENC_ATTR_MJPEG_S));


    stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_MJPEGCBR;
    stVencChnAttr.stRcAttr.stAttrMjpegeCbr.u32StatTime       = 1;
    stVencChnAttr.stRcAttr.stAttrMjpegeCbr.u32ViFrmRate      = 30;  //(VIDEO_ENCODING_MODE_PAL== enNorm)?25:30;
    stVencChnAttr.stRcAttr.stAttrMjpegeCbr.fr32TargetFrmRate = 1;   //(VIDEO_ENCODING_MODE_PAL== enNorm)?25:30;
    stVencChnAttr.stRcAttr.stAttrMjpegeCbr.u32FluctuateLevel = 0;
    stVencChnAttr.stRcAttr.stAttrMjpegeCbr.u32BitRate = 1024*1;

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VENC_CreateChn, in->venc_id, &stVencChnAttr);
    
    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VENC_RegisterChn, in->venc_id, in->venc_id);//HI_MPI_VENC_RegisterChn(VencGrp, VencChn);
    
    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VENC_StartRecvPic, in->venc_id);

    return ERR_NONE;
}

static int hi3516cv100_venc_delete_encoder(error_in *err, unsigned int venc_id) {

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VENC_StopRecvPic, venc_id);
 
    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VENC_DestroyChn, venc_id);

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
    //var err C.int

    switch encoder.Format {
    //case "h264":
    //    var in C.hi3516cv100_venc_create_h264_in
    //
    //    in.venc_id = C.uint(encoder.VencId)
    //    in.width = C.uint(encoder.Width)
    //    in.height = C.uint(encoder.Height)
    //    in.bitrate = C.uint(encoder.Bitrate)
    //
    //    err = C.hi3516cv100_venc_create_h264(&inErr, &in)
    //case "h265":
    //    var in C.hi3516cv100_venc_create_h265_in
    //
    //    in.venc_id = C.uint(encoder.VencId)
    //    in.width = C.uint(encoder.Width)
    //    in.height = C.uint(encoder.Height)
    //    in.bitrate = C.uint(encoder.Bitrate)
    //
    //    err = C.hi3516cv100_venc_create_h265(&inErr, &in)
    case "mjpeg":
        var in C.hi3516cv100_venc_create_mjpeg_in

        in.venc_id = C.uint(encoder.VencId)
        in.width = C.uint(encoder.Width)
        in.height = C.uint(encoder.Height)
        in.bitrate = C.uint(encoder.Bitrate)

        logger.Log.Trace().
            Uint("venc_id", uint(in.venc_id)).
            Uint("width", uint(in.width)).
            Uint("height", uint(in.height)).
            Uint("bitrate", uint(in.bitrate)).
            Msg("VENC create mjpeg")

        err := C.hi3516cv100_venc_sample_mjpeg(&inErr, &in)
        if err != C.ERR_NONE {
            logger.Log.Fatal().
                Str("error", errmpp.New(C.GoString(inErr.name), uint(inErr.code)).Error()).
                Msg("VENC create mjpeg failed")
        }
    default:
        logger.Log.Fatal().
            Str("codec", encoder.Format).
            Msg("VENC unknown codec")
        return errors.New("VENC unknown codec")
    }

    //if err != C.ERR_NONE {
    //    return errmpp.New(C.GoString(inErr.name), uint(inErr.code))
    //}

    addVenc(encoder.VencId)

    return nil
}

func deleteVencEncoder(encoder ActiveEncoder) error {
    var inErr C.error_in
    var err C.int

    delVenc(encoder.VencId) //first we remove fd from loop

    err = C.hi3516cv100_venc_delete_encoder(&inErr, C.uint(encoder.VencId))

    if err != C.ERR_NONE {
        return errmpp.New(C.GoString(inErr.name), uint(inErr.code))
    }

    return nil
}

