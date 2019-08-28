package hisi

import (
    _ "log"
    "sync"
    _ "bytes"
)



//#include "hisi_external.h"
//#cgo LDFLAGS: libhisi.a
import "C"

//type frames struct {
//    bytesPool sync.Pool //{ New: func() interface{} { return []byte{} },}
//}

//type frame struct {
//    lock sync.Mutex
//}

//type channel struct {
//    id uint
//}

func main() {
    var test C.struct_cmos_params
    C.hisi_init(&test)
}

type frame struct {
    lock    sync.Mutex          //lock frame while read or write
    data    []byte              //slice, underlay mem will be allocated from pool
}

type encoder struct {
    id              uint
    enabled         uint
    bytesPool       *sync.Pool  //pool to store frames
    frames          []frame     //frames itself
    numFrames       uint
    lastFrame       uint
    notification    []chan int
}

var Encoders []encoder
//var channels []channel

func myInit() {
    //channels = make([]channel, 4, 4) //TODO 4 should be taken from hisi_channels_max_num()
    Encoders = make([]encoder, 16, 16) //TODO 16 should be taken from hisi_encoders_max_num()
    /***********************************************/
    Encoders[0].notification = make([]chan int, 10)
    for i := range Encoders[0].notification {
        Encoders[0].notification[i] = make(chan int)
    }
}

//func init frame pool
//func put frame, inputs data pointer and size
//func get frame (means get pointer to data)
//func get last frame (means get pointer to data)

//TODO notification, how to deal with?

//func List
//func GetInfo
//func Create
//func Delete
//func Update

//export goDataCallback
func goDataCallback(encId C.uint, encData * C.struct_encoderData) C.int {

    return 0;
}
