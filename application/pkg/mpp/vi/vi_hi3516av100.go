//+build arm
//+build hi3516av100

package vi

/*
int vi_init(void){
        HI_S32                  ret;

   HI_S32 s32Ret;
    HI_S32 s32IspDev = 0;
    ISP_WDR_MODE_S stWdrMode;
    VI_DEV_ATTR_S  stViDevAttr;

    memset(&stViDevAttr,0,sizeof(stViDevAttr));


     //   case SONY_IMX178_LVDS_5M_30FPS:
            memcpy(&stViDevAttr,&DEV_ATTR_LVDS_BASE,sizeof(stViDevAttr));
            stViDevAttr.stDevRect.s32X = 0;
            stViDevAttr.stDevRect.s32Y = 0;
            stViDevAttr.stDevRect.u32Width  = 2592;
            stViDevAttr.stDevRect.u32Height = 1944;


    s32Ret = HI_MPI_VI_SetDevAttr(0, &stViDevAttr);
    if (s32Ret != HI_SUCCESS)
    {
        printf("HI_MPI_VI_SetDevAttr failed with %#x!\n", s32Ret);
        return HI_FAILURE;
    }

        s32Ret = HI_MPI_VI_EnableDev(0);
    if (s32Ret != HI_SUCCESS)
    {
        printf("HI_MPI_VI_EnableDev failed with %#x!\n", s32Ret);
        return HI_FAILURE;
    }

    RECT_S stCapRect;
    SIZE_S stTargetSize;

     stCapRect.s32X = 0;
        stCapRect.s32Y = 0;
                stCapRect.u32Width  = 2560;
                stCapRect.u32Height = 1440;
       stTargetSize.u32Width = stCapRect.u32Width;
        stTargetSize.u32Height = stCapRect.u32Height;





    VI_CHN_ATTR_S stChnAttr;

memcpy(&stChnAttr.stCapRect, &stCapRect, sizeof(RECT_S));
    stChnAttr.enCapSel = VI_CAPSEL_BOTH;
    stChnAttr.stDestSize.u32Width = stTargetSize.u32Width ;
    stChnAttr.stDestSize.u32Height =  stTargetSize.u32Height ;
    stChnAttr.enPixFormat = PIXEL_FORMAT_YUV_SEMIPLANAR_420;   // sp420 or sp422

    stChnAttr.bMirror = HI_FALSE;
    stChnAttr.bFlip = HI_FALSE;

    stChnAttr.s32SrcFrameRate = 25;
    stChnAttr.s32DstFrameRate = 25;
    stChnAttr.enCompressMode = COMPRESS_MODE_NONE;

    s32Ret = HI_MPI_VI_SetChnAttr(0, &stChnAttr);
    if (s32Ret != HI_SUCCESS)
    {
        printf("HI_MPI_VI_SetChnAttr failed with %#x!\n", s32Ret);
        return HI_FAILURE;
    }

    s32Ret = HI_MPI_VI_EnableChn(0);
    if (s32Ret != HI_SUCCESS)
    {
        printf(" HI_MPI_VI_EnableChn failed with %#x!\n", s32Ret);
        return HI_FAILURE;
    }


        return 0;
}


*/

func Init() {
	var errorCode C.uint

	switch err := C.mpp2_vi_init(&errorCode); err {
	case C.ERR_NONE:
		log.Println("C.mpp2_vi_init() ok")
	case C.ERR_HI_MPI_VI_SetDevAttr:
		log.Fatal("C.mpp2_vi_init() ERR_HI_MPI_VI_SetDevAttr() error ", error.Resolve(int64(errorCode)))
	case C.ERR_HI_MPI_VI_EnableDev:
		log.Fatal("C.mpp2_vi_init() ERR_HI_MPI_VI_EnableDev() error ", error.Resolve(int64(errorCode)))
	case C.ERR_HI_MPI_VI_SetChnAttr:
		log.Fatal("C.mpp2_vi_init() ERR_HI_MPI_VI_SetChnAttr() error ", error.Resolve(int64(errorCode)))
	case C.ERR_HI_MPI_VI_EnableChn:
		log.Fatal("C.mpp2_vi_init() ERR_HI_MPI_VI_EnableChn() error ", error.Resolve(int64(errorCode)))
	default:
		log.Fatal("Unexpected return ", err, " of C.mpp2_vi_init()")
	}
}
