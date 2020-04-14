//+build arm
//+build hi3516cv300

package sys

/*
#include "../include/mpp_v3.h"

#include <string.h>

#define ERR_NONE                0
#define ERR_HI_MPI_SYS_Exit     2
#define ERR_HI_MPI_VB_Exit      3
#define ERR_HI_MPI_VB_SetConf   4
#define ERR_HI_MPI_VB_Init      5
#define ERR_HI_MPI_SYS_SetConf  6
#define ERR_HI_MPI_SYS_Init     7

int mpp3_sys_init(unsigned int *error_code,
                    unsigned int width,
                    unsigned int height) {
    *error_code = 0;

        VB_CONF_S               stVbConf;
    MPP_SYS_CONF_S      stSysConf;


        memset(&stVbConf,0,sizeof(VB_CONF_S));
    stVbConf.u32MaxPoolCnt = 128;

    //video buffer
    stVbConf.astCommPool[0].u32BlkSize =(CEILING_2_POWER(width, 64) * CEILING_2_POWER(height, 64) * 1.5);
    stVbConf.astCommPool[0].u32BlkCnt = 10;

        *error_code = HI_MPI_SYS_Exit();
	if (*error_code != HI_SUCCESS) return ERR_HI_MPI_SYS_Exit;

        *error_code = HI_MPI_VB_Exit();
	if (*error_code != HI_SUCCESS) return ERR_HI_MPI_VB_Exit;

        *error_code = HI_MPI_VB_SetConf(&stVbConf);
	if (*error_code != HI_SUCCESS) return ERR_HI_MPI_VB_SetConf;
    
        *error_code = HI_MPI_VB_Init();
	if (*error_code != HI_SUCCESS) return ERR_HI_MPI_VB_Init;

        stSysConf.u32AlignWidth = 64;

        *error_code = HI_MPI_SYS_SetConf(&stSysConf);
	if (*error_code != HI_SUCCESS) return ERR_HI_MPI_SYS_SetConf;

        *error_code = HI_MPI_SYS_Init();
	if (*error_code != HI_SUCCESS) return ERR_HI_MPI_SYS_Init;

    return ERR_NONE;
}
*/
import "C"

import (
	"application/pkg/mpp/error"
	"application/pkg/logger"

    "application/pkg/mpp/cmos"
)
func Init() {
        var errorCode C.uint

        //switch err := C.mpp3_sys_init(&errorCode); err {
        switch err := C.mpp3_sys_init(&errorCode, C.uint(cmos.Width()), C.uint(cmos.Height())); err {

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
        case C.ERR_HI_MPI_VB_SetConf:
                logger.Log.Fatal().
                        Str("func", "HI_MPI_VB_SetConf()").
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

