#include "himpp3_external.h"
#include "himpp3_internal.h"
#include "himpp3_ko.h"

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <pthread.h>
#include <signal.h>
#include <sys/uio.h>
#include <stdint.h>
#include <sys/ioctl.h>
#include <fcntl.h>
#include <sys/stat.h>
#include <sys/syscall.h>
#include <sys/types.h>
#include <errno.h>
#include <ctype.h>
#include <sys/mman.h>
#include <sys/select.h>
#include <inttypes.h>

int devmem(uint32_t target, uint32_t value, uint32_t * read);
void * himpp3_venc_jpeg_test_loop(void * arg);

char * getChipFamily() {
    return "hi3516av200";
}

#ifdef TEST_C_APP
int main(void) {
        //devmem(0x120100e4);
	//himpp3_ko_init();

	//himpp3_sys_init();
        //himpp3_mipi_isp_init();
        //himpp3_vi_init();
	//himpp3_vpss_init();
        //himpp3_venc_init();
	//while(1) ;;;
	return 0;
}
#endif


int inittemperature() {
        devmem(0x120A0110, 0x60FA0000, NULL);
        return 0;
}
float gettemperature() {
        uint32_t read;
        devmem(0x120A0118, -1, &read);
        printf("C DEBUG: temperature code 0x%lx\n", read & 0x3FF);
        printf("C DEBUG: temperature C %.1f\n", (float)((( (read & 0x3FF)-125) / 806.0 ) * 165) - 40 );
        return (float)(((( (read & 0x3FF) - 125) / 806.0 ) *165) - 40);
}


#define init_module(module_image, len, param_values) syscall(__NR_init_module, module_image, len, param_values)
#define delete_module(name, flags) syscall(__NR_delete_module, name, flags)

int himpp3_ko_init() {
        //sensor0 pinmux
        devmem(0x1204017c, 0x1, NULL);  //#SENSOR0_CLK
        devmem(0x12040180, 0x0, NULL);  //#SENSOR0_RSTN
        devmem(0x12040184, 0x1, NULL);  //#SENSOR0_HS,from vi0
        devmem(0x12040188, 0x1, NULL);  //#SENSOR0_VS,from vi0
        //sensor0 drive capability
        devmem(0x12040988, 0x150, NULL);  //#SENSOR0_CLK
        devmem(0x1204098c, 0x170, NULL);  //#SENSOR0_RSTN
        devmem(0x12040990, 0x170, NULL);  //#SENSOR0_HS
        devmem(0x12040994, 0x170, NULL);  //#SENSOR0_VS 

        ////////////////////////////////////////////////

        devmem(0x120100e4, 0x1ff70000, NULL); //# I2C0-3/SSP0-3 unreset, enable clk gate
        devmem(0x1201003c, 0x31000100, NULL);     //# MIPI VI ISP unreset
        devmem(0x12010050, 0x2, NULL);            //# VEDU0 unreset 
        devmem(0x12010058, 0x2, NULL);            //# VPSS0 unreset 
        devmem(0x12010058, 0x3, NULL);            //# VPSS0 unreset 
        devmem(0x12010058, 0x2, NULL);            //# VPSS0 unreset 
        devmem(0x1201005c, 0x2, NULL);            //# VGS unreset 
        devmem(0x12010060, 0x2, NULL);            //# JPGE unreset 
        devmem(0x12010064, 0x2, NULL);            //# TDE unreset 
        devmem(0x1201006c, 0x2, NULL);            //# IVE unreset      
        devmem(0x12010070, 0x2, NULL);            //# FD unreset
        devmem(0x12010074, 0x2, NULL);            //# GDC unreset 
        devmem(0x1201007C, 0x2a, NULL);           //# HASH&SAR ADC&CIPHER unreset   
        devmem(0x12010080, 0x2, NULL);            //# AIAO unreset,clock 1188M
        devmem(0x12010084, 0x2, NULL);            //# GZIP unreset  
        devmem(0x120100d8, 0xa, NULL);            //# ddrt efuse enable clock, unreset
        devmem(0x120100e0, 0xa8, NULL);           //# rsa trng klad enable clock, unreset
        //#himm 0x120100e0 0xaa       //# rsa trng klad DMA enable clock, unreset
        devmem(0x12010040, 0x60, NULL);
        devmem(0x12010040, 0x0, NULL);            //# sensor unreset,unreset the control module with slave-mode

        //# pcie clk enable
        devmem(0x120100b0, 0x000001f0, NULL);

        ////////////////////////////////////////////////////

        #ifdef VPSS_ONLINE
		devmem(0x12030000, 0x00000204, NULL);
		
		//write priority select
		devmem(0x12030054, 0x55552356, NULL); //  each module 4bit  cci       ---        ddrt  ---    ---     gzip   ---    ---
		devmem(0x12030058, 0x16554411, NULL); // each module 4bit  vicap1    hash       ive   aio    jpge    tde   vicap0  vdp 
		devmem(0x1203005c, 0x33466314, NULL); // each module 4bit  mmc2      A17        fmc   sdio1  sdio0   A7    vpss0   vgs 
		devmem(0x12030060, 0x46266666, NULL); // each module 4bit  gdc       usb3/pcie  vedu  usb2   cipher  dma2  dma1    gsf

		//read priority select
		devmem(0x12030064, 0x55552356, NULL); // each module 4bit  cci       ---         ddrt  ---    ---     gzip   ---    ---
		devmem(0x12030068, 0x06554401, NULL); // each module 4bit  vicap1    hash        ive   aio    jpge    tde   vicap0  vdp
		devmem(0x1203006c, 0x33466304, NULL); // each module 4bit  mmc2      A17         fmc   sdio1  sdio0    A7   vpss0   vgs 
		devmem(0x12030070, 0x46266666, NULL); // each module 4bit  gdc       usb3/pcie   vedu  usb2   cipher  dma2  dma1    gsf

		devmem(0x120641f0, 0x1, NULL);	//use pri_map
		//write timeout select
		devmem(0x1206409c, 0x00000040, NULL);	
		devmem(0x120640a0, 0x00000000, NULL);	
		
		//read timeout select
		devmem(0x120640ac, 0x00000040, NULL);	
		devmem(0x120640b0, 0x00000000, NULL);	
	#else //VPSS_OFFLINE
		devmem(0x12030000, 0x00000004, NULL);

		//write priority select
		devmem(0x12030054, 0x55552366, NULL); // each module 4bit  cci       ---        ddrt  ---    ---     gzip   ---    ---
		devmem(0x12030058, 0x16556611, NULL); // each module 4bit  vicap1    hash       ive   aio    jpge    tde   vicap0  vdp
		devmem(0x1203005c, 0x43466445, NULL); // each module 4bit  mmc2      A17        fmc   sdio1  sdio0   A7    vpss0   vgs
		devmem(0x12030060, 0x56466666, NULL); // each module 4bit  gdc       usb3/pcie  vedu  usb2   cipher  dma2  dma1    gsf

		//read priority select
		devmem(0x12030064, 0x55552366, NULL); // each module 4bit  cci       ---        ddrt  ---    ---     gzip   ---    ---
		devmem(0x12030068, 0x06556600, NULL); // each module 4bit  vicap1    hash       ive   aio    jpge    tde   vicap0  vdp
		devmem(0x1203006c, 0x43466435, NULL); // each module 4bit  mmc2      A17        fmc   sdio1  sdio0   A7    vpss0   vgs
		devmem(0x12030070, 0x56266666, NULL); // each module 4bit  gdc       usb3/pcie  vedu  usb2   cipher  dma2  dma1    gsf

		devmem(0x120641f0, 0x1, NULL);	// use pri_map
		//write timeout select
		devmem(0x1206409c, 0x00000040, NULL);	
		devmem(0x120640a0, 0x00000000, NULL);	 
		
		//read timeout select
		devmem(0x120640ac, 0x00000040, NULL);	// each module 8bit
		devmem(0x120640b0, 0x00000000, NULL);
	#endif


        devmem(0x120300e0, 0xd, NULL); // internal codec: AIO MCLK out, CODEC AIO TX MCLK 
        //#himm 0x120300e0 0xe		//# external codec: AIC31 AIO MCLK out, CODEC AIO TX MCLK
        //fflush(stdout);
        ///////////////////////////////////////////////////

        unsigned char i = 0; 
        /*
        while(modules[i].start != NULL) {
                i++;
        }
        i--;

        while(i>0) {
                if (delete_module(modules[i].name, NULL) != 0) {
                        printf("C DEBUG: delete_module %s failed\n", modules[i].name);
                        //return 1;//return EXIT_FAILURE;
                }
                printf("delete_module %s OK!\n", modules[i].name);
                i--;
        }
        */

        delete_module("hi3519v101_isp", NULL);//ATTENTION THIS IS NEED FOR PROPER APP RERUN, also some info here
        //http://bbs.ebaina.com/forum.php?mod=viewthread&tid=13925&extra=&highlight=run%2Bae%2Blib%2Berr%2B0xffffffff%21&page=1
        ///////////////////////////////////////////////////

        //unsigned char i = 0; 
        while(modules[i].start != NULL) {
                if (init_module(modules[i].start, 
                                modules[i].end-modules[i].start, 
                                modules[i].params) != 0) {
                        printf("C DEBUG: init_module %s failed\n", modules[i].name);
                        //return 1;//return EXIT_FAILURE;
                }
                //printf("init_module %s loaded OK!\n", modules[i].name);
                i++;
        }

        ////////////////////////////////////////////////////
        
        //imx274)
        //tmp=0x11;
        devmem(0x12010040, 0x11, NULL);       //sensor0 clk_en, 72MHz
        //SDK config:     IVE:396M,  GDC:475M,  VGS:500M,  VEDU:600M,   VPSS:300M 
        //imx274:viu0: 600M,isp0:300M, viu1:300M,isp1:300M
        devmem(0x1201004c, 0x00094c23, NULL);
        devmem(0x12010054, 0x0004041, NULL);
        // spi0_4wire_pin_mux;
        //pinmux
        devmem(0x1204018c, 0x1, NULL); //  #SPI0_SCLK
        devmem(0x12040190, 0x1, NULL); // #SPI0_SD0
        devmem(0x12040194, 0x1, NULL); // #SPI0_SDI
        devmem(0x12040198, 0x1, NULL); // #SPI0_CSN
    
        //drive capability
        devmem(0x12040998, 0x150, NULL); // #SPI0_SCLK
        devmem(0x1204099c, 0x160, NULL); // #SPI0_SD0
        devmem(0x120409a0, 0x160, NULL); // #SPI0_SDI
        devmem(0x120409a4, 0x160, NULL); // #SPI0_CSN

        //fflush(stdout);

        // insmod /lib/modules/hi3519v101/extdrv/hi_ssp_sony.ko;
	
        if (init_module(_binary_hi_ssp_sony_ko_start, 
                                _binary_hi_ssp_sony_ko_end -_binary_hi_ssp_sony_ko_start, 
                                "") != 0) {
                        printf("C DEBUG: init_module hi_ssp_sony.ko failed\n");
                        //return 1;//return EXIT_FAILURE;
                }
	

        //single_pipe)
        devmem(0x12040184, 0x1, NULL); //   # SENSOR0 HS from VI0 HS
        devmem(0x12040188, 0x1, NULL); //   # SENSOR0 VS from VI0 VS
        devmem(0x12040010, 0x2, NULL); //   # SENSOR1 HS from VI1 HS
        devmem(0x12040014, 0x2, NULL); //   # SENSOR1 VS from VI1 VS

        //////////////////////////////////////////////////
        devmem(0x12010044, 0x4ff0, NULL);
        devmem(0x12010044, 0x4, NULL);    

        //fflush(stdout);
	inittemperature();


	return 0;
}

int devmem(uint32_t target, uint32_t value, uint32_t * read) {
        //reference https://github.com/pavel-a/devmemX/blob/master/devmem2.c

        unsigned int pagesize = (unsigned)getpagesize(); /* or sysconf(_SC_PAGESIZE)  */
        unsigned int map_size = pagesize;
        
         int access_size = 4;

        unsigned offset;

        offset = (unsigned int)(target & (pagesize-1));
        if (offset + access_size > pagesize ) {
                // Access straddles page boundary:  add another page:
                map_size += pagesize;
        }

        int fd;
        void *map_base, *virt_addr;

        fd = open("/dev/mem", O_RDWR | O_SYNC);
        if (fd == -1) {
                printf("C DEBUG: Error opening /dev/mem (%d) : %s\n", errno, strerror(errno));
                return 1;
        }

        map_base = mmap(0, map_size, PROT_READ | PROT_WRITE, MAP_SHARED,
                    fd, 
                    target & ~((typeof(target))pagesize-1));

        if (map_base == (void *) -1) {
                printf("C DEBUG: Error mapping (%d) : %s\n", errno, strerror(errno));
                return 1;//exit(1);
        }
        //printf("Memory mapped at address %p.\n", map_base);

        virt_addr = map_base + offset;

        //unsigned long read_result;
	if (read == NULL ) {
        	*((volatile uint32_t *) virt_addr) = value;
	}
	if (read != NULL) {
        	*read = *((volatile uint32_t *) virt_addr);
	}
        //printf("0x%lx value 0x%lx\n", target, read_result);

        if (munmap(map_base, map_size) != 0) {
                printf("C DEBUG: ERROR munmap (%d) %s\n", errno, strerror(errno));
        }

        close(fd);
        return 0;
}

int himpp3_sys_init() {
        //int ret;

        int error_code = 0;
        // *error_func = HIMPP3_ERROR_FUNC_NONE;
        // *error_code = 0;


        //error_code = HI_MPI_ISP_Exit(0);
        //if (error_code != HI_SUCCESS) {
        //        printf("C DEBUG: %s: HI_MPI_ISP_Exit failed with %#x!\n", __FUNCTION__, error_code);
        //        return -1;
        //}      


       error_code = HI_MPI_ISP_Exit(0);
        if (error_code != HI_SUCCESS) {
                printf("C DEBUG: %s: HI_MPI_ISP_Exit failed with %#x!\n", __FUNCTION__, error_code);
                return -1;
        }  


        ISP_DEV IspDev = 0;
    
        ISP_PUB_ATTR_S stPubAttr;
        ALG_LIB_S stLib;


    const ISP_SNS_OBJ_S *g_pstSnsObj[2] =  {&stSnsImx274Obj, HI_NULL};

       ALG_LIB_S stAeLib;
        ALG_LIB_S stAwbLib;
        ALG_LIB_S stAfLib;

        stAeLib.s32Id = 0;
        stAwbLib.s32Id = 0;
        stAfLib.s32Id = 0;
        strncpy(stAeLib.acLibName, HI_AE_LIB_NAME, sizeof(HI_AE_LIB_NAME));
        strncpy(stAwbLib.acLibName, HI_AWB_LIB_NAME, sizeof(HI_AWB_LIB_NAME));
        strncpy(stAfLib.acLibName, HI_AF_LIB_NAME, sizeof(HI_AF_LIB_NAME)); 


        stLib.s32Id = 0;
        strcpy(stLib.acLibName, HI_AE_LIB_NAME);
        error_code = HI_MPI_AE_UnRegister(IspDev, &stLib);
        if (error_code != HI_SUCCESS) {
                printf("C DEBUG: %s: HI_MPI_AE_UnRegister failed!\n", __FUNCTION__);
                //return -1;
        }
        stLib.s32Id = 0;
        strcpy(stLib.acLibName, HI_AWB_LIB_NAME);
        error_code = HI_MPI_AWB_UnRegister(IspDev, &stLib);
        if (error_code != HI_SUCCESS) {
                printf("C DEBUG: %s: HI_MPI_AWB_UnRegister failed!\n", __FUNCTION__);
                //return -1;
        }
        stLib.s32Id = 0;
        strcpy(stLib.acLibName, HI_AF_LIB_NAME);
        error_code = HI_MPI_AF_UnRegister(IspDev, &stLib);
        if (error_code != HI_SUCCESS) {
                printf("C DEBUG: %s: HI_MPI_AF_UnRegister failed!\n", __FUNCTION__);
                //return -1;
        }

   printf("C DEBUG: Try to unregister SENSOR callback\n");


    if (g_pstSnsObj[0]->pfnUnRegisterCallback != HI_NULL)
    {
        error_code = g_pstSnsObj[0]->pfnUnRegisterCallback(IspDev, &stAeLib, &stAwbLib);
        if (error_code != HI_SUCCESS)
        {
            printf("C DEBUG: %s: sensor_unregister_callback failed with %#x!\n", __FUNCTION__, error_code);
            //return error_code;
        }
    }
    else
    {
        printf("C DEBUG: %s: sensor_unregister_callback failed with HI_NULL!\n", __FUNCTION__);
    }

    printf("C DEBUG: Try to unregister SENSOR callback DONE\n");


////////////////////////////////////////////

	error_code = HI_MPI_SYS_Exit();
        if (error_code != HI_SUCCESS) {
                // *error_func = HIMPP3_ERROR_FUNC_HI_MPI_SYS_Exit;
                return -1;
        }
        printf("C DEBUG: HI_MPI_SYS_Exit ok\n");

        error_code = HI_MPI_VB_Exit();
        if (error_code != HI_SUCCESS) {
                // *error_func = HIMPP3_ERROR_FUNC_HI_MPI_VB_Exit;
                return -1;
        }
        printf("C DEBUG: HI_MPI_VB_Exit ok\n");


        VB_CONF_S stVbConf;
        
        memset(&stVbConf, 0, sizeof(VB_CONF_S));
        stVbConf.u32MaxPoolCnt                  = 128;
        stVbConf.astCommPool[0].u32BlkSize      = (CEILING_2_POWER(3840, 64) * CEILING_2_POWER(2160, 64) * 1.5);
        stVbConf.astCommPool[0].u32BlkCnt       = 10;

        error_code = HI_MPI_VB_SetConf(&stVbConf);
	if(error_code != HI_SUCCESS) {
		// *error_func = HIMPP3_ERROR_FUNC_HI_MPI_VB_SetConf;
		return -1;
	}
        printf("C DEBUG: HI_MPI_VB_SetConf ok\n");

	error_code = HI_MPI_VB_Init();
	if (error_code != HI_SUCCESS) {
		// *error_func = HIMPP3_ERROR_FUNC_???;
		return -1;
	}
        printf("C DEBUG: HI_MPI_VB_Init ok\n");

        MPP_SYS_CONF_S	stSysConf;
	
        stSysConf.u32AlignWidth = 64;

	error_code = HI_MPI_SYS_SetConf(&stSysConf);
	if (error_code != HI_SUCCESS) {
		// *error_func = HIMPP3_ERROR_FUNC_HI_MPI_SYS_SetConf;
		return -1;
	}
        printf("C DEBUG: HI_MPI_SYS_SetConf ok\n");

	error_code = HI_MPI_SYS_Init();
	if(error_code != HI_SUCCESS) {
		// *error_func = HIMPP3_ERROR_FUNC_HI_MPI_SYS_Init;
		return -1;
	}
        printf("C DEBUG: HI_MPI_SYS_Init ok\n");

        return 0;
}

HI_VOID* Test_ISP_Run(HI_VOID *param){
        int error_code;
        ISP_DEV IspDev = 0;
        printf("C DEBUG: starting HI_MPI_ISP_Run...\n");
        error_code = HI_MPI_ISP_Run(IspDev);
        printf("C DEBUG: HI_MPI_ISP_Run %d\n", error_code);

        return HI_NULL;
}
static pthread_t gs_IspPid;

int himpp3_mipi_isp_init() {
        int error_code;

        int fd;
        combo_dev_attr_t *pstcomboDevAttr, stcomboDevAttr;


/*
       error_code = HI_MPI_ISP_Exit(0);
        if (error_code != HI_SUCCESS) {
                printf("C DEBUG: %s: HI_MPI_ISP_Exit failed with %#x!\n", __FUNCTION__, error_code);
                return -1;
        }        

        */



        /* mipi reset unrest */
        fd = open("/dev/hi_mipi", O_RDWR);
        if (fd < 0) {
                //printf("warning: open hi_mipi dev failed\n");
                return -1;
        }
    
 	pstcomboDevAttr = &LVDS_6lane_SENSOR_IMX274_12BIT_8M_NOWDR_ATTR;
	
        memcpy(&stcomboDevAttr, pstcomboDevAttr, sizeof(combo_dev_attr_t));
        stcomboDevAttr.devno = 0;

        /* 1.reset mipi */
        if(ioctl(fd, HI_MIPI_RESET_MIPI, &stcomboDevAttr.devno)) {
                //printf("HI_MIPI_RESET_MIPI failed\n");
                close(fd);
                return -1;
   	}

        /* 2.reset sensor */
        if(ioctl(fd, HI_MIPI_RESET_SENSOR, &stcomboDevAttr.devno)) {
    	        //printf("HI_MIPI_RESET_SENSOR failed\n");
                close(fd);
                return -1;
        }

        if (ioctl(fd, HI_MIPI_SET_DEV_ATTR, pstcomboDevAttr)) {
                //printf("set mipi attr failed\n");
                close(fd);
                return -1;
        }

        /* 4.unreset mipi */
        if(ioctl(fd, HI_MIPI_UNRESET_MIPI, &stcomboDevAttr.devno)) {
                //printf("HI_MIPI_UNRESET_MIPI failed\n");
                close(fd);
                return -1;
        }

        /* 5.unreset sensor */
        if(ioctl(fd, HI_MIPI_UNRESET_SENSOR, &stcomboDevAttr.devno)) {
                //printf("HI_MIPI_UNRESET_SENSOR failed\n");
                close(fd);
                return -1;
        }

        close(fd);
 
/*
       error_code = HI_MPI_ISP_Exit(0);
        if (error_code != HI_SUCCESS) {
                printf("C DEBUG: %s: HI_MPI_ISP_Exit failed with %#x!\n", __FUNCTION__, error_code);
                return -1;
        }        
  */      
        ISP_DEV IspDev = 0;
    
        ISP_PUB_ATTR_S stPubAttr;
        ALG_LIB_S stLib;


	const ISP_SNS_OBJ_S *g_pstSnsObj[2] =  {&stSnsImx274Obj, HI_NULL};

       ALG_LIB_S stAeLib;
        ALG_LIB_S stAwbLib;
        ALG_LIB_S stAfLib;

        stAeLib.s32Id = 0;
        stAwbLib.s32Id = 0;
        stAfLib.s32Id = 0;
        strncpy(stAeLib.acLibName, HI_AE_LIB_NAME, sizeof(HI_AE_LIB_NAME));
        strncpy(stAwbLib.acLibName, HI_AWB_LIB_NAME, sizeof(HI_AWB_LIB_NAME));
        strncpy(stAfLib.acLibName, HI_AF_LIB_NAME, sizeof(HI_AF_LIB_NAME)); 

/*
        printf("C DEBUG: Try to unregister SENSOR callback\n");


    if (g_pstSnsObj[0]->pfnUnRegisterCallback != HI_NULL)
    {
        error_code = g_pstSnsObj[0]->pfnUnRegisterCallback(IspDev, &stAeLib, &stAwbLib);
        if (error_code != HI_SUCCESS)
        {
            printf("C DEBUG: %s: sensor_unregister_callback failed with %#x!\n", __FUNCTION__, error_code);
            //return error_code;
        }
    }
    else
    {
        printf("C DEBUG: %s: sensor_unregister_callback failed with HI_NULL!\n", __FUNCTION__);
    }

    printf("C DEBUG: Try to unregister SENSOR callback DONE\n");


*/
	/*
        ALG_LIB_S stAeLib;
        ALG_LIB_S stAwbLib;
        ALG_LIB_S stAfLib;

        stAeLib.s32Id = 0;
        stAwbLib.s32Id = 0;
        stAfLib.s32Id = 0;
        strncpy(stAeLib.acLibName, HI_AE_LIB_NAME, sizeof(HI_AE_LIB_NAME));
        strncpy(stAwbLib.acLibName, HI_AWB_LIB_NAME, sizeof(HI_AWB_LIB_NAME));
        strncpy(stAfLib.acLibName, HI_AF_LIB_NAME, sizeof(HI_AF_LIB_NAME)); 
    */
/*
        stLib.s32Id = 0;
        strcpy(stLib.acLibName, HI_AE_LIB_NAME);
        error_code = HI_MPI_AE_UnRegister(IspDev, &stLib);
        if (error_code != HI_SUCCESS) {
                printf("C DEBUG: %s: HI_MPI_AE_UnRegister failed!\n", __FUNCTION__);
                //return -1;
        }
        stLib.s32Id = 0;
        strcpy(stLib.acLibName, HI_AWB_LIB_NAME);
        error_code = HI_MPI_AWB_UnRegister(IspDev, &stLib);
        if (error_code != HI_SUCCESS) {
                printf("C DEBUG: %s: HI_MPI_AWB_UnRegister failed!\n", __FUNCTION__);
                //return -1;
        }
        stLib.s32Id = 0;
        strcpy(stLib.acLibName, HI_AF_LIB_NAME);
        error_code = HI_MPI_AF_UnRegister(IspDev, &stLib);
        if (error_code != HI_SUCCESS) {
                printf("C DEBUG: %s: HI_MPI_AF_UnRegister failed!\n", __FUNCTION__);
                //return -1;
        }
*/


	if (g_pstSnsObj[0]->pfnRegisterCallback != HI_NULL) {
                error_code = g_pstSnsObj[0]->pfnRegisterCallback(IspDev, &stAeLib, &stAwbLib);
                if (error_code != HI_SUCCESS) {
                        printf("C DEBUG: %s: sensor_register_callback failed with %#x!\n", __FUNCTION__, error_code);
                        return -1;
                }
        } else {
                printf("C DEBUG: %s: sensor_register_callback failed with HI_NULL!\n",  __FUNCTION__);
                return -1;
        }

        /* 2. register hisi ae lib */
        stLib.s32Id = 0;
        strcpy(stLib.acLibName, HI_AE_LIB_NAME);
        error_code = HI_MPI_AE_Register(IspDev, &stLib);
        if (error_code != HI_SUCCESS) {
                printf("C DEBUG: %s: HI_MPI_AE_Register failed!\n", __FUNCTION__);
                //return -1;
        }

        /* 3. register hisi awb lib */
        stLib.s32Id = 0;
        strcpy(stLib.acLibName, HI_AWB_LIB_NAME);
        error_code = HI_MPI_AWB_Register(IspDev, &stLib);
        if (error_code != HI_SUCCESS) {
                printf("C DEBUG: %s: HI_MPI_AWB_Register failed!\n", __FUNCTION__);
                //return -1;
        }

        /* 4. register hisi af lib */
        stLib.s32Id = 0;
        strcpy(stLib.acLibName, HI_AF_LIB_NAME);
        error_code = HI_MPI_AF_Register(IspDev, &stLib);
        if (error_code != HI_SUCCESS) {
                printf("C DEBUG: %s: HI_MPI_AF_Register failed!\n", __FUNCTION__);
                //return -1;
        }

        /* 5. isp mem init */
        error_code = HI_MPI_ISP_MemInit(IspDev);
        if (error_code != HI_SUCCESS) {
                printf("C DEBUG: %s: HI_MPI_ISP_Init failed!\n", __FUNCTION__);
                return -1;
        }

        /* 6. isp set WDR mode */
        ISP_WDR_MODE_S stWdrMode;
	stWdrMode.enWDRMode  = WDR_MODE_NONE;
	
        error_code = HI_MPI_ISP_SetWDRMode(0, &stWdrMode);    
        if (error_code != HI_SUCCESS) {
                printf("C DEBUG: start ISP WDR failed!\n");
                return -1;
        }

	stPubAttr.enBayer		= BAYER_RGGB;
        stPubAttr.f32FrameRate          = 30;
        stPubAttr.stWndRect.s32X        = 0;
        stPubAttr.stWndRect.s32Y        = 0;
        stPubAttr.stWndRect.u32Width    = 3840;
        stPubAttr.stWndRect.u32Height   = 2160;
        stPubAttr.stSnsSize.u32Width    = 3840;
        stPubAttr.stSnsSize.u32Height   = 2160;    

        error_code = HI_MPI_ISP_SetPubAttr(IspDev, &stPubAttr);
        if (error_code != HI_SUCCESS) {
                printf("C DEBUG: %s: HI_MPI_ISP_SetPubAttr failed with %#x!\n", __FUNCTION__, error_code);
                return -1;
        }

        /* 8. isp init */
        error_code = HI_MPI_ISP_Init(IspDev);
        if (error_code != HI_SUCCESS) {
                printf("C DEBUG: %s: HI_MPI_ISP_Init failed!\n", __FUNCTION__);
                return -1;
        }

        if (0 != pthread_create(&gs_IspPid, 0, (void* (*)(void*))Test_ISP_Run, NULL)) {
                printf("C DEBUG: %s: create isp running thread failed!\n", __FUNCTION__);
                return -1;
        }

	return 0;

}

int himpp3_vi_init() {
        int error_code;

        VI_DEV_ATTR_S  stViDevAttr;
    
        memset(&stViDevAttr,0,sizeof(stViDevAttr));
	memcpy(&stViDevAttr, &DEV_ATTR_LVDS_BASE, sizeof(stViDevAttr));

        stViDevAttr.stDevRect.s32Y                              = 0;
        stViDevAttr.stDevRect.u32Width                          = 3840;
        stViDevAttr.stDevRect.u32Height                         = 2160;
        stViDevAttr.stBasAttr.stSacleAttr.stBasSize.u32Width    = 3840;
        stViDevAttr.stBasAttr.stSacleAttr.stBasSize.u32Height   = 2160;
        stViDevAttr.stBasAttr.stSacleAttr.bCompress             = HI_FALSE;

        error_code = HI_MPI_VI_SetDevAttr(0, &stViDevAttr);
        if (error_code != HI_SUCCESS) {
                printf("C DEBUG: HI_MPI_VI_SetDevAttr failed with %#x!\n", error_code);
                return -1;
        }
 
        error_code = HI_MPI_VI_EnableDev(0);
        if (error_code != HI_SUCCESS) {
                printf("C DEBUG: HI_MPI_VI_EnableDev failed with %#x!\n", error_code);
                return -1;
        }

        RECT_S stCapRect;
        SIZE_S stTargetSize;

        stCapRect.s32X          = 0;
        stCapRect.s32Y          = 0;
        stCapRect.u32Width      = 3840;
        stCapRect.u32Height     = 2160;
        stTargetSize.u32Width   = stCapRect.u32Width;
        stTargetSize.u32Height  = stCapRect.u32Height;

        VI_CHN_ATTR_S stChnAttr;

	memcpy(&stChnAttr.stCapRect, &stCapRect, sizeof(RECT_S));
        
        stChnAttr.enCapSel              = VI_CAPSEL_BOTH;
        stChnAttr.stDestSize.u32Width   = stTargetSize.u32Width ;
        stChnAttr.stDestSize.u32Height  = stTargetSize.u32Height ;
        stChnAttr.enPixFormat           = PIXEL_FORMAT_YUV_SEMIPLANAR_420;   /* sp420 or sp422 */

        stChnAttr.bMirror       = HI_FALSE;
        stChnAttr.bFlip         = HI_FALSE;

        stChnAttr.s32SrcFrameRate       = 30;
        stChnAttr.s32DstFrameRate       = 30;
        stChnAttr.enCompressMode        = COMPRESS_MODE_NONE;

        error_code = HI_MPI_VI_SetChnAttr(0, &stChnAttr);
        if (error_code != HI_SUCCESS) {
                printf("C DEBUG: HI_MPI_VI_SetChnAttr failed with %#x!\n", error_code);
                return -1;
        }

        //#define CMOS_LDC
        //#ifdef CMOS_LDC
	/*
        VI_LDC_ATTR_S stLDCAttr;
        //First enable VI devices and VI channel.
        //Initialize LDC attributes.
        stLDCAttr.bEnable = HI_TRUE;
        stLDCAttr.stAttr.enViewType = LDC_VIEW_TYPE_ALL;
        //LDC_VIEW_TYPE_CROP;
        stLDCAttr.stAttr.s32CenterXOffset = 0;
        stLDCAttr.stAttr.s32CenterYOffset = 0;
        stLDCAttr.stAttr.s32Ratio = 58;
        stLDCAttr.stAttr.s32MinRatio = 0;
        //Set LDC attributes.
        error_code = HI_MPI_VI_SetLDCAttr(0, &stLDCAttr);
        if (error_code != HI_SUCCESS) {
                printf("Set vi LDC attr err:0x%x\n", error_code);
                return -1;
        }
        printf("HI_MPI_VI_SetLDCAttr ok\n");

        //Obtain LDC attributes.
        error_code = HI_MPI_VI_GetLDCAttr (0, &stLDCAttr);
        if (error_code != HI_SUCCESS) {
                printf("Get vi LDC attr err:0x%x\n", error_code);
                return -1;
        }
        printf("HI_MPI_VI_GetLDCAttr ok\n");
        //#endif
	*/

        error_code = HI_MPI_VI_EnableChn(0);
        if (error_code != HI_SUCCESS) {
                printf("C DEBUG: HI_MPI_VI_EnableChn failed with %#x!\n", error_code);
                return -1;
        }

	return 0;

}

int himpp3_vpss_init() {
        int error_code;

        VPSS_GRP VpssGrp = 0;
        VPSS_GRP_ATTR_S stVpssGrpAttr;

        VpssGrp = 0;

	stVpssGrpAttr.u32MaxW           = 3840;
	stVpssGrpAttr.u32MaxH           = 2160;
	stVpssGrpAttr.bIeEn             = HI_FALSE;
	stVpssGrpAttr.bNrEn             = HI_TRUE;
	stVpssGrpAttr.bHistEn           = HI_FALSE;
	stVpssGrpAttr.bDciEn            = HI_FALSE;
	stVpssGrpAttr.enDieMode         = VPSS_DIE_MODE_NODIE;
	stVpssGrpAttr.enPixFmt          = PIXEL_FORMAT_YUV_SEMIPLANAR_420;//SAMPLE_PIXEL_FORMAT;
        stVpssGrpAttr.bStitchBlendEn    = HI_FALSE;

        stVpssGrpAttr.stNrAttr.enNrType                         = VPSS_NR_TYPE_VIDEO;
	stVpssGrpAttr.stNrAttr.stNrVideoAttr.enNrRefSource      = VPSS_NR_REF_FROM_RFR;//VPSS_NR_REF_FROM_CHN0, VPSS_NR_REF_FROM_SRC
        stVpssGrpAttr.stNrAttr.stNrVideoAttr.enNrOutputMode     = VPSS_NR_OUTPUT_NORMAL;//VPSS_NR_OUTPUT_DELAY NORMAL
	stVpssGrpAttr.stNrAttr.u32RefFrameNum                   = 2;

        error_code = HI_MPI_VPSS_CreateGrp(VpssGrp, &stVpssGrpAttr);
        if (error_code != HI_SUCCESS) {
                printf("C DEBUG: HI_MPI_VPSS_CreateGrp failed with %#x!\n", error_code);
                return -1;
        }

        error_code = HI_MPI_VPSS_StartGrp(VpssGrp);
        if (error_code != HI_SUCCESS) {
                printf("C DEBUG: HI_MPI_VPSS_StartGrp failed with %#x\n", error_code);
                return -1;
        }

        MPP_CHN_S stSrcChn;
        MPP_CHN_S stDestChn;

	stSrcChn.enModId  = HI_ID_VIU;
        stSrcChn.s32DevId = 0;
        stSrcChn.s32ChnId = 0;
    
        stDestChn.enModId  = HI_ID_VPSS;
        stDestChn.s32DevId = 0;
        stDestChn.s32ChnId = 0;
    
        error_code = HI_MPI_SYS_Bind(&stSrcChn, &stDestChn);
        if (error_code != HI_SUCCESS) {
                printf("C DEBUG: HI_MPI_SYS_Bind failed with %#x!\n", error_code);
                return -1;
        }

        /////

        VPSS_CHN VpssChn;
        VPSS_CHN_ATTR_S stVpssChnAttr;
        VPSS_CHN_MODE_S stVpssChnMode;

	VpssChn = 1;
        stVpssChnMode.enChnMode      = VPSS_CHN_MODE_USER;
        stVpssChnMode.bDouble        = HI_FALSE;
        stVpssChnMode.enPixelFormat  = PIXEL_FORMAT_YUV_SEMIPLANAR_420;
        stVpssChnMode.u32Width       = 1920;// 
        stVpssChnMode.u32Height      = 1080;
        stVpssChnMode.enCompressMode = COMPRESS_MODE_NONE;//COMPRESS_MODE_SEG;
    
        memset(&stVpssChnAttr, 0, sizeof(stVpssChnAttr));
        stVpssChnAttr.s32SrcFrameRate = 30;
        stVpssChnAttr.s32DstFrameRate = 1;

	error_code = HI_MPI_VPSS_SetChnAttr(VpssGrp, VpssChn, &stVpssChnAttr);
	if (error_code != HI_SUCCESS) {
    	        printf("C DEBUG: HI_MPI_VPSS_SetChnAttr failed with %#x\n", error_code);
                return -1;
        }

	error_code = HI_MPI_VPSS_SetChnMode(VpssGrp, VpssChn, &stVpssChnMode);
        if (error_code != HI_SUCCESS) {
    	        printf("C DEBUG: %s failed with %#x\n", __FUNCTION__, error_code);
                return -1;
        }         

	/*
        VPSS_CROP_INFO_S CropInfo;
        CropInfo.bEnable = 1;
        CropInfo.enCropCoordinate = VPSS_CROP_ABS_COOR;
        CropInfo.stCropRect.s32X = 0;
        CropInfo.stCropRect.s32Y = 0;
        CropInfo.stCropRect.u32Width = 640;
        CropInfo.stCropRect.u32Height = 480;

        error_code = HI_MPI_VPSS_SetChnCrop(VpssGrp, VpssChn, &CropInfo);
    	if (error_code != HI_SUCCESS) {
        	printf("HI_MPI_VPSS_SetChnCrop failed with %#x\n", error_code);
        	return -1;
    	}
	*/

	error_code = HI_MPI_VPSS_EnableChn(VpssGrp, VpssChn);
        if (error_code != HI_SUCCESS) {
                printf("C DEBUG: HI_MPI_VPSS_EnableChn failed with %#x\n", error_code);
                return -1;
        }

	return 0;

}

static pthread_t gs_JpegPid;

//EXTERNAL
int himpp3_venc_max_chn_num() {
    return VENC_MAX_CHN_NUM;
}


//EXTERNAL
int himpp3_venc_mjpeg_params(unsigned int bitrate){
    int error_code;

    VENC_CHN_ATTR_S stVencChnAttr;


    error_code = HI_MPI_VENC_GetChnAttr(0, &stVencChnAttr);
    if (error_code != HI_SUCCESS) {
                printf("C DEBUG: HI_MPI_VENC_GetChnAttr faild with%#x!\n", error_code);
                return -1;
    }

    printf("C DEBUG: try to set %d bitrate\n", bitrate);
    stVencChnAttr.stRcAttr.stAttrMjpegeCbr.u32BitRate = bitrate;

    error_code = HI_MPI_VENC_SetChnAttr(0, &stVencChnAttr);
    if (error_code != HI_SUCCESS) {
                printf("C DEBUG: HI_MPI_VENC_SetChnAttr faild with%#x!\n", error_code);
                return -1;
    }

    return 0;
}


int himpp3_venc_init() {
        int error_code;

        VENC_CHN_ATTR_S stVencChnAttr;
        //VENC_ATTR_JPEG_S stJpegAttr;

        VENC_ATTR_MJPEG_S stMjpegAttr;
        VENC_ATTR_MJPEG_FIXQP_S stMjpegeFixQp;


        stVencChnAttr.stVeAttr.enType = PT_MJPEG;

        stMjpegAttr.u32MaxPicWidth = 3840;
        stMjpegAttr.u32MaxPicHeight = 2160;
        stMjpegAttr.u32PicWidth = 1920;//640;
        stMjpegAttr.u32PicHeight = 1080;//480;
        stMjpegAttr.u32BufSize = 3840 * 2160 * 3;
        stMjpegAttr.bByFrame = HI_TRUE;  /*get stream mode is field mode  or frame mode*/
        memcpy(&stVencChnAttr.stVeAttr.stAttrMjpege, &stMjpegAttr, sizeof(VENC_ATTR_MJPEG_S));

        stVencChnAttr.stRcAttr.enRcMode = VENC_RC_MODE_MJPEGCBR;
        stVencChnAttr.stRcAttr.stAttrMjpegeCbr.u32StatTime       = 1;
        stVencChnAttr.stRcAttr.stAttrMjpegeCbr.u32SrcFrmRate      = 1;
        stVencChnAttr.stRcAttr.stAttrMjpegeCbr.fr32DstFrmRate = 1;
        stVencChnAttr.stRcAttr.stAttrMjpegeCbr.u32FluctuateLevel = 1;
        stVencChnAttr.stRcAttr.stAttrMjpegeCbr.u32BitRate = 512 * 1;  

        /*
        stJpegAttr.u32PicWidth  = 3840;
        stJpegAttr.u32PicHeight = 2160;
        stJpegAttr.u32MaxPicWidth  = 3840;
        stJpegAttr.u32MaxPicHeight = 2160;
        stJpegAttr.u32BufSize   = 3840 * 2160 * 3;
        stJpegAttr.bByFrame     = HI_TRUE;
        stJpegAttr.bSupportDCF  = HI_FALSE;

        memcpy(&stVencChnAttr.stVeAttr.stAttrJpege, &stJpegAttr, sizeof(VENC_ATTR_JPEG_S));
        */
 	stVencChnAttr.stGopAttr.enGopMode  = VENC_GOPMODE_NORMALP;
        stVencChnAttr.stGopAttr.stNormalP.s32IPQpDelta = 0;
        



	error_code = HI_MPI_VENC_CreateChn(0, &stVencChnAttr);
        if (error_code != HI_SUCCESS) {
                printf("C DEBUG: HI_MPI_VENC_CreateChn [%d] faild with %#x!\n", 0, error_code);
                return -1;
        }

        error_code = HI_MPI_VENC_StartRecvPic(0);
        if (error_code != HI_SUCCESS) {
                printf("C DEBUG: HI_MPI_VENC_StartRecvPic faild with%#x!\n", error_code);
                return -1;
        }

        MPP_CHN_S stSrcChn;
        MPP_CHN_S stDestChn;

        stSrcChn.enModId = HI_ID_VPSS;
        stSrcChn.s32DevId = 0;
        stSrcChn.s32ChnId = 1;

        stDestChn.enModId = HI_ID_VENC;
        stDestChn.s32DevId = 0;
        stDestChn.s32ChnId = 0;

        error_code = HI_MPI_SYS_Bind(&stSrcChn, &stDestChn);
        if (error_code != HI_SUCCESS) {
                printf("C DEBUG: HI_MPI_SYS_Bind failed with %#x!\n", error_code);
                return -1;
        }

        int fd =  HI_MPI_VENC_GetFd(0);

        //////////////////////////////////////////
        /*
        int fd =  HI_MPI_VENC_GetFd(0);
        printf("HI_MPI_VENC_GetFd(0) %d\n", fd);

        VENC_STREAM_S stStream;
        VENC_CHN_STAT_S stStat;

        while(1) {
                memset(&stStream, 0, sizeof(stStream));

                error_code = HI_MPI_VENC_Query(0, &stStat);
                if (error_code != HI_SUCCESS) {
    	                printf("HI_MPI_VENC_Query chn[%d] failed with %#x!\n", 0, error_code);                    
                        return 1;
                }
                
                //stStream.pstPack = (VENC_PACK_S*)malloc(sizeof(VENC_PACK_S) * stStat.u32CurPacks);
                //stStream.u32PackCount = stStat.u32CurPacks;

                error_code = HI_MPI_VENC_GetStream(0, &stStream, -1);
                if (error_code != HI_SUCCESS) {
                        printf("HI_MPI_VENC_GetStream failed with %#x!\n", error_code);
                }
                printf("got frame\n");

                error_code = HI_MPI_VENC_ReleaseStream(0, &stStream);
                if (error_code != HI_SUCCESS) {
                        printf("failed to release stream: %#x\n", error_code);
                }
                printf("frame released\n");
        }
        */
        //////////////////////////////////////////

        if (0 != pthread_create(&gs_JpegPid, 0, &himpp3_venc_jpeg_test_loop, &fd)) {
                printf("C DEBUG: %s: create jpeg running thread failed!\n", __FUNCTION__);
                return -1;
        }


	return fd;


}


void * himpp3_venc_jpeg_test_loop(void * arg) {
	int fd = *((int *)arg);

	printf("C DEBUG: himpp3_venc_jpeg_test_loop fd %d\n", fd);

	//return 0;

 	VENC_STREAM_S stStream;
        VENC_CHN_STAT_S stStat;
	int error_code;
	fd_set read_fds;

while(1) {
        FD_ZERO(&read_fds);
        FD_SET(fd, &read_fds);

	error_code = select(fd + 1, &read_fds, NULL, NULL, NULL);
	if (error_code < 0) {
		printf("C DEBUG: select failed\n");
	}

	if (FD_ISSET(fd, &read_fds)) {

		memset(&stStream, 0, sizeof(stStream));

                error_code = HI_MPI_VENC_Query(0, &stStat);
                if (error_code != HI_SUCCESS) {
                        printf("C DEBUG: HI_MPI_VENC_Query chn[%d] failed with %#x!\n", 0, error_code);
                        return 0;
                }

                stStream.pstPack = (VENC_PACK_S*)malloc(sizeof(VENC_PACK_S) * stStat.u32CurPacks);
                stStream.u32PackCount = stStat.u32CurPacks;

                error_code = HI_MPI_VENC_GetStream(0, &stStream, -1);
                if (error_code != HI_SUCCESS) {
                        printf("C DEBUG: HI_MPI_VENC_GetStream failed with %#x!\n", error_code);
                }
                //printf("got frame\n");
		/////
		struct jpegFrame newFrame;
		newFrame.seq = stStream.u32Seq;
		newFrame.count = stStream.u32PackCount;
		for (int i =0;i<stStream.u32PackCount;i++) {
			newFrame.packs[i].length = stStream.pstPack[i].u32Len;
			newFrame.packs[i].data = stStream.pstPack[i].pu8Addr;
			newFrame.packs[i].pts = stStream.pstPack[i].u64PTS;
		}
		/*
		printf("C: ");
		for (int i =0;i<10;i++) {
			printf("%d ", stStream.pstPack[0].pu8Addr[i]);
		}
		printf("\n");
		*/
		jpegVencGetDataCallback(&newFrame);
		/////
		free(stStream.pstPack);

                error_code = HI_MPI_VENC_ReleaseStream(0, &stStream);
                if (error_code != HI_SUCCESS) {
                        printf("C DEBUG: failed to release stream: %#x\n", error_code);
                }
                //printf("frame released\n");
	}
}

}

//int himpp3_venc_jpeg_release_frame() {
//}

/*
static char test_byte = 'Y';
char * himpp3_test_func(char ** buffer) {
	printf("test byte address %p\n", &test_byte);
	*buffer = &test_byte;
	//printf("buffer pointer after %p\n", buffer);
	return &test_byte;
}
*/
