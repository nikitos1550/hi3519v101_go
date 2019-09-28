#ifndef __HI_MOD_PARAM__
#define __HI_MOD_PARAM__

#include "hi_type.h"
#include "hi_defines.h"

typedef struct hiSYS_MODULE_PARAMS_S
{
    HI_U32 u32VI_VPSS_online;
    HI_U32 u32SensorNum;
    HI_CHAR cSensor[HISI_MAX_SENSOR_NUM][32];
}SYS_MODULE_PARAMS_S;


typedef struct hiVI_MODULE_PARAMS_S
{
    HI_U32 u32DetectErrFrame;
    HI_U32 u32DropErrFrame;
    HI_U32 u32StopIntLevel;
    HI_U32 u32DiscardInt;    
    HI_U32 u32IntdetInterval;
    HI_BOOL bCscTvEnable;
    HI_BOOL bCscCtMode;
    HI_BOOL bYuvSkip;        
}VI_MODULE_PARAMS_S;

typedef struct hiVO_MODULE_PARAMS_S
{
    HI_U32  u32DetectCycle;
    HI_BOOL bTransparentTransmit;
    HI_BOOL bLowPowerMode;
}VO_MODULE_PARAMS_S;

typedef struct hiVPSS_MODULE_PARAMS_S
{
    HI_U32 u32RfrFrameCmp;
    HI_BOOL bOneBufferforLowDelay;    
}VPSS_MODULE_PARAMS_S;

typedef struct hiVGS_MODULE_PARAMS_S
{
    HI_U32 u32MaxVgsJob;
    HI_U32 u32MaxVgsTask;
    HI_U32 u32MaxVgsNode;
    HI_U32 u32WeightThreshold;
}VGS_MODULE_PARAMS_S;

typedef struct hiFISHEYE_MODULE_PARAMS_S
{
    HI_U32 u32MaxFisheyeJob;
    HI_U32 u32MaxFisheyeTask;
    HI_U32 u32MaxFisheyeNode;
    HI_U32 u32WeightThreshold;
}FISHEYE_MODULE_PARAMS_S;

typedef struct hiIVE_MODULE_PARAMS_S
{
	HI_BOOL bSavePowerEn;
}IVE_MODULE_PARAMS_S;

typedef struct hiACODEC_MODULE_PARAMS_S
{
	HI_U32  u32InitDelayTimeMs;
}ACODEC_MODULE_PARAMS_S;

#endif

