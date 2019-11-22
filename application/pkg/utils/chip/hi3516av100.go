// +build hi3516av100

package chip

var (
    chips = [...]string {
        "hi3516av100",
        "hi3516dv100",
    }
)

func RegId() uint32 {
    sysIdReg := readDevMem32(0x20050EE0) & 0xFF
    sysIdReg = sysIdReg + ((readDevMem32(0x20050EE4) & 0xFF) << 8)
    sysIdReg = sysIdReg + ((readDevMem32(0x20050EE8) & 0xFF) << 16)
    sysIdReg = sysIdReg + ((readDevMem32(0x20050EEC) & 0xFF) << 24)

    return sysIdReg
}

