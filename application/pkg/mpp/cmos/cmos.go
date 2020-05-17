package cmos

//#include "cmos.h"
import "C"

import (
	"flag"
	"unsafe"
    "strconv"

	"application/pkg/logger"
    "application/pkg/buildinfo"
)

var (
    S * cmos
    cmosNumber  int
    //cmos        string
	mode	    uint

    //data        string
    //control     string
    //controlN    uint

    cmosName        string
    cmosControl2    string
    cmosControlBus  uint
    cmosData2       string
    cmosLanes       string


    //Description from hi3516cv300
    //The VI device Dev0 supports a maximum of 16-bit inputs, and the 16-bit data lines are
    //used to connect to external video sources. In the actual application scenario, only some
    //data lines are used, which depends on hardware connection. For example, the camera
    //connects to the VI channel by using the first 12-bit data lines (data lines corresponding to
    //bit 0 to bit 11). In this case, the component type of the input data is set to single-
    //component, and the component mask is set to 0xFFF00000.
    dcZeroBitOffset   int

    i2cAddr         int
)

func init() {
    if buildinfo.Family != "hi3516cv100" {
        var defaultCmosLanes string

        for i := 0; i < (C.LVDS_LANE_NUM-1); i++ {  //TODO define custom lanes num var and set 0 for hi3516cv100 in case to successful compile
            defaultCmosLanes = defaultCmosLanes + strconv.Itoa(i) + ","
	    }
        defaultCmosLanes = defaultCmosLanes + strconv.Itoa(C.LVDS_LANE_NUM-1)

        flag.StringVar(&cmosData2, "cmos-data", "dc", "CMOS data connection type [[sub]LVDS, MIPI, DC], non case sensitive.")
        flag.StringVar(&cmosLanes, "cmos-lanes", defaultCmosLanes, "CMOS-SoC lanes connection in case of [sub]LVDS/MIPI, comma separated values, where index is CMOS lane num, value is SoC lane num.")
    }

    flag.StringVar(&cmosName, "cmos", "unknown", "CMOS model")
    
    flag.StringVar(&cmosControl2, "cmos-control", "i2c", "CMOS control connection type [i2c, spi], non case sensitive.")
    flag.UintVar(&cmosControlBus, "cmos-control-bus", 0, "CMOS control bus number.")

    flag.IntVar(&dcZeroBitOffset, "cmos-dc-connection-offset", -1, "Parallel DC connection offset, -1 means use typical.")

    flag.UintVar(&mode, "cmos-mode", 0, "CMOS mode")

    flag.IntVar(&i2cAddr, "cmos-i2c-addr", -1, "Use non standard i2c cmos address.")
}

type cmosWdr uint
const (
    WDRNone     cmosWdr = 1
    WDR2TO1     cmosWdr = 2 //TODO rename to LINE
    WDR2TO1F    cmosWdr = 3 //FRAME
    WDR2TO1FFR  cmosWdr = 4 //WDR_MODE_2To1_FRAME_FULL_RATE 
)

type crop struct {
    X0  int
    Y0  int
    Width int
    Height int
}

type cmosMode struct {
    mipiCrop crop
    viCrop crop
    ispCrop crop
	width int
	height int

	fps	int
    bitness int

    lanes []int

    //mipi	unsafe.Pointer  //will be removed, changed to mipi config
    mipiLVDSAttr unsafe.Pointer
    mipiMIPIAttr unsafe.Pointer

    wdr     cmosWdr

    dcSync  dcSyncAttr

    //pointer to driver struct, will be invoked as sensor register
    /*  TIP how to invoke C function pointer from golang space
    // typedef int (*intFunc) ();
    //
    // int
    // bridge_int_func(intFunc f)
    // {
    //		return f();
    // }
    //
    // int fortytwo()
    // {
    //	    return 42;
    // }
    import "C"
    import "fmt"

    func main() {
    	f := C.intFunc(C.fortytwo)
    	fmt.Println(int(C.bridge_int_func(f)))
	    // Output: 42
    }
    */
    //all cmos drivers (for early families) will be converted to static funcs, only sensor_register_NAME will be exported
    //we will point to that funcs, such func will be without params (bus num support will be added later, now only 0).
    //cmos.h will containt prototypes like ./hi_sns_ctrl.h:34:int  sensor_register_callback(void);
    //families that expose ISP object will have C code that do all job

    //control interface, even we know case where only interface changed, now we will use separate cmos modes for different interfaces

    control cmosControl
    data    cmosData

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
    SubLVDS cmosData = 2
	MIPI	cmosData = 3
	DC	    cmosData = 4
    HISPI   cmosData = 5
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

    mode    uint        //sould be moved to globals
	modes   []cmosMode

    control cmosControl //should be moved to cmos mode
	data	cmosData    //should be moved to cmos mode

    dcZeroBitOffset   uint //valid onlt for DC data mode

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

    
    for i:=0; i <= len(cmosItems); i++ {
        if i == len(cmosItems) {
            logger.Log.Fatal().
                Str("model", cmosName).
                Msg("CMOS not found")
        }
        if cmosItems[i].model == cmosName {
            cmosNumber = i
            S = &cmosItems[i]
            logger.Log.Debug().
                Str("model", S.model).
                Str("vendor", S.vendor).
                Msg("CMOS found")
            break
        }
    }
    

    //S = &cmosItem

    //mode = 1
    if mode >= uint(len(S.modes)) {
        logger.Log.Fatal().
            Int("mode", int(mode)).
            Msg("Cmos mode not found")
    }

    logger.Log.Debug().
        Str("description", S.modes[mode].description).
        Msg("CMOS mode found")

    S.mode = mode

    if dcZeroBitOffset != -1 {
        logger.Log.Fatal().
            Int("value", dcZeroBitOffset).
            Msg("Now only typical DC connection offset is supported.")
    }else{
        dcZeroBitOffset = int(S.dcZeroBitOffset)
    }
}

func Register() {
    var errorCode C.int

    switch err := C.mpp_cmos_init(&errorCode, C.uchar(cmosNumber)); err {
    //switch err := C.mpp_cmos_init(&errorCode); err {

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

func DCZeroBitOffset() int {
    return dcZeroBitOffset
}

func (c *cmos) ViCrop() crop {
    return c.modes[c.mode].viCrop
}

func (c *cmos) IspCrop() crop {
    return c.modes[c.mode].ispCrop
}

func (c * cmos) DCSYNC() dcSyncAttr {
    return c.modes[c.mode].dcSync
}

func (c * cmos) Model() string {
    return c.model
}

//func (c * cmos) Mipi() unsafe.Pointer {
//	return c.modes[c.mode].mipi
//}

func (c * cmos) MipiLVDSAttr() unsafe.Pointer {
  return c.modes[c.mode].mipiLVDSAttr
}

func (c * cmos) MipiMIPIAttr() unsafe.Pointer {
  return c.modes[c.mode].mipiMIPIAttr
}

//func (c * cmos) ViDev() unsafe.Pointer {
//	return c.modes[c.mode].viDev
//}

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

func (c * cmos) Bitness() int {
    return c.modes[c.mode].bitness
}

