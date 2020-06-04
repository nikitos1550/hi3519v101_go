#include <stdio.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <sys/ioctl.h>
#include <fcntl.h>
#include <unistd.h>

#include "hi_comm_video.h"

#ifdef HI_GPIO_I2C
#include "gpioi2c_ex.h"
#else
#include "hi_i2c.h"
#endif

const unsigned int h65_i2c_addr	=	0x60;		/* I2C Address of soih65 */
const unsigned int h65_addr_byte	=	1;
const unsigned int h65_data_byte	=	1;
static int g_fd = -1;
static int flag_init = 0;

extern WDR_MODE_E genSensorMode_h65;
extern HI_U8 gu8SensorImageMode_h65;
extern HI_BOOL bSensorInit_h65;

int h65_i2c_init(void)
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

    ret = ioctl(g_fd, I2C_SLAVE_FORCE, h65_i2c_addr);
    if (ret < 0)
    {
        printf("CMD_SET_DEV error!\n");
        return ret;
    }
#endif

    return 0;
}

int h65_i2c_exit(void)
{
    if (g_fd >= 0)
    {
        close(g_fd);
        g_fd = -1;
        return 0;
    }
    return -1;
}

int h65_read_register(int addr)
{
	// TODO: 
	
	return 0;
}

int h65_write_register(int addr, int data)
{
#ifdef HI_GPIO_I2C
    i2c_data.dev_addr = h65_i2c_addr;
    i2c_data.reg_addr = addr;
    i2c_data.addr_byte_num = h65_addr_byte;
    i2c_data.data = data;
    i2c_data.data_byte_num = h65_data_byte;

    ret = ioctl(g_fd, GPIO_I2C_WRITE, &i2c_data);

    if (ret)
    {
        printf("GPIO-I2C write faild!\n");
        return ret;
    }
#else
    if(flag_init == 0)
    {
    
	h65_i2c_init();
	flag_init = 1;
    }

    int idx = 0;
    int ret;
    char buf[8];

    buf[idx++] = addr & 0xFF;
    if (h65_addr_byte == 2)
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
    if (h65_data_byte == 2)
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
    usleep(ms*1000);
}


void h65_linear_960p30_init();

void h65_init()
{
    h65_i2c_init();

    h65_linear_960p30_init();

    return ;
}

void h65_exit()
{
    h65_i2c_exit();
	flag_init = 0;
    return;
}

/* 960P30*/
void h65_linear_960p30_init()
{

   // version    24MHz  MCLK  30FPS  PC:2193287995
h65_write_register(0x12,0x40);
h65_write_register(0x0E,0x11);
h65_write_register(0x0F,0x04);  
h65_write_register(0x10,0x24); 
h65_write_register(0x11,0x80);  
h65_write_register(0x5F,0x01);
h65_write_register(0x60,0x10);  
h65_write_register(0x19,0x64);  
h65_write_register(0x48,0x25);  
h65_write_register(0x20,0xD0);
h65_write_register(0x21,0x02);  
h65_write_register(0x22,0xE8);  
h65_write_register(0x23,0x03);  
h65_write_register(0x24,0x80);  
h65_write_register(0x25,0xC0);
h65_write_register(0x26,0x32);
h65_write_register(0x27,0x5C);  
h65_write_register(0x28,0x1C); 
h65_write_register(0x29,0x01);  
h65_write_register(0x2A,0x48);
h65_write_register(0x2B,0x25);  
h65_write_register(0x2C,0x00);
h65_write_register(0x2D,0x00);
h65_write_register(0x2E,0xF8);  
h65_write_register(0x2F,0x40);  
h65_write_register(0x41,0x90);  
h65_write_register(0x42,0x12);  
h65_write_register(0x39,0x90);
h65_write_register(0x1D,0xFF);
h65_write_register(0x1E,0x1F);
h65_write_register(0x6C,0x80);
h65_write_register(0x1F,0x10);
h65_write_register(0x31,0x0C);  
h65_write_register(0x32,0x20);
h65_write_register(0x33,0x0C);
h65_write_register(0x34,0x4F);
h65_write_register(0x36,0x06);
h65_write_register(0x38,0x39);
h65_write_register(0x3A,0x08);
h65_write_register(0x3B,0x50);
h65_write_register(0x3C,0xA0);
h65_write_register(0x3D,0x00);  
h65_write_register(0x3E,0x01);  
h65_write_register(0x3F,0x00);
h65_write_register(0x40,0x00);
h65_write_register(0x0D,0x50);
h65_write_register(0x5A,0x43);
h65_write_register(0x5B,0xB3);
h65_write_register(0x5C,0x0C);  
h65_write_register(0x5D,0x7E);  
h65_write_register(0x5E,0x24);  
h65_write_register(0x62,0x40);
h65_write_register(0x67,0x48);  
h65_write_register(0x6A,0x11);
h65_write_register(0x68,0x04);  
h65_write_register(0x8F,0x9F);
h65_write_register(0x0C,0x00);
h65_write_register(0x59,0x97);
h65_write_register(0x4A,0x05);   
h65_write_register(0x50,0x03);
h65_write_register(0x47,0x62);
h65_write_register(0x7E,0xCD);
h65_write_register(0x8D,0x87);
h65_write_register(0x49,0x10);
h65_write_register(0x7F,0x52);
h65_write_register(0x8E,0x00);
h65_write_register(0x8C,0xFF);
h65_write_register(0x8B,0x01);
h65_write_register(0x57,0x02);
h65_write_register(0x94,0x00);
h65_write_register(0x95,0x00);
h65_write_register(0x63,0x80);
h65_write_register(0x7B,0x46);
h65_write_register(0x7C,0x2D);
h65_write_register(0x90,0x00);
h65_write_register(0x79,0x00);
h65_write_register(0x13,0x81);
h65_write_register(0x12,0x00);
h65_write_register(0x45,0x89);
h65_write_register(0x93,0x68);
delay_ms(500);                  
h65_write_register(0x45,0x19);
h65_write_register(0x1F,0x11);


                                                                                       
                                                                                       
    bSensorInit_h65 = HI_TRUE;                                                             
    printf("=========================================================\n");             
    printf("=== soih65 sensor 960P30fps(Parallel port) init success!=====\n");         
    printf("=========================================================\n");             
                                                                                       
    return;                                                                            
}
