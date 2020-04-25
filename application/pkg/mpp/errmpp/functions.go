//+build arm
//+build debug

package errmpp

var functions = [...]string {
    "UNKNOWN",
    //SYS
    "HI_MPI_SYS_Init",
    "HI_MPI_SYS_Exit",
    "HI_MPI_SYS_SetConf",
    "HI_MPI_SYS_Bind",
    //VB
    "HI_MPI_VB_Init",
    "HI_MPI_VB_Exit",
    "HI_MPI_VB_SetConf",
    "HI_MPI_VB_SetConfig",
    //ISP
    "HI_MPI_ISP_Run",
    "HI_MPI_ISP_Exit",
    "HI_MPI_AE_Register",
    "HI_MPI_AWB_Register",
    "HI_MPI_AF_Register",
    "HI_MPI_ISP_MemInit",
    "HI_MPI_ISP_SetWDRMode",
    "HI_MPI_ISP_SetPubAttr",
    "HI_MPI_ISP_Init",
    "HI_MPI_ISP_SetImageAttr",
    "HI_MPI_ISP_SetInputTiming",
    //VI
    "HI_MPI_VI_SetDevAttr",
    "HI_MPI_VI_EnableDev",
    "HI_MPI_VI_SetChnAttr",
    "HI_MPI_VI_SetLDCAttr",
    "HI_MPI_VI_EnableChn",
    //VPSS
    "HI_MPI_VPSS_CreateGrp",
    "HI_MPI_VPSS_StartGrp",
    "HI_MPI_VPSS_SetChnAttr",
    "HI_MPI_VPSS_SetChnMode",
    "HI_MPI_VPSS_SetDepth",
    "HI_MPI_VPSS_EnableChn",
    "HI_MPI_VPSS_DisableChn",
    "HI_MPI_VPSS_GetChnFrame",
    "HI_MPI_VPSS_ReleaseChnFrame",
    //VENC
    "HI_MPI_VENC_CreateChn",
    "HI_MPI_VENC_StartRecvPic",
    "HI_MPI_VENC_DestroyChn",
    "HI_MPI_VENC_StopRecvPic",

    //TODO
    "HI_MPI_VI_SetWDRAttr",
}
