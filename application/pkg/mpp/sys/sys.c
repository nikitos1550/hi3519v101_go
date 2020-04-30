#include "sys.h"

inline static int mpp_sys_vb_conf(mpp_sys_init_in *in) __attribute__((always_inline));  //TODO

#if HI_MPP <= 3
inline static int mpp_sys_vb_conf(mpp_sys_init_in *in) {
    VB_CONF_S stVbConf;

    memset(&stVbConf, 0, sizeof(VB_CONF_S));
    stVbConf.u32MaxPoolCnt                  = 1;
    stVbConf.astCommPool[0].u32BlkSize      = (CEILING_2_POWER(in->width, 128) * CEILING_2_POWER(in->height, 128) * 1.5); //TODO
    stVbConf.astCommPool[0].u32BlkCnt       = in->cnt;
                        
    return HI_MPI_VB_SetConf(&stVbConf);
}
#elif HI_MPP == 4
inline static int mpp_sys_vb_conf(mpp_sys_init_in *in) {

    VB_CONFIG_S        stVbConf;

    memset(&stVbConf,0,sizeof(VB_CONFIG_S));
    stVbConf.u32MaxPoolCnt              = 1;
    stVbConf.astCommPool[0].u64BlkSize = COMMON_GetPicBufferSize(   in->width, 
                                                                    in->height, 
                                                                    PIXEL_FORMAT_YVU_SEMIPLANAR_420, 
                                                                    DATA_BITWIDTH_8, 
                                                                    COMPRESS_MODE_SEG, 
                                                                    DEFAULT_ALIGN);
    stVbConf.astCommPool[0].u32BlkCnt = in->cnt;

    return HI_MPI_VB_SetConfig(&stVbConf);
}
#endif

int mpp_sys_init(error_in *err, mpp_sys_init_in *in) {

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_SYS_Exit);

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VB_Exit);

    DO_OR_RETURN_ERR_MPP(err, mpp_sys_vb_conf, in);

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VB_Init);
    
    MPP_SYS_CONF_S stSysConf;

    memset(&stSysConf, 0, sizeof(MPP_SYS_CONF_S));
    stSysConf.u32AlignWidth = 64;
    
    DO_OR_RETURN_ERR_MPP(err, HI_MPI_SYS_SetConf, &stSysConf);
    
    DO_OR_RETURN_ERR_MPP(err, HI_MPI_SYS_Init);

    return ERR_NONE;
}
