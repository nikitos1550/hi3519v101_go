#include "../hisi_external.h"
#include "../hisi_utils.h"
#include "hi3516av200_mpp.h"

#include <string.h>

int hisi_channels_max_num(unsigned int * num) {
    *num = VPSS_MAX_PHY_CHN_NUM;
    return ERR_NONE;
}

int hisi_channel_info(unsigned int id, struct channel_params * chn) {
    int error_code = 0;

    if (id < 0 &&
        id >= VPSS_MAX_PHY_CHN_NUM) return ERR_OBJECT_NOT_FOUNT;

    VPSS_CHN_ATTR_S stVpssChnAttr;
    VPSS_CHN_MODE_S stVpssChnMode;

    error_code = HI_MPI_VPSS_GetChnAttr(0, id, &stVpssChnAttr);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: hisi_channel_info: HI_MPI_VPSS_GetChnAttr %d failed %#x\n", id, error_code);
        return ERR_MPP;
    }

    error_code = HI_MPI_VPSS_GetChnMode(0, id, &stVpssChnMode);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: hisi_channel_info: HI_MPI_VPSS_GetChnMode %d failed %#x\n", id, error_code);
        return ERR_MPP;
    }

    chn->width  = stVpssChnMode.u32Width;
    chn->height = stVpssChnMode.u32Height;
    chn->fps    = stVpssChnAttr.s32DstFrameRate;

    return ERR_NONE;
}

int hisi_channel_enable(unsigned int id, struct channel_params * chn) {
    int error_code = 0;

    if (id < 0 &&
        id >= VPSS_MAX_PHY_CHN_NUM) return ERR_OBJECT_NOT_FOUNT;

    if (chn->width < VPSS_MIN_IMAGE_WIDTH &&
        chn->width > 3840) { //TODO max limit from VI
        chn->width = PARAM_BAD_VALUE;
        error_code++;
    }
    if (chn->height < VPSS_MIN_IMAGE_HEIGHT &&
        chn->height > 3840) { //TODO max limit from VI
        chn->height = PARAM_BAD_VALUE;
        error_code++;
    }
    if (chn->fps < 1 &&
        chn->fps > 30) { //TODO max limit from VI
        chn->fps = PARAM_BAD_VALUE;
        error_code++;
    }
    if (error_code > 0) return ERR_BAD_PARAMS;

    VPSS_CHN_ATTR_S stVpssChnAttr;
    VPSS_CHN_MODE_S stVpssChnMode;

    stVpssChnMode.enChnMode      = VPSS_CHN_MODE_USER;
    stVpssChnMode.bDouble        = HI_FALSE;
    stVpssChnMode.enPixelFormat  = PIXEL_FORMAT_YUV_SEMIPLANAR_420;
    stVpssChnMode.u32Width       = chn->width;
    stVpssChnMode.u32Height      = chn->height;
    stVpssChnMode.enCompressMode = COMPRESS_MODE_NONE; //COMPRESS_MODE_SEG;

    memset(&stVpssChnAttr, 0, sizeof(stVpssChnAttr));
    stVpssChnAttr.s32SrcFrameRate = 30; //TODO should be got from VI
    stVpssChnAttr.s32DstFrameRate = chn->fps;

    error_code = HI_MPI_VPSS_SetChnAttr(0, id, &stVpssChnAttr);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: hisi_channel_enable: HI_MPI_VPSS_SetChnAttr %d failed %#x\n", id, error_code);
        return ERR_MPP;
    }

    error_code = HI_MPI_VPSS_SetChnMode(0, id, &stVpssChnMode);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: hisi_channel_enable: HI_MPI_VPSS_SetChnMode %d failed %#x\n", id, error_code);
        return ERR_MPP;
    }

    error_code = HI_MPI_VPSS_EnableChn(0, id);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: hisi_channel_enable: HI_MPI_VPSS_EnableChn %d failed %#x\n", id, error_code);
        return ERR_MPP;
    }

    return ERR_NONE;
}

int hisi_channel_disable(unsigned int id) {
    int error_code = 0;

    if (id < 0 &&
        id >= VPSS_MAX_PHY_CHN_NUM) return ERR_OBJECT_NOT_FOUNT;

    error_code = HI_MPI_VPSS_DisableChn(0, id);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: hisi_channel_disable: HI_MPI_VPSS_DisableChn %d failed %#x\n", id, error_code);
        return ERR_MPP;
    }

    return ERR_NONE;
}

