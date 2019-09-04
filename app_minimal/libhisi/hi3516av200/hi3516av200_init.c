#include "../hisi_external.h"
#include "../hisi_utils.h"

#include "hi3516av200_ko.h"
#include "hi3516av200_mpp.h"

#include "hi3516av200_cmos.h"

#include "hi3516av200_channels.h"
#include "hi3516av200_encoders.h"

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

int hi3516av200_init        (struct hi3516av200_cmos * c, struct capture_params * cp);
int hi3516av200_ko_init     (struct hi3516av200_cmos * c);
int hi3516av200_sys_init    (struct hi3516av200_cmos * c);
int hi3516av200_mipi_init   (struct hi3516av200_cmos * c);
int hi3516av200_isp_init    (struct hi3516av200_cmos * c);
int hi3516av200_vi_init     (struct hi3516av200_cmos * c, struct capture_params * cp);
int hi3516av200_vpss_init   (struct capture_params   * cp);

int hi3516av200_init_temperature();

int hisi_cmos(struct cmos * c) {

    c->width = vpss.width;
    c->height = vpss.height;
    c->fps = vpss.fps;

    return ERR_NONE;
}

int hisi_init(unsigned int cid, struct capture_params * cp) {
    int error_code = 0;

    struct hi3516av200_cmos * c = &hi3516av200_cmoses[cid];
    struct capture_params cp_tmp;

    cp_tmp.fps      = 30;
    cp_tmp.x0       = 0;
    cp_tmp.y0       = 0;
    cp_tmp.width    = 3840;
    cp_tmp.height   = 2160;

    error_code = hi3516av200_init(c, &cp_tmp);

    for(int i=0; i<VPSS_MAX_PHY_CHN_NUM; i++) {
        channels_enable[i] = CHANNEL_DISABLED;
    }

    for(int i=0; i<VENC_MAX_CHN_NUM; i++) {
        encoders_enable[i] = ENCODER_DISABLED;
    }

    return error_code;
}

int hi3516av200_init(struct hi3516av200_cmos * c, struct capture_params * cp) {
    //int error_code = 0;

    /*
    unsigned int        id;
    char                *name;
    char                *description;
    unsigned int        width;
    unsigned int        height;
    unsigned int        fps;
    */

    printf("c.id = %d\n", c->id);
    printf("c.name= %s\n", c->name);
    printf("c.description = %s\n", c->description);
    printf("c.width = %d\n", c->width);
    printf("c.height = %d\n", c->height);
    printf("c.fps = %d\n", c->fps);

    if(hi3516av200_ko_init  (c)     != ERR_NONE) return ERR_INTERNAL;
    if(hi3516av200_sys_init (c)     != ERR_NONE) return ERR_INTERNAL;
    if(hi3516av200_mipi_init(c)     != ERR_NONE) return ERR_INTERNAL;
    if(hi3516av200_isp_init (c)     != ERR_NONE) return ERR_INTERNAL;
    if(hi3516av200_vi_init  (c, cp) != ERR_NONE) return ERR_INTERNAL;
    if(hi3516av200_vpss_init(cp)    != ERR_NONE) return ERR_INTERNAL;

    hi3516av200_init_temperature();

    return ERR_NONE;
}

////////////////////////////////////////////////////////////////////////////////

int hi3516av200_init_temperature() {
        devmem(0x120A0110, 0x60FA0000, NULL);
        return ERR_NONE;
}

////////////////////////////////////////////////////////////////////////////////

#define init_module(module_image, len, param_values) syscall(__NR_init_module, module_image, len, param_values)
#define delete_module(name, flags) syscall(__NR_delete_module, name, flags)

int hi3516av200_ko_init(struct hi3516av200_cmos * c) {
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

    /////

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

    /////

    #ifdef VPSS_ONLINE
        /*
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

        devmem(0x120641f0, 0x1, NULL);  //use pri_map
        //write timeout select
        devmem(0x1206409c, 0x00000040, NULL);
        devmem(0x120640a0, 0x00000000, NULL);

        //read timeout select
        devmem(0x120640ac, 0x00000040, NULL);
        devmem(0x120640b0, 0x00000000, NULL);
        */
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

        devmem(0x120641f0, 0x1, NULL);  // use pri_map
        //write timeout select
        devmem(0x1206409c, 0x00000040, NULL);   
        devmem(0x120640a0, 0x00000000, NULL);    

        //read timeout select
        devmem(0x120640ac, 0x00000040, NULL);   // each module 8bit
        devmem(0x120640b0, 0x00000000, NULL);
    #endif

    devmem(0x120300e0, 0xd, NULL); // internal codec: AIO MCLK out, CODEC AIO TX MCLK 

    delete_module("hi3519v101_isp", NULL);//ATTENTION THIS IS NEED FOR PROPER APP RERUN, also some info here
    //http://bbs.ebaina.com/forum.php?mod=viewthread&tid=13925&extra=&highlight=run%2Bae%2Blib%2Berr%2B0xffffffff%21&page=1
    /////

    unsigned int i = 0;

    while(modules[i].start != NULL) {
        if (init_module(modules[i].start,
                        modules[i].end-modules[i].start,
                        modules[i].default_params) != 0) {
            printf("C DEBUG: init_module %s failed\n", modules[i].name);
            //return ERR_GENERAL; //TODO for proper restart
        }
        //printf("init_module %s loaded OK!\n", modules[i].name);
        i++;
    }

    //imx274)
    //tmp=0x11;
    devmem(0x12010040, 0x11, NULL);       //sensor0 clk_en, 72MHz
    //SDK config:     IVE:396M,  GDC:475M,  VGS:500M,  VEDU:600M,   VPSS:300M 
    //imx274:viu0: 600M,isp0:300M, viu1:300M,isp1:300M
    devmem(0x1201004c, 0x00094c23, NULL);
    devmem(0x12010054, 0x0004041, NULL);
    //spi0_4wire_pin_mux;
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
    /////////////////////////////////////////////////


    /*
        imx226)
            tmp=0x11;
            himm 0x12010040 0x11;           # sensor0 clk_en, 72MHz
            # SDK config:     IVE:396M,  GDC:475M,  VGS:500M,  VEDU:600M,   VPSS:300M 
            #imx226: viu0:600M,isp0:600M, viu1:300M,isp1:300M
            himm 0x1201004c 0x00094c23;
            himm 0x12010054 0x00024041;
            spi0_4wire_pin_mux;
            insmod extdrv/hi_ssp_sony.ko;


    */

    //fflush(stdout);

    //insmod /lib/modules/hi3519v101/extdrv/hi_ssp_sony.ko;

    /*
    if (init_module(_binary_hi_ssp_sony_ko_start,
                    _binary_hi_ssp_sony_ko_end - _binary_hi_ssp_sony_ko_start,
                    "") != 0) {
        printf("C DEBUG: init_module hi_ssp_sony.ko failed\n");
        return ERR_GENERAL;//return EXIT_FAILURE;
    }
    */

    //single_pipe)
    devmem(0x12040184, 0x1, NULL); //   # SENSOR0 HS from VI0 HS
    devmem(0x12040188, 0x1, NULL); //   # SENSOR0 VS from VI0 VS
    devmem(0x12040010, 0x2, NULL); //   # SENSOR1 HS from VI1 HS
    devmem(0x12040014, 0x2, NULL); //   # SENSOR1 VS from VI1 VS
    /////
    devmem(0x12010044, 0x4ff0, NULL);
    devmem(0x12010044, 0x4, NULL);

    //fflush(stdout);

    return ERR_NONE;
}

////////////////////////////////////////////////////////////////////////////////

int hi3516av200_sys_init(struct hi3516av200_cmos * c) {
    int error_code = 0;

    error_code = HI_MPI_SYS_Exit();
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: TODO\n");
        return ERR_MPP;
    }
    printf("C DEBUG: HI_MPI_SYS_Exit ok\n");

    error_code = HI_MPI_VB_Exit();
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: TODO\n");
        return ERR_MPP;
    }
    printf("C DEBUG: HI_MPI_VB_Exit ok\n");

    VB_CONF_S stVbConf;

    memset(&stVbConf, 0, sizeof(VB_CONF_S));
    stVbConf.u32MaxPoolCnt                  = 128;
    stVbConf.astCommPool[0].u32BlkSize      = (CEILING_2_POWER(c->width, 64) * CEILING_2_POWER(c->height, 64) * 1.5);
    stVbConf.astCommPool[0].u32BlkCnt       = 10;

    error_code = HI_MPI_VB_SetConf(&stVbConf);
    if(error_code != HI_SUCCESS) {
        printf("C DEBUG: TODO\n");
        return ERR_MPP;
    }
    printf("C DEBUG: HI_MPI_VB_SetConf ok\n");

    error_code = HI_MPI_VB_Init();
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: TODO\n");
        return ERR_MPP;
    }
    printf("C DEBUG: HI_MPI_VB_Init ok\n");

    MPP_SYS_CONF_S stSysConf;

    stSysConf.u32AlignWidth = 64;

    error_code = HI_MPI_SYS_SetConf(&stSysConf);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: TODO\n");
        return ERR_MPP;
    }
    printf("C DEBUG: HI_MPI_SYS_SetConf ok\n");

    error_code = HI_MPI_SYS_Init();
    if(error_code != HI_SUCCESS) {
        printf("C DEBUG: TODO\n");
        return ERR_MPP;
    }
    printf("C DEBUG: HI_MPI_SYS_Init ok\n");

    return ERR_NONE;
}

////////////////////////////////////////////////////////////////////////////////

int hi3516av200_mipi_init(struct hi3516av200_cmos * c) {
    int error_code = 0;
    int fd;

    fd = open("/dev/hi_mipi", O_RDWR);
    if (fd < 0) {
        printf("C DEBUG: TODO\n");
        return ERR_GENERAL;
    }

    combo_dev_attr_t stcomboDevAttr;

    memcpy(&stcomboDevAttr, c->mipidev, sizeof(combo_dev_attr_t));
    stcomboDevAttr.devno = 0;

    printf("stcomboDevAttr memcpy ok\n");

    if(ioctl(fd, HI_MIPI_RESET_MIPI, &stcomboDevAttr.devno)) {
        printf("C DEBUG: TODO\n");
        close(fd);
        return ERR_GENERAL;
    }

    if(ioctl(fd, HI_MIPI_RESET_SENSOR, &stcomboDevAttr.devno)) {
        printf("C DEBUG: TODO HI_MIPI_RESET_SENSOR failed\n");
        close(fd);
        return ERR_GENERAL;
    }

    if (ioctl(fd, HI_MIPI_SET_DEV_ATTR, &stcomboDevAttr)) {
        printf("set mipi attr failed\n");
        close(fd);
        return ERR_GENERAL;
    }

    if(ioctl(fd, HI_MIPI_UNRESET_MIPI, &stcomboDevAttr.devno)) {
        printf("C DEBUG: TODO HI_MIPI_UNRESET_MIPI failed\n");
        close(fd);
        return ERR_GENERAL;
    }

    if(ioctl(fd, HI_MIPI_UNRESET_SENSOR, &stcomboDevAttr.devno)) {
        printf("C DEBUG: TODO HI_MIPI_UNRESET_SENSOR failed\n");
        close(fd);
        return ERR_GENERAL;
    }

    close(fd);

    return ERR_NONE;
}

////////////////////////////////////////////////////////////////////////////////

HI_VOID* hi3516av200_isp_thread(HI_VOID *param){
    int error_code = 0;
    printf("C DEBUG: starting HI_MPI_ISP_Run...\n");
    error_code = HI_MPI_ISP_Run(0);
    printf("C DEBUG: HI_MPI_ISP_Run %d\n", error_code);
    //return error_code;
}

static pthread_t hi3516av200_isp_thread_pid;

int hi3516av200_isp_init(struct hi3516av200_cmos * c) {
    int error_code = 0;

    ISP_PUB_ATTR_S stPubAttr;
    ALG_LIB_S stLib;

    error_code = HI_MPI_ISP_Exit(0);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: HI_MPI_ISP_Exit failed with %#x!\n", error_code);
        return ERR_MPP;
    }

    //const ISP_SNS_OBJ_S *g_pstSnsObj[2] =  {&stSnsImx274Obj, &stSnsImx226Obj};

    ALG_LIB_S stAeLib;
    ALG_LIB_S stAwbLib;
    ALG_LIB_S stAfLib;

    stAeLib.s32Id = 0;
    stAwbLib.s32Id = 0;
    stAfLib.s32Id = 0;
    strncpy(stAeLib.acLibName,  HI_AE_LIB_NAME,     sizeof(HI_AE_LIB_NAME));
    strncpy(stAwbLib.acLibName, HI_AWB_LIB_NAME,    sizeof(HI_AWB_LIB_NAME));
    strncpy(stAfLib.acLibName,  HI_AF_LIB_NAME,     sizeof(HI_AF_LIB_NAME)); 

    if (c->snsobj->pfnRegisterCallback != HI_NULL) {
        error_code = c->snsobj->pfnRegisterCallback(0, &stAeLib, &stAwbLib);
        if (error_code != HI_SUCCESS) {
            printf("C DEBUG: sensor_register_callback failed with %#x!\n", error_code);
            return ERR_GENERAL;
        }
    } else {
        printf("C DEBUG: sensor_register_callback failed with HI_NULL!\n");
        return ERR_GENERAL;
    }

    stLib.s32Id = 0;
    strcpy(stLib.acLibName, HI_AE_LIB_NAME);
    error_code = HI_MPI_AE_Register(0, &stLib);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: HI_MPI_AE_Register failed!\n");
        return ERR_MPP;
    }

    stLib.s32Id = 0;
    strcpy(stLib.acLibName, HI_AWB_LIB_NAME);
    error_code = HI_MPI_AWB_Register(0, &stLib);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: HI_MPI_AWB_Register failed!\n");
        return ERR_MPP;
    }

    stLib.s32Id = 0;
    strcpy(stLib.acLibName, HI_AF_LIB_NAME);
    error_code = HI_MPI_AF_Register(0, &stLib);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: HI_MPI_AF_Register failed!\n");
        return ERR_MPP;
    }

    error_code = HI_MPI_ISP_MemInit(0);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: HI_MPI_ISP_Init failed!\n");
        return ERR_MPP;
    }

    ISP_WDR_MODE_S stWdrMode;
    stWdrMode.enWDRMode  = WDR_MODE_NONE;

    error_code = HI_MPI_ISP_SetWDRMode(0, &stWdrMode);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: start ISP WDR failed!\n");
        return ERR_NONE;
    }
    //TODO WDR modes support

    stPubAttr.enBayer               = c->bayer;
    stPubAttr.f32FrameRate          = c->fps;
    stPubAttr.stWndRect.s32X        = 0;
    stPubAttr.stWndRect.s32Y        = 0;
    stPubAttr.stWndRect.u32Width    = c->width;     //TODO What is WND rect?
    stPubAttr.stWndRect.u32Height   = c->height;    //TODO
    stPubAttr.stSnsSize.u32Width    = c->width;
    stPubAttr.stSnsSize.u32Height   = c->height;

    error_code = HI_MPI_ISP_SetPubAttr(0, &stPubAttr);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: HI_MPI_ISP_SetPubAttr failed with %#x!\n", error_code);
        return ERR_MPP;
    }

    error_code = HI_MPI_ISP_Init(0);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: HI_MPI_ISP_Init failed!\n");
        return ERR_MPP;
    }

    if (pthread_create(&hi3516av200_isp_thread_pid, 0, (void* (*)(void*))hi3516av200_isp_thread, NULL) != 0) {
        printf("C DEBUG: create isp running thread failed!\n");
        return ERR_GENERAL;
    }

    return ERR_NONE;
}

////////////////////////////////////////////////////////////////////////////////

int hi3516av200_vi_init(struct hi3516av200_cmos * c, struct capture_params * cp) {
    int error_code;

    VI_DEV_ATTR_S  stViDevAttr;

    memset(&stViDevAttr,0,sizeof(stViDevAttr));
    memcpy(&stViDevAttr, c->videv, sizeof(stViDevAttr));

    //stViDevAttr.stDevRect.s32Y                              = 0;
    stViDevAttr.stDevRect.u32Width                          = c->width;
    stViDevAttr.stDevRect.u32Height                         = c->height;
    stViDevAttr.stBasAttr.stSacleAttr.stBasSize.u32Width    = c->width;
    stViDevAttr.stBasAttr.stSacleAttr.stBasSize.u32Height   = c->height;
    stViDevAttr.stBasAttr.stSacleAttr.bCompress             = HI_FALSE;

    error_code = HI_MPI_VI_SetDevAttr(0, &stViDevAttr);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: HI_MPI_VI_SetDevAttr failed with %#x!\n", error_code);
        return ERR_MPP;
    }

    error_code = HI_MPI_VI_EnableDev(0);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: HI_MPI_VI_EnableDev failed with %#x!\n", error_code);
        return ERR_MPP;
    }

    RECT_S stCapRect;
    SIZE_S stTargetSize;

    stCapRect.s32X          = cp->x0;
    stCapRect.s32Y          = cp->y0;
    stCapRect.u32Width      = cp->width;
    stCapRect.u32Height     = cp->height;
    stTargetSize.u32Width   = stCapRect.u32Width;
    stTargetSize.u32Height  = stCapRect.u32Height;

    VI_CHN_ATTR_S stChnAttr;

    memcpy(&stChnAttr.stCapRect, &stCapRect, sizeof(RECT_S));

    stChnAttr.enCapSel              = VI_CAPSEL_BOTH;
    stChnAttr.stDestSize.u32Width   = stTargetSize.u32Width ;
    stChnAttr.stDestSize.u32Height  = stTargetSize.u32Height ;
    stChnAttr.enPixFormat           = PIXEL_FORMAT_YUV_SEMIPLANAR_420;   /* sp420 or sp422 */

    stChnAttr.bMirror               = HI_FALSE;
    stChnAttr.bFlip                 = HI_FALSE;

    stChnAttr.s32SrcFrameRate       = c->fps;
    stChnAttr.s32DstFrameRate       = cp->fps;
    stChnAttr.enCompressMode        = COMPRESS_MODE_NONE;

    error_code = HI_MPI_VI_SetChnAttr(0, &stChnAttr);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: HI_MPI_VI_SetChnAttr failed with %#x!\n", error_code);
        return ERR_MPP;
    }

    error_code = HI_MPI_VI_EnableChn(0);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: HI_MPI_VI_EnableChn failed with %#x!\n", error_code);
        return ERR_MPP;
    }

    vpss.width = cp->width;
    vpss.height = cp->height;
    vpss.fps = cp->fps;

    return ERR_NONE;
}

////////////////////////////////////////////////////////////////////////////////

int hi3516av200_vpss_init(struct capture_params * cp) {
    int error_code;

    VPSS_GRP_ATTR_S stVpssGrpAttr;

    stVpssGrpAttr.u32MaxW           = cp->width;
    stVpssGrpAttr.u32MaxH           = cp->width;
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

    error_code = HI_MPI_VPSS_CreateGrp(0, &stVpssGrpAttr);
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: HI_MPI_VPSS_CreateGrp failed with %#x!\n", error_code);
        return ERR_MPP;
    }

    error_code = HI_MPI_VPSS_StartGrp(0);
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
        return ERR_MPP;
    }

    return ERR_NONE;
}

