//+build arm
//+build hi3516av100

package sys

/*
#include "../include/mpp_v2.h"

#include <string.h>

#define ERR_NONE                0
#define ERR_HI_MPI_SYS_Exit     2
#define ERR_HI_MPI_VB_Exit      3
#define ERR_HI_MPI_VB_SetConf   4
#define ERR_HI_MPI_VB_Init      5
#define ERR_HI_MPI_SYS_SetConf  6
#define ERR_HI_MPI_SYS_Init     7

int mpp2_sys_init(unsigned int *error_code) {
    *error_code = 0;

    *error_code = HI_MPI_SYS_Exit();
    if (*error_code != HI_SUCCESS) return ERR_HI_MPI_SYS_Exit;

    *error_code = HI_MPI_VB_Exit();
    if (*error_code != HI_SUCCESS) return ERR_HI_MPI_VB_Exit;

    VB_CONF_S stVbConf;

    memset(&stVbConf, 0, sizeof(VB_CONF_S));
    stVbConf.u32MaxPoolCnt                  = 128;
    stVbConf.astCommPool[0].u32BlkSize      = (CEILING_2_POWER(2592, 64) * CEILING_2_POWER(1944, 64) * 1.5);
    stVbConf.astCommPool[0].u32BlkCnt       = 10;

    *error_code = HI_MPI_VB_SetConf(&stVbConf);
    if(*error_code != HI_SUCCESS) return ERR_HI_MPI_VB_SetConf;

    *error_code = HI_MPI_VB_Init();
    if (*error_code != HI_SUCCESS) return ERR_HI_MPI_VB_Init;

    MPP_SYS_CONF_S stSysConf;

    stSysConf.u32AlignWidth = 64;

    *error_code = HI_MPI_SYS_SetConf(&stSysConf);
    if (*error_code != HI_SUCCESS) return ERR_HI_MPI_SYS_SetConf;

    *error_code = HI_MPI_SYS_Init();
    if(*error_code != HI_SUCCESS) return ERR_HI_MPI_SYS_Init;

    return ERR_NONE;
}
*/
import "C"

import (
	"application/pkg/mpp/error"
	//"log"

	"application/pkg/logger"
)

func Init() {
	var errorCode C.uint

	switch err := C.mpp2_sys_init(&errorCode); err {
	case C.ERR_NONE:
		logger.Log.Debug().
			Msg("C.mpp2_sys_init ok")
	case C.ERR_HI_MPI_SYS_Exit:
		logger.Log.Fatal().
			Str("func", "HI_MPI_SYS_Exit()").
			Int("error", int(errorCode)).
			Str("error_desc", error.Resolve(int64(errorCode))).
			Msg("C.mpp2_sys_init() error")
	case C.ERR_HI_MPI_VB_Exit:
		logger.Log.Fatal().
                        Str("func", "HI_MPI_VB_Exit()").
                        Int("error", int(errorCode)).
                        Str("error_desc", error.Resolve(int64(errorCode))).
                        Msg("C.mpp2_sys_init() error")
	case C.ERR_HI_MPI_VB_SetConf:
		logger.Log.Fatal().
                        Str("func", "HI_MPI_VB_SetConf()").
                        Int("error", int(errorCode)).
                        Str("error_desc", error.Resolve(int64(errorCode))).
                        Msg("C.mpp2_sys_init() error")
	case C.ERR_HI_MPI_VB_Init:
		logger.Log.Fatal().
                        Str("func", "HI_MPI_VB_Init()").
                        Int("error", int(errorCode)).
                        Str("error_desc", error.Resolve(int64(errorCode))).
                        Msg("C.mpp2_sys_init() error")
	case C.ERR_HI_MPI_SYS_SetConf:
		logger.Log.Fatal().
                        Str("func", "HI_MPI_SYS_SetConf()").
                        Int("error", int(errorCode)).
                        Str("error_desc", error.Resolve(int64(errorCode))).
                        Msg("C.mpp2_sys_init() error")
	case C.ERR_HI_MPI_SYS_Init:
		logger.Log.Fatal().
                        Str("func", "HI_MPI_SYS_Init()").
                        Int("error", int(errorCode)).
                        Str("error_desc", error.Resolve(int64(errorCode))).
                        Msg("C.mpp2_sys_init() error")
	default:
		logger.Log.Fatal().
			Int("error", int(err)).
			Msg("C.mpp2_sys_init() Unexpected return")
	}
}
