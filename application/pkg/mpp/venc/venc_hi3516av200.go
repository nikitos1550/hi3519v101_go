//+build arm
//+build hi3516av200

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
    stMjpegAttr.u32MaxPicWidth      = 3840;
    stMjpegAttr.u32MaxPicHeight     = 2160;
    stMjpegAttr.u32PicWidth         = width;
    stMjpegAttr.u32PicHeight        = height;
    stMjpegAttr.u32BufSize          = 3840*2160*3;
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

    return ERR_NONE;
}

int mpp3_venc_sample_h264(unsigned int *error_code, int width, int height, int bitrate, int channelId) {
    *error_code = 0;

    VENC_CHN_ATTR_S stVencChnAttr;
    VENC_ATTR_H264_S stH264Attr;
    VENC_ATTR_H264_CBR_S    stH264Cbr;

    stVencChnAttr.stVeAttr.enType = PT_H264;

    stH264Attr.u32MaxPicWidth   = 3840;
    stH264Attr.u32MaxPicHeight  = 2160;
    stH264Attr.u32PicWidth      = width;         //the picture width
    stH264Attr.u32PicHeight     = height;         //the picture height
    stH264Attr.u32BufSize       = 3840*2160*1.5;  //stream buffer size
    stH264Attr.u32Profile       = 1;            //0: baseline; 1:MP; 2:HP;  3:svc_t
    stH264Attr.bByFrame         = HI_TRUE;      //get stream mode is slice mode or frame mode?
    //stH264Attr.u32BFrameNum   = 0;            //0: not support B frame; >=1: number of B frames
    //stH264Attr.u32RefNum      = 1;            //0: default; number of refrence frame

    memcpy(&stVencChnAttr.stVeAttr.stAttrH264e, &stH264Attr, sizeof(VENC_ATTR_H264_S));

    stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_H264CBR;

    stH264Cbr.u32Gop            = 25;
    stH264Cbr.u32StatTime       = 1;
    stH264Cbr.u32SrcFrmRate     = 30;       //input (vi) frame rate
    stH264Cbr.fr32DstFrmRate    = 25;//30;       //target frame rate
    stH264Cbr.u32BitRate        = bitrate;
    stH264Cbr.u32FluctuateLevel = 1;        //average bit rate

    memcpy(&stVencChnAttr.stRcAttr.stAttrH264Cbr, &stH264Cbr, sizeof(VENC_ATTR_H264_CBR_S));

    stVencChnAttr.stGopAttr.enGopMode               = VENC_GOPMODE_NORMALP;
    stVencChnAttr.stGopAttr.stNormalP.s32IPQpDelta  = 2;

    *error_code = HI_MPI_VENC_CreateChn(channelId, &stVencChnAttr);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    *error_code = HI_MPI_VENC_StartRecvPic(channelId);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    return ERR_NONE;
}

int mpp3_venc_sample_h265(unsigned int *error_code, int width, int height, int bitrate, int channelId) {
    *error_code = 0;

    VENC_CHN_ATTR_S stVencChnAttr;
    VENC_ATTR_H265_S stH265Attr;
    VENC_ATTR_H265_CBR_S    stH265Cbr;

    stVencChnAttr.stVeAttr.enType = PT_H265;

    stH265Attr.u32MaxPicWidth   = 3840;
    stH265Attr.u32MaxPicHeight  = 2160;
    stH265Attr.u32PicWidth      = width;         //the picture width
    stH265Attr.u32PicHeight     = height;         //the picture height
    stH265Attr.u32BufSize       = 3840*2160*1.5;  //stream buffer size
    stH265Attr.u32Profile       = 0;            //0: MP
    stH265Attr.bByFrame         = HI_TRUE;      //get stream mode is slice mode or frame mode?

    memcpy(&stVencChnAttr.stVeAttr.stAttrH265e, &stH265Attr, sizeof(VENC_ATTR_H265_S));

    stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_H265CBR;

    stH265Cbr.u32Gop            = 25;
    stH265Cbr.u32StatTime       = 1;
    stH265Cbr.u32SrcFrmRate     = 30;       //input (vi) frame rate
    stH265Cbr.fr32DstFrmRate    = 25;//30;       //target frame rate
    stH265Cbr.u32BitRate        = bitrate;
    stH265Cbr.u32FluctuateLevel = 1;        //average bit rate

    memcpy(&stVencChnAttr.stRcAttr.stAttrH265Cbr, &stH265Cbr, sizeof(VENC_ATTR_H265_CBR_S));

    stVencChnAttr.stGopAttr.enGopMode               = VENC_GOPMODE_NORMALP;
    stVencChnAttr.stGopAttr.stNormalP.s32IPQpDelta  = 2;

    *error_code = HI_MPI_VENC_CreateChn(channelId, &stVencChnAttr);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    *error_code = HI_MPI_VENC_StartRecvPic(channelId);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    return ERR_NONE;
}

int mpp3_venc_delete_encoder(unsigned int *error_code, int channelId) {
	*error_code = 0;

    *error_code = HI_MPI_VENC_StopRecvPic(channelId);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

	*error_code = HI_MPI_VENC_DestroyChn(channelId);
	if (*error_code != HI_SUCCESS) return ERR_MPP;

	return ERR_NONE;
}
*/
import "C"
	