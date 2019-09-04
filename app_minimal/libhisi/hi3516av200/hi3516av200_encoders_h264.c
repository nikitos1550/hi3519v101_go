#include "hi3516av200_mpp.h"
#include "../hisi_external.h"

#include "hi3516av200_encoders.h"

//

int hisi_hi3516av200_encoder_h264_cbr_fetch      (unsigned int id,
                                             struct encoder_static_params * sparams,
                                             struct encoder_h26x_cbr_params * dparams) {
    int error_code = 0;

    //get encoder struct via mpp
    //check that encoder exist
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

    //check that encoder is h264 && cbr
    if (sparams->codec != CODEC_H264) {
        return ERR_OBJECT_NOT_FOUND;//maybe add OBJECT_ANOTHER_TYPE
    }

    if (sparams->rc != RC_CBR) {
        return ERR_OBJECT_NOT_FOUND;//maybe add OBJECT_ANOTHER_TYPE
    }

    //fill sparams struct from mpp
    hi3516av200_encoders_get_sparams(&mpp_venc, sparams);

    //fill dparams struct from mpp
    //TODO

    return ERR_NONE;
}

int hisi_hi3516av200_encoder_h264_cbr_create     (unsigned int id,
                                             struct encoder_static_params * sparams,
                                             struct encoder_h26x_cbr_params * dparams) {
    //assumed that channel will be blocked in go level, to prevent creation if channel for disabled channel

    //get encoder struct via mpp
    //check that encoder unexist
    int error_code = 0;

    VENC_CHN_ATTR_S mpp_venc;

    error_code = hi3516av200_encoder_fetch(id, &mpp_venc);
    if (error_code == ERR_NONE) {
        return ERR_NOT_ALLOWED;// we can`t create encoder in busy slot
    } else { //ERR_MPP || ERR_OBJECT_NOT_FOUND
        if (error_code != ERR_OBJECT_NOT_FOUND) {
            return error_code;
        }
    }

    //validate sparams ---> hi3516av200_encoders_validate_sparams
    error_code = hi3516av200_encoders_validate_sparams(sparams);
    if (error_code != ERR_NONE) {
        return error_code;//ERR_BAD_PARAMS
    }

    //validate dparams
        //params range >0 <something
    //TODO

    //check channel -> no check, now it is upon gospace managment

    //set sparams ---> hi3516av200_encoders_set_sparams
    hi3516av200_encoders_set_sparams(sparams, &mpp_venc);

    //set dparams

    //create encoder via mpp
    error_code = hi3516av200_encoder_create(id, sparams->channel, &mpp_venc);
    if (error_code != ERR_NONE) {
        //???
    }
    return ERR_NONE;
}

int hisi_hi3516av200_encoder_h264_cbr_update     (unsigned int id,
                                             struct encoder_h26x_cbr_params * dparams) {
    //get encoder struct via mpp
    //check that encoder exist ---> hisi_encoder_fetch

    //fill sparams ---> hi3516av200_encoders_get_sparams
    //check that encoder is h264 && cbr

    //validate dparams
        //params range >=0 <something
    //set dparams, set only if param != 0

    //update encoder via mpp

}

int hisi_hi3516av200_encoder_h264_vbr_fetch      (unsigned int id,
                                             struct encoder_static_params * sparams,
                                             struct encoder_h26x_vbr_params * dparams) {

}

int hisi_hi3516av200_encoder_h264_vbr_create     (unsigned int id,
                                             struct encoder_static_params * sparams,
                                             struct encoder_h26x_vbr_params * dparams) {

}

int hisi_hi3516av200_encoder_h264_vbr_update     (unsigned int id,
                                             struct encoder_h26x_vbr_params * dparams) {

}

int hisi_hi3516av200_encoder_h264_fixqp_fetch    (unsigned int id,
                                             struct encoder_static_params * sparams,
                                             struct encoder_h26x_fixqp_params * dparams) {

}

int hisi_hi3516av200_encoder_h264_fixqp_create   (unsigned int id,
                                             struct encoder_static_params * sparams,
                                             struct encoder_h26x_fixqp_params * dparams) {

}

int hisi_hi3516av200_encoder_h264_fixqp_update   (unsigned int id,
                                             struct encoder_h26x_fixqp_params * dparams) {

}
