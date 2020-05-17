//+build nobuild

//+build arm
//+build hi3516ev200

package venc

/*
#include "../include/mpp_v4.h"

#include <string.h>
#include <stdlib.h>

#include "loop.h"

int mpp_venc_getfd(int venc_channel_id) {
    return HI_MPI_VENC_GetFd(venc_channel_id);
}

void mpp_data_loop_get_data(unsigned int venc_channel) {
}
*/
import "C"
