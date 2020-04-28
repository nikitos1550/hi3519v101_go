#include "sys.h"

#ifdef HI3516CV100
int mpp_sys_init(error_in *err, mpp_sys_init_in *in) {
    unsigned in mpp_error_code = 0;
    
    mpp_error_code = HI_MPI_SYS_Exit();
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_SYS_Exit, mpp_error_code);
    }

    mpp_error_code = HI_MPI_VB_Exit();
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_VB_Exit, mpp_error_code);
    }

    VB_CONF_S stVbConf;

    memset(&stVbConf, 0, sizeof(VB_CONF_S));
    stVbConf.u32MaxPoolCnt                  = 128;
    stVbConf.astCommPool[0].u32BlkSize      = (CEILING_2_POWER(in->width, 64) * CEILING_2_POWER(in->height, 64) * 1.5);
    stVbConf.astCommPool[0].u32BlkCnt       = in->cnt;

    
    mpp_error_code = HI_MPI_VB_SetConf(&stVbConf);
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_VB_SetConf, mpp_error_code);
    }
    
    mpp_error_code = HI_MPI_VB_Init();
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_VB_Init, mpp_error_code);
    }
    
    MPP_SYS_CONF_S stSysConf;

    memset(&stSysConf, 0, sizeof(MPP_SYS_CONF_S));
    stSysConf.u32AlignWidth = 64;
    
    mpp_error_code = HI_MPI_SYS_SetConf(&stSysConf);
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_SYS_SetConf, mpp_error_code);
    }
    
    mpp_error_code = HI_MPI_SYS_Init();
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_SYS_Init, mpp_error_code);
    }

    return ERR_NONE;
}
#endif // HI3516CV100

#ifdef HI3516CV300
int mpp_sys_init(error_in *err, mpp_sys_init_in * in) {
    unsigned int mpp_error_code = 0;

    mpp_error_code = HI_MPI_SYS_Exit();
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_SYS_Exit, mpp_error_code);
    }

    mpp_error_code = HI_MPI_VB_Exit();
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_VB_Exit, mpp_error_code);
    }

    VB_CONF_S               stVbConf;

    memset(&stVbConf,0,sizeof(VB_CONF_S));
    stVbConf.u32MaxPoolCnt = 128;
    stVbConf.astCommPool[0].u32BlkSize =(CEILING_2_POWER(in->width, 64) * CEILING_2_POWER(in->height, 64) * 1.5);
    stVbConf.astCommPool[0].u32BlkCnt = in->cnt;

    mpp_error_code = HI_MPI_VB_SetConf(&stVbConf);
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_VB_SetConf, mpp_error_code);
    }
    
    mpp_error_code = HI_MPI_VB_Init();
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_VB_Init, mpp_error_code);
    }

    MPP_SYS_CONF_S      stSysConf;

    memset(&stSysConf, 0, sizeof(MPP_SYS_CONF_S));
    stSysConf.u32AlignWidth = 64;

    mpp_error_code = HI_MPI_SYS_SetConf(&stSysConf);
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_SYS_SetConf, mpp_error_code);
    }

    mpp_error_code = HI_MPI_SYS_Init();
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_SYS_Init, mpp_error_code);
    }

    return ERR_NONE;
}
#endif // HI3516CV300

#ifdef HI3516CV500
int mpp_sys_init(error_in *err, mpp_sys_init_in *in) {
    unsigned int mpp_error_code = 0;

    mpp_error_code = HI_MPI_SYS_Exit();
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_, mpp_error_code);
    }

    mpp_error_code = HI_MPI_VB_Exit();
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_, mpp_error_code);
    }


    VB_CONFIG_S        stVbConf;

    memset(&stVbConf,0,sizeof(VB_CONFIG_S));
    stVbConf.u32MaxPoolCnt              = 2;
    stVbConf.astCommPool[0].u64BlkSize = COMMON_GetPicBufferSize(   in->width, 
                                                                    in->height, 
                                                                    PIXEL_FORMAT_YVU_SEMIPLANAR_420, 
                                                                    DATA_BITWIDTH_8, 
                                                                    COMPRESS_MODE_SEG, 
                                                                    DEFAULT_ALIGN);
    stVbConf.astCommPool[0].u32BlkCnt = in->cnt;

    mpp_error_code = HI_MPI_VB_SetConfig(&stVbConf);
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_, mpp_error_code);
    }

    mpp_error_code = HI_MPI_VB_Init();
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_, mpp_error_code);
    }

    MPP_SYS_CONF_S stSysConf;

    memset(&stSysConf, 0, sizeof(MPP_SYS_CONF_S));
    stSysConf.u32AlignWidth = 64;

    mpp_error_code = HI_MPI_SYS_SetConf(&stSysConf);
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_, mpp_error_code);
    }

    mpp_error_code = HI_MPI_SYS_Init();
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_, mpp_error_code);
    }

    return ERR_NONE;
}
#endif // HI3516CV500

#ifdef HI3516AV100
int mpp_sys_init(error_in *err, mpp_sys_init_in *in) {
    unsigned int mpp_error_code = 0;

    mpp_error_code = HI_MPI_SYS_Exit();
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_SYS_Exit, mpp_error_code);
    }

    mpp_error_code = HI_MPI_VB_Exit();
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_VB_Exit, mpp_error_code);
    }

    VB_CONF_S stVbConf;

    memset(&stVbConf, 0, sizeof(VB_CONF_S));
    stVbConf.u32MaxPoolCnt                  = 128;
    stVbConf.astCommPool[0].u32BlkSize      = (CEILING_2_POWER(in->width, 64) * CEILING_2_POWER(in->height, 64) * 1.5);
    stVbConf.astCommPool[0].u32BlkCnt       = in->cnt;

    mpp_error_code = HI_MPI_VB_SetConf(&stVbConf);
    if(mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_VB_SetConf, mpp_error_code);
    }

    mpp_error_code = HI_MPI_VB_Init();
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_VB_Init, mpp_error_code);
    }

    MPP_SYS_CONF_S stSysConf;

    memset(&stSysConf, 0, sizeof(MPP_SYS_CONF_S));
    stSysConf.u32AlignWidth = 64;

    mpp_error_code = HI_MPI_SYS_SetConf(&stSysConf);
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_SYS_SetConf, mpp_error_code);
    }

    mpp_error_code = HI_MPI_SYS_Init();
    if(mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_SYS_Init, mpp_error_code);
    }

    return ERR_NONE;
}
#endif // HI3516AV100

#ifdef HI3516AV200
int mpp_sys_init(error_in *err, mpp_sys_init_in *in) {
    unsigned int mpp_error_code = 0;

    mpp_error_code = HI_MPI_SYS_Exit();
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_SYS_Exit, mpp_error_code);
    }

    mpp_error_code = HI_MPI_VB_Exit();
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_VB_Exit, mpp_error_code);
    }

    VB_CONF_S stVbConf;

    memset(&stVbConf, 0, sizeof(VB_CONF_S));
    stVbConf.u32MaxPoolCnt                  = 128;
    stVbConf.astCommPool[0].u32BlkSize      = (CEILING_2_POWER(in->width, 64) * CEILING_2_POWER(in->height, 64) * 1.5);
    stVbConf.astCommPool[0].u32BlkCnt       = in->cnt;

    mpp_error_code = HI_MPI_VB_SetConf(&stVbConf);
    if(mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_VB_SetConf, mpp_error_code);
    }

    mpp_error_code = HI_MPI_VB_Init();
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_VB_Init, mpp_error_code);
    }
    
    MPP_SYS_CONF_S stSysConf;

    memset(&stSysConf, 0, sizeof(MPP_SYS_CONF_S));
    stSysConf.u32AlignWidth = 64;

    mpp_error_code = HI_MPI_SYS_SetConf(&stSysConf);
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_SYS_SetConf, mpp_error_code);
    }

    mpp_error_code = HI_MPI_SYS_Init();
    if(mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_SYS_Init, mpp_error_code);
    }

    return ERR_NONE;
}
#endif // HI3516AV200

#ifdef HI3516EV200
static int mpp_sys_init(error_in *err, mpp_sys_init_in *in) {
    unsigned int mpp_error_code = 0;

    return ERR_GENERAL;
}
#endif // HI3516EV200
