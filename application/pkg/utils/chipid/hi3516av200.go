// +build hi3516av200

package chipid

// #include "hi_common.h"
// HI_S32 HI_MPI_SYS_GetChipId(HI_U32 *pu32ChipId);
import "C"

import "application/pkg/utils"

var (
    chips = [...]string {
        "hi3519v101",
        "hi3516av200",
    }
)

func Reg() uint32 {
    sysIdReg := utils.ReadDevMem32(0x12020EE0) & 0xFF
    sysIdReg = sysIdReg + ((utils.ReadDevMem32(0x12020EE4) & 0xFF) << 8)
    sysIdReg = sysIdReg + ((utils.ReadDevMem32(0x12020EE8) & 0xFF) << 16)
    sysIdReg = sysIdReg + ((utils.ReadDevMem32(0x12020EEC) & 0xFF) << 24)

    return sysIdReg
}

func Mpp() uint32 {
    var id C.HI_U32
    C.HI_MPI_SYS_GetChipId(&id)

    return uint32(id)
}


