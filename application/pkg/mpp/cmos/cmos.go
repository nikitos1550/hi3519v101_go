package cmos

//#include "cmos.h"
import "C"

import (
	"flag"
	"unsafe"
	"application/pkg/logger"
)

var (
    //cmos        string
	mode	    uint
    //data        string
    //control     string
    //controlN    uint
	testmod     bool
)

func init() {
        //flag.StrVar(&cmos, "cmos", "unknown", "CMOS model") //TODO one package multiple CMOSes support

        flag.UintVar(&mode, "cmos-mode", 0, "CMOS mode") 
        //flag.StrVar(&data, "cmos-data", "LVDS", "CMOS data connection type [LVDS, MIPI, DC]")
        //flag.StrVar(&control, "cmos-control", "i2c", "CMOS control connection type [i2c, spi4wire]")
        //flag.UintVar(&controlN, "cmos-control-bus", 0, "CMOS control bus number")

        flag.BoolVar(&testmod, "cmos-info", false, "Show supported CMOS modes")
}

type cmosWdr struct {
    enabled bool
}

type cmosMode struct {
	width uint
	height uint
	fps	uint

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
    //SDK config:     IVE:396M,  GDC:475M,  VGS:500M,  VEDU:600M,   VPSS:300M 
    //      #os08a10:       viu0: 600M, isp0:300M, viu1:300M,isp1:300M
    hw  hwFreq

    description string
}

type busType uint

const (
   I2C		busType = 1
   Spi4Wire	busType = 2
   SPI      busType = 3
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

type cmos struct {
	vendor	string
	model	string

	modes   []cmosMode

    control cmosControl
	data	cmosData
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

    if mode >= uint(len(cmosItem.modes)) {
        logger.Log.Fatal().
            Int("mode", int(mode)).
            Msg("Cmod mode not found")
    }
//}
//func Setup() {
	var errorCode C.int
	//err := C.mpp_cmos_init(&errorCode)

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

func Model() string {
    return cmosItem.model
}

func Mipi() unsafe.Pointer {
	return cmosItem.modes[mode].mipi
}

func ViDev() unsafe.Pointer {
	return cmosItem.modes[mode].viDev
}

func Width() uint {
    return cmosItem.modes[mode].width
}

func Height() uint {
	return cmosItem.modes[mode].height
}

func Fps() uint {
    return cmosItem.modes[mode].fps
}

func Clock() float32 {
	return cmosItem.modes[mode].clock
}

func BusType() busType {
	return cmosItem.control.bus
}

func BusNum() uint{
	return cmosItem.control.busNum
}

func Data() cmosData {
    return cmosItem.data
}
