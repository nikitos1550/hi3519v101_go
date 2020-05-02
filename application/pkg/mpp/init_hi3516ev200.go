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

    "application/pkg/mpp/cmos"
)

const (
    DDRMemStartAddr = 0x40000000
)

//TODO rework this mess
func systemInit(devInfo DeviceInfo) {
	//This family originally pack all reg init to sy_conf ko module
	ko.UnloadAll()

    ko.Params.Add("mem_start_addr").Str("0x").Uint64Hex(DDRMemStartAddr + devInfo.MemLinuxSize)
    ko.Params.Add("mem_mpp_size").Uint64(devInfo.MemMppSize/(1024*1024)).Str("M")
    //ko.Params.Add("mem_total_size").Uint64(devInfo.MemTotalSize/(1024*1024))
    ko.Params.Add("cmos").Str(cmos.S.Model())
    ko.Params.Add("chip").Str(devInfo.Chip)
    ko.Params.Add("g_cmos_yuv_flag").Uint64(0) // 0 -- raw, 1 --DC, 3 --bt656
    ko.Params.Add("board").Str("demo")
    ko.Params.Add("save_power").Uint64(1)

	ko.LoadAll()
}
