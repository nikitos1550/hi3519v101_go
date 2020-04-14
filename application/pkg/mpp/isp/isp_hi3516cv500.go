//+build arm
//+build hi3516cv500

package isp

/*
#include "../include/mpp_v4.h"

#include <stdio.h>
#include <string.h>
#include <pthread.h>

#define ERR_NONE    0
#define ERR_GENERAL 1
#define ERR_MPP     2

static pthread_t mpp4_isp_thread_pid;

HI_VOID* mpp4_isp_thread(HI_VOID *param){
    int error_code = 0;
    printf("C DEBUG: starting HI_MPI_ISP_Run...\n");
    error_code = HI_MPI_ISP_Run(0);
    printf("C DEBUG: HI_MPI_ISP_Run %d\n", error_code);
    //return error_code;
}

ISP_PUB_ATTR_S ISP_PUB_ATTR_IMX335_MIPI_5M_30FPS =
{
    {0, 0, 2592, 1944},
    {2592, 1944},
    30,
    BAYER_RGGB,
    WDR_MODE_NONE,
    0,
};

ISP_PUB_ATTR_S ISP_PUB_ATTR_IMX335_MIPI_5M_30FPS_WDR2TO1 =
{
    {0, 0, 2592, 1944},
    {2592, 1944},
    30,
    BAYER_RGGB,
    WDR_MODE_2To1_LINE,
    0,
};

ISP_PUB_ATTR_S ISP_PUB_ATTR_IMX335_MIPI_4M_30FPS =
{
    {0, 0, 2592, 1536},
    {2592, 1944},
    30,
    BAYER_RGGB,
    WDR_MODE_NONE,
    0,
};

ISP_PUB_ATTR_S ISP_PUB_ATTR_IMX335_MIPI_4M_30FPS_WDR2TO1 =
{
    {0, 0, 2592, 1536},
    {2592, 1944},
    30,
    BAYER_RGGB,
    WDR_MODE_2To1_LINE,
    0,
};


int mpp4_isp_init(int *error_code) {
    *error_code = 0;

   //VI_CreateIsp
    ISP_PUB_ATTR_S stPubAttr;
    memcpy(&stPubAttr, &ISP_PUB_ATTR_IMX335_MIPI_5M_30FPS, sizeof(ISP_PUB_ATTR_S));    
    //stPubAttr.enWDRMode = WDR_MODE_NONE;
    
    ALG_LIB_S stAeLib;
    ALG_LIB_S stAwbLib;
    const ISP_SNS_OBJ_S* pstSnsObj;

    pstSnsObj = &stSnsImx335Obj;

    stAeLib.s32Id = 0;
    stAwbLib.s32Id = 0;
    strncpy(stAeLib.acLibName, HI_AE_LIB_NAME, sizeof(HI_AE_LIB_NAME));
    strncpy(stAwbLib.acLibName, HI_AWB_LIB_NAME, sizeof(HI_AWB_LIB_NAME));
    
    *error_code = pstSnsObj->pfnRegisterCallback(0, &stAeLib, &stAwbLib);
    if (*error_code != HI_SUCCESS) {
        return ERR_GENERAL;
    }

    ISP_SNS_COMMBUS_U uSnsBusInfo;
    ISP_SNS_TYPE_E enBusType;

    enBusType = ISP_SNS_I2C_TYPE;
    uSnsBusInfo.s8I2cDev = 0;

    *error_code = pstSnsObj->pfnSetBusInfo(0, uSnsBusInfo);
    if (*error_code != HI_SUCCESS) return ERR_GENERAL;
    
    *error_code = HI_MPI_AE_Register(0, &stAeLib);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    *error_code = HI_MPI_AWB_Register(0, &stAwbLib);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    *error_code = HI_MPI_ISP_MemInit(0);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    *error_code = HI_MPI_ISP_SetPubAttr(0, &ISP_PUB_ATTR_IMX335_MIPI_5M_30FPS);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    *error_code = HI_MPI_ISP_Init(0);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    if (pthread_create(&mpp4_isp_thread_pid, 0, (void* (*)(void*))mpp4_isp_thread, NULL) != 0) {
        return ERR_GENERAL;
    }

    //pthread_attr_t* pstAttr = NULL;
    //ISP_DEV IspDev = 0;
    //ret = pthread_create(&g_IspPid[0], pstAttr, SAMPLE_COMM_ISP_Thread, (HI_VOID*)IspDev);
    //if (0 != ret)
   // {
   //     printf("create isp running thread failed!, error: %d\n", ret);
    //    return -1;
   // }

	return ERR_NONE;
}



*/
import "C"

import (
         "application/pkg/mpp/error"
        
        "application/pkg/logger"
)

func Init() {
    var errorCode C.int

        switch err := C.mpp4_isp_init(&errorCode); err {
    case C.ERR_NONE:
        logger.Log.Debug().
                Msg("C.mpp4_isp_init() ok")
    case C.ERR_MPP:
        logger.Log.Fatal().
                Int("error", int(errorCode)).
                Str("error_desc", error.Resolve(int64(errorCode))).
                Msg("C.mpp4_isp_init() mpp error ")
    default:
            logger.Log.Fatal().
                Int("error", int(err)).
                Msg("Unexpected return of C.mpp4_isp_init()")
        }

}
