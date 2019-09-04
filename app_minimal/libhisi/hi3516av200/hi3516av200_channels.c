#include "../hisi_external.h"
#include "../hisi_utils.h"
#include "hi3516av200_mpp.h"

#include "hi3516av200_cmos.h"

#include "hi3516av200_channels.h"

#include <string.h>

int channels_enable[VPSS_MAX_PHY_CHN_NUM];
struct vpss_setup vpss;

int hisi_channels_max_num(unsigned int * num) {
    *num = VPSS_MAX_PHY_CHN_NUM;
    return ERR_NONE;
}

int hisi_channels_min_width(unsigned int * w) {
    *w = VPSS_MIN_IMAGE_WIDTH;
    return ERR_NONE;
}

int hisi_channels_max_width(unsigned int * w) {
    *w = vpss.width;
    return ERR_NONE;
}

int hisi_channels_min_height(unsigned int * h) {
    *h = VPSS_MIN_IMAGE_HEIGHT;
    return ERR_NONE;
}

int hisi_channels_max_height(unsigned int * h) {
    *h = vpss.height;
    return ERR_NONE;
}

int hisi_channels_max_fps(unsigned int * fps) {
    *fps = vpss.fps;
    return ERR_NONE;
}

int hisi_channel_fetch(unsigned int id, struct channel_params * chn) {
    int error_code = 0;

    if (chn == NULL) return ERR_GENERAL;

    if (id >= VPSS_MAX_PHY_CHN_NUM) return ERR_OBJECT_NOT_FOUND;

    if (channels_enable[id] != CHANNEL_ENABLED) return ERR_DISABLED;

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

    if (chn == NULL) return ERR_GENERAL;

    if (id >= VPSS_MAX_PHY_CHN_NUM) return ERR_OBJECT_NOT_FOUND;

    if (channels_enable[id] != CHANNEL_DISABLED) return ERR_NOT_ALLOWED;

    if (chn->width < VPSS_MIN_IMAGE_WIDTH ||
        chn->width > vpss.width) {
        chn->width = PARAM_BAD_VALUE;
        error_code++;
    }
    if (chn->height < VPSS_MIN_IMAGE_HEIGHT ||
        chn->height > vpss.height) {
        chn->height = PARAM_BAD_VALUE;
        error_code++;
    }
    if (chn->fps < 1 ||
        chn->fps > vpss.fps) {
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
    stVpssChnAttr.s32SrcFrameRate = vpss.fps;
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

    /*  //TODO Channel crop
        VPSS_CROP_INFO_S CropInfo;
        CropInfo.bEnable = 1;
        CropInfo.enCropCoordinate = VPSS_CROP_ABS_COOR;
        CropInfo.stCropRect.s32X = 0;
        CropInfo.stCropRect.s32Y = 0;
        CropInfo.stCropRect.u32Width = 640;
        CropInfo.stCropRect.u32Height = 480;

        error_code = HI_MPI_VPSS_SetChnCrop(0, id, &CropInfo);
        if (error_code != HI_SUCCESS) {
            printf("HI_MPI_VPSS_SetChnCrop failed with %#x\n", error_code);
            return -1;
        }
    */

    error_code = HI_MPI_VPSS_EnableChn(0, id);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: hisi_channel_enable: HI_MPI_VPSS_EnableChn %d failed %#x\n", id, error_code);
        return ERR_MPP;
    }

    channels_enable[id] = CHANNEL_ENABLED;

    return ERR_NONE;
}

int hisi_channel_disable(unsigned int id) {
    int error_code = 0;

    if (id >= VPSS_MAX_PHY_CHN_NUM) return ERR_OBJECT_NOT_FOUND;

    if (channels_enable[id] != CHANNEL_ENABLED) return ERR_NOT_ALLOWED;

    error_code = HI_MPI_VPSS_DisableChn(0, id);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: hisi_channel_disable: HI_MPI_VPSS_DisableChn %d failed %#x\n", id, error_code);
        return ERR_MPP;
    }

    channels_enable[id] = CHANNEL_DISABLED;

    return ERR_NONE;
}

