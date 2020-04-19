//+build arm
//+build hi3516av100

package sys

/*
#include "../include/mpp.h"
#include "../../logger/logger.h"

#include <stdint.h>
#include <string.h>

#define ERR_NONE                0
#define ERR_MPP                 1

typedef struct hi3516av100_sys_init_in_struct {
    unsigned int width;
    unsigned int height;
    unsigned int cnt;
} hi3516av100_sys_init_in;

static int hi3516av100_sys_init(int64_t *error_code, hi3516av100_sys_init_in *in) {
    *error_code = 0;

    *error_code = HI_MPI_SYS_Exit();
    if (*error_code != HI_SUCCESS) {
        GO_LOG_SYS(LOGGER_ERROR, "HI_MPI_SYS_Exit")
        return ERR_MPP;
    }

    *error_code = HI_MPI_VB_Exit();
    if (*error_code != HI_SUCCESS) {
        GO_LOG_SYS(LOGGER_ERROR, "HI_MPI_VB_Exit")
        return ERR_MPP;
    }

    VB_CONF_S stVbConf;

    memset(&stVbConf, 0, sizeof(VB_CONF_S));
    stVbConf.u32MaxPoolCnt                  = 128;
    stVbConf.astCommPool[0].u32BlkSize      = (CEILING_2_POWER(in->width, 64) * CEILING_2_POWER(in->height, 64) * 1.5);
    stVbConf.astCommPool[0].u32BlkCnt       = in->cnt;

    *error_code = HI_MPI_VB_SetConf(&stVbConf);
    if(*error_code != HI_SUCCESS) {
        GO_LOG_SYS(LOGGER_ERROR, "HI_MPI_VB_SetConf") 
        return ERR_MPP;
    }

    *error_code = HI_MPI_VB_Init();
    if (*error_code != HI_SUCCESS) {
        GO_LOG_SYS(LOGGER_ERROR, "HI_MPI_VB_Init")
        return ERR_MPP;
    }

    MPP_SYS_CONF_S stSysConf;

    memset(&stSysConf, 0, sizeof(MPP_SYS_CONF_S));
    stSysConf.u32AlignWidth = 64;

    *error_code = HI_MPI_SYS_SetConf(&stSysConf);
    if (*error_code != HI_SUCCESS) {
        GO_LOG_SYS(LOGGER_ERROR, "HI_MPI_SYS_SetConf")
        return ERR_MPP;
    }

    *error_code = HI_MPI_SYS_Init();
    if(*error_code != HI_SUCCESS) {
        GO_LOG_SYS(LOGGER_ERROR, "HI_MPI_SYS_Init")
        return ERR_MPP;
    }

    return ERR_NONE;
}
*/
import "C"

import (
	"application/pkg/mpp/error"
	"application/pkg/logger"
)

func initFamily() error {
    var errorCode C.int64_t
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
        return errmpp.New("funcname", int64(errorCode))
    }

    return nil
}
