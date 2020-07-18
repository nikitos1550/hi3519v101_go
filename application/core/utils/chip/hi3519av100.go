//+build arm
//+build hi3519av100

package chip

import "application/core/utils"

var (
    chips = [...]string {
        "hi3519av100",
    }
)

func RegId() uint32 {
    sysIdReg := utils.ReadDevMem32(0x12020EE0)
    return sysIdReg
}

