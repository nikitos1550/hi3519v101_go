#ifndef ERRMPP_H_
#define ERRMPP_H_

#define ERR_NONE                    0
#define ERR_MPP                     1
#define ERR_GENERAL                 2

typedef struct error_in_struct {
    unsigned int f;
    unsigned int mpp; 
    int general;
} error_in;  

typedef unsigned int error_mpp;
typedef int error_general;

//#define MPP_F_HI_MPI_SYS_Exit
//HI_MPI_VB_Exit
//#define MPP_F_

#define RETURN_ERR_MPP(x, y) \
    err->f = x; \
    err->mpp = y; \
    return ERR_MPP

#define ERR_F_HI_MPI_SYS_Init                   1
#define ERR_F_HI_MPI_SYS_Exit                   2
#define ERR_F_HI_MPI_SYS_SetConf                3

#define ERR_F_HI_MPI_VB_Init                    4
#define ERR_F_HI_MPI_VB_Exit                    5
#define ERR_F_HI_MPI_VB_SetConf                 6
#define ERR_F_HI_MPI_VB_SetConfig               8

#define ERR_F_HI_MPI_ISP_Run                    9
#define ERR_F_HI_MPI_ISP_Exit                   10
#define ERR_F_HI_MPI_AE_Register                11
#define ERR_F_HI_MPI_AWB_Register               12
#define ERR_F_HI_MPI_AF_Register                13
#define ERR_F_HI_MPI_ISP_MemInit                14
#define ERR_F_HI_MPI_ISP_SetWDRMode             15
#define ERR_F_HI_MPI_ISP_SetPubAttr             16
#define ERR_F_HI_MPI_ISP_Init                   17
#define ERR_F_HI_MPI_ISP_SetImageAttr           18
#define ERR_F_HI_MPI_ISP_SetInputTiming         19

#define ERR_F_HI_MPI_VI_SetDevAttr              20
#define ERR_F_HI_MPI_VI_EnableDev               21
#define ERR_F_HI_MPI_VI_SetChnAttr              22
#define ERR_F_HI_MPI_VI_SetLDCAttr              23
#define ERR_F_HI_MPI_VI_EnableChn               24

#define ERR_F_HI_MPI_VPSS_CreateGrp             25
#define ERR_F_HI_MPI_VPSS_StartGrp              26

#define ERR_F_HI_MPI_SYS_Bind                   27

#define ERR_F_HI_MPI_VPSS_SetChnAttr            280
#define ERR_F_HI_MPI_VPSS_SetChnMode            28
#define ERR_F_HI_MPI_VPSS_SetDepth              29
#define ERR_F_HI_MPI_VPSS_EnableChn             30
#define ERR_F_HI_MPI_VPSS_DisableChn            31
#define ERR_F_HI_MPI_VPSS_GetChnFrame           32
#define ERR_F_HI_MPI_VPSS_ReleaseChnFrame       33

#define ERR_F_HI_MPI_VENC_CreateChn             34
#define ERR_F_HI_MPI_VENC_StartRecvPic          35
#define ERR_F_HI_MPI_VENC_DestroyChn            36
#define ERR_F_HI_MPI_VENC_StopRecvPic           37

#endif // ERRMPP_H_
