//+build arm
//+build hi3516av100

package mpp

import (
	"application/pkg/ko"
	"application/pkg/utils"

    "application/pkg/logger"

     "application/pkg/mpp/cmos"
)

const (
    DDRMemStartAddr = 0x80000000
)

func systemInit(devInfo DeviceInfo) {
	ko.UnloadAll()

	//i2c_pin_mux()
	//{
	utils.WriteDevMem32(0x200f0070, 0x1) //;		# i2c2_sda
	utils.WriteDevMem32(0x200f0074, 0x1) //;		# i2c2_scl
	//}

	//vicap_pin_mux()
	//{
	utils.WriteDevMem32(0x200f0000, 0x1) //		# 0: GPIO0_5, 		1: SENSOR_CLK
	utils.WriteDevMem32(0x200f0004, 0x1) //		# 1£ºFLASH_TRIG,	0: GPIO0_6,		2£ºSPI1_CSN1
	utils.WriteDevMem32(0x200f0008, 0x1) //		# 1£ºSHUTTER_TRIG,	0£ºGPIO0_7,		2£ºSPI1_CSN2
	//}

	//# open module clock while you need it!
	//clk_cfg()
	//{
    if devInfo.Chip == "hi3516av100" {
    utils.WriteDevMem32(0x20030030, 0x00004025) //      # AVC-300M VGS-300M VPSS-250M VEDU-300M mda1axi 250M mda0axi 300M DDR-250
	//#himm 0x20030030 0x00004005      # AVC-300M VGS-300M VPSS-250M VEDU-300M mda1axi 250M mda0axi 300M DDR-250
	utils.WriteDevMem32(0x20030104, 0x3)     //             # VICAP-198M VPSS-198M
	utils.WriteDevMem32(0x2003002c, 0x90007) //         # VICAP-250M, ISP unreset & clk en, Sensor clk en-37.125M, clk reverse
	//#himm 0x20030034 0xffc           # VDP-1080p@60fps unreset & clk en
	//#himm 0x20030034 0xef74          #VDP-PAL/NTSC
	utils.WriteDevMem32(0x20030040, 0x2002) //          # VEDU0 AVC unreset & clk en
	utils.WriteDevMem32(0x20030048, 0x2)    //             # VPSS0 unreset & clk en

	utils.WriteDevMem32(0x20030058, 0x2) //             # TDE unreset & clk en
	utils.WriteDevMem32(0x2003005c, 0x2) //             # VGS unreset & clk en
	utils.WriteDevMem32(0x20030060, 0x2) //             # JPGE  unreset & clk en

	utils.WriteDevMem32(0x20030068, 0x2) //             # MDU unreset & clk en
	utils.WriteDevMem32(0x2003006c, 0x2) //             # IVE-300MHz unreset & clk en
	//#himm 0x20030070 0x2            # VOIE unreset & clk en

	utils.WriteDevMem32(0x2003007c, 0x2) //             # cipher unreset & clk en
	utils.WriteDevMem32(0x2003008c, 0xe) //             # aio MCLK PLL 1188M, unreset & clk en
	//#himm 0x200300d8 0xa;           # ddrt

	//# USB not set
	//# SDIO not set
	//# SFC not set
	//# NAND not set
	//# RTC use external clk
	//# PWM not set
	//# DMAC not set
	//# SPI not set
	//# I2C not set
	//# SENSE CLK not set
	//# WDG not set

	//echo "clock configure operation done!"
	}
    if devInfo.Chip == "hi3516dv100" {

    utils.WriteDevMem32(0x20030030, 0x00004005)//      # AVC-300M VGS-300M VPSS-250M VEDU-300M mda1axi 250M mda0axi 300M DDR-250
    utils.WriteDevMem32(0x20030104, 0x3)//             # VICAP-198M VPSS-198M
    utils.WriteDevMem32(0x2003002c, 0x90007)//         # VICAP-250M, ISP unreset & clk en, Sensor clk en-37.125M, clk reverse
    //#himm 0x20030034 0xffc           # VDP-1080p@60fps unreset & clk en
    //#himm 0x20030034 0xef74          #VDP-PAL/NTSC 
    utils.WriteDevMem32(0x20030040, 0x2002)//          # VEDU0 AVC unreset & clk en
    utils.WriteDevMem32(0x20030048, 0x2)//             # VPSS0 unreset & clk en

    utils.WriteDevMem32(0x20030058, 0x2)//             # TDE unreset & clk en
    utils.WriteDevMem32(0x2003005c, 0x2)//             # VGS unreset & clk en
    utils.WriteDevMem32(0x20030060, 0x2)//             # JPGE  unreset & clk en

    utils.WriteDevMem32(0x20030068, 0x2)//             # MDU unreset & clk en
    utils.WriteDevMem32(0x2003006c, 0x2)//             # IVE-300MHz unreset & clk en
    //#himm 0x20030070 0x2            # VOIE unreset & clk en

    utils.WriteDevMem32(0x2003007c, 0x2)//             # cipher unreset & clk en
    utils.WriteDevMem32(0x2003008c, 0xe)//             # aio MCLK PLL 1188M, unreset & clk en

    }


    logger.Log.Trace().
        Bool("mode", devInfo.ViVpssOnline).
        Msg("VI-VPSS mode")

	if devInfo.ViVpssOnline == true {

	    utils.WriteDevMem32(0x20120004, 0x40000000) //;			# online, SPI1 CS0
	    //#pri config
	    utils.WriteDevMem32(0x20120058, 0x26666400) //			# each module 4bit£ºvedu       ddrt_md  ive  aio    jpge    tde   vicap  vdp
	    utils.WriteDevMem32(0x2012005c, 0x66666103) //			# each module 4bit£ºsfc_nand   sfc_nor  nfc  sdio1  sdio0   a7    vpss   vgs
	    utils.WriteDevMem32(0x20120060, 0x66266666) //			# each module 4bit£ºreserve    reserve  avc  usb    cipher  dma2  dma1   gsf

	    //#timeout config
	    utils.WriteDevMem32(0x20120064, 0x00000011) //			# each module 4bit£ºvedu       ddrt_md  ive  aio    jpge    tde   vicap  vdp
	    utils.WriteDevMem32(0x20120068, 0x00000020) //			# each module 4bit£ºsfc_nand   sfc_nor  nfc  sdio1  sdio0   a7    vpss   vgs
	    utils.WriteDevMem32(0x2012006c, 0x00000000) //			# each module 4bit£ºreserve    reserve  avc  usb    cipher  dma2  dma1   gsf
	} else {
	
	    utils.WriteDevMem32(0x20120004, 0x0)//;			    # offline, mipi SPI1 CS0;
		//	# pri config
		utils.WriteDevMem32(0x20120058, 0x26666400)//     		# each module 4bit£ºvedu       ddrt_md  ive  aio    jpge    tde   vicap  vdp
		utils.WriteDevMem32(0x2012005c, 0x66666112)//     		# each module 4bit£ºsfc_nand   sfc_nor  nfc  sdio1  sdio0   a7    vpss   vgs
		utils.WriteDevMem32(0x20120060, 0x66266666)//    		# each module 4bit£ºreserve    reserve  avc  usb    cipher  dma2  dma1   gsf
		//# timeout config
		utils.WriteDevMem32(0x20120064, 0x00000011)//    		# each module 4bit£ºvedu       ddrt_md  ive  aio    jpge    tde   vicap  vdp
		utils.WriteDevMem32(0x20120068, 0x00000000)//    		# each module 4bit£ºsfc_nand   sfc_nor  nfc  sdio1  sdio0   a7    vpss   vgs
		utils.WriteDevMem32(0x2012006c, 0x00000000)//    		# each module 4bit£ºreserve    reserve  avc  usb    cipher  dma2  dma1   gsf
	}
	
	//}
	/* //insert_sns after ko load
	//imx178)
	utils.WriteDevMem32(0x200f0050, 0x2)     //;                # i2c0_scl
	utils.WriteDevMem32(0x200f0054, 0x2)     //;                # i2c0_sda
	utils.WriteDevMem32(0x2003002c, 0xF0007) //             # sensor unreset, clk 25MHz, VI 250MHz
	//utils.WriteDevMem32(0x2003002c 0x90007) //            # sensor unreset, clk 37.125MHz, VI 250MHz
	*/

    ko.Params.Add("mem_start_addr").Str("0x").Uint64Hex(DDRMemStartAddr + devInfo.MemLinuxSize)
    ko.Params.Add("mem_mpp_size").Uint64(devInfo.MemMppSize/(1024*1024)).Str("M")
    ko.Params.Add("mem_total_size").Uint64(devInfo.MemTotalSize/(1024*1024))
    ko.Params.Add("vi_vpss_online").Bool(devInfo.ViVpssOnline)
    ko.Params.Add("cmos").Str(cmos.S.Model())
    ko.Params.Add("detect_err_frame").Uint64(10)

	ko.LoadAll()

        //imx178)
        switch cmos.S.Clock() {
            case 24:
                utils.WriteDevMem32(0x2003002c, 0xE0007)             //# sensor unreset, clk 24MHz, VI 250MHz
            case 25:
                utils.WriteDevMem32(0x2003002c, 0xF0007) //           #sensor unreset, clk 25MHz, VI 250MHz
            //case 27:
            //     himm 0x2003002c 0xB0007             # sensor unreset, clk 27MHz, VI 250MHz
            case 37.125:
                utils.WriteDevMem32(0x2003002c, 0x90007)       //sensor unreset, clk 37.125MHz, VI 250MHz
            default:
                logger.Log.Fatal().
                    Float32("clock", cmos.S.Clock()).
                    Msg("CMOS clock is not supported")
        }

        switch cmos.S.BusType() {
            case cmos.SPI:
            if cmos.S.BusNum() == 0 {
                utils.WriteDevMem32(0x200f0050, 0x1)//                # spi0_sclk
                utils.WriteDevMem32(0x200f0054, 0x1)//                # spi0_sdo
                utils.WriteDevMem32(0x200f0058, 0x1)//                # spi0_sdi
                utils.WriteDevMem32(0x200f005c, 0x1)//                # spi0_csn

            } else {
                    logger.Log.Fatal().
                        Uint("bus", cmos.S.BusNum()).
                        Msg("CMOS bus num not supported")
            }
            case cmos.I2C:
                if cmos.S.BusNum() == 0 {
                    utils.WriteDevMem32(0x200f0050, 0x2)     //;                # i2c0_scl
                    utils.WriteDevMem32(0x200f0054, 0x2)     //;                # i2c0_sda
                } else {
                    logger.Log.Fatal().
                        Uint("bus", cmos.S.BusNum()).
                        Msg("CMOS bus num not supported")
            }
        default:
            logger.Log.Fatal().
                Int("type", int(cmos.S.BusType())).
                Msg("unrecognized cmos bus typy")
    }

            switch cmos.S.Model() {
            case "imx290_lvds":
            case "imx178":
            case "ov4689":
                utils.WriteDevMem32(0x20030104, 0x0)
            default:
                logger.Log.Fatal().
                    Str("name", cmos.S.Model()).
                    Msg("CMOS is not supported")
        }

}
