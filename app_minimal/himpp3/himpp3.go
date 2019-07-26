package himpp3

import (
	"fmt"
//	"unsafe"
//	"syscall"
	"../external/github.com/creack/goselect"
)

// #include "himpp3_external.h"
// #include "himpp3_mpp_includes.h"
// #cgo LDFLAGS: ${SRCDIR}/libhimpp3.a
// #cgo LDFLAGS: ${SRCDIR}/../../mpp_hi3519_v101/lib/libsns_imx274.a
// #cgo LDFLAGS: ${SRCDIR}/../../mpp_hi3519_v101/lib/libisp.a
// #cgo LDFLAGS: ${SRCDIR}/../../mpp_hi3519_v101/lib/libmpi.a
// #cgo LDFLAGS: ${SRCDIR}/../../mpp_hi3519_v101/lib/libVoiceEngine.a
// #cgo LDFLAGS: ${SRCDIR}/../../mpp_hi3519_v101/lib/lib_hiae.a
// #cgo LDFLAGS: ${SRCDIR}/../../mpp_hi3519_v101/lib/lib_hiawb.a
// #cgo LDFLAGS: ${SRCDIR}/../../mpp_hi3519_v101/lib/lib_hiaf.a
// #cgo LDFLAGS: ${SRCDIR}/../../mpp_hi3519_v101/lib/libupvqe.a
// #cgo LDFLAGS: ${SRCDIR}/../../mpp_hi3519_v101/lib/libdnvqe.a
// #cgo LDFLAGS: ${SRCDIR}/../../mpp_hi3519_v101/lib/lib_hidefog.a
// #cgo CFLAGS: -mcpu=cortex-a7 -mfloat-abi=softfp -mfpu=neon-vfpv4 -mno-unaligned-access -fno-aggressive-loop-optimizations
import "C"

func init() {
	C.himpp3_ko_init()
}

//SysInit dfsdf
func SysInit() {
	//var tmp C.
	var tmp = C.himpp3_sys_init()
	fmt.Println("SysInit %d", tmp)
}

//ViInit sdfsdf
func ViInit() {
	var tmp = C.himpp3_vi_init()
	fmt.Println("ViInit %d", tmp)
}

//MipiIspInit sdf sdfsd
func MipiIspInit() {
	var tmp = C.himpp3_mipi_isp_init()
	fmt.Println("MipiIspInit %d", tmp)
}

//VpssInit sdfsdf sd
func VpssInit() {
	var tmp = C.himpp3_vpss_init()
	fmt.Println("VpssInit %d", tmp)
}

//VencInit sdfsd f
func VencInit() {
	var tmp = C.himpp3_venc_init()
	fmt.Println("VencInit %d", tmp)
	if tmp > 0 {
		fmt.Println(tmp, " this is venc channel fd")
		// let start goroutine to get frames
		go jpegGetDataLoop((int)(tmp))
	}
}
/*
type jpegFrame struct {
	//data???
	size	uint32
	size2   uint32
}

type jpegData struct {
	maxFrames		uint8
	currentFrame	uint8
	frames 			[2]jpegFrame
}
*/
func jpegGetDataLoop(fd int) {
	/*
	var bytePtr *C.char
	var bytePtr2 **C.char
	bytePtr2 = &bytePtr
	fmt.Println("** before adress ", bytePtr2)
	fmt.Println("* before address ", bytePtr)
	C.himpp3_test_func(bytePtr2)
	fmt.Println("** after address ", bytePtr2)
	fmt.Println("* after address ", bytePtr)
	fmt.Println("* after value ", *bytePtr)
	*/
	var counter uint64
	//var read_fdset syscall.FdSet
	rFDSet := &goselect.FDSet{}
	for {
	rFDSet.Zero()
	rFDSet.Set((uintptr)(fd))
	if err := goselect.Select(fd+1, rFDSet, nil, nil, -1); err != nil {
		fmt.Println("SELECT FATAL ERROR")
	}

	for i := uintptr(0); i < goselect.FD_SETSIZE; i++ {
		if rFDSet.IsSet(i) {
			//fmt.Println(i, "is ready")
		}
	}
	C.himpp3_venc_jpeg_export_frame()
	counter++
	//fmt.Println("counter ", counter)
	}
}

func jpegGetFrame() {

}

