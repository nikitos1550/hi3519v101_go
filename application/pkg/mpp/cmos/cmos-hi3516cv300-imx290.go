//+build ignore

package cmos

//#cgo LDFLAGS: ./hi3516cv300/sony_imx290/libsns_imx290.a
//#cgo LDFLAGS: /home/nikitos1550/work/hi3519v101_go/output/hisilicon/hi3516cv300/libhi3516cv300.a
//int sensor_unregister_callback(void);
//int sensor_register_callback(void);
import "C"

import (
    "fmt"
    //"reflect"
)

func test2() {
    //var ptr uintptr = reflect.ValueOf(C.sensor_register_callback).Pointer()
    ptr := C.sensor_register_callback()
    fmt.Printf("0x%x", ptr)
}


