// +build hi3516av100

package chip

// #include "hi_common.h"
// HI_S32 HI_MPI_SYS_GetChipId(HI_U32 *pu32ChipId);
import "C"

var (
    chips = [...]string {
        "hi3516av100",
        "hi3516dv100",
    }
)

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

