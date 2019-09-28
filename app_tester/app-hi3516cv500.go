// +build hi3516cv500

package main

// #cgo LDFLAGS: ${SRCDIR}/hi3516cv500/libhi3516cv500.a
// #include "hi3516cv500/include/hi_common.h"
// #include "hi3516cv500/include/mpi_sys.h"
import "C"

const (
    chipFamily = "hi3516cv500"
)

func version() string {
    var ver C.MPP_VERSION_S
    C.HI_MPI_SYS_GetVersion(&ver)
    return C.GoString(&ver.aVersion[0])
}

func chipId() uint64 {
    return 0
}

