package himpp3

import "fmt"

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
		go jpegGetDataLoop()
	}
}

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

func jpegGetDataLoop() {

}

func jpegGetFrame() {

}

