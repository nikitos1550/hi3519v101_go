//+build arm
//+build hi3516ev200

package chip

import "application/core/utils"

var (
    chips = [...]string {
        "hi3516ev300",
        "hi3516ev200",
        "hi3516dv200",
        "hi3518ev300",
    }
)

func RegId() uint32 {
    sysIdReg := utils.ReadDevMem32(0x12020EE0)
    return sysIdReg
}

