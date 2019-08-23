package himpp3

import (
	"fmt"
	"unsafe"
//	"syscall"
//	"../external/github.com/creack/goselect"
	"bytes"
	"sync"

    "log"

    "net/http"
    "regexp"
    "strconv"
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

//TempGet dsfdsfsdf
func TempGet() float32 {
	return (float32)(C.gettemperature())
}

//SysInit dfsdf
func SysInit() {
	//var tmp C.
	var tmp = C.himpp3_sys_init()
	log.Println("SysInit %d", tmp)
}

//ViInit sdfsdf
func ViInit() {
	var tmp = C.himpp3_vi_init()
	log.Println("ViInit %d", tmp)
}

//MipiIspInit sdf sdfsd
func MipiIspInit() {
	var tmp = C.himpp3_mipi_isp_init()
	log.Println("MipiIspInit %d", tmp)
}

//VpssInit sdfsdf sd
func VpssInit() {
	var tmp = C.himpp3_vpss_init()
	log.Println("VpssInit %d", tmp)
}

//VencInit sdfsd f
func VencInit() {
	var tmp = C.himpp3_venc_init()
	log.Println("VencInit %d", tmp)
}

//vencMJpegSetBitrate description
func vencMJpegSetBitrate(b uint) int {
    var bitrate C.uint = C.uint(b)
    var tmp = C.himpp3_venc_mjpeg_params(bitrate)
    log.Println("VencMJpegSetBitrate %d", tmp)
    return int(tmp)
}

func ApiHandler (w http.ResponseWriter, r *http.Request) {
        log.Println("himpp3.ApiHandler")

        rr, _ := regexp.Compile("^/experimental/himpp3/bitrate/([0-9]+)$")
        match := rr.FindStringSubmatch(r.URL.Path)
        log.Println(match)

        if match != nil {
            var bitrate int
            bitrate, _ = strconv.Atoi(match[1])
            tmp := vencMJpegSetBitrate(uint(bitrate))
            if tmp == 0 {
                fmt.Fprintf(w, "bitrate %d is OK", bitrate)
            } else {
                fmt.Fprintf(w, "bitrate %d is bad %X", tmp)
            }
        }

}

///////////////////////////////////////////////////////////////////////////

var B1 bytes.Buffer
var Mutex sync.Mutex

//export jpegVencGetDataCallback
func jpegVencGetDataCallback(stStream * C.struct_jpegFrame) {
	//fmt.Println("New jpeg frame!")
	//fmt.Println("stStream.seq = ", stStream.seq)
	//fmt.Println("stStream.count = ", stStream.count)
	//fmt.Println("stStream.packs[0].length = ", stStream.packs[0].length)
	//fmt.Println("stStream.packs[1].length = ", stStream.packs[1].length)

	//b1.Grow((int)(stStream.packs[0].length) + (int)(stStream.packs[1].length))

	data1 := (*[1 << 30]byte)(unsafe.Pointer(stStream.packs[0].data))[:int(stStream.packs[0].length):int(stStream.packs[0].length)]
	data2 := (*[1 << 30]byte)(unsafe.Pointer(stStream.packs[1].data))[:int(stStream.packs[1].length):int(stStream.packs[1].length)]
	// or for an array if BUFF_SIZE is a constant
	//myArray := *(*[C.BUFF_SIZE]byte)(unsafe.Pointer(&C.my_buf))

	Mutex.Lock()
	B1.Reset()
	B1.Write(data1)
	B1.Write(data2)
	Mutex.Unlock()

	//fmt.Println("b1 cap = ", B1.Cap(), " len = ", B1.Len())
	//b1p := b1.Bytes()

	//fmt.Printf("%d %d %d\n", b1p[0], b1p[1], b1p[2])
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
/*
func jpegGetDataLoop(fd int) {
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
	//C.himpp3_venc_jpeg_export_frame()
	counter++
	//fmt.Println("counter ", counter)
	}
}
*/

