//+build arm
//+build hi3516cv500

package sys

/*
#include "../include/mpp_v4.h"

#include <string.h>

#define ERR_NONE                0
#define ERR_HI_MPI_SYS_Exit     2
#define ERR_HI_MPI_VB_Exit      3
#define ERR_HI_MPI_VB_SetConfig 4
#define ERR_HI_MPI_VB_Init      5
#define ERR_HI_MPI_SYS_SetConf  6
#define ERR_HI_MPI_SYS_Init     7

int mpp4_sys_init(unsigned int *error_code) {
    *error_code = 0;

    VB_CONFIG_S        stVbConf;
    HI_U32             u32BlkSize;

     hi_memset(&stVbConf, sizeof(VB_CONFIG_S), 0, sizeof(VB_CONFIG_S));
    stVbConf.u32MaxPoolCnt              = 2;


    u32BlkSize = COMMON_GetPicBufferSize(   1920, 
                                            1080, 
                                            PIXEL_FORMAT_YVU_SEMIPLANAR_420, 
                                            DATA_BITWIDTH_8, 
                                            COMPRESS_MODE_SEG, 
                                            DEFAULT_ALIGN);
    stVbConf.astCommPool[0].u64BlkSize  = u32BlkSize;
    stVbConf.astCommPool[0].u32BlkCnt   = 10;

    
//    u32BlkSize = VI_GetRawBufferSize(       1920, 
//                                            1080, 
//                                            PIXEL_FORMAT_RGB_BAYER_16BPP, 
//                                            COMPRESS_MODE_NONE, 
//                                            DEFAULT_ALIGN);
//    stVbConf.astCommPool[1].u64BlkSize  = u32BlkSize;
//    stVbConf.astCommPool[1].u32BlkCnt   = 4;

        *error_code = HI_MPI_SYS_Exit();
	if (*error_code != HI_SUCCESS) return ERR_HI_MPI_SYS_Exit;

        *error_code = HI_MPI_VB_Exit();
	if (*error_code != HI_SUCCESS) return ERR_HI_MPI_VB_Exit;

        *error_code = HI_MPI_VB_SetConfig(&stVbConf);
	if (*error_code != HI_SUCCESS) return ERR_HI_MPI_VB_SetConfig;

        *error_code = HI_MPI_VB_Init();
	if (*error_code != HI_SUCCESS) return ERR_HI_MPI_VB_Init;
	    
//        stSysConf.u32AlignWidth = 64;
//
//        ret = HI_MPI_SYS_SetConf(&stSysConf);
//        if (ret != HI_SUCCESS) {
//                fprintf(stderr, "HI_MPI_SYS_SetConf failed: ");
//                //resolve_mppv2_errors(ret);
//                return 1;
//        }

        *error_code = HI_MPI_SYS_Init();
	if (*error_code != HI_SUCCESS) return ERR_HI_MPI_SYS_Init;

    return ERR_NONE;
}
*/
import "C"

import (
	"application/pkg/mpp/error"
	"application/pkg/logger"
)
func Init() {
        var errorCode C.uint

        switch err := C.mpp4_sys_init(&errorCode); err {
        case C.ERR_NONE:
                logger.Log.Debug().
                        Msg("C.mpp4_sys_init ok")
        case C.ERR_HI_MPI_SYS_Exit:
                logger.Log.Fatal().
                        Str("func", "HI_MPI_SYS_Exit()").
                        Int("error", int(errorCode)).
                        Str("error_desc", error.Resolve(int64(errorCode))).
                        Msg("C.mpp4_sys_init() error")
        case C.ERR_HI_MPI_VB_Exit:
                logger.Log.Fatal().
                        Str("func", "HI_MPI_VB_Exit()").
                        Int("error", int(errorCode)).
                        Str("error_desc", error.Resolve(int64(errorCode))).
                        Msg("C.mpp4_sys_init() error")
        case C.ERR_HI_MPI_VB_SetConfig:
                logger.Log.Fatal().
                        Str("func", "HI_MPI_VB_SetConfig()").
                        Int("error", int(errorCode)).
                        Str("error_desc", error.Resolve(int64(errorCode))).
                        Msg("C.mpp4_sys_init() error")
        case C.ERR_HI_MPI_VB_Init:
                logger.Log.Fatal().
                        Str("func", "HI_MPI_VB_Init()").
                        Int("error", int(errorCode)).
                        Str("error_desc", error.Resolve(int64(errorCode))).
                        Msg("C.mpp4_sys_init() error")
        case C.ERR_HI_MPI_SYS_SetConf:
                logger.Log.Fatal().
                        Str("func", "HI_MPI_SYS_SetConf()").
                        Int("error", int(errorCode)).
                        Str("error_desc", error.Resolve(int64(errorCode))).
                        Msg("C.mpp4_sys_init() error")
        case C.ERR_HI_MPI_SYS_Init:
                logger.Log.Fatal().
                        Str("func", "HI_MPI_SYS_Init()").
                        Int("error", int(errorCode)).
                        Str("error_desc", error.Resolve(int64(errorCode))).
                        Msg("C.mpp4_sys_init() error")
        default:
                logger.Log.Fatal().
                        Int("error", int(err)).
                        Msg("C.mpp4_sys_init() Unexpected return")
        }
}

