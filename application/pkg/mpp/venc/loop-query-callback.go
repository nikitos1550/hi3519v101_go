package venc

/*
#include "loop.h"
//Should be here to export go callback
*/
import "C"

import (
    "log"
    "unsafe"
    //"sync"
)

var tmp int

//export go_callback_receive_data
func go_callback_receive_data(venc_channel C.int, seq C.uint, data_pointer * C.data_from_c, data_num C.int) {
    var vencChannel int
    var num         int
    var length      int
    var sequence    uint32

    vencChannel = int(venc_channel)
    num         = int(data_num)
    sequence    = uint32(seq)

    dataFromC   := (*[1 << 10]C.data_from_c)(unsafe.Pointer(data_pointer))[:num:num]

    if (vencChannel == 0) { //SampleH264
        length = 0
        pp := make([][]byte, num)
        for i := 0; i < num; i++ {
            length      = length + int(dataFromC[i].length)
        }
        if tmp == 0 {
            select {
                case start := <-SampleH264Start:
                    if start == 100 {
                        tmp = 1
                        log.Println("VENC H264 start")
                    }
                default:
                    return
            }
        }
        if tmp == 1 {
            for i := 0; i < num; i++ {
                //data := (*[1 << 28]byte)(unsafe.Pointer(dataFromC[i].data))[:dataFromC[i].length:dataFromC[i].length]
                pp[i] = (*[1 << 28]byte)(unsafe.Pointer(dataFromC[i].data))[:dataFromC[i].length:dataFromC[i].length]
                //nalType := pp[i][4] & 0x1F
                //log.Println("Found NAL ", nalType)

                /*
                SampleH264Frames.WriteNext(data, sequence)
                //log.Println("len(SampleH264Notify)", len(SampleH264Notify))
                if len(SampleH264Notify) > (10-1) {
                    //log.Println("SampleH264Notify channel is full, no send")
                } else {
                    SampleH264Notify <- int(sequence)
                    //log.Println("SampleH264Notify channel sent", sequence)
                }
                */
            }
            SampleH264Frames.WritevNext(pp, sequence)
            SampleH264Notify <- int(sequence)
        }
    } else {
    //if (vencChannel == 1) { //SampleMjpeg
        length = 0

        pp := make([][]byte, num)

        for i := 0; i < num; i++ {
            length      = length + int(dataFromC[i].length)
        }
        for i := 0; i < num; i++ {
            pp[i] = (*[1 << 28]byte)(unsafe.Pointer(dataFromC[i].data))[:dataFromC[i].length:dataFromC[i].length]
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
        SampleMjpegFrames.WritevNext(pp, sequence)
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

