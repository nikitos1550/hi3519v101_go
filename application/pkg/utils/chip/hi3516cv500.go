//+build arm
//+build hi3516cv500

package chip

import "application/pkg/utils"

var (
    chips = [...]string {
        "hi3516cv500",
        "hi3516dv300",
        "hi3516av300",
    }
)

func RegId() uint32 {
    sysIdReg := utils.ReadDevMem32(0x12020EE0)

    return sysIdReg
}
