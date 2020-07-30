#include "venc.h"

#include <string.h>

void invalidate_mpp_venc_create_encoder_in (mpp_venc_create_encoder_in *in) {
    memset(in, 0xff, sizeof(mpp_venc_create_encoder_in));
}   

