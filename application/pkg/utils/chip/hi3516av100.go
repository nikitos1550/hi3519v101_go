//+build arm
//+build hi3516av100

package chip

import "application/pkg/utils"

var (
    chips = [...]string {
        "hi3516av100",
        "hi3516dv100",
    }
)

func RegId() uint32 {
    sysIdReg := utils.ReadDevMem32(0x20050EE0) & 0xFF
    sysIdReg = sysIdReg + ((utils.ReadDevMem32(0x20050EE4) & 0xFF) << 8)
    sysIdReg = sysIdReg + ((utils.ReadDevMem32(0x20050EE8) & 0xFF) << 16)
    sysIdReg = sysIdReg + ((utils.ReadDevMem32(0x20050EEC) & 0xFF) << 24)

    return sysIdReg
}

