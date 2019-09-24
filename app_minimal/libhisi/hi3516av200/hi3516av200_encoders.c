#include "../hisi_external.h"
#include "hi3516av200_mpp.h"
#include "hi3516av200_encoders.h"

#include "hi3516av200_cmos.h"

#include <string.h>

int encoders_enable[VENC_MAX_CHN_NUM];

int hisi_encoders_max_num(unsigned int * num) {
    *num = VENC_MAX_CHN_NUM;
    return ERR_NONE;
}


int hisi_encoder_fetch  (unsigned int id,
                         struct encoder_static_params * sparams) {

    if (id > VENC_MAX_CHN_NUM) return ERR_OBJECT_NOT_FOUND;

    if (encoders_enable[id] != ENCODER_ENABLED) return ERR_OBJECT_NOT_FOUND;

    int error_code = 0;

    VENC_CHN_ATTR_S mpp_venc;

    error_code = hi3516av200_encoder_fetch(id, &mpp_venc);
    switch(error_code) {
        case ERR_NONE:
            break;
        case ERR_OBJECT_NOT_FOUND:
            return ERR_OBJECT_NOT_FOUND;
        case ERR_MPP:
            return ERR_MPP;
        default:
            return ERR_GENERAL;
    }

    hi3516av200_encoders_get_sparams(&mpp_venc, sparams);

    return ERR_NONE;
}

int hisi_encoder_delete (unsigned int id) {
    if (id > VENC_MAX_CHN_NUM) return ERR_OBJECT_NOT_FOUND;

    if (encoders_enable[id] != ENCODER_ENABLED) return ERR_OBJECT_NOT_FOUND;

    int error_code = 0;

    MPP_CHN_S stSrcChn;
    MPP_CHN_S stDestChn;

    stDestChn.enModId  = HI_ID_VENC;
    stDestChn.s32DevId = 0;
    stDestChn.s32ChnId = id;

    error_code =  HI_MPI_SYS_GetBindbyDest(&stDestChn, &stSrcChn);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: hisi_encoder_delete: HI_MPI_SYS_GetBindbyDest %d failed %#x\n", id, error_code);
        return ERR_MPP;
    }

    error_code = HI_MPI_SYS_UnBind(&stSrcChn, &stDestChn);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: hisi_encoder_delete: HI_MPI_SYS_UnBind %d failed %#x\n", id, error_code);
        return ERR_MPP;
    }

    error_code = HI_MPI_VENC_StopRecvPic(id);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: hisi_encoder_delete: HI_MPI_VENC_StopRecvPic %d failed %#x\n", id, error_code);
        return ERR_MPP;
    }

    error_code = HI_MPI_VENC_DestroyChn(id);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: hisi_encoder_delete: HI_MPI_VENC_DestroyChn %d failed %#x\n", id, error_code);
        return ERR_MPP;
    }

    return ERR_NONE;
}

int hi3516av200_encoder_create(unsigned int id, unsigned int channel, VENC_CHN_ATTR_S * mpp_venc) {

    //TODO

    return ERR_NONE;
}

//  \return error
//  \retval ERR_NONE
//  \retval ERR__MPP
//  \retval ERR_OBJECT_NOT_FOUND
int hi3516av200_encoder_fetch(unsigned int id, VENC_CHN_ATTR_S * mpp_venc) {

    if (id > VENC_MAX_CHN_NUM) return ERR_OBJECT_NOT_FOUND;

    if (encoders_enable[id] != ENCODER_ENABLED) return ERR_OBJECT_NOT_FOUND;

    int error_code = 0;

    error_code = HI_MPI_VENC_GetChnAttr(id, mpp_venc);
    switch(error_code) {
        case HI_SUCCESS:
            return ERR_NONE;
        case HI_ERR_VENC_UNEXIST:
            return ERR_OBJECT_NOT_FOUND;
        default:
            return ERR_MPP;
    }
}

int hi3516av200_encoders_get_sparams(VENC_CHN_ATTR_S * mpp_venc, struct encoder_static_params * sparams) {
/*
    int codec;
    int rc;
    int profile;
    int width;
    int height;
    int fps;
    int channel;
*/
    return ERR_NONE;
}

int hi3516av200_encoders_set_sparams(struct encoder_static_params * sparams, VENC_CHN_ATTR_S * mpp_venc) {
/*
    int codec;
    int rc;
    int profile;
    int width;
    int height;
    int fps;
    int channel;
*/
    return ERR_NONE;
}

int hi3516av200_encoders_validate_sparams(struct encoder_static_params * sparams) {
    int error_num = 0;

    //check codec
    //check rc
    //check profile
    //These should be checked by upper level functions, as rc and profile depends on codec

    struct channel_params chn;
    int error_code = hisi_channel_fetch(sparams->channel, &chn);
    if (error_code != ERR_NONE) {
        error_num++;
        sparams->channel = PARAM_BAD_VALUE;
    }

    if (sparams->width < 0 ||
        sparams->width > chn.width ||
        sparams->width & 1 != 0) { // % 2 != 0
        error_num++;
        sparams->width = PARAM_BAD_VALUE;
    }

    if (sparams->height < 0 ||
        sparams->height > chn.height ||
        sparams->height & 1 != 0) { // % 2 != 0
        error_num++;
        sparams->width = PARAM_BAD_VALUE;
    }

    if (sparams->fps < 1 ||
        sparams->fps > chn.fps) {
        error_num++;
        sparams->width = PARAM_BAD_VALUE;
    }

    if (error_num > 0) return ERR_BAD_PARAMS;

    return ERR_NONE;
}


