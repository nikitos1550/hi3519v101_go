#include "himpp3_external.h"
#include "himpp3_mpp_includes.h"


#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <pthread.h>
#include <signal.h>
#include <sys/uio.h>
#include <stdint.h>
#include <sys/ioctl.h>
#include <fcntl.h>
#include <sys/stat.h>
#include <sys/syscall.h>
#include <sys/types.h>
#include <errno.h>
#include <ctype.h>
#include <sys/mman.h>
#include <sys/select.h>
#include <inttypes.h>


int himpp3_venc_info_chn(unsigned int chn, struct encoder * e) {
    int error_code;
    if (chn > (VENC_MAX_CHN_NUM-1)) return -1;

    VENC_CHN_ATTR_S stChnAttr;

    error_code = HI_MPI_VENC_GetChnAttr(chn, &stChnAttr);
    if (error_code != HI_SUCCESS) {
        if (error_code == HI_ERR_VENC_UNEXIST) {
            e->id = chn;
            e->enabled = 0;
            return 0;
        }
        return -1;
    }

    e->id = chn;
    e->enabled = 1;

    switch (stChnAttr.stVeAttr.enType) {
        case PT_JPEG:
            e->etype = enc_type_jpeg;
            break;
        case PT_MJPEG:
            e->etype = enc_type_mjpeg;
            break;
        case PT_H264:
            e->etype = enc_type_h264;
            break;
        case PT_H265:
            e->etype = enc_type_h265;
            break;
        default:
            return -1;
    }
    e->source = 0;

    return 0;
}
