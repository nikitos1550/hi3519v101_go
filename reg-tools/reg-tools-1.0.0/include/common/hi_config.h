/******************************************************************************

  Copyright (C), 2001-2011, Hisilicon Tech. Co., Ltd.

 ******************************************************************************
  File Name     : hi_config.h
  Version       : Initial Draft
  Author        : Hisilicon multimedia software group
  Created       : 2005/5/30
  Last Modified :
  Description   : hi_config.h header file
  Function List :
  History       :
  1.Date        : 2005/5/30
    Author      : T41030
    Modification: Created file

******************************************************************************/

#ifndef __HI_CONFIG_H__
#define __HI_CONFIG_H__


#ifdef __cplusplus
#if __cplusplus
extern "C"{
#endif
#endif /* __cplusplus */


#define DEBUG
/*#define RELEASE*/

#define OS_LINUX

#define LOGQUEUE

#define AZ_STAT

#define MULTI_TASK_LOGQUEUE

#define USE_AZ_INT

#define STAT


#if 0
#define AZ_POOL_LOW
#define AZ_MAGIC_LOW
#endif

//#define H264STREAM_CORRECT
#define RBSTAT

#define BITSTREAM_ENC_CHECKSUM
#define BITSTREAM_DEC_CHECKSUM


#define RTSP_VOD

#define SAVE_VOICE 1


#define MPLAYER_NETWORK

#define DEMO_MEDIA

#define DEMO_VOICE
#define DEMO_VIDEO_ENC
#define DEMO_VIDEO_DEC

//#define SYNC_USE_COND

#define AZPOOLS_X

#define HI3510V100

#if defined(IMAGESIZE_CIF)
#define CONFIG_VIU_CAPTURE_DOWNSCALING //CIF
#endif

#define IO_ADDR_BEGIN   0x10000000
#define IO_ADDR_END		0x13000000

#ifdef __cplusplus
#if __cplusplus
}
#endif
#endif /* __cplusplus */


#endif /* __HI_CONFIG_H__ */
