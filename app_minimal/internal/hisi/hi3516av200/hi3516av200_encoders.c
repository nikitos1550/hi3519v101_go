#include "../hisi_external.h"
#include "hi3516av200_mpp.h"

#include <string.h>

int hisi_encoders_max_num(unsigned int * num) {
    *num = VENC_MAX_CHN_NUM;
    return ERR_NONE;
}

int hisi_encoder_info   (unsigned int id,
                        struct encoder_static_params * sparams,
                        struct encoder_dynamic_params * dparams) {
    int error_code = 0;

    if (id < 0 &&
        id >= VENC_MAX_CHN_NUM) return ERR_OBJECT_NOT_FOUNT;

    VENC_CHN_ATTR_S stChnAttr;

    error_code = HI_MPI_VENC_GetChnAttr(id, &stChnAttr);
    if (error_code != HI_SUCCESS) {
        if (error_code == HI_ERR_VENC_UNEXIST) {
            return ERR_OBJECT_NOT_FOUNT;
        }
        return ERR_MPP;
    }

    switch(stChnAttr.stVeAttr.enType) {
        case PT_MJPEG:
            sparams->codec = CODEC_MJPEG;
            switch(stChnAttr.stRcAttr.enRcMode) {
                case VENC_RC_MODE_MJPEGCBR:
                    sparams->rc             = RC_CBR;
                    sparams->fps            = stChnAttr.stRcAttr.stAttrMjpegeCbr.fr32DstFrmRate;

                    dparams->cbr.bitrate    = stChnAttr.stRcAttr.stAttrMjpegeCbr.u32BitRate;
                    dparams->cbr.gop        = PARAM_NO_VALUE;
                    dparams->cbr.stattime   = stChnAttr.stRcAttr.stAttrMjpegeCbr.u32StatTime;
                    dparams->cbr.fluctuate  = stChnAttr.stRcAttr.stAttrMjpegeCbr.u32FluctuateLevel;

                    break;
                case VENC_RC_MODE_MJPEGVBR:
                    sparams->rc             = RC_VBR;
                    sparams->fps            = stChnAttr.stRcAttr.stAttrMjpegeVbr.fr32DstFrmRate;

                    dparams->vbr.maxbitrate = stChnAttr.stRcAttr.stAttrMjpegeVbr.u32MaxBitRate;
                    dparams->vbr.gop        = PARAM_NO_VALUE;
                    dparams->vbr.maxqp      = stChnAttr.stRcAttr.stAttrMjpegeVbr.u32MaxQfactor;
                    dparams->vbr.minqp      = stChnAttr.stRcAttr.stAttrMjpegeVbr.u32MinQfactor;
                    dparams->vbr.stattime   = stChnAttr.stRcAttr.stAttrMjpegeVbr.u32StatTime;

                    break;
                case VENC_RC_MODE_MJPEGFIXQP:
                    sparams->rc             = RC_FIXQP;
                    sparams->fps            = stChnAttr.stRcAttr.stAttrMjpegeFixQp.fr32DstFrmRate;

                    dparams->fixqp.gop      = PARAM_NO_VALUE;
                    dparams->fixqp.iqp      = stChnAttr.stRcAttr.stAttrMjpegeFixQp.u32Qfactor;
                    dparams->fixqp.pqp      = PARAM_NO_VALUE;
                    //dparams->fixqp.stattime = PARAM_NO_VALUE;

                    break;
                default:
                    return ERR_INTERNAL;
            }
            sparams->width      = stChnAttr.stVeAttr.stAttrMjpege.u32PicWidth;
            sparams->height     = stChnAttr.stVeAttr.stAttrMjpege.u32PicHeight;
            sparams->profile    = MJPEG_PROFILE_BASELINE;
            break;
        case PT_H264:
            sparams->codec = CODEC_H264;
            switch(stChnAttr.stRcAttr.enRcMode) {
                case VENC_RC_MODE_H264CBR:
                    sparams->rc             = RC_CBR;
                    sparams->fps            = stChnAttr.stRcAttr.stAttrH264Cbr.fr32DstFrmRate;

                    dparams->cbr.bitrate    = stChnAttr.stRcAttr.stAttrH264Cbr.u32BitRate;
                    dparams->cbr.gop        = stChnAttr.stRcAttr.stAttrH264Cbr.u32Gop;
                    dparams->cbr.stattime   = stChnAttr.stRcAttr.stAttrH264Cbr.u32StatTime;
                    dparams->cbr.fluctuate  = stChnAttr.stRcAttr.stAttrH264Cbr.u32FluctuateLevel;

                    break;
                case VENC_RC_MODE_H264VBR:
                    sparams->rc             = RC_VBR;
                    sparams->fps            = stChnAttr.stRcAttr.stAttrH264Vbr.fr32DstFrmRate;

                    dparams->vbr.maxbitrate = stChnAttr.stRcAttr.stAttrH264Vbr.u32MaxBitRate;
                    dparams->vbr.gop        = stChnAttr.stRcAttr.stAttrH264Vbr.u32Gop;
                    dparams->vbr.maxqp      = stChnAttr.stRcAttr.stAttrH264Vbr.u32MaxQp;
                    dparams->vbr.minqp      = stChnAttr.stRcAttr.stAttrH264Vbr.u32MinQp;
                    dparams->vbr.stattime   = stChnAttr.stRcAttr.stAttrH264Vbr.u32StatTime;

                    break;
                case VENC_RC_MODE_H264FIXQP:
                    sparams->rc             = RC_FIXQP;
                    sparams->fps            = stChnAttr.stRcAttr.stAttrH264FixQp.fr32DstFrmRate;

                    dparams->fixqp.gop      = stChnAttr.stRcAttr.stAttrH264FixQp.u32Gop;
                    dparams->fixqp.iqp      = stChnAttr.stRcAttr.stAttrH264FixQp.u32IQp;
                    dparams->fixqp.pqp      = stChnAttr.stRcAttr.stAttrH264FixQp.u32PQp;
                    //dparams->fixqp.stattime = PARAM_NO_VALUE;

                    break;
                case VENC_RC_MODE_H264AVBR:
                case VENC_RC_MODE_H264QVBR:
                case VENC_RC_MODE_H264QPMAP:
                default:
                    return ERR_INTERNAL;
            }
            sparams->width  = stChnAttr.stVeAttr.stAttrH264e.u32PicWidth;
            sparams->height = stChnAttr.stVeAttr.stAttrH264e.u32PicHeight;
            switch(stChnAttr.stVeAttr.stAttrH265e.u32Profile) { // 0: baseline; 1:MP; 2:HP;  3:svc_t 
                case 0:
                    sparams->profile = H264_PROFILE_BASELINE;
                    break;
                case 1:
                    sparams->profile = H264_PROFILE_MAIN;
                    break;
                case 2:
                    sparams->profile = H264_PROFILE_HIGH;
                    break;
                case 3:
                default:
                    return ERR_INTERNAL;
            }
            break;
        case PT_H265:
            sparams->codec = CODEC_H265;
            switch(stChnAttr.stRcAttr.enRcMode) {
                case VENC_RC_MODE_H265CBR:
                    sparams->rc             = RC_CBR;
                    sparams->fps            = stChnAttr.stRcAttr.stAttrH265Cbr.fr32DstFrmRate;

                    dparams->cbr.bitrate    = stChnAttr.stRcAttr.stAttrH265Cbr.u32BitRate;
                    dparams->cbr.gop        = stChnAttr.stRcAttr.stAttrH265Cbr.u32Gop;
                    dparams->cbr.stattime   = stChnAttr.stRcAttr.stAttrH265Cbr.u32StatTime;
                    dparams->cbr.fluctuate  = stChnAttr.stRcAttr.stAttrH265Cbr.u32FluctuateLevel;

                    break;
                case VENC_RC_MODE_H265VBR:
                    sparams->rc             = RC_VBR;
                    sparams->fps            = stChnAttr.stRcAttr.stAttrH265Vbr.fr32DstFrmRate;

                    dparams->vbr.maxbitrate = stChnAttr.stRcAttr.stAttrH265Vbr.u32MaxBitRate;
                    dparams->vbr.gop        = stChnAttr.stRcAttr.stAttrH265Vbr.u32Gop;
                    dparams->vbr.maxqp      = stChnAttr.stRcAttr.stAttrH265Vbr.u32MaxQp;
                    dparams->vbr.minqp      = stChnAttr.stRcAttr.stAttrH265Vbr.u32MinQp;
                    dparams->vbr.stattime   = stChnAttr.stRcAttr.stAttrH265Vbr.u32StatTime;

                    break;
                case VENC_RC_MODE_H265FIXQP:
                    sparams->rc             = RC_FIXQP;
                    sparams->fps            = stChnAttr.stRcAttr.stAttrH265FixQp.fr32DstFrmRate;

                    dparams->fixqp.gop      = stChnAttr.stRcAttr.stAttrH265FixQp.u32Gop;
                    dparams->fixqp.iqp      = stChnAttr.stRcAttr.stAttrH265FixQp.u32IQp;
                    dparams->fixqp.pqp      = stChnAttr.stRcAttr.stAttrH265FixQp.u32PQp;
                    //dparams->fixqp.stattime = PARAM_NO_VALUE;

                    break;
                case VENC_RC_MODE_H265AVBR:
                case VENC_RC_MODE_H265QVBR:
                case VENC_RC_MODE_H265QPMAP:
                default:
                    return ERR_INTERNAL;
            }
            sparams->width  = stChnAttr.stVeAttr.stAttrH265e.u32PicWidth;
            sparams->height = stChnAttr.stVeAttr.stAttrH265e.u32PicHeight;
            switch(stChnAttr.stVeAttr.stAttrH265e.u32Profile) { //  0:Main
                case 0:
                    sparams->profile = H265_PROFILE_MAIN;
                    break;
                default:
                    return ERR_INTERNAL;
            }
            break;
        default:
            return ERR_INTERNAL;
    }

    //sparams->channel_id =; //TODO Get bind info

    return ERR_NONE;
}

int hisi_encoder_create (unsigned int id,
                        struct encoder_static_params * sparams,
                        struct encoder_dynamic_params * dparams) {
    int error_code = 0;

    if (id < 0 &&
        id >= VENC_MAX_CHN_NUM) return ERR_OBJECT_NOT_FOUNT;

    VENC_CHN_ATTR_S stChnAttr;

    error_code = HI_MPI_VENC_GetChnAttr(id, &stChnAttr); //TODO Check logic, maybe incorrect
    if (error_code == HI_SUCCESS) {
        return ERR_NOT_ALLOWED;
    } else {
        if (error_code != HI_ERR_VENC_UNEXIST) {
            return ERR_MPP;
        }
    }

    //Check input params...

    //fill stChnAttr with vales
    memset(&stChnAttr, 0, sizeof(stChnAttr));


    return ERR_NOT_IMPLEMENTED;
}

int hisi_encoder_update (unsigned int id,
                        struct encoder_dynamic_params * dparams) {
    int error_code = 0;

    if (id < 0 &&
        id >= VENC_MAX_CHN_NUM) return ERR_OBJECT_NOT_FOUNT;

    VENC_CHN_ATTR_S stChnAttr;

    error_code = HI_MPI_VENC_GetChnAttr(id, &stChnAttr);
    if (error_code != HI_SUCCESS) {
        if (error_code == HI_ERR_VENC_UNEXIST) {
            return ERR_OBJECT_NOT_FOUNT;
        }
        return ERR_MPP;
    }

    //TODO get encoder static params, check codec & ratecontrol
    //remember that if value == 0 (NO_CHANGE) than we using old value

    /*
    switch (ratecontrol) {
        case RC_CBR:
            if (dparams->cbr.bitrate < 0 &&
                dparams->cbr.bitrate > 1024*10) { //TODO max limit should be adjusted
                    dparams->cbr.bitrate = PARAM_BAD_VALUE;
                    error_code++;
            }
            if (dparams->cbr.gop < 0 &&
                dparams->cbr.gop > 100) { //TODO min & max limits should be adjusted
                    dparams->cbr.gop = PARAM_BAD_VALUE;
                    error_code++;
            }
            if (dparams->cbr.stattime < 0 &&
                dparams->cbr.stattime > 100) { //TODO min & max limits should be adjusted
                    dparams->cbr.stattime = PARAM_BAD_VALUE;
                    error_code++;
            }
            if (dparams->cbr.fluctuate < 0 &&
                dparams->cbr.fluctuate > 5) { //TODO min & max limits should be adjusted
                    dparams->cbr.fluctuate = PARAM_BAD_VALUE;
                    error_code++;
            }
            break;
        case RC_VBR:

            break;
        case RC_FIXQP:

            break;
        case RC_AVBR:
        case RC_QVBR:
        case RC_QMAP:
        default:
            return ERR_INTERNAL;
            break;
    }
    */
    if (error_code > 0) return ERR_BAD_PARAMS;

    //TODO update venc channel, fill new data

    error_code = HI_MPI_VENC_SetChnAttr(id, &stChnAttr);
    if (error_code != HI_SUCCESS) {
        return ERR_MPP;
    }

    return ERR_NONE;
}

int hisi_encoder_delete (unsigned int id) {
    int error_code = 0;

    if (id < 0 &&
        id >= VENC_MAX_CHN_NUM) return ERR_OBJECT_NOT_FOUNT;

    //TODO UnBind

    error_code = HI_MPI_VENC_StopRecvPic(id);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: hisi_encoder_delete: HI_MPI_VENC_StopRecvPic %d failed %#x\n", id, error_code);
        return ERR_MPP;
    }

    error_code = HI_MPI_VENC_DestroyChn(id);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: hisi_encoder_delete: HI_MPI_VENC_DestroyChn %d failed %#x\n", id, error_code);
        return -1;
    }

    return ERR_NONE;
}


