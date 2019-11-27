//+build hi3516cv300 hi3516av200

package cmos

/*
#define init_module(module_image, len, param_values) syscall(__NR_init_module, module_image, len, param_values)
#define delete_module(name, flags) syscall(__NR_delete_module, name, flags)
int hi3516av200_ko_init(struct hi3516av200_cmos * c) {

    int error_code = 0;

    error_code = HI_MPI_SYS_Exit();
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: in ko_init HI_MPI_SYS_Exit failed\n");
        //return ERR_MPP;
    }
    printf("C DEBUG: HI_MPI_SYS_Exit ok\n");

    error_code = HI_MPI_VB_Exit();
    if (error_code != HI_SUCCESS) {
        printf("C DEBUG: in ko_init HI_MPI_VB_Exit failed\n");
        //return ERR_MPP;
    }
    printf("C DEBUG: HI_MPI_VB_Exit ok\n");


    unsigned int i = 0;


    while(modules[i].start != NULL) {
        delete_module(modules[i].name, NULL);
        //if (init_module(modules[i].start,
        //                modules[i].end-modules[i].start,
        //                modules[i].default_params) != 0) {
        //    printf("C DEBUG: init_module %s failed\n", modules[i].name);
        //    //return ERR_GENERAL; //TODO for proper restart
        //}
        //printf("init_module %s loaded OK!\n", modules[i].name);
        i++;
    }


    //delete_module("hi3519v101_isp", NULL);//ATTENTION THIS IS NEED FOR PROPER APP RERUN, also some info here
    //http://bbs.ebaina.com/forum.php?mod=viewthread&tid=13925&extra=&highlight=run%2Bae%2Blib%2Berr%2B0xffffffff%21&page=1
    /////

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

        //devmem(0x12030000, 0x00000204, NULL);

        //write priority select
        //devmem(0x12030054, 0x55552356, NULL); //  each module 4bit  cci       ---        ddrt  ---    ---     gzip   ---    ---
        //devmem(0x12030058, 0x16554411, NULL); // each module 4bit  vicap1    hash       ive   aio    jpge    tde   vicap0  vdp
        //devmem(0x1203005c, 0x33466314, NULL); // each module 4bit  mmc2      A17        fmc   sdio1  sdio0   A7    vpss0   vgs
        //devmem(0x12030060, 0x46266666, NULL); // each module 4bit  gdc       usb3/pcie  vedu  usb2   cipher  dma2  dma1    gsf

        //read priority select
        //devmem(0x12030064, 0x55552356, NULL); // each module 4bit  cci       ---         ddrt  ---    ---     gzip   ---    ---
        //devmem(0x12030068, 0x06554401, NULL); // each module 4bit  vicap1    hash        ive   aio    jpge    tde   vicap0  vdp
        //devmem(0x1203006c, 0x33466304, NULL); // each module 4bit  mmc2      A17         fmc   sdio1  sdio0    A7   vpss0   vgs
        //devmem(0x12030070, 0x46266666, NULL); // each module 4bit  gdc       usb3/pcie   vedu  usb2   cipher  dma2  dma1    gsf

        //devmem(0x120641f0, 0x1, NULL);  //use pri_map
        //write timeout select
        //devmem(0x1206409c, 0x00000040, NULL);
        //devmem(0x120640a0, 0x00000000, NULL);

        //read timeout select
        //devmem(0x120640ac, 0x00000040, NULL);
        //devmem(0x120640b0, 0x00000000, NULL);

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

    //delete_module("hi3519v101_isp", NULL);//ATTENTION THIS IS NEED FOR PROPER APP RERUN, also some info here
    //http://bbs.ebaina.com/forum.php?mod=viewthread&tid=13925&extra=&highlight=run%2Bae%2Blib%2Berr%2B0xffffffff%21&page=1
    /////

    //unsigned int i = 0;

    i=0;
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



    //    imx226)
    //        tmp=0x11;
    //        himm 0x12010040 0x11;           # sensor0 clk_en, 72MHz
    //        # SDK config:     IVE:396M,  GDC:475M,  VGS:500M,  VEDU:600M,   VPSS:300M
    //        #imx226: viu0:600M,isp0:600M, viu1:300M,isp1:300M
    //        himm 0x1201004c 0x00094c23;
    //        himm 0x12010054 0x00024041;
    //        spi0_4wire_pin_mux;
    //        insmod extdrv/hi_ssp_sony.ko;


    //fflush(stdout);

    //insmod /lib/modules/hi3519v101/extdrv/hi_ssp_sony.ko;


    //if (init_module(_binary_hi_ssp_sony_ko_start,
    //                _binary_hi_ssp_sony_ko_end - _binary_hi_ssp_sony_ko_start,
    //                "") != 0) {
    //    printf("C DEBUG: init_module hi_ssp_sony.ko failed\n");
    //   return ERR_GENERAL;//return EXIT_FAILURE;
    //}


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

*/
//import "C"
