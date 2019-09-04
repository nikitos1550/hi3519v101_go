#ifndef HI3516AV200_ENCODERS_H_
#define HI3516AV200_ENCODERS_H_

#include "hi3516av200_mpp.h"

#define ENCODER_DISABLED    0
#define ENCODER_ENABLED     1

extern int encoders_enable[VENC_MAX_CHN_NUM];

//setup encoder, bind it to enabled channel and start recv images
int hi3516av200_encoder_create(unsigned int id, unsigned int channel, VENC_CHN_ATTR_S * mpp_venc) ;

//get mpp struct from MPP or returns not found or mpp error
int hi3516av200_encoder_fetch(unsigned int id, VENC_CHN_ATTR_S * mpp_venc);

//fill sparams struct <-from- mpp venc struct
int hi3516av200_encoders_get_sparams(VENC_CHN_ATTR_S * mpp_venc, struct encoder_static_params * sparams);

//fill mpp venc struct <-from- sparams struct
int hi3516av200_encoders_set_sparams(struct encoder_static_params * sparams, VENC_CHN_ATTR_S * mpp_venc);

//validate sparams structure
int hi3516av200_encoders_validate_sparams(struct encoder_static_params * sparams);

#endif //HI3516AV200_ENCODERS_H_
