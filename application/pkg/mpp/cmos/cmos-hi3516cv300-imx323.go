//+build ignore

package cmos

//#cgo LDFLAGS: ${SRCDIR}/hi3516cv300/sony_imx323/libsns_imx323.a
//int sensor_unregister_callback(void);
//int sensor_register_callback(void);
import "C"

import (
    "fmt"
    //"reflect"
)

func test() {
    //var ptr uintptr = reflect.ValueOf(C.sensor_register_callback).Pointer()
    ptr := C.sensor_register_callback()
    fmt.Printf("0x%x", ptr)
}
