// +build hi3516cv100

package main

// #cgo LDFLAGS: ${SRCDIR}/hi3516cv100/libhi3516cv100.a
// #include "hi3516cv100/include/hi_common.h"
// #include "hi3516cv100/include/mpi_sys.h"
// HI_S32 HI_MPI_SYS_GetChipId(HI_U32 *pu32ChipId);
import "C"

func Reg() uint32 {
    sysIdReg := readDevMem32(0x20050EE0) & 0xFF
    sysIdReg = sysIdReg + ((readDevMem32(0x20050EE4) & 0xFF) << 8)
    sysIdReg = sysIdReg + ((readDevMem32(0x20050EE8) & 0xFF) << 16)
    sysIdReg = sysIdReg + ((readDevMem32(0x20050EEC) & 0xFF) << 24)

    return sysIdReg
}

func Mpp() uint32 {
    var id C.HI_U32
    C.HI_MPI_SYS_GetChipId(&id)

    return uint32(id)
}

