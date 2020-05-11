#include <stdio.h>
#include <string.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <sys/ioctl.h>
#include <fcntl.h>
#include <unistd.h>
#include "hi_comm_video.h"
#include <hi_math.h>
#include <time.h>

#include "hi_spi.h"


extern WDR_MODE_E genSensorMode;
extern HI_U8 gu8SensorImageMode;
extern HI_BOOL bSensorInit;

static int g_fd= -1;

//int sensor_spi_init(void)
int imx323_spi_spi_init(void)
{
    if (g_fd >= 0)
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

//int sensor_spi_exit(void)
int imx323_spi_spi_exit(void)
{
    if (g_fd >= 0)
    {
        close(g_fd);
        g_fd = -1;
        return 0;
    }
    return -1;
}

//int sensor_write_register(unsigned int addr, unsigned char data)
int imx323_spi_write_register(unsigned int addr, unsigned char data)
{
    int ret;
    struct spi_ioc_transfer mesg[1];
    unsigned char  tx_buf[8] = {0};
    unsigned char  rx_buf[8] = {0};

    tx_buf[0] = (addr & 0x0f00) >> 8;
    tx_buf[1] = addr & 0xff;
    tx_buf[2] = data;

    memset(mesg, 0, sizeof(mesg));  
#ifndef __HuaweiLite__
    mesg[0].tx_buf = (__u32)tx_buf;  
    mesg[0].len    = 3;  
    mesg[0].rx_buf = (__u32)rx_buf; 
    mesg[0].cs_change = 1;
#else
    mesg[0].tx_buf = (char*)tx_buf;  
    mesg[0].len    = 3;  
    mesg[0].rx_buf = (char*)rx_buf; 
    mesg[0].cs_change = 1;
#endif
    ret = ioctl(g_fd, SPI_IOC_MESSAGE(1), mesg);
    if (ret < 0) {  
        printf("SPI_IOC_MESSAGE error \n");  
        return -1;  
    }
    return 0;
}

//int sensor_read_register(unsigned int addr)
int imx323_spi_read_register(unsigned int addr)
{
    int ret = 0;
    struct spi_ioc_transfer mesg[1];
    unsigned char  tx_buf[8] = {0};
    unsigned char  rx_buf[8] = {0};
    
    tx_buf[0] = (addr & 0xff00) >> 8;
    tx_buf[0] |= 0x80;
    tx_buf[1] = addr & 0xff;
    tx_buf[2] = 0;
#ifndef __HuaweiLite__
    memset(mesg, 0, sizeof(mesg));
    mesg[0].tx_buf = (__u32)tx_buf;
    mesg[0].len    = 3;
    mesg[0].rx_buf = (__u32)rx_buf;
    mesg[0].cs_change = 1;
#else
    memset(mesg, 0, sizeof(mesg));
    mesg[0].tx_buf = (char*)tx_buf;
    mesg[0].len    = 3;
    mesg[0].rx_buf = (char*)rx_buf;
    mesg[0].cs_change = 1;
#endif
    ret = ioctl(g_fd, SPI_IOC_MESSAGE(1), mesg);
    if (ret  < 0) {  
        printf("SPI_IOC_MESSAGE error \n");  
        return -1;  
    }
    
    return rx_buf[2];
}

static void delay_ms(int ms) { 
     hi_usleep(ms*1000);
}

//void sensor_prog(int* rom) 
void imx323_spi_prog(int* rom) 
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
            imx323_spi_write_register(addr, data);//sensor_write_register(addr, data);
        }
    }
}

//void sensor_linear_1080p30_RAW12_init(HI_VOID);
//void sensor_linear_720p30_RAW12_init(HI_VOID);
//void sensor_linear_720p60_RAW10_init(HI_VOID);

void imx323_spi_linear_1080p30_RAW12_init(HI_VOID);
void imx323_spi_linear_720p30_RAW12_init(HI_VOID);
void imx323_spi_linear_720p60_RAW10_init(HI_VOID);


//void sensor_init(HI_VOID)
void imx323_spi_init(HI_VOID)
{
    bSensorInit = HI_TRUE;
    /* 1. sensor spi init */
    imx323_spi_spi_init();//sensor_spi_init();

    switch (gu8SensorImageMode)
    {        
	case 0: // 1080P30
                 imx323_spi_linear_1080p30_RAW12_init();//sensor_linear_1080p30_RAW12_init();
            break;
	case 1: // 720P30
                 imx323_spi_linear_720p30_RAW12_init();//sensor_linear_720p30_RAW12_init();
            break;
	case 2: // 720P30
                 imx323_spi_linear_720p60_RAW10_init();//sensor_linear_720p60_RAW10_init();
            break;
      default:
            printf("Not support this mode\n");
            bSensorInit = HI_FALSE;
    }


}

//void sensor_exit(HI_VOID)
void imx323_spi_exit(HI_VOID)
{
    imx323_spi_spi_exit();

    return;
}
//HD 1080p mode;
//37.125MHz
//30fps
//RAW12
//void sensor_linear_1080p30_RAW12_init(HI_VOID)
void imx323_spi_linear_1080p30_RAW12_init(HI_VOID)
{
	imx323_spi_write_register(0x0200, 0x31);
	imx323_spi_write_register(0x0202, 0x0F);
	imx323_spi_write_register(0x0203, 0x4C);
	imx323_spi_write_register(0x0204, 0x04);
	imx323_spi_write_register(0x0205, 0x65);
	imx323_spi_write_register(0x0206, 0x04);
	imx323_spi_write_register(0x0212, 0x82);
	imx323_spi_write_register(0x0216, 0x3C);
	imx323_spi_write_register(0x021F, 0x73);
	imx323_spi_write_register(0x0220, 0xF0);
	imx323_spi_write_register(0x0227, 0x20);
	imx323_spi_write_register(0x022C, 0x00);
	imx323_spi_write_register(0x023F, 0x0A);
	imx323_spi_write_register(0x027A, 0x00);
	imx323_spi_write_register(0x027B, 0x00);
	imx323_spi_write_register(0x029A, 0x26);
	imx323_spi_write_register(0x029B, 0x02);
	imx323_spi_write_register(0x0317, 0x0D);
	imx323_spi_write_register(0x0200, 0x30);
	printf("-------Sony IMX323 Sensor 1080p_30fps_raw12_cmos_37p125Mhz Initial OK!-------\n");
}

//HD 720p mode
//37.125MHz
//30fps
//RAW12
//void sensor_linear_720p30_RAW12_init(HI_VOID)
void imx323_spi_linear_720p30_RAW12_init(HI_VOID)
{
	imx323_spi_write_register(0x200, 0x31);
	imx323_spi_write_register(0x202, 0x01);
	imx323_spi_write_register(0x203, 0x72);
	imx323_spi_write_register(0x204, 0x06);
	imx323_spi_write_register(0x205, 0xEE);
	imx323_spi_write_register(0x206, 0x02);
	imx323_spi_write_register(0x211, 0x01);
	imx323_spi_write_register(0x212, 0x82);
	imx323_spi_write_register(0x216, 0xF0);
	imx323_spi_write_register(0x21F, 0x73);
	imx323_spi_write_register(0x220, 0xF0);
	imx323_spi_write_register(0x222, 0xC0);
	imx323_spi_write_register(0x227, 0x20);
	imx323_spi_write_register(0x22C, 0x00);
	imx323_spi_write_register(0x23F, 0x0A);
	imx323_spi_write_register(0x2CE, 0x40);
	imx323_spi_write_register(0x2CF, 0x81);
	imx323_spi_write_register(0x2D0, 0x01);
	imx323_spi_write_register(0x317, 0x0D);
	imx323_spi_write_register(0x200, 0x30);
	printf("-------Sony IMX323 Sensor 720p_30fps_raw12_cmos_37p125Mhz Initial OK!-------\n");
}

//HD 720p mode
//37.125MHz
//60fps
//RAW10
//void sensor_linear_720p60_RAW10_init(HI_VOID)
void imx323_spi_linear_720p60_RAW10_init(HI_VOID)
{
	imx323_spi_write_register(0x200, 0x31);
	imx323_spi_write_register(0x202, 0x01);
	imx323_spi_write_register(0x203, 0x39);
	imx323_spi_write_register(0x204, 0x03);
	imx323_spi_write_register(0x205, 0xEE);
	imx323_spi_write_register(0x206, 0x02);
	imx323_spi_write_register(0x216, 0xF0);
	imx323_spi_write_register(0x21F, 0x73);
	imx323_spi_write_register(0x222, 0xC0);
	imx323_spi_write_register(0x227, 0x20);
	imx323_spi_write_register(0x22C, 0x00);
	imx323_spi_write_register(0x23F, 0x0A);
	imx323_spi_write_register(0x2CE, 0x00);
	imx323_spi_write_register(0x2CF, 0x00);
	imx323_spi_write_register(0x2D0, 0x00);
	imx323_spi_write_register(0x317, 0x0D);
	imx323_spi_write_register(0x200, 0x30);
	printf("-------Sony IMX323 Sensor 720p_60fps_raw10_cmos_37p125Mhz Initial OK!-------\n");
}
