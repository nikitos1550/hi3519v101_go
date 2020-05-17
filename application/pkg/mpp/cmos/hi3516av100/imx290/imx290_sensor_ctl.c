/******************************************************************************

  Copyright (C), 2001-2013, Hisilicon Tech. Co., Ltd.

 ******************************************************************************
  File Name     : imx290_sensor_ctl.c
  Version       : Initial Draft
  Author        : Hisilicon BVT AE group
  Created       : 2015/03/26
  Description   : Sony IMX290 sensor driver
  History       :
  1.Date        : 2015/03/26
  Author        : x00226337
  Modification  : Created file

******************************************************************************/

#include <stdio.h>
#include <string.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <sys/ioctl.h>
#include <fcntl.h>
#include <unistd.h>
#include "hi_comm_video.h"

#include "hi_spi.h"
#include "imx290_def.h"

extern WDR_MODE_E genSensorMode;
extern HI_BOOL bSensorInit;
extern HI_U8 gu8SensorImageMode;
static int g_fd = -1;


int sensor_spi_init(void)
{
    if(g_fd >= 0)
    {
        return 0;
    }    
    unsigned int value;
    int ret = 0;
    char file_name[] = "/dev/spidev0.0";
    
    g_fd = open(file_name, 0);
    if (g_fd < 0)
    {
        printf("Open %s error!\n",file_name);
        return -1;
    }

    value = SPI_MODE_3 | SPI_LSB_FIRST;// | SPI_LOOP;
    ret = ioctl(g_fd, SPI_IOC_WR_MODE, &value);
    if (ret < 0)
    {
        printf("ioctl SPI_IOC_WR_MODE err, value = %d ret = %d\n", value, ret);
        return ret;
    }

    value = 8;
    ret = ioctl(g_fd, SPI_IOC_WR_BITS_PER_WORD, &value);
    if (ret < 0)
    {
        printf("ioctl SPI_IOC_WR_BITS_PER_WORD err, value = %d ret = %d\n",value, ret);
        return ret;
    }

    value = 2000000;
    ret = ioctl(g_fd, SPI_IOC_WR_MAX_SPEED_HZ, &value);
    if (ret < 0)
    {
        printf("ioctl SPI_IOC_WR_MAX_SPEED_HZ err, value = %d ret = %d\n",value, ret);
        return ret;
    }

    return 0;
}

int sensor_spi_exit(void)
{
    if (g_fd >= 0)
    {
        close(g_fd);
        g_fd = -1;
        return 0;
    }
    return -1;
}

int sensor_write_register(unsigned int addr, unsigned char data)
{
    int ret;
    struct spi_ioc_transfer mesg[1];
    unsigned char  tx_buf[8] = {0};
    unsigned char  rx_buf[8] = {0};
    
    tx_buf[0] = (addr & 0xff00) >> 8;
    tx_buf[0] &= (~0x80);
    tx_buf[1] = addr & 0xff;
    tx_buf[2] = data;


    memset(mesg, 0, sizeof(mesg));  
    mesg[0].tx_buf = (__u32)tx_buf;  
    mesg[0].len    = 3;  
    mesg[0].rx_buf = (__u32)rx_buf; 
    mesg[0].cs_change = 1;

    ret = ioctl(g_fd, SPI_IOC_MESSAGE(1), mesg);
    if (ret < 0) {  
        printf("SPI_IOC_MESSAGE error \n");  
        return -1;  
    }

    return 0;
}

int sensor_read_register(unsigned int addr)
{
    int ret = 0;
    struct spi_ioc_transfer mesg[1];
    unsigned char  tx_buf[8] = {0};
    unsigned char  rx_buf[8] = {0};
    
    tx_buf[0] = (addr & 0xff00) >> 8;
    tx_buf[0] |= 0x80;
    tx_buf[1] = addr & 0xff;
    tx_buf[2] = 0;

    memset(mesg, 0, sizeof(mesg));
    mesg[0].tx_buf = (__u32)tx_buf;
    mesg[0].len    = 3;
    mesg[0].rx_buf = (__u32)rx_buf;
    mesg[0].cs_change = 1;

    ret = ioctl(g_fd, SPI_IOC_MESSAGE(1), mesg);
    if (ret  < 0) {  
        printf("SPI_IOC_MESSAGE error \n");  
        return -1;  
    }

    return rx_buf[2];
}

static void delay_ms(int ms) { 
    usleep(ms*1000);
}

void sensor_prog(int* rom) 
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
            sensor_write_register(addr, data);
        }
    }
}


void sensor_wdr_1080p30_init();
void sensor_wdr_720p60_init();
void sensor_linear_1080p30_init();
void sensor_linear_1080p60_init();
void sensor_linear_720p120_init();


void sensor_init()
{  
    /* 1. sensor spi init */
    sensor_spi_init();

    if(WDR_MODE_2To1_LINE == genSensorMode)
    {
        if(SENSOR_IMX290_1080P_30FPS_MODE == gu8SensorImageMode)
        {
            sensor_wdr_1080p30_init();
            bSensorInit = HI_TRUE;
        }
        else if(SENSOR_IMX290_720P_60FPS_MODE == gu8SensorImageMode)
        {
            sensor_wdr_720p60_init();
            bSensorInit = HI_TRUE;
        }
        else
        {
            printf("Not support this mode!\n");
        }
    }
    else
    {
        if(SENSOR_IMX290_1080P_30FPS_MODE == gu8SensorImageMode)
        {
             sensor_linear_1080p30_init();
             bSensorInit = HI_TRUE;
        }
        else if(SENSOR_IMX290_1080P_60FPS_MODE == gu8SensorImageMode)
        {
            sensor_linear_1080p60_init(); 
            bSensorInit = HI_TRUE;
        }  
        else if(SENSOR_IMX290_720P_120FPS_MODE == gu8SensorImageMode)
        {
            sensor_linear_720p120_init();
            bSensorInit = HI_TRUE;
        }
        else
        {
            printf("Not support this mode!\n");
        }
    }       


    return ;
}

void sensor_exit()
{
    sensor_spi_exit();
    return;
}


// Normal mode 
// INCK=37.125MHz
// 1080p 60fps
// LVDS Serial 4ch(445.5Mbps/ch)
// ADC 12bit
void sensor_linear_1080p60_init()
 {
    /* imx290 1080p30 */    
    sensor_write_register (0x200, 0x01); /* standby */

    delay_ms(200); 

    sensor_write_register (0x207, 0x00);  
    sensor_write_register (0x209, 0x01);
    sensor_write_register (0x20A, 0xF0);
    sensor_write_register (0x20F, 0x00);
    sensor_write_register (0x210, 0x21);
    sensor_write_register (0x212, 0x64);
    sensor_write_register (0x214, 0x20);//Gain
    sensor_write_register (0x216, 0x09);
    sensor_write_register (0x216, 0x09);

    sensor_write_register (0x218, 0x65); /* VMAX[7:0] */
    sensor_write_register (0x219, 0x04); /* VMAX[15:8] */
    sensor_write_register (0x21a, 0x00); /* VMAX[16] */
    sensor_write_register (0x21b, 0x00); 
    sensor_write_register (0x21c, 0x98); /* HMAX[7:0] */
    sensor_write_register (0x21d, 0x08); /* HMAX[15:8] */

    
    sensor_write_register (0x220, 0x01);//SHS1
    sensor_write_register (0x221, 0x00);

    sensor_write_register (0x246, 0xE1);//LANE CHN
    sensor_write_register (0x25C, 0x18);
    sensor_write_register (0x25E, 0x20);
    
    sensor_write_register (0x270, 0x02);
    sensor_write_register (0x271, 0x11);
    sensor_write_register (0x29B, 0x10);
    sensor_write_register (0x29C, 0x22);
    sensor_write_register (0x2A2, 0x02);
    sensor_write_register (0x2A6, 0x20);
    sensor_write_register (0x2A8, 0x20);
    sensor_write_register (0x2AA, 0x20);
    sensor_write_register (0x2AC, 0x20);
    sensor_write_register (0x2B0, 0x43);
    
    sensor_write_register (0x319, 0x9E);
    sensor_write_register (0x31C, 0x1E);
    sensor_write_register (0x31E, 0x08);
    sensor_write_register (0x328, 0x05);
    sensor_write_register (0x33D, 0x83);
    sensor_write_register (0x350, 0x03);

    sensor_write_register (0x35E, 0x1A);// 1A:37.125MHz 1B:74.25MHz
    sensor_write_register (0x364, 0x1A);// 1A:37.125MHz 1B:74.25MHz

    sensor_write_register (0x37C, 0x00);
    sensor_write_register (0x37E, 0x00);

    sensor_write_register (0x37E, 0x00);

    sensor_write_register (0x4B8, 0x50);
    sensor_write_register (0x4B9, 0x10);
    sensor_write_register (0x4BA, 0x00);
    sensor_write_register (0x4BB, 0x04);
    sensor_write_register (0x4C8, 0x50);
    sensor_write_register (0x4C9, 0x10);
    sensor_write_register (0x4CA, 0x00);
    sensor_write_register (0x4CB, 0x04);

    sensor_write_register (0x52C, 0xD3);
    sensor_write_register (0x52D, 0x10);
    sensor_write_register (0x52E, 0x0D);
    sensor_write_register (0x558, 0x06);
    sensor_write_register (0x559, 0xE1);
    sensor_write_register (0x55A, 0x11);
    sensor_write_register (0x560, 0x1E);
    sensor_write_register (0x561, 0x61);
    sensor_write_register (0x562, 0x10);
    sensor_write_register (0x5B0, 0x50);
    sensor_write_register (0x5B2, 0x1A);
    sensor_write_register (0x5B3, 0x04);

    delay_ms(200);
    sensor_write_register (0x200, 0x00); /* standby */
    sensor_write_register (0x202, 0x00); /* master mode start */
    sensor_write_register (0x249, 0x0A); /* XVSOUTSEL XHSOUTSEL */
    
    printf("--IMX290 1080P 60fps LINE Init OK!----\n");
    
}





// Normal mode 
// INCK=37.125MHz
// 1080p 30fps
// LVDS Serial 4ch (222.75Mbps/ch)
// ADC 12bit
void sensor_linear_1080p30_init()
 {
/* imx290 1080p30 */    
    sensor_write_register (0x200, 0x01); /* standby */

    delay_ms(200); 

    sensor_write_register (0x207, 0x00);  
    sensor_write_register (0x209, 0x02);
    sensor_write_register (0x20A, 0xF0);
    sensor_write_register (0x20F, 0x00);
    sensor_write_register (0x210, 0x21);
    sensor_write_register (0x212, 0x64);
    sensor_write_register (0x214, 0x20);//Gain
    sensor_write_register (0x216, 0x09);
    sensor_write_register (0x216, 0x09);

    sensor_write_register (0x218, 0x6D); /* VMAX[7:0] */
    sensor_write_register (0x219, 0x04); /* VMAX[15:8] */
    sensor_write_register (0x21a, 0x00); /* VMAX[16] */
    sensor_write_register (0x21b, 0x00); 
    sensor_write_register (0x21c, 0x30); /* HMAX[7:0] */
    sensor_write_register (0x21c, 0x11); /* HMAX[15:8] */

    
    sensor_write_register (0x220, 0x01);//SHS1
    sensor_write_register (0x221, 0x00);

    sensor_write_register (0x246, 0xE1);//LANE CHN
    sensor_write_register (0x25C, 0x18);
    sensor_write_register (0x25E, 0x20);
    
    sensor_write_register (0x270, 0x02);
    sensor_write_register (0x271, 0x11);
    sensor_write_register (0x29B, 0x10);
    sensor_write_register (0x29C, 0x22);
    sensor_write_register (0x2A2, 0x02);
    sensor_write_register (0x2A6, 0x20);
    sensor_write_register (0x2A8, 0x20);
    sensor_write_register (0x2AA, 0x20);
    sensor_write_register (0x2AC, 0x20);
    sensor_write_register (0x2B0, 0x43);
    
    sensor_write_register (0x319, 0x9E);
    sensor_write_register (0x31C, 0x1E);
    sensor_write_register (0x31E, 0x08);
    sensor_write_register (0x328, 0x05);
    sensor_write_register (0x33D, 0x83);
    sensor_write_register (0x350, 0x03);

    sensor_write_register (0x35E, 0x1A);// 1A:37.125MHz 1B:74.25MHz
    sensor_write_register (0x364, 0x1A);// 1A:37.125MHz 1B:74.25MHz

    sensor_write_register (0x37C, 0x00);
    sensor_write_register (0x37E, 0x00);

    sensor_write_register (0x37E, 0x00);

    sensor_write_register (0x4B8, 0x50);
    sensor_write_register (0x4B9, 0x10);
    sensor_write_register (0x4BA, 0x00);
    sensor_write_register (0x4BB, 0x04);
    sensor_write_register (0x4C8, 0x50);
    sensor_write_register (0x4C9, 0x10);
    sensor_write_register (0x4CA, 0x00);
    sensor_write_register (0x4CB, 0x04);

    sensor_write_register (0x52C, 0xD3);
    sensor_write_register (0x52D, 0x10);
    sensor_write_register (0x52E, 0x0D);
    sensor_write_register (0x558, 0x06);
    sensor_write_register (0x559, 0xE1);
    sensor_write_register (0x55A, 0x11);
    sensor_write_register (0x560, 0x1E);
    sensor_write_register (0x561, 0x61);
    sensor_write_register (0x562, 0x10);
    sensor_write_register (0x5B0, 0x50);
    sensor_write_register (0x5B2, 0x1A);
    sensor_write_register (0x5B3, 0x04);

    delay_ms(200);
    sensor_write_register (0x200, 0x00); /* standby */
    sensor_write_register (0x202, 0x00); /* master mode start */
    sensor_write_register (0x249, 0x0A); /* XVSOUTSEL XHSOUTSEL */
    
    printf("--IMX290 1080P 30fps LINE Init OK!----\n");
    
}

#if SENSOR_IMX290_LINE_WDR_12BIT

// DOL mode
// INCK = 37.125M
// 1080p 30fps
// LVDS Serial 4ch (445.5Mbps/ch)
// ADC 12bit

void sensor_wdr_1080p30_init()
{
//1080P @30fps DOL 
    sensor_write_register (0x200, 0x01); /* standby */

    delay_ms(200);  

    sensor_write_register (0x209, 0x01);
    sensor_write_register (0x20C, 0x11);
    sensor_write_register (0x20F, 0x00);
    sensor_write_register (0x210, 0x21); 
    sensor_write_register (0x212, 0x64); 
    sensor_write_register (0x214, 0x20); //Gain
    sensor_write_register (0x216, 0x09); 
                                                                    //Vmax default is 0x465

    sensor_write_register (0x21C, 0x98); 
    sensor_write_register (0x21D, 0x08);
    sensor_write_register (0x220, 0x02); 
    sensor_write_register (0x224, 0xC9); 
    sensor_write_register (0x225, 0x07);
    sensor_write_register (0x230, 0x0B);
    //sensor_write_register (0x245, 0x05);  //org
    //sensor_write_register (0x245, 0x0f);        // DOLSYDINFOEN [1] = '1'; HINFOEN [2]= '1'  
    sensor_write_register (0x245, 0x03);        // DOLSYDINFOEN [1] = '1'; HINFOEN [2]= '0'
    sensor_write_register (0x246, 0xE1);
    sensor_write_register (0x25C, 0x18);
    sensor_write_register (0x25E, 0x20);
    sensor_write_register (0x270, 0x02);
    sensor_write_register (0x271, 0x11);
    sensor_write_register (0x29B, 0x10);
    sensor_write_register (0x29C, 0x22);
    sensor_write_register (0x2A2, 0x02);
    sensor_write_register (0x2A6, 0x20);
    sensor_write_register (0x2A8, 0x20);
    sensor_write_register (0x2AA, 0x20);
    sensor_write_register (0x2AC, 0x20);
    sensor_write_register (0x2B0, 0x43);

    sensor_write_register (0x306, 0x1a);        //new added!
    sensor_write_register (0x319, 0x9E); 
    sensor_write_register (0x31C, 0x1E); 
    sensor_write_register (0x31E, 0x08); 
    sensor_write_register (0x328, 0x05); 
    sensor_write_register (0x33D, 0x83);
    sensor_write_register (0x350, 0x03);
    sensor_write_register (0x35E, 0x1A);
    sensor_write_register (0x364, 0x1A);
    sensor_write_register (0x37C, 0x00);
    sensor_write_register (0x37E, 0x00);

    sensor_write_register (0x4B8, 0x50);
    sensor_write_register (0x4B9, 0x10);
    sensor_write_register (0x4BA, 0x00);
    sensor_write_register (0x4BB, 0x04);
    sensor_write_register (0x4C8, 0x50);
    sensor_write_register (0x4C9, 0x10);
    sensor_write_register (0x4CA, 0x00);
    sensor_write_register (0x4CB, 0x04);
    
    sensor_write_register (0x52C, 0xD3);
    sensor_write_register (0x52D, 0x10);
    sensor_write_register (0x52E, 0x0D);
    sensor_write_register (0x558, 0x06);
    sensor_write_register (0x559, 0xE1);
    sensor_write_register (0x55A, 0x11);
    sensor_write_register (0x560, 0x1E);
    sensor_write_register (0x561, 0x61);
    sensor_write_register (0x562, 0x10);
    sensor_write_register (0x5B0, 0x50);
    sensor_write_register (0x5B2, 0x1A);
    sensor_write_register (0x5B3, 0X04);

    delay_ms(200);  
    sensor_write_register (0x200, 0x00); /* standby */
    sensor_write_register (0x202, 0x00); /* master mode start */
    sensor_write_register (0x24b, 0x0A); /* XVSOUTSEL XHSOUTSEL */

    printf("--IMX290 1080P 30fps 12bit DOL WDR Init OK!----\n");
    
}
#else


// DOL mode
// INCK = 37.125M
// 1080p 30fps
// LVDS Serial 4ch (445.5Mbps/ch)
// ADC 10bit

void sensor_wdr_1080p30_init()
{
//1080P @30fps DOL 
    sensor_write_register (0x200, 0x01); /* standby */

    delay_ms(200);  

    sensor_write_register (0x205, 0x00);  // ADbit = 10bit

    sensor_write_register (0x209, 0x01);
    sensor_write_register (0x20A, 0x3C);  // black level 60d
    sensor_write_register (0x20C, 0x11);
    sensor_write_register (0x20F, 0x00);
    sensor_write_register (0x210, 0x21); 
    sensor_write_register (0x212, 0x64); 
    sensor_write_register (0x214, 0x20);  //Gain
    sensor_write_register (0x216, 0x09); 
    
    sensor_write_register (0x218, 0xA6);       
    sensor_write_register (0x219, 0x04); 
    sensor_write_register (0x21a, 0x00);    

    sensor_write_register (0x21C, 0x20);       
    sensor_write_register (0x21D, 0x08);
    
    
    sensor_write_register (0x220, 0x02); 
    sensor_write_register (0x224, 0xC9); 
    sensor_write_register (0x225, 0x07);
    sensor_write_register (0x230, 0x0B);
    sensor_write_register (0x245, 0x03);        // DOLSYDINFOEN [1] = '1'; HINFOEN [2]= '0'
    sensor_write_register (0x246, 0xE0);        // ODbit[1:0] = 10bit
    
    sensor_write_register (0x25C, 0x18);
    sensor_write_register (0x25E, 0x20);
    sensor_write_register (0x270, 0x02);
    sensor_write_register (0x271, 0x11);
    sensor_write_register (0x29B, 0x10);
    sensor_write_register (0x29C, 0x22);
    sensor_write_register (0x2A2, 0x02);
    sensor_write_register (0x2A6, 0x20);
    sensor_write_register (0x2A8, 0x20);
    sensor_write_register (0x2AA, 0x20);
    sensor_write_register (0x2AC, 0x20);
    sensor_write_register (0x2B0, 0x43);

    sensor_write_register (0x306, 0x1a);
    
    sensor_write_register (0x319, 0x9E); 
    sensor_write_register (0x31C, 0x1E); 
    sensor_write_register (0x31E, 0x08); 
    sensor_write_register (0x328, 0x05); 
    sensor_write_register (0x329, 0x1d);  // ADBIT1 = 10bit
    sensor_write_register (0x33D, 0x83);
    sensor_write_register (0x350, 0x03);
    sensor_write_register (0x35E, 0x1A);
    sensor_write_register (0x364, 0x1A);
    sensor_write_register (0x37C, 0x12);    //ADbit2 = 10bit
    sensor_write_register (0x37E, 0x00);
    sensor_write_register (0x3EC, 0x37);    // ADbit3 = 10bit

    sensor_write_register (0x4B8, 0x50);
    sensor_write_register (0x4B9, 0x10);
    sensor_write_register (0x4BA, 0x00);
    sensor_write_register (0x4BB, 0x04);
    sensor_write_register (0x4C8, 0x50);
    sensor_write_register (0x4C9, 0x10);
    sensor_write_register (0x4CA, 0x00);
    sensor_write_register (0x4CB, 0x04);
    
    sensor_write_register (0x52C, 0xD3);
    sensor_write_register (0x52D, 0x10);
    sensor_write_register (0x52E, 0x0D);
    sensor_write_register (0x558, 0x06);
    sensor_write_register (0x559, 0xE1);
    sensor_write_register (0x55A, 0x11);
    sensor_write_register (0x560, 0x1E);
    sensor_write_register (0x561, 0x61);
    sensor_write_register (0x562, 0x10);
    sensor_write_register (0x5B0, 0x50);
    sensor_write_register (0x5B2, 0x1A);
    sensor_write_register (0x5B3, 0X04);

    delay_ms(200);  
    sensor_write_register (0x200, 0x00); /* standby */
    sensor_write_register (0x202, 0x00); /* master mode start */
    sensor_write_register (0x24b, 0x0A); /* XVSOUTSEL XHSOUTSEL */
    
    printf("--IMX290 1080P 30fps 10bit DOL WDR Init OK!----\n");  
}

#endif

void sensor_linear_720p120_init()
{
    /* imx290 720p120 */    
    sensor_write_register (0x200, 0x01); /* standby */

    delay_ms(200); 

    sensor_write_register (0x202, 0x00);
    sensor_write_register (0x205, 0x00);
    sensor_write_register (0x207, 0x10);  
    sensor_write_register (0x209, 0x00);
    sensor_write_register (0x20A, 0x3C);
    sensor_write_register (0x20F, 0x00);
    sensor_write_register (0x210, 0x21);
    sensor_write_register (0x212, 0x64);
    sensor_write_register (0x214, 0x20);//Gain
    sensor_write_register (0x216, 0x09);

    sensor_write_register (0x218, 0xEE); /* VMAX[7:0] */
    sensor_write_register (0x219, 0x02); /* VMAX[15:8] */
    sensor_write_register (0x21A, 0x00); /* VMAX[16] */
    sensor_write_register (0x21B, 0x00); 
    sensor_write_register (0x21C, 0x72); /* HMAX[7:0] */
    sensor_write_register (0x21D, 0x06); /* HMAX[15:8] */

    
    sensor_write_register (0x220, 0x01);//SHS1
    sensor_write_register (0x221, 0x00);

    sensor_write_register (0x246, 0xE0);
    sensor_write_register (0x24B, 0x0A);
    sensor_write_register (0x25C, 0x20);
    sensor_write_register (0x25D, 0x00);
    sensor_write_register (0x25E, 0x20);
    sensor_write_register (0x25F, 0x01);
    
    sensor_write_register (0x270, 0x02);
    sensor_write_register (0x271, 0x11);
    sensor_write_register (0x29B, 0x10);
    sensor_write_register (0x29C, 0x22);
    sensor_write_register (0x2A2, 0x02);
    sensor_write_register (0x2A6, 0x20);
    sensor_write_register (0x2A8, 0x20);
    sensor_write_register (0x2AA, 0x20);
    sensor_write_register (0x2AC, 0x20);
    sensor_write_register (0x2B0, 0x43);

    sensor_write_register (0x306, 0x00);
    
    sensor_write_register (0x319, 0x9E);
    sensor_write_register (0x31C, 0x1E);
    sensor_write_register (0x31E, 0x08);
    sensor_write_register (0x328, 0x05);
    sensor_write_register (0x329, 0x1D);
    sensor_write_register (0x33D, 0x83);
    sensor_write_register (0x350, 0x03);

    sensor_write_register (0x35E, 0x1A);// 1A:37.125MHz 1B:74.25MHz
    sensor_write_register (0x364, 0x1A);// 1A:37.125MHz 1B:74.25MHz

    sensor_write_register (0x37C, 0x12);
    sensor_write_register (0x37E, 0x00);
    sensor_write_register (0x3EC, 0x37);

    sensor_write_register (0x4B8, 0x50);
    sensor_write_register (0x4B9, 0x10);
    sensor_write_register (0x4BA, 0x00);
    sensor_write_register (0x4BB, 0x04);
    sensor_write_register (0x4C8, 0x50);
    sensor_write_register (0x4C9, 0x10);
    sensor_write_register (0x4CA, 0x00);
    sensor_write_register (0x4CB, 0x04);

    sensor_write_register (0x52C, 0xD3);
    sensor_write_register (0x52D, 0x10);
    sensor_write_register (0x52E, 0x0D);
    sensor_write_register (0x558, 0x06);
    sensor_write_register (0x559, 0xE1);
    sensor_write_register (0x55A, 0x11);
    sensor_write_register (0x560, 0x1E);
    sensor_write_register (0x561, 0x61);
    sensor_write_register (0x562, 0x10);
    sensor_write_register (0x5B0, 0x50);
    sensor_write_register (0x5B2, 0x1A);
    sensor_write_register (0x5B3, 0x04);

    sensor_write_register (0x618, 0xD9);
    sensor_write_register (0x619, 0x02);
    sensor_write_register (0x680, 0x49);

    delay_ms(200); 
    
    sensor_write_register (0x200, 0x00); /* standby */
    sensor_write_register (0x202, 0x00); /* master mode start */
    sensor_write_register (0x249, 0x0A); /* XVSOUTSEL XHSOUTSEL */
    
    printf("--IMX290 720P 120fps LINE Init OK!----\n");
}

void sensor_wdr_720p60_init()
{
    /* imx290 720p60 dol */    
    sensor_write_register (0x200, 0x01); /* standby */

    delay_ms(200); 

    sensor_write_register (0x202, 0x00);
    sensor_write_register (0x205, 0x00);
    sensor_write_register (0x207, 0x10);  
    sensor_write_register (0x209, 0x00);
    sensor_write_register (0x20A, 0x3C);
    sensor_write_register (0x20C, 0x11);
    sensor_write_register (0x20F, 0x00);
    sensor_write_register (0x210, 0x21);
    sensor_write_register (0x212, 0x64);
    sensor_write_register (0x214, 0x20);//Gain
    sensor_write_register (0x216, 0x09);

    sensor_write_register (0x218, 0xEE); /* VMAX[7:0] */
    sensor_write_register (0x219, 0x02); /* VMAX[15:8] */
    sensor_write_register (0x21A, 0x00); /* VMAX[16] */
    sensor_write_register (0x21B, 0x00); 
    sensor_write_register (0x21C, 0x72); /* HMAX[7:0] */
    sensor_write_register (0x21D, 0x06); /* HMAX[15:8] */

    
    sensor_write_register (0x220, 0x01);//SHS1
    sensor_write_register (0x221, 0x00);

    sensor_write_register (0x224, 0x11);//SHS2
    sensor_write_register (0x225, 0x00);
    sensor_write_register (0x226, 0x00);

    sensor_write_register (0x230, 0x09);//RHS1
    sensor_write_register (0x231, 0x00);
    sensor_write_register (0x232, 0x00);

    sensor_write_register (0x245, 0x03);
    sensor_write_register (0x246, 0xE0);
    sensor_write_register (0x24B, 0x0A);
    sensor_write_register (0x25C, 0x20);
    sensor_write_register (0x25D, 0x00);
    sensor_write_register (0x25E, 0x20);
    sensor_write_register (0x25F, 0x01);
    
    sensor_write_register (0x270, 0x02);
    sensor_write_register (0x271, 0x11);
    sensor_write_register (0x29B, 0x10);
    sensor_write_register (0x29C, 0x22);
    sensor_write_register (0x2A2, 0x02);
    sensor_write_register (0x2A6, 0x20);
    sensor_write_register (0x2A8, 0x20);
    sensor_write_register (0x2AA, 0x20);
    sensor_write_register (0x2AC, 0x20);
    sensor_write_register (0x2B0, 0x43);

    sensor_write_register (0x306, 0x1a);
    
    sensor_write_register (0x319, 0x9E);
    sensor_write_register (0x31C, 0x1E);
    sensor_write_register (0x31E, 0x08);
    sensor_write_register (0x328, 0x05);
    sensor_write_register (0x329, 0x1D);
    sensor_write_register (0x33D, 0x83);
    sensor_write_register (0x350, 0x03);

    sensor_write_register (0x35E, 0x1A);// 1A:37.125MHz 1B:74.25MHz
    sensor_write_register (0x364, 0x1A);// 1A:37.125MHz 1B:74.25MHz

    sensor_write_register (0x37C, 0x12);
    sensor_write_register (0x37E, 0x00);
    sensor_write_register (0x3EC, 0x37);

    sensor_write_register (0x4B8, 0x50);
    sensor_write_register (0x4B9, 0x10);
    sensor_write_register (0x4BA, 0x00);
    sensor_write_register (0x4BB, 0x04);
    sensor_write_register (0x4C8, 0x50);
    sensor_write_register (0x4C9, 0x10);
    sensor_write_register (0x4CA, 0x00);
    sensor_write_register (0x4CB, 0x04);

    sensor_write_register (0x52C, 0xD3);
    sensor_write_register (0x52D, 0x10);
    sensor_write_register (0x52E, 0x0D);
    sensor_write_register (0x558, 0x06);
    sensor_write_register (0x559, 0xE1);
    sensor_write_register (0x55A, 0x11);
    sensor_write_register (0x560, 0x1E);
    sensor_write_register (0x561, 0x61);
    sensor_write_register (0x562, 0x10);
    sensor_write_register (0x5B0, 0x50);
    sensor_write_register (0x5B2, 0x1A);
    sensor_write_register (0x5B3, 0x04);

    sensor_write_register (0x618, 0xC6);
    sensor_write_register (0x619, 0x05);
    sensor_write_register (0x680, 0x49);

    delay_ms(200); 
    
    sensor_write_register (0x200, 0x00); /* standby */
    sensor_write_register (0x202, 0x00); /* master mode start */
    sensor_write_register (0x249, 0x0A); /* XVSOUTSEL XHSOUTSEL */
    
    printf("--IMX290 720P 60fps DOL WDR Init OK!----\n");
}



