//+build hi3516cv300 hi3516av200

package sys

/*
int hi3516av200_sys_init(struct hi3516av200_cmos * c) {
    int error_code = 0;

    error_code = HI_MPI_SYS_Exit();
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: TODO\n");
        return ERR_MPP;
    }
    printf("C DEBUG: HI_MPI_SYS_Exit ok\n");

    error_code = HI_MPI_VB_Exit();
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: TODO\n");
        return ERR_MPP;
    }
    printf("C DEBUG: HI_MPI_VB_Exit ok\n");

    VB_CONF_S stVbConf;

    memset(&stVbConf, 0, sizeof(VB_CONF_S));
    stVbConf.u32MaxPoolCnt                  = 128;
    stVbConf.astCommPool[0].u32BlkSize      = (CEILING_2_POWER(c->width, 64) * CEILING_2_POWER(c->height, 64) * 1.5);
    stVbConf.astCommPool[0].u32BlkCnt       = 10;

    error_code = HI_MPI_VB_SetConf(&stVbConf);
    if(error_code != HI_SUCCESS) {
        printf("C DEBUG: TODO\n");
        return ERR_MPP;
    }
    printf("C DEBUG: HI_MPI_VB_SetConf ok\n");

    error_code = HI_MPI_VB_Init();
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: TODO\n");
        return ERR_MPP;
    }
    printf("C DEBUG: HI_MPI_VB_Init ok\n");

    MPP_SYS_CONF_S stSysConf;

    stSysConf.u32AlignWidth = 64;

    error_code = HI_MPI_SYS_SetConf(&stSysConf);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: TODO\n");
        return ERR_MPP;
    }
    printf("C DEBUG: HI_MPI_SYS_SetConf ok\n");

    error_code = HI_MPI_SYS_Init();
    if(error_code != HI_SUCCESS) {
        printf("C DEBUG: TODO\n");
        return ERR_MPP;
    }
    printf("C DEBUG: HI_MPI_SYS_Init ok\n");

    return ERR_NONE;
}
*/
//import "C"
