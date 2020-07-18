#include "vo.h"

int mpp_vo_init(error_in *err) {

	VO_PUB_ATTR_S          stVoPubAttr;
    
    stVoPubAttr.enIntfType  = VO_INTF_HDMI;//enVoIntfType;
    stVoPubAttr.enIntfSync  = VO_OUTPUT_1080P60;//VO_OUTPUT_720P60;//enIntfSync;

    stVoPubAttr.u32BgColor  = 0xFF0000;//0x00;//0x0000FF;//COLOR_RGB_BLUE;//0;//pstVoConfig->u32BgColor;

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VO_SetPubAttr, 0, &stVoPubAttr);

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VO_Enable, 0);

    HI_U32 u32BufLen = 3; //????

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VO_SetDisplayBufLen, 0, u32BufLen);

    VO_VIDEO_LAYER_ATTR_S stLayerAttr;

    stLayerAttr.bClusterMode    = HI_FALSE;
    stLayerAttr.bDoubleFrame    = HI_FALSE;
    stLayerAttr.enPixFormat     = PIXEL_FORMAT_YVU_SEMIPLANAR_420;

    RECT_S stDefDispRect  = {0, 0, 1920, 1080};
    stLayerAttr.stDispRect = stDefDispRect;

    stLayerAttr.stImageSize.u32Width  = 1920;
    stLayerAttr.stImageSize.u32Height = 1080;

    stLayerAttr.enDstDynamicRange     = DYNAMIC_RANGE_SDR8;

    stLayerAttr.u32DispFrmRt = 30;

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VO_SetVideoLayerAttr, 0, &stLayerAttr);

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VO_EnableVideoLayer, 0);
    
    VO_CHN_ATTR_S         stVOChnAttr;

    stVOChnAttr.stRect.s32X       = 0;//ALIGN_DOWN((u32Width / u32Square) * (i % u32Square), 2);
    stVOChnAttr.stRect.s32Y       = 0;//ALIGN_DOWN((u32Height / u32Square) * (i / u32Square), 2);
    stVOChnAttr.stRect.u32Width   = 1920;//ALIGN_DOWN(u32Width / u32Square, 2);
    stVOChnAttr.stRect.u32Height  = 1080;//ALIGN_DOWN(u32Height / u32Square, 2);
    stVOChnAttr.u32Priority       = 0;
    stVOChnAttr.bDeflicker        = HI_FALSE;

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VO_SetChnAttr, 0, 0, &stVOChnAttr);

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VO_EnableChn, 0, 0);
    
    HI_HDMI_ATTR_S      stAttr;
    HI_HDMI_ID_E        enHdmiId    = HI_HDMI_ID_0;

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_HDMI_Init);
    
    DO_OR_RETURN_ERR_MPP(err, HI_MPI_HDMI_Open, enHdmiId);
    
    DO_OR_RETURN_ERR_MPP(err, HI_MPI_HDMI_GetAttr, enHdmiId, &stAttr);
    
    stAttr.bEnableHdmi           = HI_TRUE;
    stAttr.bEnableVideo          = HI_TRUE;
    stAttr.enVideoFmt            = HI_HDMI_VIDEO_FMT_1080P_60;//HI_HDMI_VIDEO_FMT_720P_60;//HI_HDMI_VIDEO_FMT_1080P_30;//enVideoFmt;
    stAttr.enVidOutMode          = HI_HDMI_VIDEO_MODE_YCBCR444;//HI_HDMI_VIDEO_MODE_RGB444;//HI_HDMI_VIDEO_MODE_YCBCR444;
    stAttr.enDeepColorMode       = HI_HDMI_DEEP_COLOR_24BIT;
    stAttr.bxvYCCMode            = HI_FALSE;
    stAttr.enOutCscQuantization  = HDMI_QUANTIZATION_LIMITED_RANGE;

    stAttr.bEnableAudio          = HI_FALSE;
    stAttr.enSoundIntf           = HI_HDMI_SND_INTERFACE_I2S;
    stAttr.bIsMultiChannel       = HI_FALSE;

    stAttr.enBitDepth            = HI_HDMI_BIT_DEPTH_16;

    stAttr.bEnableAviInfoFrame   = HI_TRUE;
    stAttr.bEnableAudInfoFrame   = HI_TRUE;
    stAttr.bEnableSpdInfoFrame   = HI_FALSE;
    stAttr.bEnableMpegInfoFrame  = HI_FALSE;

    stAttr.bDebugFlag            = HI_FALSE;
    stAttr.bHDCPEnable           = HI_FALSE;

    stAttr.b3DEnable             = HI_FALSE;
    stAttr.enDefaultMode         = HI_HDMI_FORCE_HDMI;

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_HDMI_SetAttr, enHdmiId, &stAttr);

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_HDMI_Start, enHdmiId);
    
    return ERR_NONE;
}

int mpp_vo_bind_vpss_test(error_in *err) {

    VPSS_CHN_ATTR_S stVpssChnAttr;

    stVpssChnAttr.u32Width                     = 1920;
    stVpssChnAttr.u32Height                    = 1080;
    stVpssChnAttr.enChnMode                    = VPSS_CHN_MODE_USER;
    stVpssChnAttr.enCompressMode               = COMPRESS_MODE_NONE;//COMPRESS_MODE_SEG;
    stVpssChnAttr.enDynamicRange               = DYNAMIC_RANGE_SDR8;
    stVpssChnAttr.enPixelFormat                = PIXEL_FORMAT_YVU_SEMIPLANAR_420;
    stVpssChnAttr.stFrameRate.s32SrcFrameRate  = 30;
    stVpssChnAttr.stFrameRate.s32DstFrameRate  = 30;
    stVpssChnAttr.u32Depth                     = 1;
    stVpssChnAttr.bMirror                      = HI_FALSE;
    stVpssChnAttr.bFlip                        = HI_FALSE;
    stVpssChnAttr.enVideoFormat                = VIDEO_FORMAT_LINEAR;
    stVpssChnAttr.stAspectRatio.enMode         = ASPECT_RATIO_NONE;

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VPSS_SetChnAttr, 0, 2, &stVpssChnAttr)

    DO_OR_RETURN_ERR_MPP(err, HI_MPI_VPSS_EnableChn, 0, 2);

    MPP_CHN_S stSrcChn;
    MPP_CHN_S stDestChn;
    
    stSrcChn.enModId   = HI_ID_VPSS;
    stSrcChn.s32DevId  = 0;
    stSrcChn.s32ChnId  = 2;
   
    //stSrcChn.enModId   = HI_ID_VI;
    //stSrcChn.s32DevId  = 0;
    //stSrcChn.s32ChnId  = 0;

    stDestChn.enModId  = HI_ID_VO;
    stDestChn.s32DevId = 0;
    stDestChn.s32ChnId = 0;
    
    DO_OR_RETURN_ERR_MPP(err, HI_MPI_SYS_Bind, &stSrcChn, &stDestChn);

    return ERR_NONE;
}
