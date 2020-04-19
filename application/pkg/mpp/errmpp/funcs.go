//+build arm
//+build debug

package errmpp

/*
#include "error.h"
*/
import "C"

import (
    "application/pkg/logger"
)

func resolveFunc(f uint) string {
    switch f {
    case C.ERR_F_HI_MPI_SYS_Init:
        return "HI_MPI_SYS_Init"
    case C.ERR_F_HI_MPI_SYS_Exit:
        return "HI_MPI_SYS_Exit"
    case C.ERR_F_HI_MPI_SYS_SetConf:
        return "HI_MPI_SYS_SetConf"
    case C.ERR_F_HI_MPI_VB_Init:
        return "HI_MPI_VB_Init"
    case C.ERR_F_HI_MPI_VB_Exit:
        return "HI_MPI_VB_Exit"
    case C.ERR_F_HI_MPI_VB_SetConf:
        return "HI_MPI_VB_SetConf"
    case C.ERR_F_HI_MPI_VB_SetConfig:
        return "HI_MPI_VB_SetConfig"
    case C.ERR_F_HI_MPI_ISP_Run:
        return "HI_MPI_ISP_Run"
    case C.ERR_F_HI_MPI_ISP_Exit:
        return "HI_MPI_ISP_Exit"
    case C.ERR_F_HI_MPI_AE_Register:
        return "HI_MPI_AE_Register"
    case C.ERR_F_HI_MPI_AWB_Register:
        return "HI_MPI_AWB_Register"
    case C.ERR_F_HI_MPI_AF_Register:
        return "HI_MPI_AF_Register"
    case C.ERR_F_HI_MPI_ISP_MemInit:
        return "HI_MPI_ISP_MemInit"
    case C.ERR_F_HI_MPI_ISP_SetWDRMode:
        return "HI_MPI_ISP_SetWDRMode"
    case C.ERR_F_HI_MPI_ISP_SetPubAttr:
        return "HI_MPI_ISP_SetPubAttr"
    case C.ERR_F_HI_MPI_ISP_Init:
        return "HI_MPI_ISP_Init"
    case C.ERR_F_HI_MPI_ISP_SetImageAttr:
        return "HI_MPI_ISP_SetImageAttr"
    case C.ERR_F_HI_MPI_ISP_SetInputTiming:
        return "HI_MPI_ISP_SetInputTiming"
    default:
        logger.Log.Warn().
            Uint("func", f).
            Msg("ERRMPP missed desc")
        return "unknown"
    }
}
