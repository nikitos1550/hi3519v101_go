//+build hi3516cv300 hi3516av200

package vpss

/*
#include "../include/mpp_v3.h"

#include <string.h>

#define ERR_NONE                    0
#define ERR_MPP                     2
#define ERR_HI_MPI_VPSS_CreateGrp   3
#define ERR_HI_MPI_VPSS_StartGrp    4
#define ERR_HI_MPI_SYS_Bind         5

#define MAX_CHANNELS 10
VIDEO_FRAME_INFO_S channelFrames[MAX_CHANNELS];

typedef void (*callbackFunc) (unsigned int, VIDEO_FRAME_INFO_S*);

int mpp3_vpss_init(unsigned int *error_code) {
    *error_code = 0;

    VPSS_GRP_ATTR_S stVpssGrpAttr;

    stVpssGrpAttr.u32MaxW           = 3840;
    stVpssGrpAttr.u32MaxH           = 2160;
    stVpssGrpAttr.bIeEn             = HI_FALSE;
    stVpssGrpAttr.bNrEn             = HI_TRUE;
    stVpssGrpAttr.bHistEn           = HI_FALSE;
    stVpssGrpAttr.bDciEn            = HI_FALSE;
    stVpssGrpAttr.enDieMode         = VPSS_DIE_MODE_NODIE;
    stVpssGrpAttr.enPixFmt          = PIXEL_FORMAT_YUV_SEMIPLANAR_420;//SAMPLE_PIXEL_FORMAT;
    #ifdef HI3516AV200
    stVpssGrpAttr.bStitchBlendEn    = HI_FALSE;
    #endif

    #ifdef HI3516AV200
    stVpssGrpAttr.stNrAttr.enNrType                         = VPSS_NR_TYPE_VIDEO;
    stVpssGrpAttr.stNrAttr.stNrVideoAttr.enNrRefSource      = VPSS_NR_REF_FROM_RFR;//VPSS_NR_REF_FROM_CHN0, VPSS_NR_REF_FROM_SRC
    stVpssGrpAttr.stNrAttr.stNrVideoAttr.enNrOutputMode     = VPSS_NR_OUTPUT_NORMAL;//VPSS_NR_OUTPUT_DELAY NORMAL
    stVpssGrpAttr.stNrAttr.u32RefFrameNum                   = 2;
    #endif
    
    //    stVpssGrpAttr.u32MaxW = global_width;
    //    stVpssGrpAttr.u32MaxH = global_height;
    //    stVpssGrpAttr.bIeEn = HI_FALSE;
    //    stVpssGrpAttr.bNrEn = HI_TRUE;//HI_FALSE;//HI_TRUE;
    //    stVpssGrpAttr.bHistEn = HI_FALSE;
    //    stVpssGrpAttr.bSharpenEn = HI_FALSE;//HI_TRUE;
    //    stVpssGrpAttr.enDieMode = VPSS_DIE_MODE_NODIE;
    //    stVpssGrpAttr.enPixFmt = PIXEL_FORMAT_YUV_SEMIPLANAR_420;
    

    *error_code = HI_MPI_VPSS_CreateGrp(0, &stVpssGrpAttr);
    if (*error_code != HI_SUCCESS) return ERR_HI_MPI_VPSS_CreateGrp;

    *error_code = HI_MPI_VPSS_StartGrp(0);
    if (*error_code != HI_SUCCESS) return ERR_HI_MPI_VPSS_StartGrp;

    MPP_CHN_S stSrcChn;
    MPP_CHN_S stDestChn;

    stSrcChn.enModId  = HI_ID_VIU;
    stSrcChn.s32DevId = 0;
    stSrcChn.s32ChnId = 0;

    stDestChn.enModId  = HI_ID_VPSS;
    stDestChn.s32DevId = 0;
    stDestChn.s32ChnId = 0;

    *error_code = HI_MPI_SYS_Bind(&stSrcChn, &stDestChn);
    if (*error_code != HI_SUCCESS) return ERR_HI_MPI_SYS_Bind;

    return ERR_NONE;
}

int mpp3_vpss_sample_channel0(unsigned int *error_code) {
    *error_code = 0;

    VPSS_CHN_ATTR_S stVpssChnAttr;
    VPSS_CHN_MODE_S stVpssChnMode;

    stVpssChnMode.enChnMode      = VPSS_CHN_MODE_USER;
    stVpssChnMode.bDouble        = HI_FALSE;
    stVpssChnMode.enPixelFormat  = PIXEL_FORMAT_YUV_SEMIPLANAR_420;
    stVpssChnMode.u32Width       = 3840;
    stVpssChnMode.u32Height      = 2160;
    stVpssChnMode.enCompressMode = COMPRESS_MODE_NONE; //COMPRESS_MODE_SEG;

    memset(&stVpssChnAttr, 0, sizeof(stVpssChnAttr));

    stVpssChnAttr.s32SrcFrameRate = 30;
    stVpssChnAttr.s32DstFrameRate = 30;

    *error_code = HI_MPI_VPSS_SetChnAttr(0, 0, &stVpssChnAttr);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    *error_code = HI_MPI_VPSS_SetChnMode(0, 0, &stVpssChnMode);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    *error_code = HI_MPI_VPSS_EnableChn(0, 0);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    return ERR_NONE;
}

int mpp3_vpss_sample_channel(
        unsigned int channelId,
        unsigned int width,
        unsigned int height,
        unsigned int fps,
        unsigned int *error_code) {
    *error_code = 0;

    VPSS_CHN_ATTR_S stVpssChnAttr;
    VPSS_CHN_MODE_S stVpssChnMode;

    stVpssChnMode.enChnMode      = VPSS_CHN_MODE_USER;
    stVpssChnMode.bDouble        = HI_FALSE;
    stVpssChnMode.enPixelFormat  = PIXEL_FORMAT_YUV_SEMIPLANAR_420;
    stVpssChnMode.u32Width       = width;
    stVpssChnMode.u32Height      = height;
    stVpssChnMode.enCompressMode = COMPRESS_MODE_NONE; //COMPRESS_MODE_SEG;

    memset(&stVpssChnAttr, 0, sizeof(stVpssChnAttr));

    stVpssChnAttr.s32SrcFrameRate = 30;
    stVpssChnAttr.s32DstFrameRate = fps;

    *error_code = HI_MPI_VPSS_SetChnAttr(0, channelId, &stVpssChnAttr);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    *error_code = HI_MPI_VPSS_SetChnMode(0, channelId, &stVpssChnMode);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

 	HI_U32 u32Depth = 1;
 	*error_code = HI_MPI_VPSS_SetDepth(0, channelId, u32Depth);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    *error_code = HI_MPI_VPSS_EnableChn(0, channelId);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    return ERR_NONE;
}

int mpp3_destroy_vpss_sample_channel(unsigned int channelId, unsigned int *error_code) {
    *error_code = 0;
    *error_code = HI_MPI_VPSS_DisableChn(0, channelId);
    if (*error_code != HI_SUCCESS) return ERR_MPP;

    return ERR_NONE;
}

int mpp3_receive_frame(unsigned int channelId) {
 	int s32Ret = HI_MPI_VPSS_GetChnFrame(0, channelId, &channelFrames[channelId], -1);

 	if (HI_SUCCESS != s32Ret) {
 		printf("HI_MPI_VPSS_GetChnFrame failed with %#x!\n", s32Ret);
 	}

 	return s32Ret;
}

int mpp3_release_frame(unsigned int channelId) {
 	int s32Ret = HI_MPI_VPSS_ReleaseChnFrame(0, channelId, &channelFrames[channelId]);

 	if (HI_SUCCESS != s32Ret) {
 		printf("HI_MPI_VPSS_ReleaseChnFrame failed with %#x!\n", s32Ret);
 	}

 	return s32Ret;
}

void mpp3_send_frame_to_clients(unsigned int channelId, unsigned int processingId, void* callback) {
	callbackFunc func = callback;
	func(processingId, &channelFrames[channelId]);
}

*/
import "C"

import (
    "log"
    "application/pkg/mpp/error"
)

func Init() {
    var errorCode C.uint

    switch err := C.mpp3_vpss_init(&errorCode); err {
    case C.ERR_NONE:
        log.Println("C.mpp3_vpss_init() ok")
    case C.ERR_HI_MPI_VPSS_CreateGrp:
        log.Fatal("C.mpp3_vpss_init() HI_MPI_VPSS_CreateGrp() error ", error.Resolve(int64(errorCode)))
    case C.ERR_HI_MPI_VPSS_StartGrp:
        log.Fatal("C.mpp3_vpss_init() HI_MPI_VPSS_StartGrp() error ", error.Resolve(int64(errorCode)))
    case C.ERR_HI_MPI_SYS_Bind:
        log.Fatal("C.mpp3_vpss_init() HI_MPI_SYS_Bind() error ", error.Resolve(int64(errorCode)))
    default:
        log.Fatal("Unexpected return ", err , " of C.mpp3_vpss_init()")
    }
}

func SampleChannel0() {
    var errorCode C.uint

    switch err := C.mpp3_vpss_sample_channel0(&errorCode); err {
    case C.ERR_NONE:
        log.Println("C.mpp3_vpss_sample_channel0() ok")
    case C.ERR_MPP:
        log.Fatal("C.mpp3_vpss_sample_channel0() MPP error ", error.Resolve(int64(errorCode)))
    default:
        log.Fatal("Unexpected return ", err , " of C.mpp3_vpss_sample_channel0()")
    }
}

func CreateChannel(channel Channel) {
    var errorCode C.uint

    switch err := C.mpp3_vpss_sample_channel(C.uint(channel.ChannelId), C.uint(channel.Width), C.uint(channel.Height), C.uint(channel.Fps), &errorCode); err {
    case C.ERR_NONE:
        log.Println("C.mpp3_vpss_sample_channel() ok")
    case C.ERR_MPP:
        log.Fatal("C.mpp3_vpss_sample_channel() MPP error ", error.Resolve(int64(errorCode)))
    default:
        log.Fatal("Unexpected return ", err , " of C.mpp3_vpss_sample_channel()")
    }

    go func() {
		sendDataToClients(channel)
    }()
}

func DestroyChannel(channel Channel) {
    var errorCode C.uint

    switch err := C.mpp3_destroy_vpss_sample_channel(C.uint(channel.ChannelId), &errorCode); err {
    case C.ERR_NONE:
        log.Println("C.mpp3_destroy_vpss_sample_channel() ok")
    case C.ERR_MPP:
        log.Fatal("C.mpp3_destroy_vpss_sample_channel() MPP error ", error.Resolve(int64(errorCode)))
    default:
        log.Fatal("Unexpected return ", err , " of C.mpp3_destroy_vpss_sample_channel()")
    }
}

func sendDataToClients(channel Channel) {
	for{
		if (!channel.Started){
			break
		}

		err := C.mpp3_receive_frame(C.uint(channel.ChannelId));
		if (err != 0){
			log.Println("Failed receive frame", channel.ChannelId, error.Resolve(int64(err)))
			continue
		}

		for processingId, callback := range channel.Clients {
			C.mpp3_send_frame_to_clients(C.uint(channel.ChannelId), C.uint(processingId), callback);
		}

		err = C.mpp3_release_frame(C.uint(channel.ChannelId));
		if (err != 0){
			log.Println("Failed release frame", channel.ChannelId, error.Resolve(int64(err)))
		}
	}
}