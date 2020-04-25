//+build arm
//+build hi3516av100

package sys

/*
#include "../include/mpp.h"
#include "../errmpp/errmpp.h"

#include <stdint.h>
#include <string.h>

typedef struct hi3516av100_sys_init_in_struct {
    unsigned int width;
    unsigned int height;
    unsigned int cnt;
} hi3516av100_sys_init_in;

static int hi3516av100_sys_init(error_in *err, hi3516av100_sys_init_in *in) {
    unsigned int mpp_error_code = 0;

    mpp_error_code = HI_MPI_SYS_Exit();
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_SYS_Exit, mpp_error_code);
    }

    mpp_error_code = HI_MPI_VB_Exit();
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_VB_Exit, mpp_error_code);
    }

    VB_CONF_S stVbConf;

    memset(&stVbConf, 0, sizeof(VB_CONF_S));
    stVbConf.u32MaxPoolCnt                  = 128;
    stVbConf.astCommPool[0].u32BlkSize      = (CEILING_2_POWER(in->width, 64) * CEILING_2_POWER(in->height, 64) * 1.5);
    stVbConf.astCommPool[0].u32BlkCnt       = in->cnt;

    mpp_error_code = HI_MPI_VB_SetConf(&stVbConf);
    if(mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_VB_SetConf, mpp_error_code);
    }

    mpp_error_code = HI_MPI_VB_Init();
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_VB_Init, mpp_error_code);
    }

    MPP_SYS_CONF_S stSysConf;

    memset(&stSysConf, 0, sizeof(MPP_SYS_CONF_S));
    stSysConf.u32AlignWidth = 64;

    mpp_error_code = HI_MPI_SYS_SetConf(&stSysConf);
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_SYS_SetConf, mpp_error_code);
    }

    mpp_error_code = HI_MPI_SYS_Init();
    if(mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_SYS_Init, mpp_error_code);
    }

    return ERR_NONE;
}
*/
import "C"

import (
	"application/pkg/mpp/errmpp"
	"application/pkg/logger"
)

func initFamily() error {
    var inErr C.error_in
    var in C.hi3516av100_sys_init_int

    in.width = C.uint(width)
    in.height = C.uint(height)
    in.cnt = C.uint(cnt)

    logger.Log.Trace().
        Uint("width", uint(in.width)).
        Uint("height", uint(in.height)).
        Uint("cnt", uint(in.cnt)).
        Msg("SYS params")

    err := C.hi3516av100_sys_init(&errorCode, &in)
    if err != C.ERR_NONE {
        return errmpp.New(uint(inErr.f), uint(inErr.mpp))
    }

    return nil
}
