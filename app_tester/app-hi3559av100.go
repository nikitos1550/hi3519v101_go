// +build hi3519av100

//TODO

package main

// #cgo LDFLAGS: ${SRCDIR}/hi3559av100/libhi3559av100.a
// #include "hi3559av100/include/hi_common.h"
// #include "hi3559av100/include/mpi_sys.h"
import "C"

const (
    chipFamily = "hi3519av100"
)

var (
    chips = [...]string {"hi3559av100"}
)

/*
func version() string {
    var ver C.MPP_VERSION_S
    C.HI_MPI_SYS_GetVersion(&ver)
    return C.GoString(&ver.aVersion[0])
}

func chipId() uint64 {
    return 0
}
*/
