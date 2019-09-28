#include <stdio.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <sys/ioctl.h>
#include <fcntl.h>
#include <unistd.h>

#include "hi_comm_video.h"
#include "hi_sns_ctrl.h"


#include "hi_i2c.h"

void imx265_linear_1080p25_init_roi(ISP_DEV IspDev);
void imx265_linear_1080p25_init_orig(ISP_DEV IspDev);

const unsigned char imx265_i2c_addr     =    0x20;//0x10;        /* I2C Address of IMX265 0x10 or 0x1A */
const unsigned int  imx265_addr_byte    =    2;
const unsigned int  imx265_data_byte    =    1;
static int g_fd[ISP_MAX_DEV_NUM] = {-1, -1};        //WTF?

extern ISP_SNS_STATE_S        g_astimx265[ISP_MAX_DEV_NUM];
extern ISP_SNS_COMMBUS_U      g_aunImx265BusInfo[];


//sensor fps mode
#define IMX265_SENSOR_1080P_25FPS_LINEAR_MODE      (1)

int imx265_i2c_init(ISP_DEV IspDev) {

	printf("imx265_i2c_init IspDev = %d\n", IspDev);

    char acDevFile[16] = {0};
    HI_U8 u8DevNum;

    if(g_fd[IspDev] >= 0) {
        return 0;
    }

    int ret;

    u8DevNum = g_aunImx265BusInfo[IspDev].s8I2cDev;
    snprintf_s(acDevFile, sizeof(acDevFile), sizeof(acDevFile)-1, "/dev/i2c-%d", u8DevNum);

	printf("imx265_i2c_init opens %s \n", acDevFile);

    g_fd[IspDev] = open(acDevFile, O_RDWR);
    if(g_fd[IspDev] < 0) {
        printf("Open /dev/i2c-%d error!\n", IspDev);
        return -1;
    }

	printf("%s opened OK\n", acDevFile);

	printf("Setting i2c address %x\n", (imx265_i2c_addr>>1));
    ret = ioctl(g_fd[IspDev], I2C_SLAVE_FORCE, (imx265_i2c_addr>>1));
    if (ret < 0) {
        printf("CMD_SET_DEV error!\n");
        return ret;
    }

	printf("I2C init done\n");
	//sleep(1);
    return 0;
}

int imx265_i2c_exit(ISP_DEV IspDev) {
    if (g_fd[IspDev] >= 0) {
        close(g_fd[IspDev]);
        g_fd[IspDev] = -1;
        return 0;
    }
    return -1;
}

int imx265_read_register(ISP_DEV IspDev,int addr) {
    // TODO
    return 0;
}

int imx265_write_register(ISP_DEV IspDev,int addr, int data) {

//printf("imx265_write_register\n");
//printf("g_fd[%d] = %d\n", IspDev, g_fd[IspDev]);
//sleep(1);
    if (0 > g_fd[IspDev]) {
        return 0;
    }

    int idx = 0;
    int ret;
    char buf[8];

    if (imx265_addr_byte == 2) {
        buf[idx] = (addr >> 8) & 0xff;
        idx++;
        buf[idx] = addr & 0xff;
        idx++;
    } else {
        //buf[idx] = addr & 0xff;
        //idx++;
    }

    if (imx265_data_byte == 2) {
        //buf[idx] = (data >> 8) & 0xff;
        //idx++;
        //buf[idx] = data & 0xff;
        //idx++;
    } else {
        buf[idx] = data & 0xff;
        idx++;
    }

//printf("writing %x %x %x before shift\n", buf[0], buf[1], buf[2]);
//	buf[0] = buf[0] << 1;
//	buf[1] = buf[1] << 1;
//	buf[2] = buf[2] << 1;


	//printf("writing %x %x %x - size %d\n", buf[0], buf[1], buf[2], (imx265_addr_byte + imx265_data_byte));
    ret = write(g_fd[IspDev], buf, imx265_addr_byte + imx265_data_byte);
//sleep(1);
    if(ret < 0) {
        printf("I2C_WRITE error!\n");
        return -1;
    }

	
    return 0;
}


static void delay_ms(int ms) {
    hi_usleep(ms*1000);
}

/*void imx265_prog(ISP_DEV IspDev,int* rom) {
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
            imx265_write_register(IspDev,addr, data);
        }
    }
}*/


void imx265_standby(ISP_DEV IspDev) {
printf("imx265_standby\n");
    // TODO:
    return;
}

void imx265_restart(ISP_DEV IspDev) {
printf("imx265_restart\n");
    // TODO:
    return;
}

void imx265_linear_1080p25_init(ISP_DEV IspDev); //prototype, definitiion later

void imx265_init(ISP_DEV IspDev) {
    HI_U32           i;
    WDR_MODE_E       enWDRMode;
    HI_BOOL          bInit;

    bInit       = g_astimx265[IspDev].bInit;
    enWDRMode   = g_astimx265[IspDev].enWDRMode;

	printf("imx265_init start\n");

		printf("imx265_init -> imx265_i2c_init()\n");

    imx265_i2c_init(IspDev); //protected inside from secondary calls

    /* When sensor first init, config all registers */
    if (HI_FALSE == bInit) {
        //if (enWDRMode != NULL) {
            //TODO
        //} else {
		printf("imx265_init -> imx265_linear_1080p25_init_roi()\n");
		//sleep(1);
            imx265_linear_1080p25_init_roi(IspDev);
            //imx265_linear_1080p25_init_orig(IspDev);
        //}
    } else { /* When sensor switch mode(linear<->WDR or resolution), config different registers(if possible) */
        //???
        //imx265_linear_1080p25_init(IspDev);
        printf("imx265_init voked on inited dev\n");
    }

    /*for (i=0; i<g_astimx265[IspDev].astRegsInfo[0].u32RegNum; i++)
    {
        imx265_write_register(IspDev, g_astimx265[IspDev].astRegsInfo[0].astI2cData[i].u32RegAddr, g_astimx265[IspDev].astRegsInfo[0].astI2cData[i].u32Data);
    }*/

    g_astimx265[IspDev].bInit = HI_TRUE;


	printf("imx265_init finished\n");
sleep(1);
    return ;
}

void imx265_exit(ISP_DEV IspDev) {
printf("imx265_exit\n");
    imx265_i2c_exit(IspDev);
    return;
}

/* 1080P25 ROI */
void imx265_linear_1080p25_init_roi(ISP_DEV IspDev) {

	imx265_write_register(IspDev, 0x3000, 0x01); 
delay_ms(20);

    imx265_write_register(IspDev, 0x3001, 0xD0); //24'hD00102,
    imx265_write_register(IspDev, 0x3002, 0xAA); //24'hAA0202,
    
    imx265_write_register(IspDev, 0x300D, 0x00); //24'h0C0D02, WINMODE
    
    imx265_write_register(IspDev, 0x3010, 0x28); //24'h651002, //VMAX = 24'h628                                                               
    imx265_write_register(IspDev, 0x3011, 0x06); //24'h041102,                                                                                 
    imx265_write_register(IspDev, 0x3012, 0x00); //24'h001202,
	//14
	//15
    
    imx265_write_register(IspDev, 0x3014, 0x4e); //24'h501402, //HMAX = 0x34e                                                                                   
    imx265_write_register(IspDev, 0x3015, 0x03); //24'h0A1502,
	


    imx265_write_register(IspDev, 0x3018, 0x01); //24'h011802,
    imx265_write_register(IspDev, 0x3019, 0x00); //24'h011902, CKSEL
    imx265_write_register(IspDev, 0x3023, 0x00); //24'h002302,
    imx265_write_register(IspDev, 0x3080, 0x62); //---
    	//89*
    	//8A*
    	//8B*
    	//8C*
    	
    imx265_write_register(IspDev, 0x3089, 0x10); //24'h188902, //INCK=74.25 MHz                                                                    
    imx265_write_register(IspDev, 0x308A, 0x00); //24'h008A02,                                                                                 
    imx265_write_register(IspDev, 0x308B, 0x10); //24'h108B02,                                                                                 
    imx265_write_register(IspDev, 0x308C, 0x00); //24'h028C02,
    
    //    24'h188902, //INCK=37.125 MHz                                                                    
    //24'h008A02,                                                                                 
    //24'h108B02,                                                                                 
    //24'h028C02,
    //imx265_write_register(IspDev, 0x3089, 0x18); //24'h188902, //INCK=37.125 MHz                                                              
    //imx265_write_register(IspDev, 0x308A, 0x00); //24'h008A02,                                                                                 
    //imx265_write_register(IspDev, 0x308B, 0x10); //24'h108B02,                                                                                 
    //imx265_write_register(IspDev, 0x308C, 0x02); //24'h028C02,

    
	//8D*
    	//8E*
    	//8F*

    //imx265_write_register(IspDev, 0x309E, 0x06); //24'h069E02,
    //imx265_write_register(IspDev, 0x30A0, 0x02); //24'h02A002,
    
    imx265_write_register(IspDev, 0x30AF, 0x0E); //24'h0AAF02, All pix
    //imx265_write_register(IspDev, 0x30AF, 0x0A); //24'h0AAF02, FullHD

    imx265_write_register(IspDev, 0x3168, 0xD8); //24'hD86803,
    imx265_write_register(IspDev, 0x3169, 0xA0); //24'hA06903,
    imx265_write_register(IspDev, 0x317D, 0xA1); //24'hA17D03,
    imx265_write_register(IspDev, 0x3180, 0x62); //24'h628003,
    imx265_write_register(IspDev, 0x3190, 0x9B); //24'h9B9003,
    imx265_write_register(IspDev, 0x3191, 0xA0); //24'hA09103,
    imx265_write_register(IspDev, 0x31A4, 0x3F); //24'h3FA403,
    imx265_write_register(IspDev, 0x31A5, 0xB1); //24'hB1A503,
    imx265_write_register(IspDev, 0x31E2, 0x00); //24'h00E203,
    imx265_write_register(IspDev, 0x31EA, 0x00); //24'h00EA03,

    imx265_write_register(IspDev, 0x3226, 0x03); //24'h032604,

    imx265_write_register(IspDev, 0x35AA, 0xB3); //24'hB3AA07,
    imx265_write_register(IspDev, 0x35AC, 0x68); //24'h68AC07,
    imx265_write_register(IspDev, 0x371C, 0xB4); //24'hB41C09,
    imx265_write_register(IspDev, 0x371D, 0x00); //24'h001D09,
    imx265_write_register(IspDev, 0x371E, 0xDE); //24'hDE1E09,
    imx265_write_register(IspDev, 0x371F, 0x00); //24'h001F09,
    imx265_write_register(IspDev, 0x3728, 0xB4); //24'hB42809,
    imx265_write_register(IspDev, 0x3729, 0x00); //24'h002909,
    imx265_write_register(IspDev, 0x372A, 0xDE); //24'hDE2A09,
    imx265_write_register(IspDev, 0x372B, 0x00); //24'h002B09,
    imx265_write_register(IspDev, 0x373A, 0x36); //24'h363A09,
    imx265_write_register(IspDev, 0x3746, 0x36); //24'h364609,
    imx265_write_register(IspDev, 0x38E0, 0xEB); //24'hEBE00A,
    imx265_write_register(IspDev, 0x38E1, 0x00); //24'h00E10A,
    imx265_write_register(IspDev, 0x38E2, 0x0D); //24'h0DE20A,
    imx265_write_register(IspDev, 0x38E3, 0x01); //24'h01E30A,
    imx265_write_register(IspDev, 0x39C4, 0xEB); //24'hEBC40B,
    imx265_write_register(IspDev, 0x39C5, 0x00); //24'h00C50B,
    imx265_write_register(IspDev, 0x39C6, 0x0C); //24'h0CC60B,
    imx265_write_register(IspDev, 0x39C7, 0x01); //24'h01C70B,
    imx265_write_register(IspDev, 0x3D02, 0x6E); //24'h6E020F,
    imx265_write_register(IspDev, 0x3D04, 0xE3); //24'hE3040F,
    imx265_write_register(IspDev, 0x3D05, 0x00); //24'h00050F,
    imx265_write_register(IspDev, 0x3D0C, 0x73); //24'h730C0F,
    imx265_write_register(IspDev, 0x3D0E, 0x6E); //24'h6E0E0F,
    imx265_write_register(IspDev, 0x3D10, 0xE8); //24'hE8100F,
    imx265_write_register(IspDev, 0x3D11, 0x00); //24'h00110F,
    imx265_write_register(IspDev, 0x3D12, 0xE3); //24'hE3120F,
    imx265_write_register(IspDev, 0x3D13, 0x00); //24'h00130F,
    imx265_write_register(IspDev, 0x3D14, 0x6B); //24'h6B140F,
    imx265_write_register(IspDev, 0x3D16, 0x1C); //24'h1C160F,
    imx265_write_register(IspDev, 0x3D18, 0x1C); //24'h1C180F,
    imx265_write_register(IspDev, 0x3D1A, 0x6B); //24'h6B1A0F,
    imx265_write_register(IspDev, 0x3D1C, 0x6E); //24'h6E1C0F,
    imx265_write_register(IspDev, 0x3D1E, 0x9A); //24'h9A1E0F,
    imx265_write_register(IspDev, 0x3D20, 0x12); //24'h12200F,
    imx265_write_register(IspDev, 0x3D22, 0x3E); //24'h3E220F,
    imx265_write_register(IspDev, 0x3D28, 0xB4); //24'hB4280F,
    imx265_write_register(IspDev, 0x3D29, 0x00); //24'h00290F,
    imx265_write_register(IspDev, 0x3D2A, 0x66); //24'h662A0F,
    imx265_write_register(IspDev, 0x3D34, 0x69); //24'h69340F,
    imx265_write_register(IspDev, 0x3D36, 0x17); //24'h17360F,
    imx265_write_register(IspDev, 0x3D38, 0x6A); //24'h6A380F,
    imx265_write_register(IspDev, 0x3D3A, 0x18); //24'h183A0F,
    imx265_write_register(IspDev, 0x3D3E, 0xFF); //24'hFF3E0F,
    imx265_write_register(IspDev, 0x3D3F, 0x0F); //24'h0F3F0F,
    imx265_write_register(IspDev, 0x3D46, 0xFF); //24'hFF460F,
    imx265_write_register(IspDev, 0x3D47, 0x0F); //24'h0F470F,
    imx265_write_register(IspDev, 0x3D4E, 0x4C); //24'h4C4E0F,
    imx265_write_register(IspDev, 0x3D50, 0x50); //24'h50500F,
    imx265_write_register(IspDev, 0x3D54, 0x73); //24'h73540F,
    imx265_write_register(IspDev, 0x3D56, 0x6E); //24'h6E560F,
    imx265_write_register(IspDev, 0x3D58, 0xE8); //24'hE8580F,
    imx265_write_register(IspDev, 0x3D59, 0x00); //24'h00590F,
    imx265_write_register(IspDev, 0x3D5A, 0xCF); //24'hCF5A0F,
    imx265_write_register(IspDev, 0x3D5B, 0x00); //24'h005B0F,
    imx265_write_register(IspDev, 0x3D5E, 0x64); //24'h645E0F,
    imx265_write_register(IspDev, 0x3D66, 0x61); //24'h61660F,
    imx265_write_register(IspDev, 0x3D6E, 0x0D); //24'h0D6E0F,
    imx265_write_register(IspDev, 0x3D70, 0xFF); //24'hFF700F,
    imx265_write_register(IspDev, 0x3D71, 0x0F); //24'h0F710F,
    imx265_write_register(IspDev, 0x3D72, 0x00); //24'h00720F,
    imx265_write_register(IspDev, 0x3D73, 0x00); //24'h00730F,
    imx265_write_register(IspDev, 0x3D74, 0x11); //24'h11740F,
    imx265_write_register(IspDev, 0x3D76, 0x6A); //24'h6A760F,
    imx265_write_register(IspDev, 0x3D78, 0x7F); //24'h7F780F,
    imx265_write_register(IspDev, 0x3D7A, 0xB3); //24'hB37A0F,
    imx265_write_register(IspDev, 0x3D7C, 0x29); //24'h297C0F,
    imx265_write_register(IspDev, 0x3D7E, 0x64); //24'h647E0F,
    imx265_write_register(IspDev, 0x3D80, 0xB1); //24'hB1800F,
    imx265_write_register(IspDev, 0x3D82, 0xB3); //24'hB3820F,
    imx265_write_register(IspDev, 0x3D84, 0x62); //24'h62840F,
    imx265_write_register(IspDev, 0x3D86, 0x64); //24'h64860F,
    imx265_write_register(IspDev, 0x3D88, 0xB1); //24'hB1880F,
    imx265_write_register(IspDev, 0x3D8A, 0xB3); //24'hB38A0F,
    imx265_write_register(IspDev, 0x3D8C, 0x62); //24'h628C0F,
    imx265_write_register(IspDev, 0x3D8E, 0x64); //24'h648E0F,
    imx265_write_register(IspDev, 0x3D90, 0x6D); //24'h6D900F,
    imx265_write_register(IspDev, 0x3D92, 0x65); //24'h65920F,
    imx265_write_register(IspDev, 0x3D94, 0x65); //24'h65940F,
    imx265_write_register(IspDev, 0x3D96, 0x6D); //24'h6D960F,
    imx265_write_register(IspDev, 0x3D98, 0x20); //24'h20980F,
    imx265_write_register(IspDev, 0x3D9A, 0x28); //24'h289A0F,
    imx265_write_register(IspDev, 0x3D9C, 0x81); //24'h819C0F,
    imx265_write_register(IspDev, 0x3D9E, 0x89); //24'h899E0F,
    imx265_write_register(IspDev, 0x3D9F, 0x01); //24'h019F0F,
    imx265_write_register(IspDev, 0x3DA0, 0x66); //24'h66A00F,
    imx265_write_register(IspDev, 0x3DA2, 0x7B); //24'h7BA20F,
    imx265_write_register(IspDev, 0x3DA4, 0x21); //24'h21A40F,
    imx265_write_register(IspDev, 0x3DA6, 0x27); //24'h27A60F,
    imx265_write_register(IspDev, 0x3DA8, 0x8B); //24'h8BA80F,
    imx265_write_register(IspDev, 0x3DA9, 0x01); //24'h01A90F,
    imx265_write_register(IspDev, 0x3DAA, 0x95); //24'h95AA0F,
    imx265_write_register(IspDev, 0x3DAB, 0x01); //24'h01AB0F,
    imx265_write_register(IspDev, 0x3DAC, 0x12); //24'h12AC0F,
    imx265_write_register(IspDev, 0x3DAE, 0x1C); //24'h1CAE0F,
    imx265_write_register(IspDev, 0x3DB0, 0x98); //24'h98B00F,
    imx265_write_register(IspDev, 0x3DB1, 0x01); //24'h01B10F,
    imx265_write_register(IspDev, 0x3DB2, 0xA0); //24'hA0B20F,
    imx265_write_register(IspDev, 0x3DB3, 0x01); //24'h01B30F,
    imx265_write_register(IspDev, 0x3DB4, 0x13); //24'h13B40F,
    imx265_write_register(IspDev, 0x3DB6, 0x1D); //24'h1DB60F,
    imx265_write_register(IspDev, 0x3DB8, 0x99); //24'h99B80F,
    imx265_write_register(IspDev, 0x3DB9, 0x01); //24'h01B90F,
    imx265_write_register(IspDev, 0x3DBA, 0xA1); //24'hA1BA0F,
    imx265_write_register(IspDev, 0x3DBB, 0x01); //24'h01BB0F,
    imx265_write_register(IspDev, 0x3DBC, 0x14); //24'h14BC0F,
    imx265_write_register(IspDev, 0x3DBE, 0x1E); //24'h1EBE0F,
    imx265_write_register(IspDev, 0x3DC0, 0x9A); //24'h9AC00F,
    imx265_write_register(IspDev, 0x3DC1, 0x01); //24'h01C10F,
    imx265_write_register(IspDev, 0x3DC2, 0xA2); //24'hA2C20F,
    imx265_write_register(IspDev, 0x3DC3, 0x01); //24'h01C30F,
    imx265_write_register(IspDev, 0x3DC4, 0x64); //24'h64C40F,
    imx265_write_register(IspDev, 0x3DC6, 0x6E); //24'h6EC60F,
    imx265_write_register(IspDev, 0x3DC8, 0x17); //24'h17C80F,
    imx265_write_register(IspDev, 0x3DCA, 0x26); //24'h26CA0F,
    imx265_write_register(IspDev, 0x3DCC, 0x9D); //24'h9DCC0F,
    imx265_write_register(IspDev, 0x3DCD, 0x01); //24'h01CD0F,
    imx265_write_register(IspDev, 0x3DCE, 0xAC); //24'hACCE0F,
    imx265_write_register(IspDev, 0x3DCF, 0x01); //24'h01CF0F,
    imx265_write_register(IspDev, 0x3DD0, 0x65); //24'h65D00F,
    imx265_write_register(IspDev, 0x3DD2, 0x6F); //24'h6FD20F,
    imx265_write_register(IspDev, 0x3DD4, 0x18); //24'h18D40F,
    imx265_write_register(IspDev, 0x3DD6, 0x27); //24'h27D60F,
    imx265_write_register(IspDev, 0x3DD8, 0x9E); //24'h9ED80F,
    imx265_write_register(IspDev, 0x3DD9, 0x01); //24'h01D90F,
    imx265_write_register(IspDev, 0x3DDA, 0xAD); //24'hADDA0F,
    imx265_write_register(IspDev, 0x3DDB, 0x01); //24'h01DB0F,
    imx265_write_register(IspDev, 0x3DDC, 0x66); //24'h66DC0F,
    imx265_write_register(IspDev, 0x3DDE, 0x70); //24'h70DE0F,
    imx265_write_register(IspDev, 0x3DE0, 0x19); //24'h19E00F,
    imx265_write_register(IspDev, 0x3DE2, 0x28); //24'h28E20F,
    imx265_write_register(IspDev, 0x3DE4, 0x9F); //24'h9FE40F,
    imx265_write_register(IspDev, 0x3DE5, 0x01); //24'h01E50F,
    imx265_write_register(IspDev, 0x3DE6, 0xAE); //24'hAEE60F,
    imx265_write_register(IspDev, 0x3DE7, 0x01); //24'h01E70F,
    imx265_write_register(IspDev, 0x3E04, 0x9D); //24'h9D0410,
    imx265_write_register(IspDev, 0x3E06, 0xB0); //24'hB00610,
    imx265_write_register(IspDev, 0x3E07, 0x00); //24'h000710,
    imx265_write_register(IspDev, 0x3E08, 0x6B); //24'h6B0810,
    imx265_write_register(IspDev, 0x3E0A, 0x7E); //24'h7E0A10,
    imx265_write_register(IspDev, 0x3E24, 0xE3); //24'hE32410,
    imx265_write_register(IspDev, 0x3E25, 0x00); //24'h002510,
    imx265_write_register(IspDev, 0x3E26, 0x9A); //24'h9A2610,
    imx265_write_register(IspDev, 0x3E27, 0x01); //24'h012710,
    imx265_write_register(IspDev, 0x3F20, 0x00); //24'h002011,
    imx265_write_register(IspDev, 0x3F21, 0x00); //24'h002111,
    imx265_write_register(IspDev, 0x3F22, 0xFF); //24'hFF2211,
    imx265_write_register(IspDev, 0x3F23, 0x3F); //24'h3F2311,
    imx265_write_register(IspDev, 0x4003, 0x55); //24'h550312,
    imx265_write_register(IspDev, 0x4005, 0xFF); //24'hFF0512,
    imx265_write_register(IspDev, 0x400B, 0x00); //24'h000B12,
    imx265_write_register(IspDev, 0x400C, 0x54); //24'h540C12,
    imx265_write_register(IspDev, 0x400D, 0xB8); //24'hB80D12,
    imx265_write_register(IspDev, 0x400E, 0x48); //24'h480E12,
    imx265_write_register(IspDev, 0x400F, 0xA2); //24'hA20F12,
    imx265_write_register(IspDev, 0x4012, 0x53); //24'h531212,
    imx265_write_register(IspDev, 0x4013, 0x0A); //24'h0A1312,
    imx265_write_register(IspDev, 0x4014, 0x0C); //24'h0C1412,
    imx265_write_register(IspDev, 0x4015, 0x0A); //24'h0A1512,
    imx265_write_register(IspDev, 0x402A, 0x7F); //24'h7F2A12,
    imx265_write_register(IspDev, 0x402C, 0x29); //24'h292C12,
    imx265_write_register(IspDev, 0x4030, 0x73); //24'h733012,
    imx265_write_register(IspDev, 0x4032, 0x8D); //24'h8D3212,
    imx265_write_register(IspDev, 0x4033, 0x01); //24'h013312,
    imx265_write_register(IspDev, 0x4049, 0x02); //24'h024912,
    imx265_write_register(IspDev, 0x4056, 0x18); //24'h185612,
    imx265_write_register(IspDev, 0x408C, 0x9A); //24'h9A8C12,
    imx265_write_register(IspDev, 0x408E, 0xAA); //24'hAA8E12,
    imx265_write_register(IspDev, 0x4090, 0x3E); //24'h3E9012,
    imx265_write_register(IspDev, 0x4092, 0x5F); //24'h5F9212,
    imx265_write_register(IspDev, 0x4094, 0x0A); //24'h0A9412,
    imx265_write_register(IspDev, 0x4096, 0x0A); //24'h0A9612,
    imx265_write_register(IspDev, 0x4098, 0x7F); //24'h7F9812,
    imx265_write_register(IspDev, 0x409A, 0xB3); //24'hB39A12,
    imx265_write_register(IspDev, 0x409C, 0x29); //24'h299C12,
    imx265_write_register(IspDev, 0x409E, 0x64); //24'h649E12,

	/////////////////////
	imx265_write_register(IspDev, 0x309E, 0x08); //24'h069E02,
    imx265_write_register(IspDev, 0x30A0, 0x04); //24'h02A002,
    ////////////
	
	
    imx265_write_register(IspDev, 0x3254, 0x80); //24'h805404 //black level
    imx265_write_register(IspDev, 0x3255, 0x00); //24'h805404
    
    imx265_write_register(IspDev, 0x3022, 0xF0); //24'h012202 //appnote black level auto adjust NO (F0) or YES (01)
    //imx265_write_register(IspDev, 0x3238, 0x06); //24'h063804 // 1D or 06 ???
	
//ROI

        //imx265_write_register(IspDev, 0x3300, 0x03);

        //tmp = CRB.roi_org_h;
        //Sensor_Set(0x0510, LSB(tmp));
        //Sensor_Set(0x0511, MSB(tmp));
        //imx265_write_register(IspDev, 0x3310, 0x0);
        //imx265_write_register(IspDev, 0x3311, 0x0);

        //Sensor_Set(0x0512, LSB(CRB.roi_org_v));
        //Sensor_Set(0x0513, MSB(CRB.roi_org_v));
        //imx265_write_register(IspDev, 0x3312, 0x0);
        //imx265_write_register(IspDev, 0x3313, 0x0);
        
        //tmp = CRB.roi_size_h;
        //Sensor_Set(0x0514, LSB(tmp));
        //Sensor_Set(0x0515, MSB(tmp));
        //unsigned int width = 1920;
        //imx265_write_register(IspDev, 0x3314, width & 0xFF);
        //imx265_write_register(IspDev, 0x3315, (width & 0xFF00) >> 8);

        //Sensor_Set(0x0516, LSB(CRB.roi_size_v+4));
        //Sensor_Set(0x0517, MSB(CRB.roi_size_v+4));
        //unsigned int height = 1080;
        //imx265_write_register(IspDev, 0x3316, height & 0xFF);
        //imx265_write_register(IspDev, 0x3317, (height & 0xFF00) >> 8);

        
        //imx265_write_register(IspDev, 0x3010, (height+32) & 0xFF); // //VMAX                                                         
        //imx265_write_register(IspDev, 0x3011, ((height+32) & 0xFF00) >> 8);                      
        //imx265_write_register(IspDev, 0x3012, ((height+32) & 0xFF0000) >> 16); 
        //unsigned int vmax = 0x465;
        //imx265_write_register(IspDev, 0x3010, (vmax) & 0xFF); // //VMAX                                                         
        //imx265_write_register(IspDev, 0x3011, ((vmax) & 0xFF00) >> 8);                      
        //imx265_write_register(IspDev, 0x3012, ((vmax) & 0xFF0000) >> 16); 
       // unsigned int vmax = 0x628;
       // imx265_write_register(IspDev, 0x3010, (vmax) & 0xFF); // //VMAX                                                         
       // imx265_write_register(IspDev, 0x3011, ((vmax) & 0xFF00) >> 8);                      
        //imx265_write_register(IspDev, 0x3012, ((vmax) & 0xFF0000) >> 16); 

	//test
	/*
    imx265_write_register(IspDev, 0x3254, 0x00); //24'h805404 //black level
    imx265_write_register(IspDev, 0x3255, 0x00); //24'h805404
    imx265_write_register(IspDev, 0x3022, 0xF0); //24'h012202 //appnote black level auto adjust NO (F0) or YES (01)
    //imx265_write_register(IspDev, 0x3238, 0x55); //24'h063804 // 1D or 06 ??? //55 hor color bars, 5d vert color bars
    //imx265_write_register(IspDev, 0x3248, 0x0F); //ColorWidth
    
    imx265_write_register(IspDev, 0x3238, 0x1D); //24'h063804 // 1D or 06 ??? //55 hor color bars, 5d vert color bars
    //imx265_write_register(IspDev, 0x3238, 0x45); //24'h063804 // 1D or 06 ??? //55 hor color bars, 5d vert color bars
    //imx265_write_register(IspDev, 0x3238, 0x45); //24'h063804 // 1D or 06 ??? //55 hor color bars, 5d vert color bars        
    //imx265_write_register(IspDev, 0x3244, 240);
    //imx265_write_register(IspDev, 0x3246, 135);
    
    //imx265_write_register(IspDev, 0x3248, 0x0F); //24'h063804 // COLORWIDTH 
	*/


	imx265_write_register(IspDev, 0x300B, 0x00); //Normal mode

////////
 	delay_ms(20);
    	imx265_write_register(IspDev, 0x3000, 0x00); //24'h000002,  //standby off
    	delay_ms(20);
    	imx265_write_register(IspDev, 0x300A, 0x00); //24'h000A02  //master mode start

///////
	sleep(1);
	
	//Exposure time [s] = (1 H period) × (Number of lines per frame - SHS) + 13.73μs
	//0.020 = ()x(1125-X) + 13.73
    unsigned int shutter = 400; // [6:1124], 6 max exposure time, 1124 min exposure time
	imx265_write_register(IspDev, 0x308D, shutter & 0xFF); //24'h012202 //???	
    imx265_write_register(IspDev, 0x308E, (shutter & 0xFF00) >> 8); //24'h012202 //???
	imx265_write_register(IspDev, 0x308F, (shutter & 0x0F0000) >> 16); //24'h012202 //???
	
	//imx265_write_register(IspDev, 0x3212, 0x09);//GAINDLY 08 or 09 (next frame as SHS)
	
	unsigned int gain = 0; //(gain = 60 = 6db * 10 ), min = 0 , max = 480d
	imx265_write_register(IspDev, 0x3204, gain & 0xFF); //gain itself
	imx265_write_register(IspDev, 0x3205, (gain & 0x100) >> 8); //gain itself
	
    printf("===ROI IMX265 FULLSCAN 25fps 12bit LINEAR Init OK!===\n");
    return;
}
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////
void imx265_linear_1080p25_init_orig(ISP_DEV IspDev) {

	imx265_write_register(IspDev, 0x3000, 0x01); 
delay_ms(20);

    imx265_write_register(IspDev, 0x3001, 0xD0); //24'hD00102,
    imx265_write_register(IspDev, 0x3002, 0xAA); //24'hAA0202,
    imx265_write_register(IspDev, 0x300D, 0x0C); //24'h0C0D02,
    
    imx265_write_register(IspDev, 0x3010, 0x65); //24'h651002, //VMAX = 24'h465                                                               
    imx265_write_register(IspDev, 0x3011, 0x04); //24'h041102,                                                                                 
    imx265_write_register(IspDev, 0x3012, 0x00); //24'h001202,
	//14
	//15
    
    //imx265_write_register(IspDev, 0x3014, 0x4C); //24'h501402, //HMAX = 0x44c for 60 fps                                                                                  
    //imx265_write_register(IspDev, 0x3015, 0x04); //24'h0A1502,
	
	imx265_write_register(IspDev, 0x3014, 0x98); //24'h501402, //HMAX = 898h for 30fps
    imx265_write_register(IspDev, 0x3015, 0x08); //24'h0A1502,


    imx265_write_register(IspDev, 0x3018, 0x01); //24'h011802,
    imx265_write_register(IspDev, 0x3019, 0x01); //24'h011902,
    imx265_write_register(IspDev, 0x3023, 0x00); //24'h002302,
    imx265_write_register(IspDev, 0x3080, 0x62); //---
    	//89*
    	//8A*
    	//8B*
    	//8C*
    	
    imx265_write_register(IspDev, 0x3089, 0x0C); //24'h188902, //INCK=74.25 MHz                                                                    
    imx265_write_register(IspDev, 0x308A, 0x00); //24'h008A02,                                                                                 
    imx265_write_register(IspDev, 0x308B, 0x10); //24'h108B02,                                                                                 
    imx265_write_register(IspDev, 0x308C, 0x00); //24'h028C02,
    
    //    24'h188902, //INCK=37.125 MHz                                                                    
    //24'h008A02,                                                                                 
    //24'h108B02,                                                                                 
    //24'h028C02,
    //imx265_write_register(IspDev, 0x3089, 0x18); //24'h188902, //INCK=37.125 MHz                                                              
    //imx265_write_register(IspDev, 0x308A, 0x00); //24'h008A02,                                                                                 
    //imx265_write_register(IspDev, 0x308B, 0x10); //24'h108B02,                                                                                 
    //imx265_write_register(IspDev, 0x308C, 0x02); //24'h028C02,

    
	//8D*
    	//8E*
    	//8F*

    //imx265_write_register(IspDev, 0x309E, 0x06); //24'h069E02,
    //imx265_write_register(IspDev, 0x30A0, 0x02); //24'h02A002,
    
    imx265_write_register(IspDev, 0x30AF, 0x0A); //24'h0AAF02,

    imx265_write_register(IspDev, 0x3168, 0xD8); //24'hD86803,
    imx265_write_register(IspDev, 0x3169, 0xA0); //24'hA06903,
    imx265_write_register(IspDev, 0x317D, 0xA1); //24'hA17D03,
    imx265_write_register(IspDev, 0x3180, 0x62); //24'h628003,
    imx265_write_register(IspDev, 0x3190, 0x9B); //24'h9B9003,
    imx265_write_register(IspDev, 0x3191, 0xA0); //24'hA09103,
    imx265_write_register(IspDev, 0x31A4, 0x3F); //24'h3FA403,
    imx265_write_register(IspDev, 0x31A5, 0xB1); //24'hB1A503,
    imx265_write_register(IspDev, 0x31E2, 0x00); //24'h00E203,
    imx265_write_register(IspDev, 0x31EA, 0x00); //24'h00EA03,

    imx265_write_register(IspDev, 0x3226, 0x03); //24'h032604,

    imx265_write_register(IspDev, 0x35AA, 0xB3); //24'hB3AA07,
    imx265_write_register(IspDev, 0x35AC, 0x68); //24'h68AC07,
    imx265_write_register(IspDev, 0x371C, 0xB4); //24'hB41C09,
    imx265_write_register(IspDev, 0x371D, 0x00); //24'h001D09,
    imx265_write_register(IspDev, 0x371E, 0xDE); //24'hDE1E09,
    imx265_write_register(IspDev, 0x371F, 0x00); //24'h001F09,
    imx265_write_register(IspDev, 0x3728, 0xB4); //24'hB42809,
    imx265_write_register(IspDev, 0x3729, 0x00); //24'h002909,
    imx265_write_register(IspDev, 0x372A, 0xDE); //24'hDE2A09,
    imx265_write_register(IspDev, 0x372B, 0x00); //24'h002B09,
    imx265_write_register(IspDev, 0x373A, 0x36); //24'h363A09,
    imx265_write_register(IspDev, 0x3746, 0x36); //24'h364609,
    imx265_write_register(IspDev, 0x38E0, 0xEB); //24'hEBE00A,
    imx265_write_register(IspDev, 0x38E1, 0x00); //24'h00E10A,
    imx265_write_register(IspDev, 0x38E2, 0x0D); //24'h0DE20A,
    imx265_write_register(IspDev, 0x38E3, 0x01); //24'h01E30A,
    imx265_write_register(IspDev, 0x39C4, 0xEB); //24'hEBC40B,
    imx265_write_register(IspDev, 0x39C5, 0x00); //24'h00C50B,
    imx265_write_register(IspDev, 0x39C6, 0x0C); //24'h0CC60B,
    imx265_write_register(IspDev, 0x39C7, 0x01); //24'h01C70B,
    imx265_write_register(IspDev, 0x3D02, 0x6E); //24'h6E020F,
    imx265_write_register(IspDev, 0x3D04, 0xE3); //24'hE3040F,
    imx265_write_register(IspDev, 0x3D05, 0x00); //24'h00050F,
    imx265_write_register(IspDev, 0x3D0C, 0x73); //24'h730C0F,
    imx265_write_register(IspDev, 0x3D0E, 0x6E); //24'h6E0E0F,
    imx265_write_register(IspDev, 0x3D10, 0xE8); //24'hE8100F,
    imx265_write_register(IspDev, 0x3D11, 0x00); //24'h00110F,
    imx265_write_register(IspDev, 0x3D12, 0xE3); //24'hE3120F,
    imx265_write_register(IspDev, 0x3D13, 0x00); //24'h00130F,
    imx265_write_register(IspDev, 0x3D14, 0x6B); //24'h6B140F,
    imx265_write_register(IspDev, 0x3D16, 0x1C); //24'h1C160F,
    imx265_write_register(IspDev, 0x3D18, 0x1C); //24'h1C180F,
    imx265_write_register(IspDev, 0x3D1A, 0x6B); //24'h6B1A0F,
    imx265_write_register(IspDev, 0x3D1C, 0x6E); //24'h6E1C0F,
    imx265_write_register(IspDev, 0x3D1E, 0x9A); //24'h9A1E0F,
    imx265_write_register(IspDev, 0x3D20, 0x12); //24'h12200F,
    imx265_write_register(IspDev, 0x3D22, 0x3E); //24'h3E220F,
    imx265_write_register(IspDev, 0x3D28, 0xB4); //24'hB4280F,
    imx265_write_register(IspDev, 0x3D29, 0x00); //24'h00290F,
    imx265_write_register(IspDev, 0x3D2A, 0x66); //24'h662A0F,
    imx265_write_register(IspDev, 0x3D34, 0x69); //24'h69340F,
    imx265_write_register(IspDev, 0x3D36, 0x17); //24'h17360F,
    imx265_write_register(IspDev, 0x3D38, 0x6A); //24'h6A380F,
    imx265_write_register(IspDev, 0x3D3A, 0x18); //24'h183A0F,
    imx265_write_register(IspDev, 0x3D3E, 0xFF); //24'hFF3E0F,
    imx265_write_register(IspDev, 0x3D3F, 0x0F); //24'h0F3F0F,
    imx265_write_register(IspDev, 0x3D46, 0xFF); //24'hFF460F,
    imx265_write_register(IspDev, 0x3D47, 0x0F); //24'h0F470F,
    imx265_write_register(IspDev, 0x3D4E, 0x4C); //24'h4C4E0F,
    imx265_write_register(IspDev, 0x3D50, 0x50); //24'h50500F,
    imx265_write_register(IspDev, 0x3D54, 0x73); //24'h73540F,
    imx265_write_register(IspDev, 0x3D56, 0x6E); //24'h6E560F,
    imx265_write_register(IspDev, 0x3D58, 0xE8); //24'hE8580F,
    imx265_write_register(IspDev, 0x3D59, 0x00); //24'h00590F,
    imx265_write_register(IspDev, 0x3D5A, 0xCF); //24'hCF5A0F,
    imx265_write_register(IspDev, 0x3D5B, 0x00); //24'h005B0F,
    imx265_write_register(IspDev, 0x3D5E, 0x64); //24'h645E0F,
    imx265_write_register(IspDev, 0x3D66, 0x61); //24'h61660F,
    imx265_write_register(IspDev, 0x3D6E, 0x0D); //24'h0D6E0F,
    imx265_write_register(IspDev, 0x3D70, 0xFF); //24'hFF700F,
    imx265_write_register(IspDev, 0x3D71, 0x0F); //24'h0F710F,
    imx265_write_register(IspDev, 0x3D72, 0x00); //24'h00720F,
    imx265_write_register(IspDev, 0x3D73, 0x00); //24'h00730F,
    imx265_write_register(IspDev, 0x3D74, 0x11); //24'h11740F,
    imx265_write_register(IspDev, 0x3D76, 0x6A); //24'h6A760F,
    imx265_write_register(IspDev, 0x3D78, 0x7F); //24'h7F780F,
    imx265_write_register(IspDev, 0x3D7A, 0xB3); //24'hB37A0F,
    imx265_write_register(IspDev, 0x3D7C, 0x29); //24'h297C0F,
    imx265_write_register(IspDev, 0x3D7E, 0x64); //24'h647E0F,
    imx265_write_register(IspDev, 0x3D80, 0xB1); //24'hB1800F,
    imx265_write_register(IspDev, 0x3D82, 0xB3); //24'hB3820F,
    imx265_write_register(IspDev, 0x3D84, 0x62); //24'h62840F,
    imx265_write_register(IspDev, 0x3D86, 0x64); //24'h64860F,
    imx265_write_register(IspDev, 0x3D88, 0xB1); //24'hB1880F,
    imx265_write_register(IspDev, 0x3D8A, 0xB3); //24'hB38A0F,
    imx265_write_register(IspDev, 0x3D8C, 0x62); //24'h628C0F,
    imx265_write_register(IspDev, 0x3D8E, 0x64); //24'h648E0F,
    imx265_write_register(IspDev, 0x3D90, 0x6D); //24'h6D900F,
    imx265_write_register(IspDev, 0x3D92, 0x65); //24'h65920F,
    imx265_write_register(IspDev, 0x3D94, 0x65); //24'h65940F,
    imx265_write_register(IspDev, 0x3D96, 0x6D); //24'h6D960F,
    imx265_write_register(IspDev, 0x3D98, 0x20); //24'h20980F,
    imx265_write_register(IspDev, 0x3D9A, 0x28); //24'h289A0F,
    imx265_write_register(IspDev, 0x3D9C, 0x81); //24'h819C0F,
    imx265_write_register(IspDev, 0x3D9E, 0x89); //24'h899E0F,
    imx265_write_register(IspDev, 0x3D9F, 0x01); //24'h019F0F,
    imx265_write_register(IspDev, 0x3DA0, 0x66); //24'h66A00F,
    imx265_write_register(IspDev, 0x3DA2, 0x7B); //24'h7BA20F,
    imx265_write_register(IspDev, 0x3DA4, 0x21); //24'h21A40F,
    imx265_write_register(IspDev, 0x3DA6, 0x27); //24'h27A60F,
    imx265_write_register(IspDev, 0x3DA8, 0x8B); //24'h8BA80F,
    imx265_write_register(IspDev, 0x3DA9, 0x01); //24'h01A90F,
    imx265_write_register(IspDev, 0x3DAA, 0x95); //24'h95AA0F,
    imx265_write_register(IspDev, 0x3DAB, 0x01); //24'h01AB0F,
    imx265_write_register(IspDev, 0x3DAC, 0x12); //24'h12AC0F,
    imx265_write_register(IspDev, 0x3DAE, 0x1C); //24'h1CAE0F,
    imx265_write_register(IspDev, 0x3DB0, 0x98); //24'h98B00F,
    imx265_write_register(IspDev, 0x3DB1, 0x01); //24'h01B10F,
    imx265_write_register(IspDev, 0x3DB2, 0xA0); //24'hA0B20F,
    imx265_write_register(IspDev, 0x3DB3, 0x01); //24'h01B30F,
    imx265_write_register(IspDev, 0x3DB4, 0x13); //24'h13B40F,
    imx265_write_register(IspDev, 0x3DB6, 0x1D); //24'h1DB60F,
    imx265_write_register(IspDev, 0x3DB8, 0x99); //24'h99B80F,
    imx265_write_register(IspDev, 0x3DB9, 0x01); //24'h01B90F,
    imx265_write_register(IspDev, 0x3DBA, 0xA1); //24'hA1BA0F,
    imx265_write_register(IspDev, 0x3DBB, 0x01); //24'h01BB0F,
    imx265_write_register(IspDev, 0x3DBC, 0x14); //24'h14BC0F,
    imx265_write_register(IspDev, 0x3DBE, 0x1E); //24'h1EBE0F,
    imx265_write_register(IspDev, 0x3DC0, 0x9A); //24'h9AC00F,
    imx265_write_register(IspDev, 0x3DC1, 0x01); //24'h01C10F,
    imx265_write_register(IspDev, 0x3DC2, 0xA2); //24'hA2C20F,
    imx265_write_register(IspDev, 0x3DC3, 0x01); //24'h01C30F,
    imx265_write_register(IspDev, 0x3DC4, 0x64); //24'h64C40F,
    imx265_write_register(IspDev, 0x3DC6, 0x6E); //24'h6EC60F,
    imx265_write_register(IspDev, 0x3DC8, 0x17); //24'h17C80F,
    imx265_write_register(IspDev, 0x3DCA, 0x26); //24'h26CA0F,
    imx265_write_register(IspDev, 0x3DCC, 0x9D); //24'h9DCC0F,
    imx265_write_register(IspDev, 0x3DCD, 0x01); //24'h01CD0F,
    imx265_write_register(IspDev, 0x3DCE, 0xAC); //24'hACCE0F,
    imx265_write_register(IspDev, 0x3DCF, 0x01); //24'h01CF0F,
    imx265_write_register(IspDev, 0x3DD0, 0x65); //24'h65D00F,
    imx265_write_register(IspDev, 0x3DD2, 0x6F); //24'h6FD20F,
    imx265_write_register(IspDev, 0x3DD4, 0x18); //24'h18D40F,
    imx265_write_register(IspDev, 0x3DD6, 0x27); //24'h27D60F,
    imx265_write_register(IspDev, 0x3DD8, 0x9E); //24'h9ED80F,
    imx265_write_register(IspDev, 0x3DD9, 0x01); //24'h01D90F,
    imx265_write_register(IspDev, 0x3DDA, 0xAD); //24'hADDA0F,
    imx265_write_register(IspDev, 0x3DDB, 0x01); //24'h01DB0F,
    imx265_write_register(IspDev, 0x3DDC, 0x66); //24'h66DC0F,
    imx265_write_register(IspDev, 0x3DDE, 0x70); //24'h70DE0F,
    imx265_write_register(IspDev, 0x3DE0, 0x19); //24'h19E00F,
    imx265_write_register(IspDev, 0x3DE2, 0x28); //24'h28E20F,
    imx265_write_register(IspDev, 0x3DE4, 0x9F); //24'h9FE40F,
    imx265_write_register(IspDev, 0x3DE5, 0x01); //24'h01E50F,
    imx265_write_register(IspDev, 0x3DE6, 0xAE); //24'hAEE60F,
    imx265_write_register(IspDev, 0x3DE7, 0x01); //24'h01E70F,
    imx265_write_register(IspDev, 0x3E04, 0x9D); //24'h9D0410,
    imx265_write_register(IspDev, 0x3E06, 0xB0); //24'hB00610,
    imx265_write_register(IspDev, 0x3E07, 0x00); //24'h000710,
    imx265_write_register(IspDev, 0x3E08, 0x6B); //24'h6B0810,
    imx265_write_register(IspDev, 0x3E0A, 0x7E); //24'h7E0A10,
    imx265_write_register(IspDev, 0x3E24, 0xE3); //24'hE32410,
    imx265_write_register(IspDev, 0x3E25, 0x00); //24'h002510,
    imx265_write_register(IspDev, 0x3E26, 0x9A); //24'h9A2610,
    imx265_write_register(IspDev, 0x3E27, 0x01); //24'h012710,
    imx265_write_register(IspDev, 0x3F20, 0x00); //24'h002011,
    imx265_write_register(IspDev, 0x3F21, 0x00); //24'h002111,
    imx265_write_register(IspDev, 0x3F22, 0xFF); //24'hFF2211,
    imx265_write_register(IspDev, 0x3F23, 0x3F); //24'h3F2311,
    imx265_write_register(IspDev, 0x4003, 0x55); //24'h550312,
    imx265_write_register(IspDev, 0x4005, 0xFF); //24'hFF0512,
    imx265_write_register(IspDev, 0x400B, 0x00); //24'h000B12,
    imx265_write_register(IspDev, 0x400C, 0x54); //24'h540C12,
    imx265_write_register(IspDev, 0x400D, 0xB8); //24'hB80D12,
    imx265_write_register(IspDev, 0x400E, 0x48); //24'h480E12,
    imx265_write_register(IspDev, 0x400F, 0xA2); //24'hA20F12,
    imx265_write_register(IspDev, 0x4012, 0x53); //24'h531212,
    imx265_write_register(IspDev, 0x4013, 0x0A); //24'h0A1312,
    imx265_write_register(IspDev, 0x4014, 0x0C); //24'h0C1412,
    imx265_write_register(IspDev, 0x4015, 0x0A); //24'h0A1512,
    imx265_write_register(IspDev, 0x402A, 0x7F); //24'h7F2A12,
    imx265_write_register(IspDev, 0x402C, 0x29); //24'h292C12,
    imx265_write_register(IspDev, 0x4030, 0x73); //24'h733012,
    imx265_write_register(IspDev, 0x4032, 0x8D); //24'h8D3212,
    imx265_write_register(IspDev, 0x4033, 0x01); //24'h013312,
    imx265_write_register(IspDev, 0x4049, 0x02); //24'h024912,
    imx265_write_register(IspDev, 0x4056, 0x18); //24'h185612,
    imx265_write_register(IspDev, 0x408C, 0x9A); //24'h9A8C12,
    imx265_write_register(IspDev, 0x408E, 0xAA); //24'hAA8E12,
    imx265_write_register(IspDev, 0x4090, 0x3E); //24'h3E9012,
    imx265_write_register(IspDev, 0x4092, 0x5F); //24'h5F9212,
    imx265_write_register(IspDev, 0x4094, 0x0A); //24'h0A9412,
    imx265_write_register(IspDev, 0x4096, 0x0A); //24'h0A9612,
    imx265_write_register(IspDev, 0x4098, 0x7F); //24'h7F9812,
    imx265_write_register(IspDev, 0x409A, 0xB3); //24'hB39A12,
    imx265_write_register(IspDev, 0x409C, 0x29); //24'h299C12,
    imx265_write_register(IspDev, 0x409E, 0x64); //24'h649E12,

	/////////////////////
	imx265_write_register(IspDev, 0x309E, 0x06); //24'h069E02,
    imx265_write_register(IspDev, 0x30A0, 0x02); //24'h02A002,
    ////////////
	
	
    imx265_write_register(IspDev, 0x3254, 0x80); //24'h805404 //black level
    imx265_write_register(IspDev, 0x3255, 0x00); //24'h805404
    
    imx265_write_register(IspDev, 0x3022, 0xF0); //24'h012202 //appnote black level auto adjust NO (F0) or YES (01)
    //imx265_write_register(IspDev, 0x3238, 0x06); //24'h063804 // 1D or 06 ???
	

	//test
	/*
    imx265_write_register(IspDev, 0x3254, 0x00); //24'h805404 //black level
    imx265_write_register(IspDev, 0x3255, 0x00); //24'h805404
    imx265_write_register(IspDev, 0x3022, 0xF0); //24'h012202 //appnote black level auto adjust NO (F0) or YES (01)
    //imx265_write_register(IspDev, 0x3238, 0x55); //24'h063804 // 1D or 06 ??? //55 hor color bars, 5d vert color bars
    //imx265_write_register(IspDev, 0x3248, 0x0F); //ColorWidth
    
    imx265_write_register(IspDev, 0x3238, 0x1D); //24'h063804 // 1D or 06 ??? //55 hor color bars, 5d vert color bars
    //imx265_write_register(IspDev, 0x3238, 0x45); //24'h063804 // 1D or 06 ??? //55 hor color bars, 5d vert color bars
    //imx265_write_register(IspDev, 0x3238, 0x45); //24'h063804 // 1D or 06 ??? //55 hor color bars, 5d vert color bars        
    //imx265_write_register(IspDev, 0x3244, 240);
    //imx265_write_register(IspDev, 0x3246, 135);
    
    //imx265_write_register(IspDev, 0x3248, 0x0F); //24'h063804 // COLORWIDTH 
	*/


	imx265_write_register(IspDev, 0x300B, 0x00); //Normal mode

////////
 	delay_ms(20);
    	imx265_write_register(IspDev, 0x3000, 0x00); //24'h000002,  //standby off
    	delay_ms(20);
    	imx265_write_register(IspDev, 0x300A, 0x00); //24'h000A02  //master mode start

///////
	sleep(1);
	
	//Exposure time [s] = (1 H period) × (Number of lines per frame - SHS) + 13.73μs
	//0.020 = ()x(1125-X) + 13.73
    unsigned int shutter = 400; // [6:1124], 6 max exposure time, 1124 min exposure time
	imx265_write_register(IspDev, 0x308D, shutter & 0xFF); //24'h012202 //???	
    imx265_write_register(IspDev, 0x308E, (shutter & 0xFF00) >> 8); //24'h012202 //???
	imx265_write_register(IspDev, 0x308F, (shutter & 0x0F0000) >> 16); //24'h012202 //???
	
	//imx265_write_register(IspDev, 0x3212, 0x09);//GAINDLY 08 or 09 (next frame as SHS)
	
	unsigned int gain = 0; //(gain = 60 = 6db * 10 ), min = 0 , max = 480d
	imx265_write_register(IspDev, 0x3204, gain & 0xFF); //gain itself
	imx265_write_register(IspDev, 0x3205, (gain & 0x100) >> 8); //gain itself
	
    printf("===IMX265 1080P 25fps 12bit LINEAR Init OK!===\n");
    return;
}


