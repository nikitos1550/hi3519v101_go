//+build arm
//+build hi3516cv500

package sys

/*
#include "../include/mpp.h"
#include "../errmpp/errmpp.h"

#include <string.h>

typedef struct hi3516cv500_sys_init_in_struct {
    unsigned int width; 
    unsigned int height;
    unsigned int cnt;
} hi3516cv500_sys_init_in;

static int hi3516cv500_sys_init(error_in *err, hi3516cv500_sys_init_in *in) {
    unsigned int mpp_error_code = 0;

    mpp_error_code = HI_MPI_SYS_Exit();
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_, mpp_error_code);
    }

    mpp_error_code = HI_MPI_VB_Exit();
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_, mpp_error_code);
    }


    VB_CONFIG_S        stVbConf;

    memset(&stVbConf,0,sizeof(VB_CONFIG_S));
    stVbConf.u32MaxPoolCnt              = 2;
    stVbConf.astCommPool[0].u64BlkSize = COMMON_GetPicBufferSize(   in->width, 
                                                                    in->height, 
                                                                    PIXEL_FORMAT_YVU_SEMIPLANAR_420, 
                                                                    DATA_BITWIDTH_8, 
                                                                    COMPRESS_MODE_SEG, 
                                                                    DEFAULT_ALIGN);
    stVbConf.astCommPool[0].u32BlkCnt = in->cnt;

    mpp_error_code = HI_MPI_VB_SetConfig(&stVbConf);
	if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_, mpp_error_code);
    }

    mpp_error_code = HI_MPI_VB_Init();
	if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_, mpp_error_code);
    }

    MPP_SYS_CONF_S stSysConf;

    memset(&stSysConf, 0, sizeof(MPP_SYS_CONF_S));
    stSysConf.u32AlignWidth = 64;

    mpp_error_code = HI_MPI_SYS_SetConf(&stSysConf);
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_, mpp_error_code);
    }

    mpp_error_code = HI_MPI_SYS_Init();
	if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_, mpp_error_code);
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
        return errmpp.New(uint(inErr.f), uint(inErr.mpp))
    }

    return nil
}

