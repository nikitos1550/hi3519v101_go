package main

import (
    _ "log"
    _ "sync"
    _ "bytes"
)



//#include "../hisi_external.h"
//#cgo LDFLAGS: ${SRCDIR}/../libhisi.a
import "C"

func main() {
    var test C.struct_cmos_params
    C.hisi_init(&test)
}
