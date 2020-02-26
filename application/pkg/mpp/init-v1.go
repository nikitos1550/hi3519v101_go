//+build hi3516cv100

package mpp

import (
	//"log"
	//"os"

	"application/pkg/ko"
	"application/pkg/utils"
	//"application/pkg/mpp/error"
)

func systemInit() {
	ko.UnloadAll()

	//#USB PHY
	utils.WriteDevMem32(0x20050080, 0x000121a8)
	//#USB PHY
	utils.WriteDevMem32(0x20050084, 0x005d2188)
	//NANDC 0x200300D0 [1:0]'b01
	utils.WriteDevMem32(0x200300D0, 0x5)
	//NANDC gpio
	utils.WriteDevMem32(0x200F00C8, 0x1)
	utils.WriteDevMem32(0x200f00cc, 0x1)
	utils.WriteDevMem32(0x200f00d0, 0x1)
	utils.WriteDevMem32(0x200f00d4, 0x1)
	utils.WriteDevMem32(0x200f00d8, 0x1)
	utils.WriteDevMem32(0x200f00dc, 0x1)
	utils.WriteDevMem32(0x200f00e0, 0x1)
	utils.WriteDevMem32(0x200f00e4, 0x1)
	utils.WriteDevMem32(0x200f00e8, 0x1)
	utils.WriteDevMem32(0x200f00ec, 0x1)
	utils.WriteDevMem32(0x200f00f4, 0x1)
	utils.WriteDevMem32(0x200f00f8, 0x1)
	//#�ر�SAR ADC ʱ��
	utils.WriteDevMem32(0x20030080, 0x1)
	//#�ر�SAR ADC
	utils.WriteDevMem32(0x200b0008, 0x1)
	//#��PWM
	utils.WriteDevMem32(0x20030038, 0x2)
	//#�ر�IR
	utils.WriteDevMem32(0x20070000, 0x0)
	//#IR �ܽŸ��ó�gpio
	utils.WriteDevMem32(0x200f00c4, 0x1)
	//#UART2��ʹ�ܣ�0x200A0000 [9][8][0]bit������Ϊ0
	utils.WriteDevMem32(0x200A0030, 0x0)
	//#UART2�ܽŸ��ó�gpio
	utils.WriteDevMem32(0x200f0108, 0x0)
	utils.WriteDevMem32(0x200f010c, 0x0)
	//#�ر�SPI0��SPI1
	utils.WriteDevMem32(0x200C0004, 0x7F00)
	utils.WriteDevMem32(0x200E0004, 0x7F00)
	//#spi0 �ܽŸ��ó�gpio
	utils.WriteDevMem32(0x200f000c, 0x0)
	utils.WriteDevMem32(0x200f0010, 0x0)
	utils.WriteDevMem32(0x200f0014, 0x0)
	//#spi1 �ܽŸ��ó�gpio
	utils.WriteDevMem32(0x200f0110, 0x0)
	utils.WriteDevMem32(0x200f0114, 0x0)
	utils.WriteDevMem32(0x200f0118, 0x0)
	utils.WriteDevMem32(0x200f011c, 0x0)
	//#AUDIO CODEC LINE IN �ر�������
	//utils.WriteDevMem32(0x20050068, 0xa8022c2c)
	//utils.WriteDevMem32(0x2005006c, 0xf5035a4a)

	//#RMII
	//net_rmii_mode()
	//{
	utils.WriteDevMem32(0x200f0030, 0x1)
	utils.WriteDevMem32(0x200f0034, 0x1)
	utils.WriteDevMem32(0x200f0038, 0x1)
	utils.WriteDevMem32(0x200f003C, 0x1)
	utils.WriteDevMem32(0x200f0040, 0x1)
	utils.WriteDevMem32(0x200f0044, 0x1)
	utils.WriteDevMem32(0x200f0048, 0x1)
	utils.WriteDevMem32(0x200f004C, 0x1)
	utils.WriteDevMem32(0x200f0050, 0x1)
	utils.WriteDevMem32(0x200f0054, 0x1)
	utils.WriteDevMem32(0x200f0058, 0x1)
	utils.WriteDevMem32(0x200f005C, 0x3)
	utils.WriteDevMem32(0x200f0060, 0x1)
	utils.WriteDevMem32(0x200f0064, 0x1)
	utils.WriteDevMem32(0x200f0068, 0x1) //  #MII_TXER 0x1,GPIO2_6 0x0
	utils.WriteDevMem32(0x200f006C, 0x1) //  #MII_RXER 0x1,GPIO2_7 0x0
	utils.WriteDevMem32(0x200f0070, 0x1)
	utils.WriteDevMem32(0x200f0074, 0x1)
	utils.WriteDevMem32(0x200f0078, 0x1)
	//}

	//#I2C default setting is I2C
	//i2c_type_select()
	//{
	utils.WriteDevMem32(0x200f0018, 0x00000001) // # 0:GPIO2_0   / 1:I2C_SDA
	utils.WriteDevMem32(0x200f001c, 0x00000001) // # 0:GPIO2_1   / 1:I2C_SCL
	//}

	//#SENSOR default setting is Sensor Clk
	//sensor_clock_select()
	//{
	utils.WriteDevMem32(0x200f0008, 0x00000001) // # 0:GPIO1_2   /1:SENSOR_CLK
	//}

	//open module clock while you need it!
	//clk_cfg()
	//{
	utils.WriteDevMem32(0x2003002c, 0x2a)  //          # VICAP, ISP unreset & clock enable
	utils.WriteDevMem32(0x20030048, 0x2)   //            # VPSS unreset, code also config
	utils.WriteDevMem32(0x20030034, 0x510) //  # VDP  unreset & HD clock enable
	utils.WriteDevMem32(0x20030040, 0x2)   //            # VEDU unreset
	utils.WriteDevMem32(0x20030060, 0x2)   //            # JPEG unreset
	utils.WriteDevMem32(0x20030058, 0x2)   //            # TDE  unreset
	utils.WriteDevMem32(0x20030068, 0x2)   //            # MDU  unreset

	ko.LoadAll()

	//imx104|imx122|imx138|imx225)
	utils.WriteDevMem32(0x200f000c, 0x1) //;              #pinmux SPI0
	utils.WriteDevMem32(0x200f0010, 0x1) //;              #pinmux SPI0
	utils.WriteDevMem32(0x200f0014, 0x1) //;              #pinmux SPI0
	utils.WriteDevMem32(0x20030030, 0x6) //;              #Sensor clock 37.125 MHz
	//insmod extdrv/ssp_sony.ko;;

	//# This is a sample, you should rewrite it according to your chip #
	//# mddrc pri&timeout setting
	utils.WriteDevMem32(0x20110150, 0x03ff6) //         #DMA1 DMA2
	utils.WriteDevMem32(0x20110154, 0x03ff6) //         #ETH
	utils.WriteDevMem32(0x20110158, 0x03ff6) //         #USB
	utils.WriteDevMem32(0x2011015C, 0x03ff6) //         #CIPHER   0x15C
	utils.WriteDevMem32(0x20110160, 0x03ff6) //         #SDIO   0X160
	utils.WriteDevMem32(0x20110164, 0x03ff6) //         #NAND SFC   0X164
	utils.WriteDevMem32(0x20110168, 0x10201) //         #ARMD  0X168
	utils.WriteDevMem32(0x2011016C, 0x10201) //         #ARMI  0X16C
	utils.WriteDevMem32(0x20110170, 0x03ff6) //         #IVE  0X170
	utils.WriteDevMem32(0x20110174, 0x03ff6) //                                 #MD, DDR_TEST  0x174
	utils.WriteDevMem32(0x20110178, 0x03ff6) //         #JPGE #0x178
	utils.WriteDevMem32(0x2011017C, 0x03ff3) //         #TDE0 0X17C
	utils.WriteDevMem32(0x20110180, 0x03ff4) //         #VPSS  0X180
	utils.WriteDevMem32(0x20110184, 0x10c82) //         #VENC  0X184
	utils.WriteDevMem32(0x20110188, 0x10101) //         #VICAP FPGA
	//#himm 0x20110188 0x10640         #VICAP ESL
	utils.WriteDevMem32(0x2011018c, 0x10100) //         #VDP
	utils.WriteDevMem32(0x20110100, 0x67)    //            #mddrc order enable mddrc idmap mode select

	//#himm 0x20050054 0x123564        #[2:0] VENC [6:4] VPSS [10:8] TDE [14:12] JPGE [18:16] MD

	//#himm 0x200500d8 0x3             #DDR0Ö»Ê¹ÄÜVICAPºÍVDPÂÒÐò
	utils.WriteDevMem32(0x20050038, 0x3) //             #DDR0Ö»Ê¹ÄÜVICAPºÍVDPÂÒÐò

}
