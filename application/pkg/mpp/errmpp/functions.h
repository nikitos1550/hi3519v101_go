#ifndef FUNCTIONS_H_
#define FUNCTIONS_H_

//SYS
#define ERR_F_HI_MPI_SYS_Init               1
#define ERR_F_HI_MPI_SYS_Exit               2
#define ERR_F_HI_MPI_SYS_SetConf            3
#define ERR_F_HI_MPI_SYS_Bind               4
//VB
#define ERR_F_HI_MPI_VB_Init                5
#define ERR_F_HI_MPI_VB_Exit                6
#define ERR_F_HI_MPI_VB_SetConf             7
#define ERR_F_HI_MPI_VB_SetConfig           8
//ISP
#define ERR_F_HI_MPI_ISP_Run                9
#define ERR_F_HI_MPI_ISP_Exit               10
#define ERR_F_HI_MPI_AE_Register            11
#define ERR_F_HI_MPI_AWB_Register           12
#define ERR_F_HI_MPI_AF_Register            13
#define ERR_F_HI_MPI_ISP_MemInit            14
#define ERR_F_HI_MPI_ISP_SetWDRMode         15
#define ERR_F_HI_MPI_ISP_SetPubAttr         16
#define ERR_F_HI_MPI_ISP_Init               17
#define ERR_F_HI_MPI_ISP_SetImageAttr       18
#define ERR_F_HI_MPI_ISP_SetInputTiming     19
//VI
#define ERR_F_HI_MPI_VI_SetDevAttr          20
#define ERR_F_HI_MPI_VI_EnableDev           21
#define ERR_F_HI_MPI_VI_SetChnAttr          22
#define ERR_F_HI_MPI_VI_SetLDCAttr          23
#define ERR_F_HI_MPI_VI_EnableChn           24
//VPSS
#define ERR_F_HI_MPI_VPSS_CreateGrp         25
#define ERR_F_HI_MPI_VPSS_StartGrp          26
#define ERR_F_HI_MPI_VPSS_SetChnAttr        27
#define ERR_F_HI_MPI_VPSS_SetChnMode        28
#define ERR_F_HI_MPI_VPSS_SetDepth          29
#define ERR_F_HI_MPI_VPSS_EnableChn         30
#define ERR_F_HI_MPI_VPSS_DisableChn        31
#define ERR_F_HI_MPI_VPSS_GetChnFrame       32   
#define ERR_F_HI_MPI_VPSS_ReleaseChnFrame   33
//VENC
#define ERR_F_HI_MPI_VENC_CreateChn         34 
#define ERR_F_HI_MPI_VENC_StartRecvPic      35
#define ERR_F_HI_MPI_VENC_DestroyChn        36 
#define ERR_F_HI_MPI_VENC_StopRecvPic       37

//TODO
#define ERR_F_HI_MPI_VI_SetWDRAttr          38

#endif // FUNCTIONS_H_
