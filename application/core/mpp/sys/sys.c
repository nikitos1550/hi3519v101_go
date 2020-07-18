#include "sys.h"

inline static int mpp_sys_vb_conf(mpp_sys_init_in *in) __attribute__((always_inline));  //TODO

#define CEILING_2_POWER_L(x,a)     ( ((x) + ((a) - 1) ) & ( ~((a) - 1) ) )

//TODO deal with calculation and do it HERE!
inline static int mpp_sys_vb_conf(mpp_sys_init_in *in) {
    #if HI_MPP <= 3
        VB_CONF_S   stVbConf;
    #elif HI_MPP == 4
        VB_CONFIG_S stVbConf;
    #endif

    memset(&stVbConf, 0, sizeof(stVbConf));

    //stVbConf.u32MaxPoolCnt                  = 128;
    #if HI_MPP <= 3
        stVbConf.u32MaxPoolCnt                  = 128;
        stVbConf.astCommPool[0].u32BlkSize  = (CEILING_2_POWER_L(in->width, 128) * CEILING_2_POWER_L(in->height, 128) * 1.5);
    #elif HI_MPP == 4
        stVbConf.u32MaxPoolCnt                  = 512;
        stVbConf.astCommPool[0].u64BlkSize  = (CEILING_2_POWER_L(in->width, 128) * CEILING_2_POWER_L(in->height, 128) * 1.5);
        //stVbConf.astCommPool[0].u64BlkSize  = 7558272;
        //stVbConf.astCommPool[1].u64BlkSize  = 5529600;
    #endif
    stVbConf.astCommPool[0].u32BlkCnt       = in->cnt;
    //stVbConf.astCommPool[1].u32BlkCnt       = 1;

    //stVbConf.astCommPool[0].u32BlkSize      = (CEILING_2_POWER(in->width, 64) * CEILING_2_POWER(in->height, 64) * 1.5); //TODO
    //stVbConf.astCommPool[0].u32BlkSize      = (CEILING_2_POWER(in->width, 128) * CEILING_2_POWER(in->height, 128) * 1.5); //TODO
    //stVbConf.astCommPool[0].u32BlkCnt       = in->cnt;
    //#if defined(HI3516CV500)
    //stVbConf.astCommPool[0].u64BlkSize = COMMON_GetPicBufferSize(   in->width, 
    //                                                                in->height, 
    //                                                                PIXEL_FORMAT_YVU_SEMIPLANAR_420, 
    //                                                                DATA_BITWIDTH_8, 
    //                                                                COMPRESS_MODE_NONE, 
    //                                                                DEFAULT_ALIGN);
    //#elif defined(HI3516EV200)
    //stVbConf.astCommPool[0].u64BlkSize = COMMON_GetPicBufferSize(   in->width,
    //                                                                in->height,
    //                                                                PIXEL_FORMAT_YVU_SEMIPLANAR_420,
    //                                                                DATA_BITWIDTH_8,
    //                                                                COMPRESS_MODE_NONE,
    //                                                                DEFAULT_ALIGN);
    //#endif
    //stVbConf.astCommPool[0].u64BlkSize  = in->width * in->height * 1.5;
    //stVbConf.astCommPool[0].u32BlkCnt   = in->cnt;
    //printf("TMP: %ull\n", stVbConf.astCommPool[0].u64BlkSize);

    #if HI_MPP <= 3
        return HI_MPI_VB_SetConf(&stVbConf);
    #elif HI_MPP == 4
        return HI_MPI_VB_SetConfig(&stVbConf);
    #endif
}

inline static int mpp_sys_sys_conf() {
    #if HI_MPP <= 3
        MPP_SYS_CONF_S stSysConf;

        memset(&stSysConf, 0, sizeof(MPP_SYS_CONF_S));
        stSysConf.u32AlignWidth = 64;

        return HI_MPI_SYS_SetConf(&stSysConf);
    #elif HI_MPP == 4 //TODO
        //MPP_SYS_CONFIG_S stSysConfig;
    
        //memset(&stSysConfig, 0, sizeof(MPP_SYS_CONFIG_S));
        //stSysConf.u32AlignWidth = 64;

        //return HI_MPI_SYS_SetConfig(&stSysConfig);
        return HI_SUCCESS;
    #endif
}

int mpp_sys_init(error_in *err, mpp_sys_init_in *in) {

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_SYS_Exit);

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VB_Exit);

    DO_OR_RETURN_ERR_MPP(err, mpp_sys_vb_conf, in);

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VB_Init);
    
    DO_OR_RETURN_ERR_MPP(err, mpp_sys_sys_conf);
    
    DO_OR_RETURN_ERR_MPP(err, HI_MPI_SYS_Init);

    #if HI_MPP == 4
        VI_VPSS_MODE_S      stVIVPSSMode;
        DO_OR_RETURN_ERR_MPP(err, HI_MPI_SYS_GetVIVPSSMode, &stVIVPSSMode);

        if (1) {
            for (int i = 0; i < VI_MAX_PIPE_NUM; i++) {
                stVIVPSSMode.aenMode[i] = VI_OFFLINE_VPSS_OFFLINE;    //ok
                //stVIVPSSMode.aenMode[i] = VI_OFFLINE_VPSS_ONLINE;       //ok
                //stVIVPSSMode.aenMode[i] = VI_ONLINE_VPSS_OFFLINE;
                //stVIVPSSMode.aenMode[i] = VI_ONLINE_VPSS_ONLINE;
            }
        } else {
            stVIVPSSMode.aenMode[0] = VI_ONLINE_VPSS_ONLINE;
			for (int i = 1; i < VI_MAX_PIPE_NUM; i++) {
            	stVIVPSSMode.aenMode[i] = VI_OFFLINE_VPSS_OFFLINE;
			}
        }
    
        DO_OR_RETURN_ERR_MPP(err, HI_MPI_SYS_SetVIVPSSMode, &stVIVPSSMode);
    #endif

    return ERR_NONE;
}
