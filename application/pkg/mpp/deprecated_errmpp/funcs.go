//+build arm
//+build debug

package errmpp

/*
#include "errmpp.h"
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
    case C.ERR_F_HI_MPI_VI_SetDevAttr:
        return "HI_MPI_VI_SetDevAttr"
    case C.ERR_F_HI_MPI_VI_EnableDev:
        return "HI_MPI_VI_EnableDev"
    case C.ERR_F_HI_MPI_VI_SetChnAttr:
        return "HI_MPI_VI_SetChnAttr"
    case C.ERR_F_HI_MPI_VI_SetLDCAttr:
        return "HI_MPI_VI_SetLDCAttr"
    case C.ERR_F_HI_MPI_VI_EnableChn:
        return "HI_MPI_VI_EnableChn"
    case C.ERR_F_HI_MPI_VPSS_CreateGrp:
        return "HI_MPI_VPSS_CreateGrp"
    case C.ERR_F_HI_MPI_VPSS_StartGrp:
        return "HI_MPI_VPSS_StartGrp"
    case C.ERR_F_HI_MPI_SYS_Bind:
        return "HI_MPI_SYS_Bind"
    case C.ERR_F_HI_MPI_VPSS_SetChnAttr:
        return "HI_MPI_VPSS_SetChnAttr"
    case C.ERR_F_HI_MPI_VPSS_SetChnMode:
        return "HI_MPI_VPSS_SetChnMode"
    case C.ERR_F_HI_MPI_VPSS_SetDepth:
        return "HI_MPI_VPSS_SetDepth"
    case C.ERR_F_HI_MPI_VPSS_EnableChn:
        return "HI_MPI_VPSS_EnableChn"
    case C.ERR_F_HI_MPI_VPSS_DisableChn:
        return "HI_MPI_VPSS_DisableChn"
    case C.ERR_F_HI_MPI_VPSS_GetChnFrame:
        return "HI_MPI_VPSS_GetChnFrame"
    case C.ERR_F_HI_MPI_VPSS_ReleaseChnFrame:
        return "HI_MPI_VPSS_ReleaseChnFrame"
    case C.ERR_F_HI_MPI_VENC_CreateChn:
        return "HI_MPI_VENC_CreateChn"
    case C.ERR_F_HI_MPI_VENC_StartRecvPic:
        return "HI_MPI_VENC_StartRecvPic"
    case C.ERR_F_HI_MPI_VENC_DestroyChn:
        return "HI_MPI_VENC_DestroyChn"
    case C.ERR_F_HI_MPI_VENC_StopRecvPic:
        return "HI_MPI_VENC_StopRecvPic"
    default:
        logger.Log.Warn().
            Uint("func", f).
            Msg("ERRMPP missed desc")
        return "unknown"
    }
}
