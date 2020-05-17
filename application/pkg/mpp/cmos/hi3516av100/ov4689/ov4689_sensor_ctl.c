/******************************************************************************

  Copyright (C), 2001-2013, Hisilicon Tech. Co., Ltd.

 ******************************************************************************
  File Name     : ov4689_sensor_ctl.c
  Version       : Initial Draft
  Author        : Hisilicon BVT ISP group
  Created       : 2014/02/20
  Description   : Omnivision OV4689 sensor driver
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

const unsigned char ov4689_i2c_addr     =    0x42; /* I2C Address of ov4689 */
const unsigned int  ov4689_addr_byte    =    2;
const unsigned int  ov4689_data_byte    =    1;
static int g_fd = -1;

extern WDR_MODE_E genSensorMode_ov4689;
extern HI_U8 gu8SensorImageMode_ov4689;
extern HI_BOOL bSensorInit_ov4689;

//int sensor_i2c_init(void)
int ov4689_i2c_init(void)
{
    if(g_fd >= 0)
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

    ret = ioctl(g_fd, I2C_SLAVE_FORCE, ov4689_i2c_addr);
    if (ret < 0)
    {
        printf("CMD_SET_DEV error!\n");
        return ret;
    }
#endif

    return 0;
}

//int sensor_i2c_exit(void)
int ov4689_i2c_exit(void)
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
int ov4689_read_register(int addr)
{
    // TODO:

    return 0;
}

//int ov4689_write_register(int addr, int data)
int ov4689_write_register(int addr, int data)
{
#ifdef HI_GPIO_I2C
    i2c_data.dev_addr = ov4689_i2c_addr;
    i2c_data.reg_addr = addr;
    i2c_data.addr_byte_num = ov4689_addr_byte;
    i2c_data.data = data;
    i2c_data.data_byte_num = ov4689_data_byte;

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
    if (ov4689_addr_byte == 2)
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
    if (ov4689_data_byte == 2)
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
void ov4689_prog(int* rom)
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
            ov4689_write_register(addr, data);
        }
    }
}


static void ov4689_init_4M50();
static void ov4689_init_4M30();
static void ov4689_wdr_4M_25_init();
static void ov4689_wdr_1080p30_init();
static void ov4689_wdr_2304_1296_30_init();
static void ov4689_wdr_2048_1520_30_init();
static void ov4689_linear_2304_1296_30_init();
static void ov4689_linear_2048_1520_30_init();
static void ov4689_linear_1080p60_init();
static void ov4689_linear_720p_180fps_init();

//void sensor_init()
void ov4689_init()
{
    ov4689_i2c_init();

    switch (gu8SensorImageMode_ov4689)
    {
        case 0: // 4M30
            if(WDR_MODE_2To1_LINE == genSensorMode_ov4689)
            {
                ov4689_wdr_4M_25_init();
            }
            else if(WDR_MODE_NONE == genSensorMode_ov4689)
            {
                ov4689_init_4M30();
            }else if((WDR_MODE_2To1_FRAME_FULL_RATE == genSensorMode_ov4689) ||WDR_MODE_2To1_FRAME == genSensorMode_ov4689)
            {
                ov4689_init_4M30();
            }
        break;

        case 1: // 1080P60
            if(WDR_MODE_2To1_LINE == genSensorMode_ov4689)
            {
                ov4689_wdr_1080p30_init();
            }
            else
            {
                ov4689_linear_1080p60_init();
            }
            break;

        case 2: // 2304*1296 30
            if(WDR_MODE_2To1_LINE == genSensorMode_ov4689)
            {
                ov4689_wdr_2304_1296_30_init();
            }
            else
            {
                ov4689_linear_2304_1296_30_init();
            }
            break;
        case 3: // 2048*1520 30
            if(WDR_MODE_2To1_LINE == genSensorMode_ov4689)
            {
                ov4689_wdr_2048_1520_30_init();
            }
            else
            {
                ov4689_linear_2048_1520_30_init();
            }
            break;

        case 4: // 720 180
            if(WDR_MODE_NONE == genSensorMode_ov4689)
            {
                ov4689_linear_720p_180fps_init();
            }
            else
            {
                printf("Not support this mode yet!\n");
            }
        break;


        default:
            printf("Not support this mode\n");
    }

}

//void sensor_exit()
void ov4689_exit()
{
    ov4689_i2c_exit();

    return;
}

//void sensor_enable_wdr_mode()
static void ov4689_enable_wdr_mode()
{
    ov4689_write_register(0x0100,0x01);
    hi_usleep(2000);

    ov4689_write_register(0x3841,0x03);
    ov4689_write_register(0x3846,0x08);
    ov4689_write_register(0x3847,0x07);
    ov4689_write_register(0x4800,0x0C);
    ov4689_write_register(0x376e,0x01);
    ov4689_write_register(0x350B,0x08);
    ov4689_write_register(0x3511,0x01);
    ov4689_write_register(0x3517,0x00);
    ov4689_write_register(0x351d,0x00);

    ov4689_write_register(0x3841,0x03);
    ov4689_write_register(0x3847,0x06);
}

static void ov4689_init_4M50()
{
    delay_ms(10);
    ov4689_write_register(0x0103,0x01);
    ov4689_write_register(0x3638,0x00);
    ov4689_write_register(0x0300,0x00);
    ov4689_write_register(0x0302, 0x19);// 2a
    ov4689_write_register(0x0303,0x00);
    ov4689_write_register(0x0304,0x03);
    ov4689_write_register(0x030b,0x00);
    ov4689_write_register(0x030d, 0x1e);
    ov4689_write_register(0x030e,0x04);
    ov4689_write_register(0x030f,0x01);
    ov4689_write_register(0x0312,0x01);
    ov4689_write_register(0x031e,0x00);
    ov4689_write_register(0x3000,0x20);
    ov4689_write_register(0x3002,0x00);
    ov4689_write_register(0x3018,0x72);
    ov4689_write_register(0x3020,0x93);
    ov4689_write_register(0x3021,0x03);
    ov4689_write_register(0x3022,0x01);
    ov4689_write_register(0x3031,0x0a);
    ov4689_write_register(0x303f,0x0c);
    ov4689_write_register(0x3305,0xf1);
    ov4689_write_register(0x3307,0x04);
    ov4689_write_register(0x3309,0x29);
    ov4689_write_register(0x3500,0x00);
    ov4689_write_register(0x3501,0x60);
    ov4689_write_register(0x3502,0x00);
    ov4689_write_register(0x3503,0x77);
    ov4689_write_register(0x3504,0x00);
    ov4689_write_register(0x3505,0x00);
    ov4689_write_register(0x3506,0x00);
    ov4689_write_register(0x3507,0x00);
    ov4689_write_register(0x3508,0x00);
    ov4689_write_register(0x3509,0x80);
    ov4689_write_register(0x350a,0x00);
    ov4689_write_register(0x350b,0x00);
    ov4689_write_register(0x350c,0x00);
    ov4689_write_register(0x350d,0x00);
    ov4689_write_register(0x350e,0x00);
    ov4689_write_register(0x350f,0x80);
    ov4689_write_register(0x3510,0x00);
    ov4689_write_register(0x3511,0x00);
    ov4689_write_register(0x3512,0x00);
    ov4689_write_register(0x3513,0x00);
    ov4689_write_register(0x3514,0x00);
    ov4689_write_register(0x3515,0x80);
    ov4689_write_register(0x3516,0x00);
    ov4689_write_register(0x3517,0x00);
    ov4689_write_register(0x3518,0x00);
    ov4689_write_register(0x3519,0x00);
    ov4689_write_register(0x351a,0x00);
    ov4689_write_register(0x351b,0x80);
    ov4689_write_register(0x351c,0x00);
    ov4689_write_register(0x351d,0x00);
    ov4689_write_register(0x351e,0x00);
    ov4689_write_register(0x351f,0x00);
    ov4689_write_register(0x3520,0x00);
    ov4689_write_register(0x3521,0x80);
    ov4689_write_register(0x3522,0x08);
    ov4689_write_register(0x3524,0x08);
    ov4689_write_register(0x3526,0x08);
    ov4689_write_register(0x3528,0x08);
    ov4689_write_register(0x352a,0x08);
    ov4689_write_register(0x3602,0x00);
    ov4689_write_register(0x3603,0x40);
    ov4689_write_register(0x3604,0x02);
    ov4689_write_register(0x3605,0x00);
    ov4689_write_register(0x3606,0x00);
    ov4689_write_register(0x3607,0x00);
    ov4689_write_register(0x3609,0x12);
    ov4689_write_register(0x360a,0x40);
    ov4689_write_register(0x360c,0x08);
    ov4689_write_register(0x360f,0xe5);
    ov4689_write_register(0x3608,0x8f);
    ov4689_write_register(0x3611,0x00);
    ov4689_write_register(0x3613,0xf7);
    ov4689_write_register(0x3616,0x58);
    ov4689_write_register(0x3619,0x99);
    ov4689_write_register(0x361b,0x60);
    ov4689_write_register(0x361c,0x7a);
    ov4689_write_register(0x361e,0x79);
    ov4689_write_register(0x361f,0x02);
    ov4689_write_register(0x3632,0x00);
    ov4689_write_register(0x3633,0x10);
    ov4689_write_register(0x3634,0x10);
    ov4689_write_register(0x3635,0x10);
    ov4689_write_register(0x3636,0x15);
    ov4689_write_register(0x3646,0x86);
    ov4689_write_register(0x364a,0x0b);
    ov4689_write_register(0x3700,0x17);
    ov4689_write_register(0x3701,0x22);
    ov4689_write_register(0x3703,0x10);
    ov4689_write_register(0x370a,0x37);
    ov4689_write_register(0x3705,0x00);
    ov4689_write_register(0x3706,0x63);
    ov4689_write_register(0x3709,0x3c);
    ov4689_write_register(0x370b,0x01);
    ov4689_write_register(0x370c,0x30);
    ov4689_write_register(0x3710,0x24);
    ov4689_write_register(0x3711,0x0c);
    ov4689_write_register(0x3716,0x00);
    ov4689_write_register(0x3720,0x28);
    ov4689_write_register(0x3729,0x7b);
    ov4689_write_register(0x372a,0x84);
    ov4689_write_register(0x372b,0xbd);
    ov4689_write_register(0x372c,0xbc);
    ov4689_write_register(0x372e,0x52);
    ov4689_write_register(0x373c,0x0e);
    ov4689_write_register(0x373e,0x33);
    ov4689_write_register(0x3743,0x10);
    ov4689_write_register(0x3744,0x88);
    ov4689_write_register(0x3745,0xc0);
    ov4689_write_register(0x374a,0x43);
    ov4689_write_register(0x374c,0x00);
    ov4689_write_register(0x374e,0x23);
    ov4689_write_register(0x3751,0x7b);
    ov4689_write_register(0x3752,0x84);
    ov4689_write_register(0x3753,0xbd);
    ov4689_write_register(0x3754,0xbc);
    ov4689_write_register(0x3756,0x52);
    ov4689_write_register(0x375c,0x00);
    ov4689_write_register(0x3760,0x00);
    ov4689_write_register(0x3761,0x00);
    ov4689_write_register(0x3762,0x00);
    ov4689_write_register(0x3763,0x00);
    ov4689_write_register(0x3764,0x00);
    ov4689_write_register(0x3767,0x04);
    ov4689_write_register(0x3768,0x04);
    ov4689_write_register(0x3769,0x08);
    ov4689_write_register(0x376a,0x08);
    ov4689_write_register(0x376b,0x20);
    ov4689_write_register(0x376c,0x00);
    ov4689_write_register(0x376d,0x00);
    ov4689_write_register(0x376e,0x00);
    ov4689_write_register(0x3773,0x00);
    ov4689_write_register(0x3774,0x51);
    ov4689_write_register(0x3776,0xbd);
    ov4689_write_register(0x3777,0xbd);
    ov4689_write_register(0x3781,0x18);
    ov4689_write_register(0x3783,0x25);
    ov4689_write_register(0x3798,0x1b);
    ov4689_write_register(0x3800,0x00);
    ov4689_write_register(0x3801,0x38);
    ov4689_write_register(0x3802,0x00);
    ov4689_write_register(0x3803,0x04);
    ov4689_write_register(0x3804,0x0A);
    ov4689_write_register(0x3805,0x67);
    ov4689_write_register(0x3806,0x05);
    ov4689_write_register(0x3807,0xFB);
    ov4689_write_register(0x3808,0x0A);
    ov4689_write_register(0x3809,0x20);
    ov4689_write_register(0x380A,0x05);
    ov4689_write_register(0x380B,0xF0);
    ov4689_write_register(0x380c,0x05);//03
    ov4689_write_register(0x380d,0xc0);//5c
    ov4689_write_register(0x380E,0x06);
    ov4689_write_register(0x380F,0x5F);
    ov4689_write_register(0x3810,0x00);
    ov4689_write_register(0x3811,0x08);
    ov4689_write_register(0x3812,0x00);
    ov4689_write_register(0x3813,0x04);
    ov4689_write_register(0x3814,0x01);
    ov4689_write_register(0x3815,0x01);
    ov4689_write_register(0x3819,0x01);
    ov4689_write_register(0x3820,0x00);
    ov4689_write_register(0x3821,0x06);
    ov4689_write_register(0x3829,0x00);
    ov4689_write_register(0x382a,0x01);
    ov4689_write_register(0x382b,0x01);
    ov4689_write_register(0x382d,0x7f);
    ov4689_write_register(0x3830,0x04);
    ov4689_write_register(0x3836,0x01);
    ov4689_write_register(0x3837,0x00);
    ov4689_write_register(0x3841,0x02);
    ov4689_write_register(0x3846,0x08);
    ov4689_write_register(0x3847,0x07);
    ov4689_write_register(0x3d85,0x36);
    ov4689_write_register(0x3d8c,0x71);
    ov4689_write_register(0x3d8d,0xcb);
    ov4689_write_register(0x3f0a,0x00);
    ov4689_write_register(0x4000,0xf1);
    ov4689_write_register(0x4001,0x40);
    ov4689_write_register(0x4002,0x04);
    ov4689_write_register(0x4003,0x14);
    ov4689_write_register(0x400e,0x00);
    ov4689_write_register(0x4011,0x00);
    ov4689_write_register(0x401a,0x00);
    ov4689_write_register(0x401b,0x00);
    ov4689_write_register(0x401c,0x00);
    ov4689_write_register(0x401d,0x00);
    ov4689_write_register(0x401f,0x00);
    ov4689_write_register(0x4020,0x00);
    ov4689_write_register(0x4021,0x10);
    ov4689_write_register(0x4022,0x08);
    ov4689_write_register(0x4023,0xB3);
    ov4689_write_register(0x4024,0x09);
    ov4689_write_register(0x4025,0xE0);
    ov4689_write_register(0x4026,0x09);
    ov4689_write_register(0x4027,0xF0);
    ov4689_write_register(0x4028,0x00);
    ov4689_write_register(0x4029,0x02);
    ov4689_write_register(0x402a,0x02);
    ov4689_write_register(0x402b,0x02);
    ov4689_write_register(0x402c,0x02);
    ov4689_write_register(0x402d,0x02);
    ov4689_write_register(0x402e,0x0e);
    ov4689_write_register(0x402f,0x04);
    ov4689_write_register(0x4302,0xff);
    ov4689_write_register(0x4303,0xff);
    ov4689_write_register(0x4304,0x00);
    ov4689_write_register(0x4305,0x00);
    ov4689_write_register(0x4306,0x00);
    ov4689_write_register(0x4308,0x02);
    ov4689_write_register(0x4500,0x6c);
    ov4689_write_register(0x4501,0xc4);
    ov4689_write_register(0x4502,0x40);
    ov4689_write_register(0x4503,0x01);
    ov4689_write_register(0x4601,0x20);
    ov4689_write_register(0x4800,0x04);
    ov4689_write_register(0x4813,0x08);
    ov4689_write_register(0x481f,0x40);
    ov4689_write_register(0x4829,0x78);
    ov4689_write_register(0x4837,0x1a);
    ov4689_write_register(0x4b00,0x2a);
    ov4689_write_register(0x4b0d,0x00);
    ov4689_write_register(0x4d00,0x04);
    ov4689_write_register(0x4d01,0x42);
    ov4689_write_register(0x4d02,0xd1);
    ov4689_write_register(0x4d03,0x93);
    ov4689_write_register(0x4d04,0xf5);
    ov4689_write_register(0x4d05,0xc1);
    ov4689_write_register(0x5000,0xf3);
    ov4689_write_register(0x5001,0x11);
    ov4689_write_register(0x5004,0x00);
    ov4689_write_register(0x500a,0x00);
    ov4689_write_register(0x500b,0x00);
    ov4689_write_register(0x5032,0x00);
    ov4689_write_register(0x5040,0x00);
    ov4689_write_register(0x5050,0x0c);
    ov4689_write_register(0x5500,0x00);
    ov4689_write_register(0x5501,0x10);
    ov4689_write_register(0x5502,0x01);
    ov4689_write_register(0x5503,0x0f);
    ov4689_write_register(0x8000,0x00);
    ov4689_write_register(0x8001,0x00);
    ov4689_write_register(0x8002,0x00);
    ov4689_write_register(0x8003,0x00);
    ov4689_write_register(0x8004,0x00);
    ov4689_write_register(0x8005,0x00);
    ov4689_write_register(0x8006,0x00);
    ov4689_write_register(0x8007,0x00);
    ov4689_write_register(0x8008,0x00);
    ov4689_write_register(0x3638,0x00);

    bSensorInit_ov4689 = HI_TRUE;
}

static void ov4689_init_4M30()
{
    delay_ms(10);
    ov4689_write_register(0x0103,0x01);

    ov4689_write_register(0x3638,0x00);
    ov4689_write_register(0x0300,0x00);
    ov4689_write_register(0x0302,0x19);//;2a
    ov4689_write_register(0x0303,0x00);
    ov4689_write_register(0x0304,0x03);
    ov4689_write_register(0x030b,0x00);
    ov4689_write_register(0x030d,0x1e);
    ov4689_write_register(0x030e,0x04);
    ov4689_write_register(0x030f,0x01);
    ov4689_write_register(0x0312,0x01);
    ov4689_write_register(0x031e,0x00);
    ov4689_write_register(0x3000,0x20);
    ov4689_write_register(0x3002,0x00);
    ov4689_write_register(0x3018,0x72);
    ov4689_write_register(0x3020,0x93);
    ov4689_write_register(0x3021,0x03);
    ov4689_write_register(0x3022,0x01);
    ov4689_write_register(0x3031,0x0a);
    ov4689_write_register(0x303f,0x0c);
    ov4689_write_register(0x3305,0xf1);
    ov4689_write_register(0x3307,0x04);
    ov4689_write_register(0x3309,0x29);
    ov4689_write_register(0x3500,0x00);
    ov4689_write_register(0x3501,0x60);
    ov4689_write_register(0x3502,0x00);
    ov4689_write_register(0x3503,0x77);
    ov4689_write_register(0x3504,0x00);
    ov4689_write_register(0x3505,0x00);
    ov4689_write_register(0x3506,0x00);
    ov4689_write_register(0x3507,0x00);
    ov4689_write_register(0x3508,0x00);
    ov4689_write_register(0x3509,0x80);
    ov4689_write_register(0x350a,0x00);
    ov4689_write_register(0x350b,0x00);
    ov4689_write_register(0x350c,0x00);
    ov4689_write_register(0x350d,0x00);
    ov4689_write_register(0x350e,0x00);
    ov4689_write_register(0x350f,0x80);
    ov4689_write_register(0x3510,0x00);
    ov4689_write_register(0x3511,0x00);
    ov4689_write_register(0x3512,0x00);
    ov4689_write_register(0x3513,0x00);
    ov4689_write_register(0x3514,0x00);
    ov4689_write_register(0x3515,0x80);
    ov4689_write_register(0x3516,0x00);
    ov4689_write_register(0x3517,0x00);
    ov4689_write_register(0x3518,0x00);
    ov4689_write_register(0x3519,0x00);
    ov4689_write_register(0x351a,0x00);
    ov4689_write_register(0x351b,0x80);
    ov4689_write_register(0x351c,0x00);
    ov4689_write_register(0x351d,0x00);
    ov4689_write_register(0x351e,0x00);
    ov4689_write_register(0x351f,0x00);
    ov4689_write_register(0x3520,0x00);
    ov4689_write_register(0x3521,0x80);
    ov4689_write_register(0x3522,0x08);
    ov4689_write_register(0x3524,0x08);
    ov4689_write_register(0x3526,0x08);
    ov4689_write_register(0x3528,0x08);
    ov4689_write_register(0x352a,0x08);
    ov4689_write_register(0x3602,0x00);
    ov4689_write_register(0x3603,0x40);
    ov4689_write_register(0x3604,0x02);
    ov4689_write_register(0x3605,0x00);
    ov4689_write_register(0x3606,0x00);
    ov4689_write_register(0x3607,0x00);
    ov4689_write_register(0x3609,0x12);
    ov4689_write_register(0x360a,0x40);
    ov4689_write_register(0x360c,0x08);
    ov4689_write_register(0x360f,0xe5);
    ov4689_write_register(0x3608,0x8f);
    ov4689_write_register(0x3611,0x00);
    ov4689_write_register(0x3613,0xf7);
    ov4689_write_register(0x3616,0x58);
    ov4689_write_register(0x3619,0x99);
    ov4689_write_register(0x361b,0x60);
    ov4689_write_register(0x361c,0x7a);
    ov4689_write_register(0x361e,0x79);
    ov4689_write_register(0x361f,0x02);
    ov4689_write_register(0x3632,0x00);
    ov4689_write_register(0x3633,0x10);
    ov4689_write_register(0x3634,0x10);
    ov4689_write_register(0x3635,0x10);
    ov4689_write_register(0x3636,0x15);
    ov4689_write_register(0x3646,0x86);
    ov4689_write_register(0x364a,0x0b);
    ov4689_write_register(0x3700,0x17);
    ov4689_write_register(0x3701,0x22);
    ov4689_write_register(0x3703,0x10);
    ov4689_write_register(0x370a,0x37);
    ov4689_write_register(0x3705,0x00);
    ov4689_write_register(0x3706,0x63);
    ov4689_write_register(0x3709,0x3c);
    ov4689_write_register(0x370b,0x01);
    ov4689_write_register(0x370c,0x30);
    ov4689_write_register(0x3710,0x24);
    ov4689_write_register(0x3711,0x0c);
    ov4689_write_register(0x3716,0x00);
    ov4689_write_register(0x3720,0x28);
    ov4689_write_register(0x3729,0x7b);
    ov4689_write_register(0x372a,0x84);
    ov4689_write_register(0x372b,0xbd);
    ov4689_write_register(0x372c,0xbc);
    ov4689_write_register(0x372e,0x52);
    ov4689_write_register(0x373c,0x0e);
    ov4689_write_register(0x373e,0x33);
    ov4689_write_register(0x3743,0x10);
    ov4689_write_register(0x3744,0x88);
    ov4689_write_register(0x3745,0xc0);
    ov4689_write_register(0x374a,0x43);
    ov4689_write_register(0x374c,0x00);
    ov4689_write_register(0x374e,0x23);
    ov4689_write_register(0x3751,0x7b);
    ov4689_write_register(0x3752,0x84);
    ov4689_write_register(0x3753,0xbd);
    ov4689_write_register(0x3754,0xbc);
    ov4689_write_register(0x3756,0x52);
    ov4689_write_register(0x375c,0x00);
    ov4689_write_register(0x3760,0x00);
    ov4689_write_register(0x3761,0x00);
    ov4689_write_register(0x3762,0x00);
    ov4689_write_register(0x3763,0x00);
    ov4689_write_register(0x3764,0x00);
    ov4689_write_register(0x3767,0x04);
    ov4689_write_register(0x3768,0x04);
    ov4689_write_register(0x3769,0x08);
    ov4689_write_register(0x376a,0x08);
    ov4689_write_register(0x376b,0x20);
    ov4689_write_register(0x376c,0x00);
    ov4689_write_register(0x376d,0x00);
    ov4689_write_register(0x376e,0x00);
    ov4689_write_register(0x3773,0x00);
    ov4689_write_register(0x3774,0x51);
    ov4689_write_register(0x3776,0xbd);
    ov4689_write_register(0x3777,0xbd);
    ov4689_write_register(0x3781,0x18);
    ov4689_write_register(0x3783,0x25);
    ov4689_write_register(0x3798,0x1b);
    ov4689_write_register(0x3800,0x00);
    ov4689_write_register(0x3801,0x38);
    ov4689_write_register(0x3802,0x00);
    ov4689_write_register(0x3803,0x04);
    ov4689_write_register(0x3804,0x0A);
    ov4689_write_register(0x3805,0x67);
    ov4689_write_register(0x3806,0x05);
    ov4689_write_register(0x3807,0xFB);
    ov4689_write_register(0x3808,0x0A);
    ov4689_write_register(0x3809,0x20);
    ov4689_write_register(0x380A,0x05);
    ov4689_write_register(0x380B,0xF0);
    ov4689_write_register(0x380c,0x0A);//;03
    ov4689_write_register(0x380d,0x0B);//;5c
    ov4689_write_register(0x380E,0x06);
    ov4689_write_register(0x380F,0x14);
    ov4689_write_register(0x3810,0x00);
    ov4689_write_register(0x3811,0x08);
    ov4689_write_register(0x3812,0x00);
    ov4689_write_register(0x3813,0x04);
    ov4689_write_register(0x3814,0x01);
    ov4689_write_register(0x3815,0x01);
    ov4689_write_register(0x3819,0x01);
    ov4689_write_register(0x3820,0x00);
    ov4689_write_register(0x3821,0x06);
    ov4689_write_register(0x3829,0x00);
    ov4689_write_register(0x382a,0x01);
    ov4689_write_register(0x382b,0x01);
    ov4689_write_register(0x382d,0x7f);
    ov4689_write_register(0x3830,0x04);
    ov4689_write_register(0x3836,0x01);
    ov4689_write_register(0x3837,0x00);
    ov4689_write_register(0x3841,0x02);
    ov4689_write_register(0x3846,0x08);
    ov4689_write_register(0x3847,0x07);
    ov4689_write_register(0x3d85,0x36);
    ov4689_write_register(0x3d8c,0x71);
    ov4689_write_register(0x3d8d,0xcb);
    ov4689_write_register(0x3f0a,0x00);
    ov4689_write_register(0x4000,0xf1);
    ov4689_write_register(0x4001,0x40);
    ov4689_write_register(0x4002,0x04);
    ov4689_write_register(0x4003,0x14);
    ov4689_write_register(0x400e,0x00);
    ov4689_write_register(0x4011,0x00);
    ov4689_write_register(0x401a,0x00);
    ov4689_write_register(0x401b,0x00);
    ov4689_write_register(0x401c,0x00);
    ov4689_write_register(0x401d,0x00);
    ov4689_write_register(0x401f,0x00);
    ov4689_write_register(0x4020,0x00);
    ov4689_write_register(0x4021,0x10);
    ov4689_write_register(0x4022,0x08);
    ov4689_write_register(0x4023,0xB3);
    ov4689_write_register(0x4024,0x09);
    ov4689_write_register(0x4025,0xE0);
    ov4689_write_register(0x4026,0x09);
    ov4689_write_register(0x4027,0xF0);
    ov4689_write_register(0x4028,0x00);
    ov4689_write_register(0x4029,0x02);
    ov4689_write_register(0x402a,0x02);
    ov4689_write_register(0x402b,0x02);
    ov4689_write_register(0x402c,0x02);
    ov4689_write_register(0x402d,0x02);
    ov4689_write_register(0x402e,0x0e);
    ov4689_write_register(0x402f,0x04);
    ov4689_write_register(0x4302,0xff);
    ov4689_write_register(0x4303,0xff);
    ov4689_write_register(0x4304,0x00);
    ov4689_write_register(0x4305,0x00);
    ov4689_write_register(0x4306,0x00);
    ov4689_write_register(0x4308,0x02);
    ov4689_write_register(0x4500,0x6c);
    ov4689_write_register(0x4501,0xc4);
    ov4689_write_register(0x4502,0x40);
    ov4689_write_register(0x4503,0x01);
    ov4689_write_register(0x4601,0x41);
    ov4689_write_register(0x4800,0x04);
    ov4689_write_register(0x4813,0x08);
    ov4689_write_register(0x481f,0x40);
    ov4689_write_register(0x4829,0x78);
    ov4689_write_register(0x4837,0x1a);//;10
    ov4689_write_register(0x4b00,0x2a);
    ov4689_write_register(0x4b0d,0x00);
    ov4689_write_register(0x4d00,0x04);
    ov4689_write_register(0x4d01,0x42);
    ov4689_write_register(0x4d02,0xd1);
    ov4689_write_register(0x4d03,0x93);
    ov4689_write_register(0x4d04,0xf5);
    ov4689_write_register(0x4d05,0xc1);
    ov4689_write_register(0x5000,0xf3);
    ov4689_write_register(0x5001,0x11);
    ov4689_write_register(0x5004,0x00);
    ov4689_write_register(0x500a,0x00);
    ov4689_write_register(0x500b,0x00);
    ov4689_write_register(0x5032,0x00);
    ov4689_write_register(0x5040,0x00);
    ov4689_write_register(0x5050,0x0c);
    ov4689_write_register(0x5500,0x00);
    ov4689_write_register(0x5501,0x10);
    ov4689_write_register(0x5502,0x01);
    ov4689_write_register(0x5503,0x0f);
    ov4689_write_register(0x8000,0x00);
    ov4689_write_register(0x8001,0x00);
    ov4689_write_register(0x8002,0x00);
    ov4689_write_register(0x8003,0x00);
    ov4689_write_register(0x8004,0x00);
    ov4689_write_register(0x8005,0x00);
    ov4689_write_register(0x8006,0x00);
    ov4689_write_register(0x8007,0x00);
    ov4689_write_register(0x8008,0x00);
    ov4689_write_register(0x3638,0x00);
    ov4689_write_register(0x0100,0x01);

    bSensorInit_ov4689 = HI_TRUE;
    printf("-------OV4689 Sensor 4M30fps Initial OK!!-------\n");
}

void ov4689_wdr_4M_25_init()
{
    ov4689_init_4M50();
    ov4689_enable_wdr_mode();
    bSensorInit_ov4689 = HI_TRUE;
    printf("-------OV4689 Sensor 4M25fps  WDR Initial OK!-------\n");
}

void ov4689_linear_1080p60_init()
{
    delay_ms(10);
    ov4689_write_register(0x0103,0x01);

    ov4689_write_register(0x3638,0x00);
    ov4689_write_register(0x0300,0x00);
    ov4689_write_register(0x0302,0x19);//;2a
    ov4689_write_register(0x0303,0x00);
    ov4689_write_register(0x0304,0x03);
    ov4689_write_register(0x030b,0x00);
    ov4689_write_register(0x030d,0x1e);
    ov4689_write_register(0x030e,0x04);
    ov4689_write_register(0x030f,0x01);
    ov4689_write_register(0x0312,0x01);
    ov4689_write_register(0x031e,0x00);
    ov4689_write_register(0x3000,0x20);
    ov4689_write_register(0x3002,0x00);
    ov4689_write_register(0x3018,0x72);
    ov4689_write_register(0x3020,0x93);
    ov4689_write_register(0x3021,0x03);
    ov4689_write_register(0x3022,0x01);
    ov4689_write_register(0x3031,0x0a);
    ov4689_write_register(0x303f,0x0c);
    ov4689_write_register(0x3305,0xf1);
    ov4689_write_register(0x3307,0x04);
    ov4689_write_register(0x3309,0x29);
    ov4689_write_register(0x3500,0x00);
    ov4689_write_register(0x3501,0x4c);
    ov4689_write_register(0x3502,0x00);
    ov4689_write_register(0x3503,0x77);
    ov4689_write_register(0x3504,0x00);
    ov4689_write_register(0x3505,0x00);
    ov4689_write_register(0x3506,0x00);
    ov4689_write_register(0x3507,0x00);
    ov4689_write_register(0x3508,0x00);
    ov4689_write_register(0x3509,0x80);
    ov4689_write_register(0x350a,0x00);
    ov4689_write_register(0x350b,0x00);
    ov4689_write_register(0x350c,0x00);
    ov4689_write_register(0x350d,0x00);
    ov4689_write_register(0x350e,0x00);
    ov4689_write_register(0x350f,0x80);
    ov4689_write_register(0x3510,0x00);
    ov4689_write_register(0x3511,0x00);
    ov4689_write_register(0x3512,0x00);
    ov4689_write_register(0x3513,0x00);
    ov4689_write_register(0x3514,0x00);
    ov4689_write_register(0x3515,0x80);
    ov4689_write_register(0x3516,0x00);
    ov4689_write_register(0x3517,0x00);
    ov4689_write_register(0x3518,0x00);
    ov4689_write_register(0x3519,0x00);
    ov4689_write_register(0x351a,0x00);
    ov4689_write_register(0x351b,0x80);
    ov4689_write_register(0x351c,0x00);
    ov4689_write_register(0x351d,0x00);
    ov4689_write_register(0x351e,0x00);
    ov4689_write_register(0x351f,0x00);
    ov4689_write_register(0x3520,0x00);
    ov4689_write_register(0x3521,0x80);
    ov4689_write_register(0x3522,0x08);
    ov4689_write_register(0x3524,0x08);
    ov4689_write_register(0x3526,0x08);
    ov4689_write_register(0x3528,0x08);
    ov4689_write_register(0x352a,0x08);
    ov4689_write_register(0x3602,0x00);
    ov4689_write_register(0x3603,0x40);
    ov4689_write_register(0x3604,0x02);
    ov4689_write_register(0x3605,0x00);
    ov4689_write_register(0x3606,0x00);
    ov4689_write_register(0x3607,0x00);
    ov4689_write_register(0x3609,0x12);
    ov4689_write_register(0x360a,0x40);
    ov4689_write_register(0x360c,0x08);
    ov4689_write_register(0x360f,0xe5);
    ov4689_write_register(0x3608,0x8f);
    ov4689_write_register(0x3611,0x00);
    ov4689_write_register(0x3613,0xf7);
    ov4689_write_register(0x3616,0x58);
    ov4689_write_register(0x3619,0x99);
    ov4689_write_register(0x361b,0x60);
    ov4689_write_register(0x361c,0x7a);
    ov4689_write_register(0x361e,0x79);
    ov4689_write_register(0x361f,0x02);
    ov4689_write_register(0x3632,0x00);
    ov4689_write_register(0x3633,0x10);
    ov4689_write_register(0x3634,0x10);
    ov4689_write_register(0x3635,0x10);
    ov4689_write_register(0x3636,0x15);
    ov4689_write_register(0x3646,0x86);
    ov4689_write_register(0x364a,0x0b);
    ov4689_write_register(0x3700,0x17);
    ov4689_write_register(0x3701,0x22);
    ov4689_write_register(0x3703,0x10);
    ov4689_write_register(0x370a,0x37);
    ov4689_write_register(0x3705,0x00);
    ov4689_write_register(0x3706,0x63);
    ov4689_write_register(0x3709,0x3c);
    ov4689_write_register(0x370b,0x01);
    ov4689_write_register(0x370c,0x30);
    ov4689_write_register(0x3710,0x24);
    ov4689_write_register(0x3711,0x0c);
    ov4689_write_register(0x3716,0x00);
    ov4689_write_register(0x3720,0x28);
    ov4689_write_register(0x3729,0x7b);
    ov4689_write_register(0x372a,0x84);
    ov4689_write_register(0x372b,0xbd);
    ov4689_write_register(0x372c,0xbc);
    ov4689_write_register(0x372e,0x52);
    ov4689_write_register(0x373c,0x0e);
    ov4689_write_register(0x373e,0x33);
    ov4689_write_register(0x3743,0x10);
    ov4689_write_register(0x3744,0x88);
    ov4689_write_register(0x3745,0xc0);
    ov4689_write_register(0x374a,0x43);
    ov4689_write_register(0x374c,0x00);
    ov4689_write_register(0x374e,0x23);
    ov4689_write_register(0x3751,0x7b);
    ov4689_write_register(0x3752,0x84);
    ov4689_write_register(0x3753,0xbd);
    ov4689_write_register(0x3754,0xbc);
    ov4689_write_register(0x3756,0x52);
    ov4689_write_register(0x375c,0x00);
    ov4689_write_register(0x3760,0x00);
    ov4689_write_register(0x3761,0x00);
    ov4689_write_register(0x3762,0x00);
    ov4689_write_register(0x3763,0x00);
    ov4689_write_register(0x3764,0x00);
    ov4689_write_register(0x3767,0x04);
    ov4689_write_register(0x3768,0x04);
    ov4689_write_register(0x3769,0x08);
    ov4689_write_register(0x376a,0x08);
    ov4689_write_register(0x376b,0x20);
    ov4689_write_register(0x376c,0x00);
    ov4689_write_register(0x376d,0x00);
    ov4689_write_register(0x376e,0x00);
    ov4689_write_register(0x3773,0x00);
    ov4689_write_register(0x3774,0x51);
    ov4689_write_register(0x3776,0xbd);
    ov4689_write_register(0x3777,0xbd);
    ov4689_write_register(0x3781,0x18);
    ov4689_write_register(0x3783,0x25);
    ov4689_write_register(0x3798,0x1b);
    ov4689_write_register(0x3800,0x00);
    ov4689_write_register(0x3801,0x86);
    ov4689_write_register(0x3802,0x00);
    ov4689_write_register(0x3803,0xe0);
    ov4689_write_register(0x3804,0x09);
    ov4689_write_register(0x3805,0x17);
    ov4689_write_register(0x3806,0x05);
    ov4689_write_register(0x3807,0x1f);
    ov4689_write_register(0x3808,0x07);
    ov4689_write_register(0x3809,0x80);
    ov4689_write_register(0x380a,0x04);
    ov4689_write_register(0x380b,0x38);
    ov4689_write_register(0x380c,0x05);//;03
    ov4689_write_register(0x380d,0x07);//;5c
    ov4689_write_register(0x380e,0x06);//;04
    ov4689_write_register(0x380f,0x12);//;8A
    ov4689_write_register(0x3810,0x00);
    ov4689_write_register(0x3811,0x08);
    ov4689_write_register(0x3812,0x00);
    ov4689_write_register(0x3813,0x04);
    ov4689_write_register(0x3814,0x01);
    ov4689_write_register(0x3815,0x01);
    ov4689_write_register(0x3819,0x01);
    ov4689_write_register(0x3820,0x00);
    ov4689_write_register(0x3821,0x06);
    ov4689_write_register(0x3829,0x00);
    ov4689_write_register(0x382a,0x01);
    ov4689_write_register(0x382b,0x01);
    ov4689_write_register(0x382d,0x7f);
    ov4689_write_register(0x3830,0x04);
    ov4689_write_register(0x3836,0x01);
    ov4689_write_register(0x3837,0x00);
    ov4689_write_register(0x3841,0x02);
    ov4689_write_register(0x3846,0x08);
    ov4689_write_register(0x3847,0x07);
    ov4689_write_register(0x3d85,0x36);
    ov4689_write_register(0x3d8c,0x71);
    ov4689_write_register(0x3d8d,0xcb);
    ov4689_write_register(0x3f0a,0x00);
    ov4689_write_register(0x4000,0xF1);
    ov4689_write_register(0x4001,0x40);
    ov4689_write_register(0x4002,0x04);
    ov4689_write_register(0x4003,0x14);
    ov4689_write_register(0x400e,0x00);
    ov4689_write_register(0x4011,0x00);
    ov4689_write_register(0x401a,0x00);
    ov4689_write_register(0x401b,0x00);
    ov4689_write_register(0x401c,0x00);
    ov4689_write_register(0x401d,0x00);
    ov4689_write_register(0x401f,0x00);
    ov4689_write_register(0x4020,0x00);
    ov4689_write_register(0x4021,0x10);
    ov4689_write_register(0x4022,0x06);
    ov4689_write_register(0x4023,0x13);
    ov4689_write_register(0x4024,0x07);
    ov4689_write_register(0x4025,0x40);
    ov4689_write_register(0x4026,0x07);
    ov4689_write_register(0x4027,0x50);
    ov4689_write_register(0x4028,0x00);
    ov4689_write_register(0x4029,0x02);
    ov4689_write_register(0x402a,0x02);
    ov4689_write_register(0x402b,0x02);
    ov4689_write_register(0x402c,0x02);
    ov4689_write_register(0x402d,0x02);
    ov4689_write_register(0x402e,0x0e);
    ov4689_write_register(0x402f,0x04);
    ov4689_write_register(0x4302,0xff);
    ov4689_write_register(0x4303,0xff);
    ov4689_write_register(0x4304,0x00);
    ov4689_write_register(0x4305,0x00);
    ov4689_write_register(0x4306,0x00);
    ov4689_write_register(0x4308,0x02);
    ov4689_write_register(0x4500,0x6c);
    ov4689_write_register(0x4501,0xc4);
    ov4689_write_register(0x4502,0x40);
    ov4689_write_register(0x4503,0x01);
    ov4689_write_register(0x4601,0x41);
    ov4689_write_register(0x4800,0x04);
    ov4689_write_register(0x4813,0x08);
    ov4689_write_register(0x481f,0x40);
    ov4689_write_register(0x4829,0x78);
    ov4689_write_register(0x4837,0x1a);//;10
    ov4689_write_register(0x4b00,0x2a);
    ov4689_write_register(0x4b0d,0x00);
    ov4689_write_register(0x4d00,0x04);
    ov4689_write_register(0x4d01,0x42);
    ov4689_write_register(0x4d02,0xd1);
    ov4689_write_register(0x4d03,0x93);
    ov4689_write_register(0x4d04,0xf5);
    ov4689_write_register(0x4d05,0xc1);
    ov4689_write_register(0x5000,0xf3);
    ov4689_write_register(0x5001,0x11);
    ov4689_write_register(0x5004,0x00);
    ov4689_write_register(0x500a,0x00);
    ov4689_write_register(0x500b,0x00);
    ov4689_write_register(0x5032,0x00);
    ov4689_write_register(0x5040,0x00);
    ov4689_write_register(0x5050,0x0c);
    ov4689_write_register(0x5500,0x00);
    ov4689_write_register(0x5501,0x10);
    ov4689_write_register(0x5502,0x01);
    ov4689_write_register(0x5503,0x0f);
    ov4689_write_register(0x8000,0x00);
    ov4689_write_register(0x8001,0x00);
    ov4689_write_register(0x8002,0x00);
    ov4689_write_register(0x8003,0x00);
    ov4689_write_register(0x8004,0x00);
    ov4689_write_register(0x8005,0x00);
    ov4689_write_register(0x8006,0x00);
    ov4689_write_register(0x8007,0x00);
    ov4689_write_register(0x8008,0x00);
    ov4689_write_register(0x3638,0x00);
    ov4689_write_register(0x0100,0x01);
    bSensorInit_ov4689 = HI_TRUE;
    printf("-------OV4689 Sensor 1080P 60fps Linear Mode Initial OK!-------\n");
}

void ov4689_wdr_1080p30_init()
{
    delay_ms(10);
    ov4689_write_register(0x0103,0x01);

    ov4689_write_register(0x3638,0x00);
    ov4689_write_register(0x0300,0x00);
    ov4689_write_register(0x0302,0x19);//;2a
    ov4689_write_register(0x0303,0x00);
    ov4689_write_register(0x0304,0x03);
    ov4689_write_register(0x030b,0x00);
    ov4689_write_register(0x030d,0x1e);
    ov4689_write_register(0x030e,0x04);
    ov4689_write_register(0x030f,0x01);
    ov4689_write_register(0x0312,0x01);
    ov4689_write_register(0x031e,0x00);
    ov4689_write_register(0x3000,0x20);
    ov4689_write_register(0x3002,0x00);
    ov4689_write_register(0x3018,0x72);
    ov4689_write_register(0x3020,0x93);
    ov4689_write_register(0x3021,0x03);
    ov4689_write_register(0x3022,0x01);
    ov4689_write_register(0x3031,0x0a);
    ov4689_write_register(0x303f,0x0c);
    ov4689_write_register(0x3305,0xf1);
    ov4689_write_register(0x3307,0x04);
    ov4689_write_register(0x3309,0x29);
    ov4689_write_register(0x3500,0x00);
    ov4689_write_register(0x3501,0x4c);
    ov4689_write_register(0x3502,0x00);
    ov4689_write_register(0x3503,0x77);
    ov4689_write_register(0x3504,0x00);
    ov4689_write_register(0x3505,0x00);
    ov4689_write_register(0x3506,0x00);
    ov4689_write_register(0x3507,0x00);
    ov4689_write_register(0x3508,0x00);
    ov4689_write_register(0x3509,0x80);
    ov4689_write_register(0x350a,0x00);
    ov4689_write_register(0x350b,0x00);
    ov4689_write_register(0x350c,0x00);
    ov4689_write_register(0x350d,0x00);
    ov4689_write_register(0x350e,0x00);
    ov4689_write_register(0x350f,0x80);
    ov4689_write_register(0x3510,0x00);
    ov4689_write_register(0x3511,0x00);
    ov4689_write_register(0x3512,0x00);
    ov4689_write_register(0x3513,0x00);
    ov4689_write_register(0x3514,0x00);
    ov4689_write_register(0x3515,0x80);
    ov4689_write_register(0x3516,0x00);
    ov4689_write_register(0x3517,0x00);
    ov4689_write_register(0x3518,0x00);
    ov4689_write_register(0x3519,0x00);
    ov4689_write_register(0x351a,0x00);
    ov4689_write_register(0x351b,0x80);
    ov4689_write_register(0x351c,0x00);
    ov4689_write_register(0x351d,0x00);
    ov4689_write_register(0x351e,0x00);
    ov4689_write_register(0x351f,0x00);
    ov4689_write_register(0x3520,0x00);
    ov4689_write_register(0x3521,0x80);
    ov4689_write_register(0x3522,0x08);
    ov4689_write_register(0x3524,0x08);
    ov4689_write_register(0x3526,0x08);
    ov4689_write_register(0x3528,0x08);
    ov4689_write_register(0x352a,0x08);
    ov4689_write_register(0x3602,0x00);
    ov4689_write_register(0x3603,0x40);
    ov4689_write_register(0x3604,0x02);
    ov4689_write_register(0x3605,0x00);
    ov4689_write_register(0x3606,0x00);
    ov4689_write_register(0x3607,0x00);
    ov4689_write_register(0x3609,0x12);
    ov4689_write_register(0x360a,0x40);
    ov4689_write_register(0x360c,0x08);
    ov4689_write_register(0x360f,0xe5);
    ov4689_write_register(0x3608,0x8f);
    ov4689_write_register(0x3611,0x00);
    ov4689_write_register(0x3613,0xf7);
    ov4689_write_register(0x3616,0x58);
    ov4689_write_register(0x3619,0x99);
    ov4689_write_register(0x361b,0x60);
    ov4689_write_register(0x361c,0x7a);
    ov4689_write_register(0x361e,0x79);
    ov4689_write_register(0x361f,0x02);
    ov4689_write_register(0x3632,0x00);
    ov4689_write_register(0x3633,0x10);
    ov4689_write_register(0x3634,0x10);
    ov4689_write_register(0x3635,0x10);
    ov4689_write_register(0x3636,0x15);
    ov4689_write_register(0x3646,0x86);
    ov4689_write_register(0x364a,0x0b);
    ov4689_write_register(0x3700,0x17);
    ov4689_write_register(0x3701,0x22);
    ov4689_write_register(0x3703,0x10);
    ov4689_write_register(0x370a,0x37);
    ov4689_write_register(0x3705,0x00);
    ov4689_write_register(0x3706,0x63);
    ov4689_write_register(0x3709,0x3c);
    ov4689_write_register(0x370b,0x01);
    ov4689_write_register(0x370c,0x30);
    ov4689_write_register(0x3710,0x24);
    ov4689_write_register(0x3711,0x0c);
    ov4689_write_register(0x3716,0x00);
    ov4689_write_register(0x3720,0x28);
    ov4689_write_register(0x3729,0x7b);
    ov4689_write_register(0x372a,0x84);
    ov4689_write_register(0x372b,0xbd);
    ov4689_write_register(0x372c,0xbc);
    ov4689_write_register(0x372e,0x52);
    ov4689_write_register(0x373c,0x0e);
    ov4689_write_register(0x373e,0x33);
    ov4689_write_register(0x3743,0x10);
    ov4689_write_register(0x3744,0x88);
    ov4689_write_register(0x3745,0xc0);
    ov4689_write_register(0x374a,0x43);
    ov4689_write_register(0x374c,0x00);
    ov4689_write_register(0x374e,0x23);
    ov4689_write_register(0x3751,0x7b);
    ov4689_write_register(0x3752,0x84);
    ov4689_write_register(0x3753,0xbd);
    ov4689_write_register(0x3754,0xbc);
    ov4689_write_register(0x3756,0x52);
    ov4689_write_register(0x375c,0x00);
    ov4689_write_register(0x3760,0x00);
    ov4689_write_register(0x3761,0x00);
    ov4689_write_register(0x3762,0x00);
    ov4689_write_register(0x3763,0x00);
    ov4689_write_register(0x3764,0x00);
    ov4689_write_register(0x3767,0x04);
    ov4689_write_register(0x3768,0x04);
    ov4689_write_register(0x3769,0x08);
    ov4689_write_register(0x376a,0x08);
    ov4689_write_register(0x376b,0x20);
    ov4689_write_register(0x376c,0x00);
    ov4689_write_register(0x376d,0x00);
    ov4689_write_register(0x376e,0x00);
    ov4689_write_register(0x3773,0x00);
    ov4689_write_register(0x3774,0x51);
    ov4689_write_register(0x3776,0xbd);
    ov4689_write_register(0x3777,0xbd);
    ov4689_write_register(0x3781,0x18);
    ov4689_write_register(0x3783,0x25);
    ov4689_write_register(0x3798,0x1b);
    ov4689_write_register(0x3800,0x00);
    ov4689_write_register(0x3801,0x86);
    ov4689_write_register(0x3802,0x00);
    ov4689_write_register(0x3803,0xe0);
    ov4689_write_register(0x3804,0x09);
    ov4689_write_register(0x3805,0x17);
    ov4689_write_register(0x3806,0x05);
    ov4689_write_register(0x3807,0x1f);
    ov4689_write_register(0x3808,0x07);
    ov4689_write_register(0x3809,0x80);
    ov4689_write_register(0x380a,0x04);
    ov4689_write_register(0x380b,0x38);
    ov4689_write_register(0x380c,0x05);//;03
    ov4689_write_register(0x380d,0x07);//;5c
    ov4689_write_register(0x380e,0x06);//;04
    ov4689_write_register(0x380f,0x12);//;8A
    ov4689_write_register(0x3810,0x00);
    ov4689_write_register(0x3811,0x08);
    ov4689_write_register(0x3812,0x00);
    ov4689_write_register(0x3813,0x04);
    ov4689_write_register(0x3814,0x01);
    ov4689_write_register(0x3815,0x01);
    ov4689_write_register(0x3819,0x01);
    ov4689_write_register(0x3820,0x00);
    ov4689_write_register(0x3821,0x06);
    ov4689_write_register(0x3829,0x00);
    ov4689_write_register(0x382a,0x01);
    ov4689_write_register(0x382b,0x01);
    ov4689_write_register(0x382d,0x7f);
    ov4689_write_register(0x3830,0x04);
    ov4689_write_register(0x3836,0x01);
    ov4689_write_register(0x3837,0x00);
    ov4689_write_register(0x3841,0x02);
    ov4689_write_register(0x3846,0x08);
    ov4689_write_register(0x3847,0x07);
    ov4689_write_register(0x3d85,0x36);
    ov4689_write_register(0x3d8c,0x71);
    ov4689_write_register(0x3d8d,0xcb);
    ov4689_write_register(0x3f0a,0x00);
    ov4689_write_register(0x4000,0xF1);
    ov4689_write_register(0x4001,0x40);
    ov4689_write_register(0x4002,0x04);
    ov4689_write_register(0x4003,0x14);
    ov4689_write_register(0x400e,0x00);
    ov4689_write_register(0x4011,0x00);
    ov4689_write_register(0x401a,0x00);
    ov4689_write_register(0x401b,0x00);
    ov4689_write_register(0x401c,0x00);
    ov4689_write_register(0x401d,0x00);
    ov4689_write_register(0x401f,0x00);
    ov4689_write_register(0x4020,0x00);
    ov4689_write_register(0x4021,0x10);
    ov4689_write_register(0x4022,0x06);
    ov4689_write_register(0x4023,0x13);
    ov4689_write_register(0x4024,0x07);
    ov4689_write_register(0x4025,0x40);
    ov4689_write_register(0x4026,0x07);
    ov4689_write_register(0x4027,0x50);
    ov4689_write_register(0x4028,0x00);
    ov4689_write_register(0x4029,0x02);
    ov4689_write_register(0x402a,0x02);
    ov4689_write_register(0x402b,0x02);
    ov4689_write_register(0x402c,0x02);
    ov4689_write_register(0x402d,0x02);
    ov4689_write_register(0x402e,0x0e);
    ov4689_write_register(0x402f,0x04);
    ov4689_write_register(0x4302,0xff);
    ov4689_write_register(0x4303,0xff);
    ov4689_write_register(0x4304,0x00);
    ov4689_write_register(0x4305,0x00);
    ov4689_write_register(0x4306,0x00);
    ov4689_write_register(0x4308,0x02);
    ov4689_write_register(0x4500,0x6c);
    ov4689_write_register(0x4501,0xc4);
    ov4689_write_register(0x4502,0x40);
    ov4689_write_register(0x4503,0x01);
    ov4689_write_register(0x4601,0x41);
    ov4689_write_register(0x4800,0x04);
    ov4689_write_register(0x4813,0x08);
    ov4689_write_register(0x481f,0x40);
    ov4689_write_register(0x4829,0x78);
    ov4689_write_register(0x4837,0x1a);//;10
    ov4689_write_register(0x4b00,0x2a);
    ov4689_write_register(0x4b0d,0x00);
    ov4689_write_register(0x4d00,0x04);
    ov4689_write_register(0x4d01,0x42);
    ov4689_write_register(0x4d02,0xd1);
    ov4689_write_register(0x4d03,0x93);
    ov4689_write_register(0x4d04,0xf5);
    ov4689_write_register(0x4d05,0xc1);
    ov4689_write_register(0x5000,0xf3);
    ov4689_write_register(0x5001,0x11);
    ov4689_write_register(0x5004,0x00);
    ov4689_write_register(0x500a,0x00);
    ov4689_write_register(0x500b,0x00);
    ov4689_write_register(0x5032,0x00);
    ov4689_write_register(0x5040,0x00);
    ov4689_write_register(0x5050,0x0c);
    ov4689_write_register(0x5500,0x00);
    ov4689_write_register(0x5501,0x10);
    ov4689_write_register(0x5502,0x01);
    ov4689_write_register(0x5503,0x0f);
    ov4689_write_register(0x8000,0x00);
    ov4689_write_register(0x8001,0x00);
    ov4689_write_register(0x8002,0x00);
    ov4689_write_register(0x8003,0x00);
    ov4689_write_register(0x8004,0x00);
    ov4689_write_register(0x8005,0x00);
    ov4689_write_register(0x8006,0x00);
    ov4689_write_register(0x8007,0x00);
    ov4689_write_register(0x8008,0x00);
    ov4689_write_register(0x3638,0x00);
    //ov4689_write_register(0x0100,0x01);

    ov4689_enable_wdr_mode();
    bSensorInit_ov4689 = HI_TRUE;
    printf("-------OV4689 Sensor 1080P 30fps WDR Mode Initial OK!-------\n");
}

void ov4689_wdr_2304_1296_30_init()
{
    delay_ms(10);
    ov4689_write_register(0x0103, 0x01);

    ov4689_write_register(0x3638, 0x00);
    ov4689_write_register(0x0300, 0x00);
    ov4689_write_register(0x0301, 0x00);
    ov4689_write_register(0x0302, 0x19);
    ov4689_write_register(0x0303, 0x00);
    ov4689_write_register(0x0304, 0x03);
    ov4689_write_register(0x0305, 0x01);
    ov4689_write_register(0x0306, 0x01);
    ov4689_write_register(0x030A, 0x00);
    ov4689_write_register(0x030b, 0x00);
    ov4689_write_register(0x030c, 0x00);
    ov4689_write_register(0x030d, 0x1e);
    ov4689_write_register(0x030e, 0x04);
    ov4689_write_register(0x030f, 0x01);


    ov4689_write_register(0x0311, 0x00);
    ov4689_write_register(0x0312, 0x01);
    ov4689_write_register(0x031e, 0x00);
    ov4689_write_register(0x3000, 0x20);
    ov4689_write_register(0x3002, 0x00);
    ov4689_write_register(0x3018, 0x72);
    ov4689_write_register(0x3020, 0x93);
    ov4689_write_register(0x3021, 0x03);
    ov4689_write_register(0x3022, 0x01);
    ov4689_write_register(0x3031, 0x0a);
    ov4689_write_register(0x3305, 0xf1);
    ov4689_write_register(0x3307, 0x04);
    ov4689_write_register(0x3309, 0x29);
    ov4689_write_register(0x3500, 0x00);
    ov4689_write_register(0x3501, 0x45);
    ov4689_write_register(0x3502, 0xB0);
    ov4689_write_register(0x3503, 0x77);
    ov4689_write_register(0x3504, 0x00);
    ov4689_write_register(0x3505, 0x00);
    ov4689_write_register(0x3506, 0x00);
    ov4689_write_register(0x3507, 0x00);
    ov4689_write_register(0x3508, 0x00);
    ov4689_write_register(0x3509, 0x80);
    ov4689_write_register(0x350a, 0x00);
    ov4689_write_register(0x350b, 0x00);
    ov4689_write_register(0x350c, 0x00);
    ov4689_write_register(0x350d, 0x00);
    ov4689_write_register(0x350e, 0x00);
    ov4689_write_register(0x350f, 0x80);
    ov4689_write_register(0x3510, 0x00);
    ov4689_write_register(0x3511, 0x00);
    ov4689_write_register(0x3512, 0x00);
    ov4689_write_register(0x3513, 0x00);
    ov4689_write_register(0x3514, 0x00);
    ov4689_write_register(0x3515, 0x80);
    ov4689_write_register(0x3516, 0x00);
    ov4689_write_register(0x3517, 0x00);
    ov4689_write_register(0x3518, 0x00);
    ov4689_write_register(0x3519, 0x00);
    ov4689_write_register(0x351a, 0x00);
    ov4689_write_register(0x351b, 0x80);
    ov4689_write_register(0x351c, 0x00);
    ov4689_write_register(0x351d, 0x00);
    ov4689_write_register(0x351e, 0x00);
    ov4689_write_register(0x351f, 0x00);
    ov4689_write_register(0x3520, 0x00);
    ov4689_write_register(0x3521, 0x80);
    ov4689_write_register(0x3522, 0x08);
    ov4689_write_register(0x3524, 0x08);
    ov4689_write_register(0x3526, 0x08);
    ov4689_write_register(0x3528, 0x08);
    ov4689_write_register(0x352a, 0x08);
    ov4689_write_register(0x3602, 0x00);
    ov4689_write_register(0x3604, 0x02);
    ov4689_write_register(0x3605, 0x00);
    ov4689_write_register(0x3606, 0x00);
    ov4689_write_register(0x3607, 0x00);
    ov4689_write_register(0x3609, 0x12);
    ov4689_write_register(0x360a, 0x40);
    ov4689_write_register(0x360c, 0x08);
    ov4689_write_register(0x360f, 0xe5);
    ov4689_write_register(0x3608, 0x8f);
    ov4689_write_register(0x3611, 0x00);
    ov4689_write_register(0x3613, 0xf7);
    ov4689_write_register(0x3616, 0x58);
    ov4689_write_register(0x3619, 0x99);
    ov4689_write_register(0x361b, 0x60);
    ov4689_write_register(0x361c, 0x7a);
    ov4689_write_register(0x361e, 0x79);
    ov4689_write_register(0x361f, 0x02);
    ov4689_write_register(0x3632, 0x00);
    ov4689_write_register(0x3633, 0x10);
    ov4689_write_register(0x3634, 0x10);
    ov4689_write_register(0x3635, 0x10);
    ov4689_write_register(0x3636, 0x15);
    ov4689_write_register(0x3646, 0x86);
    ov4689_write_register(0x364a, 0x0b);
    ov4689_write_register(0x3700, 0x17);
    ov4689_write_register(0x3701, 0x22);
    ov4689_write_register(0x3703, 0x10);
    ov4689_write_register(0x370a, 0x37);
    ov4689_write_register(0x3705, 0x00);
    ov4689_write_register(0x3706, 0x63);
    ov4689_write_register(0x3709, 0x3c);
    ov4689_write_register(0x370b, 0x01);
    ov4689_write_register(0x370c, 0x30);
    ov4689_write_register(0x3710, 0x24);
    ov4689_write_register(0x3711, 0x0c);
    ov4689_write_register(0x3716, 0x00);
    ov4689_write_register(0x3720, 0x28);
    ov4689_write_register(0x3729, 0x7b);
    ov4689_write_register(0x372a, 0x84);
    ov4689_write_register(0x372b, 0xbd);
    ov4689_write_register(0x372c, 0xbc);
    ov4689_write_register(0x372e, 0x52);
    ov4689_write_register(0x373c, 0x0e);
    ov4689_write_register(0x373e, 0x33);
    ov4689_write_register(0x3743, 0x10);
    ov4689_write_register(0x3744, 0x88);
    ov4689_write_register(0x3745, 0xc0);
    ov4689_write_register(0x374a, 0x43);
    ov4689_write_register(0x374c, 0x00);
    ov4689_write_register(0x374e, 0x23);
    ov4689_write_register(0x3751, 0x7b);
    ov4689_write_register(0x3752, 0x84);
    ov4689_write_register(0x3753, 0xbd);
    ov4689_write_register(0x3754, 0xbc);
    ov4689_write_register(0x3756, 0x52);
    ov4689_write_register(0x375c, 0x00);
    ov4689_write_register(0x3760, 0x00);
    ov4689_write_register(0x3761, 0x00);
    ov4689_write_register(0x3762, 0x00);
    ov4689_write_register(0x3763, 0x00);
    ov4689_write_register(0x3764, 0x00);
    ov4689_write_register(0x3767, 0x04);
    ov4689_write_register(0x3768, 0x04);
    ov4689_write_register(0x3769, 0x08);
    ov4689_write_register(0x376a, 0x08);
    ov4689_write_register(0x376b, 0x20);
    ov4689_write_register(0x376c, 0x00);
    ov4689_write_register(0x376d, 0x00);
    ov4689_write_register(0x376e, 0x00);
    ov4689_write_register(0x3773, 0x00);
    ov4689_write_register(0x3774, 0x51);
    ov4689_write_register(0x3776, 0xbd);
    ov4689_write_register(0x3777, 0xbd);
    ov4689_write_register(0x3781, 0x18);
    ov4689_write_register(0x3783, 0x25);

    ov4689_write_register(0x3800,0x00);
    ov4689_write_register(0x3801,0xc6);
    ov4689_write_register(0x3802,0x00);
    ov4689_write_register(0x3803,0x74);
    ov4689_write_register(0x3804,0x09);
    ov4689_write_register(0x3805,0xd7);
    ov4689_write_register(0x3806,0x05);
    ov4689_write_register(0x3807,0x8b);
    ov4689_write_register(0x3808,0x09);
    ov4689_write_register(0x3809,0x00);
    ov4689_write_register(0x380a,0x05);
    ov4689_write_register(0x380b,0x10);

    ov4689_write_register(0x380c,0x05);
    ov4689_write_register(0x380d,0xee);
    ov4689_write_register(0x380e,0x05);
    ov4689_write_register(0x380f,0x2f);
    ov4689_write_register(0x3810,0x00);
    ov4689_write_register(0x3811,0x08);
    ov4689_write_register(0x3812,0x00);
    ov4689_write_register(0x3813,0x04);

    ov4689_write_register(0x3814, 0x01);
    ov4689_write_register(0x3815, 0x01);
    ov4689_write_register(0x3819, 0x01);
    ov4689_write_register(0x3820, 0x00);
    ov4689_write_register(0x3821, 0x06);
    ov4689_write_register(0x3829, 0x00);
    ov4689_write_register(0x382a, 0x01);
    ov4689_write_register(0x382b, 0x01);
    ov4689_write_register(0x382d, 0x7f);
    ov4689_write_register(0x3830, 0x04);
    ov4689_write_register(0x3836, 0x01);
    ov4689_write_register(0x3841, 0x02);
    ov4689_write_register(0x3846, 0x08);
    ov4689_write_register(0x3847, 0x07);
    ov4689_write_register(0x3d85, 0x36);
    ov4689_write_register(0x3d8c, 0x71);
    ov4689_write_register(0x3d8d, 0xcb);
    ov4689_write_register(0x3f0a, 0x00);
    ov4689_write_register(0x4000, 0x71);
    ov4689_write_register(0x4001, 0x40);
    ov4689_write_register(0x4002, 0x04);
    ov4689_write_register(0x4003, 0x14);
    ov4689_write_register(0x400e, 0x00);
    ov4689_write_register(0x4011, 0x00);
    ov4689_write_register(0x401a, 0x00);
    ov4689_write_register(0x401b, 0x00);
    ov4689_write_register(0x401c, 0x00);
    ov4689_write_register(0x401d, 0x00);
    ov4689_write_register(0x401f, 0x00);

    ov4689_write_register(0x4020,0x00);
    ov4689_write_register(0x4021,0x10);
    ov4689_write_register(0x4022,0x07);
    ov4689_write_register(0x4023,0x93);
    ov4689_write_register(0x4024,0x08);
    ov4689_write_register(0x4025,0xc0);
    ov4689_write_register(0x4026,0x08);
    ov4689_write_register(0x4027,0xd0);

    ov4689_write_register(0x4028, 0x00);
    ov4689_write_register(0x4029, 0x02);
    ov4689_write_register(0x402a, 0x02);
    ov4689_write_register(0x402b, 0x02);
    ov4689_write_register(0x402c, 0x02);
    ov4689_write_register(0x402d, 0x02);
    ov4689_write_register(0x402e, 0x0e);
    ov4689_write_register(0x402f, 0x04);
    ov4689_write_register(0x4302, 0xff);
    ov4689_write_register(0x4303, 0xff);
    ov4689_write_register(0x4304, 0x00);
    ov4689_write_register(0x4305, 0x00);
    ov4689_write_register(0x4306, 0x00);
    ov4689_write_register(0x4308, 0x02);
    ov4689_write_register(0x4500, 0x6c);
    ov4689_write_register(0x4501, 0xc4);
    ov4689_write_register(0x4502, 0x40);
    ov4689_write_register(0x4503, 0x01);

    ov4689_write_register(0x4600,0x00);
    ov4689_write_register(0x4601,0x41);

    ov4689_write_register(0x4800, 0x04);
    ov4689_write_register(0x4813, 0x08);
    ov4689_write_register(0x481f, 0x40);
    ov4689_write_register(0x4829, 0x78);
    ov4689_write_register(0x4837, 0x1a);
    ov4689_write_register(0x4b00, 0x2a);
    ov4689_write_register(0x4b0d, 0x00);
    ov4689_write_register(0x4d00, 0x04);
    ov4689_write_register(0x4d01, 0x42);
    ov4689_write_register(0x4d02, 0xd1);
    ov4689_write_register(0x4d03, 0x93);
    ov4689_write_register(0x4d04, 0xf5);
    ov4689_write_register(0x4d05, 0xc1);
    ov4689_write_register(0x5000, 0xf3);
    ov4689_write_register(0x5001, 0x11);
    ov4689_write_register(0x5004, 0x00);
    ov4689_write_register(0x500a, 0x00);
    ov4689_write_register(0x500b, 0x00);
    ov4689_write_register(0x5032, 0x00);
    ov4689_write_register(0x5040, 0x00);
    ov4689_write_register(0x5050, 0x0c);
    ov4689_write_register(0x8000, 0x00);
    ov4689_write_register(0x8001, 0x00);
    ov4689_write_register(0x8002, 0x00);
    ov4689_write_register(0x8003, 0x00);
    ov4689_write_register(0x8004, 0x00);
    ov4689_write_register(0x8005, 0x00);
    ov4689_write_register(0x8006, 0x00);
    ov4689_write_register(0x8007, 0x00);
    ov4689_write_register(0x8008, 0x00);
    ov4689_write_register(0x3638, 0x00);

    ov4689_enable_wdr_mode();
    bSensorInit_ov4689 = HI_TRUE;
    printf("-------OV4689 Sensor 2304_1296 30fps WDR Mode Initial OK!-------\n");
}


void ov4689_wdr_2048_1520_30_init()
{
    delay_ms(10);
    ov4689_write_register(0x0103,0x01);
    ov4689_write_register(0x3638,0x00);
    ov4689_write_register(0x0300,0x00);
    ov4689_write_register(0x0302,0x1a);//  ;2a
    ov4689_write_register(0x0303,0x00);
    ov4689_write_register(0x0304,0x03);
    ov4689_write_register(0x030b,0x00);
    ov4689_write_register(0x030d,0x1f);//  ;1e
    ov4689_write_register(0x030e,0x04);
    ov4689_write_register(0x030f,0x01);
    ov4689_write_register(0x0312,0x01);
    ov4689_write_register(0x031e,0x00);
    ov4689_write_register(0x3000,0x20);
    ov4689_write_register(0x3002,0x00);
    ov4689_write_register(0x3018,0x72);
    ov4689_write_register(0x3020,0x93);
    ov4689_write_register(0x3021,0x03);
    ov4689_write_register(0x3022,0x01);
    ov4689_write_register(0x3031,0x0a);
    ov4689_write_register(0x303f,0x0c);
    ov4689_write_register(0x3305,0xf1);
    ov4689_write_register(0x3307,0x04);
    ov4689_write_register(0x3309,0x29);
    ov4689_write_register(0x3500,0x00);
    ov4689_write_register(0x3501,0x60);
    ov4689_write_register(0x3502,0x00);
    ov4689_write_register(0x3503,0x77);
    ov4689_write_register(0x3504,0x00);
    ov4689_write_register(0x3505,0x00);
    ov4689_write_register(0x3506,0x00);
    ov4689_write_register(0x3507,0x00);
    ov4689_write_register(0x3508,0x00);
    ov4689_write_register(0x3509,0x80);
    ov4689_write_register(0x350a,0x00);
    ov4689_write_register(0x350b,0x00);
    ov4689_write_register(0x350c,0x00);
    ov4689_write_register(0x350d,0x00);
    ov4689_write_register(0x350e,0x00);
    ov4689_write_register(0x350f,0x80);
    ov4689_write_register(0x3510,0x00);
    ov4689_write_register(0x3511,0x00);
    ov4689_write_register(0x3512,0x00);
    ov4689_write_register(0x3513,0x00);
    ov4689_write_register(0x3514,0x00);
    ov4689_write_register(0x3515,0x80);
    ov4689_write_register(0x3516,0x00);
    ov4689_write_register(0x3517,0x00);
    ov4689_write_register(0x3518,0x00);
    ov4689_write_register(0x3519,0x00);
    ov4689_write_register(0x351a,0x00);
    ov4689_write_register(0x351b,0x80);
    ov4689_write_register(0x351c,0x00);
    ov4689_write_register(0x351d,0x00);
    ov4689_write_register(0x351e,0x00);
    ov4689_write_register(0x351f,0x00);
    ov4689_write_register(0x3520,0x00);
    ov4689_write_register(0x3521,0x80);
    ov4689_write_register(0x3522,0x08);
    ov4689_write_register(0x3524,0x08);
    ov4689_write_register(0x3526,0x08);
    ov4689_write_register(0x3528,0x08);
    ov4689_write_register(0x352a,0x08);
    ov4689_write_register(0x3602,0x00);
    ov4689_write_register(0x3603,0x40);
    ov4689_write_register(0x3604,0x02);
    ov4689_write_register(0x3605,0x00);
    ov4689_write_register(0x3606,0x00);
    ov4689_write_register(0x3607,0x00);
    ov4689_write_register(0x3609,0x12);
    ov4689_write_register(0x360a,0x40);
    ov4689_write_register(0x360c,0x08);
    ov4689_write_register(0x360f,0xe5);
    ov4689_write_register(0x3608,0x8f);
    ov4689_write_register(0x3611,0x00);
    ov4689_write_register(0x3613,0xf7);
    ov4689_write_register(0x3616,0x58);
    ov4689_write_register(0x3619,0x99);
    ov4689_write_register(0x361b,0x60);
    ov4689_write_register(0x361c,0x7a);
    ov4689_write_register(0x361e,0x79);
    ov4689_write_register(0x361f,0x02);
    ov4689_write_register(0x3632,0x00);
    ov4689_write_register(0x3633,0x10);
    ov4689_write_register(0x3634,0x10);
    ov4689_write_register(0x3635,0x10);
    ov4689_write_register(0x3636,0x15);
    ov4689_write_register(0x3646,0x86);
    ov4689_write_register(0x364a,0x0b);
    ov4689_write_register(0x3700,0x17);
    ov4689_write_register(0x3701,0x22);
    ov4689_write_register(0x3703,0x10);
    ov4689_write_register(0x370a,0x37);
    ov4689_write_register(0x3705,0x00);
    ov4689_write_register(0x3706,0x63);
    ov4689_write_register(0x3709,0x3c);
    ov4689_write_register(0x370b,0x01);
    ov4689_write_register(0x370c,0x30);
    ov4689_write_register(0x3710,0x24);
    ov4689_write_register(0x3711,0x0c);
    ov4689_write_register(0x3716,0x00);
    ov4689_write_register(0x3720,0x28);
    ov4689_write_register(0x3729,0x7b);
    ov4689_write_register(0x372a,0x84);
    ov4689_write_register(0x372b,0xbd);
    ov4689_write_register(0x372c,0xbc);
    ov4689_write_register(0x372e,0x52);
    ov4689_write_register(0x373c,0x0e);
    ov4689_write_register(0x373e,0x33);
    ov4689_write_register(0x3743,0x10);
    ov4689_write_register(0x3744,0x88);
    ov4689_write_register(0x3745,0xc0);
    ov4689_write_register(0x374a,0x43);
    ov4689_write_register(0x374c,0x00);
    ov4689_write_register(0x374e,0x23);
    ov4689_write_register(0x3751,0x7b);
    ov4689_write_register(0x3752,0x84);
    ov4689_write_register(0x3753,0xbd);
    ov4689_write_register(0x3754,0xbc);
    ov4689_write_register(0x3756,0x52);
    ov4689_write_register(0x375c,0x00);
    ov4689_write_register(0x3760,0x00);
    ov4689_write_register(0x3761,0x00);
    ov4689_write_register(0x3762,0x00);
    ov4689_write_register(0x3763,0x00);
    ov4689_write_register(0x3764,0x00);
    ov4689_write_register(0x3767,0x04);
    ov4689_write_register(0x3768,0x04);
    ov4689_write_register(0x3769,0x08);
    ov4689_write_register(0x376a,0x08);
    ov4689_write_register(0x376b,0x20);
    ov4689_write_register(0x376c,0x00);
    ov4689_write_register(0x376d,0x00);
    ov4689_write_register(0x376e,0x00);
    ov4689_write_register(0x3773,0x00);
    ov4689_write_register(0x3774,0x51);
    ov4689_write_register(0x3776,0xbd);
    ov4689_write_register(0x3777,0xbd);
    ov4689_write_register(0x3781,0x18);
    ov4689_write_register(0x3783,0x25);
    ov4689_write_register(0x3798,0x1b);
    ov4689_write_register(0x3800,0x01);// ;00
    ov4689_write_register(0x3801,0x48); //;08
    ov4689_write_register(0x3802,0x00);
    ov4689_write_register(0x3803,0x04);
    ov4689_write_register(0x3804,0x09); //;0a
    ov4689_write_register(0x3805,0x57); //;97
    ov4689_write_register(0x3806,0x05);
    ov4689_write_register(0x3807,0xfb);
    ov4689_write_register(0x3808,0x08);// ;0a
    ov4689_write_register(0x3809,0x0 );// ;20
    ov4689_write_register(0x380a,0x05);
    ov4689_write_register(0x380b,0xf0);
    ov4689_write_register(0x380c,0x05);// ;03
    ov4689_write_register(0x380d,0x30);// ;5c
    ov4689_write_register(0x380e,0x06);
    ov4689_write_register(0x380f,0x14);
    ov4689_write_register(0x3810,0x00);
    ov4689_write_register(0x3811,0x08);
    ov4689_write_register(0x3812,0x00);
    ov4689_write_register(0x3813,0x04);
    ov4689_write_register(0x3814,0x01);
    ov4689_write_register(0x3815,0x01);
    ov4689_write_register(0x3819,0x01);
    ov4689_write_register(0x3820,0x00);
    ov4689_write_register(0x3821,0x06);
    ov4689_write_register(0x3829,0x00);
    ov4689_write_register(0x382a,0x01);
    ov4689_write_register(0x382b,0x01);
    ov4689_write_register(0x382d,0x7f);
    ov4689_write_register(0x3830,0x04);
    ov4689_write_register(0x3836,0x01);
    ov4689_write_register(0x3837,0x00);
    ov4689_write_register(0x3841,0x02);
    ov4689_write_register(0x3846,0x08);
    ov4689_write_register(0x3847,0x07);
    ov4689_write_register(0x3d85,0x36);
    ov4689_write_register(0x3d8c,0x71);
    ov4689_write_register(0x3d8d,0xcb);
    ov4689_write_register(0x3f0a,0x00);
    ov4689_write_register(0x4000,0xf1);
    ov4689_write_register(0x4001,0x40);
    ov4689_write_register(0x4002,0x04);
    ov4689_write_register(0x4003,0x14);
    ov4689_write_register(0x400e,0x00);
    ov4689_write_register(0x4011,0x00);
    ov4689_write_register(0x401a,0x00);
    ov4689_write_register(0x401b,0x00);
    ov4689_write_register(0x401c,0x00);
    ov4689_write_register(0x401d,0x00);
    ov4689_write_register(0x401f,0x00);
    ov4689_write_register(0x4020,0x00);
    ov4689_write_register(0x4021,0x10);
    ov4689_write_register(0x4022,0x06);// ;07
    ov4689_write_register(0x4023,0x93);// ;cf
    ov4689_write_register(0x4024,0x07);// ;09
    ov4689_write_register(0x4025,0xc0);// ;60
    ov4689_write_register(0x4026,0x07);//;09
    ov4689_write_register(0x4027,0xD0);// ;6f
    ov4689_write_register(0x4028,0x00);
    ov4689_write_register(0x4029,0x02);
    ov4689_write_register(0x402a,0x02);
    ov4689_write_register(0x402b,0x02);
    ov4689_write_register(0x402c,0x02);
    ov4689_write_register(0x402d,0x02);
    ov4689_write_register(0x402e,0x0e);
    ov4689_write_register(0x402f,0x04);
    ov4689_write_register(0x4302,0xff);
    ov4689_write_register(0x4303,0xff);
    ov4689_write_register(0x4304,0x00);
    ov4689_write_register(0x4305,0x00);
    ov4689_write_register(0x4306,0x00);
    ov4689_write_register(0x4308,0x02);
    ov4689_write_register(0x4500,0x6c);
    ov4689_write_register(0x4501,0xc4);
    ov4689_write_register(0x4502,0x40);
    ov4689_write_register(0x4503,0x01);
    ov4689_write_register(0x4601,0x41);
    ov4689_write_register(0x4800,0x04);
    ov4689_write_register(0x4813,0x08);
    ov4689_write_register(0x481f,0x40);
    ov4689_write_register(0x4829,0x78);
    ov4689_write_register(0x4837,0x1a);// ;10
    ov4689_write_register(0x4b00,0x2a);
    ov4689_write_register(0x4b0d,0x00);
    ov4689_write_register(0x4d00,0x04);
    ov4689_write_register(0x4d01,0x42);
    ov4689_write_register(0x4d02,0xd1);
    ov4689_write_register(0x4d03,0x93);
    ov4689_write_register(0x4d04,0xf5);
    ov4689_write_register(0x4d05,0xc1);
    ov4689_write_register(0x5000,0xf3);
    ov4689_write_register(0x5001,0x11);
    ov4689_write_register(0x5004,0x00);
    ov4689_write_register(0x500a,0x00);
    ov4689_write_register(0x500b,0x00);
    ov4689_write_register(0x5032,0x00);
    ov4689_write_register(0x5040,0x00);
    ov4689_write_register(0x5050,0x0c);
    ov4689_write_register(0x5500,0x00);
    ov4689_write_register(0x5501,0x10);
    ov4689_write_register(0x5502,0x01);
    ov4689_write_register(0x5503,0x0f);
    ov4689_write_register(0x8000,0x00);
    ov4689_write_register(0x8001,0x00);
    ov4689_write_register(0x8002,0x00);
    ov4689_write_register(0x8003,0x00);
    ov4689_write_register(0x8004,0x00);
    ov4689_write_register(0x8005,0x00);
    ov4689_write_register(0x8006,0x00);
    ov4689_write_register(0x8007,0x00);
    ov4689_write_register(0x8008,0x00);
    ov4689_write_register(0x3638,0x00);
    //ov4689_write_register(0x0100,0x01);

    ov4689_enable_wdr_mode();
    bSensorInit_ov4689 = HI_TRUE;
    printf("-------OV4689 Sensor 2304_1296 30fps WDR Mode Initial OK!-------\n");
}

void ov4689_linear_2304_1296_30_init()
{
    delay_ms(10);
    ov4689_write_register(0x0103,0x01);

    ov4689_write_register(0x3638,0x00);
    ov4689_write_register(0x0300,0x00);
    ov4689_write_register(0x0302,0x19);//;2a
    ov4689_write_register(0x0303,0x00);
    ov4689_write_register(0x0304,0x03);
    ov4689_write_register(0x030b,0x00);
    ov4689_write_register(0x030d,0x1e);
    ov4689_write_register(0x030e,0x04);
    ov4689_write_register(0x030f,0x01);
    ov4689_write_register(0x0312,0x01);
    ov4689_write_register(0x031e,0x00);
    ov4689_write_register(0x3000,0x20);
    ov4689_write_register(0x3002,0x00);
    ov4689_write_register(0x3018,0x72);
    ov4689_write_register(0x3020,0x93);
    ov4689_write_register(0x3021,0x03);
    ov4689_write_register(0x3022,0x01);
    ov4689_write_register(0x3031,0x0a);
    ov4689_write_register(0x303f,0x0c);
    ov4689_write_register(0x3305,0xf1);
    ov4689_write_register(0x3307,0x04);
    ov4689_write_register(0x3309,0x29);
    ov4689_write_register(0x3500,0x00);
    ov4689_write_register(0x3501,0x44);
    ov4689_write_register(0x3502,0x00);
    ov4689_write_register(0x3503,0x77);
    ov4689_write_register(0x3504,0x00);
    ov4689_write_register(0x3505,0x00);
    ov4689_write_register(0x3506,0x00);
    ov4689_write_register(0x3507,0x00);
    ov4689_write_register(0x3508,0x00);
    ov4689_write_register(0x3509,0x80);
    ov4689_write_register(0x350a,0x00);
    ov4689_write_register(0x350b,0x00);
    ov4689_write_register(0x350c,0x00);
    ov4689_write_register(0x350d,0x00);
    ov4689_write_register(0x350e,0x00);
    ov4689_write_register(0x350f,0x80);
    ov4689_write_register(0x3510,0x00);
    ov4689_write_register(0x3511,0x00);
    ov4689_write_register(0x3512,0x00);
    ov4689_write_register(0x3513,0x00);
    ov4689_write_register(0x3514,0x00);
    ov4689_write_register(0x3515,0x80);
    ov4689_write_register(0x3516,0x00);
    ov4689_write_register(0x3517,0x00);
    ov4689_write_register(0x3518,0x00);
    ov4689_write_register(0x3519,0x00);
    ov4689_write_register(0x351a,0x00);
    ov4689_write_register(0x351b,0x80);
    ov4689_write_register(0x351c,0x00);
    ov4689_write_register(0x351d,0x00);
    ov4689_write_register(0x351e,0x00);
    ov4689_write_register(0x351f,0x00);
    ov4689_write_register(0x3520,0x00);
    ov4689_write_register(0x3521,0x80);
    ov4689_write_register(0x3522,0x08);
    ov4689_write_register(0x3524,0x08);
    ov4689_write_register(0x3526,0x08);
    ov4689_write_register(0x3528,0x08);
    ov4689_write_register(0x352a,0x08);
    ov4689_write_register(0x3602,0x00);
    ov4689_write_register(0x3603,0x40);
    ov4689_write_register(0x3604,0x02);
    ov4689_write_register(0x3605,0x00);
    ov4689_write_register(0x3606,0x00);
    ov4689_write_register(0x3607,0x00);
    ov4689_write_register(0x3609,0x12);
    ov4689_write_register(0x360a,0x40);
    ov4689_write_register(0x360c,0x08);
    ov4689_write_register(0x360f,0xe5);
    ov4689_write_register(0x3608,0x8f);
    ov4689_write_register(0x3611,0x00);
    ov4689_write_register(0x3613,0xf7);
    ov4689_write_register(0x3616,0x58);
    ov4689_write_register(0x3619,0x99);
    ov4689_write_register(0x361b,0x60);
    ov4689_write_register(0x361c,0x7a);
    ov4689_write_register(0x361e,0x79);
    ov4689_write_register(0x361f,0x02);
    ov4689_write_register(0x3632,0x00);
    ov4689_write_register(0x3633,0x10);
    ov4689_write_register(0x3634,0x10);
    ov4689_write_register(0x3635,0x10);
    ov4689_write_register(0x3636,0x15);
    ov4689_write_register(0x3646,0x86);
    ov4689_write_register(0x364a,0x0b);
    ov4689_write_register(0x3700,0x17);
    ov4689_write_register(0x3701,0x22);
    ov4689_write_register(0x3703,0x10);
    ov4689_write_register(0x370a,0x37);
    ov4689_write_register(0x3705,0x00);
    ov4689_write_register(0x3706,0x63);
    ov4689_write_register(0x3709,0x3c);
    ov4689_write_register(0x370b,0x01);
    ov4689_write_register(0x370c,0x30);
    ov4689_write_register(0x3710,0x24);
    ov4689_write_register(0x3711,0x0c);
    ov4689_write_register(0x3716,0x00);
    ov4689_write_register(0x3720,0x28);
    ov4689_write_register(0x3729,0x7b);
    ov4689_write_register(0x372a,0x84);
    ov4689_write_register(0x372b,0xbd);
    ov4689_write_register(0x372c,0xbc);
    ov4689_write_register(0x372e,0x52);
    ov4689_write_register(0x373c,0x0e);
    ov4689_write_register(0x373e,0x33);
    ov4689_write_register(0x3743,0x10);
    ov4689_write_register(0x3744,0x88);
    ov4689_write_register(0x3745,0xc0);
    ov4689_write_register(0x374a,0x43);
    ov4689_write_register(0x374c,0x00);
    ov4689_write_register(0x374e,0x23);
    ov4689_write_register(0x3751,0x7b);
    ov4689_write_register(0x3752,0x84);
    ov4689_write_register(0x3753,0xbd);
    ov4689_write_register(0x3754,0xbc);
    ov4689_write_register(0x3756,0x52);
    ov4689_write_register(0x375c,0x00);
    ov4689_write_register(0x3760,0x00);
    ov4689_write_register(0x3761,0x00);
    ov4689_write_register(0x3762,0x00);
    ov4689_write_register(0x3763,0x00);
    ov4689_write_register(0x3764,0x00);
    ov4689_write_register(0x3767,0x04);
    ov4689_write_register(0x3768,0x04);
    ov4689_write_register(0x3769,0x08);
    ov4689_write_register(0x376a,0x08);
    ov4689_write_register(0x376b,0x20);
    ov4689_write_register(0x376c,0x00);
    ov4689_write_register(0x376d,0x00);
    ov4689_write_register(0x376e,0x00);
    ov4689_write_register(0x3773,0x00);
    ov4689_write_register(0x3774,0x51);
    ov4689_write_register(0x3776,0xbd);
    ov4689_write_register(0x3777,0xbd);
    ov4689_write_register(0x3781,0x18);
    ov4689_write_register(0x3783,0x25);
    ov4689_write_register(0x3798,0x1b);

    ov4689_write_register(0x3800,0x00);
    ov4689_write_register(0x3801,0xC8);
    ov4689_write_register(0x3802,0x00);
    ov4689_write_register(0x3803,0x74);
    ov4689_write_register(0x3804,0x09);
    ov4689_write_register(0x3805,0xD7);
    ov4689_write_register(0x3806,0x05);
    ov4689_write_register(0x3807,0x8B);
    ov4689_write_register(0x3808,0x09);
    ov4689_write_register(0x3809,0x0 );
    ov4689_write_register(0x380A,0x05);
    ov4689_write_register(0x380B,0x10);

    ov4689_write_register(0x380c,0x0B);//a;03
    ov4689_write_register(0x380d,0xBB);//26;5c
    ov4689_write_register(0x380E,0x05);
    ov4689_write_register(0x380F,0x34);
    ov4689_write_register(0x3810,0x00);
    ov4689_write_register(0x3811,0x08);
    ov4689_write_register(0x3812,0x00);
    ov4689_write_register(0x3813,0x04);
    ov4689_write_register(0x3814,0x01);
    ov4689_write_register(0x3815,0x01);
    ov4689_write_register(0x3819,0x01);
    ov4689_write_register(0x3820,0x00);
    ov4689_write_register(0x3821,0x06);
    ov4689_write_register(0x3829,0x00);
    ov4689_write_register(0x382a,0x01);
    ov4689_write_register(0x382b,0x01);
    ov4689_write_register(0x382d,0x7f);
    ov4689_write_register(0x3830,0x04);
    ov4689_write_register(0x3836,0x01);
    ov4689_write_register(0x3837,0x00);
    ov4689_write_register(0x3841,0x02);
    ov4689_write_register(0x3846,0x08);
    ov4689_write_register(0x3847,0x07);
    ov4689_write_register(0x3d85,0x36);
    ov4689_write_register(0x3d8c,0x71);
    ov4689_write_register(0x3d8d,0xcb);
    ov4689_write_register(0x3f0a,0x00);
    ov4689_write_register(0x4000,0xf1);
    ov4689_write_register(0x4001,0x40);
    ov4689_write_register(0x4002,0x04);
    ov4689_write_register(0x4003,0x14);
    ov4689_write_register(0x400e,0x00);
    ov4689_write_register(0x4011,0x00);
    ov4689_write_register(0x401a,0x00);
    ov4689_write_register(0x401b,0x00);
    ov4689_write_register(0x401c,0x00);
    ov4689_write_register(0x401d,0x00);
    ov4689_write_register(0x401f,0x00);
    ov4689_write_register(0x4020,0x00);
    ov4689_write_register(0x4021,0x10);
    ov4689_write_register(0x4022,0x07);
    ov4689_write_register(0x4023,0x93);
    ov4689_write_register(0x4024,0x08);
    ov4689_write_register(0x4025,0xC0);
    ov4689_write_register(0x4026,0x08);
    ov4689_write_register(0x4027,0xD0);
    ov4689_write_register(0x4028,0x00);
    ov4689_write_register(0x4029,0x02);
    ov4689_write_register(0x402a,0x02);
    ov4689_write_register(0x402b,0x02);
    ov4689_write_register(0x402c,0x02);
    ov4689_write_register(0x402d,0x02);
    ov4689_write_register(0x402e,0x0e);
    ov4689_write_register(0x402f,0x04);
    ov4689_write_register(0x4302,0xff);
    ov4689_write_register(0x4303,0xff);
    ov4689_write_register(0x4304,0x00);
    ov4689_write_register(0x4305,0x00);
    ov4689_write_register(0x4306,0x00);
    ov4689_write_register(0x4308,0x02);
    ov4689_write_register(0x4500,0x6c);
    ov4689_write_register(0x4501,0xc4);
    ov4689_write_register(0x4502,0x40);
    ov4689_write_register(0x4503,0x01);
    ov4689_write_register(0x4600,0x00);
    ov4689_write_register(0x4601,0x6F);
    ov4689_write_register(0x4800,0x04);
    ov4689_write_register(0x4813,0x08);
    ov4689_write_register(0x481f,0x40);
    ov4689_write_register(0x4829,0x78);
    ov4689_write_register(0x4837,0x1a);//;10
    ov4689_write_register(0x4b00,0x2a);
    ov4689_write_register(0x4b0d,0x00);
    ov4689_write_register(0x4d00,0x04);
    ov4689_write_register(0x4d01,0x42);
    ov4689_write_register(0x4d02,0xd1);
    ov4689_write_register(0x4d03,0x93);
    ov4689_write_register(0x4d04,0xf5);
    ov4689_write_register(0x4d05,0xc1);
    ov4689_write_register(0x5000,0xf3);
    ov4689_write_register(0x5001,0x11);
    ov4689_write_register(0x5004,0x00);
    ov4689_write_register(0x500a,0x00);
    ov4689_write_register(0x500b,0x00);
    ov4689_write_register(0x5032,0x00);
    ov4689_write_register(0x5040,0x00);
    ov4689_write_register(0x5050,0x0c);
    ov4689_write_register(0x5500,0x00);
    ov4689_write_register(0x5501,0x10);
    ov4689_write_register(0x5502,0x01);
    ov4689_write_register(0x5503,0x0f);
    ov4689_write_register(0x8000,0x00);
    ov4689_write_register(0x8001,0x00);
    ov4689_write_register(0x8002,0x00);
    ov4689_write_register(0x8003,0x00);
    ov4689_write_register(0x8004,0x00);
    ov4689_write_register(0x8005,0x00);
    ov4689_write_register(0x8006,0x00);
    ov4689_write_register(0x8007,0x00);
    ov4689_write_register(0x8008,0x00);
    ov4689_write_register(0x3638,0x00);
    ov4689_write_register(0x0100,0x01);
    //ov4689_write_register(0x0300,0x60);
    bSensorInit_ov4689 = HI_TRUE;
    printf("-------OV4689 Sensor 3M 2304*1296 30fps Linear Mode Initial OK!-------\n");
}


void ov4689_linear_2048_1520_30_init()
{
    delay_ms(10);
    ov4689_write_register(0x0103,0x01);

    ov4689_write_register(0x3638,0x00);
    ov4689_write_register(0x0300,0x00);
    ov4689_write_register(0x0302,0x19);//;2a
    ov4689_write_register(0x0303,0x00);
    ov4689_write_register(0x0304,0x03);
    ov4689_write_register(0x030b,0x00);
    ov4689_write_register(0x030d,0x1e);
    ov4689_write_register(0x030e,0x04);
    ov4689_write_register(0x030f,0x01);
    ov4689_write_register(0x0312,0x01);
    ov4689_write_register(0x031e,0x00);
    ov4689_write_register(0x3000,0x20);
    ov4689_write_register(0x3002,0x00);
    ov4689_write_register(0x3018,0x72);
    ov4689_write_register(0x3020,0x93);
    ov4689_write_register(0x3021,0x03);
    ov4689_write_register(0x3022,0x01);
    ov4689_write_register(0x3031,0x0a);
    ov4689_write_register(0x303f,0x0c);
    ov4689_write_register(0x3305,0xf1);
    ov4689_write_register(0x3307,0x04);
    ov4689_write_register(0x3309,0x29);
    ov4689_write_register(0x3500,0x00);
    ov4689_write_register(0x3501,0x60);
    ov4689_write_register(0x3502,0x00);
    ov4689_write_register(0x3503,0x77);
    ov4689_write_register(0x3504,0x00);
    ov4689_write_register(0x3505,0x00);
    ov4689_write_register(0x3506,0x00);
    ov4689_write_register(0x3507,0x00);
    ov4689_write_register(0x3508,0x00);
    ov4689_write_register(0x3509,0x80);
    ov4689_write_register(0x350a,0x00);
    ov4689_write_register(0x350b,0x00);
    ov4689_write_register(0x350c,0x00);
    ov4689_write_register(0x350d,0x00);
    ov4689_write_register(0x350e,0x00);
    ov4689_write_register(0x350f,0x80);
    ov4689_write_register(0x3510,0x00);
    ov4689_write_register(0x3511,0x00);
    ov4689_write_register(0x3512,0x00);
    ov4689_write_register(0x3513,0x00);
    ov4689_write_register(0x3514,0x00);
    ov4689_write_register(0x3515,0x80);
    ov4689_write_register(0x3516,0x00);
    ov4689_write_register(0x3517,0x00);
    ov4689_write_register(0x3518,0x00);
    ov4689_write_register(0x3519,0x00);
    ov4689_write_register(0x351a,0x00);
    ov4689_write_register(0x351b,0x80);
    ov4689_write_register(0x351c,0x00);
    ov4689_write_register(0x351d,0x00);
    ov4689_write_register(0x351e,0x00);
    ov4689_write_register(0x351f,0x00);
    ov4689_write_register(0x3520,0x00);
    ov4689_write_register(0x3521,0x80);
    ov4689_write_register(0x3522,0x08);
    ov4689_write_register(0x3524,0x08);
    ov4689_write_register(0x3526,0x08);
    ov4689_write_register(0x3528,0x08);
    ov4689_write_register(0x352a,0x08);
    ov4689_write_register(0x3602,0x00);
    ov4689_write_register(0x3603,0x40);
    ov4689_write_register(0x3604,0x02);
    ov4689_write_register(0x3605,0x00);
    ov4689_write_register(0x3606,0x00);
    ov4689_write_register(0x3607,0x00);
    ov4689_write_register(0x3609,0x12);
    ov4689_write_register(0x360a,0x40);
    ov4689_write_register(0x360c,0x08);
    ov4689_write_register(0x360f,0xe5);
    ov4689_write_register(0x3608,0x8f);
    ov4689_write_register(0x3611,0x00);
    ov4689_write_register(0x3613,0xf7);
    ov4689_write_register(0x3616,0x58);
    ov4689_write_register(0x3619,0x99);
    ov4689_write_register(0x361b,0x60);
    ov4689_write_register(0x361c,0x7a);
    ov4689_write_register(0x361e,0x79);
    ov4689_write_register(0x361f,0x02);
    ov4689_write_register(0x3632,0x00);
    ov4689_write_register(0x3633,0x10);
    ov4689_write_register(0x3634,0x10);
    ov4689_write_register(0x3635,0x10);
    ov4689_write_register(0x3636,0x15);
    ov4689_write_register(0x3646,0x86);
    ov4689_write_register(0x364a,0x0b);
    ov4689_write_register(0x3700,0x17);
    ov4689_write_register(0x3701,0x22);
    ov4689_write_register(0x3703,0x10);
    ov4689_write_register(0x370a,0x37);
    ov4689_write_register(0x3705,0x00);
    ov4689_write_register(0x3706,0x63);
    ov4689_write_register(0x3709,0x3c);
    ov4689_write_register(0x370b,0x01);
    ov4689_write_register(0x370c,0x30);
    ov4689_write_register(0x3710,0x24);
    ov4689_write_register(0x3711,0x0c);
    ov4689_write_register(0x3716,0x00);
    ov4689_write_register(0x3720,0x28);
    ov4689_write_register(0x3729,0x7b);
    ov4689_write_register(0x372a,0x84);
    ov4689_write_register(0x372b,0xbd);
    ov4689_write_register(0x372c,0xbc);
    ov4689_write_register(0x372e,0x52);
    ov4689_write_register(0x373c,0x0e);
    ov4689_write_register(0x373e,0x33);
    ov4689_write_register(0x3743,0x10);
    ov4689_write_register(0x3744,0x88);
    ov4689_write_register(0x3745,0xc0);
    ov4689_write_register(0x374a,0x43);
    ov4689_write_register(0x374c,0x00);
    ov4689_write_register(0x374e,0x23);
    ov4689_write_register(0x3751,0x7b);
    ov4689_write_register(0x3752,0x84);
    ov4689_write_register(0x3753,0xbd);
    ov4689_write_register(0x3754,0xbc);
    ov4689_write_register(0x3756,0x52);
    ov4689_write_register(0x375c,0x00);
    ov4689_write_register(0x3760,0x00);
    ov4689_write_register(0x3761,0x00);
    ov4689_write_register(0x3762,0x00);
    ov4689_write_register(0x3763,0x00);
    ov4689_write_register(0x3764,0x00);
    ov4689_write_register(0x3767,0x04);
    ov4689_write_register(0x3768,0x04);
    ov4689_write_register(0x3769,0x08);
    ov4689_write_register(0x376a,0x08);
    ov4689_write_register(0x376b,0x20);
    ov4689_write_register(0x376c,0x00);
    ov4689_write_register(0x376d,0x00);
    ov4689_write_register(0x376e,0x00);
    ov4689_write_register(0x3773,0x00);
    ov4689_write_register(0x3774,0x51);
    ov4689_write_register(0x3776,0xbd);
    ov4689_write_register(0x3777,0xbd);
    ov4689_write_register(0x3781,0x18);
    ov4689_write_register(0x3783,0x25);
    ov4689_write_register(0x3798,0x1b);
    ov4689_write_register(0x3800,0x01);
    ov4689_write_register(0x3801,0x48);
    ov4689_write_register(0x3802,0x00);
    ov4689_write_register(0x3803,0x04);
    ov4689_write_register(0x3804,0x09);
    ov4689_write_register(0x3805,0x57);
    ov4689_write_register(0x3806,0x05);
    ov4689_write_register(0x3807,0xFB);
    ov4689_write_register(0x3808,0x08);
    ov4689_write_register(0x3809,0x00);
    ov4689_write_register(0x380A,0x05);
    ov4689_write_register(0x380B,0xF0);
    ov4689_write_register(0x380c,0x0A);//;03
    ov4689_write_register(0x380d,0x0B);//;5c
    ov4689_write_register(0x380E,0x06);
    ov4689_write_register(0x380F,0x14);
    ov4689_write_register(0x3810,0x00);
    ov4689_write_register(0x3811,0x08);
    ov4689_write_register(0x3812,0x00);
    ov4689_write_register(0x3813,0x04);
    ov4689_write_register(0x3814,0x01);
    ov4689_write_register(0x3815,0x01);
    ov4689_write_register(0x3819,0x01);
    ov4689_write_register(0x3820,0x00);
    ov4689_write_register(0x3821,0x06);
    ov4689_write_register(0x3829,0x00);
    ov4689_write_register(0x382a,0x01);
    ov4689_write_register(0x382b,0x01);
    ov4689_write_register(0x382d,0x7f);
    ov4689_write_register(0x3830,0x04);
    ov4689_write_register(0x3836,0x01);
    ov4689_write_register(0x3837,0x00);
    ov4689_write_register(0x3841,0x02);
    ov4689_write_register(0x3846,0x08);
    ov4689_write_register(0x3847,0x07);
    ov4689_write_register(0x3d85,0x36);
    ov4689_write_register(0x3d8c,0x71);
    ov4689_write_register(0x3d8d,0xcb);
    ov4689_write_register(0x3f0a,0x00);
    ov4689_write_register(0x4000,0xf1);
    ov4689_write_register(0x4001,0x40);
    ov4689_write_register(0x4002,0x04);
    ov4689_write_register(0x4003,0x14);
    ov4689_write_register(0x400e,0x00);
    ov4689_write_register(0x4011,0x00);
    ov4689_write_register(0x401a,0x00);
    ov4689_write_register(0x401b,0x00);
    ov4689_write_register(0x401c,0x00);
    ov4689_write_register(0x401d,0x00);
    ov4689_write_register(0x401f,0x00);
    ov4689_write_register(0x4020,0x00);
    ov4689_write_register(0x4021,0x10);
    ov4689_write_register(0x4022,0x06);
    ov4689_write_register(0x4023,0x93);
    ov4689_write_register(0x4024,0x07);
    ov4689_write_register(0x4025,0xC0);
    ov4689_write_register(0x4026,0x07);
    ov4689_write_register(0x4027,0xD0);
    ov4689_write_register(0x4028,0x00);
    ov4689_write_register(0x4029,0x02);
    ov4689_write_register(0x402a,0x02);
    ov4689_write_register(0x402b,0x02);
    ov4689_write_register(0x402c,0x02);
    ov4689_write_register(0x402d,0x02);
    ov4689_write_register(0x402e,0x0e);
    ov4689_write_register(0x402f,0x04);
    ov4689_write_register(0x4302,0xff);
    ov4689_write_register(0x4303,0xff);
    ov4689_write_register(0x4304,0x00);
    ov4689_write_register(0x4305,0x00);
    ov4689_write_register(0x4306,0x00);
    ov4689_write_register(0x4308,0x02);
    ov4689_write_register(0x4500,0x6c);
    ov4689_write_register(0x4501,0xc4);
    ov4689_write_register(0x4502,0x40);
    ov4689_write_register(0x4503,0x01);
    ov4689_write_register(0x4600,0x00);
    ov4689_write_register(0x4601,0x5f);
    ov4689_write_register(0x4800,0x04);
    ov4689_write_register(0x4813,0x08);
    ov4689_write_register(0x481f,0x40);
    ov4689_write_register(0x4829,0x78);
    ov4689_write_register(0x4837,0x1a);//10
    ov4689_write_register(0x4b00,0x2a);
    ov4689_write_register(0x4b0d,0x00);
    ov4689_write_register(0x4d00,0x04);
    ov4689_write_register(0x4d01,0x42);
    ov4689_write_register(0x4d02,0xd1);
    ov4689_write_register(0x4d03,0x93);
    ov4689_write_register(0x4d04,0xf5);
    ov4689_write_register(0x4d05,0xc1);
    ov4689_write_register(0x5000,0xf3);
    ov4689_write_register(0x5001,0x11);
    ov4689_write_register(0x5004,0x00);
    ov4689_write_register(0x500a,0x00);
    ov4689_write_register(0x500b,0x00);
    ov4689_write_register(0x5032,0x00);
    ov4689_write_register(0x5040,0x00);
    ov4689_write_register(0x5050,0x0c);
    ov4689_write_register(0x5500,0x00);
    ov4689_write_register(0x5501,0x10);
    ov4689_write_register(0x5502,0x01);
    ov4689_write_register(0x5503,0x0f);
    ov4689_write_register(0x8000,0x00);
    ov4689_write_register(0x8001,0x00);
    ov4689_write_register(0x8002,0x00);
    ov4689_write_register(0x8003,0x00);
    ov4689_write_register(0x8004,0x00);
    ov4689_write_register(0x8005,0x00);
    ov4689_write_register(0x8006,0x00);
    ov4689_write_register(0x8007,0x00);
    ov4689_write_register(0x8008,0x00);
    ov4689_write_register(0x3638,0x00);
    ov4689_write_register(0x0100,0x01);
    bSensorInit_ov4689 = HI_TRUE;
    printf("-------OV4689 Sensor 3M 2048*1520 30fps Linear Mode Initial OK!-------\n");
}
void ov4689_linear_720p_180fps_init()
{
    delay_ms(10);
    ov4689_write_register(0x0103,0x01);

    ov4689_write_register(0x3638,0x00);
    ov4689_write_register(0x0300,0x02);
    ov4689_write_register(0x0302,0x32);
    ov4689_write_register(0x0303,0x00);
    ov4689_write_register(0x0304,0x03);
    ov4689_write_register(0x030b,0x00);
    ov4689_write_register(0x030d,0x1e);
    ov4689_write_register(0x030e,0x04);
    ov4689_write_register(0x030f,0x01);
    ov4689_write_register(0x0312,0x01);
    ov4689_write_register(0x031e,0x00);
    ov4689_write_register(0x3000,0x20);
    ov4689_write_register(0x3002,0x00);
    ov4689_write_register(0x3018,0x72);
    ov4689_write_register(0x3020,0x93);
    ov4689_write_register(0x3021,0x03);
    ov4689_write_register(0x3022,0x01);
    ov4689_write_register(0x3031,0x0a);
    ov4689_write_register(0x303f,0x0c);
    ov4689_write_register(0x3305,0xf1);
    ov4689_write_register(0x3307,0x04);
    ov4689_write_register(0x3309,0x29);
    ov4689_write_register(0x3500,0x00);
    ov4689_write_register(0x3501,0x30);
    ov4689_write_register(0x3502,0x00);
    ov4689_write_register(0x3503,0x77);
    ov4689_write_register(0x3504,0x00);
    ov4689_write_register(0x3505,0x00);
    ov4689_write_register(0x3506,0x00);
    ov4689_write_register(0x3507,0x00);
    ov4689_write_register(0x3508,0x00);
    ov4689_write_register(0x3509,0x80);
    ov4689_write_register(0x350a,0x00);
    ov4689_write_register(0x350b,0x00);
    ov4689_write_register(0x350c,0x00);
    ov4689_write_register(0x350d,0x00);
    ov4689_write_register(0x350e,0x00);
    ov4689_write_register(0x350f,0x80);
    ov4689_write_register(0x3510,0x00);
    ov4689_write_register(0x3511,0x00);
    ov4689_write_register(0x3512,0x00);
    ov4689_write_register(0x3513,0x00);
    ov4689_write_register(0x3514,0x00);
    ov4689_write_register(0x3515,0x80);
    ov4689_write_register(0x3516,0x00);
    ov4689_write_register(0x3517,0x00);
    ov4689_write_register(0x3518,0x00);
    ov4689_write_register(0x3519,0x00);
    ov4689_write_register(0x351a,0x00);
    ov4689_write_register(0x351b,0x80);
    ov4689_write_register(0x351c,0x00);
    ov4689_write_register(0x351d,0x00);
    ov4689_write_register(0x351e,0x00);
    ov4689_write_register(0x351f,0x00);
    ov4689_write_register(0x3520,0x00);
    ov4689_write_register(0x3521,0x80);
    ov4689_write_register(0x3522,0x08);
    ov4689_write_register(0x3524,0x08);
    ov4689_write_register(0x3526,0x08);
    ov4689_write_register(0x3528,0x08);
    ov4689_write_register(0x352a,0x08);
    ov4689_write_register(0x3602,0x00);
    ov4689_write_register(0x3603,0x40);
    ov4689_write_register(0x3604,0x02);
    ov4689_write_register(0x3605,0x00);
    ov4689_write_register(0x3606,0x00);
    ov4689_write_register(0x3607,0x00);
    ov4689_write_register(0x3609,0x12);
    ov4689_write_register(0x360a,0x40);
    ov4689_write_register(0x360c,0x08);
    ov4689_write_register(0x360f,0xe5);
    ov4689_write_register(0x3608,0x8f);
    ov4689_write_register(0x3611,0x00);
    ov4689_write_register(0x3613,0xf7);
    ov4689_write_register(0x3616,0x58);
    ov4689_write_register(0x3619,0x99);
    ov4689_write_register(0x361b,0x60);
    ov4689_write_register(0x361c,0x7a);
    ov4689_write_register(0x361e,0x79);
    ov4689_write_register(0x361f,0x02);
    ov4689_write_register(0x3632,0x05);
    ov4689_write_register(0x3633,0x10);
    ov4689_write_register(0x3634,0x10);
    ov4689_write_register(0x3635,0x10);
    ov4689_write_register(0x3636,0x15);
    ov4689_write_register(0x3646,0x86);
    ov4689_write_register(0x364a,0x0b);
    ov4689_write_register(0x3700,0x17);
    ov4689_write_register(0x3701,0x22);
    ov4689_write_register(0x3703,0x10);
    ov4689_write_register(0x370a,0x37);
    ov4689_write_register(0x3705,0x00);
    ov4689_write_register(0x3706,0x63);
    ov4689_write_register(0x3709,0x3c);
    ov4689_write_register(0x370b,0x01);
    ov4689_write_register(0x370c,0x30);
    ov4689_write_register(0x3710,0x24);
    ov4689_write_register(0x3711,0x0c);
    ov4689_write_register(0x3716,0x00);
    ov4689_write_register(0x3720,0x28);
    ov4689_write_register(0x3729,0x7b);
    ov4689_write_register(0x372a,0x84);
    ov4689_write_register(0x372b,0xbd);
    ov4689_write_register(0x372c,0xbc);
    ov4689_write_register(0x372e,0x52);
    ov4689_write_register(0x373c,0x0e);
    ov4689_write_register(0x373e,0x33);
    ov4689_write_register(0x3743,0x10);
    ov4689_write_register(0x3744,0x88);
    ov4689_write_register(0x3745,0xc0);
    ov4689_write_register(0x374a,0x43);
    ov4689_write_register(0x374c,0x00);
    ov4689_write_register(0x374e,0x23);
    ov4689_write_register(0x3751,0x7b);
    ov4689_write_register(0x3752,0x84);
    ov4689_write_register(0x3753,0xbd);
    ov4689_write_register(0x3754,0xbc);
    ov4689_write_register(0x3756,0x52);
    ov4689_write_register(0x375c,0x00);
    ov4689_write_register(0x3760,0x00);
    ov4689_write_register(0x3761,0x00);
    ov4689_write_register(0x3762,0x00);
    ov4689_write_register(0x3763,0x00);
    ov4689_write_register(0x3764,0x00);
    ov4689_write_register(0x3767,0x04);
    ov4689_write_register(0x3768,0x04);
    ov4689_write_register(0x3769,0x08);
    ov4689_write_register(0x376a,0x08);
    ov4689_write_register(0x376b,0x40);
    ov4689_write_register(0x376c,0x00);
    ov4689_write_register(0x376d,0x00);
    ov4689_write_register(0x376e,0x00);
    ov4689_write_register(0x3773,0x00);
    ov4689_write_register(0x3774,0x51);
    ov4689_write_register(0x3776,0xbd);
    ov4689_write_register(0x3777,0xbd);
    ov4689_write_register(0x3781,0x18);
    ov4689_write_register(0x3783,0x25);
    ov4689_write_register(0x3798,0x1b);
    ov4689_write_register(0x3800,0x00);
    ov4689_write_register(0x3801,0x48);
    ov4689_write_register(0x3802,0x00);
    ov4689_write_register(0x3803,0x2C);
    ov4689_write_register(0x3804,0x0a);
    ov4689_write_register(0x3805,0x57);
    ov4689_write_register(0x3806,0x05);
    ov4689_write_register(0x3807,0xD3);
    ov4689_write_register(0x3808,0x05);
    ov4689_write_register(0x3809,0x00);
    ov4689_write_register(0x380a,0x02);
    ov4689_write_register(0x380b,0xD0);
    ov4689_write_register(0x380c,0x03);
    ov4689_write_register(0x380d,0x5e);
    ov4689_write_register(0x380e,0x03);
    ov4689_write_register(0x380f,0x05);
    ov4689_write_register(0x3810,0x00);
    ov4689_write_register(0x3811,0x04);
    ov4689_write_register(0x3812,0x00);
    ov4689_write_register(0x3813,0x02);
    ov4689_write_register(0x3814,0x03);
    ov4689_write_register(0x3815,0x01);
    ov4689_write_register(0x3819,0x01);
    ov4689_write_register(0x3820,0x10);
    ov4689_write_register(0x3821,0x07);
    ov4689_write_register(0x3829,0x00);
    ov4689_write_register(0x382a,0x03);
    ov4689_write_register(0x382b,0x01);
    ov4689_write_register(0x382d,0x7f);
    ov4689_write_register(0x3830,0x08);
    ov4689_write_register(0x3836,0x02);
    ov4689_write_register(0x3837,0x00);
    ov4689_write_register(0x3841,0x02);
    ov4689_write_register(0x3846,0x08);
    ov4689_write_register(0x3847,0x07);
    ov4689_write_register(0x3d85,0x36);
    ov4689_write_register(0x3d8c,0x71);
    ov4689_write_register(0x3d8d,0xcb);
    ov4689_write_register(0x3f0a,0x00);
    ov4689_write_register(0x4000,0xf1);
    ov4689_write_register(0x4001,0x50);
    ov4689_write_register(0x4002,0x04);
    ov4689_write_register(0x4003,0x14);
    ov4689_write_register(0x400e,0x00);
    ov4689_write_register(0x4011,0x00);
    ov4689_write_register(0x401a,0x00);
    ov4689_write_register(0x401b,0x00);
    ov4689_write_register(0x401c,0x00);
    ov4689_write_register(0x401d,0x00);
    ov4689_write_register(0x401f,0x00);
    ov4689_write_register(0x4020,0x00);
    ov4689_write_register(0x4021,0x10);
    ov4689_write_register(0x4022,0x03);
    ov4689_write_register(0x4023,0x93);
    ov4689_write_register(0x4024,0x04);
    ov4689_write_register(0x4025,0xC0);
    ov4689_write_register(0x4026,0x04);
    ov4689_write_register(0x4027,0xD0);
    ov4689_write_register(0x4028,0x00);
    ov4689_write_register(0x4029,0x02);
    ov4689_write_register(0x402a,0x02);
    ov4689_write_register(0x402b,0x02);
    ov4689_write_register(0x402c,0x02);
    ov4689_write_register(0x402d,0x02);
    ov4689_write_register(0x402e,0x0e);
    ov4689_write_register(0x402f,0x04);
    ov4689_write_register(0x4302,0xff);
    ov4689_write_register(0x4303,0xff);
    ov4689_write_register(0x4304,0x00);
    ov4689_write_register(0x4305,0x00);
    ov4689_write_register(0x4306,0x00);
    ov4689_write_register(0x4308,0x02);
    ov4689_write_register(0x4500,0x6c);
    ov4689_write_register(0x4501,0xc4);
    ov4689_write_register(0x4502,0x44);
    ov4689_write_register(0x4503,0x01);
    ov4689_write_register(0x4601,0x2F);//;4F
    ov4689_write_register(0x4800,0x04);
    ov4689_write_register(0x4813,0x08);
    ov4689_write_register(0x481f,0x40);
    ov4689_write_register(0x4829,0x78);
    ov4689_write_register(0x4837,0x1b);
    ov4689_write_register(0x4b00,0x2a);
    ov4689_write_register(0x4b0d,0x00);
    ov4689_write_register(0x4d00,0x04);
    ov4689_write_register(0x4d01,0x42);
    ov4689_write_register(0x4d02,0xd1);
    ov4689_write_register(0x4d03,0x93);
    ov4689_write_register(0x4d04,0xf5);
    ov4689_write_register(0x4d05,0xc1);
    ov4689_write_register(0x5000,0xf3);
    ov4689_write_register(0x5001,0x11);
    ov4689_write_register(0x5004,0x00);
    ov4689_write_register(0x500a,0x00);
    ov4689_write_register(0x500b,0x00);
    ov4689_write_register(0x5032,0x00);
    ov4689_write_register(0x5040,0x00);
    ov4689_write_register(0x5050,0x3c);
    ov4689_write_register(0x5500,0x00);
    ov4689_write_register(0x5501,0x10);
    ov4689_write_register(0x5502,0x01);
    ov4689_write_register(0x5503,0x0f);
    ov4689_write_register(0x8000,0x00);
    ov4689_write_register(0x8001,0x00);
    ov4689_write_register(0x8002,0x00);
    ov4689_write_register(0x8003,0x00);
    ov4689_write_register(0x8004,0x00);
    ov4689_write_register(0x8005,0x00);
    ov4689_write_register(0x8006,0x00);
    ov4689_write_register(0x8007,0x00);
    ov4689_write_register(0x8008,0x00);
    ov4689_write_register(0x3638,0x00);
    ov4689_write_register(0x0100,0x01);
    bSensorInit_ov4689 = HI_TRUE;
    printf("-------OV4689 Sensor 720180fps Linear Mode Initial OK!-------\n");
}
