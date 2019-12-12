#if !defined(__IMX265_CMOS_H_)
#define __IMX265_CMOS_H_

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

#ifdef __cplusplus
#if __cplusplus
extern "C"{
#endif
#endif /* End of #ifdef __cplusplus */

#define IMX265_ID 265

#define IMX265_SENSOR_1080P_25FPS_LINEAR_MODE (1)


//#define IMX265_INCREASE_LINES (1) /* make real fps less than stand fps because NVR require*/
//#define IMX265_VMAX_1080P25_LINEAR (1125) //(1125+IMX265_INCREASE_LINES)

/****************************************************************************
 * global variables                                                            *
 ****************************************************************************/
ISP_SNS_STATE_S             g_astimx265[ISP_MAX_DEV_NUM] = {{0}};
static ISP_SNS_STATE_S     *g_apstSnsState[ISP_MAX_DEV_NUM] = {&g_astimx265[0], &g_astimx265[1]};
ISP_SNS_COMMBUS_U     g_aunImx265BusInfo[ISP_MAX_DEV_NUM] = {
    [0] = { .s8I2cDev = 0},
    [1] = { .s8I2cDev = 1}
};

/*
static ISP_FSWDR_MODE_E genFSWDRMode[ISP_MAX_DEV_NUM] = {ISP_FSWDR_NORMAL_MODE,ISP_FSWDR_NORMAL_MODE};
static HI_U32 gu32MaxTimeGetCnt[ISP_MAX_DEV_NUM] = {0,0};
*/



static HI_U32 g_au32InitExposure[ISP_MAX_DEV_NUM]  = {0};
static HI_U32 g_au32LinesPer500ms[ISP_MAX_DEV_NUM] = {0};
static HI_U16 g_au16InitWBGain[ISP_MAX_DEV_NUM][3] = {{0}};
static HI_U16 g_au16SampleRgain[ISP_MAX_DEV_NUM] = {0};
static HI_U16 g_au16SampleBgain[ISP_MAX_DEV_NUM] = {0};

extern const unsigned int imx265_i2c_addr;
extern unsigned int imx265_addr_byte;
extern unsigned int imx265_data_byte;


typedef struct hiIMX265_STATE_S
{
    HI_U8       u8Hcg;
    HI_U32      u32BRL;
    HI_U32      u32RHS1_MAX;
    HI_U32      u32RHS2_MAX;
} IMX265_STATE_S;


IMX265_STATE_S g_astimx265State[ISP_MAX_DEV_NUM] = {{0}};


extern void imx265_init(ISP_DEV IspDev);
extern void imx265_exit(ISP_DEV IspDev);
extern void imx265_standby(ISP_DEV IspDev);
extern void imx265_restart(ISP_DEV IspDev);
extern int  imx265_write_register(ISP_DEV IspDev, int addr, int data);
extern int  imx265_read_register(ISP_DEV IspDev, int addr);


#define IMX265_INCREASE_LINES (1) /* make real fps less than stand fps because NVR require*/

#define IMX265_VMAX_1080P25_LINEAR  (1125+IMX265_INCREASE_LINES)
//#define IMX265_VMAX_1080P25_LINEAR  (1536+IMX265_INCREASE_LINES)

/////////////////////////////////////////////////////////////////////////////////

/////////////////////////////////////////////////////////////////////////////////
static ISP_CMOS_AGC_TABLE_S g_stIspAgcTable =
{
    /* bvalid */
    1,

    /* snr_thresh */
    {0x08,0x0c,0x10,0x14,0x18,0x20,0x28,0x30,0x30,0x30,0x30,0x30,0x30,0x30,0x30,0x30},
               
    /* demosaic_np_offset */
    {0x0,0xa,0x12,0x1a,0x20,0x28,0x30,0x30,0x30,0x30,0x30,0x30,0x30,0x30,0x30,0x30},
        
    /* ge_strength */
    {0x55,0x55,0x55,0x55,0x55,0x55,0x37,0x37,0x37,0x37,0x37,0x37,0x37,0x37,0x37,0x37}
    
};

static ISP_CMOS_BAYER_SHARPEN_S g_stIspBayerSharpen = 
{
    /* bvalid */
    0,

    /* ShpAlgSel = 1 is Demosaic SharpenEx, else Demosaic sharpen. */ 
    0,
    
    /* sharpen_alt_d to Sharpen */
    {0x48,0x48,0x48,0x48,0x48,0x48,0x48,0x48,0x48,0x48,0x48,0x48,0x40,0x30,0x20,0x10},
        
    /* sharpen_alt_ud to Sharpen */
    {0x3a,0x38,0x36,0x34,0x32,0x30,0x30,0x30,0x28,0x24,0x24,0x20,0x20,0x20,0x10,0x10},
        
    /* demosaic_lum_thresh to Sharpen */
    {0x50,0x50,0x4c,0x48,0x42,0x3a,0x32,0x28,0x20,0x20,0x20,0x20,0x20,0x20,0x20,0x20},
        
    /* SharpenHF to SharpenEx */
    {0x30,0x30,0x30,0x30,0x30,0x30,0x2c,0x28,0x20,0x18,0x14,0x10,0x10,0x10,0x10,0x10},
        
    /* SharpenMF to SharpenEx */
    {0x30,0x30,0x30,0x30,0x28,0x20,0x20,0x20,0x20,0x20,0x10,0x10,0x10,0x10,0x10,0x10},
        
    /* SharpenLF to SharpenEx */
    {0x18,0x18,0x18,0x18,0x18,0x18,0x18,0x18,0x18,0x18,0x18,0x18,0x18,0x18,0x18,0x18},

    /* SadAmplifier to SharpenEx */
    {0x10,0x10,0x10,0x10,0x10,0x10,0x10,0x10,0x10,0x10,0x10,0x10,0x10,0x10,0x10,0x10}   

};

static ISP_CMOS_YUV_SHARPEN_S g_stIspYuvSharpen = 
{
    /* bvalid */
     0,

     /* 100,  200,    400,     800,    1600,    3200,    6400,    12800,    25600,   51200,  102400,  204800,   409600,   819200,   1638400,  3276800 */
    
     /* bEnLowLumaShoot */ 
     {0,     0,     0,     0,     0,     0,     0,     0,     0,     1,     1,     1,     1,     1,     1,     1},
     
     /* TextureSt */
     {52,    52,   52,    50,    50,    50,    46,    40,    36,    26,    20,    12,    12,     8,     8,     8},
         
     /* EdgeSt */
     {56,    56,   56,    56,    56,    56,    50,    50,    44,    44,    38,    38,    38,    20,     20,    20},      
         
     /* OverShoot */
     {64,   64,    60,    60,    56,    56,    56,    56,    56,    50,    44,    40,    40,    40,    40,    40},
        
     /* UnderShoot */
     {64,   64,    60,    60,    56,    56,    56,    56,    56,    50,    44,    40,    40,    40,    40,    40},
#if 0
     /* TextureThd */
     {10,   16,    20,    24,    32,    40,    48,    56,    64,   128,   156,    156,    156,    160,    160,   160},
        
     /* EdgeThd */
     {0,     0,     0,    10,    10,    10,    16,    32,    64,   128,   156,    156,    156,    160,    160,   160},
#else        
     /* TextureThd */
     {80,   80,    80,   80,   80,    80,   90,   90,   90,     100,    100,    100,    100,    100,    110, 110},

     /* EdgeThd */
     {0,    0,     5,    10,   10,   10,   16,   20,   30,   40,  50,    50,     60,     60,     60,   60},        
#endif    
     /* JagCtrl */
     {16,   14,    12,    10,     8,     6,     4,     4,     4,     4,     4,      4,      2,      2,      2,     2},
    
     /* SaltCtrl */
     {50,   50,    50,    90,    90,    90,    90,    90,    90,    90,     90,    90,     90,     50,     50,    50},
    
     /* PepperCtrl */
     {0,     0,      0,     20,     60,     60,     60,     80,    120,    160,    180,     180,   180,     180,    180,   180},
    
     /* DetailCtrl */
     {150,  148,   146,    144,    140,    136,    132,    130,    128,    126,    124,     122,    120,    100,     80,    80}, 
    
    /* LumaThd */
    {
        {20,    20,     20,     20,     20,     20,     20,     20,     20,     20,     20,     20,     20,     20,     20,     20}, /* LumaThd0 */
        {40,    40,     40,     40,     40,     40,     40,     40,     40,     40,     40,     40,     40,     40,     40,     40}, /* LumaThd1 */
        {65,    65,     65,     65,     65,     65,     65,     65,     65,     65,     65,     65,     65,     65,     65,     65}, /* LumaThd2 */
        {90,    90,     90,     90,     90,     90,     90,     90,     90,     90,     90,     90,     90,     90,     90,     90}  /* LumaThd3 */  
    }, 
    
    /* LumaWgt */
    {
        {160,   160,    160,    150,    140,    130,    120,    110,    100,    100,     90,     90,     80,     80,     80,     80},
        {200,   200,    200,    180,    170,    160,    150,    150,    150,    150,    120,    120,    120,    120,    120,    120},
        {240,   240,    240,    200,    200,    190,    180,    180,    180,    180,    160,    160,    160,    160,    160,    160},
        {255,   255,    255,    255,    255,    255,    255,    255,    255,    255,    255,    255,    255,    255,    255,    255},
    } 

};



static ISP_CMOS_NOISE_TABLE_S g_stIspNoiseTable =
{
    /* bvalid */
    0,

    /* nosie_profile_weight_lut */
    {
        0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xA, 0xF, 0x12, 0x15, 0x17, 0x19, 0x1A, 0x1B, 0x1D, 0x1E,
        0x1F, 0x20, 0x21, 0x21, 0x22, 0x23, 0x24, 0x24, 0x24, 0x25, 0x26, 0x26, 0x27, 0x27, 0x28, 0x28,
        0x29, 0x29, 0x29, 0x2A, 0x2A, 0x2A, 0x2B, 0x2B, 0x2B, 0x2C, 0x2C, 0x2C, 0x2D, 0x2D, 0x2D, 0x2D,
        0x2E, 0x2E, 0x2E, 0x2E, 0x2F, 0x2F, 0x2F, 0x2F, 0x30, 0x30, 0x30, 0x30, 0x30, 0x31, 0x31, 0x31,
        0x31, 0x31, 0x32, 0x32, 0x32, 0x32, 0x32, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x34, 0x34, 0x34,
        0x34, 0x34, 0x34, 0x34, 0x35, 0x35, 0x35, 0x35, 0x35, 0x35, 0x35, 0x36, 0x36, 0x36, 0x36, 0x36,
        0x36, 0x36, 0x36, 0x37, 0x37, 0x37, 0x37, 0x37, 0x37, 0x37, 0x37, 0x37, 0x38, 0x38, 0x38, 0x38,
        0x38, 0x38, 0x38, 0x38, 0x38, 0x39, 0x39, 0x39, 0x39, 0x39, 0x39, 0x39, 0x39, 0x39, 0x39
    },

    /* demosaic_weight_lut */
    {
        0x3, 0xA, 0xF, 0x12, 0x15, 0x17, 0x19, 0x1A, 0x1B, 0x1D, 0x1E, 0x1F, 0x20, 0x21, 0x21, 0x22,
        0x23, 0x24, 0x24, 0x24, 0x25, 0x26, 0x26, 0x27, 0x27, 0x28, 0x28, 0x29, 0x29, 0x29, 0x2A, 0x2A,
        0x2A, 0x2B, 0x2B, 0x2B, 0x2C, 0x2C, 0x2C, 0x2D, 0x2D, 0x2D, 0x2D, 0x2E, 0x2E, 0x2E, 0x2E, 0x2F,
        0x2F, 0x2F, 0x2F, 0x30, 0x30, 0x30, 0x30, 0x30, 0x31, 0x31, 0x31, 0x31, 0x31, 0x32, 0x32, 0x32,
        0x32, 0x32, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33, 0x34, 0x34, 0x34, 0x34, 0x34, 0x34, 0x34, 0x35,
        0x35, 0x35, 0x35, 0x35, 0x35, 0x35, 0x36, 0x36, 0x36, 0x36, 0x36, 0x36, 0x36, 0x36, 0x37, 0x37,
        0x37, 0x37, 0x37, 0x37, 0x37, 0x37, 0x37, 0x38, 0x38, 0x38, 0x38, 0x38, 0x38, 0x38, 0x38, 0x38,
        0x39, 0x39, 0x39, 0x39, 0x39, 0x39, 0x39, 0x39, 0x39, 0x39, 0x39, 0x39, 0x39, 0x39, 0x39, 0x39
    }
    
};


static ISP_CMOS_DEMOSAIC_S g_stIspDemosaic =
{
    /* bvalid */
    0,
    
    /*vh_slope*/
    128,

    /*aa_slope*/
    128,

    /*va_slope*/
    128,

    /*uu_slope*/
    0x80,

    /*sat_slope*/
    128,

    /*ac_slope*/
    128,
    
    /*fc_slope*/
    128,

    /*vh_thresh*/
    0x0,

    /*aa_thresh*/
    0x00,

    /*va_thresh*/
    0x00,

    /*uu_thresh*/
    0,

    /*sat_thresh*/
    0x00,

    /*ac_thresh*/
    0
};    



static ISP_CMOS_GAMMA_S g_stIspGamma =
{
    /* bvalid */
    0,

    {
        0, 180, 320, 426, 516, 590, 660, 730, 786, 844, 896, 946, 994, 1040, 1090, 1130, 1170, 1210, 1248,
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
        3994, 4002, 4006, 4010, 4018, 4022, 4032, 4038, 4046, 4050, 4056, 4062, 4072, 4076, 4084, 4090, 4095
    }
};




/////////////////////////////////////////////////////////////////////////////////

static HI_U32 cmos_get_isp_default(ISP_DEV IspDev,ISP_CMOS_DEFAULT_S *pstDef) {
	printf("cmos_get_isp_default\n");

    if (HI_NULL == pstDef) {
        printf("null pointer when get isp default value!\n");
        return -1;
    }

    memset(pstDef, 0, sizeof(ISP_CMOS_DEFAULT_S));



    
   
            pstDef->stDrc.bEnable               = HI_FALSE;

             pstDef->stDrc.u32BlackLevel         = 0x00;
            pstDef->stDrc.u32WhiteLevel         = 0xD0000; 
            pstDef->stDrc.u32SlopeMax           = 0x30;
            pstDef->stDrc.u32SlopeMin           = 0x00;
            pstDef->stDrc.u32VarianceSpace      = 0x04;
            pstDef->stDrc.u32VarianceIntensity  = 0x01;
            pstDef->stDrc.u32Asymmetry          = 0x08;
            pstDef->stDrc.u32BrightEnhance      = 0xE6;
            pstDef->stDrc.bFilterMux            = 0x1;
            pstDef->stDrc.u32Svariance          = 0x8;
            pstDef->stDrc.u32BrightPr           = 0xB0;
            pstDef->stDrc.u32Contrast           = 0xB0;
            pstDef->stDrc.u32DarkEnhance        = 0x8000;
            
            memcpy(&pstDef->stAgcTbl, &g_stIspAgcTable, sizeof(ISP_CMOS_AGC_TABLE_S));
            memcpy(&pstDef->stNoiseTbl, &g_stIspNoiseTable, sizeof(ISP_CMOS_NOISE_TABLE_S));
            memcpy(&pstDef->stDemosaic, &g_stIspDemosaic, sizeof(ISP_CMOS_DEMOSAIC_S));
            memcpy(&pstDef->stBayerSharpen, &g_stIspBayerSharpen, sizeof(ISP_CMOS_BAYER_SHARPEN_S));
            memcpy(&pstDef->stYuvSharpen, &g_stIspYuvSharpen, sizeof(ISP_CMOS_YUV_SHARPEN_S));
            memcpy(&pstDef->stGamma, &g_stIspGamma, sizeof(ISP_CMOS_GAMMA_S));


    pstDef->stSensorMaxResolution.u32MaxWidth  = 1920;//1920;
    pstDef->stSensorMaxResolution.u32MaxHeight = 1080;//1080;
    pstDef->stSensorMode.u32SensorID = IMX265_ID;
    pstDef->stSensorMode.u8SensorMode = g_apstSnsState[IspDev]->u8ImgMode;

    return 0;
}

static HI_U32 cmos_get_isp_black_level(ISP_DEV IspDev,ISP_CMOS_BLACK_LEVEL_S *pstBlackLevel) {
	printf("cmos_get_isp_black_level\n");

    HI_S32  i;

    if (HI_NULL == pstBlackLevel) {
        printf("null pointer when get isp black level value!\n");
        return -1;
    }

    /* Don't need to update black level when iso change */
    pstBlackLevel->bUpdate = HI_FALSE;

    /* black level of linear mode */
   // if (WDR_MODE_NONE == g_apstSnsState[IspDev]->enWDRMode) {
        for (i=0; i<4; i++) {
            pstBlackLevel->au16BlackLevel[i] = 0x80;//0x3c;//0xc8;//0xf0;    // 240
        }
    //}

    return 0;
}

static HI_VOID cmos_set_pixel_detect(ISP_DEV IspDev,HI_BOOL bEnable) {
	printf("cmos_set_pixel_detect\n");
    return;
}


static HI_VOID cmos_set_wdr_mode(ISP_DEV IspDev,HI_U8 u8Mode) {
	printf("cmos_set_wdr_mode\n");

    g_apstSnsState[IspDev]->bSyncInit = HI_FALSE;

    switch(u8Mode) {
        case WDR_MODE_NONE:
            g_apstSnsState[IspDev]->enWDRMode = WDR_MODE_NONE;
            g_apstSnsState[IspDev]->u32FLStd = IMX265_VMAX_1080P25_LINEAR;
            g_apstSnsState[IspDev]->u8ImgMode = IMX265_SENSOR_1080P_25FPS_LINEAR_MODE;
            //g_astimx265State[IspDev].u8Hcg = 0x2; //???
            printf("linear mode\n");
        break;

        default:
            printf("NOT support this mode!\n");
            return;
        break;
    }

    g_apstSnsState[IspDev]->au32FL[0]= g_apstSnsState[IspDev]->u32FLStd;
    g_apstSnsState[IspDev]->au32FL[1] = g_apstSnsState[IspDev]->au32FL[0];
    memset(g_apstSnsState[IspDev]->au32WDRIntTime, 0, sizeof(g_apstSnsState[IspDev]->au32WDRIntTime));

    return;
}

static HI_U32 cmos_get_sns_regs_info(ISP_DEV IspDev,ISP_SNS_REGS_INFO_S *pstSnsRegsInfo) {
  	//printf("cmos_get_sns_regs_info\n");
    if (HI_NULL == pstSnsRegsInfo) {
        printf("null pointer when get sns reg info!\n");
        return -1;
    }

int i;

     if ((HI_FALSE == g_apstSnsState[IspDev]->bSyncInit) )//|| (HI_FALSE == pstSnsRegsInfo->bConfig))
    {
        g_apstSnsState[IspDev]->astRegsInfo[0].enSnsType = ISP_SNS_I2C_TYPE;
        g_apstSnsState[IspDev]->astRegsInfo[0].unComBus.s8I2cDev = g_aunImx265BusInfo[IspDev].s8I2cDev;        
        g_apstSnsState[IspDev]->astRegsInfo[0].u8Cfg2ValidDelayMax = 2;
        g_apstSnsState[IspDev]->astRegsInfo[0].u32RegNum = 0;

        for (i=0; i<g_apstSnsState[IspDev]->astRegsInfo[0].u32RegNum; i++)
        {
            g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[i].bUpdate = HI_TRUE;
            g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[i].u8DevAddr = imx265_i2c_addr;
            g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[i].u32AddrByteNum = imx265_addr_byte;
            g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[i].u32DataByteNum = imx265_data_byte;
        }

        //SHR registers
        g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[0].u8DelayFrmNum = 0;
        g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[0].u32RegAddr = 0x308d;//IMX326_SHR_L;
        g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[1].u8DelayFrmNum = 0;
        g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[1].u32RegAddr = 0x308e;//IMX326_SHR_H;
        g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[2].u8DelayFrmNum = 0;
        g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[2].u32RegAddr = 0x308f;//IMX326_SHR_H;


        // gain related
        g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[3].u32RegAddr = 0x3204;//IMX326_PGC_L;
        g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[3].u8DelayFrmNum = 0;
        g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[4].u32RegAddr = 0x3205;//IMX326_PGC_H;
        g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[4].u8DelayFrmNum = 0;
       

        //VMAX registers
        //g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[5].u8DelayFrmNum = 0;
        //g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[5].u32RegAddr = 0x3010;//IMX326_VMAX_L;
        //g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[6].u8DelayFrmNum = 0;
        //g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[6].u32RegAddr = 0x3011;//IMX326_VMAX_M;
        //g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[7].u8DelayFrmNum = 0;
        //g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[7].u32RegAddr = 0x3012;//IMX326_VMAX_H;
   
        
        g_apstSnsState[IspDev]->bSyncInit = HI_TRUE;
    }
    else
    {
	//printf("checking new regs\n");
        for (i=0; i<g_apstSnsState[IspDev]->astRegsInfo[0].u32RegNum; i++)
        {
		//printf("i = %d old = %d new = %d\n", i,
		//	g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[i].u32Data,
		//	g_apstSnsState[IspDev]->astRegsInfo[1].astI2cData[i].u32Data );
            if (g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[i].u32Data == g_apstSnsState[IspDev]->astRegsInfo[1].astI2cData[i].u32Data)
            {
                g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[i].bUpdate = HI_FALSE;
            }
            else
            {
                g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[i].bUpdate = HI_FALSE;//HI_TRUE;
		//printf("SNS need update\n");
            }
        }
    }

    memcpy(pstSnsRegsInfo, &g_apstSnsState[IspDev]->astRegsInfo[0], sizeof(ISP_SNS_REGS_INFO_S)); 
    memcpy(&g_apstSnsState[IspDev]->astRegsInfo[1], &g_apstSnsState[IspDev]->astRegsInfo[0], sizeof(ISP_SNS_REGS_INFO_S)); 
    g_apstSnsState[IspDev]->au32FL[1] = g_apstSnsState[IspDev]->au32FL[0];


    return 0;
}

static HI_S32 cmos_set_image_mode(ISP_DEV IspDev,ISP_CMOS_SENSOR_IMAGE_MODE_S *pstSensorImageMode) {
	printf("cmos_set_image_mode\n");

    HI_U8 u8SensorImageMode = g_apstSnsState[IspDev]->u8ImgMode;

    g_apstSnsState[IspDev]->bSyncInit = HI_FALSE;

    if (HI_NULL == pstSensorImageMode ) {
        printf("null pointer when set image mode\n");
        return -1;
    }

    if ((pstSensorImageMode->u16Width == 1920) && (pstSensorImageMode->u16Height == 1080)) {
        //OK!
    } else if ((pstSensorImageMode->u16Width == 2048) && (pstSensorImageMode->u16Height == 1536)) {
            //OK
    } else {
        printf("Not support! Width:%d, Height:%d, Fps:%f, WDRMode:%d\n",
            pstSensorImageMode->u16Width,
            pstSensorImageMode->u16Height,
            pstSensorImageMode->f32Fps,
            g_apstSnsState[IspDev]->enWDRMode);

        return -1;
    }

    if ((HI_TRUE == g_apstSnsState[IspDev]->bInit) && (u8SensorImageMode == g_apstSnsState[IspDev]->u8ImgMode)) {
        /* Don't need to switch SensorImageMode */
        return -1;
    }

    return 0;
}

static HI_VOID sensor_global_init(ISP_DEV IspDev) {

	printf("sensor_global_init\n");

    g_apstSnsState[IspDev]->bInit = HI_FALSE;
    g_apstSnsState[IspDev]->bSyncInit = HI_FALSE;
    g_apstSnsState[IspDev]->u8ImgMode = IMX265_SENSOR_1080P_25FPS_LINEAR_MODE;
    g_apstSnsState[IspDev]->enWDRMode = WDR_MODE_NONE;
    g_apstSnsState[IspDev]->u32FLStd = IMX265_VMAX_1080P25_LINEAR;
    g_apstSnsState[IspDev]->au32FL[0] = IMX265_VMAX_1080P25_LINEAR;
    g_apstSnsState[IspDev]->au32FL[1] = IMX265_VMAX_1080P25_LINEAR;

    memset(&g_apstSnsState[IspDev]->astRegsInfo[0], 0, sizeof(ISP_SNS_REGS_INFO_S));
    memset(&g_apstSnsState[IspDev]->astRegsInfo[1], 0, sizeof(ISP_SNS_REGS_INFO_S));
}

static HI_S32 cmos_init_sensor_exp_function(ISP_SENSOR_EXP_FUNC_S *pstSensorExpFunc) {

	printf("cmos_init_sensor_exp_function\n");

    memset(pstSensorExpFunc, 0, sizeof(ISP_SENSOR_EXP_FUNC_S));

    pstSensorExpFunc->pfn_cmos_sensor_init = imx265_init; //imx265_sensor_ctl OK
    pstSensorExpFunc->pfn_cmos_sensor_exit = imx265_exit; //imx265_sensor_ctl DEFAULT
    pstSensorExpFunc->pfn_cmos_sensor_global_init = sensor_global_init;//imx265_cmos DEFAULT (constants customized)
    pstSensorExpFunc->pfn_cmos_set_image_mode = cmos_set_image_mode;//imx265_cmos DEFAULT
    pstSensorExpFunc->pfn_cmos_set_wdr_mode = cmos_set_wdr_mode;//imx265_cmos g_astimx265State[IspDev].u8Hcg ???

    pstSensorExpFunc->pfn_cmos_get_isp_default = cmos_get_isp_default;//imx265_cmos ???
    pstSensorExpFunc->pfn_cmos_get_isp_black_level = cmos_get_isp_black_level;//imx265_cmos ???
    pstSensorExpFunc->pfn_cmos_set_pixel_detect = cmos_set_pixel_detect;//imx265_cmos ???
    pstSensorExpFunc->pfn_cmos_get_sns_reg_info = cmos_get_sns_regs_info;//imx265_cmos ???

    return 0;
}


///////////////////////////////////////////////////////////////////////////////
#define IMX290_FULL_LINES_MAX  (0x3FFFF)

static HI_S32 cmos_get_ae_default(ISP_DEV IspDev,AE_SENSOR_DEFAULT_S *pstAeSnsDft)
{
	printf("cmos_get_ae_default\n");
    if (HI_NULL == pstAeSnsDft)
    {
        printf("null pointer when get ae default value!\n");
        return - 1;
    }
	//return 0;
    memset(&pstAeSnsDft->stAERouteAttr, 0, sizeof(ISP_AE_ROUTE_S));
      
    pstAeSnsDft->u32FullLinesStd = IMX265_VMAX_1080P25_LINEAR;//g_apstSnsState[IspDev]->u32FLStd;
    pstAeSnsDft->u32FlickerFreq = 0;//50*256;
    //pstAeSnsDft->u32FullLinesMax = IMX290_FULL_LINES_MAX;

    pstAeSnsDft->stIntTimeAccu.enAccuType = AE_ACCURACY_LINEAR;
    pstAeSnsDft->stIntTimeAccu.f32Accuracy = 1;
    pstAeSnsDft->stIntTimeAccu.f32Offset = 0.320;

    pstAeSnsDft->stAgainAccu.enAccuType = AE_ACCURACY_TABLE;
    pstAeSnsDft->stAgainAccu.f32Accuracy = 1;

    pstAeSnsDft->stDgainAccu.enAccuType = AE_ACCURACY_TABLE;
    pstAeSnsDft->stDgainAccu.f32Accuracy = 1;
    
    pstAeSnsDft->u32ISPDgainShift = 3;
    pstAeSnsDft->u32MinISPDgainTarget = 1 << pstAeSnsDft->u32ISPDgainShift;
    pstAeSnsDft->u32MaxISPDgainTarget = 8 << pstAeSnsDft->u32ISPDgainShift;

  
	pstAeSnsDft->u32LinesPer500ms = g_apstSnsState[IspDev]->u32FLStd*30/2;//(u32Fll * U32MaxFps) >> 1
	


    pstAeSnsDft->enMaxIrisFNO = ISP_IRIS_F_NO_1_0;
    pstAeSnsDft->enMinIrisFNO = ISP_IRIS_F_NO_32_0;

    pstAeSnsDft->bAERouteExValid = HI_FALSE;
    pstAeSnsDft->stAERouteAttr.u32TotalNum = 0;
    pstAeSnsDft->stAERouteAttrEx.u32TotalNum = 0;

    switch(g_apstSnsState[IspDev]->enWDRMode)
    {
        default:
        case WDR_MODE_NONE:   /*linear mode*/
            pstAeSnsDft->au8HistThresh[0] = 0xd;
            pstAeSnsDft->au8HistThresh[1] = 0x28;
            pstAeSnsDft->au8HistThresh[2] = 0x60;
            pstAeSnsDft->au8HistThresh[3] = 0x80;

		
            pstAeSnsDft->u32MaxAgain = 62564; 
            pstAeSnsDft->u32MinAgain = 2024;
            pstAeSnsDft->u32MaxAgainTarget = pstAeSnsDft->u32MaxAgain;
            pstAeSnsDft->u32MinAgainTarget = pstAeSnsDft->u32MinAgain;
		
		/*
            pstAeSnsDft->u32MaxAgain = 15888;//15888;//22795;
            pstAeSnsDft->u32MinAgain = 0x400;
            pstAeSnsDft->u32MaxAgainTarget = pstAeSnsDft->u32MaxAgain;
            pstAeSnsDft->u32MinAgainTarget = pstAeSnsDft->u32MinAgain;
*/
		/*
            pstAeSnsDft->u32MaxDgain = 38577;  
            pstAeSnsDft->u32MinDgain = 1024;
            pstAeSnsDft->u32MaxDgainTarget = 20013;
            pstAeSnsDft->u32MinDgainTarget = pstAeSnsDft->u32MinDgain;
            */
            pstAeSnsDft->u32MaxDgain = 16000;//15888;  
            pstAeSnsDft->u32MinDgain = 0x800;//1024;
            pstAeSnsDft->u32MaxDgainTarget = pstAeSnsDft->u32MaxDgain;
            pstAeSnsDft->u32MinDgainTarget = pstAeSnsDft->u32MinDgain;

            pstAeSnsDft->u8AeCompensation = 30;//0x38;//20
            pstAeSnsDft->enAeExpMode = AE_EXP_HIGHLIGHT_PRIOR; 

            pstAeSnsDft->u32InitExposure = 6;//6;//160000;//g_au32InitExposure[IspDev] ? g_au32InitExposure[IspDev] : 148859;
            
            pstAeSnsDft->u32MaxIntTime = 1119;//g_apstSnsState[IspDev]->u32FLStd;//2;
            pstAeSnsDft->u32MinIntTime = 6;//4;//1;

            pstAeSnsDft->u32MaxIntTimeTarget = 1119;//g_apstSnsState[IspDev]->u32FLStd;//65535;
            pstAeSnsDft->u32MinIntTimeTarget = 6;//1;

////


///
        break;

    }

    return 0;
}

static HI_VOID cmos_fps_set(ISP_DEV IspDev, HI_FLOAT f32Fps, AE_SENSOR_DEFAULT_S *pstAeSnsDft)
{
	printf("cmos_fps_set %f\n", f32Fps);
    HI_U32 u32VMAX = IMX265_VMAX_1080P25_LINEAR;                                                                            
          
	/*                                                                                                                
    switch (g_apstSnsState[IspDev]->u8ImgMode)                                                                                           
    {
      
      case IMX265_SENSOR_1080P_25FPS_LINEAR_MODE:
           if ((f32Fps == 30))                                                                            
           {
               u32VMAX = IMX265_VMAX_1080P25_LINEAR;// * 30 / f32Fps;  
           }
           else                                                                                                              
           {                                                                                                                 
               printf("Not support Fps: %f\n", f32Fps);                                                                      
               return;                                                                                                       
           } 
           u32VMAX = (u32VMAX > IMX265_VMAX_1080P25_LINEAR) ? IMX265_VMAX_1080P25_LINEAR : u32VMAX;   
           break;
                                                                                                       
      default:                                                                                                              
          return;
          break;                                                                                                            
    }                                                                                                                                                                                          
       */
	                                                                                                                 
    //if (WDR_MODE_NONE == g_apstSnsState[IspDev]->enWDRMode)                                                                                   
   // {                                                                                                                     
        //g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[5].u32Data = (u32VMAX & 0xFF);                                                         
        //g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[6].u32Data = ((u32VMAX & 0xFF00) >> 8);                                                
        //g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[7].u32Data = ((u32VMAX & 0xF0000) >> 16);                                              
   // }                                                                                                                  
                                                                                                          
    g_apstSnsState[IspDev]->u32FLStd = u32VMAX;                                                                                       
                                                                                                                   
                                                                                                                          
    pstAeSnsDft->f32Fps = f32Fps;                                    
    pstAeSnsDft->u32LinesPer500ms = g_apstSnsState[IspDev]->u32FLStd * f32Fps / 2;                                                        
    pstAeSnsDft->u32FullLinesStd = g_apstSnsState[IspDev]->u32FLStd;  
    pstAeSnsDft->u32MaxIntTime = g_apstSnsState[IspDev]->u32FLStd - 2;  
 
    g_apstSnsState[IspDev]->au32FL[0] = g_apstSnsState[IspDev]->u32FLStd;

    pstAeSnsDft->u32FullLines = g_apstSnsState[IspDev]->au32FL[0];
                                                                                                                          
    return;                                                                                                               

}

static HI_VOID cmos_slow_framerate_set(ISP_DEV IspDev,HI_U32 u32FullLines,
    AE_SENSOR_DEFAULT_S *pstAeSnsDft)
{
   printf("cmos_slow_framerate_set\n");
   
    	u32FullLines = (u32FullLines > IMX265_VMAX_1080P25_LINEAR) ? IMX265_VMAX_1080P25_LINEAR : u32FullLines;
        g_apstSnsState[IspDev]->au32FL[0] = u32FullLines;  
  

    if(WDR_MODE_NONE == g_apstSnsState[IspDev]->enWDRMode)
    {
        //g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[5].u32Data = (g_apstSnsState[IspDev]->au32FL[0] & 0xFF);
        //g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[6].u32Data = ((g_apstSnsState[IspDev]->au32FL[0] & 0xFF00) >> 8);
        //g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[7].u32Data = ((g_apstSnsState[IspDev]->au32FL[0] & 0xF0000) >> 16);
    } 

    pstAeSnsDft->u32FullLines = g_apstSnsState[IspDev]->au32FL[0];
    pstAeSnsDft->u32MaxIntTime = g_apstSnsState[IspDev]->au32FL[0] - 2; 

    return;
}

static HI_VOID cmos_inttime_update(ISP_DEV IspDev,HI_U32 u32IntTime)
{
	//printf("cmos_inttime_update ");
HI_U32 u32Value = 0;
	/*
    static HI_BOOL bFirst[ISP_MAX_DEV_NUM] ={1, 1}; 
    HI_U32 u32Value = 0;

    static HI_U8 u8Count[ISP_MAX_DEV_NUM] = {0};

    static HI_U32 u32ShortIntTime[ISP_MAX_DEV_NUM] = {0};
    static HI_U32 u32ShortIntTime1[ISP_MAX_DEV_NUM] = {0}; 
    static HI_U32 u32ShortIntTime2[ISP_MAX_DEV_NUM] = {0};
    static HI_U32 u32LongIntTime[ISP_MAX_DEV_NUM] = {0};         

    static HI_U32 u32RHS1[ISP_MAX_DEV_NUM]  = {0}; 
    static HI_U32 u32RHS2[ISP_MAX_DEV_NUM]  = {0};  

    static HI_U32 u32SHS1[ISP_MAX_DEV_NUM]  = {0};                      
    static HI_U32 u32SHS2[ISP_MAX_DEV_NUM]  = {0};
    static HI_U32 u32SHS3[ISP_MAX_DEV_NUM]  = {0};

    HI_U32 u32YOU *pu32DgainDb = 0
new again 0 old 0
*pu32AgainLin = 1024, *pu32AgainDb = 0
*pu32DgainLin = 1024, *pu32DgainDb = 0
new again 0 old 0
*pu32AgainLin = 1024, *pu32AgainDb = 0
*pu32DgainLin = 1024, *pu32DgainDb = 0
new again 0 old 0
*pu32AgainLin = 1024, *pu32AgainDb = 0
*pu32DgainLin = 1024, *pu32DgainDb = 0
new again 0 old 0
*pu32AgainLin = 1024, *pu32AgainDb = 0
*pu32DgainLin = 1024, *pu32DgainDb = 0
new again 0 old 0
*pu32AgainLin = 1024, *pu32AgainDb = 0
*pu32DgainLin = 1024, *pu32DgainDb = 0
new again 0 old 0
TSIZE;
                                                                                               
         u32Value = g_apstSnsState[IspDev]->au32FL[0] - u32IntTime - 1; 
                                                                                                       
         g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[0].u32Data = (u32Value & 0xFF);                                    
         g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[1].u32Data = ((u32Value & 0xFF00) >> 8);                           
         g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[2].u32Data = ((u32Value & 0x30000) >> 16); 
         
         bFirst[IspDev] = HI_TRUE;                                                                             
        */    


	//u32IntTime = 500;

	HI_U32 old = g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[0].u32Data +
		(g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[1].u32Data << 8) +
		(g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[2].u32Data << 16);

	u32Value = g_apstSnsState[IspDev]->au32FL[0] - u32IntTime - 1; 
	//u32Value = u32IntTime;

	//printf("(%d) new %d old %d\n", u32IntTime, u32Value, old);

	if (u32Value < 6 ) u32Value = 6;
	if (u32Value > 1119 ) u32Value = 1119;

//printf("u32IntTime = %d, u32Value = %d\n", u32IntTime, u32Value);
	                   
     //g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[0].u32Data = ((u32Value) & 0xFF);    
     // g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[1].u32Data = (((u32Value) & 0xFF00) >> 8); 
     // g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[2].u32Data = (((u32Value) & 0x0F0000) >>  16);     
      
    //imx265_write_register(IspDev, 0x308D, u32Value & 0xFF); //24'h012202 //???	
    //imx265_write_register(IspDev, 0x308E, (u32Value & 0xFF00) >> 8); //24'h012202 //???
	//imx265_write_register(IspDev, 0x308F, (u32Value & 0x0F0000) >> 16); //24'h012202 //???                                                                               
                                                                                                              
  return;                                                                                           

}

static HI_VOID cmos_gains_update(ISP_DEV IspDev,HI_U32 u32Again, HI_U32 u32Dgain)
{  
	//printf("cmos_gains_update ");
 HI_U32 u32Tmp;
u32Tmp=0;//u32Again+u32Dgain;
/*
    HI_U32 u32HCG = g_astimx265State[IspDev].u8Hcg;
    HI_U32 u32Tmp;
    
    if(u32Again >= 21)
    {
        u32HCG = u32HCG | 0x10;  // bit[4] HCG  .Reg0x3009[7:0]
        u32Again = u32Again - 21;
    }

    u32Tmp=u32Again+u32Dgain;
        
    g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[3].u32Data = (u32Tmp & 0xFF);
    g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[4].u32Data = (u32HCG & 0xFF);
    */

	/*
    g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[2].u32Data = (u32Again & 0xFF);
    g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[3].u32Data = ((u32Again >> 8) & 0x00FF);
    g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[4].u32Data = (u32Dgain & 0x7);
	*/
	//HI_U32 oldAgain = (g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[3].u32Data + (g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[4].u32Data << 8));
	//printf("new again %d old %d\n", u32Tmp, oldAgain);
       // g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[3].u32Data = (u32Tmp & 0xFF);
        //g_apstSnsState[IspDev]->astRegsInfo[0].astI2cData[4].u32Data = ((u32Tmp >> 8) & 0x0001);

    return;
}

static HI_U32 again_table[240]=
{
0, 1036, 1048, 1061, 1074, 1087, 1100, 1113, 1126, 1140, 1154, 1168, 1182, 1196, //1024
1210, 1225, 1240, 1255, 1270, 1285, 1300, 1316, 1332, 1348, 1364, 1380, 1397, 1414, 
1431, 1448, 1465, 1483, 1501, 1519, 1537, 1555, 1574, 1593, 1612, 1631, 1651, 1671, 
1691, 1711, 1732, 1753, 1774, 1795, 1817, 1839, 1861, 1883, 1906, 1929, 1952, 1975, 
1999, 2023, 2047, 2072, 2097, 2122, 2147, 2173, 2199, 2225, 2252, 2279, 2306, 2334, 
2362, 2390, 2419, 2448, 2477, 2507, 2537, 2567, 2598, 2629, 2661, 2693, 2725, 2758, 
2791, 2824, 2858, 2892, 2927, 2962, 2998, 3034, 3070, 3107, 3144, 3182, 3220, 3259, 
3298, 3338, 3378, 3419, 3460, 3502, 3544, 3587, 3630, 3674, 3718, 3763, 3808, 3854, 
3900, 3947, 3994, 4042, 4091, 4140, 4190, 4240, 4291, 4342, 4394, 4447, 4500, 4554, 
4609, 4664, 4720, 4777, 4834, 4892, 4951, 5010, 5070, 5131, 5193, 5255, 5318, 5382, 
5447, 5512, 5578, 5645, 5713, 5782, 5851, 5921, 5992, 6064, 6137, 6211, 6286, 6361, 
6437, 6514, 6592, 6671, 6751, 6832, 6914, 6997, 7081, 7166, 7252, 7339, 7427, 7516, 
7606, 7697, 7789, 7882, 7977, 8073, 8170, 8268, 8367, 8467, 8569, 8672, 8776, 8881, 
8988, 9096, 9205, 9315, 9427, 9540, 9654, 9770, 9887, 10006, 10126, 10248, 10371, 
10495, 10621, 10748, 10877, 11008, 11140, 11274, 11409, 11546, 11685, 11825, 11967, 
12111, 12256, 12403, 12552, 12703, 12855, 13009, 13165, 13323, 13483, 13645, 13809, 
13975, 14143, 14313, 14485, 14659, 14835, 15013, 15193, 15375, 15560, 15747, 15936, 
16127, 16321, 16517, 16715, 16916, 17119, 17324, 17532, 17742
};

static HI_VOID cmos_again_calc_table(ISP_DEV IspDev,HI_U32 *pu32AgainLin, HI_U32 *pu32AgainDb)
{
	//printf("cmos_again_calc_table\n");
    int i;

    if((HI_NULL == pu32AgainLin) ||(HI_NULL == pu32AgainDb))
    {
        printf("null pointer when get ae sensor gain info  value!\n");
        return;
    }

    if (*pu32AgainLin >= again_table[239])
    {
         *pu32AgainLin = again_table[239];
         *pu32AgainDb = 239;
         goto calc_table_end;
    }

    for (i = 1; i < 240; i++)
    {
        if (*pu32AgainLin < again_table[i])
        {
            *pu32AgainLin = again_table[i - 1];
            *pu32AgainDb = i - 1;
            goto calc_table_end;
        }
    }

calc_table_end:

    //*pu32AgainDb <<= 2;

    return;
}

static HI_VOID cmos_dgain_calc_table(ISP_DEV IspDev,HI_U32 *pu32DgainLin, HI_U32 *pu32DgainDb)
{
	//printf("cmos_dgain_calc_table\n");
    int i;

    if((HI_NULL == pu32DgainLin) ||(HI_NULL == pu32DgainDb))
    {
        printf("null pointer when get ae sensor gain info  value!\n");
        return;
    }

    if (*pu32DgainLin >= again_table[239])
    {
         *pu32DgainLin = again_table[239];
         *pu32DgainDb = 239;
         goto calc_table_end;
    }

    for (i = 1; i < 240; i++)
    {
        if (*pu32DgainLin < again_table[i])
        {
            *pu32DgainLin = again_table[i - 1];
            *pu32DgainDb = i - 1;
            goto calc_table_end;
        }
    }

calc_table_end:

    //*pu32DgainDb <<= 2;

    return;
    return;
}

static HI_VOID cmos_get_inttime_max(ISP_DEV IspDev,HI_U16 u16ManRatioEnable, HI_U32 *au32Ratio, HI_U32 *au32IntTimeMax, HI_U32 *au32IntTimeMin, HI_U32 *pu32LFMaxIntTime)
{
	//printf("cmos_get_inttime_max\n");
    HI_U32 i = 0;
    HI_U32 u32IntTimeMaxTmp0 = 0;
    HI_U32 u32IntTimeMaxTmp  = 0;
    HI_U32 u32RHS2_Max=0;
    HI_U32 u32RatioTmp = 0x40;
    HI_U32 u32ShortTimeMinLimit = 0;

    u32ShortTimeMinLimit = (WDR_MODE_2To1_LINE == g_apstSnsState[IspDev]->enWDRMode) ? 2 : ((WDR_MODE_3To1_LINE == g_apstSnsState[IspDev]->enWDRMode) ? 3 : 2);
        
    if((HI_NULL == au32Ratio) || (HI_NULL == au32IntTimeMax) || (HI_NULL == au32IntTimeMin))
    {
        printf("null pointer when get ae sensor ExpRatio/IntTimeMax/IntTimeMin value!\n");
        return;
    }                                                                                                                       
   
   
 

    if(u32IntTimeMaxTmp >= u32ShortTimeMinLimit)
    {
        if (IS_LINE_WDR_MODE(g_apstSnsState[IspDev]->enWDRMode))
        {
            au32IntTimeMax[0] = u32IntTimeMaxTmp;
            au32IntTimeMax[1] = au32IntTimeMax[0] * au32Ratio[0] >> 6;
            au32IntTimeMax[2] = au32IntTimeMax[1] * au32Ratio[1] >> 6;
            au32IntTimeMax[3] = au32IntTimeMax[2] * au32Ratio[2] >> 6;
            au32IntTimeMin[0] = u32ShortTimeMinLimit;
            au32IntTimeMin[1] = au32IntTimeMin[0] * au32Ratio[0] >> 6;
            au32IntTimeMin[2] = au32IntTimeMin[1] * au32Ratio[1] >> 6;
            au32IntTimeMin[3] = au32IntTimeMin[2] * au32Ratio[2] >> 6;
        }
        else
        {
        }
    }
    else
    {
        if(1 == u16ManRatioEnable)
        {
            printf("Manaul ExpRatio is too large!\n");
            return;
        }
        else
        {
            u32IntTimeMaxTmp = u32ShortTimeMinLimit; 

            au32IntTimeMin[0] = au32IntTimeMax[0];
            au32IntTimeMin[1] = au32IntTimeMax[1];
            au32IntTimeMin[2] = au32IntTimeMax[2];
            au32IntTimeMin[3] = au32IntTimeMax[3];
        }
    }

    return;
                                                                                                                
}

static HI_VOID cmos_ae_fswdr_attr_set(ISP_DEV IspDev,AE_FSWDR_ATTR_S *pstAeFSWDRAttr)
{
    //genFSWDRMode[IspDev] = pstAeFSWDRAttr->enFSWDRMode;
   // gu32MaxTimeGetCnt[IspDev] = 0;
}


static HI_S32 cmos_init_ae_exp_function(AE_SENSOR_EXP_FUNC_S *pstExpFuncs)
{
    memset(pstExpFuncs, 0, sizeof(AE_SENSOR_EXP_FUNC_S));

    pstExpFuncs->pfn_cmos_get_ae_default    = cmos_get_ae_default;
    pstExpFuncs->pfn_cmos_fps_set           = cmos_fps_set;
    pstExpFuncs->pfn_cmos_slow_framerate_set= cmos_slow_framerate_set;    
    pstExpFuncs->pfn_cmos_inttime_update    = cmos_inttime_update;
    pstExpFuncs->pfn_cmos_gains_update      = cmos_gains_update;
    pstExpFuncs->pfn_cmos_again_calc_table  = cmos_again_calc_table;
    pstExpFuncs->pfn_cmos_dgain_calc_table  = cmos_dgain_calc_table;
    pstExpFuncs->pfn_cmos_get_inttime_max   = cmos_get_inttime_max; 
    pstExpFuncs->pfn_cmos_ae_fswdr_attr_set = cmos_ae_fswdr_attr_set; 

    return 0;
}
///////////////////////////////////////////////////////////////////////////////
static AWB_CCM_S g_stAwbCcm =
{
    5100,
    {      
        0x01AB, 0x80A4, 0x8007,
        0x804B, 0x0182, 0x8037,
        0x8006, 0x80A4, 0x01AA   
    },

    3600,
    {
        0x018E, 0x807A, 0x8014,
        0x8066, 0x019C, 0x8036,
        0x0002, 0x80A2, 0x01A0,       
    }, 

    2500,
    {     
        0x01B0, 0x80A1, 0x800F,
        0x807D, 0x018B, 0x800E,
        0x0010, 0x81C0, 0x02B0        
    }
};


static AWB_AGC_TABLE_S g_stAwbAgcTable = 
{
    0,
    {0x80,0x80,0x78,0x74,0x68,0x60,0x58,0x50,0x48,0x40,0x38,0x38,0x38,0x38,0x38,0x38}
};




static HI_S32 cmos_get_awb_default(ISP_DEV IspDev, AWB_SENSOR_DEFAULT_S *pstAwbSnsDft)
{

	printf("cmos_get_awb_default\n");

    if (HI_NULL == pstAwbSnsDft)
    {
        printf("null pointer when get awb default value!\n");
        return -1;
    }

    memset(pstAwbSnsDft, 0, sizeof(AWB_SENSOR_DEFAULT_S));
    pstAwbSnsDft->u16WbRefTemp = 5120;

/*
#define CALIBRATE_STATIC_WB_R_GAIN (550)
#define CALIBRATE_STATIC_WB_GR_GAIN (256)
#define CALIBRATE_STATIC_WB_GB_GAIN (256)
#define CALIBRATE_STATIC_WB_B_GAIN (440)


#define CALIBRATE_AWB_P1 (256)
#define CALIBRATE_AWB_P2 (-257)
#define CALIBRATE_AWB_Q1 (-253)
#define CALIBRATE_AWB_A1 (0)
#define CALIBRATE_AWB_B1 (128)
#define CALIBRATE_AWB_C1 (256000)


#define GOLDEN_RGAIN 273
#define GOLDEN_BGAIN 592

*/
/* Calibration results for Static WB */
#define CALIBRATE_STATIC_WB_R_GAIN 0x226 
#define CALIBRATE_STATIC_WB_GR_GAIN 0x100 
#define CALIBRATE_STATIC_WB_GB_GAIN 0x100 
#define CALIBRATE_STATIC_WB_B_GAIN 0x1B8 
/* Calibration results for Auto WB Planck */
#define CALIBRATE_AWB_P1 -27  
#define CALIBRATE_AWB_P2 283 
#define CALIBRATE_AWB_Q1 0 
#define CALIBRATE_AWB_A1 165329  
#define CALIBRATE_AWB_B1 128  
#define CALIBRATE_AWB_C1 -118136 
/* Rgain and Bgain of the golden sample */
#define GOLDEN_RGAIN 0   
#define GOLDEN_BGAIN 0   

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
	/*
    pstAwbSnsDft->au16GainOffset[0] = 0x1C3;
    pstAwbSnsDft->au16GainOffset[1] = 0x100;
    pstAwbSnsDft->au16GainOffset[2] = 0x100;
    pstAwbSnsDft->au16GainOffset[3] = 0x1D4;

    pstAwbSnsDft->as32WbPara[0] = -37;
    pstAwbSnsDft->as32WbPara[1] = 293;
    pstAwbSnsDft->as32WbPara[2] = 0;
    pstAwbSnsDft->as32WbPara[3] = 179537;
    pstAwbSnsDft->as32WbPara[4] = 128;
    pstAwbSnsDft->as32WbPara[5] = -123691;
	*/
   
    pstAwbSnsDft->u16GoldenRgain = GOLDEN_RGAIN;
    pstAwbSnsDft->u16GoldenBgain = GOLDEN_BGAIN;
    
    switch (g_apstSnsState[IspDev]->enWDRMode)
    {
        default:
        case WDR_MODE_NONE:
            memcpy(&pstAwbSnsDft->stAgcTbl, &g_stAwbAgcTable, sizeof(AWB_AGC_TABLE_S));
            memcpy(&pstAwbSnsDft->stCcm, &g_stAwbCcm, sizeof(AWB_CCM_S));
        break;

    }

    pstAwbSnsDft->u16SampleRgain = g_au16SampleRgain[IspDev];
    pstAwbSnsDft->u16SampleBgain = g_au16SampleBgain[IspDev];
    pstAwbSnsDft->u16InitRgain = g_au16InitWBGain[IspDev][0];
    pstAwbSnsDft->u16InitGgain = g_au16InitWBGain[IspDev][1];
    pstAwbSnsDft->u16InitBgain = g_au16InitWBGain[IspDev][2];

    pstAwbSnsDft->u8AWBRunInterval = 2;

    return 0;
}

static HI_S32 cmos_init_awb_exp_function(AWB_SENSOR_EXP_FUNC_S *pstExpFuncs)
{
    memset(pstExpFuncs, 0, sizeof(AWB_SENSOR_EXP_FUNC_S));

    pstExpFuncs->pfn_cmos_get_awb_default = cmos_get_awb_default;

    return 0;
}


/****************************************************************************
 * callback structure                                                       *
 ****************************************************************************/

static int imx265_set_bus_info(ISP_DEV IspDev, ISP_SNS_COMMBUS_U unSNSBusInfo) {
printf("imx265_set_bus_info\n");
    g_aunImx265BusInfo[IspDev].s8I2cDev = unSNSBusInfo.s8I2cDev;

    return 0;
}




static int sensor_register_callback(ISP_DEV IspDev, ALG_LIB_S *pstAeLib, ALG_LIB_S *pstAwbLib) {
printf("sensor_register_callback\n");
    HI_S32 s32Ret;
    ISP_SENSOR_REGISTER_S stIspRegister;
    AE_SENSOR_REGISTER_S  stAeRegister;
    AWB_SENSOR_REGISTER_S stAwbRegister;

    cmos_init_sensor_exp_function(&stIspRegister.stSnsExp);
    s32Ret = HI_MPI_ISP_SensorRegCallBack(IspDev, IMX265_ID, &stIspRegister);
    if (s32Ret) {
        printf("sensor register callback function failed!\n");
        return s32Ret;
    }

	
    cmos_init_ae_exp_function(&stAeRegister.stSnsExp);
    s32Ret = HI_MPI_AE_SensorRegCallBack(IspDev, pstAeLib, IMX265_ID, &stAeRegister);
    if (s32Ret) {
        printf("sensor register callback function to ae lib failed!\n");
        return s32Ret;
    }

    cmos_init_awb_exp_function(&stAwbRegister.stSnsExp);
    s32Ret = HI_MPI_AWB_SensorRegCallBack(IspDev, pstAwbLib, IMX265_ID, &stAwbRegister);
    if (s32Ret) {
        printf("sensor register callback function to awb lib failed!\n");
        return s32Ret;
    }

    return 0;
}

static int sensor_unregister_callback(ISP_DEV IspDev, ALG_LIB_S *pstAeLib, ALG_LIB_S *pstAwbLib) {
	printf("sensor_unregister_callback\n");
    HI_S32 s32Ret;

    s32Ret = HI_MPI_ISP_SensorUnRegCallBack(IspDev, IMX265_ID);
    if (s32Ret) {
        printf("sensor unregister callback function failed!\n");
        return s32Ret;
    }

    s32Ret = HI_MPI_AE_SensorUnRegCallBack(IspDev, pstAeLib, IMX265_ID);
    if (s32Ret) {
        printf("sensor unregister callback function to ae lib failed!\n");
        return s32Ret;
    }

    s32Ret = HI_MPI_AWB_SensorUnRegCallBack(IspDev, pstAwbLib, IMX265_ID);
    if (s32Ret) {
        printf("sensor unregister callback function to awb lib failed!\n");
        return s32Ret;
    }

    return 0;
}

static int sensor_set_init(ISP_DEV IspDev, ISP_INIT_ATTR_S *pstInitAttr) {
	printf("sensor_set_init\n");
    g_au32InitExposure[IspDev] = pstInitAttr->u32Exposure;
    g_au32LinesPer500ms[IspDev] = pstInitAttr->u32LinesPer500ms;
    g_au16InitWBGain[IspDev][0] = pstInitAttr->u16WBRgain;
    g_au16InitWBGain[IspDev][1] = pstInitAttr->u16WBGgain;
    g_au16InitWBGain[IspDev][2] = pstInitAttr->u16WBBgain;
    g_au16SampleRgain[IspDev] = pstInitAttr->u16SampleRgain;
    g_au16SampleBgain[IspDev] = pstInitAttr->u16SampleBgain;

    return 0;
}

ISP_SNS_OBJ_S stSnsImx265Obj = {
    .pfnRegisterCallback    = sensor_register_callback, //imx265_cmos DEFAULT
    .pfnUnRegisterCallback  = sensor_unregister_callback, //imx265_cmos DEFAULT
    .pfnStandby             = imx265_standby, //imx265_sensor_Ctl EMPTY
    .pfnRestart             = imx265_restart, //imx265_sensor_Ctl EMPTY
    .pfnWriteReg            = imx265_write_register, //imx265_sensor_ctl DEFAULT
    .pfnReadReg             = imx265_read_register, //imx265_sensor_ctl EMPTY
    .pfnSetBusInfo          = imx265_set_bus_info, //imx265_cmos DEFAULT
    .pfnSetInit             = sensor_set_init //imx265_cmos DEFAULT
};


#ifdef __cplusplus
#if __cplusplus
}
#endif
#endif /* End of #ifdef __cplusplus */

#endif /* __IMX265_CMOS_H_ */



