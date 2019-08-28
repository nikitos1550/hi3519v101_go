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


/*
struct channel {
    unsigned int    id;
    unsigned char   enabled;
    unsigned int    width;
    unsigned int    height;
    unsigned int    fps;

}
*/
int himpp3_vpss_info_chn(unsigned int chn, struct channel * c) {
    int error_code;

    if (chn >= VPSS_MAX_PHY_CHN_NUM) {
        return -1;
    }

    c->id = chn;

    VPSS_CHN_ATTR_S stVpssChnAttr;
    VPSS_CHN_MODE_S stVpssChnMode;

    error_code = HI_MPI_VPSS_GetChnAttr(0, chn, &stVpssChnAttr);
    if (error_code != HI_SUCCESS) {

        if (error_code == HI_ERR_VPSS_UNEXIST) {
            c->enabled = 0;
            return 0;
        }

        printf("C DEBUG: HI_MPI_VPSS_GetChnAttr failed with %#x\n", error_code);
        return -1;
    }

    error_code = HI_MPI_VPSS_GetChnMode(0, chn, &stVpssChnMode);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: HI_MPI_VPSS_GetChnMode failed with %#x\n", error_code);
        return -1;
    }

    c->enabled  = 1;
    c->width    = stVpssChnMode.u32Width;
    c->height   = stVpssChnMode.u32Height;
    c->fps      = stVpssChnAttr.s32DstFrameRate;

    return 0;
}
