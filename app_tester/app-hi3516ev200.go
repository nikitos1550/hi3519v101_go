// +build hi3516ev200

package main

// #cgo LDFLAGS: ${SRCDIR}/hi3516ev200/libhi3516ev200.a
// #include "hi3516ev200/include/hi_common.h"
// #include "hi3516ev200/include/mpi_sys.h"
import "C"

const (
    chipFamily = "hi3516ev200"
)

func version() string {
    var ver C.MPP_VERSION_S
    C.HI_MPI_SYS_GetVersion(&ver)
    return C.GoString(&ver.aVersion[0])
}

