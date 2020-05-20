//+build arm
//+build hi3516av100 hi3516av200 hi3516cv100 hi3516cv200 hi3516cv300 hi3516cv500 hi3516ev200 hi3519av100 hi3559av100

package venc

//#include "venc.h"
import "C"

import (
	"application/pkg/logger"

    "unsafe"

    "fmt"
    "net/http"
    "application/pkg/openapi"

    "sync"
)

var mutex = &sync.Mutex{}

func init() {
    openapi.AddApiRoute("serveDebugLoop", "/mpp/venc/loop", "GET", serveDebugLoop)
}

func serveDebugLoop(w http.ResponseWriter, r *http.Request) {
    logger.Log.Trace().
	    Msg("mpp.venc.serveDebugLoop")

    w.Header().Set("Content-Type", "test/plain; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    fmt.Fprintf(w, "%s", "TODO")
}


//export go_callback_receive_data
func go_callback_receive_data(venc_channel C.int, info_pointer *C.info_from_c, data_pointer *C.data_from_c, data_num C.int) { 
    //vencChannel := int(venc_channel)

    encoder, exists := ActiveEncoders[int(venc_channel)]//[vencChannel]
    if (!exists) {
      return
    }

    num := int(data_num) 

    //var infoFromC *C.info_from_c = info_pointer
    //infoFromC.ref_type:
    //BASE_IDRSLICE = 0           //IDR frame at the base layer
    //BASE_PSLICE_REFTOIDR        //P-frame at the base layer, referenced by other frames at the base layer and references only IDR frames
    //BASE_PSLICE_REFBYBASE       //P-frame at the base layer, referenced by other frames at the base layer
    //BASE_PSLICE_REFBYENHANCE    //P-frame at the base layer, referenced by frames at the enhance layer
    //ENHANCE_PSLICE_REFBYENHANCE
    //ENHANCE_PSLICE_NOTFORREF

    dataFromC := (*[1 << 10]C.data_from_c)(unsafe.Pointer(data_pointer))[:num:num]
    length := 0
    for i := 0; i < num; i++ {
        length = length + int(dataFromC[i].length)
    }

    data := make([]byte, length)

    offset := 0
    for i := 0; i < num; i++ {
        p := (*[1 << 28]byte)(unsafe.Pointer(dataFromC[i].data))[:dataFromC[i].length:dataFromC[i].length]
        n := copy(data[offset:], p)
        offset = offset + n
    }

    for ch,enabled := range encoder.Channels {
        if (enabled){
            if (cap(ch) <= len(ch)) {
                <-ch
            }

            ch <- data
        }
    }
}


//Rules:
//addVenc/delVenc functions should operate one in a time.
//There should be some sync for them run:
//1) exlusive mutex
//2) query (based obviosly on go channels)

func addVenc(venc int) {
    mutex.Lock()
	defer mutex.Unlock()

    var errorCode C.uint
    var vencChannelId C.uint
    vencChannelId = C.uint(venc)

    switch err := C.mpp_data_loop_add(&errorCode, vencChannelId, 1); err { //TODO
    case C.ERR_NONE:
	    logger.Log.Debug().
		    Msg("C.mpp_data_loop_add() ok")
    case C.ERR_GENERAL: //C.ERR_SYS:
	    logger.Log.Fatal().
		    Msg("C.mpp_data_loop_add() SYS error")
    default:
	    logger.Log.Fatal().
		    Int("error", int(err)).
		    Msg("C.mpp_data_loop_add() Unexpected return")
    }

    logger.Log.Debug().
	    Int("channel", venc).
	    Msg("VENC channel added to loop")
}

func delVenc(venc int) {
    mutex.Lock()
	defer mutex.Unlock()

    var errorCode C.uint
    var vencChannelId C.uint
    vencChannelId = C.uint(venc)

    switch err := C.mpp_data_loop_del(&errorCode, vencChannelId); err {
    case C.ERR_NONE:
	    logger.Log.Debug().
		    Msg("C.mpp_data_loop_del() ok")
    case C.ERR_GENERAL: //C.ERR_SYS:
	    logger.Log.Fatal().
		    Msg("C.mpp_data_loop_del() SYS error")
    default:
	    logger.Log.Fatal().
		    Int("error", int(err)).
		    Msg("C.mpp_data_loop_del() Unexpected return")
    }

    logger.Log.Debug().
        Int("channel", venc).
        Msg("VENC channel deleted from loop")
}


func loopInit() {
    var errorCode C.uint

    switch err := C.mpp_data_loop_init(&errorCode); err {
    case C.ERR_NONE:
	    logger.Log.Debug().
		    Msg("C.mpp_data_loop_init() ok")
    case C.ERR_GENERAL: //C.ERR_SYS:
	    logger.Log.Fatal().
		    Msg("C.mpp_data_loop_init() SYS error")
    default:
	    logger.Log.Fatal().
		    Int("error", int(err)).
		    Msg("C.mpp_data_loop_init() Unexpected return")
    }

}
