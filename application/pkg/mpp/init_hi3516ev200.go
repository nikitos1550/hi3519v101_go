//+build arm
//+build hi3516ev200

package mpp

import (
    //"log"
    //"application/pkg/logger"
    //"os"

	"application/pkg/ko"
    //"application/pkg/utils"
    //"application/pkg/mpp/error"
)

//TODO rework this mess
func systemInit(devInfo DeviceInfo) {
	//This family originally pack all reg init to sy_conf ko module
	ko.UnloadAll()

    ko.Params.Add("mem_start_addr").Str("0x").Uint64Hex(DDRMemStartAddr + devInfo.MemLinuxSize)
    ko.Params.Add("mem_mpp_size").Uint64(devInfo.MemMppSize/(1024*1024)).Str("M")
    ko.Params.Add("mem_total_size").Uint64(devInfo.MemTotalSize/(1024*1024))
	ko.LoadAll()
}
