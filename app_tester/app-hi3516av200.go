// +build hi3516av200

package main

// #cgo LDFLAGS: ${SRCDIR}/hi3516av200/libhi3516av200.a
// #include "hi3516av200/include/hi_common.h"
// #include "hi3516av200/include/mpi_sys.h"
import "C"

const (
    chipFamily = "hi3516av200"
)

func version() string {
    var ver C.MPP_VERSION_S
    C.HI_MPI_SYS_GetVersion(&ver)
    return C.GoString(&ver.aVersion[0])
}
