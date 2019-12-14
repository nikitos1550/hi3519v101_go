//+build hi3516av200 hi3516cv300

package getloop

/*
#include "getloop.h"
//Should be here to export go callback
*/
import "C"

import (
    //"log"
    "unsafe"
    //"reflect"
    "sync"
    //"io"
)

var (
    mux sync.Mutex
    array []uint8
)

//export go_callback_receive_data
func go_callback_receive_data(venc_channel C.int, data_pointer * C.data_from_c, data_num C.int, data_length C.int) {
    var vencChannel int
    var length      int
    var num         int
    //var dataFromC   *C.data_from_c

    vencChannel = int(venc_channel)
    length      = int(data_length)
    num         = int(data_num)
    //dataFromC   = data_pointer

    dataFromC   := (*[1 << 10]C.data_from_c)(unsafe.Pointer(data_pointer))[:num:num]

    if (vencChannel == 1) {
        mux.Lock()
        //log.Println("Total ", length, " bytes")
        array = make([]uint8, length)
        var offset int
        for i := 0; i < num; i++ {
            //log.Println("Item length ", dataFromC[i].length)
            data        := (*[1 << 28]uint8)(unsafe.Pointer(dataFromC[i].data))[:dataFromC[i].length:dataFromC[i].length]
            /*copied      :=*/ copy(array[offset:(offset+int(dataFromC[i].length))], data)
            //log.Println("copied ", copied, " bytes")
            offset      = offset + int(dataFromC[i].length)
        }
        mux.Unlock()
    }

    //log.Println("go_callback_receive_data(",vencChannel,") done")
    //log.Println(reflect.TypeOf(data))
}

func TmpLock() {
    mux.Lock()
}
func TmpUnlock() {
    mux.Unlock()
}
func TmpGet() []uint8 { //target io.Writer) {
    return array
    /*
    mux.Lock()
        copied := copy(target, array)
        //io.Copy(target, array)
        log.Println("TmpMjpegCopyTo copied ")//, copied)
    mux.Unlock()
    */
}


/*
        var theCArray *C.YourType = C.getTheArray()
        length := C.getTheArrayLength()
        slice := (*[1 << 28]C.YourType)(unsafe.Pointer(theCArray))[:length:length]
*/
