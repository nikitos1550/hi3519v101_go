//+build hi3516av200 hi3516cv300

package vi

/*
int hi3516av200_vi_init(struct hi3516av200_cmos * c, struct capture_params * cp) {
    int error_code;

    VI_DEV_ATTR_S  stViDevAttr;

    memset(&stViDevAttr, 0, sizeof(stViDevAttr));
    memcpy(&stViDevAttr, c->videv, sizeof(stViDevAttr));

    //stViDevAttr.stDevRect.s32X                              = 0;
    //stViDevAttr.stDevRect.s32Y                              = 0;
    stViDevAttr.stDevRect.u32Width                          = c->width;
    stViDevAttr.stDevRect.u32Height                         = c->height;
    stViDevAttr.stBasAttr.stSacleAttr.stBasSize.u32Width    = c->width;
    stViDevAttr.stBasAttr.stSacleAttr.stBasSize.u32Height   = c->height;
    stViDevAttr.stBasAttr.stSacleAttr.bCompress             = HI_FALSE;

    error_code = HI_MPI_VI_SetDevAttr(0, &stViDevAttr);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: HI_MPI_VI_SetDevAttr failed with %#x!\n", error_code);
        return ERR_MPP;
    }

    error_code = HI_MPI_VI_EnableDev(0);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: HI_MPI_VI_EnableDev failed with %#x!\n", error_code);
        return ERR_MPP;
    }

    RECT_S stCapRect;
    SIZE_S stTargetSize;

    stCapRect.s32X          = cp->x0;
    stCapRect.s32Y          = cp->y0;
    stCapRect.u32Width      = cp->width;
    stCapRect.u32Height     = cp->height;
    stTargetSize.u32Width   = stCapRect.u32Width;
    stTargetSize.u32Height  = stCapRect.u32Height;

    VI_CHN_ATTR_S stChnAttr;

    memcpy(&stChnAttr.stCapRect, &stCapRect, sizeof(RECT_S));

    stChnAttr.enCapSel              = VI_CAPSEL_BOTH;
    stChnAttr.stDestSize.u32Width   = stTargetSize.u32Width ;
    stChnAttr.stDestSize.u32Height  = stTargetSize.u32Height ;
    stChnAttr.enPixFormat           = PIXEL_FORMAT_YUV_SEMIPLANAR_420;   // sp420 or sp422

    stChnAttr.bMirror               = HI_FALSE;
    stChnAttr.bFlip                 = HI_FALSE;

    stChnAttr.s32SrcFrameRate       = c->fps;
    stChnAttr.s32DstFrameRate       = cp->fps;
    stChnAttr.enCompressMode        = COMPRESS_MODE_NONE;

    error_code = HI_MPI_VI_SetChnAttr(0, &stChnAttr);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: HI_MPI_VI_SetChnAttr failed with %#x!\n", error_code);
        return ERR_MPP;
    }

    error_code = HI_MPI_VI_EnableChn(0);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: HI_MPI_VI_EnableChn failed with %#x!\n", error_code);
        return ERR_MPP;
    }

    return ERR_NONE;
}
*/
//import "C"
