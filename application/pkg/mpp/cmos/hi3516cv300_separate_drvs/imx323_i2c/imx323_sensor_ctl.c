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

#ifdef HI_GPIO_I2C
#include "gpioi2c_ex.h"
#else
#include "hi_i2c.h"
#endif

const unsigned char sensor_i2c_addr     =    0x34;        /* I2C Address of IMX323 */
const unsigned int  sensor_addr_byte    =    2;
const unsigned int  sensor_data_byte    =    1;


extern WDR_MODE_E genSensorMode;
extern HI_U8 gu8SensorImageMode;
extern HI_BOOL bSensorInit;

static int g_fd= -1;

int sensor_i2c_init(void)
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
    
    ret = ioctl(g_fd, I2C_SLAVE_FORCE, (sensor_i2c_addr>>1));
    if (ret < 0)
    {
        printf("CMD_SET_DEV error!\n");
        return ret;
    } 
#endif
    
    return 0;
}

int sensor_i2c_exit(void)
{
    if (g_fd >= 0)
    {
        close(g_fd);
        g_fd = -1;
        return 0;
    }
    return -1;
}

int sensor_write_register(int addr, int data)
{
#ifdef HI_GPIO_I2C
    i2c_data.dev_addr = sensor_i2c_addr;
    i2c_data.reg_addr = addr;
    i2c_data.addr_byte_num = sensor_addr_byte;
    i2c_data.data = data;
    i2c_data.data_byte_num = sensor_data_byte;
    
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
    
    if (sensor_addr_byte == 2) {
                buf[idx] = (addr >> 8) & 0xff;
                idx++;
                buf[idx] = addr & 0xff;
                idx++;
        } else {
                buf[idx] = addr & 0xff;
                idx++;
        }
        
        if (sensor_data_byte == 2) {
                buf[idx] = (data >> 8) & 0xff;
                idx++;
                buf[idx] = data & 0xff;
                idx++;
        } else {
                buf[idx] = data & 0xff;
                idx++;
        }
    
    ret = write(g_fd, buf, (sensor_addr_byte + sensor_data_byte));
    if(ret < 0)
    {
        printf("I2C_WRITE error!\n");
        return -1;
    } 
#endif
    //printf("sensor_write_register addr:0x%X data:0x%X\n", addr, data);
    return 0;
}


int TODO_sensor_read_register(unsigned int addr)
{
    //TODO
}

static void delay_ms(int ms) { 
     hi_usleep(ms*1000);
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

void sensor_linear_1080p30_RAW12_init(HI_VOID);
void sensor_linear_720p30_RAW12_init(HI_VOID);
void sensor_linear_720p60_RAW10_init(HI_VOID);

void sensor_init(HI_VOID)
{
    bSensorInit = HI_TRUE;
    /* 1. sensor spi init */
    sensor_i2c_init();

    switch (gu8SensorImageMode)
    {        
	case 0: // 1080P30
                 sensor_linear_1080p30_RAW12_init();
            break;
	case 1: // 720P30
                 //sensor_linear_720p30_RAW12_init();
            break;
	case 2: // 720P30
                 //sensor_linear_720p60_RAW10_init();
            break;
      default:
            printf("Not support this mode\n");
            bSensorInit = HI_FALSE;
    }


}

void sensor_exit(HI_VOID)
{
    sensor_i2c_exit();
    return;
}

//HD 1080p mode;
//37.125MHz
//30fps
//RAW12
void sensor_linear_1080p30_RAW12_init(HI_VOID)
{
        sensor_write_register(0x0100, 0x00);//sensor_write_register(0x3000, 0x31);  //STANDBY
        sensor_write_register(0x3002, 0x0F);                                        //MODE 1080p
        sensor_write_register(0x0342, 0x04);                                        //HMAX MSB
        sensor_write_register(0x0343, 0x4C);//sensor_write_register(0x3004, 0x04);  //HMAX LSB
        sensor_write_register(0x0340, 0x04);//sensor_write_register(0x3005, 0x65);  //VMAX MSB
        sensor_write_register(0x0341, 0x65);//sensor_write_register(0x3006, 0x04);  //VMAX LSB
        sensor_write_register(0x0112, 0x0C);//sensor_write_register(0x3012, 0x82);  //AD gradation setting: 12bit
        sensor_write_register(0x0113, 0x0C);                                        // ---//---
        sensor_write_register(0x3016, 0x3C);                                        //HD1080p
        sensor_write_register(0x301F, 0x73);                                        //magic
        sensor_write_register(0x0008, 0x01);//0x01//sensor_write_register(0x3020, 0xF0);  //BLKLEVEL [0]
        sensor_write_register(0x0009, 0x70);//0x70//sensor_write_register(0x3020, 0xF0);  // ---//--- [0:7]
        sensor_write_register(0x3027, 0x20);                                        //magic
        sensor_write_register(0x302C, 0x00);                                        //XMSTA
        sensor_write_register(0x303F, 0x0A);                                        //magic
        sensor_write_register(0x307A, 0x00);                                        //10BITC Setting registers for 10 bit
        sensor_write_register(0x307B, 0x00);                                        // ---//---
        sensor_write_register(0x309A, 0x26);                                        //12B1080 P [11:0]
        sensor_write_register(0x309B, 0x02);                                        // ---//---
        sensor_write_register(0x3117, 0x0D);                                        //magic
        sensor_write_register(0x0100, 0x01);//sensor_write_register(0x3000, 0x30);
        printf("-------Sony IMX323 Sensor 1080p_30fps_raw12_cmos_37p125Mhz Initial OK!-------\n");
}


//HD 720p mode
//37.125MHz
//30fps
//RAW12
void TODO_sensor_linear_720p30_RAW12_init(HI_VOID)
{
	sensor_write_register(0x200, 0x31);
	sensor_write_register(0x202, 0x01);
	sensor_write_register(0x203, 0x72);
	sensor_write_register(0x204, 0x06);
	sensor_write_register(0x205, 0xEE);
	sensor_write_register(0x206, 0x02);
	sensor_write_register(0x211, 0x01);
	sensor_write_register(0x212, 0x82);
	sensor_write_register(0x216, 0xF0);
	sensor_write_register(0x21F, 0x73);
	sensor_write_register(0x220, 0xF0);
	sensor_write_register(0x222, 0xC0);
	sensor_write_register(0x227, 0x20);
	sensor_write_register(0x22C, 0x00);
	sensor_write_register(0x23F, 0x0A);
	sensor_write_register(0x2CE, 0x40);
	sensor_write_register(0x2CF, 0x81);
	sensor_write_register(0x2D0, 0x01);
	sensor_write_register(0x317, 0x0D);
	sensor_write_register(0x200, 0x30);
	printf("-------Sony IMX323 Sensor 720p_30fps_raw12_cmos_37p125Mhz Initial OK!-------\n");
}

//HD 720p mode
//37.125MHz
//60fps
//RAW10
void TODO_sensor_linear_720p60_RAW10_init(HI_VOID)
{
	sensor_write_register(0x200, 0x31);
	sensor_write_register(0x202, 0x01);
	sensor_write_register(0x203, 0x39);
	sensor_write_register(0x204, 0x03);
	sensor_write_register(0x205, 0xEE);
	sensor_write_register(0x206, 0x02);
	sensor_write_register(0x216, 0xF0);
	sensor_write_register(0x21F, 0x73);
	sensor_write_register(0x222, 0xC0);
	sensor_write_register(0x227, 0x20);
	sensor_write_register(0x22C, 0x00);
	sensor_write_register(0x23F, 0x0A);
	sensor_write_register(0x2CE, 0x00);
	sensor_write_register(0x2CF, 0x00);
	sensor_write_register(0x2D0, 0x00);
	sensor_write_register(0x317, 0x0D);
	sensor_write_register(0x200, 0x30);
	printf("-------Sony IMX323 Sensor 720p_60fps_raw10_cmos_37p125Mhz Initial OK!-------\n");
}
