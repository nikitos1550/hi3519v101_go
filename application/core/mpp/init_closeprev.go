package mpp

/*
#include "./include/mpp.h"

//#include "isp_main.h"

#include <string.h>
#include <stdio.h>

#define ERR_NONE    0
#define ERR_MPP     1

int mpp3_sys_exit(int *error_code) {
    *error_code = 0;
    *error_code = HI_MPI_SYS_Exit();
    printf("sys exit error code %d\n", *error_code);
    if (*error_code != HI_SUCCESS) return ERR_MPP;
    return ERR_NONE;
}

int mpp3_vb_exit(int *error_code) {
    *error_code = 0;
    *error_code = HI_MPI_VB_Exit();
    if (*error_code != HI_SUCCESS) return ERR_MPP;
    return ERR_NONE;
}

static inline unsigned int mpp_isp_unregister_lib_ae(char * lib) {
    ALG_LIB_S stLib;

    strcpy(stLib.acLibName, lib);
    stLib.s32Id = 0;

    //printf("%s\n", stLib.acLibName);

    #if HI_MPP == 1
        return HI_MPI_AE_UnRegister(&stLib);
    #elif HI_MPP >= 2
        return HI_MPI_AE_UnRegister(0, &stLib);
    #endif
}

static inline unsigned int mpp_isp_unregister_lib_awb(char * lib) {
    ALG_LIB_S stLib;

    strcpy(stLib.acLibName, lib);
    stLib.s32Id = 0;

    //printf("%s\n", stLib.acLibName);

    #if HI_MPP == 1
        return HI_MPI_AWB_UnRegister(&stLib);
    #elif HI_MPP >= 2
        return HI_MPI_AWB_UnRegister(0, &stLib);
    #endif
}

static inline unsigned int mpp_isp_unregister_lib_af(char * lib) {
    ALG_LIB_S stLib;

    strcpy(stLib.acLibName, lib);
    stLib.s32Id = 0;

    //printf("%s\n", stLib.acLibName);

    #if HI_MPP == 1
        return HI_MPI_AF_UnRegister(&stLib);
    #elif HI_MPP == 4
        return HI_SUCCESS;
    #elif HI_MPP >=2
        return HI_MPI_AF_UnRegister(0, &stLib);
    #endif
} 

int mpp3_isp_exit(int *error_code) {
    *error_code = 0;

    //ISP_CTX_S *pstIspCtx = HI_NULL;
    //ISP_DBG_CTRL_S *pstDbg = HI_NULL;

    //pstIspCtx = &g_astIspCtx[0];

    //pstDbg = &pstIspCtx->stIspDbg;

    //pstDbg->bDebugEn = HI_FALSE;

    //mpp_isp_unregister_lib_ae(HI_AE_LIB_NAME);
    //mpp_isp_unregister_lib_awb(HI_AWB_LIB_NAME);
    //mpp_isp_unregister_lib_af(HI_AF_LIB_NAME);

    //*error_code = HI_MPI_ISP_Exit(0);
    //if (*error_code != HI_SUCCESS) return ERR_MPP;
    return ERR_NONE;
}
*/
import "C"

import (
    "os"

    "application/core/logger"
)
/*
func IspClose() {
        var errorCode C.int
        switch err := C.mpp3_isp_exit(&errorCode); err {
        case C.ERR_NONE:
            logger.Log.Debug().
                Msg("C.mpp3_isp_exit() ok")
        case C.ERR_MPP:
            logger.Log.Fatal().
                Str("func", "HI_MPI_ISP_Exit()").
                Int("error", int(errorCode)).
                Msg("C.mpp3_isp_exit() error")
        default:
            logger.Log.Fatal().
                Int("error", int(err)).
                Msg("Unexpected return of C.mpp3_isp_exit()")
        }

}
*/
func closePrev() {
    //NOTE should be done, otherwise there can be kernel panic on module unload
    //ATTENTION maybe isp exit should be added as well
    //TODO rework, add error codes, deal with C includes

    /*
    if _, err := os.Stat("/dev/isp_dev"); err == nil { //kernel panic !?
        var errorCode C.int
        switch err := C.mpp3_isp_exit(&errorCode); err {
        case C.ERR_NONE:
            logger.Log.Debug().
                Msg("C.mpp3_isp_exit() ok")
        case C.ERR_MPP:
            logger.Log.Fatal().
                Str("func", "HI_MPI_ISP_Exit()").
                Int("error", int(errorCode)).
                Msg("C.mpp3_isp_exit() error")
        default:
            logger.Log.Fatal().
                Int("error", int(err)).
                Msg("Unexpected return of C.mpp3_isp_exit()")
        }
    }
    */

    if _, err := os.Stat("/dev/sys"); err == nil { 
        var errorCode C.int
        
        switch err := C.mpp3_sys_exit(&errorCode); err {
        case C.ERR_NONE:
            logger.Log.Debug().
                Msg("C.mpp3_sys_exit() ok")
        case C.ERR_MPP:
            logger.Log.Warn().
                Str("func", "HI_MPI_SYS_Exit()").
                Uint("error", uint(errorCode)).
                Msg("C.mpp3_sys_exit() error")
        default:
            logger.Log.Warn().
                Int("error", int(err)).
                Msg("Unexpected return of C.mpp3_sys_exit()")
        } 
        
    }

    if _, err := os.Stat("/dev/vb"); err == nil {      
        var errorCode C.int
        switch err := C.mpp3_vb_exit(&errorCode); err {
        case C.ERR_NONE:
            logger.Log.Debug().
                Msg("C.mpp3_vb_exit() ok")
        case C.ERR_MPP:
            logger.Log.Warn().
                Str("func", "HI_MPI_VB_Exit()").
                Uint("error", uint(errorCode)).
                Msg("C.mpp3_vb_exit() error")
        default:
            logger.Log.Warn().
                Int("error", int(err)).
                Msg("Unexpected return of C.mpp3_vb_exit()")
        }
    }

}
