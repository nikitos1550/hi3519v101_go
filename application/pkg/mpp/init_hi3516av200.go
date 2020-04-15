//+build arm
//+build hi3516av200

package mpp

/*
#include "./include/mpp_v3.h"

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
    "application/pkg/logger"
    "os"

	"application/pkg/ko"
    "application/pkg/utils"
    "application/pkg/mpp/error"

    "application/pkg/mpp/cmos"
    "application/pkg/utils/regs"
)

const (
	DDRMemStartAddr = 0x80000000
)

//TODO rework this mess
func systemInit(devInfo DeviceInfo) {

    //NOTE should be done, otherwise there can be kernel panic on module unload
    //ATTENTION maybe isp exit should be added as well
    //TODO rework, add error codes, deal with C includes

    if _, err := os.Stat("/dev/sys"); err == nil { 
        var errorCode C.int
        switch err := C.mpp3_sys_exit(&errorCode); err {
        case C.ERR_NONE:
            //log.Println("C.mpp3_sys_exit() ok")
	    logger.Log.Debug().
	    	Msg("C.mpp3_sys_exit() ok")
        case C.ERR_MPP:
            //log.Fatal("C.mpp3_sys_exit() HI_MPI_SYS_Exit() error ", error.Resolve(int64(errorCode)))
	    logger.Log.Fatal().
	    	Str("func", "HI_MPI_SYS_Exit()").
		Int("error", int(errorCode)).
		Str("error_desc", error.Resolve(int64(errorCode))).
		Msg("C.mpp3_sys_exit() error")
        default:
            //log.Fatal("Unexpected return ", err , " of C.mpp3_sys_exit()")
	    logger.Log.Fatal().
	    	Int("error", int(err)).
		Msg("Unexpected return of C.mpp3_sys_exit()")
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
            //log.Println("C.mpp3_vb_exit() ok")
	    logger.Log.Debug().
	    	Msg("C.mpp3_vb_exit() ok")
        case C.ERR_MPP:
            //log.Fatal("C.mpp3_vb_exit() HI_MPI_VB_Exit() error ", error.Resolve(int64(errorCode))) 
	    logger.Log.Fatal().
	    	Str("func", "HI_MPI_VB_Exit()").
	    	Int("error", int(errorCode)).
		Str("error_desc", error.Resolve(int64(errorCode))).
		Msg("C.mpp3_vb_exit() error")
        default:
            //log.Fatal("Unexpected return ", err , " of C.mpp3_vb_exit()")
	    logger.Log.Fatal().
	    		Int("error", int(err)).
			Msg("Unexpected return of C.mpp3_vb_exit()")
        }
    }
	//delete_module("hi3519v101_isp", NULL);//ATTENTION THIS IS NEED FOR PROPER APP RERUN, also some info here
    //http://bbs.ebaina.com/forum.php?mod=viewthread&tid=13925&extra=&highlight=run%2Bae%2Blib%2Berr%2B0xffffffff%21&page=1
    ko.UnloadAll()

	//sensor0 pinmux
	//utils.WriteDevMem32(0x1204017c, 0x1);  //#SENSOR0_CLK
    regs.ByNameConst(regs.MUXCTRL_REG95).Set(0x1)

	//utils.WriteDevMem32(0x12040180, 0x0);  //#SENSOR0_RSTN
    regs.ByNameStr("muxctrl_reg96").Set(0x0)
    
    //utils.WriteDevMem32(0x12040184, 0x1);  //#SENSOR0_HS,from vi0
    regs.ByNameStr("muxctrl_reg97").Set(0x1)

	//utils.WriteDevMem32(0x12040188, 0x1);  //#SENSOR0_VS,from vi0
    regs.ByNameStr("muxctrl_reg98").Set(0x1)

	//sensor0 drive capability
	//utils.WriteDevMem32(0x12040988, 0x150);  //#SENSOR0_CLK
    regs.ByNameStr("pad_ctrl_reg98").Set(0x150)

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
    logger.Log.Trace().
            Bool("mode", devInfo.ViVpssOnline).
            Msg("VI-VPSS mode")

    if devInfo.ViVpssOnline == true {

	    utils.WriteDevMem32(0x12030000, 0x00000204);

	    //write priority select
	    utils.WriteDevMem32(0x12030054, 0x55552356); //  each module 4bit  cci       ---        ddrt  ---    ---     gzip   ---    ---
	    utils.WriteDevMem32(0x12030058, 0x16554411); // each module 4bit  vicap1    hash       ive   aio    jpge    tde   vicap0  vdp 
	    utils.WriteDevMem32(0x1203005c, 0x33466314); // each module 4bit  mmc2      A17        fmc   sdio1  sdio0   A7    vpss0   vgs 
	    utils.WriteDevMem32(0x12030060, 0x46266666); // each module 4bit  gdc       usb3/pcie  vedu  usb2   cipher  dma2  dma1    gsf

	    //read priority select
	    utils.WriteDevMem32(0x12030064, 0x55552356); // each module 4bit  cci       ---         ddrt  ---    ---     gzip   ---    ---
	    utils.WriteDevMem32(0x12030068, 0x06554401); // each module 4bit  vicap1    hash        ive   aio    jpge    tde   vicap0  vdp
	    utils.WriteDevMem32(0x1203006c, 0x33466304); // each module 4bit  mmc2      A17         fmc   sdio1  sdio0    A7   vpss0   vgs 
	    utils.WriteDevMem32(0x12030070, 0x46266666); // each module 4bit  gdc       usb3/pcie   vedu  usb2   cipher  dma2  dma1    gsf

	    utils.WriteDevMem32(0x120641f0, 0x1);  //use pri_map
	    //write timeout select
	    utils.WriteDevMem32(0x1206409c, 0x00000040);
	    utils.WriteDevMem32(0x120640a0, 0x00000000);

	    //read timeout select
	    utils.WriteDevMem32(0x120640ac, 0x00000040);
	    utils.WriteDevMem32(0x120640b0, 0x00000000);
    } else {

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
    }

	utils.WriteDevMem32(0x120300e0, 0xd); // internal codec: AIO MCLK out, CODEC AIO TX MCLK 

    //KO
    ko.Params.Add("mem_start_addr").Str("0x").Uint64Hex(DDRMemStartAddr + devInfo.MemLinuxSize)
    ko.Params.Add("mem_mpp_size").Uint64(devInfo.MemMppSize/(1024*1024)).Str("M")
    ko.Params.Add("mem_total_size").Uint64(devInfo.MemTotalSize/(1024*1024))
    ko.Params.Add("vi_vpss_online").Bool(devInfo.ViVpssOnline)
    ko.Params.Add("cmos").Str(cmos.Model())
	ko.Params.Add("proc_param").Uint64(30)
    ko.Params.Add("detect_err_frame").Uint64(10)
    ko.Params.Add("save_power").Uint64(1)

	ko.LoadAll()


		switch cmos.Model() {
            case "imx326":
                utils.WriteDevMem32(0x1201004c, 0x00094c21)
                utils.WriteDevMem32(0x12010054, 0x0004041)
                utils.WriteDevMem32(0x12010040, 0x14)
			case "imx274":
                utils.WriteDevMem32(0x1201004c, 0x00094c23)
                utils.WriteDevMem32(0x12010054, 0x0004041)
			case "imx226":
                utils.WriteDevMem32(0x1201004c, 0x00094c23)
                utils.WriteDevMem32(0x12010054, 0x0004041)
			default:
                logger.Log.Fatal().
                    Str("name", cmos.Model()).
                    Msg("CMOS is not supported")
		}


		switch cmos.Clock() {
			case 24:
				utils.WriteDevMem32(0x12010040, 0x14) //           # sensor0 clk_en, 24MHz
			case 72:
				utils.WriteDevMem32(0x12010040, 0x11);       //sensor0 clk_en, 72MHz
			default:
				logger.Log.Fatal().
					Float32("clock", cmos.Clock()).
					Msg("CMOS clock is not supported")
		}

		switch cmos.BusType() {
		case cmos.I2C:
			if cmos.BusNum() == 0 {
				//i2c0_pin_mux()
        			utils.WriteDevMem32(0x12040190, 0x2)    //;  #I2C0_SDA
        			utils.WriteDevMem32(0x1204018c, 0x2)    //;  #I2C0_SCL
    
        			//#drive capability
        			utils.WriteDevMem32(0x1204099c, 0x120)  //; #I2C0_SDA
        			utils.WriteDevMem32(0x12040998, 0x120)  //; #I2C0_SCL    

			} else {
				logger.Log.Fatal().
                    Uint("bus", cmos.BusNum()).
                    Msg("CMOS bus num not supported")
			}
		case cmos.Spi4Wire:
			if cmos.BusNum() == 0 {
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

			} else {
                logger.Log.Fatal().
                    Uint("bus", cmos.BusNum()).
                    Msg("CMOS bus num not supported")
			}
		default:
			logger.Log.Fatal().
				Int("type", int(cmos.BusType())).
				Msg("unrecognized cmos bus type")
	}

    if false {
	//imx274)
    //tmp=0x11;
    //utils.WriteDevMem32(0x12010040, 0x11);       //sensor0 clk_en, 72MHz
    //SDK config:     IVE:396M,  GDC:475M,  VGS:500M,  VEDU:600M,   VPSS:300M 
    //imx274:viu0: 600M,isp0:300M, viu1:300M,isp1:300M
    //utils.WriteDevMem32(0x1201004c, 0x00094c23);
    //utils.WriteDevMem32(0x12010054, 0x0004041);
    //spi0_4wire_pin_mux;
    //pinmux
    //utils.WriteDevMem32(0x1204018c, 0x1); //  #SPI0_SCLK
    //utils.WriteDevMem32(0x12040190, 0x1); // #SPI0_SD0
    //utils.WriteDevMem32(0x12040194, 0x1); // #SPI0_SDI
    //utils.WriteDevMem32(0x12040198, 0x1); // #SPI0_CSN

    //drive capability
    //utils.WriteDevMem32(0x12040998, 0x150); // #SPI0_SCLK
    //utils.WriteDevMem32(0x1204099c, 0x160); // #SPI0_SD0
    //utils.WriteDevMem32(0x120409a0, 0x160); // #SPI0_SDI
    //utils.WriteDevMem32(0x120409a4, 0x160); // #SPI0_CSN

    } else {

	    //os05a)
            //tmp=0x14;
            //            # SDK config:     IVE:396M,  GDC:475M,  VGS:500M,  VEDU:600M,   VPSS:300M 
            //            himm 0x1201004c 0x00094c21;
            //            himm 0x12010054 0x0004041;
            //himm 0x12010040 0x14;           # sensor0 clk_en, 24MHz
            //i2c0_pin_mux;


	    //os08a10)
            //tmp=0x14;
            //            # SDK config:     IVE:396M,  GDC:475M,  VGS:500M,  VEDU:600M,   VPSS:300M 
            //            #os08a10:       viu0: 600M, isp0:300M, viu1:300M,isp1:300M
            //            himm 0x1201004c 0x00094c23;
            //            himm 0x12010054 0x0004041;
            //himm 0x12010040 0x14;           # sensor0 clk_en, 24MHz
            //i2c0_pin_mux;


    	//utils.WriteDevMem32(0x1201004c, 0x00094c23)
    	//utils.WriteDevMem32(0x12010054, 0x0004041)
    	//utils.WriteDevMem32(0x12010040, 0x14) //           # sensor0 clk_en, 24MHz


	//i2c0_pin_mux()
	//{
    	//#pinmux
    	//utils.WriteDevMem32(0x12040190, 0x2)	//;  #I2C0_SDA
    	//utils.WriteDevMem32(0x1204018c, 0x2)	//;  #I2C0_SCL
    
    	//#drive capability
    	//utils.WriteDevMem32(0x1204099c, 0x120)	//; #I2C0_SDA
    	//utils.WriteDevMem32(0x12040998, 0x120)	//; #I2C0_SCL    
	//}

    }

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
