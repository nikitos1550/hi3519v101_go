package cmos

//#include "cmos.h"
import "C"

import (
	"flag"
	"unsafe"
	"application/pkg/logger"
)

var (
	mode	uint
	testmod bool
)

func init() {
        flag.UintVar(&mode, "cmos-mode", 0, "CMOS mode") 
        flag.BoolVar(&testmod, "cmos-list-modes", false, "Show availible CMOS modes")
}

//func PrintInfo() {
//    fmt.Println("Here will be list of supported cmoses soon...")
//}

type cmosMode struct {
	width uint
	height uint
	fps	uint
	
	mipi	unsafe.Pointer
}

type busType uint

const (
   I2C		busType = 0
   Spi4Wire	busType = 1
)

type cmosControl struct {
	bus	busType
	busNum	uint
}

type cmos struct {
	vendor	string
	model	string

	modes   []cmosMode

	viDev 	unsafe.Pointer
        /*
                72 MHz
                37.125 MHz
                24 Mhz
        */
	clock   float32
        control cmosControl

	//SDK config:     IVE:396M,  GDC:475M,  VGS:500M,  VEDU:600M,   VPSS:300M 
        //      #os08a10:       viu0: 600M, isp0:300M, viu1:300M,isp1:300M
	hw	hwFreq

}

type hwFreq struct {
	//ive	uint
	//gdc	uint
	//vgs	uint
	//vedu	uint
	//vpss	uint

	viu	uint
	isp0	uint
	viui	uint
	isp1	uint
}

func Init() {
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

func Mipi() unsafe.Pointer {
	return cmosItem.modes[0].mipi
}

func ViDev() unsafe.Pointer {
	return cmosItem.viDev
}

func Width() uint {
        return cmosItem.modes[0].width
}

func Height() uint {
	return cmosItem.modes[0].height
}

func Fps() uint {
        return cmosItem.modes[0].fps
}

func Clock() float32 {
	return cmosItem.clock
}

func BusType() busType {
	return cmosItem.control.bus
}

func BusNum() uint{
	return cmosItem.control.busNum
}
