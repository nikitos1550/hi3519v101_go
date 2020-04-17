//+build arm
//+build hi3516cv300

package venc

/*
#include "../include/mpp_v3.h"

#include <string.h>

#define ERR_NONE                0
#define ERR_MPP                 2

int mpp3_venc_sample_mjpeg(unsigned int *error_code, int width, int height, int bitrate, int channelId) {
    *error_code = 0;

    VENC_CHN_ATTR_S stVencChnAttr;
    VENC_ATTR_MJPEG_S stMjpegAttr;
    VENC_ATTR_MJPEG_CBR_S stMjpegCbr;

    stVencChnAttr.stVeAttr.enType   = PT_MJPEG;
    stMjpegAttr.u32MaxPicWidth      = 1920;
    stMjpegAttr.u32MaxPicHeight     = 1080;
    stMjpegAttr.u32PicWidth         = width;
    stMjpegAttr.u32PicHeight        = height;
    stMjpegAttr.u32BufSize          = 1920*1080*3;
    stMjpegAttr.bByFrame            = HI_TRUE;

    memcpy(&stVencChnAttr.stVeAttr.stAttrH264e, &stMjpegAttr, sizeof(VENC_ATTR_MJPEG_S));

    stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_MJPEGCBR;
    stMjpegCbr.u32StatTime          = 1;
    stMjpegCbr.u32SrcFrmRate        = 30;//30;// input (vi) frame rate
    stMjpegCbr.fr32DstFrmRate       = 1;//30;// target frame rate
    stMjpegCbr.u32BitRate           = bitrate;
    stMjpegCbr.u32FluctuateLevel    = 1; // average bit rate

    memcpy(&stVencChnAttr.stRcAttr.stAttrMjpegeCbr, &stMjpegCbr, sizeof(VENC_ATTR_MJPEG_CBR_S));

    stVencChnAttr.stGopAttr.enGopMode  = VENC_GOPMODE_NORMALP;
    stVencChnAttr.stGopAttr.stNormalP.s32IPQpDelta = 0;

    *error_code = HI_MPI_VENC_CreateChn(channelId, &stVencChnAttr);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    *error_code = HI_MPI_VENC_StartRecvPic(channelId);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    MPP_CHN_S stSrcChn;
    MPP_CHN_S stDestChn;

    stSrcChn.enModId    = HI_ID_VPSS;
    stSrcChn.s32DevId   = 0;
    stSrcChn.s32ChnId   = 0;
    stDestChn.enModId   = HI_ID_VENC;
    stDestChn.s32DevId  = 0;
    stDestChn.s32ChnId  = channelId;

    *error_code = HI_MPI_SYS_Bind(&stSrcChn, &stDestChn);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    return ERR_NONE;
}


int mpp3_venc_sample_h264(unsigned int *error_code, int width, int height, int bitrate, int channelId) {
    *error_code = 0;

    VENC_CHN_ATTR_S stVencChnAttr;
    VENC_ATTR_H264_S stH264Attr;
    VENC_ATTR_H264_CBR_S    stH264Cbr;
VENC_ATTR_H264_VBR_S    stH264Vbr;
VENC_ATTR_H264_AVBR_S    stH264AVbr;
VENC_ATTR_H264_FIXQP_S  stH264FixQp;

  VENC_ATTR_H265_S        stH265Attr;
    VENC_ATTR_H265_CBR_S    stH265Cbr;
    VENC_ATTR_H265_VBR_S    stH265Vbr;
    VENC_ATTR_H265_FIXQP_S  stH265FixQp;


   stVencChnAttr.stVeAttr.enType = PT_H264;

            stH264Attr.u32MaxPicWidth = 1920;
            stH264Attr.u32MaxPicHeight = 1080;
            stH264Attr.u32PicWidth = width;//the picture width
            stH264Attr.u32PicHeight = height;//the picture height
            stH264Attr.u32BufSize  = 1920 * 1080 * 2;//stream buffer size
            stH264Attr.u32Profile  = 2;//0: baseline; 1:MP; 2:HP;  3:svc_t 
            stH264Attr.bByFrame = HI_TRUE;//get stream mode is slice mode or frame mode?
            //stH264Attr.u32BFrameNum = 0;// 0: not support B frame; >=1: number of B frames 
            //stH264Attr.u32RefNum = 1;// 0: default; number of refrence frame
            memcpy(&stVencChnAttr.stVeAttr.stAttrH264e, &stH264Attr, sizeof(VENC_ATTR_H264_S));
  stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_H264CBR;
                stH264Cbr.u32Gop            = 30;//fps;
                stH264Cbr.u32StatTime       = 1; // stream rate statics time(s) 
                stH264Cbr.u32SrcFrmRate      = 30;//fps; // input (vi) frame rate 
                stH264Cbr.fr32DstFrmRate = 30;//fps; // target frame rate 
stH264Cbr.u32BitRate = bitrate;//1024*1;//1024*1;//30;
          
        stH264Cbr.u32FluctuateLevel = 1; // average bit rate 
         memcpy(&stVencChnAttr.stRcAttr.stAttrH264Cbr, &stH264Cbr, sizeof(VENC_ATTR_H264_CBR_S));


   stVencChnAttr.stGopAttr.enGopMode  = VENC_GOPMODE_NORMALP;
    stVencChnAttr.stGopAttr.stNormalP.s32IPQpDelta = 0;

     *error_code = HI_MPI_VENC_CreateChn(channelId, &stVencChnAttr);
     if (*error_code != HI_SUCCESS) return ERR_MPP;

    *error_code = HI_MPI_VENC_StartRecvPic(channelId);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    MPP_CHN_S stSrcChn;
    MPP_CHN_S stDestChn;

    stSrcChn.enModId = HI_ID_VPSS;
    stSrcChn.s32DevId = 0;
    stSrcChn.s32ChnId = 0;

    stDestChn.enModId = HI_ID_VENC;
    stDestChn.s32DevId = 0;
    stDestChn.s32ChnId = channelId;

    *error_code = HI_MPI_SYS_Bind(&stSrcChn, &stDestChn);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    return ERR_NONE;
}

int mpp3_venc_delete_encoder(unsigned int *error_code, int channelId) {
    *error_code = 0;

    //HI_S32 HI_MPI_VENC_StopRecvPic(VENC_CHN VeChn);
    //HI_S32 HI_MPI_VENC_CloseFd(VENC_CHN VeChn);
    //HI_S32 HI_MPI_VENC_DestroyChn(VENC_CHN VeChn);

    MPP_CHN_S stSrcChn;
    MPP_CHN_S stDestChn;

    stSrcChn.enModId    = HI_ID_VPSS;
    stSrcChn.s32DevId   = 0;
    stSrcChn.s32ChnId   = 0;
    stDestChn.enModId   = HI_ID_VENC;
    stDestChn.s32DevId  = 0;
    stDestChn.s32ChnId  = channelId;

    *error_code = HI_MPI_SYS_UnBind(&stSrcChn, &stDestChn);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    *error_code = HI_MPI_VENC_StopRecvPic(channelId);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    *error_code = HI_MPI_VENC_DestroyChn(channelId);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    return ERR_NONE;
}

*/
import "C"

import (
	//"application/pkg/mpp/error"
	//"log"
    //"application/pkg/logger"
)
/*
var (
	SampleMjpegFrames *frames
	SampleH264Frames  *frames
	SampleH264Notify  chan int
	SampleH264Start   chan int
)
*/
/*
func SampleMjpeg() {
	var errorCode C.uint

	switch err := C.mpp3_venc_sample_mjpeg(&errorCode); err {
	case C.ERR_NONE:
		log.Println("C.mpp3_venc_sample_mjpeg() ok")
	case C.ERR_MPP:
		log.Fatal("C.mpp3_venc_sample_mjpeg() MPP error ", error.Resolve(int64(errorCode)))
	default:
		log.Fatal("Unexpected return ", err, " of C.mpp3_venc_sample_mjpeg()")
	}

	//TODO //create corresponding encoder object
	SampleMjpegFrames = CreateFrames(3)
	addVenc(1)
}

func SampleH264() {
	var errorCode C.uint

	switch err := C.mpp3_venc_sample_h264(&errorCode); err {
	case C.ERR_NONE:
		log.Println("C.mpp3_venc_sample_h264() ok")
	case C.ERR_MPP:
		log.Fatal("C.mpp3_venc_sample_h264() MPP error ", error.Resolve(int64(errorCode)))
	default:
		log.Fatal("Unexpected return ", err, " of C.mpp3_venc_sample_h264()")
	}

	//TODO //create corresponding encoder object
	SampleH264Frames = CreateFrames(10)
	SampleH264Notify = make(chan int, 10)
	SampleH264Start = make(chan int, 1)
	addVenc(0) //add venc to get loop
}
*/
/*
func deleteEncoder(encoder Encoder) {
    var errorCode C.uint
    var err C.int

    delVenc(encoder.VencId) //first we remove fd from loop

    err = C.mpp3_venc_delete_encoder(&errorCode, C.int(encoder.VencId))
    switch err {
    case C.ERR_NONE:
        //log.Println("Encoder deleted ", encoder.VencId)
        logger.Log.Debug().
            Int("vencId", encoder.VencId).
            Msg("Encoder deleted")
    case C.ERR_MPP:
        //log.Fatal("Failed to delete encoder ", encoder.VencId, " error ", error.Resolve(int64(errorCode)))
        logger.Log.Fatal().
            Int("vencId", encoder.VencId).
            Int("error", int(errorCode)).
            Str("error_code", error.Resolve(int64(errorCode))).
            Msg("Failed to delete encoder")
    default:
        //log.Fatal("Failed to delete encoder ", encoder.VencId, "Unexpected return ", err)
        logger.Log.Fatal().
            Int("error", int(err)).
            Msg("Failed to delete encoder, unexpected return")

    }

}


func createEncoder(encoder Encoder) {
    var errorCode C.uint
    var err C.int
    switch encoder.Format {
    case "h264":
        err = C.mpp3_venc_sample_h264(&errorCode, C.int(encoder.Width), C.int(encoder.Height), C.int(encoder.Bitrate), C.int(encoder.VencId))
    //case "h265":
    //    err = C.mpp3_venc_sample_h265(&errorCode, C.int(encoder.Width), C.int(encoder.Height), C.int(encoder.Bitrate), C.int(encoder.VencId))
    case "mjpeg":
        err = C.mpp3_venc_sample_mjpeg(&errorCode, C.int(encoder.Width), C.int(encoder.Height), C.int(encoder.Bitrate), C.int(encoder.VencId))
    default:
        //log.Println("Unknown encoder format ", encoder.Format)
        logger.Log.Warn().
            Str("codec", encoder.Format).
            Msg("Unknown encoder format")
    }

    switch err {
    case C.ERR_NONE:
        //log.Println("Encoder created ", encoder.Format)
        logger.Log.Debug(). //TODO encoderId
            Str("codec", encoder.Format).
            Msg("Encoder created")

    case C.ERR_MPP:
        //log.Fatal("Failed to create encoder ", encoder.Format, " error ", error.Resolve(int64(errorCode)))
        logger.Log.Fatal().
            Str("codec", encoder.Format).
            Int("error", int(errorCode)).
            Str("error-dec", error.Resolve(int64(errorCode))).
            Msg("Failed to create encoder")
    default:
        //log.Fatal("Failed to create encoder ", encoder.Format, "Unexpected return ", err)
        logger.Log.Fatal().
            Str("codec", encoder.Format).
            Int("error", int(err)).
            Msg("Failed to create encoder, unexpected return")
    }

    addVenc(encoder.VencId)
}

*/
