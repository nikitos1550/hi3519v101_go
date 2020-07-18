#if !defined(__IMX290_CMOS_H_)
#define __IMX290_CMOS_H_

#include <stdio.h>
#include <string.h>
#include <assert.h>
#include "hi_comm_sns.h"
#include "hi_comm_video.h"
#include "hi_sns_ctrl.h"
#include "mpi_isp.h"
#include "mpi_ae.h"
#include "mpi_awb.h"
#include "mpi_af.h"
#include "imx290_def.h"

#ifdef __cplusplus
#if __cplusplus
extern "C"{
#endif
#endif /* End of #ifdef __cplusplus */

#define IMX290_ID 290


/****************************************************************************
 * local variables                                                            *
 ****************************************************************************/

#define FULL_LINES_MAX  (0x3FFFF)

#define SHS1_ADDR (0x220) 
#define SHS2_ADDR (0x224) 
#define GAIN_ADDR (0x214)
#define HCG_ADDR  (0x209)
#define VMAX_ADDR (0x218)
#define HMAX_ADDR (0x21c)
#define RHS1_ADDR (0x230) 


#define VMAX_IMX290_1080P30_LINE (1133) 
#define VMAX_IMX290_1080P60_LINE (1125)
#define VMAX_IMX290_720P120_LINE (750)
#define VMAX_IMX290_720P60_WDR   (750)

#if SENSOR_IMX290_LINE_WDR_12BIT
#define VMAX_IMX290_1080P30_WDR  (1125)
#else
#define VMAX_IMX290_1080P30_WDR  (1190)
#endif


static HI_U32 gu32BRL = 1109;
static HI_U32 gu32RHS1_Max = (VMAX_IMX290_1080P30_WDR - 1109) * 2 - 21;


static HI_BOOL bInit = HI_FALSE;
HI_BOOL bSensorInit = HI_FALSE;
static HI_U32 gu32FullLinesStd = VMAX_IMX290_1080P30_LINE;
static HI_U32 gu32FullLines = VMAX_IMX290_1080P30_LINE;
static HI_U32 gu32PreFullLines = VMAX_IMX290_1080P30_LINE;


WDR_MODE_E genSensorMode = WDR_MODE_NONE;
HI_U8 gu8SensorImageMode = SENSOR_IMX290_1080P_30FPS_MODE;

static HI_U8 gu8HCGReg = 0x01;

static HI_BOOL gbFPSUp = HI_FALSE;
static HI_BOOL gbVMAXDelay = HI_FALSE;


ISP_SNS_REGS_INFO_S g_stSnsRegsInfo = {0};
ISP_SNS_REGS_INFO_S g_stPreSnsRegsInfo = {0};

static HI_U32 au32WDRIntTime[2] = {0};

#define PATHLEN_MAX 256
#define CMOS_CFG_INI "imx290_cfg.ini"
static char pcName[PATHLEN_MAX] = "configs/imx290_cfg.ini";

/* 2to1 WDR*/
static ISP_AE_ROUTE_EX_S gstAERouteExAttr = 
{
    14,
    {
        {2, 1024, 1024, 1024, 0},
        {9, 1024, 1024, 1024, 0},        //for ghost
        {9, 3584, 1024, 1024, 0},
        {43, 3584, 1024, 1024, 0},       //for flicker
        {43, 5120, 1024, 1024, 0},      
        {65536, 5120, 1024, 1024, 0},    //for noise
        {65536, 5120, 1024, 4096, 0},    
        {65536, 10240, 1024, 4096, 0},   //balance sensor input and isp gain
        {65536, 10240, 1024, 8192, 0},
        {65536, 14336, 1024, 8192, 0},
        {65536, 14336, 1024, 10240, 0},
        {65536, 32768, 1024, 10240, 0},
        {65536, 32768, 1024, 16384, 0},
        {65536, 8153234, 1024, 16384, 0}
    }
};

static HI_S32 cmos_get_ae_default(AE_SENSOR_DEFAULT_S *pstAeSnsDft)
{
    if (HI_NULL == pstAeSnsDft)
    {
        printf("null pointer when get ae default value!\n");
        return -1;
    }
    
    pstAeSnsDft->u32LinesPer500ms = gu32FullLinesStd * 30 / 2;
    pstAeSnsDft->u32FullLinesStd = gu32FullLinesStd;
    pstAeSnsDft->u32FlickerFreq = 0;
    pstAeSnsDft->u32FullLinesMax = FULL_LINES_MAX;

    pstAeSnsDft->stIntTimeAccu.enAccuType = AE_ACCURACY_LINEAR;
    pstAeSnsDft->stIntTimeAccu.f32Accuracy = 1;
    pstAeSnsDft->stIntTimeAccu.f32Offset = 0;

    pstAeSnsDft->stAgainAccu.enAccuType = AE_ACCURACY_TABLE;
    pstAeSnsDft->stAgainAccu.f32Accuracy = 1;

    pstAeSnsDft->stDgainAccu.enAccuType = AE_ACCURACY_LINEAR;
    pstAeSnsDft->stDgainAccu.f32Accuracy = 1;
    
    pstAeSnsDft->u32ISPDgainShift = 8;
    pstAeSnsDft->u32MinISPDgainTarget = 1 << pstAeSnsDft->u32ISPDgainShift;
    pstAeSnsDft->u32MaxISPDgainTarget = 16 << pstAeSnsDft->u32ISPDgainShift; 

    pstAeSnsDft->u32MaxAgain = 8153234;  
    pstAeSnsDft->u32MinAgain = 1024;
    pstAeSnsDft->u32MaxAgainTarget = pstAeSnsDft->u32MaxAgain;
    pstAeSnsDft->u32MinAgainTarget = pstAeSnsDft->u32MinAgain;

    pstAeSnsDft->u32MaxDgain = 1;  
    pstAeSnsDft->u32MinDgain = 1;
    pstAeSnsDft->u32MaxDgainTarget = pstAeSnsDft->u32MaxDgain;
    pstAeSnsDft->u32MinDgainTarget = pstAeSnsDft->u32MinDgain;

    pstAeSnsDft->bAERouteExValid = HI_TRUE;
    
    switch(genSensorMode)
    {
        default:
        case WDR_MODE_NONE:   /*linear mode*/
            pstAeSnsDft->au8HistThresh[0] = 0xd;
            pstAeSnsDft->au8HistThresh[1] = 0x28;
            pstAeSnsDft->au8HistThresh[2] = 0x60;
            pstAeSnsDft->au8HistThresh[3] = 0x80;
            
            pstAeSnsDft->u8AeCompensation = 0x38;
            
            pstAeSnsDft->u32MaxIntTime = gu32FullLinesStd - 2;
            pstAeSnsDft->u32MinIntTime = 1;
            pstAeSnsDft->u32MaxIntTimeTarget = 65535;
            pstAeSnsDft->u32MinIntTimeTarget = 1;
              
        break;  
        case WDR_MODE_2To1_LINE:
            pstAeSnsDft->au8HistThresh[0] = 0xC;
            pstAeSnsDft->au8HistThresh[1] = 0x18;
            pstAeSnsDft->au8HistThresh[2] = 0x60;
            pstAeSnsDft->au8HistThresh[3] = 0x80;

        #if SENSOR_IMX290_LINE_WDR_12BIT
            pstAeSnsDft->u32MaxIntTime = 8;
            pstAeSnsDft->u32MinIntTime = 1; 
        #else
            pstAeSnsDft->u32MaxIntTime = 138;
            pstAeSnsDft->u32MinIntTime = 1; 
        #endif
            pstAeSnsDft->u32MaxIntTimeTarget = 65535;
            pstAeSnsDft->u32MinIntTimeTarget = pstAeSnsDft->u32MinIntTime;

            pstAeSnsDft->u32MaxAgain = 8153234;  
            pstAeSnsDft->u32MinAgain = 1024;
            pstAeSnsDft->u32MaxAgainTarget = 8153234;
            pstAeSnsDft->u32MinAgainTarget = pstAeSnsDft->u32MinAgain;
           
            pstAeSnsDft->u8AeCompensation = 40;
            pstAeSnsDft->u16ManRatioEnable = HI_TRUE;
            pstAeSnsDft->u32Ratio = 0x400;       

            memcpy(&pstAeSnsDft->stAERouteAttrEx, &gstAERouteExAttr, sizeof(ISP_AE_ROUTE_EX_S));
        break;
        case WDR_MODE_2To1_FRAME:
        case WDR_MODE_2To1_FRAME_FULL_RATE:
            pstAeSnsDft->au8HistThresh[0] = 0xC;
            pstAeSnsDft->au8HistThresh[1] = 0x18;
            pstAeSnsDft->au8HistThresh[2] = 0x60;
            pstAeSnsDft->au8HistThresh[3] = 0x80;

            pstAeSnsDft->u32MaxIntTime = gu32FullLinesStd - 2;
            pstAeSnsDft->u32MinIntTime = 1;
            pstAeSnsDft->u32MaxIntTimeTarget = 65535;
            pstAeSnsDft->u32MinIntTimeTarget = pstAeSnsDft->u32MinIntTime;
            
            pstAeSnsDft->u8AeCompensation = 0x38;
       break;
    }

    return 0;
}


/* the function of sensor set fps */
static HI_VOID cmos_fps_set(HI_FLOAT f32Fps, AE_SENSOR_DEFAULT_S *pstAeSnsDft)
{
    HI_U32 u32VMAX = VMAX_IMX290_1080P30_LINE;
    
    switch (gu8SensorImageMode)
    {   
    case SENSOR_IMX290_1080P_30FPS_MODE:
        if ((f32Fps <= 30) && (f32Fps >= 0.5))
        {            
            if (WDR_MODE_2To1_LINE == genSensorMode)
            {
                u32VMAX = VMAX_IMX290_1080P30_WDR * 30 / f32Fps;  
            }
            else
            {
                u32VMAX = VMAX_IMX290_1080P30_LINE * 30 / f32Fps;
            }
        }
        else
        {
            printf("Not support Fps: %f\n", f32Fps);
            return;
        }
        break;
        
    case SENSOR_IMX290_1080P_60FPS_MODE:
        if ((f32Fps <= 60) && (f32Fps >= 0.5))
        { 
            u32VMAX = VMAX_IMX290_1080P60_LINE * 60 / f32Fps;
        }
        else
        {
            printf("Not support Fps: %f\n", f32Fps);
            return;
        }
        break;

    case SENSOR_IMX290_720P_60FPS_MODE:
        if ((f32Fps <= 60) && (f32Fps >= 0.5))
        {            
            if (WDR_MODE_2To1_LINE == genSensorMode)
            {
                u32VMAX = VMAX_IMX290_720P60_WDR * 60 / f32Fps;  
            }
            else
            {
            }
        }
        else
        {
            printf("Not support Fps: %f\n", f32Fps);
            return;
        }
        break;
    case SENSOR_IMX290_720P_120FPS_MODE:
        if ((f32Fps <= 120) && (f32Fps >= 0.5))
        {            
            u32VMAX = VMAX_IMX290_720P120_LINE * 120 / f32Fps;
        }
        else
        {
            printf("Not support Fps: %f\n", f32Fps);
            return;
        }
        break;
        
    default:
        break;
    }
    
    u32VMAX = (u32VMAX > FULL_LINES_MAX) ? FULL_LINES_MAX : u32VMAX;  

    if (WDR_MODE_NONE == genSensorMode)
    {
        g_stSnsRegsInfo.astSspData[5].u32Data = (u32VMAX & 0xFF);
        g_stSnsRegsInfo.astSspData[6].u32Data = ((u32VMAX & 0xFF00) >> 8);
        g_stSnsRegsInfo.astSspData[7].u32Data = ((u32VMAX & 0x30000) >> 16);
    }    
    else
    {
        g_stSnsRegsInfo.astSspData[8].u32Data = (u32VMAX & 0xFF);
        g_stSnsRegsInfo.astSspData[9].u32Data = ((u32VMAX & 0xFF00) >> 8);
        g_stSnsRegsInfo.astSspData[10].u32Data = ((u32VMAX & 0x30000) >> 16);      
    }

    if (WDR_MODE_2To1_LINE == genSensorMode)
    {
        gu32FullLinesStd = u32VMAX * 2;

        /*
            RHS1 limitation:
            2n + 5
            RHS1 <= FSC - BRL*2 -21
            (2 * VMAX_IMX290_1080P30_WDR - 2 * gu32BRL - 21) - (((2 * VMAX_IMX290_1080P30_WDR - 2 * 1109 - 21) - 5) %2)
        */
        gu32RHS1_Max = (u32VMAX - gu32BRL) * 2 - 21;
        
        //adjust in cmos_get_inttime_max
        pstAeSnsDft->u32MaxIntTime = gu32RHS1_Max - 3;
      
    }
    else
    {
        gu32FullLinesStd = u32VMAX;
        pstAeSnsDft->u32MaxIntTime = gu32FullLinesStd - 2;
    }

   gbFPSUp |=( (pstAeSnsDft->f32Fps > f32Fps) || (pstAeSnsDft->f32Fps < f32Fps));
   
    pstAeSnsDft->f32Fps = f32Fps;
    pstAeSnsDft->u32LinesPer500ms = gu32FullLinesStd * f32Fps / 2;
    pstAeSnsDft->u32FullLinesStd = gu32FullLinesStd;
    //pstAeSnsDft->u32MaxIntTime= gu32FullLinesStd - 2;
    gu32FullLines = gu32FullLinesStd;

    return;
}

static HI_VOID cmos_slow_framerate_set(HI_U32 u32FullLines,
    AE_SENSOR_DEFAULT_S *pstAeSnsDft)
{
    if (WDR_MODE_2To1_LINE == genSensorMode)
    {
        u32FullLines >>= 1;
        u32FullLines = (u32FullLines > FULL_LINES_MAX) ? FULL_LINES_MAX : u32FullLines;
        gu32FullLines = u32FullLines << 1;
    }
    else
    {
        u32FullLines = (u32FullLines > FULL_LINES_MAX) ? FULL_LINES_MAX : u32FullLines;
        gu32FullLines = u32FullLines;
    }
    
    if (WDR_MODE_NONE == genSensorMode)
    {
        g_stSnsRegsInfo.astSspData[5].u32Data = (u32FullLines & 0xFF);
        g_stSnsRegsInfo.astSspData[6].u32Data = ((u32FullLines & 0xFF00) >> 8);
        g_stSnsRegsInfo.astSspData[7].u32Data = ((u32FullLines & 0x30000) >> 16);
    }    
    else
    {
        g_stSnsRegsInfo.astSspData[8].u32Data = (u32FullLines & 0xFF);
        g_stSnsRegsInfo.astSspData[9].u32Data = ((u32FullLines & 0xFF00) >> 8);
        g_stSnsRegsInfo.astSspData[10].u32Data = ((u32FullLines & 0x30000) >> 16);      
    }
    
   gbFPSUp |= (gu32PreFullLines < gu32FullLines);

    if (WDR_MODE_2To1_LINE == genSensorMode)
    {
        /*
            RHS1 limitation:
            2n + 5
            RHS1 <= FSC - BRL*2 -21
            (2 * VMAX_IMX290_1080P30_WDR - 2 * gu32BRL - 21) - (((2 * VMAX_IMX290_1080P30_WDR - 2 * 1109 - 21) - 5) %2)
        */
        gu32RHS1_Max = (u32FullLines - gu32BRL) * 2 - 21;
        
        //adjust in cmos_get_inttime_max
       pstAeSnsDft->u32MaxIntTime = gu32RHS1_Max - 3;
      
    }
    else
    {
        pstAeSnsDft->u32MaxIntTime = gu32FullLines - 2;
    }
    
    return;
}

/* while isp notify ae to update sensor regs, ae call these funcs. */
static HI_VOID cmos_inttime_update(HI_U32 u32IntTime)
{
    static HI_BOOL bFirst = HI_TRUE;
    HI_U32 u32Value = 0;

    static HI_U32 u32ShortIntTime;
    static HI_U32 u32LongIntTime;
    static HI_U32 u32RHS1 = 0x3e;
    HI_U32 u32SHS1;
    HI_U32 u32SHS2;
    
    if (WDR_MODE_2To1_LINE == genSensorMode)
    {
        if (bFirst) /* short exposure */
        {
            u32ShortIntTime = u32IntTime;      
            au32WDRIntTime[0] = u32IntTime;
            bFirst = HI_FALSE;
        }
        else /* long exposure */
        {            
            u32LongIntTime = u32IntTime;       
            au32WDRIntTime[1] = u32IntTime;

            u32SHS2 = gu32PreFullLines - u32LongIntTime - 1;
                
            //allocate the RHS1
            u32SHS1 = (u32ShortIntTime % 2) + 2;
            u32RHS1 = u32ShortIntTime + u32SHS1 + 1;
       

            g_stSnsRegsInfo.astSspData[0].u32Data = (u32SHS1 & 0xFF);
            g_stSnsRegsInfo.astSspData[1].u32Data = ((u32SHS1 & 0xFF00) >> 8);
            g_stSnsRegsInfo.astSspData[2].u32Data = ((u32SHS1 & 0x30000) >> 16);

            g_stSnsRegsInfo.astSspData[5].u32Data = (u32SHS2 & 0xFF);
            g_stSnsRegsInfo.astSspData[6].u32Data = ((u32SHS2 & 0xFF00) >> 8);
            g_stSnsRegsInfo.astSspData[7].u32Data = ((u32SHS2 & 0x30000) >> 16);

            g_stSnsRegsInfo.astSspData[11].u32Data = (u32RHS1 & 0xFF);
            g_stSnsRegsInfo.astSspData[12].u32Data = ((u32RHS1 & 0xFF00) >> 8);
            g_stSnsRegsInfo.astSspData[13].u32Data = ((u32RHS1 & 0xF0000) >> 16);

            

            
            bFirst = HI_TRUE;
        }
    }
    else if (WDR_MODE_2To1_FRAME == genSensorMode ||WDR_MODE_2To1_FRAME_FULL_RATE ==genSensorMode)
    {
        if (bFirst) /* short exposure */
        {
            bFirst = HI_FALSE;
            u32Value = gu32FullLines - u32IntTime - 1;
            g_stSnsRegsInfo.astSspData[0].u32Data = u32Value & 0xFF;
            g_stSnsRegsInfo.astSspData[1].u32Data = (u32Value & 0xFF00) >> 8;
            g_stSnsRegsInfo.astSspData[2].u32Data = (u32Value & 0x30000) >> 16;
            
        }
        else /* long exposure */
        {
            bFirst = HI_TRUE;
            u32Value = gu32FullLines - u32IntTime - 1;
            g_stSnsRegsInfo.astSspData[5].u32Data = u32Value & 0xFF;
            g_stSnsRegsInfo.astSspData[6].u32Data = (u32Value & 0xFF00) >> 8;
            g_stSnsRegsInfo.astSspData[7].u32Data = (u32Value & 0x30000) >> 16; 
        }
    }
     else
    {        
        u32Value = gu32FullLines - u32IntTime - 1;

        g_stSnsRegsInfo.astSspData[0].u32Data = (u32Value & 0xFF);
        g_stSnsRegsInfo.astSspData[1].u32Data = ((u32Value & 0xFF00) >> 8);
        g_stSnsRegsInfo.astSspData[2].u32Data = ((u32Value & 0x30000) >> 16);
        bFirst = HI_TRUE;
    }

    return;
}



static HI_U32 gain_table[262]=
{
    1024,1059,1097,1135,1175,1217,1259,1304,1349,1397,1446,1497,1549,1604,1660,1719,1779,1842,1906,
    1973,2043,2048,2119,2194,2271,2351,2434,2519,2608,2699,2794,2892,2994,3099,3208,3321,3438,3559,
    3684,3813,3947,4086,4229,4378,4532,4691,4856,5027,5203,5386,5576,5772,5974,6184,6402,6627,6860,
    7101,7350,7609,7876,8153,8439,8736,9043,9361,9690,10030,10383,10748,11125,11516,11921,12340,12774,
    13222,13687,14168,14666,15182,15715,16267,16839,17431,18043,18677,19334,20013,20717,21445,22198,
    22978,23786,24622,25487,26383,27310,28270,29263,30292,31356,32458,33599,34780,36002,37267,38577,
    39932,41336,42788,44292,45849,47460,49128,50854,52641,54491,56406,58388,60440,62564,64763,67039,
    69395,71833,74358,76971,79676,82476,85374,88375,91480,94695,98023,101468,105034,108725,112545,
    116501,120595,124833,129220,133761,138461,143327,148364,153578,158975,164562,170345,176331,182528,
    188942,195582,202455,209570,216935,224558,232450,240619,249074,257827,266888,276267,285976,296026,
    306429,317197,328344,339883,351827,364191,376990,390238,403952,418147,432842,448053,463799,480098,
    496969,514434,532512,551226,570597,590649,611406,632892,655133,678156,701988,726657,752194,778627,
    805990,834314,863634,893984,925400,957921,991585,1026431,1062502,1099841,1138491,1178500,1219916,
    1262786,1307163,1353100,1400651,1449872,1500824,1553566,1608162,1664676,1723177,1783733,1846417,
    1911304,1978472,2048000,2119971,2194471,2271590,2351418,2434052,2519590,2608134,2699789,2794666,
    2892876,2994538,3099773,3208706,3321467,3438190,3559016,3684087,3813554,3947571,4086297,4229898,
    4378546,4532417,4691696,4856573,5027243,5203912,5386788,5576092,5772048,5974890,6184861,6402210,
    6627198,6860092,7101170,7350721,7609041,7876439,8153234
};

static HI_VOID cmos_again_calc_table(HI_U32 *pu32AgainLin, HI_U32 *pu32AgainDb)
{
    int i;

    if (*pu32AgainLin >= gain_table[261])
    {
         *pu32AgainLin = gain_table[261];
         *pu32AgainDb = 261;
         return ;
    }
    
    for (i = 1; i < 262; i++)
    {
        if (*pu32AgainLin < gain_table[i])
        {
            *pu32AgainLin = gain_table[i - 1];
            *pu32AgainDb = i - 1;
            break;
        }
    }
    return;
}



static HI_VOID cmos_gains_update(HI_U32 u32Again, HI_U32 u32Dgain)
{  
    HI_U32 u32HCG = gu8HCGReg;
    
    if(u32Again >= 21)
    {
        u32HCG = u32HCG | 0x10;  // bit[4] HCG  .Reg0x209[7:0]
        u32Again = u32Again - 21;
    }
        
    
    g_stSnsRegsInfo.astSspData[3].u32Data = (u32Again & 0xFF);
    g_stSnsRegsInfo.astSspData[4].u32Data = (u32HCG & 0xFF);
    

    return;
}

/* Only used in WDR_MODE_2To1_LINE and WDR_MODE_2To1_FRAME mode */
static HI_VOID cmos_get_inttime_max(HI_U32 u32Ratio, HI_U32 *pu32IntTimeMax)
{
    HI_U32  u32ShortIntTimeMax;  
    HI_U32  u32IntTimeMaxTmp = 0;
    
    if(HI_NULL == pu32IntTimeMax)
    {
        printf("null pointer when get ae sensor IntTimeMax value!\n");
        return;
    }
    
    if ((WDR_MODE_2To1_FRAME_FULL_RATE == genSensorMode) || (WDR_MODE_2To1_FRAME == genSensorMode))
    {
        *pu32IntTimeMax = (gu32FullLines - 2) * 0x40 / DIV_0_TO_1(u32Ratio);
    }
    else if(WDR_MODE_2To1_LINE == genSensorMode) 
    {
        /*  limitation for line base WDR
    
            SHS1 limitation:
            2 or more   
            RHS1 - 2 or less

            SHS2 limitation:
            RHS1 + 2 or more
            FSC - 2 or less

            RHS1 Limitation
            2n + 5 (n = 0,1,2...)
            RHS1 <= FSC - BRL * 2 - 21

            short exposure time = RHS1 - (SHS1 + 1) <= RHS1 - 3
            long exposure time = FSC - (SHS2 + 1) <= FSC - (RHS1 + 3)
            ExposureShort + ExposureLong <= FSC - 6
            short exposure time <= (FSC - 6) / (ratio + 1)
        */

        u32IntTimeMaxTmp = ((gu32PreFullLines - 6) * 0x40)  / DIV_0_TO_1(u32Ratio + 0x40);
        u32ShortIntTimeMax = ((gu32FullLines - 6) * 0x40)  / DIV_0_TO_1(u32Ratio + 0x40);      
        u32ShortIntTimeMax = (u32IntTimeMaxTmp < u32ShortIntTimeMax)? u32IntTimeMaxTmp : u32ShortIntTimeMax;        
        u32ShortIntTimeMax = (u32ShortIntTimeMax > (gu32RHS1_Max - 3))? (gu32RHS1_Max - 3): u32ShortIntTimeMax;
        u32ShortIntTimeMax = (0 == u32ShortIntTimeMax)? 1: u32ShortIntTimeMax;

        *pu32IntTimeMax = u32ShortIntTimeMax;
    }   

    return;
}

HI_S32 cmos_init_ae_exp_function(AE_SENSOR_EXP_FUNC_S *pstExpFuncs)
{
    memset(pstExpFuncs, 0, sizeof(AE_SENSOR_EXP_FUNC_S));

    pstExpFuncs->pfn_cmos_get_ae_default     = cmos_get_ae_default;
    pstExpFuncs->pfn_cmos_fps_set            = cmos_fps_set;
    pstExpFuncs->pfn_cmos_slow_framerate_set = cmos_slow_framerate_set;    
    pstExpFuncs->pfn_cmos_inttime_update     = cmos_inttime_update;
    pstExpFuncs->pfn_cmos_gains_update       = cmos_gains_update;
    pstExpFuncs->pfn_cmos_again_calc_table   = cmos_again_calc_table; 
    pstExpFuncs->pfn_cmos_dgain_calc_table   = NULL;
    pstExpFuncs->pfn_cmos_get_inttime_max    = cmos_get_inttime_max;  
    return 0;
}


static AWB_CCM_S g_stAwbCcm =
{
 
   5291,
   {
      0x01DC, 0x80B2, 0x8029,
      0x805A, 0x0198, 0x803D,
      0x0013, 0x811B, 0x0208
   },
   3184,
   {
      0x01C3, 0x808B, 0x8038,
      0x8085, 0x01A5, 0x8020,
      0x003A, 0x814C, 0x0211
   }, 
   2457,
   {     
      0x0172, 0x801C, 0x8056,
      0x808C, 0x01A0, 0x8014,
      0x801D, 0x8157, 0x0274,
   }
};

static AWB_AGC_TABLE_S g_stAwbAgcTable =
{
    /* bvalid */
    1,

    /* saturation */ 
    {128,128,128,128,128,120,112,104,96,88,80,72,64,58,58,58}
};

static HI_S32 cmos_get_awb_default(AWB_SENSOR_DEFAULT_S *pstAwbSnsDft)
{
    if (HI_NULL == pstAwbSnsDft)
    {
        printf("null pointer when get awb default value!\n");
        return -1;
    }

    memset(pstAwbSnsDft, 0, sizeof(AWB_SENSOR_DEFAULT_S));
    pstAwbSnsDft->u16WbRefTemp = 4900;

    pstAwbSnsDft->au16GainOffset[0] = 0x1A2;
    pstAwbSnsDft->au16GainOffset[1] = 0x100;
    pstAwbSnsDft->au16GainOffset[2] = 0x100;
    pstAwbSnsDft->au16GainOffset[3] = 0x21D;

    pstAwbSnsDft->as32WbPara[0] = 72;
    pstAwbSnsDft->as32WbPara[1] = 36;
    pstAwbSnsDft->as32WbPara[2] = -148;
    pstAwbSnsDft->as32WbPara[3] = 190509;
    pstAwbSnsDft->as32WbPara[4] = 128;
    pstAwbSnsDft->as32WbPara[5] = -141264;

    memcpy(&pstAwbSnsDft->stCcm, &g_stAwbCcm, sizeof(AWB_CCM_S));
    memcpy(&pstAwbSnsDft->stAgcTbl, &g_stAwbAgcTable, sizeof(AWB_AGC_TABLE_S));
    
    return 0;
}


HI_S32 cmos_init_awb_exp_function(AWB_SENSOR_EXP_FUNC_S *pstExpFuncs)
{
    memset(pstExpFuncs, 0, sizeof(AWB_SENSOR_EXP_FUNC_S));

    pstExpFuncs->pfn_cmos_get_awb_default = cmos_get_awb_default;

    return 0;
}


static ISP_CMOS_AGC_TABLE_S g_stIspAgcTable =
{
    /* bvalid */
    1,
    
    /* 100, 200, 400, 800, 1600, 3200, 6400, 12800, 25600, 51200, 102400, 204800, 409600, 819200, 1638400, 3276800 */

    /* sharpen_alt_d */
    {0x38,0x38,0x36,0x32,0x2b,0x26,0x20,0x1c,0x1a,0x16,0x14,0x12,0x12,0x12,0x12,0x12},
        
    /* sharpen_alt_ud */
    {0x32,0x32,0x30,0x30,0x30,0x2a,0x2a,0x2a,0x26,0x26,0x20,0x16,0x12,0x12,0x12,0x12},
        
    /* snr_thresh Max=0x54 */
    {0x08,0x0a,0x0f,0x12,0x16,0x1a,0x22,0x28,0x2e,0x36,0x3a,0x40,0x40,0x40,0x40,0x40},
        
    /* demosaic_lum_thresh */
    {0x58,0x58,0x56,0x4e,0x46,0x3a,0x30,0x28,0x24,0x20,0x20,0x20,0x20,0x20,0x20,0x20},
        
    /* demosaic_np_offset */
    {0x00,0x0a,0x12,0x1a,0x20,0x28,0x30,0x32,0x34,0x36,0x38,0x38,0x38,0x38,0x38,0x38},
        
    /* ge_strength */
    {0x55,0x55,0x55,0x55,0x55,0x55,0x37,0x37,0x37,0x35,0x35,0x35,0x35,0x35,0x35,0x35},

    /* rgb_sharpen_strength */
    {0x58,0x58,0x56,0x4e,0x46,0x36,0x36,0x34,0x34,0x32,0x30,0x20,0x20,0x20,0x20,0x20}
};

static ISP_CMOS_AGC_TABLE_S g_stIspAgcTableFSWDR =
{
    /* bvalid */
    1,
    
    /* 100, 200, 400, 800, 1600, 3200, 6400, 12800, 25600, 51200, 102400, 204800, 409600, 819200, 1638400, 3276800 */

    /* sharpen_alt_d */
    {0x3C,0x3C,0x3C,0x3C,0x38,0x30,0x28,0x20,0x20,0x20,0x20,0x20,0x20,0x20,0x20,0x20},
    
    /* sharpen_alt_ud */
    {0x60,0x60,0x60,0x60,0x50,0x40,0x30,0x20,0x10,0x10,0x10,0x10,0x10,0x10,0x10,0x10},
        
    /* snr_thresh */
    {0x8,0xC,0x10,0x14,0x18,0x20,0x28,0x30,0x30,0x30,0x30,0x30,0x30,0x30,0x30,0x30},
        
    /* demosaic_lum_thresh */
    {0x50,0x50,0x40,0x40,0x30,0x30,0x20,0x20,0x20,0x20,0x20,0x20,0x20,0x20,0x20,0x20},
        
    /* demosaic_np_offset */
    {0x0,0xa,0x12,0x1a,0x20,0x28,0x30,0x30,0x30,0x30,0x30,0x30,0x30,0x30,0x30,0x30},
        
    /* ge_strength */
    {0x55,0x55,0x55,0x55,0x55,0x55,0x37,0x37,0x37,0x37,0x37,0x37,0x37,0x37,0x37,0x37},

    /* RGBsharpen_strength */
    {0x60,0x60,0x60,0x60,0x50,0x40,0x30,0x20,0x20,0x20,0x20,0x20,0x20,0x20,0x20,0x20}
};


static ISP_CMOS_NOISE_TABLE_S g_stIspNoiseTable =
{
    /* bvalid */
    1,

    /* nosie_profile_weight_lut */
    {          
        0, 0, 0, 0, 0, 0, 0, 3, 11, 16, 19, 22, 24, 26, 27, 28, 30, 31, 32, 32, 34, 35, 35, 36, 37, 37, 37,
        38, 39, 39, 40, 40, 41, 41, 42, 42, 42, 43, 43, 43, 44, 44, 44, 45, 45, 45, 46, 46, 46, 47, 47, 47,
        47, 48, 48, 48, 48, 48, 49, 49, 49, 49, 50, 50, 50, 50, 50, 51, 51, 51, 51, 51, 51, 52, 52, 52, 52,
        52, 52, 53, 53, 53, 53, 53, 53, 54, 54, 54, 54, 54, 54, 54, 54, 55, 55, 55, 55, 55, 55, 55, 56, 56,
        56, 56, 56, 56, 56, 56, 56, 57, 57, 57, 57, 57, 57, 57, 57, 57, 58, 58, 58, 58, 58, 58, 58, 58, 58,
        58
    },

    /* demosaic_weight_lut */
    {
        3, 11, 16, 19, 22, 24, 26, 27, 28, 30, 31, 32, 32, 34, 35, 35, 36, 37, 37, 37, 38, 39, 39, 40, 40,
        41, 41, 42, 42, 42, 43, 43, 43, 44, 44, 44, 45, 45, 45, 46, 46, 46, 47, 47, 47, 47, 48, 48, 48, 48,
        48, 49, 49, 49, 49, 50, 50, 50, 50, 50, 51, 51, 51, 51, 51, 51, 52, 52, 52, 52, 52, 52, 53, 53, 53,
        53, 53, 53, 54, 54, 54, 54, 54, 54, 54, 54, 55, 55, 55, 55, 55, 55, 55, 56, 56, 56, 56, 56, 56, 56, 
        56, 56, 57, 57, 57, 57, 57, 57, 57, 57, 57, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58,
        58, 58, 58
    }        
    
};

static ISP_CMOS_NOISE_TABLE_S g_stIspNoiseTableFSWDR =
{
    /* bvalid */
    1,
    
    /* nosie_profile_weight_lut */
    {
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 45,
    },

    /* demosaic_weight_lut */
    {
        3, 11, 16, 19, 22, 24, 26, 27, 28, 30, 31, 32, 32, 34, 35, 35, 36, 37, 37, 37, 38, 39, 39, 40, 40,
        41, 41, 42, 42, 42, 43, 43, 43, 44, 44, 44, 45, 45, 45, 46, 46, 46, 47, 47, 47, 47, 48, 48, 48, 48,
        48, 49, 49, 49, 49, 50, 50, 50, 50, 50, 51, 51, 51, 51, 51, 51, 52, 52, 52, 52, 52, 52, 53, 53, 53,
        53, 53, 53, 54, 54, 54, 54, 54, 54, 54, 54, 55, 55, 55, 55, 55, 55, 55, 56, 56, 56, 56, 56, 56, 56, 
        56, 56, 57, 57, 57, 57, 57, 57, 57, 57, 57, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58, 58,
        58, 58, 58
    }
};

static ISP_CMOS_DEMOSAIC_S g_stIspDemosaic =
{
    /* bvalid */
    1,
    
    /*vh_slope*/
    0xd0,

    /*aa_slope*/
    0xbd,

    /*va_slope*/
    0xca,

    /*uu_slope*/
    0xbd,

    /*sat_slope*/
    0x5d,

    /*ac_slope*/
    0xa0,
    
    /*fc_slope*/
    0x8a,

    /*vh_thresh*/
    0x0,

    /*aa_thresh*/
    0x00,

    /*va_thresh*/
    0x00,

    /*uu_thresh*/
    0x08,

    /*sat_thresh*/
    0x00,

    /*ac_thresh*/
    0x1b3
};    

static ISP_CMOS_DEMOSAIC_S g_stIspDemosaicFSWDR =
{
    /* bvalid */
    1,
    
    /*vh_slope*/
    0xdc,

    /*aa_slope*/
    0xc8,

    /*va_slope*/
    0xb9,

    /*uu_slope*/
    0xa8,

    /*sat_slope*/
    0x5d,

    /*ac_slope*/
    0xa0,
    
    /*fc_slope*/
    0x80,

    /*vh_thresh*/
    0x00,

    /*aa_thresh*/
    0x00,

    /*va_thresh*/
    0x00,

    /*uu_thresh*/
    0x08,

    /*sat_thresh*/
    0x00,

    /*ac_thresh*/
    0x1b3
};

static ISP_CMOS_RGBSHARPEN_S g_stIspRgbSharpen =
{   
    /* bvalid */   
    1,   
    
    /*lut_core*/   
    192,  
    
    /*lut_strength*/  
    127, 
    
    /*lut_magnitude*/   
    6      
};

static ISP_CMOS_GAMMA_S g_stIspGamma =
{
    /* bvalid */
    1,
#if 1 /* Normal mode */   
    {  0, 180, 320, 426, 516, 590, 660, 730, 786, 844, 896, 946, 994, 1040, 1090, 1130, 1170, 1210, 1248,
    1296, 1336, 1372, 1416, 1452, 1486, 1516, 1546, 1580, 1616, 1652, 1678, 1714, 1742, 1776, 1798, 1830,
    1862, 1886, 1912, 1940, 1968, 1992, 2010, 2038, 2062, 2090, 2114, 2134, 2158, 2178, 2202, 2222, 2246,
    2266, 2282, 2300, 2324, 2344, 2360, 2372, 2390, 2406, 2422, 2438, 2458, 2478, 2494, 2510, 2526, 2546,
    2562, 2582, 2598, 2614, 2630, 2648, 2660, 2670, 2682, 2698, 2710, 2724, 2736, 2752, 2764, 2780, 2792,
    2808, 2820, 2836, 2848, 2864, 2876, 2888, 2896, 2908, 2920, 2928, 2940, 2948, 2960, 2972, 2984, 2992,
    3004, 3014, 3028, 3036, 3048, 3056, 3068, 3080, 3088, 3100, 3110, 3120, 3128, 3140, 3148, 3160, 3168,
    3174, 3182, 3190, 3202, 3210, 3218, 3228, 3240, 3256, 3266, 3276, 3288, 3300, 3306, 3318, 3326, 3334,
    3342, 3350, 3360, 3370, 3378, 3386, 3394, 3398, 3406, 3414, 3422, 3426, 3436, 3444, 3454, 3466, 3476,
    3486, 3498, 3502, 3510, 3518, 3526, 3530, 3538, 3546, 3554, 3558, 3564, 3570, 3574, 3582, 3590, 3598,
    3604, 3610, 3618, 3628, 3634, 3640, 3644, 3652, 3656, 3664, 3670, 3678, 3688, 3696, 3700, 3708, 3712,
    3716, 3722, 3730, 3736, 3740, 3748, 3752, 3756, 3760, 3766, 3774, 3778, 3786, 3790, 3800, 3808, 3812,
    3816, 3824, 3830, 3832, 3842, 3846, 3850, 3854, 3858, 3862, 3864, 3870, 3874, 3878, 3882, 3888, 3894,
    3900, 3908, 3912, 3918, 3924, 3928, 3934, 3940, 3946, 3952, 3958, 3966, 3974, 3978, 3982, 3986, 3990,
    3994, 4002, 4006, 4010, 4018, 4022, 4032, 4038, 4046, 4050, 4056, 4062, 4072, 4076, 4084, 4090, 4095}
#else  /* Infrared or Spotlight mode */
    {  0, 120, 220, 320, 416, 512, 592, 664, 736, 808, 880, 944, 1004, 1062, 1124, 1174,
    1226, 1276, 1328, 1380, 1432, 1472, 1516, 1556, 1596, 1636, 1680, 1720, 1756, 1792,
    1828, 1864, 1896, 1932, 1968, 2004, 2032, 2056, 2082, 2110, 2138, 2162, 2190, 2218,
    2242, 2270, 2294, 2314, 2338, 2358, 2382, 2402, 2426, 2446, 2470, 2490, 2514, 2534,
    2550, 2570, 2586, 2606, 2622, 2638, 2658, 2674, 2694, 2710, 2726, 2746, 2762, 2782,
    2798, 2814, 2826, 2842, 2854, 2870, 2882, 2898, 2910, 2924, 2936, 2952, 2964, 2980,
    2992, 3008, 3020, 3036, 3048, 3064, 3076, 3088, 3096, 3108, 3120, 3128, 3140, 3152,
    3160, 3172, 3184, 3192, 3204, 3216, 3224, 3236, 3248, 3256, 3268, 3280, 3288, 3300,
    3312, 3320, 3332, 3340, 3348, 3360, 3368, 3374, 3382, 3390, 3402, 3410, 3418, 3426,
    3434, 3446, 3454, 3462, 3470, 3478, 3486, 3498, 3506, 3514, 3522, 3530, 3542, 3550,
    3558, 3566, 3574, 3578, 3586, 3594, 3602, 3606, 3614, 3622, 3630, 3634, 3642, 3650,
    3658, 3662, 3670, 3678, 3686, 3690, 3698, 3706, 3710, 3718, 3722, 3726, 3734, 3738,
    3742, 3750, 3754, 3760, 3764, 3768, 3776, 3780, 3784, 3792, 3796, 3800, 3804, 3812,
    3816, 3820, 3824, 3832, 3836, 3840, 3844, 3848, 3856, 3860, 3864, 3868, 3872, 3876,
    3880, 3884, 3892, 3896, 3900, 3904, 3908, 3912, 3916, 3920, 3924, 3928, 3932, 3936,
    3940, 3944, 3948, 3952, 3956, 3960, 3964, 3968, 3972, 3972, 3976, 3980, 3984, 3988,
    3992, 3996, 4000, 4004, 4008, 4012, 4016, 4020, 4024, 4028, 4032, 4032, 4036, 4040,
    4044, 4048, 4052, 4056, 4056, 4060, 4064, 4068, 4072, 4072, 4076, 4080, 4084, 4086,
    4088, 4092, 4095} 
#endif

};

static ISP_CMOS_GAMMA_S g_stIspGammaFSWDR =
{
    /* bvalid */
    1,
    
    {
    #if 0
        /* low contrast */
        0,   1,   2,   4,   8,  12,  17,  23,  30,  38,  47,  57,  68,  79,  92, 105, 120, 133, 147, 161, 176, 192,
        209, 226, 243, 260, 278, 296, 315, 333, 351, 370, 390, 410, 431, 453, 474, 494, 515, 536, 558, 580, 602,
        623, 644, 665, 686, 708, 730, 751, 773, 795, 818, 840, 862, 884, 907, 929, 951, 974, 998,1024,1051,1073,
        1096,1117,1139,1159,1181,1202,1223,1243,1261,1275,1293,1313,1332,1351,1371,1389,1408,1427,1446,1464,1482,
        1499,1516,1533,1549,1567,1583,1600,1616,1633,1650,1667,1683,1700,1716,1732,1749,1766,1782,1798,1815,1831,
        1847,1863,1880,1896,1912,1928,1945,1961,1977,1993,2009,2025,2041,2057,2073,2089,2104,2121,2137,2153,2168,
        2184,2200,2216,2231,2248,2263,2278,2294,2310,2326,2341,2357,2373,2388,2403,2419,2434,2450,2466,2481,2496,
        2512,2527,2543,2558,2573,2589,2604,2619,2635,2650,2665,2680,2696,2711,2726,2741,2757,2771,2787,2801,2817,
        2832,2847,2862,2877,2892,2907,2922,2937,2952,2967,2982,2997,3012,3027,3041,3057,3071,3086,3101,3116,3130,
        3145,3160,3175,3190,3204,3219,3234,3248,3263,3278,3293,3307,3322,3337,3351,3365,3380,3394,3409,3424,3438,
        3453,3468,3482,3497,3511,3525,3540,3554,3569,3584,3598,3612,3626,3641,3655,3670,3684,3699,3713,3727,3742,
        3756,3770,3784,3799,3813,3827,3841,3856,3870,3884,3898,3912,3927,3941,3955,3969,3983,3997,4011,4026,4039,
        4054,4068,4082,4095

        /* higher  contrast */
        0,1,2,4,8,12,17,23,30,38,47,57,68,79,92,105,120,133,147,161,176,192,209,226,243,260,278,296,317,340,365,
        390,416,440,466,491,517,538,561,584,607,631,656,680,705,730,756,784,812,835,858,882,908,934,958,982,1008,
        1036,1064,1092,1119,1143,1167,1192,1218,1243,1269,1296,1323,1351,1379,1408,1434,1457,1481,1507,1531,1554,
        1579,1603,1628,1656,1683,1708,1732,1756,1780,1804,1829,1854,1877,1901,1926,1952,1979,2003,2024,2042,2062,
        2084,2106,2128,2147,2168,2191,2214,2233,2256,2278,2296,2314,2335,2352,2373,2391,2412,2431,2451,2472,2492,
        2513,2531,2547,2566,2581,2601,2616,2632,2652,2668,2688,2705,2721,2742,2759,2779,2796,2812,2826,2842,2857,
        2872,2888,2903,2920,2934,2951,2967,2983,3000,3015,3033,3048,3065,3080,3091,3105,3118,3130,3145,3156,3171,
        3184,3197,3213,3224,3240,3252,3267,3281,3295,3310,3323,3335,3347,3361,3372,3383,3397,3409,3421,3432,3447,
        3459,3470,3482,3497,3509,3521,3534,3548,3560,3572,3580,3592,3602,3613,3625,3633,3646,3657,3667,3679,3688,
        3701,3709,3719,3727,3736,3745,3754,3764,3773,3781,3791,3798,3806,3816,3823,3833,3840,3847,3858,3865,3872,
        3879,3888,3897,3904,3911,3919,3926,3933,3940,3948,3955,3962,3970,3973,3981,3988,3996,4003,4011,4018,4026,
        4032,4037,4045,4053,4057,4064,4072,4076,4084,4088,4095
    #endif
        /* mid contrast */
        0,11,22,33,45,57,70,82,95,109,122,136,151,165,180,195,211,226,241,256,272,288,304,321,337,353,370,387,
        404,422,439,457,476,494,512,531,549,567,585,603,621,639,658,676,694,713,731,750,769,793,816,839,863,887,
        910,933,958,983,1008,1033,1059,1084,1109,1133,1157,1180,1204,1228,1252,1277,1301,1326,1348,1366,1387,
        1410,1432,1453,1475,1496,1518,1542,1565,1586,1607,1628,1648,1669,1689,1711,1730,1751,1771,1793,1815,1835,
        1854,1871,1889,1908,1928,1947,1965,1983,2003,2023,2040,2060,2079,2096,2113,2132,2149,2167,2184,2203,2220,
        2238,2257,2275,2293,2310,2326,2344,2359,2377,2392,2408,2426,2442,2460,2477,2492,2510,2527,2545,2561,2577,
        2592,2608,2623,2638,2654,2669,2685,2700,2716,2732,2748,2764,2779,2796,2811,2827,2842,2855,2870,2884,2898,
        2913,2926,2941,2955,2969,2985,2998,3014,3027,3042,3057,3071,3086,3100,3114,3127,3142,3155,3168,3182,3196,
        3209,3222,3237,3250,3264,3277,3292,3305,3319,3332,3347,3360,3374,3385,3398,3411,3424,3437,3448,3462,3475,
        3487,3501,3513,3526,3537,3550,3561,3573,3585,3596,3609,3621,3632,3644,3655,3666,3678,3689,3701,3712,3723,
        3735,3746,3757,3767,3779,3791,3802,3812,3823,3834,3845,3855,3866,3877,3888,3899,3907,3919,3929,3940,3951,
        3962,3973,3984,3994,4003,4014,4025,4034,4045,4056,4065,4076,4085,4095
    }

};

static ISP_CMOS_GAMMAFE_S g_stGammafeFSWDR = 
{
    /* bvalid */
    1,

    /* gamma_fe0 */
    {
        0, 38406, 39281, 40156, 41031, 41907, 42782, 43657, 44532, 45407, 46282, 47158, 48033, 48908, 49783, 
        50658, 51533, 52409, 53284, 54159, 55034, 55909, 56784, 57660, 58535, 59410, 60285, 61160, 62035, 62911,
        63786, 64661, 65535
    },

    /* gamma_fe1 */
    {
        0, 72, 145, 218, 293, 369, 446, 524, 604, 685, 767, 851, 937, 1024, 1113, 1204, 1297, 1391, 1489, 1590,
            1692, 1798, 1907, 2020, 2136, 2258, 2383, 2515, 2652, 2798, 2952, 3116, 3295, 3490, 3708, 3961, 4272,
            4721, 5954, 6407, 6719, 6972, 7190, 7386, 7564, 7729, 7884, 8029, 8167, 8298, 8424, 8545, 8662, 8774,
            8883, 8990, 9092, 9192, 9289, 9385, 9478, 9569, 9658, 9745, 9831, 9915, 9997, 10078, 10158, 10236,
            10313, 10389, 10463, 10538, 10610, 10682, 10752, 10823, 10891, 10959, 11026, 11094, 11159, 11224,
            11289, 11352, 11415, 11477, 11539, 11600, 11660, 11720, 11780, 11838, 11897, 11954, 12012, 12069, 
            12125, 12181, 12236, 12291, 12346, 12399, 12453, 12507, 12559, 12612, 12664, 12716, 12768, 12818,
            12869, 12919, 12970, 13020, 13069, 13118, 13166, 13215, 13263, 13311, 13358, 13405, 13453, 13500, 
            13546, 13592, 13638, 13684, 13730, 13775, 13820, 13864, 13909, 13953, 13997, 14041, 14085, 14128,
            14172, 14214, 14257, 14299, 14342, 14384, 14426, 14468, 14509, 14551, 14592, 16213, 17654, 18942,
            20118, 21208, 22227, 23189, 24101, 24971, 25804, 26603, 27373, 28118, 28838, 29538, 30219, 30881, 
            31527, 32156, 32772, 33375, 33964, 34541, 35107, 35663, 36208, 36745, 37272, 37790, 38301, 38803, 
            39298, 39785, 40267, 40741, 41210, 41672, 42128, 42580, 43026, 43466, 43901, 44332, 44757, 45179, 
            45596, 46008, 46417, 46821, 47222, 47619, 48011, 48400, 48785, 49168, 49547, 49924, 50296, 50666, 
            51033, 51397, 51758, 52116, 52472, 52825, 53175, 53522, 53868, 54211, 54551, 54889, 55225, 55558, 
            55889, 56218, 56545, 56870, 57193, 57514, 57833, 58150, 58465, 58778, 59090, 59399, 59708, 60014, 
            60318, 60621, 60922, 61222, 61520, 61816, 62111, 62403, 62695, 62985, 63275, 63562, 63848, 64132, 
            64416, 64698, 64978, 65258, 65535
    }
        

};

HI_U32 cmos_get_isp_default(ISP_CMOS_DEFAULT_S *pstDef)
{
    if (HI_NULL == pstDef)
    {
        printf("null pointer when get isp default value!\n");
        return -1;
    }

    memset(pstDef, 0, sizeof(ISP_CMOS_DEFAULT_S));

    switch (genSensorMode)
    {
        default:
        case WDR_MODE_NONE:
            pstDef->stDrc.bEnable               = HI_FALSE;
            pstDef->stDrc.u32BlackLevel         = 0x00;
            pstDef->stDrc.u32WhiteLevel         = 0x4FF; 
            pstDef->stDrc.u32SlopeMax           = 0x30;
            pstDef->stDrc.u32SlopeMin           = 0x00;
            pstDef->stDrc.u32VarianceSpace      = 0x04;
            pstDef->stDrc.u32VarianceIntensity  = 0x01;
            pstDef->stDrc.u32Asymmetry          = 0x14;
            pstDef->stDrc.u32BrightEnhance      = 0xC8;

            memcpy(&pstDef->stNoiseTbl, &g_stIspNoiseTable, sizeof(ISP_CMOS_NOISE_TABLE_S));            
            memcpy(&pstDef->stAgcTbl, &g_stIspAgcTable, sizeof(ISP_CMOS_AGC_TABLE_S));
            memcpy(&pstDef->stDemosaic, &g_stIspDemosaic, sizeof(ISP_CMOS_DEMOSAIC_S));
            memcpy(&pstDef->stGamma, &g_stIspGamma, sizeof(ISP_CMOS_GAMMA_S));
            memcpy(&pstDef->stRgbSharpen, &g_stIspRgbSharpen, sizeof(ISP_CMOS_RGBSHARPEN_S));
        break;
        case WDR_MODE_2To1_LINE:
        case WDR_MODE_2To1_FRAME:
        case WDR_MODE_2To1_FRAME_FULL_RATE:   
            pstDef->stDrc.bEnable               = HI_TRUE;
            pstDef->stDrc.u32BlackLevel         = 0x00;
            pstDef->stDrc.u32WhiteLevel         = 0xFFF; 
            pstDef->stDrc.u32SlopeMax           = 0x38;
            pstDef->stDrc.u32SlopeMin           = 0xC0;
            pstDef->stDrc.u32VarianceSpace      = 0x06;
            pstDef->stDrc.u32VarianceIntensity  = 0x08;
            pstDef->stDrc.u32Asymmetry          = 0x14;
            pstDef->stDrc.u32BrightEnhance      = 0xC8;

            memcpy(&pstDef->stNoiseTbl, &g_stIspNoiseTableFSWDR, sizeof(ISP_CMOS_NOISE_TABLE_S));            
            memcpy(&pstDef->stAgcTbl, &g_stIspAgcTableFSWDR, sizeof(ISP_CMOS_AGC_TABLE_S));
            memcpy(&pstDef->stDemosaic, &g_stIspDemosaicFSWDR, sizeof(ISP_CMOS_DEMOSAIC_S));
            memcpy(&pstDef->stRgbSharpen, &g_stIspRgbSharpen, sizeof(ISP_CMOS_RGBSHARPEN_S));        
            memcpy(&pstDef->stGamma, &g_stIspGammaFSWDR, sizeof(ISP_CMOS_GAMMA_S));
            memcpy(&pstDef->stGammafe, &g_stGammafeFSWDR, sizeof(ISP_CMOS_GAMMAFE_S));
        break;
    }

    pstDef->stSensorMaxResolution.u32MaxWidth  = 1920;
    pstDef->stSensorMaxResolution.u32MaxHeight = 1080;

    return 0;
}


HI_U32 cmos_get_isp_black_level(ISP_CMOS_BLACK_LEVEL_S *pstBlackLevel)
{
    HI_S32  i;
    
    if (HI_NULL == pstBlackLevel)
    {
        printf("null pointer when get isp black level value!\n");
        return -1;
    }

    /* Don't need to update black level when iso change */
    pstBlackLevel->bUpdate = HI_FALSE;

    if (WDR_MODE_NONE == genSensorMode)
    {
        for (i=0; i<4; i++)
        {
            pstBlackLevel->au16BlackLevel[i] = 0xF0;    // 240
        }
    }
    else
    {
        pstBlackLevel->au16BlackLevel[0] = 0xEF;
        pstBlackLevel->au16BlackLevel[1] = 0xEF;
        pstBlackLevel->au16BlackLevel[2] = 0xEF;
        pstBlackLevel->au16BlackLevel[3] = 0xEF;
    }

    return 0;    
}

HI_VOID cmos_set_pixel_detect(HI_BOOL bEnable)
{
    HI_U32 u32FullLines_5Fps, u32MaxIntTime_5Fps;
    if (WDR_MODE_2To1_LINE == genSensorMode)
    {   
        return;
    }
    else
    {
        if(SENSOR_IMX290_1080P_60FPS_MODE == gu8SensorImageMode)
        {
            u32FullLines_5Fps = (VMAX_IMX290_1080P60_LINE * 60) / 5;
        }
        else  if(SENSOR_IMX290_1080P_30FPS_MODE == gu8SensorImageMode)
        {
            u32FullLines_5Fps = (VMAX_IMX290_1080P30_LINE* 30) / 5;
        }
        else if(SENSOR_IMX290_720P_120FPS_MODE == gu8SensorImageMode)
        {
            u32FullLines_5Fps = (VMAX_IMX290_720P120_LINE * 120) / 5;
        }
        else
        {
            return;
        }
    }

    u32FullLines_5Fps = (u32FullLines_5Fps > 0x3FFFF) ? 0x3FFFF : u32FullLines_5Fps;
    u32MaxIntTime_5Fps =  u32FullLines_5Fps - 4;

    if (bEnable) /* setup for ISP pixel calibration mode */
    {
        sensor_write_register (GAIN_ADDR,0x00);
        
        sensor_write_register (VMAX_ADDR, u32FullLines_5Fps & 0xFF); 
        sensor_write_register (VMAX_ADDR + 1, (u32FullLines_5Fps & 0xFF00) >> 8); 
        sensor_write_register (VMAX_ADDR + 2, (u32FullLines_5Fps & 0x30000) >> 16);

        sensor_write_register (SHS1_ADDR, u32MaxIntTime_5Fps & 0xFF);
        sensor_write_register (SHS1_ADDR + 1,  (u32MaxIntTime_5Fps & 0xFF00) >> 8); 
        sensor_write_register (SHS1_ADDR + 2, (u32MaxIntTime_5Fps & 0x30000) >> 16); 
          
    }
    else /* setup for ISP 'normal mode' */
    {
        gu32FullLinesStd = (gu32FullLinesStd > 0x1FFFF) ? 0x1FFFF : gu32FullLinesStd;
        sensor_write_register (VMAX_ADDR, gu32FullLines & 0xFF); 
        sensor_write_register (VMAX_ADDR + 1, (gu32FullLines & 0xFF00) >> 8); 
        sensor_write_register (VMAX_ADDR + 2, (gu32FullLines & 0x30000) >> 16);
        bInit = HI_FALSE;
    }

    return;
}

HI_VOID cmos_set_wdr_mode(HI_U8 u8Mode)
{
    bInit = HI_FALSE;
    
    switch(u8Mode)
    {
        case WDR_MODE_NONE:
            genSensorMode = WDR_MODE_NONE;
            if(SENSOR_IMX290_1080P_30FPS_MODE == gu8SensorImageMode)
            {            
                gu32FullLinesStd = VMAX_IMX290_1080P30_LINE;
            }
            else if(SENSOR_IMX290_1080P_60FPS_MODE == gu8SensorImageMode)
            {
                gu32FullLinesStd = VMAX_IMX290_1080P60_LINE;
            }
            else if(SENSOR_IMX290_720P_120FPS_MODE == gu8SensorImageMode)
            {
                gu32FullLinesStd = VMAX_IMX290_720P120_LINE;
            }
            
            printf("linear mode\n");
        break;

        case WDR_MODE_2To1_LINE:
            genSensorMode = WDR_MODE_2To1_LINE;

            if(SENSOR_IMX290_1080P_30FPS_MODE == gu8SensorImageMode)
            {
                gu32FullLinesStd = VMAX_IMX290_1080P30_WDR * 2;
                gu32BRL = 1109;
            }
            else if(SENSOR_IMX290_720P_60FPS_MODE == gu8SensorImageMode)
            {
                gu32FullLinesStd = VMAX_IMX290_720P60_WDR * 2;
                gu32BRL = 735;
            }
            else
            {
            }
            
            printf("2to1 line WDR mode\n");
        break;

        case WDR_MODE_2To1_FRAME:           //half rate
            genSensorMode = WDR_MODE_2To1_FRAME;
            
            if(SENSOR_IMX290_1080P_30FPS_MODE == gu8SensorImageMode)
            {            
                gu32FullLinesStd = VMAX_IMX290_1080P30_LINE;
            }
            else if(SENSOR_IMX290_1080P_60FPS_MODE == gu8SensorImageMode)
            {
                gu32FullLinesStd = VMAX_IMX290_1080P60_LINE;
            }
            else if(SENSOR_IMX290_720P_120FPS_MODE == gu8SensorImageMode)
            {
                gu32FullLinesStd = VMAX_IMX290_720P120_LINE;
            }
            
            printf("2to1 half frame WDR mode\n");
        break;

        case WDR_MODE_2To1_FRAME_FULL_RATE:
            genSensorMode = WDR_MODE_2To1_FRAME_FULL_RATE;

            if(SENSOR_IMX290_1080P_30FPS_MODE == gu8SensorImageMode)
            {            
                gu32FullLinesStd = VMAX_IMX290_1080P30_LINE;
            }
            else if(SENSOR_IMX290_1080P_60FPS_MODE == gu8SensorImageMode)
            {
                gu32FullLinesStd = VMAX_IMX290_1080P60_LINE;
            }
            else if(SENSOR_IMX290_720P_120FPS_MODE == gu8SensorImageMode)
            {
                gu32FullLinesStd = VMAX_IMX290_720P120_LINE;
            }
            
            printf("2to1 full frame WDR mode\n");
        break;
        
        

        default:
            printf("NOT support this mode!\n");
            return;
        break;
    }

    gu32FullLines = gu32FullLinesStd;
    gu32PreFullLines = gu32FullLines;
    memset(au32WDRIntTime, 0, sizeof(au32WDRIntTime));
    
    return;
}

HI_U32 cmos_get_sns_regs_info(ISP_SNS_REGS_INFO_S *pstSnsRegsInfo)
{
    HI_S32 i;

    if (HI_FALSE == bInit)
    {
        g_stSnsRegsInfo.enSnsType = ISP_SNS_SSP_TYPE;
        g_stSnsRegsInfo.u8Cfg2ValidDelayMax = 2;        
        g_stSnsRegsInfo.u32RegNum = 8;
        
        if (WDR_MODE_2To1_LINE == genSensorMode)
        {
            g_stSnsRegsInfo.u32RegNum += 6;
            g_stSnsRegsInfo.u8Cfg2ValidDelayMax = 2;
        }
        else if ((WDR_MODE_2To1_FRAME_FULL_RATE == genSensorMode) || (WDR_MODE_2To1_FRAME == genSensorMode))
        {
            g_stSnsRegsInfo.u32RegNum += 3;
            g_stSnsRegsInfo.u8Cfg2ValidDelayMax = 2;
        }
        
        for (i=0; i<g_stSnsRegsInfo.u32RegNum; i++)
        {
            g_stSnsRegsInfo.astSspData[i].bUpdate = HI_TRUE;
            g_stSnsRegsInfo.astSspData[i].u32DevAddr = 0x02;
            g_stSnsRegsInfo.astSspData[i].u32DevAddrByteNum = 1;
            g_stSnsRegsInfo.astSspData[i].u32RegAddrByteNum = 1;
            g_stSnsRegsInfo.astSspData[i].u32DataByteNum = 1;
        }        
        
        g_stSnsRegsInfo.astSspData[0].u8DelayFrmNum = 0;
        g_stSnsRegsInfo.astSspData[0].u32RegAddr = SHS1_ADDR;       
        g_stSnsRegsInfo.astSspData[1].u8DelayFrmNum = 0;
        g_stSnsRegsInfo.astSspData[1].u32RegAddr = SHS1_ADDR + 1;        
        g_stSnsRegsInfo.astSspData[2].u8DelayFrmNum = 0;
        g_stSnsRegsInfo.astSspData[2].u32RegAddr = SHS1_ADDR + 2;   
        
        g_stSnsRegsInfo.astSspData[3].u8DelayFrmNum = 0;       //make shutter and gain effective at the same time
        g_stSnsRegsInfo.astSspData[3].u32RegAddr = GAIN_ADDR;  //gain     
        g_stSnsRegsInfo.astSspData[4].u8DelayFrmNum = 1;       //make shutter and gain effective at the same time
        g_stSnsRegsInfo.astSspData[4].u32RegAddr = HCG_ADDR;   //gain   
        
        
        g_stSnsRegsInfo.astSspData[5].u8DelayFrmNum = 0;
        g_stSnsRegsInfo.astSspData[5].u32RegAddr = VMAX_ADDR;
        g_stSnsRegsInfo.astSspData[6].u8DelayFrmNum = 0;
        g_stSnsRegsInfo.astSspData[6].u32RegAddr = VMAX_ADDR + 1;
        g_stSnsRegsInfo.astSspData[7].u8DelayFrmNum = 0;
        g_stSnsRegsInfo.astSspData[7].u32RegAddr = VMAX_ADDR + 2;
      
        if (WDR_MODE_2To1_LINE == genSensorMode)
        {
            g_stSnsRegsInfo.astSspData[5].u8DelayFrmNum = 0;
            g_stSnsRegsInfo.astSspData[5].u32RegAddr = SHS2_ADDR;
            g_stSnsRegsInfo.astSspData[6].u8DelayFrmNum = 0;
            g_stSnsRegsInfo.astSspData[6].u32RegAddr = SHS2_ADDR + 1;
            g_stSnsRegsInfo.astSspData[7].u8DelayFrmNum = 0;
            g_stSnsRegsInfo.astSspData[7].u32RegAddr = SHS2_ADDR + 2;
            
            g_stSnsRegsInfo.astSspData[8].u8DelayFrmNum = 1;
            g_stSnsRegsInfo.astSspData[8].u32RegAddr = VMAX_ADDR;
            g_stSnsRegsInfo.astSspData[9].u8DelayFrmNum = 1;
            g_stSnsRegsInfo.astSspData[9].u32RegAddr = VMAX_ADDR + 1;      
            g_stSnsRegsInfo.astSspData[10].u8DelayFrmNum = 1;
            g_stSnsRegsInfo.astSspData[10].u32RegAddr = VMAX_ADDR + 2;

            g_stSnsRegsInfo.astSspData[11].u8DelayFrmNum = 0;
            g_stSnsRegsInfo.astSspData[11].u32RegAddr = RHS1_ADDR;          
            g_stSnsRegsInfo.astSspData[12].u8DelayFrmNum = 0;       
            g_stSnsRegsInfo.astSspData[12].u32RegAddr = RHS1_ADDR + 1;
            g_stSnsRegsInfo.astSspData[13].u8DelayFrmNum = 0;
            g_stSnsRegsInfo.astSspData[13].u32RegAddr = RHS1_ADDR + 2;

            if(gbFPSUp)
            {
                g_stSnsRegsInfo.astSspData[11].u8DelayFrmNum = 1;
                g_stSnsRegsInfo.astSspData[12].u8DelayFrmNum = 1;
                g_stSnsRegsInfo.astSspData[13].u8DelayFrmNum = 1;
            }


        }
        else if ((WDR_MODE_2To1_FRAME == genSensorMode) || (WDR_MODE_2To1_FRAME_FULL_RATE == genSensorMode))
        {
            g_stSnsRegsInfo.astSspData[5].u8DelayFrmNum = 1;
            g_stSnsRegsInfo.astSspData[5].u32RegAddr = SHS1_ADDR;
            g_stSnsRegsInfo.astSspData[6].u8DelayFrmNum = 1;
            g_stSnsRegsInfo.astSspData[6].u32RegAddr = SHS1_ADDR + 1;
            g_stSnsRegsInfo.astSspData[7].u8DelayFrmNum = 1;
            g_stSnsRegsInfo.astSspData[7].u32RegAddr = SHS1_ADDR + 2;

            g_stSnsRegsInfo.astSspData[8].u8DelayFrmNum = 0;
            g_stSnsRegsInfo.astSspData[8].u32RegAddr = VMAX_ADDR;
            g_stSnsRegsInfo.astSspData[9].u8DelayFrmNum = 0;       
            g_stSnsRegsInfo.astSspData[9].u32RegAddr = VMAX_ADDR + 1;
            g_stSnsRegsInfo.astSspData[10].u8DelayFrmNum = 0;
            g_stSnsRegsInfo.astSspData[10].u32RegAddr = VMAX_ADDR + 2;
        }

        if(gbFPSUp)
        {
            gbFPSUp = HI_FALSE;
            bInit = HI_FALSE;
        }
        else
        {
            bInit = HI_TRUE;
        }
    }
    else
    {
        for (i=0; i<g_stSnsRegsInfo.u32RegNum; i++)
        {
            if (g_stSnsRegsInfo.astSspData[i].u32Data == g_stPreSnsRegsInfo.astSspData[i].u32Data)
            {
                g_stSnsRegsInfo.astSspData[i].bUpdate = HI_FALSE;
            }
            else
            {
                g_stSnsRegsInfo.astSspData[i].bUpdate = HI_TRUE;
            }
        }

        if ((WDR_MODE_2To1_FRAME == genSensorMode) || (WDR_MODE_2To1_FRAME_FULL_RATE == genSensorMode))
        {
            g_stSnsRegsInfo.astSspData[0].bUpdate = HI_TRUE;
            g_stSnsRegsInfo.astSspData[1].bUpdate = HI_TRUE;
            g_stSnsRegsInfo.astSspData[2].bUpdate = HI_TRUE;
            g_stSnsRegsInfo.astSspData[5].bUpdate = HI_TRUE;
            g_stSnsRegsInfo.astSspData[6].bUpdate = HI_TRUE;
            g_stSnsRegsInfo.astSspData[7].bUpdate = HI_TRUE;
        }
        
    }

    #if 0
    //to avoid bad frame at fps up
    if(WDR_MODE_2To1_LINE== genSensorMode)
    {
        static HI_BOOL bVMAXDelay = HI_FALSE;
        if(gbFPSUp)
        {
            gbFPSUp = HI_FALSE;
            bVMAXDelay = HI_TRUE;
            g_stSnsRegsInfo.astSspData[8].u8DelayFrmNum = 1;
            g_stSnsRegsInfo.astSspData[9].u8DelayFrmNum = 1;
            g_stSnsRegsInfo.astSspData[10].u8DelayFrmNum = 1;
        }
        else if(bVMAXDelay)
        {
            g_stSnsRegsInfo.astSspData[8].u8DelayFrmNum = 0;
            g_stSnsRegsInfo.astSspData[8].bUpdate = HI_TRUE;
            g_stSnsRegsInfo.astSspData[9].u8DelayFrmNum = 0;
            g_stSnsRegsInfo.astSspData[9].bUpdate = HI_TRUE;
            g_stSnsRegsInfo.astSspData[10].u8DelayFrmNum = 0;
            g_stSnsRegsInfo.astSspData[10].bUpdate = HI_TRUE;

            bVMAXDelay = HI_FALSE;
        }
    }
    #endif
    
    if (HI_NULL == pstSnsRegsInfo)
    {
        printf("null pointer when get sns reg info!\n");
        return -1;
    }


    
    memcpy(pstSnsRegsInfo, &g_stSnsRegsInfo, sizeof(ISP_SNS_REGS_INFO_S)); 
    memcpy(&g_stPreSnsRegsInfo, &g_stSnsRegsInfo, sizeof(ISP_SNS_REGS_INFO_S)); 

    
    gu32PreFullLines = gu32FullLines;

    return 0;
}

static HI_S32 cmos_set_image_mode(ISP_CMOS_SENSOR_IMAGE_MODE_S *pstSensorImageMode)
{
    HI_U8 u8SensorImageMode = gu8SensorImageMode;
    
    bInit = HI_FALSE;    

    if (HI_NULL == pstSensorImageMode )
    {
        printf("null pointer when set image mode\n");
        return -1;
    }
    if((pstSensorImageMode->u16Width <= 1280) && (pstSensorImageMode->u16Height <= 720)
        && (pstSensorImageMode->f32Fps <= 120))
    {
        if(WDR_MODE_2To1_LINE == genSensorMode)
        {
            if(pstSensorImageMode->f32Fps <= 60)
            {
                u8SensorImageMode = SENSOR_IMX290_720P_60FPS_MODE;
                gu32FullLinesStd = VMAX_IMX290_720P60_WDR * 2;
                gu8HCGReg = 0x0;
                gu32BRL = 735;
            }
            else
            {
                printf("Not support! Width:%d, Height:%d, Fps:%f\n", 
                    pstSensorImageMode->u16Width, 
                    pstSensorImageMode->u16Height,
                    pstSensorImageMode->f32Fps);

                return -1;                
            }
        }
        else
        {
            u8SensorImageMode = SENSOR_IMX290_720P_120FPS_MODE;
            gu32FullLinesStd = VMAX_IMX290_720P120_LINE;
            gu8HCGReg = 0x0;
        }
        
    }
    else if((pstSensorImageMode->u16Width <= 1920) && (pstSensorImageMode->u16Height <= 1080)
        && (pstSensorImageMode->f32Fps <= 60))
    {
        if(pstSensorImageMode->f32Fps <= 30)
        {
            u8SensorImageMode = SENSOR_IMX290_1080P_30FPS_MODE;
            
            if(WDR_MODE_2To1_LINE == genSensorMode)
            {
                gu32FullLinesStd = VMAX_IMX290_1080P30_WDR * 2;
                gu8HCGReg = 0x1;
                gu32BRL = 1109;
            } 
            //when current mode is 60fps linear/fs wdr, won't go into 30fps linear mode
            else if(SENSOR_IMX290_1080P_60FPS_MODE == gu8SensorImageMode)
            {
                u8SensorImageMode = SENSOR_IMX290_1080P_60FPS_MODE;
                gu8HCGReg = 0x1;
            }
            else
            {
                gu32FullLinesStd = VMAX_IMX290_1080P30_LINE;
                gu8HCGReg = 0x2;
            }
        }
        else
        {
            if(WDR_MODE_2To1_LINE == genSensorMode)
            {
                printf("Not support! Width:%d, Height:%d, Fps:%f\n", 
                pstSensorImageMode->u16Width, 
                    pstSensorImageMode->u16Height,
                    pstSensorImageMode->f32Fps);

                return -1;
            }

            u8SensorImageMode = SENSOR_IMX290_1080P_60FPS_MODE;
            gu32FullLinesStd = VMAX_IMX290_1080P60_LINE;
            gu8HCGReg = 0x1;
            
        }
    }
    else
    {
        printf("Not support! Width:%d, Height:%d, Fps:%f\n", 
            pstSensorImageMode->u16Width, 
            pstSensorImageMode->u16Height,
            pstSensorImageMode->f32Fps);

        return -1;
    }

    if ((HI_TRUE == bSensorInit) && (u8SensorImageMode == gu8SensorImageMode))
    {
        /* Don't need to switch SensorImageMode */
        return -1;
    }

    gu8SensorImageMode = u8SensorImageMode;
    gu32FullLines = gu32FullLinesStd;
    gu32PreFullLines = gu32FullLines;
    memset(au32WDRIntTime, 0, sizeof(au32WDRIntTime));
    

    return 0;
    
}


int  sensor_set_inifile_path(const char *pcPath)
{
    memset(pcName, 0, sizeof(pcName));
        
    if (HI_NULL == pcPath)
    {        
        strncat(pcName, "configs/", strlen("configs/"));
        strncat(pcName, CMOS_CFG_INI, sizeof(CMOS_CFG_INI));
    }
    else
    {
        if(strlen(pcPath) > (PATHLEN_MAX - 30))
        {
            printf("Set inifile path is larger PATHLEN_MAX!\n");
            return -1;        
        }
        
        strncat(pcName, pcPath, strlen(pcPath));
        strncat(pcName, CMOS_CFG_INI, sizeof(CMOS_CFG_INI));
    }
    
    return 0;
}



HI_VOID sensor_global_init()
{
    bInit = HI_FALSE;
    bSensorInit = HI_FALSE;
    gbVMAXDelay = HI_FALSE;
    gbFPSUp = HI_FALSE;
    gu32FullLinesStd = VMAX_IMX290_1080P30_LINE;
    gu32FullLines = VMAX_IMX290_1080P30_LINE;     
    gu32PreFullLines = VMAX_IMX290_1080P30_LINE;
    genSensorMode = WDR_MODE_NONE;
    gu8SensorImageMode = SENSOR_IMX290_1080P_30FPS_MODE;
    
    memset(&g_stSnsRegsInfo, 0, sizeof(ISP_SNS_REGS_INFO_S));
    memset(&g_stPreSnsRegsInfo, 0, sizeof(ISP_SNS_REGS_INFO_S));

}

HI_S32 cmos_init_sensor_exp_function(ISP_SENSOR_EXP_FUNC_S *pstSensorExpFunc)
{
    memset(pstSensorExpFunc, 0, sizeof(ISP_SENSOR_EXP_FUNC_S));

    pstSensorExpFunc->pfn_cmos_sensor_init                  = sensor_init;
    pstSensorExpFunc->pfn_cmos_sensor_exit                  = sensor_exit;
    pstSensorExpFunc->pfn_cmos_sensor_global_init           = sensor_global_init;
    pstSensorExpFunc->pfn_cmos_set_image_mode               = cmos_set_image_mode;
    pstSensorExpFunc->pfn_cmos_set_wdr_mode                 = cmos_set_wdr_mode;
    pstSensorExpFunc->pfn_cmos_get_isp_default              = cmos_get_isp_default;
    pstSensorExpFunc->pfn_cmos_get_isp_black_level          = cmos_get_isp_black_level;
    pstSensorExpFunc->pfn_cmos_set_pixel_detect             = cmos_set_pixel_detect;
    pstSensorExpFunc->pfn_cmos_get_sns_reg_info             = cmos_get_sns_regs_info;

    return 0;
}

/****************************************************************************
 * callback structure                                                       *
 ****************************************************************************/

int sensor_register_callback(void)
{
    HI_S32 s32Ret;
    ALG_LIB_S stLib;
    ISP_DEV IspDev=0;
    ISP_SENSOR_REGISTER_S stIspRegister;
    AE_SENSOR_REGISTER_S  stAeRegister;
    AWB_SENSOR_REGISTER_S stAwbRegister;

    cmos_init_sensor_exp_function(&stIspRegister.stSnsExp);
    s32Ret = HI_MPI_ISP_SensorRegCallBack(IspDev, IMX290_ID, &stIspRegister);
    if (s32Ret)
    {
        printf("sensor register callback function failed!\n");
        return s32Ret;
    }
    
    stLib.s32Id = 0;
    strncpy(stLib.acLibName, HI_AE_LIB_NAME, sizeof(HI_AE_LIB_NAME));
    cmos_init_ae_exp_function(&stAeRegister.stSnsExp);
    s32Ret = HI_MPI_AE_SensorRegCallBack(IspDev, &stLib, IMX290_ID, &stAeRegister);
    if (s32Ret)
    {
        printf("sensor register callback function to ae lib failed!\n");
        return s32Ret;
    }

    stLib.s32Id = 0;
    strncpy(stLib.acLibName, HI_AWB_LIB_NAME, sizeof(HI_AWB_LIB_NAME));
    cmos_init_awb_exp_function(&stAwbRegister.stSnsExp);
    s32Ret = HI_MPI_AWB_SensorRegCallBack(IspDev, &stLib, IMX290_ID, &stAwbRegister);
    if (s32Ret)
    {
        printf("sensor register callback function to awb lib failed!\n");
        return s32Ret;
    }
    
    return 0;
}

int sensor_unregister_callback(void)
{
    HI_S32 s32Ret;
    ALG_LIB_S stLib;
    ISP_DEV IspDev=0;

    s32Ret = HI_MPI_ISP_SensorUnRegCallBack(IspDev, IMX290_ID);
    if (s32Ret)
    {
        printf("sensor unregister callback function failed!\n");
        return s32Ret;
    }
    
    stLib.s32Id = 0;
    strncpy(stLib.acLibName, HI_AE_LIB_NAME, sizeof(HI_AE_LIB_NAME));
    s32Ret = HI_MPI_AE_SensorUnRegCallBack(IspDev, &stLib, IMX290_ID);
    if (s32Ret)
    {
        printf("sensor unregister callback function to ae lib failed!\n");
        return s32Ret;
    }

    stLib.s32Id = 0;
    strncpy(stLib.acLibName, HI_AWB_LIB_NAME, sizeof(HI_AWB_LIB_NAME));
    s32Ret = HI_MPI_AWB_SensorUnRegCallBack(IspDev, &stLib, IMX290_ID);
    if (s32Ret)
    {
        printf("sensor unregister callback function to awb lib failed!\n");
        return s32Ret;
    }
    
    return 0;
}

#ifdef __cplusplus
#if __cplusplus
}
#endif
#endif /* End of #ifdef __cplusplus */

#endif 
