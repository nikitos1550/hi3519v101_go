#include "isp.h"

static pthread_t mpp_isp_thread_pid;

void* mpp_isp_thread(HI_VOID *param){   //now we start it from go space
    GO_LOG_ISP(LOGGER_TRACE, "HI_MPI_ISP_Run");
    #if HI_MPP == 1
        HI_MPI_ISP_Run();
    #elif HI_MPP >= 2
        HI_MPI_ISP_Run(0);
    #endif
    GO_LOG_ISP(LOGGER_ERROR, "HI_MPI_ISP_Run failed");

    return NULL;
}

static inline int64_t mpp_isp_register_lib_ae(char * lib) {
    ALG_LIB_S stLib;

    strcpy(stLib.acLibName, lib);
    stLib.s32Id = 0;

    //printf("%s\n", stLib.acLibName);

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

    //printf("%s\n", stLib.acLibName);

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

    //printf("%s\n", stLib.acLibName);

    #if HI_MPP == 1
        return HI_MPI_AF_Register(&stLib);
    #elif HI_MPP == 4
        return HI_SUCCESS;
    #elif HI_MPP >=2
        return HI_MPI_AF_Register(0, &stLib);
    #endif
} 

int mpp_isp_init(error_in *err, mpp_isp_init_in *in) { 
    //int general_error_code = 0;

    DO_OR_RETURN_ERR_MPP(err, mpp_isp_register_lib_ae, HI_AE_LIB_NAME);

    DO_OR_RETURN_ERR_MPP(err, mpp_isp_register_lib_awb, HI_AWB_LIB_NAME);

    DO_OR_RETURN_ERR_MPP(err, mpp_isp_register_lib_af, HI_AF_LIB_NAME);

    #if HI_MPP == 1
        DO_OR_RETURN_ERR_MPP(err, HI_MPI_ISP_Init); //TODO check is it possible to move func call after HI_MPI_ISP_SetInputTiming,

        ISP_IMAGE_ATTR_S stImageAttr;

        memset(&stImageAttr, 0, sizeof(stImageAttr));

        stImageAttr.enBayer      = in->bayer;
        stImageAttr.u16FrameRate = in->fps;
        stImageAttr.u16Width     = in->width;
        stImageAttr.u16Height    = in->height;

        DO_OR_RETURN_ERR_MPP(err, HI_MPI_ISP_SetImageAttr, &stImageAttr);

        ISP_INPUT_TIMING_S stInputTiming;

        memset(&stInputTiming, 0, sizeof(stInputTiming));

        stInputTiming.enWndMode         = ISP_WIND_ALL;
        stInputTiming.u16HorWndStart    = in->isp_crop_x0; 
        stInputTiming.u16VerWndStart    = in->isp_crop_y0; 
        stInputTiming.u16HorWndLength   = in->isp_crop_width;
        stInputTiming.u16VerWndLength   = in->isp_crop_height;

        DO_OR_RETURN_ERR_MPP(err, HI_MPI_ISP_SetInputTiming, &stInputTiming);

    #elif HI_MPP >= 2

        DO_OR_RETURN_ERR_MPP(err, HI_MPI_ISP_MemInit, 0);

        #if HI_MPP <= 3
            ISP_WDR_MODE_S stWdrMode;

            memset(&stWdrMode, 0, sizeof(stWdrMode));

            stWdrMode.enWDRMode         = in->wdr;
        
            DO_OR_RETURN_ERR_MPP(err, HI_MPI_ISP_SetWDRMode, 0, &stWdrMode);
        #endif


        ISP_PUB_ATTR_S stPubAttr;

        memset(&stPubAttr, 0, sizeof(stPubAttr));

        //#if defined(HI3516AV200) || HI_MPP == 4
        //    stPubAttr.stSnsSize.u32Width    = in->isp_crop_width; 
        //    stPubAttr.stSnsSize.u32Height   = in->isp_crop_height; 
        //#endif

        //#if HI_MPP == 4
        //    //Selecting the initialization sequence of the sensor. When the
        //    //resolution and frame rates of the two sequences are the same,
        //    //different u8SnsMode values map different initialization sequences.
        //    //In other cases, u8SnsMode is set to 0 by default, and the
        //    //initialization sequence can be selected based on stSnsSize and
        //    //f32FrameRate.
        //    stPubAttr.u8SnsMode             = 0;
        //#endif

        stPubAttr.enBayer               = in->bayer;
        stPubAttr.f32FrameRate          = in->fps;
        //Start position of the cropping window, image width, and image height
        stPubAttr.stWndRect.s32X        = in->isp_crop_x0;
        stPubAttr.stWndRect.s32Y        = in->isp_crop_y0;
        stPubAttr.stWndRect.u32Width    = in->isp_crop_width;
        stPubAttr.stWndRect.u32Height   = in->isp_crop_height;

        #if defined(HI3516AV200) || HI_MPP == 4
            stPubAttr.stSnsSize.u32Width    = in->isp_crop_width; 
            stPubAttr.stSnsSize.u32Height   = in->isp_crop_height; 
        #endif

        #if HI_MPP == 4
            //Selecting the initialization sequence of the sensor. When the
            //resolution and frame rates of the two sequences are the same,
            //different u8SnsMode values map different initialization sequences.
            //In other cases, u8SnsMode is set to 0 by default, and the
            //initialization sequence can be selected based on stSnsSize and
            //f32FrameRate.
            stPubAttr.u8SnsMode             = 0;
            stPubAttr.enWDRMode         = in->wdr;
        #endif

        DO_OR_RETURN_ERR_MPP(err, HI_MPI_ISP_SetPubAttr, 0, &stPubAttr);

        DO_OR_RETURN_ERR_MPP(err, HI_MPI_ISP_Init, 0);
    #endif

    //thread start moved to go space, tmp
    DO_OR_RETURN_ERR_GENERAL(err, pthread_create, &mpp_isp_thread_pid, 0, (void* (*)(void*))mpp_isp_thread, NULL);
	DO_OR_RETURN_ERR_GENERAL(err, pthread_setname_np, mpp_isp_thread_pid, "ISP");

    return ERR_NONE;
}
