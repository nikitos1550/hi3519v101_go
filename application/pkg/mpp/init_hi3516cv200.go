//+build arm
//+build hi3516cv200

package mpp

/*
#include "./include/mpp_v2.h"

#define ERR_NONE    0
#define ERR_MPP     1

int mpp2_sys_exit(int *error_code) {
    *error_code = 0;
    *error_code = HI_MPI_SYS_Exit();
    if (*error_code != HI_SUCCESS) return ERR_MPP;
    return ERR_NONE;
}

int mpp2_vb_exit(int *error_code) {
    *error_code = 0;
    *error_code = HI_MPI_VB_Exit();
    if (*error_code != HI_SUCCESS) return ERR_MPP;
    return ERR_NONE;
}

int mpp2_isp_exit(int *error_code) {
    *error_code = 0;
    *error_code = HI_MPI_ISP_Exit(0);
    if (*error_code != HI_SUCCESS) return ERR_MPP;
    return ERR_NONE;
}
*/
import "C"

import (
	"application/pkg/ko"
)

func systemInit(devInfo DeviceInfo) {
	//This family originally pack all reg init to sy_conf ko module
	ko.UnloadAll()

    ko.Params.Add("mem_start_addr").Str("0x").Uint64Hex(DDRMemStartAddr + devInfo.MemLinuxSize)
    ko.Params.Add("mem_mpp_size").Uint64(devInfo.MemMppSize/(1024*1024)).Str("M")
    ko.Params.Add("mem_total_size").Uint64(devInfo.MemTotalSize/(1024*1024))
	ko.LoadAll()
}
