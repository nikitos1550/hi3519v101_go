#include "isp.h"

static pthread_t mpp_isp_thread_pid;

static void* mpp_isp_thread(HI_VOID *param){
    int error_code = 0;
    GO_LOG_ISP(LOGGER_TRACE, "HI_MPI_ISP_Run");
    #if defined(HI_MPP_V1)
        error_code = HI_MPI_ISP_Run();
    #elif defined(HI_MPP_V2) \
        || defined(HI_MPP_V3) \
        || defined(HI_MPP_V4)
        error_code = HI_MPI_ISP_Run(0);
    #endif
    error_code = HI_MPI_ISP_Run(0);
    GO_LOG_ISP(LOGGER_ERROR, "HI_MPI_ISP_Run failed");
}

static inline int mpp_isp_register_lib_ae(char * lib) {
    ALG_LIB_S stLib;
    strcpy(stLib.acLibName, lib);
    stLib.s32Id = 0;
    #if defined(HI_MPP_V1)
        return HI_MPI_AE_Register(&stLib);    
    #elif defined(HI_MPP_V2) \
        || defined(HI_MPP_V3) \
        || defined(HI_MPP_V4) 
        return HI_MPI_AE_Register(0, &stLib);
    #endif
}

static inline int mpp_isp_register_lib_awb(char * lib) {
    ALG_LIB_S stLib;
    strcpy(stLib.acLibName, lib);           
    stLib.s32Id = 0;
    #if defined(HI_MPP_V1)
        return HI_MPI_AWB_Register(&stLib);
    #elif defined(HI_MPP_V2) \
        || defined(HI_MPP_V3) \
        || defined(HI_MPP_V4)  
        return HI_MPI_AWB_Register(0, &stLib);
    #endif
} 

static inline int mpp_isp_register_lib_af(char * lib) {
    ALG_LIB_S stLib;
    strcpy(stLib.acLibName, lib);           
    stLib.s32Id = 0;
    #if defined(HI_MPP_V1)
        return HI_MPI_AF_Register(&stLib);
    #elif defined(HI_MPP_V2) \
        || defined(HI_MPP_V3) \
        || defined(HI_MPP_V4) 
        return HI_MPI_AF_Register(0, &stLib);
    #endif
} 

int mpp_isp_init(error_in *err, mpp_isp_init_in *in) { 
    unsigned int mpp_error_code = 0;
    int general_error_code = 0;

    mpp_error_code = mpp_isp_register_lib_ae(HI_AE_LIB_NAME);
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_AE_Register, mpp_error_code);
    }

    mpp_error_code = mpp_isp_register_lib_awb(HI_AWB_LIB_NAME);
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_AWB_Register, mpp_error_code);
    }

    mpp_error_code = mpp_isp_register_lib_af(HI_AF_LIB_NAME);
    if (mpp_error_code != HI_SUCCESS) {
        RETURN_ERR_MPP(ERR_F_HI_MPI_AF_Register, mpp_error_code); 
    }

    #if defined(HI_MPP_V1)
        ISP_IMAGE_ATTR_S stImageAttr;
        ISP_INPUT_TIMING_S stInputTiming;

        mpp_error_code = HI_MPI_ISP_Init();
        if (mpp_error_code != HI_SUCCESS) {
            RETURN_ERR_MPP(ERR_F_HI_MPI_ISP_Init, mpp_error_code);
        }

        stImageAttr.enBayer      = BAYER_RGGB;
        stImageAttr.u16FrameRate = in->fps;
        stImageAttr.u16Width     = in->width;
        stImageAttr.u16Height    = in->height;

        DO_OR_RETURN_MPP(HI_MPI_ISP_SetImageAttr, &stImageAttr);

        stInputTiming.enWndMode         = ISP_WIND_ALL;
        stInputTiming.u16HorWndStart    = 200;          //TODO
        stInputTiming.u16HorWndLength   = in->width;
        stInputTiming.u16VerWndStart    = 18;
        stInputTiming.u16VerWndLength   = in->height;   //TODO

        DO_OR_RETURN_MPP(HI_MPI_ISP_SetInputTiming, &stInputTiming);

    #elif defined(HI_MPP_V2) \
        || defined(HI_MPP_V3) \
        || defined(HI_MPP_V4)

        DO_OR_RETURN_MPP(HI_MPI_ISP_MemInit, 0);

        #if defined(HI_MPP_V2) \
            || defined(HI_MPP_V3)
            ISP_WDR_MODE_S stWdrMode;

            stWdrMode.enWDRMode  = in->wdr;

            DO_OR_RETURN_MPP(HI_MPI_ISP_SetWDRMode, 0, &stWdrMode);
        #endif

        ISP_PUB_ATTR_S stPubAttr;

        #if defined(HI3516AV200) \
            || defined(HI_MPP_V4) \

            stPubAttr.stSnsSize.u32Width    = in->width; 
            stPubAttr.stSnsSize.u32Height   = in->height; 
        #endif

        #if defined(HI_MPP_V4)
            stPubAttr.enWDRMode             = WDR_MODE_NONE;
            stPubAttr.u8SnsMode             = 0; //TODO
        #endif

        stPubAttr.enBayer               = in->bayer;
        stPubAttr.f32FrameRate          = in->fps;
        //Start position of the cropping window, image width, and image height
        stPubAttr.stWndRect.s32X        = 0;
        stPubAttr.stWndRect.s32Y        = 0;
        stPubAttr.stWndRect.u32Width    = in->width;
        stPubAttr.stWndRect.u32Height   = in->height;

        DO_OR_RETURN_MPP(HI_MPI_ISP_SetPubAttr, 0, &stPubAttr);

        DO_OR_RETURN_MPP(HI_MPI_ISP_Init, 0);
    #endif

    general_error_code = pthread_create(&mpp_isp_thread_pid, 0, (void* (*)(void*))mpp_isp_thread, NULL);
    if (general_error_code != 0) {
        GO_LOG_ISP(LOGGER_ERROR, "pthread_create");
        err->general = general_error_code;
        return ERR_GENERAL;
    }

    return ERR_NONE;
}
