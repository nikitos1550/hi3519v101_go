#if !defined(__SOIH65_CMOS_H_)
#define __SOIH65_CMOS_H_

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


#define SOIH65_ID 65

#ifdef INIFILE_CONFIG_MODE

extern AE_SENSOR_DEFAULT_S  g_AeDft[];
extern AWB_SENSOR_DEFAULT_S g_AwbDft[];
extern ISP_CMOS_DEFAULT_S   g_IspDft[];
extern HI_S32 Cmos_LoadINIPara(const HI_CHAR *pcName);
#else

#endif


/****************************************************************************
 * local variables                                                          *
 ****************************************************************************/
//extern	int sensor_write_register(int addr, int data);

extern void h65_init();
extern void h65_exit();
extern int h65_write_register(int addr, int data);


extern const unsigned int h65_i2c_addr;
extern unsigned int h65_addr_byte;
extern unsigned int h65_data_byte;

#define SENSOR_960P_30FPS_MODE (1)


#define VMAX_SOIH65_960P30_LINEAR     (1000)   

#define FULL_LINES_MAX  (0xFFFF)


HI_U8 gu8SensorImageMode_h65 = SENSOR_960P_30FPS_MODE;
WDR_MODE_E genSensorMode_h65 = WDR_MODE_NONE;

static HI_U32 gu32FullLinesStd = VMAX_SOIH65_960P30_LINEAR;
static HI_U32 gu32FullLines = VMAX_SOIH65_960P30_LINEAR;

static HI_BOOL bInit = HI_FALSE;
HI_BOOL bSensorInit_h65 = HI_FALSE;
static ISP_SNS_REGS_INFO_S g_stSnsRegsInfo = {0};
static ISP_SNS_REGS_INFO_S g_stPreSnsRegsInfo = {0};

/* Piris attr */
static ISP_PIRIS_ATTR_S gstPirisAttr=
{
    0,      // bStepFNOTableChange
    1,      // bZeroIsMax
    93,     // u16TotalStep
    62,     // u16StepCount
    /* Step-F number mapping table. Must be from small to large. F1.0 is 1024 and F32.0 is 1 */
    {30,35,40,45,50,56,61,67,73,79,85,92,98,105,112,120,127,135,143,150,158,166,174,183,191,200,208,217,225,234,243,252,261,270,279,289,298,307,316,325,335,344,353,362,372,381,390,399,408,417,426,435,444,453,462,470,478,486,493,500,506,512},
    ISP_IRIS_F_NO_1_4, // enMaxIrisFNOTarget
    ISP_IRIS_F_NO_5_6  // enMinIrisFNOTarget
};


#define PATHLEN_MAX 256
#define CMOS_CFG_INI "soih65_cfg.ini"
static char pcName[PATHLEN_MAX] = "configs/soih65_cfg.ini";


/* AE default parameter and function */
static HI_S32 cmos_get_ae_default(AE_SENSOR_DEFAULT_S *pstAeSnsDft)
{
    if (HI_NULL == pstAeSnsDft)
    {
        printf("null pointer when get ae default value!\n");
        return -1;
    }

    pstAeSnsDft->u32LinesPer500ms = VMAX_SOIH65_960P30_LINEAR * 30 / 2;
    pstAeSnsDft->u32FullLinesStd = gu32FullLinesStd;
    pstAeSnsDft->u32FlickerFreq = 0;

    pstAeSnsDft->au8HistThresh[0] = 0xd;
    pstAeSnsDft->au8HistThresh[1] = 0x28;
    pstAeSnsDft->au8HistThresh[2] = 0x60;
    pstAeSnsDft->au8HistThresh[3] = 0x80;
            
    pstAeSnsDft->u8AeCompensation = 0x38;

    pstAeSnsDft->stIntTimeAccu.enAccuType = AE_ACCURACY_LINEAR;
    pstAeSnsDft->stIntTimeAccu.f32Accuracy = 1;
    pstAeSnsDft->stIntTimeAccu.f32Offset = 0;
    pstAeSnsDft->u32MaxIntTime = gu32FullLinesStd - 2;
    pstAeSnsDft->u32MinIntTime = 2;
    pstAeSnsDft->u32MaxIntTimeTarget = 65535;
    pstAeSnsDft->u32MinIntTimeTarget = 2;


    pstAeSnsDft->stAgainAccu.enAccuType = AE_ACCURACY_DB;
    pstAeSnsDft->stAgainAccu.f32Accuracy = 6;
    pstAeSnsDft->u32MaxAgain = 4;
    pstAeSnsDft->u32MinAgain = 0;
    pstAeSnsDft->u32MaxAgainTarget = pstAeSnsDft->u32MaxAgain;
    pstAeSnsDft->u32MinAgainTarget = pstAeSnsDft->u32MinAgain;
        


    pstAeSnsDft->stDgainAccu.enAccuType = AE_ACCURACY_LINEAR;
    pstAeSnsDft->stDgainAccu.f32Accuracy =  0.0625;
    pstAeSnsDft->u32MaxDgain = 31;
    pstAeSnsDft->u32MinDgain = 16;
    pstAeSnsDft->u32MaxDgainTarget = 31;  
    pstAeSnsDft->u32MinDgainTarget = 16; 
    

    
    pstAeSnsDft->u32ISPDgainShift = 8;
    pstAeSnsDft->u32MinISPDgainTarget = 1 << pstAeSnsDft->u32ISPDgainShift;
    pstAeSnsDft->u32MaxISPDgainTarget = 2 << pstAeSnsDft->u32ISPDgainShift;


	
    pstAeSnsDft->u32LinesPer500ms = gu32FullLinesStd*30/2;

    pstAeSnsDft->enIrisType = ISP_IRIS_DC_TYPE;
    memcpy(&pstAeSnsDft->stPirisAttr, &gstPirisAttr, sizeof(ISP_PIRIS_ATTR_S));
    pstAeSnsDft->enMaxIrisFNO = ISP_IRIS_F_NO_1_4;
    pstAeSnsDft->enMinIrisFNO = ISP_IRIS_F_NO_5_6;

    /*For some SOI sensors, AERunInterval needs to be set more than 1*/
    pstAeSnsDft->u8AERunInterval = 2;
    
    return 0;
}

/* the function of sensor set fps */
static HI_VOID cmos_fps_set(HI_FLOAT f32Fps, AE_SENSOR_DEFAULT_S *pstAeSnsDft)
{
    if ((f32Fps <= 30) && (f32Fps >= 0.5))
    {
        if(SENSOR_960P_30FPS_MODE == gu8SensorImageMode_h65)
        {
            gu32FullLinesStd = VMAX_SOIH65_960P30_LINEAR * 30 / f32Fps;

            g_stSnsRegsInfo.astI2cData[3].u32Data = gu32FullLinesStd  & 0xff;
            g_stSnsRegsInfo.astI2cData[4].u32Data = (gu32FullLinesStd >> 8) & 0xff;

		//	sensor_write_register(0x22,gu32FullLines & 0xff);
		//	sensor_write_register(0x23,(gu32FullLines >> 8) & 0xff);

            pstAeSnsDft->f32Fps = f32Fps;
            pstAeSnsDft->u32MaxIntTime = gu32FullLinesStd - 2;
            pstAeSnsDft->u32FullLinesStd = gu32FullLinesStd;
            pstAeSnsDft->u32LinesPer500ms = gu32FullLinesStd * f32Fps/2;
            gu32FullLines = gu32FullLinesStd;
            pstAeSnsDft->u32FullLines = gu32FullLines;
        }
    }
    else
    {
        printf("Not support Fps: %f\n", f32Fps);
        return;
    }
  
    return;
}

static HI_VOID cmos_slow_framerate_set(HI_U32 u32FullLines,
    AE_SENSOR_DEFAULT_S *pstAeSnsDft)
{
    u32FullLines = (u32FullLines > FULL_LINES_MAX) ? FULL_LINES_MAX : u32FullLines;

    gu32FullLines = u32FullLines;
    
    g_stSnsRegsInfo.astI2cData[3].u32Data = gu32FullLines & 0xff;
    g_stSnsRegsInfo.astI2cData[4].u32Data = (gu32FullLines >> 8) & 0xff;
//	sensor_write_register(0x22,gu32FullLines & 0xff);
//	sensor_write_register(0x23,(gu32FullLines >> 8) & 0xff);
	
    pstAeSnsDft->u32FullLines = gu32FullLines;
    pstAeSnsDft->u32MaxIntTime = gu32FullLines - 2;
    
    return;
}

/* while isp notify ae to update sensor regs, ae call these funcs. */
static HI_VOID cmos_inttime_update(HI_U32 u32IntTime)
{


    g_stSnsRegsInfo.astI2cData[0].u32Data =  u32IntTime & 0xFF;
    g_stSnsRegsInfo.astI2cData[1].u32Data = (u32IntTime >> 8) & 0xFF;

	
    return;
}

static HI_U32 analog_gain_table[64] =
{
    1024, 1088, 1152, 1216, 1280, 1344, 1408, 1472, 1536, 1600, 1664, 1728, 1792, 1856, 1920, 1984, 2048,
    2176, 2304, 2432, 2560, 2688, 2816, 2944, 3072, 3200, 3328, 3456, 3584, 3712, 3840, 3968, 4096, 4352,
    4608, 4864, 5120, 5376, 5632, 5888, 6144, 6400, 6656, 6912, 7168, 7424, 7680, 7936, 8192, 8704, 9216,
    9728, 10240, 10752, 11264, 11776, 12288, 12800, 13312, 13824, 14336, 14848, 15360, 15872
};

static HI_U32 analog_gain_reg_table[64] =
{     
    0x010, 0x011, 0x012, 0x013, 0x014, 0x015, 0x016, 0x017, 0x018, 0x019, 0x01A, 0x01B, 0x01C, 0x01D,
    0x01E, 0x01F, 0x020, 0x022, 0x024, 0x026, 0x028, 0x02A, 0x02C, 0x02E, 0x030, 0x032, 0x034, 0x036,
    0x038, 0x03A, 0x03C, 0x03E, 0x040, 0x044, 0x048, 0x04C, 0x050, 0x054, 0x058, 0x05C, 0x060, 0x064,
    0x068, 0x06C, 0x070, 0x074, 0x078, 0x07C, 0x080, 0x088, 0x090, 0x098, 0x0A0, 0x0A8, 0x0B0, 0x0B8,
    0x0C0, 0x0C8, 0x0D0, 0x0D8, 0x0E0, 0x0E8, 0x0F0, 0x0F8
};


static HI_VOID cmos_again_calc_table(HI_U32 *pu32AgainLin, HI_U32 *pu32AgainDb)
{
    int i;

    if (*pu32AgainLin >= analog_gain_table[63])
    {
         *pu32AgainLin = analog_gain_table[63];
         *pu32AgainDb = 63;
         return ;
    }
    
    for (i = 1; i < 64; i++)
    {
        if (*pu32AgainLin < analog_gain_table[i])
        {
            *pu32AgainLin = analog_gain_table[i - 1];
            *pu32AgainDb = i - 1;
            break;
        }
    }

    return;

}

static HI_VOID cmos_gains_update(HI_U32 u32Again, HI_U32 u32Dgain)
{  
   HI_U8 u8High, u8Low;
    switch (u32Again)
    {
        case 0 :    /* 0db, 1 multiplies */
            u8High = 0x00;
            break;
        case 1 :    /* 6db, 2 multiplies */
            u8High = 0x10;
            break;
        case 2 :    /* 12db, 4 multiplies */
            u8High = 0x20;
            break;
        case 3 :    /* 18db, 8 multiplies */
            u8High = 0x30;
            break;
        case 4 :    /* 24db, 16 multiplies */
            u8High = 0x40;
            break;
        default:
            u8High = 0x00;
            break;
    }

    u8Low = (u32Dgain - 16) & 0xf;
    g_stSnsRegsInfo.astI2cData[2].u32Data = (u8High | u8Low);
    
    return;
}


static HI_S32 cmos_init_ae_exp_function(AE_SENSOR_EXP_FUNC_S *pstExpFuncs)
{
    memset(pstExpFuncs, 0, sizeof(AE_SENSOR_EXP_FUNC_S));

    pstExpFuncs->pfn_cmos_get_ae_default    = cmos_get_ae_default;
    pstExpFuncs->pfn_cmos_fps_set           = cmos_fps_set;
    pstExpFuncs->pfn_cmos_slow_framerate_set= cmos_slow_framerate_set;    
    pstExpFuncs->pfn_cmos_inttime_update    = cmos_inttime_update;
    pstExpFuncs->pfn_cmos_gains_update      = cmos_gains_update;
    pstExpFuncs->pfn_cmos_again_calc_table  = NULL; // cmos_again_calc_table;
    pstExpFuncs->pfn_cmos_dgain_calc_table  = NULL;

    return 0;
}


/* AWB default parameter and function */
static AWB_CCM_S g_stAwbCcm =
{  

    5000,
    {
        0x01A2,  0x807D,  0x8025,
	0x803B,  0x0198,  0x805D,
	0x8001,  0x80BB,  0x01BC,          
    },
    4000,
    {
        0x01BF,  0x809B,  0x8024,
	0x804B,  0x0197,  0x804C,
	0x8005,  0x80D5,  0x01DA,
    },
    2856,
    {
        0x01D6,  0x80B8,  0x801E,
	0x8057,  0x0175,  0x801E,
	0x8036,  0x81D1,  0x0307,

    }

};

static AWB_AGC_TABLE_S g_stAwbAgcTable =
{
    /* bvalid */
    1,
	
    /*1,  2,  4,  8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768*/
    /* saturation */   
    //{0x80,0x80,0x7a,0x78,0x70,0x68,0x60,0x58,0x50,0x48,0x40,0x38,0x38,0x38,0x38,0x38}
    {0x9C,0x96,0x91,0x80,0x70,0x60,0x58,0x50,0x40,0x38,0x38,0x38,0x38,0x38,0x38,0x38}

};

static HI_S32 cmos_get_awb_default(AWB_SENSOR_DEFAULT_S *pstAwbSnsDft)
{
    if (HI_NULL == pstAwbSnsDft)
    {
        printf("null pointer when get awb default value!\n");
        return -1;
    }

    memset(pstAwbSnsDft, 0, sizeof(AWB_SENSOR_DEFAULT_S));

    pstAwbSnsDft->u16WbRefTemp = 5200;
    pstAwbSnsDft->au16GainOffset[0] = 0x1B8;//0X1C2    
    pstAwbSnsDft->au16GainOffset[1] = 0x100;    
    pstAwbSnsDft->au16GainOffset[2] = 0x100;    
    pstAwbSnsDft->au16GainOffset[3] = 0x22D;//0X1C0    
    pstAwbSnsDft->as32WbPara[0] = 101;    //128
    pstAwbSnsDft->as32WbPara[1] = 37;    //-26
    pstAwbSnsDft->as32WbPara[2] = -119;    //-154
    pstAwbSnsDft->as32WbPara[3] = 147020;    //233501
    pstAwbSnsDft->as32WbPara[4] = 128;    
    pstAwbSnsDft->as32WbPara[5] = -99417; //-184710
    
    memcpy(&pstAwbSnsDft->stCcm, &g_stAwbCcm, sizeof(AWB_CCM_S));
    memcpy(&pstAwbSnsDft->stAgcTbl, &g_stAwbAgcTable, sizeof(AWB_AGC_TABLE_S));
    
    return 0;
}

static HI_S32 cmos_init_awb_exp_function(AWB_SENSOR_EXP_FUNC_S *pstExpFuncs)
{
    memset(pstExpFuncs, 0, sizeof(AWB_SENSOR_EXP_FUNC_S));

    pstExpFuncs->pfn_cmos_get_awb_default = cmos_get_awb_default;

    return 0;
}

#define DMNR_CALIB_CARVE_NUM_SOIH65 8

float g_coef_calib_soih65[DMNR_CALIB_CARVE_NUM_SOIH65][4] = 
{
    
	    {100.000000f, 2.000000f, 0.027490f, 11.997045f, }, 
	
		{464.000000f, 2.666518f, 0.026187f, 13.384201f, }, 
	
		{745.000000f, 2.872156f, 0.028402f, 14.261521f, }, 
	
		{1230.000000f, 3.089905f, 0.032918f, 16.007761f, }, 
	
		{1424.000000f, 3.153510f, 0.034163f, 16.493176f, }, 
	
		{2217.000000f, 3.345766f, 0.037814f, 18.170021f, }, 
	
		{3496.000000f, 3.543571f, 0.040788f, 22.397856f, }, 
	
		{5623.000000f, 3.749968f, 0.008956f, 42.169674f, }, 
    
};



static ISP_NR_ISO_PARA_TABLE_S g_stNrIsoParaTab[HI_ISP_NR_ISO_LEVEL_MAX] = 
{
     //u16Threshold//u8varStrength//u8fixStrength//u8LowFreqSlope	
       {0x5DC,       0x64,             0x28,            1 },  //100    //                      //                                                
       {0x850,       0x6E,             0x50,            2 },  //200    // ISO                  // ISO //u8LowFreqSlope
       {0xC38,       0x78,             0x78,            4 },  //400    //{400,  1200, 96,256}, //{400 , 0  }
       {0x1020,      0xAF,             0xA0,            8 },  //800    //{800,  1400, 80,256}, //{600 , 2  }
       {0x1408,      0xB9,             0xD2,            12 },  //1600   //{1600, 1200, 72,256}, //{800 , 8  }
       {0x17F0,      0xC3,             0xDC,            14 },  //3200   //{3200, 1200, 64,256}, //{1000, 12 }
       {0x17F0,      0xCD,             0xE6,            16 },  //6400   //{6400, 1100, 56,256}, //{1600, 6  }
       {1375,       70,             256-256,            0 },  //12800  //{12000,1100, 48,256}, //{2400, 0  }
       {1375,       65,             256-256,            0 },  //25600  //{36000,1100, 48,256}, //
       {1375,       70,             256-256,            0 },  //51200  //{64000,1100, 96,256}, //
       {1250,       70,             256-256,            0 },  //102400 //{82000,1000,240,256}, //
       {1250,       70,             256-256,            0 },  //204800 //                           //
       {1250,       70,             256-256,            0 },  //409600 //                           //
       {1250,       70,             256-256,            0 },  //819200 //                           //
       {1250,       70,             256-256,            0 },  //1638400//                           //
       {1250,       70,             256-256,            0 },  //3276800//                           //
};

static ISP_CMOS_DEMOSAIC_S g_stIspDemosaic =
{
	/*For Demosaic*/
	1, /*bEnable*/			
	12,/*u16VhLimit*/	
	8,/*u16VhOffset*/
	48,   /*u16VhSlope*/
	/*False Color*/
	1,    /*bFcrEnable*/
	{ 8, 8, 8, 8, 8, 8, 8, 8, 3, 0, 0, 0, 0, 0, 0, 0},    /*au8FcrStrength[ISP_AUTO_ISO_STENGTH_NUM]*/
	{24,24,24,24,24,24,24,24,24,24,24,24,24,24,24,24},    /*au8FcrThreshold[ISP_AUTO_ISO_STENGTH_NUM]*/
	/*For Ahd*/
	400, /*u16UuSlope*/	
	{512,512,512,512,512,512,512,  400,  0,0,0,0,0,0,0,0}    /*au16NpOffset[ISP_AUTO_ISO_STENGTH_NUM]*/
};

static ISP_CMOS_GE_S g_stIspGe =
{
	/*For GE*/
	0,    /*bEnable*/			
	7,    /*u8Slope*/	
	7,    /*u8Sensitivity*/
	4096, /*u16Threshold*/
	4096, /*u16SensiThreshold*/	
	{1024,1024,1024,2048,2048,2048,2048,  2048,  2048,2048,2048,2048,2048,2048,2048,2048}    /*au16Strength[ISP_AUTO_ISO_STENGTH_NUM]*/	
};

static ISP_CMOS_RGBSHARPEN_S g_stIspRgbSharpen =
{      
    {   0,   0,  0,  0,  1,  1,  1,  1,  1,  1,  1,  1,  1,  1,  1,  1},/* bEnLowLumaShoot */
    {  0x6E, 0x64, 0x5A, 0x50, 0x46, 0x3C, 0x3C, 0x3C, 0x3C, 0x3C, 0x3C, 0x3C, 0x3C, 0x3C, 0x3C, 0x3C},/*SharpenUD*/
    {  0x7D, 0x78, 0x6E, 0x64, 0x5A, 0x50, 0x46, 0x46, 0x46, 0x46, 0x46, 0x46, 0x46, 0x46, 0x46, 0x46},/*SharpenD*/
    {   0,   0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0},/*TextureNoiseThd*/
    {   0,   0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0},/*EdgeNoiseThd*/
    { 0x64,  0x5F, 0x5A, 0x55, 0x50, 0x4B, 0x46, 0x41, 0x2D, 0x2D, 0x2D, 0x2D, 0x2D, 0x2D, 0x2D, 0x2D},/*overshoot*/
    { 0x64, 0x64, 0x5F, 0x5F, 0x55, 0x50, 0x4B, 0x46, 0x41, 0x3C, 0x3C, 0x3C, 0x3C, 0x3C, 0x3C, 0x3C},/*undershoot*/

};


static ISP_CMOS_UVNR_S g_stIspUVNR = 
{
  //{100,	200,	400,	800,	1600,	3200,	6400,	12800,	25600,	51200,	102400,	204800,	409600,	819200,	1638400,	3276800};
	{1,	    2,       4,      5,      7,      48,     32,     16,     16,     16,      16,     16,     16,     16,     16,        16},      /*UVNRThreshold*/
 	{0,		0,		0,		0,		0,		0,		0,		0,		0,		1,			1,		2,		2,		2,		2,		2},  /*Coring_lutLimit*/
	{0,		0,		0,		16,		34,		34,		34,		34,		34,		34,		34,		34,		34,		34,		34,			34}  /*UVNR_blendRatio*/
};

static ISP_CMOS_DPC_S g_stCmosDpc = 
{
	//1,/*IR_channel*/
	//1,/*IR_position*/
	{252,252,252,252,252,252,252,252,252,252,252,252,252,252,252,252},/*au16Strength[16]*/
	{0,0,0,0,0,0,0,0,0x24,0x80,0x80,0x80,0xE5,0xE5,0xE5,0xE5},/*au16BlendRatio[16]*/
};


static ISP_CMOS_DRC_S g_stIspDrc =
{
    0,
    10,
    0,
    2,
    192,
    60,
    0,
    0,
    0,
    {1024,1024,1024,1024,1024,1024,1024,1024,1024,1024,1024,1024,1024,1024,1024,1024,1024,1024,1024,1024,1024,1024,1024,1024,1024,1024,1024,1024,1024,1024,1024,1024,1024}
};


static HI_U32 cmos_get_isp_default(ISP_CMOS_DEFAULT_S *pstDef)
{
    if (HI_NULL == pstDef)
    {
        printf("null pointer when get isp default value!\n");
        return -1;
    }

    memset(pstDef, 0, sizeof(ISP_CMOS_DEFAULT_S));

    memcpy(&pstDef->stDrc, &g_stIspDrc, sizeof(ISP_CMOS_DRC_S));
    memcpy(&pstDef->stDemosaic, &g_stIspDemosaic, sizeof(ISP_CMOS_DEMOSAIC_S));
    memcpy(&pstDef->stRgbSharpen, &g_stIspRgbSharpen, sizeof(ISP_CMOS_RGBSHARPEN_S));
    memcpy(&pstDef->stGe, &g_stIspGe, sizeof(ISP_CMOS_GE_S));			
  //  pstDef->stNoiseTbl.u8SensorIndex = HI_ISP_NR_SENSOR_INDEX_SOIH65;
    pstDef->stNoiseTbl.stNrCaliPara.u8CalicoefRow = DMNR_CALIB_CARVE_NUM_SOIH65;
    pstDef->stNoiseTbl.stNrCaliPara.pCalibcoef    = (HI_FLOAT (*)[4])g_coef_calib_soih65;

    memcpy(&pstDef->stNoiseTbl.stIsoParaTable[0], &g_stNrIsoParaTab[0],sizeof(ISP_NR_ISO_PARA_TABLE_S)*HI_ISP_NR_ISO_LEVEL_MAX);

    memcpy(&pstDef->stUvnr,       &g_stIspUVNR,       sizeof(ISP_CMOS_UVNR_S));
    memcpy(&pstDef->stDpc,       &g_stCmosDpc,       sizeof(ISP_CMOS_DPC_S));

    pstDef->stSensorMaxResolution.u32MaxWidth  = 1280;
    pstDef->stSensorMaxResolution.u32MaxHeight = 960;

    return 0;
}

static HI_U32 cmos_get_isp_black_level(ISP_CMOS_BLACK_LEVEL_S *pstBlackLevel)
{
    if (HI_NULL == pstBlackLevel)
    {
        printf("null pointer when get isp black level value!\n");
        return -1;
    }

    /* Don't need to update black level when iso change */
    pstBlackLevel->bUpdate = HI_FALSE;
          
    pstBlackLevel->au16BlackLevel[0] = 64;
    pstBlackLevel->au16BlackLevel[1] = 64; //58;
    pstBlackLevel->au16BlackLevel[2] = 64; //58;
    pstBlackLevel->au16BlackLevel[3] = 64;
    

    return 0;  
    
}

static HI_VOID cmos_set_pixel_detect(HI_BOOL bEnable)
{
    HI_U32 u32FullLines_5Fps = 0; 
    HI_U32 u32MaxIntTime_5Fps = 0;

    if (SENSOR_960P_30FPS_MODE == gu8SensorImageMode_h65)
    {
        u32FullLines_5Fps = (VMAX_SOIH65_960P30_LINEAR * 30) / 5;
    }
    else
    {
        return;
    }

    u32FullLines_5Fps = (u32FullLines_5Fps > FULL_LINES_MAX) ? FULL_LINES_MAX : u32FullLines_5Fps;
    u32MaxIntTime_5Fps = u32FullLines_5Fps - 2;

    if (bEnable) /* setup for ISP pixel calibration mode */
    {
        h65_write_register(0x22, u32FullLines_5Fps & 0xFF );  /* 5fps */
        h65_write_register(0x23, (u32FullLines_5Fps >> 8 ) & 0xFF);           /* 5fps */
        h65_write_register(0x01, u32MaxIntTime_5Fps & 0xFF);               /* max exposure lines */
        h65_write_register(0x02, (u32MaxIntTime_5Fps >> 8) & 0xFF);     /* max exposure lines */
        h65_write_register(0x00, 0x00);                                    /* min AG */
    }
    else /* setup for ISP 'normal mode' */
    {
        h65_write_register(0x22, gu32FullLinesStd & 0xFF);
        h65_write_register(0x23, (gu32FullLinesStd >> 8 ) & 0xFF);
        
        bInit = HI_FALSE;
    }

    return;
}

static HI_VOID cmos_set_wdr_mode(HI_U8 u8Mode)
{
    bInit = HI_FALSE;
    
    switch(u8Mode)
    {
        case WDR_MODE_NONE:
            if (SENSOR_960P_30FPS_MODE == gu8SensorImageMode_h65)
            {
                gu32FullLinesStd = VMAX_SOIH65_960P30_LINEAR;
            }
            genSensorMode_h65 = WDR_MODE_NONE;
            printf("linear mode\n");
        break;
        default:
            printf("NOT support this mode!\n");
            return;
        break;
    }
    
    return;
}

static HI_U32 cmos_get_sns_regs_info(ISP_SNS_REGS_INFO_S *pstSnsRegsInfo)
{
    HI_S32 i;

    if (HI_FALSE == bInit)
    {
        g_stSnsRegsInfo.enSnsType = ISP_SNS_I2C_TYPE;
        g_stSnsRegsInfo.u8Cfg2ValidDelayMax = 2;		
        g_stSnsRegsInfo.u32RegNum = 5;
	
        for (i=0; i<g_stSnsRegsInfo.u32RegNum; i++)
        {	
            g_stSnsRegsInfo.astI2cData[i].bUpdate = HI_TRUE;
            g_stSnsRegsInfo.astI2cData[i].u8DevAddr = h65_i2c_addr;
            g_stSnsRegsInfo.astI2cData[i].u32AddrByteNum = h65_addr_byte;
            g_stSnsRegsInfo.astI2cData[i].u32DataByteNum = h65_data_byte;
        }

        g_stSnsRegsInfo.astI2cData[0].u8DelayFrmNum = 0;         //exposure time: 
        g_stSnsRegsInfo.astI2cData[0].u32RegAddr = 0x01;
        g_stSnsRegsInfo.astI2cData[1].u8DelayFrmNum = 0;
        g_stSnsRegsInfo.astI2cData[1].u32RegAddr = 0x02;
		
        g_stSnsRegsInfo.astI2cData[2].u8DelayFrmNum = 0;	// gain
        g_stSnsRegsInfo.astI2cData[2].u32RegAddr = 0x00;


        g_stSnsRegsInfo.astI2cData[3].u8DelayFrmNum = 0;       //VTS
        g_stSnsRegsInfo.astI2cData[3].u32RegAddr = 0x22;
        g_stSnsRegsInfo.astI2cData[4].u8DelayFrmNum = 0;       
        g_stSnsRegsInfo.astI2cData[4].u32RegAddr = 0x23;

        bInit = HI_TRUE;
    }
    else    
    {        
        for (i = 0; i < g_stSnsRegsInfo.u32RegNum; i++)        
        {            
            if (g_stSnsRegsInfo.astI2cData[i].u32Data == g_stPreSnsRegsInfo.astI2cData[i].u32Data)            
            {                
                g_stSnsRegsInfo.astI2cData[i].bUpdate = HI_FALSE;
            }            
            else            
            {                
                g_stSnsRegsInfo.astI2cData[i].bUpdate = HI_TRUE;
            }        
        }    
    }

    if (HI_NULL == pstSnsRegsInfo)
    {
        printf("null pointer when get sns reg info!\n");
        return -1;
    }

    memcpy(pstSnsRegsInfo, &g_stSnsRegsInfo, sizeof(ISP_SNS_REGS_INFO_S)); 
    memcpy(&g_stPreSnsRegsInfo, &g_stSnsRegsInfo, sizeof(ISP_SNS_REGS_INFO_S)); 

    return 0;
}

/*
static int sensor_set_inifile_path(const char *pcPath)
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
*/

static HI_S32 cmos_set_image_mode(ISP_CMOS_SENSOR_IMAGE_MODE_S *pstSensorImageMode)
{
    HI_U8 u8SensorImageMode = gu8SensorImageMode_h65;

    bInit = HI_FALSE;
    
    if (HI_NULL == pstSensorImageMode )
    {
        printf("null pointer when set image mode\n");
        return -1;
    }

    if ((pstSensorImageMode->u16Width <= 1280) && (pstSensorImageMode->u16Height <= 960))
    {
        if (WDR_MODE_NONE == genSensorMode_h65)
        {
            if (pstSensorImageMode->f32Fps <= 30)
            {
                u8SensorImageMode = SENSOR_960P_30FPS_MODE;
            }
            else
            {
                printf("Not support! Width:%d, Height:%d, Fps:%f, WDRMode:%d\n", 
                    pstSensorImageMode->u16Width, 
                    pstSensorImageMode->u16Height,
                    pstSensorImageMode->f32Fps,
                    genSensorMode_h65);

                return -1;
            }
        }
        else
        {
            printf("Not support! Width:%d, Height:%d, Fps:%f, WDRMode:%d\n", 
                pstSensorImageMode->u16Width, 
                pstSensorImageMode->u16Height,
                pstSensorImageMode->f32Fps,
                genSensorMode_h65);

            return -1;
        }
    }
    else
    {
        printf("Not support! Width:%d, Height:%d, Fps:%f, WDRMode:%d\n", 
            pstSensorImageMode->u16Width, 
            pstSensorImageMode->u16Height,
            pstSensorImageMode->f32Fps,
            genSensorMode_h65);

        return -1;
    }

    /* Sensor first init */
    if (HI_FALSE == bSensorInit_h65)
    {
        gu8SensorImageMode_h65 = u8SensorImageMode;
        
        return 0;
    }

    /* Switch SensorImageMode */
    if (u8SensorImageMode == gu8SensorImageMode_h65)
    {
        /* Don't need to switch SensorImageMode */
        return -1;
    }
    
    gu8SensorImageMode_h65 = u8SensorImageMode;

    return 0;
}

static HI_VOID sensor_global_init()
{   
    gu8SensorImageMode_h65 = SENSOR_960P_30FPS_MODE;
    genSensorMode_h65 = WDR_MODE_NONE;
    gu32FullLinesStd = VMAX_SOIH65_960P30_LINEAR; 
    gu32FullLines = VMAX_SOIH65_960P30_LINEAR;
    bInit = HI_FALSE;
    bSensorInit_h65 = HI_FALSE; 

    memset(&g_stSnsRegsInfo, 0, sizeof(ISP_SNS_REGS_INFO_S));
    memset(&g_stPreSnsRegsInfo, 0, sizeof(ISP_SNS_REGS_INFO_S));
}

static HI_S32 cmos_init_sensor_exp_function(ISP_SENSOR_EXP_FUNC_S *pstSensorExpFunc)
{
    memset(pstSensorExpFunc, 0, sizeof(ISP_SENSOR_EXP_FUNC_S));

    pstSensorExpFunc->pfn_cmos_sensor_init = h65_init;
    pstSensorExpFunc->pfn_cmos_sensor_exit = h65_exit;
    pstSensorExpFunc->pfn_cmos_sensor_global_init = sensor_global_init;
    pstSensorExpFunc->pfn_cmos_set_image_mode = cmos_set_image_mode;
    pstSensorExpFunc->pfn_cmos_set_wdr_mode = cmos_set_wdr_mode;
    
    pstSensorExpFunc->pfn_cmos_get_isp_default = cmos_get_isp_default;
    pstSensorExpFunc->pfn_cmos_get_isp_black_level = cmos_get_isp_black_level;
    pstSensorExpFunc->pfn_cmos_set_pixel_detect = cmos_set_pixel_detect;
    pstSensorExpFunc->pfn_cmos_get_sns_reg_info = cmos_get_sns_regs_info;

    return 0;
}

/****************************************************************************
 * callback structure                                                       *
 ****************************************************************************/
 
int h65_register_callback(void)
{
    ISP_DEV IspDev = 0;
    HI_S32 s32Ret;
    ALG_LIB_S stLib;
    ISP_SENSOR_REGISTER_S stIspRegister;
    AE_SENSOR_REGISTER_S  stAeRegister;
    AWB_SENSOR_REGISTER_S stAwbRegister;

    cmos_init_sensor_exp_function(&stIspRegister.stSnsExp);
    s32Ret = HI_MPI_ISP_SensorRegCallBack(IspDev, SOIH65_ID, &stIspRegister);
    if (s32Ret)
    {
        printf("sensor register callback function failed!\n");
        return s32Ret;
    }
    
    stLib.s32Id = 0;
    strncpy(stLib.acLibName, HI_AE_LIB_NAME, sizeof(HI_AE_LIB_NAME));
    cmos_init_ae_exp_function(&stAeRegister.stSnsExp);
    s32Ret = HI_MPI_AE_SensorRegCallBack(IspDev, &stLib, SOIH65_ID, &stAeRegister);
    if (s32Ret)
    {
        printf("sensor register callback function to ae lib failed!\n");
        return s32Ret;
    }

    stLib.s32Id = 0;
    strncpy(stLib.acLibName, HI_AWB_LIB_NAME, sizeof(HI_AWB_LIB_NAME));
    cmos_init_awb_exp_function(&stAwbRegister.stSnsExp);
    s32Ret = HI_MPI_AWB_SensorRegCallBack(IspDev, &stLib, SOIH65_ID, &stAwbRegister);
    if (s32Ret)
    {
        printf("sensor register callback function to ae lib failed!\n");
        return s32Ret;
    }
    
    return 0;
}

int h65_unregister_callback(void)
{
    ISP_DEV IspDev = 0;
    HI_S32 s32Ret;
    ALG_LIB_S stLib;

    s32Ret = HI_MPI_ISP_SensorUnRegCallBack(IspDev, SOIH65_ID);
    if (s32Ret)
    {
        printf("sensor unregister callback function failed!\n");
        return s32Ret;
    }
    
    stLib.s32Id = 0;
    strncpy(stLib.acLibName, HI_AE_LIB_NAME, sizeof(HI_AE_LIB_NAME));
    s32Ret = HI_MPI_AE_SensorUnRegCallBack(IspDev, &stLib, SOIH65_ID);
    if (s32Ret)
    {
        printf("sensor unregister callback function to ae lib failed!\n");
        return s32Ret;
    }

    stLib.s32Id = 0;
    strncpy(stLib.acLibName, HI_AWB_LIB_NAME, sizeof(HI_AWB_LIB_NAME));
    s32Ret = HI_MPI_AWB_SensorUnRegCallBack(IspDev, &stLib, SOIH65_ID);
    if (s32Ret)
    {
        printf("sensor unregister callback function to ae lib failed!\n");
        return s32Ret;
    }
    
    return 0;
}

#ifdef __cplusplus
#if __cplusplus
}
#endif
#endif /* End of #ifdef __cplusplus */

#endif /* __SOI_H65_CMOS_H_ */
