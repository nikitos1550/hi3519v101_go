//+build arm
//+build hi3516cv300

package sys

/*
#include "../include/mpp_v3.h"

#include <string.h>

#define ERR_NONE                0
#define ERR_HI_MPI_SYS_Exit     2
#define ERR_HI_MPI_VB_Exit      3
#define ERR_HI_MPI_VB_SetConfig 4
#define ERR_HI_MPI_VB_Init      5
#define ERR_HI_MPI_SYS_SetConf  6
#define ERR_HI_MPI_SYS_Init     7

int mpp3_sys_init(unsigned int *error_code) {
    *error_code = 0;


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

        switch err := C.mpp3_sys_init(&errorCode); err {
        case C.ERR_NONE:
                logger.Log.Debug().
                        Msg("C.mpp3_sys_init ok")
        case C.ERR_HI_MPI_SYS_Exit:
                logger.Log.Fatal().
                        Str("func", "HI_MPI_SYS_Exit()").
                        Int("error", int(errorCode)).
                        Str("error_desc", error.Resolve(int64(errorCode))).
                        Msg("C.mpp3_sys_init() error")
        case C.ERR_HI_MPI_VB_Exit:
                logger.Log.Fatal().
                        Str("func", "HI_MPI_VB_Exit()").
                        Int("error", int(errorCode)).
                        Str("error_desc", error.Resolve(int64(errorCode))).
                        Msg("C.mpp3_sys_init() error")
        case C.ERR_HI_MPI_VB_SetConfig:
                logger.Log.Fatal().
                        Str("func", "HI_MPI_VB_SetConfig()").
                        Int("error", int(errorCode)).
                        Str("error_desc", error.Resolve(int64(errorCode))).
                        Msg("C.mpp3_sys_init() error")
        case C.ERR_HI_MPI_VB_Init:
                logger.Log.Fatal().
                        Str("func", "HI_MPI_VB_Init()").
                        Int("error", int(errorCode)).
                        Str("error_desc", error.Resolve(int64(errorCode))).
                        Msg("C.mpp3_sys_init() error")
        case C.ERR_HI_MPI_SYS_SetConf:
                logger.Log.Fatal().
                        Str("func", "HI_MPI_SYS_SetConf()").
                        Int("error", int(errorCode)).
                        Str("error_desc", error.Resolve(int64(errorCode))).
                        Msg("C.mpp3_sys_init() error")
        case C.ERR_HI_MPI_SYS_Init:
                logger.Log.Fatal().
                        Str("func", "HI_MPI_SYS_Init()").
                        Int("error", int(errorCode)).
                        Str("error_desc", error.Resolve(int64(errorCode))).
                        Msg("C.mpp3_sys_init() error")
        default:
                logger.Log.Fatal().
                        Int("error", int(err)).
                        Msg("C.mpp3_sys_init() Unexpected return")
        }
}

