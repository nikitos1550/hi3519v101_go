//+build arm
//+build hi3516av200

package chip

import "application/core/utils"
//import "log"

var (
    chips = [...]string {
        "hi3519v101",
        "hi3516av200",
    }
)

func RegId() uint32 {
    sysIdReg := utils.ReadDevMem32(0x12020EE0) 
    //log.Println(sysIdReg)
    if sysIdReg != 0x06000001 {
	sysIdReg = sysIdReg & 0xFF
    	sysIdReg = sysIdReg + ((utils.ReadDevMem32(0x12020EE4) & 0xFF) << 8)
    	sysIdReg = sysIdReg + ((utils.ReadDevMem32(0x12020EE8) & 0xFF) << 16)
    	sysIdReg = sysIdReg + ((utils.ReadDevMem32(0x12020EEC) & 0xFF) << 24)
    }
    return sysIdReg
    
}
