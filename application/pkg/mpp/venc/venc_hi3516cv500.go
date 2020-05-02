//+build arm
//+build hi3516cv500 hi3516ev200

package venc

/*
#include "../include/mpp.h"
#include "../errmpp/errmpp.h"
#include "../../logger/logger.h"

#include <string.h>


typedef struct hi3516cv500_venc_create_mjpeg_in_struct {
    unsigned int venc_id;
    unsigned int width;
    unsigned int height;
    unsigned int bitrate;
    
    unsigned int vpss_fps;
    unsigned int fps;
} hi3516cv500_venc_create_mjpeg_in;

int hi3516cv500_venc_sample_mjpeg(error_in *err, hi3516cv500_venc_create_mjpeg_in *in) {

    VENC_CHN_ATTR_S        stVencChnAttr;

    stVencChnAttr.stVencAttr.enType          = PT_MJPEG;
    stVencChnAttr.stVencAttr.u32MaxPicWidth  = in->width;   //stPicSize.u32Width;
    stVencChnAttr.stVencAttr.u32MaxPicHeight = in->height;  //stPicSize.u32Height;
    stVencChnAttr.stVencAttr.u32PicWidth     = in->width;   //stPicSize.u32Width;//the picture width
    stVencChnAttr.stVencAttr.u32PicHeight    = in->height;  //stPicSize.u32Height;//the picture height
    stVencChnAttr.stVencAttr.u32BufSize      = in->width * in->height * 2;  //stream buffer size
    stVencChnAttr.stVencAttr.u32Profile      = 0;
    stVencChnAttr.stVencAttr.bByFrame        = HI_TRUE;     //get stream mode is slice mode or frame mode?


    VENC_MJPEG_CBR_S stMjpegeCbr;

    stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_MJPEGCBR;
    stMjpegeCbr.u32StatTime         = 1; //u32StatTime;
    stMjpegeCbr.u32SrcFrameRate     = 30; //in->?;
    stMjpegeCbr.fr32DstFrameRate    = 1; //in->fps;
    stMjpegeCbr.u32BitRate          = 1024*1;

    memcpy(&stVencChnAttr.stRcAttr.stMjpegCbr, &stMjpegeCbr,sizeof(VENC_MJPEG_CBR_S));

    stVencChnAttr.stGopAttr.enGopMode  = VENC_GOPMODE_NORMALP;
    stVencChnAttr.stGopAttr.stNormalP.s32IPQpDelta = 0;

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VENC_CreateChn, in->venc_id, &stVencChnAttr);

    VENC_RECV_PIC_PARAM_S  stRecvParam;
    
    stRecvParam.s32RecvPicNum = -1;
    
    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VENC_StartRecvFrame, in->venc_id, &stRecvParam);
    
    //MPP_CHN_S stSrcChn;
    //MPP_CHN_S stDestChn;
    //
    //stSrcChn.enModId   = HI_ID_VPSS;
    //stSrcChn.s32DevId  = 1;
    //stSrcChn.s32ChnId  = 0;
    //
    //stDestChn.enModId  = HI_ID_VENC;
    //stDestChn.s32DevId = 0;
    //stDestChn.s32ChnId = 0;
    //
    //DO_OR_RETURN_ERR_MPP(err, HI_MPI_SYS_Bind, &stSrcChn, &stDestChn);

    return ERR_NONE;
}

static int hi3516cv500_venc_delete_encoder(error_in *err, unsigned int venc_id) {


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
    case "mjpeg":
        var in C.hi3516cv500_venc_create_mjpeg_in

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

        err := C.hi3516cv500_venc_sample_mjpeg(&inErr, &in)
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

    err = C.hi3516cv500_venc_delete_encoder(&inErr, C.uint(encoder.VencId))

    if err != C.ERR_NONE {
        return errmpp.New(C.GoString(inErr.name), uint(inErr.code))
    }

    return nil
}

