//+build arm
//+build hi3516cv500

package sys

/*
#include "../include/mpp_v4.h"

#include <string.h>

#define ERR_NONE                0
#define ERR_MPP                 1

typedef struct hi3516cv500_sys_init_in_struct {
    unsigned int width; 
    unsigned int height;
    unsigned int cnt;
} hi3516cv500_sys_init_in;

static int hi3516cv500_sys_init(unsigned int *error_code, hi3516cv500_sys_init_in *in) {
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


    VB_CONFIG_S        stVbConf;

    //hi_memset(&stVbConf, sizeof(VB_CONFIG_S), 0, sizeof(VB_CONFIG_S));   
    memset(&stVbConf,0,sizeof(VB_CONFIG_S));
    stVbConf.u32MaxPoolCnt              = 2;
    stVbConf.astCommPool[0].u64BlkSize = COMMON_GetPicBufferSize(   in->width, 
                                                                    in->height, 
                                                                    PIXEL_FORMAT_YVU_SEMIPLANAR_420, 
                                                                    DATA_BITWIDTH_8, 
                                                                    COMPRESS_MODE_SEG, 
                                                                    DEFAULT_ALIGN);
    stVbConf.astCommPool[0].u32BlkCnt = in->cnt;

    *error_code = HI_MPI_VB_SetConfig(&stVbConf);
	if (*error_code != HI_SUCCESS) {
        GO_LOG_SYS(LOGGER_ERROR, "HI_MPI_VB_SetConfig")
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
	if (*error_code != HI_SUCCESS) {
        GO_LOG_SYS(LOGGER_ERROR, "HI_MPI_SYS_Init")
        return ERR_MPP;
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
    var errorCode C.uint
    var in C.hi3516cv500_sys_init_int

    in.width = C.uint(width)
    in.height = C.uint(height)
    in.cnt = C.uint(cnt)

    logger.Log.Trace().
        Uint("width", uint(in.width)).
        Uint("height", uint(in.height)).
        Uint("cnt", uint(in.cnt)).
        Msg("SYS params")

    err := C.hi3516cv500_sys_init(&errorCode, &in)
    if err != C.ERR_NONE {
        return errmpp.New("funcname", int64(errorCode))
    }

    return nil
}

