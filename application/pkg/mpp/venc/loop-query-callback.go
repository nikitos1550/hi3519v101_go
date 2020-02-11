package venc

/*
#include "loop.h"
//Should be here to export go callback
*/
import "C"

import (
    //"log"
    "unsafe"
    //"sync"
)

//export go_callback_receive_data
func go_callback_receive_data(venc_channel C.int, data_pointer * C.data_from_c, data_num C.int) {
    var vencChannel int
    var num         int
    var length      int

    vencChannel = int(venc_channel)
    num         = int(data_num)

    dataFromC   := (*[1 << 10]C.data_from_c)(unsafe.Pointer(data_pointer))[:num:num]

    if (vencChannel == 1) {
        //mux.Lock() //move near encoder object access
        //log.Println("Total ", length, " bytes")
        length = 0
        for i := 0; i < num; i++ {
            length      = length + int(dataFromC[i].length)
        }
        //array = make([]byte, length)
        //var offset int
        //F.Delete()
        for i := 0; i < num; i++ {
            //log.Println("Item length ", dataFromC[i].length)
            //this data slice is safe!!!
            data := (*[1 << 28]byte)(unsafe.Pointer(dataFromC[i].data))[:dataFromC[i].length:dataFromC[i].length]
            ///*copied := */ copy(array[offset:(offset+int(dataFromC[i].length))], data)
            //F.Append(data)
            if i == 0 {
                SampleMjpegFrames.Write(data)
            } else {
                SampleMjpegFrames.Append(data)
            }
            /*TODO
                find corresponding go space encoder object and copy data there
                encoder := findEncoder(vencChannel)
                if encoder != nil {
                    //encoder found
                    frame := encoder.frames.getNextFrame //find frame to write
                    frame.Write(data)
                } else {
                    //encoder not found, seems BUG!
                }
            */
            //log.Println("copied ", copied, " bytes")
            //offset      = offset + int(dataFromC[i].length)
        }
        //mux.Unlock() //move near encoder object access
    }

    //log.Println("go_callback_receive_data(",vencChannel,") done")
    //log.Println(reflect.TypeOf(data))
}

//DEPRECATED

var (
    //mux     sync.Mutex
    //array   []uint8
    F       frame
)


//func TmpLock() {
//    mux.Lock()
//}
//func TmpUnlock() {
//    mux.Unlock()
//}

//func TmpGet() []uint8 { //target io.Writer) {
//    return array
    /*
    mux.Lock()
        copied := copy(target, array)
        //io.Copy(target, array)
        log.Println("TmpMjpegCopyTo copied ")//, copied)
    mux.Unlock()
    */
//}


/*
        var theCArray *C.YourType = C.getTheArray()
        length := C.getTheArrayLength()
        slice := (*[1 << 28]C.YourType)(unsafe.Pointer(theCArray))[:length:length]
*/

