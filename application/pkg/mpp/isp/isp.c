#include "isp.h"

static pthread_t mpp_isp_thread_pid;

static void* mpp_isp_thread(HI_VOID *param){
    GO_LOG_ISP(LOGGER_TRACE, "HI_MPI_ISP_Run");
    #if HI_MPP == 1
        HI_MPI_ISP_Run();
    #elif HI_MPP >= 2
        HI_MPI_ISP_Run(0);
    #endif
    GO_LOG_ISP(LOGGER_ERROR, "HI_MPI_ISP_Run failed");
}

static inline int64_t mpp_isp_register_lib_ae(char * lib) {
    ALG_LIB_S stLib;

    strcpy(stLib.acLibName, lib);
    stLib.s32Id = 0;

    #if HI_MPP == 1
        return HI_MPI_AE_Register(&stLib);    
    #elif HI_MPP >= 2
        return HI_MPI_AE_Register(0, &stLib);
    #endif
}

static inline int mpp_isp_register_lib_awb(char * lib) {
    ALG_LIB_S stLib;

    strcpy(stLib.acLibName, lib);           
    stLib.s32Id = 0;

    #if HI_MPP == 1
        return HI_MPI_AWB_Register(&stLib);
    #elif HI_MPP >= 2
        return HI_MPI_AWB_Register(0, &stLib);
    #endif
} 

static inline int mpp_isp_register_lib_af(char * lib) {
    ALG_LIB_S stLib;

    strcpy(stLib.acLibName, lib);           
    stLib.s32Id = 0;

    #if HI_MPP == 1
        return HI_MPI_AF_Register(&stLib);
    #elif HI_MPP >=2
        return HI_MPI_AF_Register(0, &stLib);
    #endif
} 

int mpp_isp_init(error_in *err, mpp_isp_init_in *in) { 
    int general_error_code = 0;

    DO_OR_RETURN_ERR_MPP(err, mpp_isp_register_lib_ae, HI_AE_LIB_NAME);

    DO_OR_RETURN_ERR_MPP(err, mpp_isp_register_lib_awb, HI_AWB_LIB_NAME);

    DO_OR_RETURN_ERR_MPP(err, mpp_isp_register_lib_af, HI_AF_LIB_NAME);

    #if HI_MPP == 1
        DO_OR_RETURN_ERR_MPP(err, HI_MPI_ISP_Init); //TODO check is it possible to move func call after HI_MPI_ISP_SetInputTiming,

        ISP_IMAGE_ATTR_S stImageAttr;

        stImageAttr.enBayer      = in->bayer;
        stImageAttr.u16FrameRate = in->fps;
        stImageAttr.u16Width     = in->width;
        stImageAttr.u16Height    = in->height;

        DO_OR_RETURN_ERR_MPP(err, HI_MPI_ISP_SetImageAttr, &stImageAttr);

        ISP_INPUT_TIMING_S stInputTiming;

        stInputTiming.enWndMode         = ISP_WIND_ALL;
        stInputTiming.u16HorWndStart    = 0; //200;         //TODO
        stInputTiming.u16VerWndStart    = 0; //18;          //TODO
        stInputTiming.u16HorWndLength   = in->width;
        stInputTiming.u16VerWndLength   = in->height;

        DO_OR_RETURN_ERR_MPP(err, HI_MPI_ISP_SetInputTiming, &stInputTiming);

    #elif HI_MPP >= 2

        DO_OR_RETURN_ERR_MPP(err, HI_MPI_ISP_MemInit, 0);

        #if HI_MPP == 2 || HI_MPP == 3
            ISP_WDR_MODE_S stWdrMode;

            stWdrMode.enWDRMode  = in->wdr;

            DO_OR_RETURN_ERR_MPP(err, HI_MPI_ISP_SetWDRMode, 0, &stWdrMode);
        #endif

        ISP_PUB_ATTR_S stPubAttr;

        #if defined(HI3516AV200) || HI_MPP == 4
            stPubAttr.stSnsSize.u32Width    = in->width; 
            stPubAttr.stSnsSize.u32Height   = in->height; 
        #endif

        #if HI_MPP == 4
            stPubAttr.enWDRMode             = in->wdr; //WDR_MODE_NONE;
            stPubAttr.u8SnsMode             = 0; //TODO
        #endif

        stPubAttr.enBayer               = in->bayer;
        stPubAttr.f32FrameRate          = in->fps;
        //Start position of the cropping window, image width, and image height
        stPubAttr.stWndRect.s32X        = 0;
        stPubAttr.stWndRect.s32Y        = 0;
        stPubAttr.stWndRect.u32Width    = in->width;
        stPubAttr.stWndRect.u32Height   = in->height;

        DO_OR_RETURN_ERR_MPP(err, HI_MPI_ISP_SetPubAttr, 0, &stPubAttr);

        DO_OR_RETURN_ERR_MPP(err, HI_MPI_ISP_Init, 0);
    #endif

    DO_OR_RETURN_ERR_GENERAL(err, pthread_create, &mpp_isp_thread_pid, 0, (void* (*)(void*))mpp_isp_thread, NULL);

    return ERR_NONE;
}
