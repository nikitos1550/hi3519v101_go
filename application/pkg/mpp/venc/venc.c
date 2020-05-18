#include "venc.h"

#include <string.h>

int mpp_venc_create(error_in *err, mpp_venc_create_in *in) {


    VENC_CHN_ATTR_S stVencChnAttr;
    
    stVencChnAttr.stVeAttr.enType   = PT_MJPEG;
    //stVencChnAttr.stVeAttr.enType   = PT_H264;
    //stVencChnAttr.stVeAttr.enType   = PT_H265;

    stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_MJPEGCBR;

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

    //VENC_CHN_ATTR_S stVencChnAttr;  

    //stVencChnAttr.stVeAttr.enType   = PT_MJPEG;
    //stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_MJPEGCBR;

    memcpy(&stVencChnAttr.stVeAttr.stAttrH264e, &stMjpegAttr, sizeof(VENC_ATTR_MJPEG_S));
    memcpy(&stVencChnAttr.stRcAttr.stAttrMjpegeCbr, &stMjpegCbr, sizeof(VENC_ATTR_MJPEG_CBR_S));
    
    #if defined(HI_MPP_V3)
    stVencChnAttr.stGopAttr.enGopMode  = VENC_GOPMODE_NORMALP;
    stVencChnAttr.stGopAttr.stNormalP.s32IPQpDelta = 0;
    #endif

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VENC_CreateChn, in->venc_id, &stVencChnAttr);

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VENC_StartRecvPic, in->venc_id);

    return ERR_NONE;
}

