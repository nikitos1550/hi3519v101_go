// +build hi3516av100

package main

// #cgo LDFLAGS: ${SRCDIR}/hi3516av100/libhi3516av100.a
// #include "hi3516av100/include/hi_common.h"
// #include "hi3516av100/include/mpi_sys.h"
// HI_S32 HI_MPI_SYS_GetChipId(HI_U32 *pu32ChipId);
import "C"

const (
    chipFamily = "hi3516av100"
)

func version() string {
    var ver C.MPP_VERSION_S
    C.HI_MPI_SYS_GetVersion(&ver)
    return C.GoString(&ver.aVersion[0])
}

func chipId() uint64 {
    var id C.HI_U32
    C.HI_MPI_SYS_GetChipId(&id)
    return uint64(id)
}


