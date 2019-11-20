// +build hi3516cv300

package chipid

// #include "hi_common.h"
// HI_S32 HI_MPI_SYS_GetChipId(HI_U32 *pu32ChipId);
import "C"

func Reg() uint32 {
    sysIdReg := readDevMem32(0x12020EE0) & 0xFF
    sysIdReg = sysIdReg + ((readDevMem32(0x12020EE4) & 0xFF) << 8)
    sysIdReg = sysIdReg + ((readDevMem32(0x12020EE8) & 0xFF) << 16)
    sysIdReg = sysIdReg + ((readDevMem32(0x12020EEC) & 0xFF) << 24)

    return sysIdReg
}

func Mpp() uint32 {
    var id C.HI_U32
    C.HI_MPI_SYS_GetChipId(&id)

    return uint32(id)
}

