package cmos

//#include "cmos.h"
import "C"

import (
	"flag"
	"unsafe"
	"application/pkg/logger"
)

var (
    S * cmos
    //cmos        string
	mode	    uint
    //data        string
    //control     string
    //controlN    uint
)

func init() {
        //flag.StrVar(&cmos, "cmos", "unknown", "CMOS model") //TODO one package multiple CMOSes support
        //flag.StrVar(&data, "cmos-data", "LVDS", "CMOS data connection type [LVDS, MIPI, DC]")
        //flag.StrVar(&control, "cmos-control", "i2c", "CMOS control connection type [i2c, spi4wire]")
        //flag.UintVar(&controlN, "cmos-control-bus", 0, "CMOS control bus number")


    flag.UintVar(&mode, "cmos-mode", 0, "CMOS mode") 
}

type cmosWdr uint
const (
    WDRNone cmosWdr = 1
    WDR2TO1 cmosWdr = 2
)

type cmosMode struct {
	width int
	height int
	fps	int

    mipi	unsafe.Pointer
    viDev   unsafe.Pointer
    wdr     cmosWdr

    /*    
        72 MHz    
        37.125 MHz
        25 MHz
        24 Mhz
    */
    clock   float32
    //clock   uint //TODO
    
    //SDK config:     IVE:396M,  GDC:475M,  VGS:500M,  VEDU:600M,   VPSS:300M 
    //      #os08a10:       viu0: 600M, isp0:300M, viu1:300M,isp1:300M
    //hw  hwFreq

    description string
}

type busType uint

const (
   I2C		busType = 1
   SPI      busType = 2
)

type cmosControl struct {
	bus	busType
	busNum	uint
}

type cmosData   uint

const (
	LVDS	cmosData = 1
	MIPI	cmosData = 2
	DC	    cmosData = 3
)

type cmosBayer uint8
const (
    RGGB    cmosBayer   = 1
    GRBG    cmosBayer   = 2
    GBRG    cmosBayer   = 3
    BGGR    cmosBayer   = 4
)

type cmos struct {
	vendor	string
	model	string

    mode    uint
	modes   []cmosMode

    control cmosControl
	data	cmosData
    bayer   cmosBayer
}

type hwFreq struct {
	//ive	uint
	//gdc	uint
	//vgs	uint
	//vedu	uint
	//vpss	uint

	viu	    uint
	isp0	uint
	viui	uint
	isp1	uint
}

func Init() {
    //mode = 1
    if mode >= uint(len(cmosItem.modes)) {
        logger.Log.Fatal().
            Int("mode", int(mode)).
            Msg("Cmos mode not found")
    }

    S = &cmosItem
    S.mode = mode
}

func Register() {
    var errorCode C.int

    switch err := C.mpp_cmos_init(&errorCode); err {
    case C.ERR_NONE:
        logger.Log.Debug().
                Msg("C.mpp_cmos_init() ok")
//    case C.ERR_MPP:
//        logger.Log.Fatal().
//                Int("error", int(errorCode)).
//                Str("error_desc", error.Resolve(int64(errorCode))).
//                Msg("C.C.mpp_cmos_init() mpp error ")
    default:
            logger.Log.Fatal().
                Int("error", int(err)).
                Msg("Unexpected return of C.mpp_cmos_init()")
        }


}

func (c * cmos) Model() string {
    return c.model
}

func (c * cmos) Mipi() unsafe.Pointer {
	return c.modes[c.mode].mipi
}

func (c * cmos) ViDev() unsafe.Pointer {
	return c.modes[c.mode].viDev
}

func (c * cmos) Width() int {
    return c.modes[c.mode].width
}

func (c * cmos) Height() int {
	return c.modes[c.mode].height
}

func (c * cmos) Fps() int {
    return c.modes[c.mode].fps
}

func (c * cmos) Clock() float32 {
	return c.modes[c.mode].clock
}

func (c * cmos) BusType() busType {
	return c.control.bus
}

func (c * cmos) BusNum() uint{
	return c.control.busNum
}

func (c * cmos) Data() cmosData {
    return c.data
}

func (c * cmos) Bayer() cmosBayer {
    return c.bayer
}

func (c * cmos) Wdr() cmosWdr {
    return c.modes[c.mode].wdr

}
