//+build hi3516av200 hi3516cv300

package mpp

/*
#include "./include/hi3516av200_mpp.h"

#define ERR_NONE    0
#define ERR_MPP     1

int mpp3_sys_exit(int *error_code) {
    *error_code = 0;
    *error_code = HI_MPI_SYS_Exit();
    if (*error_code != HI_SUCCESS) return ERR_MPP;
    return ERR_NONE;
}

int mpp3_vb_exit(int *error_code) {
    *error_code = 0;
    *error_code = HI_MPI_VB_Exit();
    if (*error_code != HI_SUCCESS) return ERR_MPP;
    return ERR_NONE;
}

int mpp3_isp_exit(int *error_code) {
    *error_code = 0;
    *error_code = HI_MPI_ISP_Exit(0);
    if (*error_code != HI_SUCCESS) return ERR_MPP;
    return ERR_NONE;
}
*/
import "C"

import (
    "log"
    "os"

	"application/pkg/koloader"
    "application/pkg/utils"
    "application/pkg/mpp/error"
)

//TODO rework this mess
func systemInit() {
    //NOTE should be done, otherwise there can be kernel panic on module unload
    //ATTENTION maybe isp exit should be added as well
    //TODO rework, add error codes, deal with C includes

    if _, err := os.Stat("/dev/sys"); err == nil { 
        var errorCode C.int
        switch err := C.mpp3_sys_exit(&errorCode); err {
        case C.ERR_NONE:
            log.Println("C.mpp3_sys_exit() ok")
        case C.ERR_MPP:
            log.Fatal("C.mpp3_sys_exit() HI_MPI_SYS_Exit() error ", error.Resolve(uint(errorCode))) 
        default:
            log.Fatal("Unexpected return ", err , " of C.mpp3_sys_exit()")
        } 
    }

    /*if _, err := os.Stat("/dev/isp_dev"); err == nil { //kernel panic !?
        var errorCode C.int
        switch err := C.mpp3_isp_exit(&errorCode); err {
        case C.ERR_NONE:
            log.Println("C.mpp3_isp_exit() ok")
        case C.ERR_MPP:
            log.Fatal("C.mpp3_isp_exit() HI_MPI_ISP_Exit() error ", error.Resolve(uint(errorCode))) 
        default:
            log.Fatal("Unexpected return ", err , " of C.mpp3_isp_exit()")
        }
    }*/

    if _, err := os.Stat("/dev/vb"); err == nil {      
        var errorCode C.int
        switch err := C.mpp3_vb_exit(&errorCode); err {
        case C.ERR_NONE:
            log.Println("C.mpp3_vb_exit() ok")
        case C.ERR_MPP:
            log.Fatal("C.mpp3_vb_exit() HI_MPI_VB_Exit() error ", error.Resolve(uint(errorCode))) 
        default:
            log.Fatal("Unexpected return ", err , " of C.mpp3_vb_exit()")
        }
    }
	//delete_module("hi3519v101_isp", NULL);//ATTENTION THIS IS NEED FOR PROPER APP RERUN, also some info here
    //http://bbs.ebaina.com/forum.php?mod=viewthread&tid=13925&extra=&highlight=run%2Bae%2Blib%2Berr%2B0xffffffff%21&page=1
	koloader.UnloadAll()

	//sensor0 pinmux
	utils.WriteDevMem32(0x1204017c, 0x1);  //#SENSOR0_CLK
	utils.WriteDevMem32(0x12040180, 0x0);  //#SENSOR0_RSTN
	utils.WriteDevMem32(0x12040184, 0x1);  //#SENSOR0_HS,from vi0
	utils.WriteDevMem32(0x12040188, 0x1);  //#SENSOR0_VS,from vi0
	//sensor0 drive capability
	utils.WriteDevMem32(0x12040988, 0x150);  //#SENSOR0_CLK
	utils.WriteDevMem32(0x1204098c, 0x170);  //#SENSOR0_RSTN
	utils.WriteDevMem32(0x12040990, 0x170);  //#SENSOR0_HS
	utils.WriteDevMem32(0x12040994, 0x170);  //#SENSOR0_VS 	

	/////

    utils.WriteDevMem32(0x120100e4, 0x1ff70000); //# I2C0-3/SSP0-3 unreset, enable clk gate
    utils.WriteDevMem32(0x1201003c, 0x31000100);     //# MIPI VI ISP unreset
    utils.WriteDevMem32(0x12010050, 0x2);            //# VEDU0 unreset 
    utils.WriteDevMem32(0x12010058, 0x2);            //# VPSS0 unreset 
    utils.WriteDevMem32(0x12010058, 0x3);            //# VPSS0 unreset 
    utils.WriteDevMem32(0x12010058, 0x2);            //# VPSS0 unreset 
    utils.WriteDevMem32(0x1201005c, 0x2);            //# VGS unreset 
    utils.WriteDevMem32(0x12010060, 0x2);            //# JPGE unreset 
    utils.WriteDevMem32(0x12010064, 0x2);            //# TDE unreset 
    utils.WriteDevMem32(0x1201006c, 0x2);            //# IVE unreset      
    utils.WriteDevMem32(0x12010070, 0x2);            //# FD unreset
    utils.WriteDevMem32(0x12010074, 0x2);            //# GDC unreset 
    utils.WriteDevMem32(0x1201007C, 0x2a);           //# HASH&SAR ADC&CIPHER unreset   
    utils.WriteDevMem32(0x12010080, 0x2);            //# AIAO unreset,clock 1188M
    utils.WriteDevMem32(0x12010084, 0x2);            //# GZIP unreset  
    utils.WriteDevMem32(0x120100d8, 0xa);            //# ddrt efuse enable clock, unreset
    utils.WriteDevMem32(0x120100e0, 0xa8);           //# rsa trng klad enable clock, unreset
    //#himm 0x120100e0 0xaa       //# rsa trng klad DMA enable clock, unreset
    utils.WriteDevMem32(0x12010040, 0x60);
    utils.WriteDevMem32(0x12010040, 0x0);            //# sensor unreset,unreset the control module with slave-mode

	//# pcie clk enable
    utils.WriteDevMem32(0x120100b0, 0x000001f0);

	//#ifdef VPSS_ONLINE
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
    //#else //VPSS_OFFLINE
	utils.WriteDevMem32(0x12030000, 0x00000004);

    //write priority select
    utils.WriteDevMem32(0x12030054, 0x55552366); // each module 4bit  cci       ---        ddrt  ---    ---     gzip   ---    ---
    utils.WriteDevMem32(0x12030058, 0x16556611); // each module 4bit  vicap1    hash       ive   aio    jpge    tde   vicap0  vdp
    utils.WriteDevMem32(0x1203005c, 0x43466445); // each module 4bit  mmc2      A17        fmc   sdio1  sdio0   A7    vpss0   vgs
    utils.WriteDevMem32(0x12030060, 0x56466666); // each module 4bit  gdc       usb3/pcie  vedu  usb2   cipher  dma2  dma1    gsf

    //read priority select
    utils.WriteDevMem32(0x12030064, 0x55552366); // each module 4bit  cci       ---        ddrt  ---    ---     gzip   ---    ---
    utils.WriteDevMem32(0x12030068, 0x06556600); // each module 4bit  vicap1    hash       ive   aio    jpge    tde   vicap0  vdp
    utils.WriteDevMem32(0x1203006c, 0x43466435); // each module 4bit  mmc2      A17        fmc   sdio1  sdio0   A7    vpss0   vgs
    utils.WriteDevMem32(0x12030070, 0x56266666); // each module 4bit  gdc       usb3/pcie  vedu  usb2   cipher  dma2  dma1    gsf

    utils.WriteDevMem32(0x120641f0, 0x1);  // use pri_map
    //write timeout select
    utils.WriteDevMem32(0x1206409c, 0x00000040);   
    utils.WriteDevMem32(0x120640a0, 0x00000000);    

    //read timeout select
    utils.WriteDevMem32(0x120640ac, 0x00000040);   // each module 8bit
    utils.WriteDevMem32(0x120640b0, 0x00000000);
    //#endif

	utils.WriteDevMem32(0x120300e0, 0xd); // internal codec: AIO MCLK out, CODEC AIO TX MCLK 

	koloader.LoadAll()

	//imx274)
    //tmp=0x11;
    utils.WriteDevMem32(0x12010040, 0x11);       //sensor0 clk_en, 72MHz
    //SDK config:     IVE:396M,  GDC:475M,  VGS:500M,  VEDU:600M,   VPSS:300M 
    //imx274:viu0: 600M,isp0:300M, viu1:300M,isp1:300M
    utils.WriteDevMem32(0x1201004c, 0x00094c23);
    utils.WriteDevMem32(0x12010054, 0x0004041);
    //spi0_4wire_pin_mux;
    //pinmux
    utils.WriteDevMem32(0x1204018c, 0x1); //  #SPI0_SCLK
    utils.WriteDevMem32(0x12040190, 0x1); // #SPI0_SD0
    utils.WriteDevMem32(0x12040194, 0x1); // #SPI0_SDI
    utils.WriteDevMem32(0x12040198, 0x1); // #SPI0_CSN

    //drive capability
    utils.WriteDevMem32(0x12040998, 0x150); // #SPI0_SCLK
    utils.WriteDevMem32(0x1204099c, 0x160); // #SPI0_SD0
    utils.WriteDevMem32(0x120409a0, 0x160); // #SPI0_SDI
    utils.WriteDevMem32(0x120409a4, 0x160); // #SPI0_CSN
    /////////////////////////////////////////////////
  	//single_pipe)
  	utils.WriteDevMem32(0x12040184, 0x1); //   # SENSOR0 HS from VI0 HS
  	utils.WriteDevMem32(0x12040188, 0x1); //   # SENSOR0 VS from VI0 VS
  	utils.WriteDevMem32(0x12040010, 0x2); //   # SENSOR1 HS from VI1 HS
	utils.WriteDevMem32(0x12040014, 0x2); //   # SENSOR1 VS from VI1 VS
  	/////
	utils.WriteDevMem32(0x12010044, 0x4ff0);
	utils.WriteDevMem32(0x12010044, 0x4);

}