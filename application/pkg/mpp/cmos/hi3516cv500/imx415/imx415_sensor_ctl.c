
/******************************************************************************

  Copyright (C), 2016, Hisilicon Tech. Co., Ltd.

 ******************************************************************************
  File Name     : imx415_sensor_ctl.c

  Version       : Initial Draft
  Author        : Hisilicon multimedia software group
  Created       : 2013/11/07
  Description   :
  History       :
  1.Date        : 2013/11/07
    Author      :
    Modification: Created file

******************************************************************************/
#include <stdio.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <sys/ioctl.h>
#include <fcntl.h>
#include <unistd.h>

#include <ctype.h>
//#include <linux/fb.h>
#include <sys/mman.h>
#include <memory.h>

//#include "hi_comm_video.h"
#include "hi_sns_ctrl.h"
#include "imx415_cmos_priv.h"

#ifdef HI_GPIO_I2C
#include "gpioi2c_ex.h"
#else
#include "hi_i2c.h"
//#include "drv_i2c.h"
#endif


const unsigned char imx415_i2c_addr     =    0x34;        /* I2C Address of imx415 */
const unsigned int  imx415_addr_byte    =    2;
const unsigned int  imx415_data_byte    =    1;
static int g_fd[ISP_MAX_PIPE_NUM] = {[0 ... (ISP_MAX_PIPE_NUM - 1)] = -1};
static HI_BOOL g_bStandby[ISP_MAX_PIPE_NUM] = {0};

extern WDR_MODE_E genSensorMode;
extern HI_U8 gu8SensorImageMode;
extern HI_BOOL bSensorInit;
extern const IMX415_VIDEO_MODE_TBL_S g_astImx415ModeTbl[];

extern ISP_SNS_STATE_S   *g_pastImx415[ISP_MAX_PIPE_NUM];
extern ISP_SNS_COMMBUS_U  g_aunImx415BusInfo[];
#ifndef NA
#define NA 0xFFFF
#endif

int imx415_i2c_init(VI_PIPE ViPipe)
{
    int ret;
    HI_U8 u8DevNum;
    char acDevFile[16];

    if (g_fd[ViPipe] >= 0) {
        return HI_SUCCESS;
    }

    u8DevNum = g_aunImx415BusInfo[ViPipe].s8I2cDev;

    snprintf(acDevFile, sizeof(acDevFile),  "/dev/i2c-%u", u8DevNum);

    g_fd[ViPipe] = open(acDevFile, O_RDWR, S_IRUSR | S_IWUSR);

    if (g_fd[ViPipe] < 0) {
        ISP_ERR_TRACE("Open /dev/hi_i2c_drv-%u error!\n", u8DevNum);
        return HI_FAILURE;
    }

    ret = ioctl(g_fd[ViPipe], I2C_SLAVE_FORCE, (imx415_i2c_addr >> 1));

    if (ret < 0) {
        ISP_ERR_TRACE("I2C_SLAVE_FORCE error!\n");
        close(g_fd[ViPipe]);
        g_fd[ViPipe] = -1;
        return ret;
    }

    return HI_SUCCESS;
}

int imx415_i2c_exit(VI_PIPE ViPipe)
{
    if (g_fd[ViPipe] >= 0) {
        close(g_fd[ViPipe]);
        g_fd[ViPipe] = -1;
        return HI_SUCCESS;
    }

    return HI_FAILURE;
}

int imx415_read_register(VI_PIPE ViPipe, int addr)
{
    return HI_SUCCESS;
}
int imx415_write_register(VI_PIPE ViPipe, int addr, int data)
{
    int ret;
    int idx = 0;
    char buf[8];

    if (g_fd[ViPipe] < 0) {
        return HI_SUCCESS;
    }

    if (imx415_addr_byte == 2) {
        buf[idx] = (addr >> 8) & 0xff;
        idx++;
        buf[idx] = addr & 0xff;
        idx++;
    } else {
        buf[idx] = addr & 0xff;
        idx++;
    }

    if (imx415_data_byte == 2) {
        buf[idx] = (data >> 8) & 0xff;
        idx++;
        buf[idx] = data & 0xff;
        idx++;
    } else {
        buf[idx] = data & 0xff;
        idx++;
    }

    ret = write(g_fd[ViPipe], buf, (imx415_addr_byte + imx415_data_byte));

    if (ret < 0) {
        ISP_ERR_TRACE("I2C_WRITE DATA error!\n");
        return HI_FAILURE;
    }

    return HI_SUCCESS;
}

void imx415_standby(VI_PIPE ViPipe)
{
    imx415_write_register(ViPipe, 0x3000, 0x00);
    g_bStandby[ViPipe] = HI_TRUE;
    return;
}

void imx415_restart(VI_PIPE ViPipe)
{
    imx415_write_register(ViPipe, 0x3000, 0x01);
    g_bStandby[ViPipe] = HI_FALSE;
    return;
}

static void delay_ms(int ms)
{
    usleep(ms * 1000);
}

void imx415_mirror_flip(VI_PIPE ViPipe, ISP_SNS_MIRRORFLIP_TYPE_E eSnsMirrorFlip)
{
    switch (eSnsMirrorFlip) {
        default:
        case ISP_SNS_NORMAL:
            imx415_write_register(ViPipe, 0x3030, 0x0);
            break;

        case ISP_SNS_MIRROR:
            imx415_write_register(ViPipe, 0x3030, 0x01);
            break;

        case ISP_SNS_FLIP:
            imx415_write_register(ViPipe, 0x3030, 0x02);
            break;

        case ISP_SNS_MIRROR_FLIP:
            imx415_write_register(ViPipe, 0x3030, 0x03);
            break;
    }

    return;
}

static const HI_U16 gs_au16SensorCfgSeq[][IMX415_MODE_BUTT + 1] = {
    {0x7F, 0x7F, 0x3008},
    {0x5B, 0x5B, 0x300A},
    {0x4C, 0x72, 0x3028},
    {0x04, 0x06, 0x3029},
    {0x05, 0x07, 0x3033},
    {0x08, 0x08, 0x3050},
    {0x00, 0x00, 0x30C1},
    {0x24, 0x24, 0x3116},
    {0xC0, 0x80, 0x3118},
    {0x24, 0x24, 0x311E},
    {0x21, 0x21, 0x32D4},
    {0xA1, 0xA1, 0x32EC},
    {0x7F, 0x7F, 0x3452},
    {0x03, 0x03, 0x3453},
    {0x04, 0x04, 0x358A},
    {0x02, 0x02, 0x35A1},
    {0x0C, 0x0C, 0x36BC},
    {0x53, 0x53, 0x36CC},
    {0x00, 0x00, 0x36CD},
    {0x3C, 0x3C, 0x36CE},
    {0x8C, 0x8C, 0x36D0},
    {0x00, 0x00, 0x36D1},
    {0x71, 0x71, 0x36D2},
    {0x3C, 0x3C, 0x36D4},
    {0x53, 0x53, 0x36D6},
    {0x00, 0x00, 0x36D7},
    {0x71, 0x71, 0x36D8},
    {0x8C, 0x8C, 0x36DA},
    {0x00, 0x00, 0x36DB},
    {0x02, 0x02, 0x3724},
    {0x02, 0x02, 0x3726},
    {0x02, 0x02, 0x3732},
    {0x03, 0x03, 0x3734},
    {0x03, 0x03, 0x3736},
    {0x03, 0x03, 0x3742},
    {0xE0, 0xE0, 0x3862},
    {0x30, 0x30, 0x38CC},
    {0x2F, 0x2F, 0x38CD},
    {0x0C, 0x0C, 0x395C},
    {0xD1, 0xD1, 0x3A42},
    {0x77, 0x77, 0x3A4C},
    {0x02, 0x02, 0x3AE0},
    {0x0C, 0x0C, 0x3AEC},
    {0x2E, 0x2E, 0x3B00},
    {0x29, 0x29, 0x3B06},
    {0x25, 0x25, 0x3B98},
    {0x21, 0x21, 0x3B99},
    {0x13, 0x13, 0x3B9B},
    {0x13, 0x13, 0x3B9C},
    {0x13, 0x13, 0x3B9D},
    {0x13, 0x13, 0x3B9E},
    {0x00, 0x00, 0x3BA1},
    {0x06, 0x06, 0x3BA2},
    {0x0B, 0x0B, 0x3BA3},
    {0x10, 0x10, 0x3BA4},
    {0x14, 0x14, 0x3BA5},
    {0x18, 0x18, 0x3BA6},
    {0x1A, 0x1A, 0x3BA7},
    {0x1A, 0x1A, 0x3BA8},
    {0x1A, 0x1A, 0x3BA9},
    {0xED, 0xED, 0x3BAC},
    {0x01, 0x01, 0x3BAD},
    {0xF6, 0xF6, 0x3BAE},
    {0x02, 0x02, 0x3BAF},
    {0xA2, 0xA2, 0x3BB0},
    {0x03, 0x03, 0x3BB1},
    {0xE0, 0xE0, 0x3BB2},
    {0x03, 0x03, 0x3BB3},
    {0xE0, 0xE0, 0x3BB4},
    {0x03, 0x03, 0x3BB5},
    {0xE0, 0xE0, 0x3BB6},
    {0x03, 0x03, 0x3BB7},
    {0xE0, 0xE0, 0x3BB8},
    {0xE0, 0xE0, 0x3BBA},
    {0xDA, 0xDA, 0x3BBC},
    {0x88, 0x88, 0x3BBE},
    {0x44, 0x44, 0x3BC0},
    {0x7B, 0x7B, 0x3BC2},
    {0xA2, 0xA2, 0x3BC4},
    {0xBD, 0xBD, 0x3BC8},
    {0xBD, 0xBD, 0x3BCA},
    {0x48, 0x48, 0x4004},
    {0x09, 0x09, 0x4005},
    {0x00, 0x00, 0x400C},
    {0x7F, 0x67, 0x4018},
    {0x37, 0x27, 0x401A},
    {0x37, 0x27, 0x401C},
    {0xF7, 0xB7, 0x401E},
    {0x00, 0x00, 0x401F},
    {0x3F, 0x2F, 0x4020},
    {0x6F, 0x4F, 0x4022},
    {0x3F, 0x2F, 0x4024},
    {0x5F, 0x47, 0x4026},
    {0x2F, 0x27, 0x4028},
    {0x01, 0x01, 0x4074},
};

void imx415_init(VI_PIPE ViPipe)
{
    HI_U32 i;
    HI_U8 u8ImgMode;
    HI_BOOL bInit;

    bInit       = g_pastImx415[ViPipe]->bInit;
    u8ImgMode   = g_pastImx415[ViPipe]->u8ImgMode;
    HI_U8 u16RegData;
    HI_U16 u16RegAddr;
    HI_U32 u32SeqEntries;

    if (bInit == HI_FALSE) {
        /* 2. sensor i2c init */
        imx415_i2c_init(ViPipe);
    }

    u32SeqEntries = sizeof(gs_au16SensorCfgSeq) / sizeof(gs_au16SensorCfgSeq[0]);

    for ( i = 0 ; i < u32SeqEntries; i++ ) {
        u16RegAddr = gs_au16SensorCfgSeq[i][IMX415_MODE_BUTT];
        u16RegData = gs_au16SensorCfgSeq[i][u8ImgMode];

        if (u16RegData == NA) {
            continue;
        }

        imx415_write_register(ViPipe, u16RegAddr, u16RegData);
    }

    for (i = 0; i < g_pastImx415[ViPipe]->astRegsInfo[0].u32RegNum; i++) {
        imx415_write_register(ViPipe, g_pastImx415[ViPipe]->astRegsInfo[0].astI2cData[i].u32RegAddr, g_pastImx415[ViPipe]->astRegsInfo[0].astI2cData[i].u32Data);
    }
    imx415_write_register(ViPipe, 0x3000, 0x00); //Standby Cancel
    delay_ms(24);
    imx415_write_register(ViPipe, 0x3002, 0x00);

    g_pastImx415[ViPipe]->bInit = HI_TRUE;
    //imx415_restart(ViPipe);

    printf("IMX415 %s init succuss!\n", g_astImx415ModeTbl[u8ImgMode].pszModeName);

    return ;

}

void imx415_exit(VI_PIPE ViPipe)
{
    imx415_i2c_exit(ViPipe);
    g_bStandby[ViPipe] = HI_FALSE;

    return;
}


