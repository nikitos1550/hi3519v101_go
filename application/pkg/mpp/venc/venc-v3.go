//+build hi3516cv300 hi3516av200

package venc

/*
#include "../include/hi3516av200_mpp.h"
#include <string.h>

#define ERR_NONE                0
#define ERR_MPP                 2

int mpp3_venc_sample_mjpeg(unsigned int *error_code) {
    *error_code = 0;

    VENC_CHN_ATTR_S stVencChnAttr;
    VENC_ATTR_MJPEG_S stMjpegAttr;
    VENC_ATTR_MJPEG_CBR_S stMjpegCbr;

    stVencChnAttr.stVeAttr.enType   = PT_MJPEG;
    stMjpegAttr.u32MaxPicWidth      = 3840;
    stMjpegAttr.u32MaxPicHeight     = 2160;
    stMjpegAttr.u32PicWidth         = 3840;
    stMjpegAttr.u32PicHeight        = 2160;
    stMjpegAttr.u32BufSize          = 3840*2160*3;
    stMjpegAttr.bByFrame            = HI_TRUE;

    memcpy(&stVencChnAttr.stVeAttr.stAttrH264e, &stMjpegAttr, sizeof(VENC_ATTR_MJPEG_S));

    stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_MJPEGCBR;
    stMjpegCbr.u32StatTime          = 1;
    stMjpegCbr.u32SrcFrmRate        = 30;//30;// input (vi) frame rate
    stMjpegCbr.fr32DstFrmRate       = 1;//30;// target frame rate
    stMjpegCbr.u32BitRate           = 1024*2;
    stMjpegCbr.u32FluctuateLevel    = 1; // average bit rate

    memcpy(&stVencChnAttr.stRcAttr.stAttrMjpegeCbr, &stMjpegCbr, sizeof(VENC_ATTR_MJPEG_CBR_S));

    stVencChnAttr.stGopAttr.enGopMode  = VENC_GOPMODE_NORMALP;
    stVencChnAttr.stGopAttr.stNormalP.s32IPQpDelta = 0;

    *error_code = HI_MPI_VENC_CreateChn(1, &stVencChnAttr);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    *error_code = HI_MPI_VENC_StartRecvPic(1);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    MPP_CHN_S stSrcChn;
    MPP_CHN_S stDestChn;

    stSrcChn.enModId    = HI_ID_VPSS;
    stSrcChn.s32DevId   = 0;
    stSrcChn.s32ChnId   = 0;
    stDestChn.enModId   = HI_ID_VENC;
    stDestChn.s32DevId  = 0;
    stDestChn.s32ChnId  = 1;

    *error_code = HI_MPI_SYS_Bind(&stSrcChn, &stDestChn);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    return ERR_NONE;
}

int mpp3_venc_sample_h264(unsigned int *error_code) {
    *error_code = 0;

    VENC_CHN_ATTR_S stVencChnAttr;
    VENC_ATTR_H264_S stH264Attr;
    VENC_ATTR_H264_CBR_S    stH264Cbr;

    stVencChnAttr.stVeAttr.enType = PT_H264;

    stH264Attr.u32MaxPicWidth   = 3840;
    stH264Attr.u32MaxPicHeight  = 2160;
    stH264Attr.u32PicWidth      = 3840;         //the picture width
    stH264Attr.u32PicHeight     = 2160;         //the picture height
    stH264Attr.u32BufSize       = 3840*2160*4;  //stream buffer size
    stH264Attr.u32Profile       = 2;            //0: baseline; 1:MP; 2:HP;  3:svc_t
    stH264Attr.bByFrame         = HI_TRUE;      //get stream mode is slice mode or frame mode?
    //stH264Attr.u32BFrameNum   = 0;            //0: not support B frame; >=1: number of B frames 
    //stH264Attr.u32RefNum      = 1;            //0: default; number of refrence frame

    memcpy(&stVencChnAttr.stVeAttr.stAttrH264e, &stH264Attr, sizeof(VENC_ATTR_H264_S));

    stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_H264CBR;

    stH264Cbr.u32Gop            = 30;
    stH264Cbr.u32StatTime       = 1;
    stH264Cbr.u32SrcFrmRate     = 30;       //input (vi) frame rate
    stH264Cbr.fr32DstFrmRate    = 1;//30;       //target frame rate
    stH264Cbr.u32BitRate        = 1024*1;
    stH264Cbr.u32FluctuateLevel = 1;        //average bit rate

    memcpy(&stVencChnAttr.stRcAttr.stAttrH264Cbr, &stH264Cbr, sizeof(VENC_ATTR_H264_CBR_S));

    stVencChnAttr.stGopAttr.enGopMode               = VENC_GOPMODE_NORMALP;
    stVencChnAttr.stGopAttr.stNormalP.s32IPQpDelta  = 2;

    *error_code = HI_MPI_VENC_CreateChn(0, &stVencChnAttr);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    *error_code = HI_MPI_VENC_StartRecvPic(0);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    MPP_CHN_S stSrcChn;
    MPP_CHN_S stDestChn;

    stSrcChn.enModId    = HI_ID_VPSS;
    stSrcChn.s32DevId   = 0;
    stSrcChn.s32ChnId   = 0;
    stDestChn.enModId   = HI_ID_VENC;
    stDestChn.s32DevId  = 0;
    stDestChn.s32ChnId  = 0;

    *error_code = HI_MPI_SYS_Bind(&stSrcChn, &stDestChn);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    return ERR_NONE;
}
*/
import "C"

import (
    "log"
    "application/pkg/mpp/error"
    "application/pkg/mpp/getloop"
)

func SampleMjpeg() {
    var errorCode C.uint

    switch err := C.mpp3_venc_sample_mjpeg(&errorCode); err {
    case C.ERR_NONE:
        log.Println("C.mpp3_venc_sample_mjpeg() ok")
    case C.ERR_MPP:
        log.Fatal("C.mpp3_venc_sample_mjpeg() MPP error ", error.Resolve(int64(errorCode)))
    default:
        log.Fatal("Unexpected return ", err , " of C.mpp3_venc_sample_mjpeg()")
    }

    getloop.AddVenc(1)
}

func SampleH264() {
    var errorCode C.uint

    switch err := C.mpp3_venc_sample_h264(&errorCode); err {
    case C.ERR_NONE:
        log.Println("C.mpp3_venc_sample_h264() ok")
    case C.ERR_MPP:
        log.Fatal("C.mpp3_venc_sample_h264() MPP error ", error.Resolve(int64(errorCode)))
    default:
        log.Fatal("Unexpected return ", err , " of C.mpp3_venc_sample_h264()")
    }

    getloop.AddVenc(0)
}

