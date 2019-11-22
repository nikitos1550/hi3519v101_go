// +build hi3516cv300

package chip

var (
    chips = [...]string {
        "hi3516cv300",
        "hi3516ev100",
    }
)

func RegId() uint32 {
    sysIdReg := readDevMem32(0x12020EE0) & 0xFF
    sysIdReg = sysIdReg + ((readDevMem32(0x12020EE4) & 0xFF) << 8)
    sysIdReg = sysIdReg + ((readDevMem32(0x12020EE8) & 0xFF) << 16)
    sysIdReg = sysIdReg + ((readDevMem32(0x12020EEC) & 0xFF) << 24)

    return sysIdReg
}
