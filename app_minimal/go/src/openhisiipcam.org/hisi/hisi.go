package hisi

import (
	"log"
    "flag"
)

//#include "../../../../libhisi/hisi_external.h"
//#cgo LDFLAGS: ${SRCDIR}/../../../../libhisi/libhisi.a
import "C"

////////////////////////////////////////////////////////////////////////////////

const (
	ERR_NONE            = 0
	ERR_OBJ_NOT_FOUND   = -1
	ERR_NOT_ALLOWED     = -2
	ERR_NOT_IMPLEMENTED = -3
)

////////////////////////////////////////////////////////////////////////////////


var flagCmos *int

func init() {
    flagCmos = flag.Int("cmos", 1, "CMOS id to init")
}

func MppInit() {
	log.Println("hisi MppInit start...")
    var captureParams C.struct_capture_params
	C.hisi_init(0, &captureParams) //hardcoded imx274
	log.Println("done")

	/////
	/*
	   var tmp C.uint

	   C.hisi_channels_max_num(&tmp)
	   chns.maxNum = uint(tmp)
	   C.hisi_channels_min_width(&tmp)
	   chns.minWidth = uint(tmp)
	   C.hisi_channels_max_width(&tmp)
	   chns.maxWidth = uint(tmp)
	   C.hisi_channels_min_height(&tmp)
	   chns.minHeight = uint(tmp)
	   C.hisi_channels_max_height(&tmp)
	   chns.maxHeight = uint(tmp)
	   C.hisi_channels_max_fps(&tmp)
	   chns.maxFps = uint(tmp)
	*/

	/*
	   chns.chn = make([]channel, chns.maxNum)
	   for i := 0; i < int(chns.maxNum); i++ {
	       chns.chn[i].id = uint(i)
	   }

	   C.hisi_encoders_max_num(&tmp)
	   encs.maxNum = uint(tmp)

	   encs.enc = make([]encoder, encs.maxNum)
	   for i := 0; i < int(encs.maxNum); i++ {
	       encs.enc[i].id = uint(i)
	   }

	   var err int

	   err = ChannelEnable(0, 3840, 2160, 30)
	   log.Println("enable first = ", err)
	*/

}

