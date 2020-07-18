
/******************************************************************************

  Copyright (C), 2016, Hisilicon Tech. Co., Ltd.

 ******************************************************************************
  File Name     : imx415_cmos.c

  Version       : Initial Draft
  Author        : Hisilicon multimedia software group
  Created       : 2013/11/07
  Description   :
  History       :
  1.Date        : 2013/11/07
    Author      :
    Modification: Created file

******************************************************************************/
#if !defined(__IMX415_CMOS_H_)
#define __IMX415_CMOS_H_

#include <stdio.h>
#include <string.h>
#include <assert.h>
#include "hi_comm_sns.h"
#include "hi_comm_video.h"
#include "hi_sns_ctrl.h"
#include "mpi_isp.h"
#include "mpi_ae.h"
#include "mpi_awb.h"

#include "imx415_cmos_ex.h"
#include "imx415_cmos_priv.h"
#include "isp_math_utils.h"

#ifdef __cplusplus
#if __cplusplus
extern "C" {
#endif
#endif /* End of #ifdef __cplusplus */

#define CONFIG_HI_ISP_CA_SUPPORT        1
#define CONFIG_HI_ISP_PREGAMMA_SUPPORT  1
#define CONFIG_HI_ISP_CR_SUPPORT        1

//TODO
//#define CONFIG_HI_ISP_HLC_SUPPORT
//#define CONFIG_HI_ISP_EDGEMARK_SUPPORT


#define IMX415_ID 415

ISP_SNS_STATE_S *g_pastImx415[ISP_MAX_PIPE_NUM] = {HI_NULL};

#define IMX415_SENSOR_GET_CTX(dev, pstCtx)   (pstCtx = g_pastImx415[dev])
#define IMX415_SENSOR_SET_CTX(dev, pstCtx)   (g_pastImx415[dev] = pstCtx)
#define IMX415_SENSOR_RESET_CTX(dev)         (g_pastImx415[dev] = HI_NULL)

static HI_U32 g_au32InitExposure[ISP_MAX_PIPE_NUM]  = {0};
static HI_U32 g_au32LinesPer500ms[ISP_MAX_PIPE_NUM] = {0};
static HI_U16 g_au16InitWBGain[ISP_MAX_PIPE_NUM][3] = {{0}};
static HI_U16 g_au16SampleRgain[ISP_MAX_PIPE_NUM] = {0};
static HI_U16 g_au16SampleBgain[ISP_MAX_PIPE_NUM] = {0};

ISP_SNS_COMMBUS_U g_aunImx415BusInfo[ISP_MAX_PIPE_NUM] = {
    [0] = { .s8I2cDev = 0},
    [1 ... ISP_MAX_PIPE_NUM - 1] = { .s8I2cDev = -1}
};

/****************************************************************************
 * extern                                                                   *
 ****************************************************************************/
extern unsigned char imx415_i2c_addr;
extern unsigned int  imx415_addr_byte;
extern unsigned int  imx415_data_byte;

const IMX415_VIDEO_MODE_TBL_S g_astImx415ModeTbl[IMX415_MODE_BUTT] = {
    {0x08CA, 0, 2, 8, 30.00,  0, "4K2K_10BIT_30FPS"  }, //891Mbps 4lane 37.125MHz
    {0x08CA, 0, 2, 8, 20.00,  1, "4K2K_10BIT_20FPS"  }, //594Mbps 4lane 37.125MHz
};

extern void imx415_init(VI_PIPE ViPipe);
extern void imx415_exit(VI_PIPE ViPipe);
extern void imx415_standby(VI_PIPE ViPipe);
extern void imx415_restart(VI_PIPE ViPipe);
extern void imx415_mirror_flip(VI_PIPE ViPipe, ISP_SNS_MIRRORFLIP_TYPE_E eSnsMirrorFlip);
extern int  imx415_write_register(VI_PIPE ViPipe, int addr, int data);
extern int  imx415_read_register(VI_PIPE ViPipe, int addr);

/****************************************************************************
 * local variables                                                            *
 ****************************************************************************/
#define IMX415_FULL_LINES_MAX  (0xFFFFF)

#define IMX415_ERR_MODE_PRINT(pstSensorImageMode)\
    do{\
        ISP_ERR_TRACE("Not support! Width:%d, Height:%d, Fps:%f, SnsMode:%d\n",\
                  pstSensorImageMode->u16Width,  \
                  pstSensorImageMode->u16Height, \
                  pstSensorImageMode->f32Fps,    \
                  pstSensorImageMode->u8SnsMode);\
    }while(0)

static HI_S32 cmos_get_ae_default(VI_PIPE ViPipe, AE_SENSOR_DEFAULT_S *pstAeSnsDft)
{
    HI_U32 u32Fll = 0;
    HI_FLOAT f32MaxFps = 30;

    ISP_SNS_STATE_S *pstSnsState = HI_NULL;

    CMOS_CHECK_POINTER(pstAeSnsDft);
    IMX415_SENSOR_GET_CTX(ViPipe, pstSnsState);
    CMOS_CHECK_POINTER(pstSnsState);

    u32Fll = g_astImx415ModeTbl[pstSnsState->u8ImgMode].u32VMax;
    f32MaxFps = g_astImx415ModeTbl[pstSnsState->u8ImgMode].f32MaxFps;
    pstAeSnsDft->stIntTimeAccu.f32Offset = g_astImx415ModeTbl[pstSnsState->u8ImgMode].u32Offset;
    pstSnsState->u32FLStd = u32Fll;
    pstAeSnsDft->u32MaxIntTime = pstSnsState->u32FLStd - g_astImx415ModeTbl[pstSnsState->u8ImgMode].u32ExpLineLimit;

    pstAeSnsDft->u32FullLinesStd = pstSnsState->u32FLStd;
    pstAeSnsDft->u32FlickerFreq = 0;
    pstAeSnsDft->u32HmaxTimes = (1000000*1000) / DIV_0_TO_1_FLOAT(u32Fll * f32MaxFps);

    pstAeSnsDft->stIntTimeAccu.enAccuType = AE_ACCURACY_LINEAR;
    pstAeSnsDft->stIntTimeAccu.f32Accuracy = 1;

    pstAeSnsDft->stAgainAccu.enAccuType = AE_ACCURACY_TABLE;
    pstAeSnsDft->stAgainAccu.f32Accuracy = 1;

    pstAeSnsDft->stDgainAccu.enAccuType = AE_ACCURACY_TABLE;
    pstAeSnsDft->stDgainAccu.f32Accuracy = 1;

    pstAeSnsDft->u32ISPDgainShift = 8;
    pstAeSnsDft->u32MinISPDgainTarget = 1 << pstAeSnsDft->u32ISPDgainShift;
    pstAeSnsDft->u32MaxISPDgainTarget = 2 << pstAeSnsDft->u32ISPDgainShift;

    memcpy(&pstAeSnsDft->stPirisAttr, &gstPirisAttr, sizeof(ISP_PIRIS_ATTR_S));
    pstAeSnsDft->enMaxIrisFNO = ISP_IRIS_F_NO_1_4;
    pstAeSnsDft->enMinIrisFNO = ISP_IRIS_F_NO_5_6;

    pstAeSnsDft->bAERouteExValid = HI_FALSE;
    pstAeSnsDft->stAERouteAttr.u32TotalNum = 0;
    pstAeSnsDft->stAERouteAttrEx.u32TotalNum = 0;

    if (g_au32InitExposure[ViPipe] == 0) {
        pstAeSnsDft->u32InitExposure = 1000000;
    } else {
        pstAeSnsDft->u32InitExposure = g_au32InitExposure[ViPipe];;
    }

    if (g_au32LinesPer500ms[ViPipe] == 0) {
        pstAeSnsDft->u32LinesPer500ms = ((HI_U64)(u32Fll * f32MaxFps)) >> 1;
    } else {
        pstAeSnsDft->u32LinesPer500ms = g_au32LinesPer500ms[ViPipe];
    }

    switch (pstSnsState->enWDRMode) {
        default:
        case WDR_MODE_NONE: {
            pstAeSnsDft->au8HistThresh[0] = 0xd;
            pstAeSnsDft->au8HistThresh[1] = 0x28;
            pstAeSnsDft->au8HistThresh[2] = 0x60;
            pstAeSnsDft->au8HistThresh[3] = 0x80;

            pstAeSnsDft->u8AeCompensation = 0x2D;
            pstAeSnsDft->enAeExpMode = AE_EXP_HIGHLIGHT_PRIOR;

            pstAeSnsDft->u32MinIntTime = g_astImx415ModeTbl[pstSnsState->u8ImgMode].u32ExpLineMin;
            pstAeSnsDft->u32MaxIntTimeTarget = 65535;
            pstAeSnsDft->u32MinIntTimeTarget = pstAeSnsDft->u32MinIntTime;

            pstAeSnsDft->u32MaxAgain = IMX415_AGAIN_MAX;
            pstAeSnsDft->u32MinAgain = IMX415_AGAIN_MIN;
            pstAeSnsDft->u32MaxAgainTarget = pstAeSnsDft->u32MaxAgain;
            pstAeSnsDft->u32MinAgainTarget = pstAeSnsDft->u32MinAgain;

            pstAeSnsDft->u32MaxDgain = IMX415_DGAIN_MAX;  /* if Dgain enable,please set ispdgain bigger than 1*/
            pstAeSnsDft->u32MinDgain = IMX415_DGAIN_MIN;
            pstAeSnsDft->u32MaxDgainTarget = pstAeSnsDft->u32MaxDgain;
            pstAeSnsDft->u32MinDgainTarget = pstAeSnsDft->u32MinDgain;

            break;
        }
    }

    return HI_SUCCESS;
}


/* the function of sensor set fps */
static HI_VOID cmos_fps_set(VI_PIPE ViPipe, HI_FLOAT f32Fps, AE_SENSOR_DEFAULT_S *pstAeSnsDft)
{
    HI_FLOAT f32MaxFps;
    HI_U32 u32Lines;
    ISP_SNS_STATE_S *pstSnsState = HI_NULL;

    CMOS_CHECK_POINTER_VOID(pstAeSnsDft);
    IMX415_SENSOR_GET_CTX(ViPipe, pstSnsState);
    CMOS_CHECK_POINTER_VOID(pstSnsState);

    f32MaxFps = g_astImx415ModeTbl[pstSnsState->u8ImgMode].f32MaxFps;
    u32Lines = g_astImx415ModeTbl[pstSnsState->u8ImgMode].u32VMax * (f32MaxFps / DIV_0_TO_1_FLOAT(f32Fps));
    pstSnsState->u32FLStd = u32Lines;
    pstAeSnsDft->u32MaxIntTime = pstSnsState->u32FLStd - g_astImx415ModeTbl[pstSnsState->u8ImgMode].u32ExpLineLimit;

    /* SHR 16bit, So limit full_lines as 0xFFFF */
    if (f32Fps > f32MaxFps || u32Lines > IMX415_FULL_LINES_MAX) {
        ISP_ERR_TRACE("Not support Fps: %f\n", f32Fps);
        return;
    }

    pstAeSnsDft->f32Fps = f32Fps;
    pstAeSnsDft->u32LinesPer500ms = pstSnsState->u32FLStd * f32Fps / 2;
    pstAeSnsDft->u32FullLinesStd = pstSnsState->u32FLStd;

    pstSnsState->au32FL[0] = pstSnsState->u32FLStd;
    pstAeSnsDft->u32FullLines = pstSnsState->au32FL[0];
    pstAeSnsDft->u32HmaxTimes = (1000000*1000) / DIV_0_TO_1_FLOAT(pstSnsState->u32FLStd * f32Fps);

    pstSnsState->astRegsInfo[0].astI2cData[5].u32Data = IMX415_LOW_8BITS(pstSnsState->au32FL[0]);
    pstSnsState->astRegsInfo[0].astI2cData[6].u32Data = IMX415_MID_8BITS(pstSnsState->au32FL[0]);
    pstSnsState->astRegsInfo[0].astI2cData[7].u32Data = IMX415_HIG_4BITS(pstSnsState->au32FL[0]);

    return;
}

static HI_VOID cmos_slow_framerate_set(VI_PIPE ViPipe, HI_U32 u32FullLines,
                                       AE_SENSOR_DEFAULT_S *pstAeSnsDft)
{
    ISP_SNS_STATE_S *pstSnsState = HI_NULL;

    CMOS_CHECK_POINTER_VOID(pstAeSnsDft);
    IMX415_SENSOR_GET_CTX(ViPipe, pstSnsState);
    CMOS_CHECK_POINTER_VOID(pstSnsState);

    u32FullLines = (u32FullLines > IMX415_FULL_LINES_MAX) ? IMX415_FULL_LINES_MAX : u32FullLines;
    pstSnsState->au32FL[0] = u32FullLines;
    pstAeSnsDft->u32FullLines = pstSnsState->au32FL[0];
    pstAeSnsDft->u32MaxIntTime = pstSnsState->au32FL[0] - g_astImx415ModeTbl[pstSnsState->u8ImgMode].u32ExpLineLimit;

    pstSnsState->astRegsInfo[0].astI2cData[5].u32Data = IMX415_LOW_8BITS(pstSnsState->au32FL[0]);
    pstSnsState->astRegsInfo[0].astI2cData[6].u32Data = IMX415_MID_8BITS(pstSnsState->au32FL[0]);
    pstSnsState->astRegsInfo[0].astI2cData[7].u32Data = IMX415_HIG_4BITS(pstSnsState->au32FL[0]);

    return;
}

/* while isp notify ae to update sensor regs, ae call these funcs. */
static HI_VOID cmos_inttime_update(VI_PIPE ViPipe, HI_U32 u32IntTime)
{
    HI_U32 u32Value = 0;
    ISP_SNS_STATE_S *pstSnsState = HI_NULL;
    IMX415_SENSOR_GET_CTX(ViPipe, pstSnsState);
    CMOS_CHECK_POINTER_VOID(pstSnsState);

    u32Value = pstSnsState->au32FL[0] - u32IntTime;
    u32Value = MIN(u32Value, IMX415_FULL_LINES_MAX);
    u32Value = MIN(MAX(u32Value, 9), (pstSnsState->au32FL[0] - 1));

    pstSnsState->astRegsInfo[0].astI2cData[0].u32Data = IMX415_LOW_8BITS(u32Value);   //SHR0
    pstSnsState->astRegsInfo[0].astI2cData[1].u32Data = IMX415_MID_8BITS(u32Value);
    pstSnsState->astRegsInfo[0].astI2cData[2].u32Data = IMX415_HIG_4BITS(u32Value);

    return;
}

/*Again and Dgain use the same table, Step is 0.3db*/
static HI_U32 ad_gain_table[IMX415_AD_GAIN_TBL_RANGE] = {
    1024,   1059,   1097,   1135,   1175,   1217,   1259,   1304,   1349,   1397,   1446,   1497,   1549,   1604, //3.9dB
    1660,   1719,   1779,   1842,   1906,   1973,   2043,   2114,   2189,   2266,   2345,   2428,   2513,   2601, //8.1dB
    2693,   2788,   2886,   2987,   3092,   3201,   3313,   3430,   3550,   3675,   3804,   3938,   4076,   4219, //12.3dB
    4368,   4521,   4680,   4845,   5015,   5191,   5374,   5562,   5758,   5960,   6170,   6387,   6611,   6843, //16.5dB
    7084,   7333,   7591,   7857,   8133,   8419,   8715,   9021,   9338,   9667,  10006,  10358,  10722,  11099, //20.7dB
    11489,  11893,  12311,  12743,  13191,  13655,  14135,  14631,  15146,  15678,  16229,  16799,  17390,  18001, //24.9dB
    18633,  19288,  19966,  20668,  21394,  22146,  22924,  23730,  24564,  25427,  26320,  27245,  28203,  29194, //29.1dB
    30220,  31282,  32381,  41519,  34697,  35917,  37179,  38485,  39838,  41238,  42687,  44187,  45740,  47347, //33.3dB
    49011,  50734,  52517,  54362,  56273,  58250,  60297,  62416,  64610,  66880,  69230,  71663,  74182,  76789, //37.5dB
    79487,  82281,  85172,  88165,  91264,  94471,  97791, 101228, 104785, 108468, 112279, 116225, 120310, 124537, //41.7dB
    128914, 133444, 138134, 142988, 148013, 153215, 158599, 164172, 169942, 175914, 182096, 188495, 195119, 201976, //45.9dB
    209074, 216421, 224027, 231900, 240049, 248485, 257217, 266256, 275613, 285299, 295325, 305703, 316446, 327567, //50.1dB
    339078, 350994, 363329, 376097, 389314, 402995, 417157, 431817, 446992, 462700, 478961, 495793, 513216, 531251, //54.3dB
    549921, 569246, 589250, 609958, 631393, 653582, 676550, 700326, 724936, 750412, 776783, 804081, 832338, 861589, //58.5dB
    891867, 923209, 955652, 989236, 1024000, 1059985, 1097236, 1135795, 1175709, 1217026, 1259795, 1304067, 1349895, 1397333, //62.7dB
    1446438, 1497269, 1549887, 1604353, 1660734, 1719095, 1779508, 1842044, 1906777, 1973786, 2043149, 2114949, 2189273, 2266209, //66.9dB
    2345848, 2428287, 2513622, 2601956, 2693394, 2788046, 2886024, 2987445, 3092431, 3201105, 3313599, 3430046, 3550585, 3675361, //71.1dB
    3804521, 3938220, 4076617                                                                                        //72.0dB
};

static HI_VOID cmos_again_calc_table(VI_PIPE ViPipe, HI_U32 *pu32AgainLin, HI_U32 *pu32AgainDb)
{
    int i;

    CMOS_CHECK_POINTER_VOID(pu32AgainLin);
    CMOS_CHECK_POINTER_VOID(pu32AgainDb);

    if (*pu32AgainLin >= ad_gain_table[IMX415_AGAIN_TBL_RANGE]) {
        *pu32AgainLin = ad_gain_table[IMX415_AGAIN_TBL_RANGE];
        *pu32AgainDb = IMX415_AGAIN_TBL_RANGE;
    } else {
        for (i = 1; i < (IMX415_AGAIN_TBL_RANGE + 1); i++) {
            if (*pu32AgainLin < ad_gain_table[i]) {
                *pu32AgainLin = ad_gain_table[i - 1];
                *pu32AgainDb = i - 1;
                break;
            }
        }
    }

    return;
}

static HI_VOID cmos_dgain_calc_table(VI_PIPE ViPipe, HI_U32 *pu32DgainLin, HI_U32 *pu32DgainDb)
{
    int i;

    CMOS_CHECK_POINTER_VOID(pu32DgainLin);
    CMOS_CHECK_POINTER_VOID(pu32DgainDb);

    if (*pu32DgainLin >= ad_gain_table[IMX415_DGAIN_TBL_RANGE]) {
        *pu32DgainLin = ad_gain_table[IMX415_DGAIN_TBL_RANGE];
        *pu32DgainDb = IMX415_DGAIN_TBL_RANGE;
    } else {
        for (i = 1; i < (IMX415_DGAIN_TBL_RANGE + 1); i++) {
            if (*pu32DgainLin < ad_gain_table[i]) {
                *pu32DgainLin = ad_gain_table[i - 1];
                *pu32DgainDb = i - 1;
                break;
            }
        }
    }

    return;
}

static HI_VOID cmos_gains_update(VI_PIPE ViPipe, HI_U32 u32Again, HI_U32 u32Dgain)
{
    HI_U32 u32Tmp;

    u32Tmp = u32Again + u32Dgain;

    ISP_SNS_STATE_S *pstSnsState = HI_NULL;

    IMX415_SENSOR_GET_CTX(ViPipe, pstSnsState);
    CMOS_CHECK_POINTER_VOID(pstSnsState);

    pstSnsState->astRegsInfo[0].astI2cData[3].u32Data = IMX415_LOW_8BITS(u32Tmp);
    pstSnsState->astRegsInfo[0].astI2cData[4].u32Data = IMX415_MID_8BITS(u32Tmp);

    return;
}



static HI_S32 cmos_init_ae_exp_function(AE_SENSOR_EXP_FUNC_S *pstExpFuncs)
{
    CMOS_CHECK_POINTER(pstExpFuncs);

    memset(pstExpFuncs, 0, sizeof(AE_SENSOR_EXP_FUNC_S));

    pstExpFuncs->pfn_cmos_get_ae_default    = cmos_get_ae_default;
    pstExpFuncs->pfn_cmos_fps_set           = cmos_fps_set;
    pstExpFuncs->pfn_cmos_slow_framerate_set = cmos_slow_framerate_set;
    pstExpFuncs->pfn_cmos_inttime_update    = cmos_inttime_update;
    pstExpFuncs->pfn_cmos_gains_update      = cmos_gains_update;
    pstExpFuncs->pfn_cmos_again_calc_table  = cmos_again_calc_table;
    pstExpFuncs->pfn_cmos_dgain_calc_table  = cmos_dgain_calc_table;

    return HI_SUCCESS;
}


/* AWB default parameter and function */
#define CALIBRATE_STATIC_WB_R_GAIN 518
#define CALIBRATE_STATIC_WB_GR_GAIN 256
#define CALIBRATE_STATIC_WB_GB_GAIN 256
#define CALIBRATE_STATIC_WB_B_GAIN 409

/* Calibration results for Auto WB Planck */
#define CALIBRATE_AWB_P1 30
#define CALIBRATE_AWB_P2 188
#define CALIBRATE_AWB_Q1 -38
#define CALIBRATE_AWB_A1 153679
#define CALIBRATE_AWB_B1 128
#define CALIBRATE_AWB_C1 -101709

/* Rgain and Bgain of the golden sample */
#define GOLDEN_RGAIN 0
#define GOLDEN_BGAIN 0
static HI_S32 cmos_get_awb_default(VI_PIPE ViPipe, AWB_SENSOR_DEFAULT_S *pstAwbSnsDft)
{
    ISP_SNS_STATE_S *pstSnsState = HI_NULL;

    CMOS_CHECK_POINTER(pstAwbSnsDft);
    IMX415_SENSOR_GET_CTX(ViPipe, pstSnsState);
    CMOS_CHECK_POINTER(pstSnsState);

    memset(pstAwbSnsDft, 0, sizeof(AWB_SENSOR_DEFAULT_S));

    pstAwbSnsDft->u16WbRefTemp = 4983;

    pstAwbSnsDft->au16GainOffset[0] = CALIBRATE_STATIC_WB_R_GAIN;
    pstAwbSnsDft->au16GainOffset[1] = CALIBRATE_STATIC_WB_GR_GAIN;
    pstAwbSnsDft->au16GainOffset[2] = CALIBRATE_STATIC_WB_GB_GAIN;
    pstAwbSnsDft->au16GainOffset[3] = CALIBRATE_STATIC_WB_B_GAIN;

    pstAwbSnsDft->as32WbPara[0] = CALIBRATE_AWB_P1;
    pstAwbSnsDft->as32WbPara[1] = CALIBRATE_AWB_P2;
    pstAwbSnsDft->as32WbPara[2] = CALIBRATE_AWB_Q1;
    pstAwbSnsDft->as32WbPara[3] = CALIBRATE_AWB_A1;
    pstAwbSnsDft->as32WbPara[4] = CALIBRATE_AWB_B1;
    pstAwbSnsDft->as32WbPara[5] = CALIBRATE_AWB_C1;
    pstAwbSnsDft->u16GoldenRgain = GOLDEN_RGAIN;
    pstAwbSnsDft->u16GoldenBgain = GOLDEN_BGAIN;


    switch (pstSnsState->enWDRMode) {
        default:
        case WDR_MODE_NONE:
            memcpy(&pstAwbSnsDft->stCcm, &g_stAwbCcm, sizeof(AWB_CCM_S));
            memcpy(&pstAwbSnsDft->stAgcTbl, &g_stAwbAgcTable, sizeof(AWB_AGC_TABLE_S));
            break;
    }
    pstAwbSnsDft->u16InitRgain = g_au16InitWBGain[ViPipe][0];
    pstAwbSnsDft->u16InitGgain = g_au16InitWBGain[ViPipe][1];
    pstAwbSnsDft->u16InitBgain = g_au16InitWBGain[ViPipe][2];
    pstAwbSnsDft->u16SampleRgain = g_au16SampleRgain[ViPipe];
    pstAwbSnsDft->u16SampleBgain = g_au16SampleBgain[ViPipe];

    return HI_SUCCESS;
}

static HI_S32 cmos_get_awb_spec_default(VI_PIPE ViPipe, AWB_SPEC_SENSOR_DEFAULT_S *pstAwbSpecSnsDft)
{
    CMOS_CHECK_POINTER(pstAwbSpecSnsDft);



    return HI_SUCCESS;
}

static HI_S32 cmos_init_awb_exp_function(AWB_SENSOR_EXP_FUNC_S *pstExpFuncs)
{
    CMOS_CHECK_POINTER(pstExpFuncs);

    memset(pstExpFuncs, 0, sizeof(AWB_SENSOR_EXP_FUNC_S));

    pstExpFuncs->pfn_cmos_get_awb_default = cmos_get_awb_default;
    pstExpFuncs->pfn_cmos_get_awb_spec_default = cmos_get_awb_spec_default;

    return HI_SUCCESS;
}


static ISP_CMOS_DNG_COLORPARAM_S g_stDngColorParam = {
    {378, 256, 430},
    {439, 256, 439}
};

static HI_S32 cmos_get_isp_default(VI_PIPE ViPipe, ISP_CMOS_DEFAULT_S *pstDef)
{
    ISP_SNS_STATE_S *pstSnsState = HI_NULL;

    CMOS_CHECK_POINTER(pstDef);
    IMX415_SENSOR_GET_CTX(ViPipe, pstSnsState);
    CMOS_CHECK_POINTER(pstSnsState);

    memset(pstDef, 0, sizeof(ISP_CMOS_DEFAULT_S));
#ifdef CONFIG_HI_ISP_CA_SUPPORT
    pstDef->unKey.bit1Ca       = 1;
    pstDef->pstCa              = &g_stIspCA;
#endif
    pstDef->unKey.bit1Clut     = 1;
    pstDef->pstClut            = &g_stIspCLUT;

    pstDef->unKey.bit1Dpc      = 1;
    pstDef->pstDpc             = &g_stCmosDpc;

    pstDef->unKey.bit1Wdr      = 1;
    pstDef->pstWdr             = &g_stIspWDR;
#ifdef CONFIG_HI_ISP_HLC_SUPPORT
    pstDef->unKey.bit1Hlc      = 0;
    pstDef->pstHlc             = &g_stIspHlc;
#endif
    pstDef->unKey.bit1Lsc      = 1;
    pstDef->pstLsc             = &g_stCmosLsc_8M;

#ifdef CONFIG_HI_ISP_EDGEMARK_SUPPORT
    pstDef->unKey.bit1EdgeMark = 0;
    pstDef->pstEdgeMark        = &g_stIspEdgeMark;
#endif
#ifdef CONFIG_HI_ISP_CR_SUPPORT
    pstDef->unKey.bit1Ge       = 1;
    pstDef->pstGe              = &g_stIspGe;
#endif
    pstDef->unKey.bit1Detail   = 1;
    pstDef->pstDetail          = &g_stIspDetail;

    switch (pstSnsState->enWDRMode) {
        default:
        case WDR_MODE_NONE:
            pstDef->unKey.bit1Demosaic       = 1;
            pstDef->pstDemosaic              = &g_stIspDemosaic;
            pstDef->unKey.bit1Sharpen        = 1;
            pstDef->pstSharpen               = &g_stIspYuvSharpen;
            pstDef->unKey.bit1Drc            = 1;
            pstDef->pstDrc                   = &g_stIspDRC;
            pstDef->unKey.bit1Gamma          = 1;
            pstDef->pstGamma                 = &g_stIspGamma;
            pstDef->unKey.bit1BayerNr        = 1;
            pstDef->pstBayerNr               = &g_stIspBayerNr;
            pstDef->unKey.bit1AntiFalseColor = 1;
            pstDef->pstAntiFalseColor        = &g_stIspAntiFalseColor;
            pstDef->unKey.bit1Ldci           = 1;
            pstDef->pstLdci                  = &g_stIspLdci;
            pstDef->unKey.bit1Dehaze         = 1;
            pstDef->pstDehaze                = &g_stIspDehaze;
            memcpy(&pstDef->stNoiseCalibration, &g_stIspNoiseCalibration, sizeof(ISP_CMOS_NOISE_CALIBRATION_S));
            break;
    }

    pstDef->stSensorMode.u32SensorID = IMX415_ID;
    pstDef->stSensorMode.u8SensorMode = pstSnsState->u8ImgMode;

    memcpy(&pstDef->stDngColorParam, &g_stDngColorParam, sizeof(ISP_CMOS_DNG_COLORPARAM_S));

    switch (pstSnsState->u8ImgMode) {
        default:
            pstDef->stSensorMode.stDngRawFormat.u8BitsPerSample = 12;
            pstDef->stSensorMode.stDngRawFormat.u32WhiteLevel = 4096;
            break;
    }

    pstDef->stSensorMode.stDngRawFormat.stDefaultScale.stDefaultScaleH.u32Denominator = 1;
    pstDef->stSensorMode.stDngRawFormat.stDefaultScale.stDefaultScaleH.u32Numerator = 1;
    pstDef->stSensorMode.stDngRawFormat.stDefaultScale.stDefaultScaleV.u32Denominator = 1;
    pstDef->stSensorMode.stDngRawFormat.stDefaultScale.stDefaultScaleV.u32Numerator = 1;
    pstDef->stSensorMode.stDngRawFormat.stCfaRepeatPatternDim.u16RepeatPatternDimRows = 2;
    pstDef->stSensorMode.stDngRawFormat.stCfaRepeatPatternDim.u16RepeatPatternDimCols = 2;
    pstDef->stSensorMode.stDngRawFormat.stBlcRepeatDim.u16BlcRepeatRows = 2;
    pstDef->stSensorMode.stDngRawFormat.stBlcRepeatDim.u16BlcRepeatCols = 2;
    pstDef->stSensorMode.stDngRawFormat.enCfaLayout = CFALAYOUT_TYPE_RECTANGULAR;
    pstDef->stSensorMode.stDngRawFormat.au8CfaPlaneColor[0] = 0;
    pstDef->stSensorMode.stDngRawFormat.au8CfaPlaneColor[1] = 1;
    pstDef->stSensorMode.stDngRawFormat.au8CfaPlaneColor[2] = 2;
    pstDef->stSensorMode.stDngRawFormat.au8CfaPattern[0] = 1;
    pstDef->stSensorMode.stDngRawFormat.au8CfaPattern[1] = 2;
    pstDef->stSensorMode.stDngRawFormat.au8CfaPattern[2] = 0;
    pstDef->stSensorMode.stDngRawFormat.au8CfaPattern[3] = 1;
    pstDef->stSensorMode.bValidDngRawFormat = HI_TRUE;

    return HI_SUCCESS;
}


static HI_S32 cmos_get_isp_black_level(VI_PIPE ViPipe, ISP_CMOS_BLACK_LEVEL_S *pstBlackLevel)
{
    CMOS_CHECK_POINTER(pstBlackLevel);

    /* It need to update black level when iso change */
    pstBlackLevel->bUpdate = HI_FALSE;

    pstBlackLevel->au16BlackLevel[0] = 200;
    pstBlackLevel->au16BlackLevel[1] = 200;
    pstBlackLevel->au16BlackLevel[2] = 200;
    pstBlackLevel->au16BlackLevel[3] = 200;


    return HI_SUCCESS;

}
static HI_VOID cmos_set_pixel_detect(VI_PIPE ViPipe, HI_BOOL bEnable)
{
    HI_U32 u32FullLines_5Fps = 0;
    HI_U32 u32MaxIntTime_5Fps = 0;

    ISP_SNS_STATE_S *pstSnsState = HI_NULL;

    IMX415_SENSOR_GET_CTX(ViPipe, pstSnsState);
    CMOS_CHECK_POINTER_VOID(pstSnsState);

    if (WDR_MODE_2To1_LINE == pstSnsState->enWDRMode) {
        return;
    } else {
        u32FullLines_5Fps = (g_astImx415ModeTbl[pstSnsState->u8ImgMode].u32VMax * 30) / 5;
    }

    u32MaxIntTime_5Fps = u32FullLines_5Fps - 8;
    return;

    if (bEnable) { /* setup for ISP pixel calibration mode */
        imx415_write_register(ViPipe, IMX415_GAIN_LONG_LOW, 0x00); //gain
        imx415_write_register(ViPipe, IMX415_GAIN_LONG_HIGH, 0x00);

        imx415_write_register(ViPipe, IMX415_VMAX_ADDR_L, IMX415_LOW_8BITS(u32FullLines_5Fps));    //VMAX
        imx415_write_register(ViPipe, IMX415_VMAX_ADDR_M, IMX415_MID_8BITS(u32FullLines_5Fps));
        imx415_write_register(ViPipe, IMX415_VMAX_ADDR_H, IMX415_HIG_4BITS(u32FullLines_5Fps));

        imx415_write_register(ViPipe, IMX415_SHR0_LOW, IMX415_LOW_8BITS(u32MaxIntTime_5Fps));
        imx415_write_register(ViPipe, IMX415_SHR0_MIDDLE, IMX415_MID_8BITS(u32MaxIntTime_5Fps));
        imx415_write_register(ViPipe, IMX415_SHR0_MIDDLE, IMX415_HIG_4BITS(u32MaxIntTime_5Fps));
    } else {/* setup for ISP 'normal mode' */
        pstSnsState->u32FLStd = (pstSnsState->u32FLStd > IMX415_FULL_LINES_MAX) ? IMX415_FULL_LINES_MAX : pstSnsState->u32FLStd;
        imx415_write_register (ViPipe, IMX415_VMAX_ADDR_L, IMX415_LOW_8BITS(pstSnsState->u32FLStd));
        imx415_write_register (ViPipe, IMX415_VMAX_ADDR_M, IMX415_MID_8BITS(pstSnsState->u32FLStd));
        imx415_write_register (ViPipe, IMX415_VMAX_ADDR_H, IMX415_HIG_4BITS(pstSnsState->u32FLStd));
        pstSnsState->bSyncInit = HI_FALSE ;
    }

    return;
}

static HI_S32 cmos_set_wdr_mode(VI_PIPE ViPipe, HI_U8 u8Mode)
{
    ISP_SNS_STATE_S *pstSnsState = HI_NULL;

    IMX415_SENSOR_GET_CTX(ViPipe, pstSnsState);
    CMOS_CHECK_POINTER(pstSnsState);

    pstSnsState->bSyncInit = HI_FALSE;

    switch (u8Mode & 0x3F) {
        case WDR_MODE_NONE:
            pstSnsState->enWDRMode = WDR_MODE_NONE;
            printf("linear mode\n");
            break;

        default:
            ISP_ERR_TRACE("NOT support this mode!\n");
            return HI_FAILURE;
    }

    memset(pstSnsState->au32WDRIntTime, 0, sizeof(pstSnsState->au32WDRIntTime));

    return HI_SUCCESS;
}

static HI_S32 cmos_get_sns_regs_info(VI_PIPE ViPipe, ISP_SNS_REGS_INFO_S *pstSnsRegsInfo)
{
    HI_S32 i;
    ISP_SNS_STATE_S *pstSnsState = HI_NULL;

    CMOS_CHECK_POINTER(pstSnsRegsInfo);
    IMX415_SENSOR_GET_CTX(ViPipe, pstSnsState);
    CMOS_CHECK_POINTER(pstSnsState);

    if ((pstSnsState->bSyncInit == HI_FALSE) || (pstSnsRegsInfo->bConfig == HI_FALSE)) {
        pstSnsState->astRegsInfo[0].enSnsType = ISP_SNS_I2C_TYPE;
        pstSnsState->astRegsInfo[0].unComBus.s8I2cDev = g_aunImx415BusInfo[ViPipe].s8I2cDev;
        pstSnsState->astRegsInfo[0].u8Cfg2ValidDelayMax = 2;
        pstSnsState->astRegsInfo[0].u32RegNum = 8;

        for (i = 0; i < pstSnsState->astRegsInfo[0].u32RegNum; i++) {
            pstSnsState->astRegsInfo[0].astI2cData[i].bUpdate = HI_TRUE;
            pstSnsState->astRegsInfo[0].astI2cData[i].u8DevAddr = imx415_i2c_addr;
            pstSnsState->astRegsInfo[0].astI2cData[i].u32AddrByteNum = imx415_addr_byte;
            pstSnsState->astRegsInfo[0].astI2cData[i].u32DataByteNum = imx415_data_byte;
        }

        //shutter related
        pstSnsState->astRegsInfo[0].astI2cData[0].u8DelayFrmNum = 0;
        pstSnsState->astRegsInfo[0].astI2cData[0].u32RegAddr = IMX415_SHR0_LOW;        //SHR0
        pstSnsState->astRegsInfo[0].astI2cData[1].u8DelayFrmNum = 0;
        pstSnsState->astRegsInfo[0].astI2cData[1].u32RegAddr = IMX415_SHR0_MIDDLE;
        pstSnsState->astRegsInfo[0].astI2cData[2].u8DelayFrmNum = 0;
        pstSnsState->astRegsInfo[0].astI2cData[2].u32RegAddr = IMX415_SHR0_HIGH;

        pstSnsState->astRegsInfo[0].astI2cData[3].u8DelayFrmNum = 0;        //Long Gain
        pstSnsState->astRegsInfo[0].astI2cData[3].u32RegAddr = IMX415_GAIN_LONG_LOW;
        pstSnsState->astRegsInfo[0].astI2cData[4].u8DelayFrmNum = 0;
        pstSnsState->astRegsInfo[0].astI2cData[4].u32RegAddr = IMX415_GAIN_LONG_HIGH;

        //Vmax
        pstSnsState->astRegsInfo[0].astI2cData[5].u8DelayFrmNum = 0;
        pstSnsState->astRegsInfo[0].astI2cData[5].u32RegAddr = IMX415_VMAX_ADDR_L;
        pstSnsState->astRegsInfo[0].astI2cData[6].u8DelayFrmNum = 0;
        pstSnsState->astRegsInfo[0].astI2cData[6].u32RegAddr = IMX415_VMAX_ADDR_M;
        pstSnsState->astRegsInfo[0].astI2cData[7].u8DelayFrmNum = 0;
        pstSnsState->astRegsInfo[0].astI2cData[7].u32RegAddr = IMX415_VMAX_ADDR_H;

        pstSnsState->bSyncInit = HI_TRUE;
    } else {
        for (i = 0; i < pstSnsState->astRegsInfo[0].u32RegNum; i++) {
            if (pstSnsState->astRegsInfo[0].astI2cData[i].u32Data == pstSnsState->astRegsInfo[1].astI2cData[i].u32Data) {
                pstSnsState->astRegsInfo[0].astI2cData[i].bUpdate = HI_FALSE;
            } else {

                pstSnsState->astRegsInfo[0].astI2cData[i].bUpdate = HI_TRUE;
            }
        }
    }
    pstSnsRegsInfo->bConfig = HI_FALSE;

    memcpy(pstSnsRegsInfo, &pstSnsState->astRegsInfo[0], sizeof(ISP_SNS_REGS_INFO_S));
    memcpy(&pstSnsState->astRegsInfo[1], &pstSnsState->astRegsInfo[0], sizeof(ISP_SNS_REGS_INFO_S));

    pstSnsState->au32FL[1] = pstSnsState->au32FL[0];

    return HI_SUCCESS;
}

static HI_S32 cmos_set_image_mode(VI_PIPE ViPipe, ISP_CMOS_SENSOR_IMAGE_MODE_S *pstSensorImageMode)
{
    HI_U8 u8SensorImageMode;
    ISP_SNS_STATE_S *pstSnsState = HI_NULL;
    HI_U8 u8SnsMode;
    HI_U32 u32W, u32H;

    CMOS_CHECK_POINTER(pstSensorImageMode);
    IMX415_SENSOR_GET_CTX(ViPipe, pstSnsState);
    CMOS_CHECK_POINTER(pstSnsState);

    pstSnsState->bSyncInit = HI_FALSE;

    u32H      = pstSensorImageMode->u16Height;
    u32W      = pstSensorImageMode->u16Width;
    u8SnsMode = pstSensorImageMode->u8SnsMode;

    if (IMX415_RES_IS_8M(u32W, u32H)) {
        if (0 == u8SnsMode) {
            u8SensorImageMode = IMX415_8M_30FPS_10BIT_LINEAR_MODE;
        } else if (1 == u8SnsMode) {
            u8SensorImageMode = IMX415_8M_20FPS_10BIT_LINEAR_MODE;
        } else {
            IMX415_ERR_MODE_PRINT(pstSensorImageMode);
            return HI_FAILURE;
        }
    } else {
        IMX415_ERR_MODE_PRINT(pstSensorImageMode);
        return HI_FAILURE;
    }

    /* Switch SensorImageMode */
    if ((pstSnsState->bInit == HI_TRUE) && (u8SensorImageMode == pstSnsState->u8ImgMode)) {
        /* Don't need to switch SensorImageMode */
        return ISP_DO_NOT_NEED_SWITCH_IMAGEMODE;
    }

    pstSnsState->u8ImgMode = u8SensorImageMode;
    pstSnsState->u32FLStd  = g_astImx415ModeTbl[pstSnsState->u8ImgMode].u32VMax;
    pstSnsState->au32FL[0] = pstSnsState->u32FLStd;
    pstSnsState->au32FL[1] = pstSnsState->au32FL[0];

    return HI_SUCCESS;
}

static HI_VOID sensor_global_init(VI_PIPE ViPipe)
{
    ISP_SNS_STATE_S *pstSnsState = HI_NULL;

    IMX415_SENSOR_GET_CTX(ViPipe, pstSnsState);
    CMOS_CHECK_POINTER_VOID(pstSnsState);

    pstSnsState->bInit = HI_FALSE;
    pstSnsState->bSyncInit = HI_FALSE;
    pstSnsState->u8ImgMode = IMX415_8M_30FPS_10BIT_LINEAR_MODE;
    pstSnsState->enWDRMode = WDR_MODE_NONE;
    pstSnsState->u32FLStd = g_astImx415ModeTbl[pstSnsState->u8ImgMode].u32VMax;
    pstSnsState->au32FL[0] = g_astImx415ModeTbl[pstSnsState->u8ImgMode].u32VMax;
    pstSnsState->au32FL[1] = g_astImx415ModeTbl[pstSnsState->u8ImgMode].u32VMax;

    memset(&pstSnsState->astRegsInfo[0], 0, sizeof(ISP_SNS_REGS_INFO_S));
    memset(&pstSnsState->astRegsInfo[1], 0, sizeof(ISP_SNS_REGS_INFO_S));
}

static HI_S32 cmos_init_sensor_exp_function(ISP_SENSOR_EXP_FUNC_S *pstSensorExpFunc)
{
    CMOS_CHECK_POINTER(pstSensorExpFunc);

    memset(pstSensorExpFunc, 0, sizeof(ISP_SENSOR_EXP_FUNC_S));

    pstSensorExpFunc->pfn_cmos_sensor_init = imx415_init;
    pstSensorExpFunc->pfn_cmos_sensor_exit = imx415_exit;
    pstSensorExpFunc->pfn_cmos_sensor_global_init = sensor_global_init;
    pstSensorExpFunc->pfn_cmos_set_image_mode = cmos_set_image_mode;
    pstSensorExpFunc->pfn_cmos_set_wdr_mode = cmos_set_wdr_mode;
    pstSensorExpFunc->pfn_cmos_get_isp_default = cmos_get_isp_default;
    pstSensorExpFunc->pfn_cmos_get_isp_black_level = cmos_get_isp_black_level;
    pstSensorExpFunc->pfn_cmos_set_pixel_detect = cmos_set_pixel_detect;
    pstSensorExpFunc->pfn_cmos_get_sns_reg_info = cmos_get_sns_regs_info;

    return HI_SUCCESS;
}

/****************************************************************************
 * callback structure                                                       *
 ****************************************************************************/

static HI_S32 imx415_set_bus_info(VI_PIPE ViPipe, ISP_SNS_COMMBUS_U unSNSBusInfo)
{
    g_aunImx415BusInfo[ViPipe].s8I2cDev = unSNSBusInfo.s8I2cDev;

    return HI_SUCCESS;
}

static HI_S32 sensor_ctx_init(VI_PIPE ViPipe)
{
    ISP_SNS_STATE_S *pastSnsStateCtx = HI_NULL;

    IMX415_SENSOR_GET_CTX(ViPipe, pastSnsStateCtx);

    if (pastSnsStateCtx == HI_NULL) {
        pastSnsStateCtx = (ISP_SNS_STATE_S *)malloc(sizeof(ISP_SNS_STATE_S));
        if (pastSnsStateCtx == HI_NULL) {
            ISP_ERR_TRACE("Isp[%d] SnsCtx malloc memory failed!\n", ViPipe);
            return HI_ERR_ISP_NOMEM;
        }
    }

    memset(pastSnsStateCtx, 0, sizeof(ISP_SNS_STATE_S));

    IMX415_SENSOR_SET_CTX(ViPipe, pastSnsStateCtx);

    return HI_SUCCESS;
}

static HI_VOID sensor_ctx_exit(VI_PIPE ViPipe)
{
    ISP_SNS_STATE_S *pastSnsStateCtx = HI_NULL;

    IMX415_SENSOR_GET_CTX(ViPipe, pastSnsStateCtx);
    SENSOR_FREE(pastSnsStateCtx);
    IMX415_SENSOR_RESET_CTX(ViPipe);
}

static HI_S32 sensor_register_callback(VI_PIPE ViPipe, ALG_LIB_S *pstAeLib, ALG_LIB_S *pstAwbLib)
{
    HI_S32 s32Ret;
    ISP_SENSOR_REGISTER_S stIspRegister;
    AE_SENSOR_REGISTER_S  stAeRegister;
    AWB_SENSOR_REGISTER_S stAwbRegister;
    ISP_SNS_ATTR_INFO_S   stSnsAttrInfo;

    CMOS_CHECK_POINTER(pstAeLib);
    CMOS_CHECK_POINTER(pstAwbLib);

    s32Ret = sensor_ctx_init(ViPipe);

    if (s32Ret != HI_SUCCESS) {
        return HI_FAILURE;
    }

    stSnsAttrInfo.eSensorId = IMX415_ID;

    s32Ret  = cmos_init_sensor_exp_function(&stIspRegister.stSnsExp);
    s32Ret |= HI_MPI_ISP_SensorRegCallBack(ViPipe, &stSnsAttrInfo, &stIspRegister);

    if (s32Ret != HI_SUCCESS) {
        ISP_ERR_TRACE("sensor register callback function failed!\n");
        return s32Ret;
    }

    s32Ret  = cmos_init_ae_exp_function(&stAeRegister.stSnsExp);
    s32Ret |= HI_MPI_AE_SensorRegCallBack(ViPipe, pstAeLib, &stSnsAttrInfo, &stAeRegister);

    if (s32Ret != HI_SUCCESS) {
        ISP_ERR_TRACE("sensor register callback function to ae lib failed!\n");
        return s32Ret;
    }

    s32Ret  = cmos_init_awb_exp_function(&stAwbRegister.stSnsExp);
    s32Ret |= HI_MPI_AWB_SensorRegCallBack(ViPipe, pstAwbLib, &stSnsAttrInfo, &stAwbRegister);

    if (s32Ret != HI_SUCCESS) {
        ISP_ERR_TRACE("sensor register callback function to awb lib failed!\n");
        return s32Ret;
    }

    return HI_SUCCESS;
}

static HI_S32 sensor_unregister_callback(VI_PIPE ViPipe, ALG_LIB_S *pstAeLib, ALG_LIB_S *pstAwbLib)
{
    HI_S32 s32Ret;

    CMOS_CHECK_POINTER(pstAeLib);
    CMOS_CHECK_POINTER(pstAwbLib);

    s32Ret = HI_MPI_ISP_SensorUnRegCallBack(ViPipe, IMX415_ID);

    if (s32Ret != HI_SUCCESS) {
        ISP_ERR_TRACE("sensor unregister callback function failed!\n");
        return s32Ret;
    }

    s32Ret = HI_MPI_AE_SensorUnRegCallBack(ViPipe, pstAeLib, IMX415_ID);

    if (s32Ret != HI_SUCCESS) {
        ISP_ERR_TRACE("sensor unregister callback function to ae lib failed!\n");
        return s32Ret;
    }

    s32Ret = HI_MPI_AWB_SensorUnRegCallBack(ViPipe, pstAwbLib, IMX415_ID);

    if (s32Ret != HI_SUCCESS) {
        ISP_ERR_TRACE("sensor unregister callback function to awb lib failed!\n");
        return s32Ret;
    }

    sensor_ctx_exit(ViPipe);

    return HI_SUCCESS;
}

static HI_S32 sensor_set_init(VI_PIPE ViPipe, ISP_INIT_ATTR_S *pstInitAttr)
{
    CMOS_CHECK_POINTER(pstInitAttr);

    g_au32InitExposure[ViPipe] = pstInitAttr->u32Exposure;
    g_au32LinesPer500ms[ViPipe] = pstInitAttr->u32LinesPer500ms;
    g_au16InitWBGain[ViPipe][0] = pstInitAttr->u16WBRgain;
    g_au16InitWBGain[ViPipe][1] = pstInitAttr->u16WBGgain;
    g_au16InitWBGain[ViPipe][2] = pstInitAttr->u16WBBgain;
    g_au16SampleRgain[ViPipe] = pstInitAttr->u16SampleRgain;
    g_au16SampleBgain[ViPipe] = pstInitAttr->u16SampleBgain;

    return HI_SUCCESS;
}

ISP_SNS_OBJ_S stSnsImx415Obj = {
    .pfnRegisterCallback    = sensor_register_callback,
    .pfnUnRegisterCallback  = sensor_unregister_callback,
    .pfnStandby             = imx415_standby,
    .pfnRestart             = imx415_restart,
    .pfnMirrorFlip          = imx415_mirror_flip,
    .pfnWriteReg            = imx415_write_register,
    .pfnReadReg             = imx415_read_register,
    .pfnSetBusInfo          = imx415_set_bus_info,
    .pfnSetInit             = sensor_set_init
};


#ifdef __cplusplus
#if __cplusplus
}
#endif
#endif /* End of #ifdef __cplusplus */

#endif /* __IMX415_CMOS_H_ */
