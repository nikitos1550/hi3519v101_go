//+build arm
//+build hi3516av200

package vpss

/*
#include "../include/mpp.h"
#include "../errmpp/error.h"
#include "../../logger/logger.h"

#include <stdint.h>
#include <string.h>

#define MAX_CHANNELS VPSS_MAX_PHY_CHN_NUM
VIDEO_FRAME_INFO_S channelFrames[MAX_CHANNELS];

typedef void (*callbackFunc) (unsigned int, VIDEO_FRAME_INFO_S*);

typedef struct hi3516av200_vpss_init_in_struct {
    unsigned int width;
    unsigned int height;
    unsigned char nr;
    unsigned char nr_frames;
} hi3516av200_vpss_init_in;

static int hi3516av200_vpss_init(error_in *err, hi3516av200_vpss_init_in *in) {
    unsigned int mpp_error_code = 0;

    VPSS_GRP_ATTR_S stVpssGrpAttr;

    stVpssGrpAttr.u32MaxW           = in->width;
    stVpssGrpAttr.u32MaxH           = in->height;
    stVpssGrpAttr.bIeEn             = HI_FALSE;                         //reserved
    stVpssGrpAttr.bHistEn           = HI_FALSE;                         //reserved
    stVpssGrpAttr.bDciEn            = HI_FALSE;                         //reserved
    stVpssGrpAttr.enDieMode         = VPSS_DIE_MODE_NODIE;              //reserved
    stVpssGrpAttr.enPixFmt          = PIXEL_FORMAT_YUV_SEMIPLANAR_420;  //yuv420 or yuv422
    stVpssGrpAttr.bStitchBlendEn    = HI_FALSE;                         
    
    if (in->nr == 1) {
        GO_LOG_VPSS(LOGGER_TRACE, "VPSS NR on")
        stVpssGrpAttr.bNrEn = HI_TRUE;
    } else {
        GO_LOG_VPSS(LOGGER_TRACE, "VPSS NR off")
        stVpssGrpAttr.bNrEn = HI_FALSE;
    }

    stVpssGrpAttr.stNrAttr.enNrType                         = VPSS_NR_TYPE_VIDEO;       //video or snapshot, we use video (i don`t know anything about snapshot mode)

    //VPSS_NR_REF_FROM_RFR Reconstruction frame as the reference frame
    //VPSS_NR_REF_FROM_CHN0 Channel 0 output as the reference frame
    //VPSS_NR_REF_FROM_SRC Input source picture as the reference frame
    stVpssGrpAttr.stNrAttr.stNrVideoAttr.enNrRefSource      = VPSS_NR_REF_FROM_RFR;

    //VPSS_NR_OUTPUT_NORMAL Normal output mode. There is no delay.
    //VPSS_NR_OUTPUT_DELAY Delay output mode. The output is one-frame delayed.
    //Only Hi3519 V101 supports this data structure. TODO what about hi3516av200 chip?
    //In delay output mode, the number of reference frames cannot be set to 1, the single component is not supported, and the large stream cannot be used as the reference frame.

    stVpssGrpAttr.stNrAttr.stNrVideoAttr.enNrOutputMode     = VPSS_NR_OUTPUT_NORMAL;
    
    //1 or 2 for video mode
    stVpssGrpAttr.stNrAttr.u32RefFrameNum                   = in->nr_frames; //2;
    
    mpp_error_code = HI_MPI_VPSS_CreateGrp(0, &stVpssGrpAttr);
    if (mpp_error_code != HI_SUCCESS) {
        GO_LOG_VPSS(LOGGER_ERROR, "HI_MPI_VPSS_CreateGrp")
        err->mpp = mpp_error_code;
        return ERR_MPP;
    }

    mpp_error_code = HI_MPI_VPSS_StartGrp(0);
    if (mpp_error_code != HI_SUCCESS) {
        GO_LOG_VPSS(LOGGER_ERROR, "HI_MPI_VPSS_StartGrp")   
        err->mpp = mpp_error_code;
        return ERR_MPP;
    }

    MPP_CHN_S stSrcChn;
    MPP_CHN_S stDestChn;

    stSrcChn.enModId  = HI_ID_VIU;
    stSrcChn.s32DevId = 0;
    stSrcChn.s32ChnId = 0;

    stDestChn.enModId  = HI_ID_VPSS;
    stDestChn.s32DevId = 0;
    stDestChn.s32ChnId = 0;

    mpp_error_code = HI_MPI_SYS_Bind(&stSrcChn, &stDestChn);
    if (mpp_error_code != HI_SUCCESS) {
        GO_LOG_VPSS(LOGGER_ERROR, "HI_MPI_SYS_Bind")   
        err->mpp = mpp_error_code;
        return ERR_MPP;
    }

    return ERR_NONE;
}

typedef struct hi3516av200_vpss_create_channel_in_struct {
    unsigned int channel_id;
    unsigned int width;
    unsigned int height;
    unsigned int vi_fps;
    unsigned int fps;
} hi3516av200_vpss_create_channel_in;

static int hi3516av200_vpss_create_channel(error_in *err, hi3516av200_vpss_create_channel_in * in) {
    unsigned int mpp_error_code = 0;

    VPSS_CHN_MODE_S stVpssChnMode;

    stVpssChnMode.enChnMode      = VPSS_CHN_MODE_USER;
    stVpssChnMode.bDouble        = HI_FALSE;
    stVpssChnMode.enPixelFormat  = PIXEL_FORMAT_YUV_SEMIPLANAR_420;
    stVpssChnMode.u32Width       = in->width;
    stVpssChnMode.u32Height      = in->height;
    stVpssChnMode.enCompressMode = COMPRESS_MODE_NONE; //COMPRESS_MODE_SEG;

    VPSS_CHN_ATTR_S stVpssChnAttr;

    memset(&stVpssChnAttr, 0, sizeof(stVpssChnAttr));

    //typedef struct hiVPSS_CHN_ATTR_S
    //{
    //    HI_BOOL bSpEn;            //SP enable. It must be set to HI_FALSE.
    //    HI_BOOL bBorderEn;        //Border enable. It must be set to HI_FALSE.
    //    HI_BOOL bMirror;
    //    HI_BOOL bFlip;
    //    HI_S32 s32SrcFrameRate;
    //    HI_S32 s32DstFrameRate;
    //    BORDER_S
    //    stBorder;
    //}VPSS_CHN_ATTR_S;

    stVpssChnAttr.s32SrcFrameRate = in->vi_fps;
    stVpssChnAttr.s32DstFrameRate = in->fps;

    mpp_error_code = (uint64_t)HI_MPI_VPSS_SetChnAttr(0, in->channel_id, &stVpssChnAttr);
    if (mpp_error_code != HI_SUCCESS) {
        GO_LOG_VPSS(LOGGER_ERROR, "HI_MPI_VPSS_SetChnAttr")   
        err->mpp = mpp_error_code;
        return ERR_MPP;
    }

    mpp_error_code = HI_MPI_VPSS_SetChnMode(0, in->channel_id, &stVpssChnMode);
    if (mpp_error_code != HI_SUCCESS) {
        GO_LOG_VPSS(LOGGER_ERROR, "HI_MPI_VPSS_SetChnMode")   
        err->mpp = mpp_error_code;
        return ERR_MPP;
    }

 	HI_U32 u32Depth = 1; //TODO
 	mpp_error_code = HI_MPI_VPSS_SetDepth(0, in->channel_id, u32Depth);
    if (mpp_error_code != HI_SUCCESS) {
        GO_LOG_VPSS(LOGGER_ERROR, "HI_MPI_VPSS_SetDepth")
        err->mpp = mpp_error_code;
        return ERR_MPP;
    }

    mpp_error_code = HI_MPI_VPSS_EnableChn(0, in->channel_id);
    if (mpp_error_code != HI_SUCCESS) {
        GO_LOG_VPSS(LOGGER_ERROR, "HI_MPI_VPSS_EnableChn")   
        err->mpp = mpp_error_code;
        return ERR_MPP;
    }

    return ERR_NONE;
}

typedef struct hi3516av200_vpss_destroy_channel_in_struct {
    unsigned int channel_id;
} hi3516av200_vpss_destroy_channel_in;

static int hi3516av200_vpss_destroy_channel(error_in * err, hi3516av200_vpss_destroy_channel_in *in) {
    unsigned int mpp_error_code = 0;

    mpp_error_code = HI_MPI_VPSS_DisableChn(0, in->channel_id);
    if (mpp_error_code != HI_SUCCESS) {
        GO_LOG_VPSS(LOGGER_ERROR, "HI_MPI_VPSS_DisableChn")   
        err->mpp = mpp_error_code;
        return ERR_MPP;
    }

    return ERR_NONE;
}

typedef struct hi3516av200_receive_frame_out_struct {

} hi3516av200_receive_frame_out;

static int hi3516av200_receive_frame(error_in *err, unsigned int channel_id) {
    unsigned int mpp_error_code;

    mpp_error_code = HI_MPI_VPSS_GetChnFrame(0, channel_id, &channelFrames[channel_id], -1); //blocking mode call

 	if (mpp_error_code != HI_SUCCESS) {
        GO_LOG_VPSS(LOGGER_ERROR, "HI_MPI_VPSS_GetChnFrame")
        err->mpp = mpp_error_code;
        return ERR_MPP;
 	}

 	return ERR_NONE;
}

static int hi3516av200_release_frame(error_in *err, unsigned int channel_id) {
    unsigned int mpp_error_code;

 	mpp_error_code = HI_MPI_VPSS_ReleaseChnFrame(0, channel_id, &channelFrames[channel_id]);

 	if (mpp_error_code != HI_SUCCESS) {
        GO_LOG_VPSS(LOGGER_ERROR, "HI_MPI_VPSS_ReleaseChnFrame")
        err->mpp = mpp_error_code;
        return ERR_MPP;
 	}

 	return ERR_NONE;
}

void mpp3_send_frame_to_clients(unsigned int channelId, unsigned int processingId, void* callback) { //TODO move to go space
	callbackFunc func = callback;
	func(processingId, &channelFrames[channelId]);
}

*/
import "C"

import (
    "flag"
    "application/pkg/mpp/vi"
    "application/pkg/mpp/errmpp"
    "application/pkg/logger"
)

var (
    nr bool
    nrFrmNum uint
)
func init() {
    flag.BoolVar(&nr, "vpss-nr", true, "Noise remove enable")
    flag.UintVar(&nrFrmNum, "vpss-nr-frames", 2, "Noise remove reference frames number [1;2]")
}

func maxChannels() uint {
    return uint(C.VPSS_MAX_PHY_CHN_NUM)
}

func initFamily() error {
    var inErr C.error_in
    var in C.hi3516av200_vpss_init_in

    in.width = C.uint(vi.Width())
    in.height = C.uint(vi.Height())

    if nr == true {

        if nrFrmNum < 1 || nrFrmNum > 2 {
            logger.Log.Fatal().
                Uint("vpss-nr-frames", nrFrmNum).
                Msg("vpss-nr-frames shoud be 1 or 2")
        }
        in.nr_frames = C.uchar(nrFrmNum)
        in.nr = 1
    } else {
        in.nr = 0
    }

    logger.Log.Trace().
        Uint("width", uint(in.width)).
        Uint("height", uint(in.height)).
        Uint("nr", uint(in.nr)).
        Uint("nr_frames", uint(in.nr_frames)).
        Msg("VPSS params")


    err := C.hi3516av200_vpss_init(&inErr, &in)

    if err != 0 {
        return errmpp.New("funcname", uint(inErr.mpp))
    }

    return nil
}

func createChannel(channel Channel) { //TODO return error
    var inErr C.error_in
    var in C.hi3516av200_vpss_create_channel_in

    in.channel_id = C.uint(channel.ChannelId)
    in.width = C.uint(channel.Width)
    in.height = C.uint(channel.Height)
    in.vi_fps = C.uint(vi.Fps())
    in.fps = C.uint(channel.Fps)

    logger.Log.Trace().
        Int("channelId", channel.ChannelId).
        Uint("width", uint(in.width)).
        Uint("height", uint(in.height)).
        Uint("vi_fps", uint(in.vi_fps)).
        Uint("fps", uint(in.fps)).
        Msg("VPSS channel params")

    err := C.hi3516av200_vpss_create_channel(&inErr, &in)
    
    if err != 0 {
        logger.Log.Fatal(). //log temporary, should generate and return error
            Str("error", errmpp.New("funcname", uint(inErr.mpp)).Error()).
            Msg("VPSS")
    }

    go func() {
		sendDataToClients(channel)
    }()

    //return nil
}

func destroyChannel(channel Channel) { //TODO return error
    var inErr C.error_in
    var in C.hi3516av200_vpss_destroy_channel_in

    in.channel_id = C.uint(channel.ChannelId)

    err := C.hi3516av200_vpss_destroy_channel(&inErr, &in)

    if err != 0 {
        logger.Log.Fatal(). //log temporary, should generate and return error
            Str("error", errmpp.New("funcname", uint(inErr.mpp)).Error()).
            Msg("VPSS")
    }

    //return nil
}

func sendDataToClients(channel Channel) {
    logger.Log.Trace().
        Int("channelId", channel.ChannelId).
        Str("name", "sendDataToClients").
        Msg("VPSS rutine started")

	for {
		if (!channel.Started){
			break
		}

        var err C.int
        var inErr C.error_in

		err = C.hi3516av200_receive_frame(&inErr, C.uint(channel.ChannelId));
		if err != C.ERR_NONE {
			logger.Log.Warn().
				Int("channelId", channel.ChannelId).
				Str("error", errmpp.New("funcname", uint(inErr.mpp)).Error()).
				Msg("VPSS failed receive frame")
			continue
		}

		for processingId, callback := range channel.Clients {
			C.mpp3_send_frame_to_clients(C.uint(channel.ChannelId), C.uint(processingId), callback);
		}

		err = C.hi3516av200_release_frame(&inErr, C.uint(channel.ChannelId));
		if err != C.ERR_NONE {
			logger.Log.Error().
				Int("channelId", channel.ChannelId).
                Str("error", errmpp.New("funcname", uint(inErr.mpp)).Error()).
                Msg("VPSS failed release frame")
		}
	}

    logger.Log.Trace().        
        Int("channelId", channel.ChannelId).    
        Str("name", "sendDataToClients").
        Msg("VPSS rutine stopped")
}
