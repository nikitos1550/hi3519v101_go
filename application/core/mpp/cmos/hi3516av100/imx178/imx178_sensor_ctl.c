/******************************************************************************

  Copyright (C), 2001-2013, Hisilicon Tech. Co., Ltd.

 ******************************************************************************
  File Name     : imx178_sensor_ctl.c
  Version       : Initial Draft
  Author        : Hisilicon BVT ISP group
  Created       : 2014/02/20
  Description   : Sony IMX178 sensor driver
  History       :
  1.Date        : 2014/02/20
  Author        : yy
  Modification  : Created file

******************************************************************************/

#include <stdio.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <sys/ioctl.h>
#include <fcntl.h>
#include <unistd.h>

#include "hi_comm_video.h"
#include "hi_math.h"

#ifdef HI_GPIO_I2C
#include "gpioi2c_ex.h"
#else
#include "hi_i2c.h"
#endif

const unsigned char imx178_i2c_addr     =    0x34; /* I2C Address of IMX178 */
const unsigned int  imx178_addr_byte    =    2;
const unsigned int  imx178_data_byte    =    1;
static int g_fd = -1;

extern HI_U8 gu8SensorImageMode_imx178;
extern HI_BOOL bSensorInit_imx178;

//int sensor_i2c_init(void)
int imx178_i2c_init(void)
{
    if (g_fd >= 0)
    {
        return 0;
    }
#ifdef HI_GPIO_I2C
    int ret;

    g_fd = open("/dev/gpioi2c_ex", 0);
    if(g_fd < 0)
    {
        printf("Open gpioi2c_ex error!\n");
        return -1;
    }
#else
    int ret;

    g_fd = open("/dev/i2c-0", O_RDWR);
    if(g_fd < 0)
    {
        printf("Open /dev/i2c-0 error!\n");
        return -1;
    }

    ret = ioctl(g_fd, I2C_SLAVE_FORCE, imx178_i2c_addr);
    if (ret < 0)
    {
        printf("CMD_SET_DEV error!\n");
        return ret;
    }
#endif

    return 0;
}

//int sensor_i2c_exit(void)
int imx178_i2c_exit(void)
{
    if (g_fd >= 0)
    {
        close(g_fd);
        g_fd = -1;
        return 0;
    }
    return -1;
}

//int sensor_read_register(int addr)
int imx178_read_register(int addr)
{
    // TODO:

    return 0;
}

//int sensor_write_register(int addr, int data)
int imx178_write_register(int addr, int data)
{
#ifdef HI_GPIO_I2C
    i2c_data.dev_addr = imx178_i2c_addr;
    i2c_data.reg_addr = addr;
    i2c_data.addr_byte_num = imx178_addr_byte;
    i2c_data.data = data;
    i2c_data.data_byte_num = imx178_data_byte;

    ret = ioctl(g_fd, GPIO_I2C_WRITE, &i2c_data);

    if (ret)
    {
        printf("GPIO-I2C write faild!\n");
        return ret;
    }
#else
    int idx = 0;
    int ret;
    char buf[8];

    buf[idx++] = addr & 0xFF;
    if (imx178_addr_byte == 2)
    {
        ret = ioctl(g_fd, I2C_16BIT_REG, 1);
        buf[idx++] = addr >> 8;
    }
    else
    {
        ret = ioctl(g_fd, I2C_16BIT_REG, 0);
    }

    if (ret < 0)
    {
        printf("CMD_SET_REG_WIDTH error!\n");
        return -1;
    }

    buf[idx++] = data;
    if (imx178_data_byte == 2)
    {
        ret = ioctl(g_fd, I2C_16BIT_DATA, 1);
        buf[idx++] = data >> 8;
    }
    else
    {
        ret = ioctl(g_fd, I2C_16BIT_DATA, 0);
    }

    if (ret)
    {
        printf("hi_i2c write faild!\n");
        return -1;
    }

    ret = write(g_fd, buf, idx);
    if(ret < 0)
    {
        printf("I2C_WRITE error!\n");
        return -1;
    }
#endif
    return 0;
}

static void delay_ms(int ms) {
    hi_usleep(ms*1000);
}

//void sensor_prog(int* rom)
void imx178_prog(int* rom)
{
    int i = 0;
    while (1) {
        int lookup = rom[i++];
        int addr = (lookup >> 16) & 0xFFFF;
        int data = lookup & 0xFFFF;
        if (addr == 0xFFFE) {
            delay_ms(data);
        } else if (addr == 0xFFFF) {
            return;
        } else {
            imx178_write_register(addr, data);
        }
    }
}

//void sensor_init_5M30();
//void sensor_init_1080p60();

static void imx178_init_5M30();
static void imx178_init_1080p60();

//void sensor_init()
void imx178_init()
{
    imx178_i2c_init();

    if (1 == gu8SensorImageMode_imx178)    /* SENSOR_5M_30FPS_MODE */
    {
        imx178_init_5M30();
        bSensorInit_imx178 = HI_TRUE;
    }
    else if (2 == gu8SensorImageMode_imx178) /* SENSOR_1080P_60FPS_MODE */
    {
        imx178_init_1080p60();
        bSensorInit_imx178 = HI_TRUE;
    }
    else
    {
        printf("Not support this mode\n");
    }
}

//void sensor_exit()
void imx178_exit()
{
    imx178_i2c_exit();

    return;
}

//void sensor_init_5M30()
static void imx178_init_5M30()
{
    /* imx178 5M@30fps */
    imx178_write_register (0x3000, 0x07); /* standby */

    imx178_write_register (0x300E, 0x00); /* MODE, Window cropping 5.0M (4:3) */
    imx178_write_register (0x300F, 0x10); /* WINMODE, Window cropping 5.0M (4:3) */
    imx178_write_register (0x3010, 0x00); /* TCYCLE */
    imx178_write_register (0x3066, 0x06); /* VNDMY */
    imx178_write_register (0x302C, 0xF4);
    imx178_write_register (0x302D, 0x08); /* VMAX */
    imx178_write_register (0x302F, 0xE9);
    imx178_write_register (0x3030, 0x03); /* HMAX */
    imx178_write_register (0x300D, 0x05); /* ADBIT, ADBITFREQ  (ADC 12-bit) */
    imx178_write_register (0x3059, 0x31); /* ODBIT, OPORTSEL   (12-BIT) */
    imx178_write_register (0x3004, 0x03); /* STBLVDS, 4CH ACTIVE */

    /* register setting details */
    imx178_write_register (0x3101, 0x30); /* FREQ[1:0] */

    /* FREQ setting (INCK=27MHz) */
    imx178_write_register (0x310C, 0x00);
    imx178_write_register (0x33BE, 0x21);
    imx178_write_register (0x33BF, 0x21);
    imx178_write_register (0x33C0, 0x2C);
    imx178_write_register (0x33C1, 0x2C);
    imx178_write_register (0x33C2, 0x21);
    imx178_write_register (0x33C3, 0x2C);
    imx178_write_register (0x33C4, 0x2C);
    imx178_write_register (0x33C5, 0x00);
    imx178_write_register (0x311C, 0x34);
    imx178_write_register (0x311D, 0x28);
    imx178_write_register (0x311E, 0xAB);
    imx178_write_register (0x311F, 0x00);
    imx178_write_register (0x3120, 0x95);
    imx178_write_register (0x3121, 0x00);
    imx178_write_register (0x3122, 0xB4);
    imx178_write_register (0x3123, 0x00);
    imx178_write_register (0x3124, 0x8c);
    imx178_write_register (0x3125, 0x02);
    imx178_write_register (0x312D, 0x03);
    imx178_write_register (0x312E, 0x0C);
    imx178_write_register (0x312F, 0x28);
    imx178_write_register (0x3131, 0x2D);
    imx178_write_register (0x3132, 0x00);
    imx178_write_register (0x3133, 0xB4);
    imx178_write_register (0x3134, 0x00);
    imx178_write_register (0x3137, 0x50);
    imx178_write_register (0x3138, 0x08);
    imx178_write_register (0x3139, 0x00);
    imx178_write_register (0x313A, 0x07);
    imx178_write_register (0x313D, 0x05);
    imx178_write_register (0x3140, 0x06);
    imx178_write_register (0x3220, 0x8B);
    imx178_write_register (0x3221, 0x00);
    imx178_write_register (0x3222, 0x74);
    imx178_write_register (0x3223, 0x00);
    imx178_write_register (0x3226, 0xC2);
    imx178_write_register (0x3227, 0x00);
    imx178_write_register (0x32A9, 0x1B);
    imx178_write_register (0x32AA, 0x00);
    imx178_write_register (0x32B3, 0x0E);
    imx178_write_register (0x32B4, 0x00);
    imx178_write_register (0x33D6, 0x16);
    imx178_write_register (0x33D7, 0x15);
    imx178_write_register (0x33D8, 0x14);
    imx178_write_register (0x33D9, 0x10);
    imx178_write_register (0x33DA, 0x08);

    /* registers must be changed */
    imx178_write_register (0x3011, 0x00);
    imx178_write_register (0x301B, 0x00);
    imx178_write_register (0x3037, 0x08);
    imx178_write_register (0x3038, 0x00);
    imx178_write_register (0x3039, 0x00);
    imx178_write_register (0x30AD, 0x49);
    imx178_write_register (0x30AF, 0x54);
    imx178_write_register (0x30B0, 0x33);
    imx178_write_register (0x30B3, 0x0A);
    imx178_write_register (0x30C4, 0x30);
    imx178_write_register (0x3103, 0x03);
    imx178_write_register (0x3104, 0x08);
    imx178_write_register (0x3107, 0x10);
    imx178_write_register (0x310F, 0x01);
    imx178_write_register (0x32E5, 0x06);
    imx178_write_register (0x32E6, 0x00);
    imx178_write_register (0x32E7, 0x1F);
    imx178_write_register (0x32E8, 0x00);
    imx178_write_register (0x32E9, 0x00);
    imx178_write_register (0x32EA, 0x00);
    imx178_write_register (0x32EB, 0x00);
    imx178_write_register (0x32EC, 0x00);
    imx178_write_register (0x32EE, 0x00);
    imx178_write_register (0x32F2, 0x02);
    imx178_write_register (0x32F4, 0x00);
    imx178_write_register (0x32F5, 0x00);
    imx178_write_register (0x32F6, 0x00);
    imx178_write_register (0x32F7, 0x00);
    imx178_write_register (0x32F8, 0x00);
    imx178_write_register (0x32FC, 0x02);
    imx178_write_register (0x3310, 0x11);
    imx178_write_register (0x3338, 0x81);
    imx178_write_register (0x333D, 0x00);
    imx178_write_register (0x3362, 0x00);
    imx178_write_register (0x336B, 0x02);
    imx178_write_register (0x336E, 0x11);
    imx178_write_register (0x33B4, 0xFE);
    imx178_write_register (0x33B5, 0x06);
    imx178_write_register (0x33B9, 0x00);

    /*shutter and gain */
    imx178_write_register (0x3034, 0x08);
    imx178_write_register (0x3035, 0x00); /* SHS1 */
    imx178_write_register (0x301F, 0xA0);
    imx178_write_register (0x3020, 0x00); /* GAIN */

    imx178_write_register (0x3000, 0x00); /* standby */
    imx178_write_register (0x3008, 0x00); /* master mode start */
    imx178_write_register (0x305E, 0x0A); /* XVSOUTSEL XHSOUTSEL */
    imx178_write_register (0x3015, 0xC8); /* BLKLEVEL */

    printf("-------Sony IMX178 Sensor 5M30fps Initial OK!-------\n");

}

//void sensor_init_1080p60()
static void imx178_init_1080p60()
{
    /* imx178 1080@60fps */
    imx178_write_register (0x3000, 0x07); /* standby */

    imx178_write_register (0x300E, 0x01); /* MODE, Window cropping 5.0M (4:3) */
    imx178_write_register (0x300F, 0x00); /* WINMODE, Window cropping 5.0M (4:3) */
    imx178_write_register (0x3010, 0x00); /* TCYCLE */
    imx178_write_register (0x3066, 0x03); /* VNDMY */
    imx178_write_register (0x302C, 0xF8);
    imx178_write_register (0x302D, 0x05); /* VMAX */
    imx178_write_register (0x302F, 0xEE);
    imx178_write_register (0x3030, 0x02); /* HMAX */
    imx178_write_register (0x300D, 0x05); /* ADBIT, ADBITFREQ  (ADC 12-bit) */
    imx178_write_register (0x3059, 0x31); /* ODBIT, OPORTSEL   (12-BIT) */
    imx178_write_register (0x3004, 0x03); /* STBLVDS, 4CH ACTIVE */

    /* register setting details */
    imx178_write_register (0x3101, 0x30); /* FREQ[1:0] */

    /* FREQ setting (INCK=37.125MHz) */
    imx178_write_register (0x310C, 0x00);
    imx178_write_register (0x33BE, 0x21);
    imx178_write_register (0x33BF, 0x21);
    imx178_write_register (0x33C0, 0x2C);
    imx178_write_register (0x33C1, 0x2C);
    imx178_write_register (0x33C2, 0x21);
    imx178_write_register (0x33C3, 0x2C);
    imx178_write_register (0x33C4, 0x2C);
    imx178_write_register (0x33C5, 0x00);
    imx178_write_register (0x311C, 0x34);
    imx178_write_register (0x311D, 0x28);
    imx178_write_register (0x311E, 0xAB);
    imx178_write_register (0x311F, 0x00);
    imx178_write_register (0x3120, 0x95);
    imx178_write_register (0x3121, 0x00);
    imx178_write_register (0x3122, 0xB4);
    imx178_write_register (0x3123, 0x00);
    imx178_write_register (0x3124, 0x8c);
    imx178_write_register (0x3125, 0x02);
    imx178_write_register (0x312D, 0x03);
    imx178_write_register (0x312E, 0x0C);
    imx178_write_register (0x312F, 0x28);
    imx178_write_register (0x3131, 0x2D);
    imx178_write_register (0x3132, 0x00);
    imx178_write_register (0x3133, 0xB4);
    imx178_write_register (0x3134, 0x00);
    imx178_write_register (0x3137, 0x50);
    imx178_write_register (0x3138, 0x08);
    imx178_write_register (0x3139, 0x00);
    imx178_write_register (0x313A, 0x07);
    imx178_write_register (0x313D, 0x05);
    imx178_write_register (0x3140, 0x06);
    imx178_write_register (0x3220, 0x8B);
    imx178_write_register (0x3221, 0x00);
    imx178_write_register (0x3222, 0x74);
    imx178_write_register (0x3223, 0x00);
    imx178_write_register (0x3226, 0xC2);
    imx178_write_register (0x3227, 0x00);
    imx178_write_register (0x32A9, 0x1B);
    imx178_write_register (0x32AA, 0x00);
    imx178_write_register (0x32B3, 0x0E);
    imx178_write_register (0x32B4, 0x00);
    imx178_write_register (0x33D6, 0x16);
    imx178_write_register (0x33D7, 0x15);
    imx178_write_register (0x33D8, 0x14);
    imx178_write_register (0x33D9, 0x10);
    imx178_write_register (0x33DA, 0x08);

    /* registers must be changed */
    imx178_write_register (0x3011, 0x00);
    imx178_write_register (0x301B, 0x00);
    imx178_write_register (0x3037, 0x08);
    imx178_write_register (0x3038, 0x00);
    imx178_write_register (0x3039, 0x00);
    imx178_write_register (0x30AD, 0x49);
    imx178_write_register (0x30AF, 0x54);
    imx178_write_register (0x30B0, 0x33);
    imx178_write_register (0x30B3, 0x0A);
    imx178_write_register (0x30C4, 0x30);
    imx178_write_register (0x3103, 0x03);
    imx178_write_register (0x3104, 0x08);
    imx178_write_register (0x3107, 0x10);
    imx178_write_register (0x310F, 0x01);
    imx178_write_register (0x32E5, 0x06);
    imx178_write_register (0x32E6, 0x00);
    imx178_write_register (0x32E7, 0x1F);
    imx178_write_register (0x32E8, 0x00);
    imx178_write_register (0x32E9, 0x00);
    imx178_write_register (0x32EA, 0x00);
    imx178_write_register (0x32EB, 0x00);
    imx178_write_register (0x32EC, 0x00);
    imx178_write_register (0x32EE, 0x00);
    imx178_write_register (0x32F2, 0x02);
    imx178_write_register (0x32F4, 0x00);
    imx178_write_register (0x32F5, 0x00);
    imx178_write_register (0x32F6, 0x00);
    imx178_write_register (0x32F7, 0x00);
    imx178_write_register (0x32F8, 0x00);
    imx178_write_register (0x32FC, 0x02);
    imx178_write_register (0x3310, 0x11);
    imx178_write_register (0x3338, 0x81);
    imx178_write_register (0x333D, 0x00);
    imx178_write_register (0x3362, 0x00);
    imx178_write_register (0x336B, 0x02);
    imx178_write_register (0x336E, 0x11);
    imx178_write_register (0x33B4, 0xFE);
    imx178_write_register (0x33B5, 0x06);
    imx178_write_register (0x33B9, 0x00);

    /*shutter and gain */
    imx178_write_register (0x3034, 0x08);
    imx178_write_register (0x3035, 0x00); /* SHS1 */
    imx178_write_register (0x301F, 0xA0);
    imx178_write_register (0x3020, 0x00); /* GAIN */

    imx178_write_register (0x3000, 0x00); /* standby */
    imx178_write_register (0x3008, 0x00); /* master mode start */
    imx178_write_register (0x305E, 0x0A); /* XVSOUTSEL XHSOUTSEL */
    imx178_write_register (0x3015, 0xC8); /* BLKLEVEL */

    printf("-------Sony IMX178 Sensor 1080p60fps Initial OK!-------\n");

}
