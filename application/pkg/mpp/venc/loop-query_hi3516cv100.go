//+build arm
//+build hi3516cv100

package venc

/*
#include "../include/mpp_v1.h"

#include <string.h>
#include <stdlib.h>

//#include "getfunc.h"
#include "loop.h"

int mpp_venc_getfd(int venc_channel_id) {
    return HI_MPI_VENC_GetFd(venc_channel_id);
}

void mpp_data_loop_get_data(unsigned int venc_channel) {
}
*/
import "C"
