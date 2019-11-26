//+build hi3516cv300 hi3516av200

package isp

/*
HI_VOID* hi3516av200_isp_thread(HI_VOID *param){
    int error_code = 0;
    printf("C DEBUG: starting HI_MPI_ISP_Run...\n");
    error_code = HI_MPI_ISP_Run(0);
    printf("C DEBUG: HI_MPI_ISP_Run %d\n", error_code);
    //return error_code;
}

static pthread_t hi3516av200_isp_thread_pid;

int hi3516av200_isp_init(struct hi3516av200_cmos * c) {
    int error_code = 0;

    ISP_PUB_ATTR_S stPubAttr;
    ALG_LIB_S stLib;

    error_code = HI_MPI_ISP_Exit(0);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: HI_MPI_ISP_Exit failed with %#x!\n", error_code);
        return ERR_MPP;
    }

    ALG_LIB_S stAeLib;
    ALG_LIB_S stAwbLib;
    ALG_LIB_S stAfLib;

    stAeLib.s32Id = 0;
    stAwbLib.s32Id = 0;
    stAfLib.s32Id = 0;
    strncpy(stAeLib.acLibName,  HI_AE_LIB_NAME,     sizeof(HI_AE_LIB_NAME));
    strncpy(stAwbLib.acLibName, HI_AWB_LIB_NAME,    sizeof(HI_AWB_LIB_NAME));
    strncpy(stAfLib.acLibName,  HI_AF_LIB_NAME,     sizeof(HI_AF_LIB_NAME));

    if (c->snsobj->pfnRegisterCallback != HI_NULL) {
        error_code = c->snsobj->pfnRegisterCallback(0, &stAeLib, &stAwbLib);
        if (error_code != HI_SUCCESS) {
            printf("C DEBUG: sensor_register_callback failed with %#x!\n", error_code);
            return ERR_GENERAL;
        }
    } else {
        printf("C DEBUG: sensor_register_callback failed with HI_NULL!\n");
        return ERR_GENERAL;
    }

    stLib.s32Id = 0;
    strcpy(stLib.acLibName, HI_AE_LIB_NAME);
    error_code = HI_MPI_AE_Register(0, &stLib);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: HI_MPI_AE_Register failed!\n");
        return ERR_MPP;
    }

    stLib.s32Id = 0;
    strcpy(stLib.acLibName, HI_AWB_LIB_NAME);
    error_code = HI_MPI_AWB_Register(0, &stLib);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: HI_MPI_AWB_Register failed!\n");
        return ERR_MPP;
    }

    stLib.s32Id = 0;
    strcpy(stLib.acLibName, HI_AF_LIB_NAME);
    error_code = HI_MPI_AF_Register(0, &stLib);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: HI_MPI_AF_Register failed!\n");
        return ERR_MPP;
    }

    error_code = HI_MPI_ISP_MemInit(0);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: HI_MPI_ISP_Init failed!\n");
        return ERR_MPP;
    }

    ISP_WDR_MODE_S stWdrMode;
    stWdrMode.enWDRMode  = WDR_MODE_NONE;

    error_code = HI_MPI_ISP_SetWDRMode(0, &stWdrMode);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: start ISP WDR failed!\n");
        return ERR_NONE;
    }
    //TODO WDR modes support

    stPubAttr.enBayer               = c->bayer;
    stPubAttr.f32FrameRate          = c->fps;
    stPubAttr.stWndRect.s32X        = 0;
    stPubAttr.stWndRect.s32Y        = 0;
    stPubAttr.stWndRect.u32Width    = c->width;     //TODO What is WND rect?
    stPubAttr.stWndRect.u32Height   = c->height;    //TODO
    stPubAttr.stSnsSize.u32Width    = c->width;
    stPubAttr.stSnsSize.u32Height   = c->height;

    error_code = HI_MPI_ISP_SetPubAttr(0, &stPubAttr);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: HI_MPI_ISP_SetPubAttr failed with %#x!\n", error_code);
        return ERR_MPP;
    }

    error_code = HI_MPI_ISP_Init(0);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: HI_MPI_ISP_Init failed!\n");
        return ERR_MPP;
    }

    if (pthread_create(&hi3516av200_isp_thread_pid, 0, (void* (*)(void*))hi3516av200_isp_thread, NULL) != 0) {
        printf("C DEBUG: create isp running thread failed!\n");
        return ERR_GENERAL;
    }

    return ERR_NONE;
}

*/
//import "C"
