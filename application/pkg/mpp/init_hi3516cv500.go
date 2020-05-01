//+build arm
//+build hi3516cv500

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
    DDRMemStartAddr = 0x80000000
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

	ko.LoadAll()

    /*
         VI_VPSS_MODE_S      stVIVPSSMode;

    *error_code = HI_MPI_SYS_GetVIVPSSMode(&stVIVPSSMode);
    if (*error_code != HI_SUCCESS) {
        
        return ERR_GENERAL;
    }   

    stVIVPSSMode.aenMode[0] = VI_OFFLINE_VPSS_OFFLINE;

    *error_code = HI_MPI_SYS_SetVIVPSSMode(&stVIVPSSMode);
    if (*error_code != HI_SUCCESS) {
        
        return ERR_GENERAL; 
    }   

    */
}
