//+build arm
//+build hi3516cv200

package mpp

/*
#include "./include/mpp.h"

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

    "application/pkg/mpp/cmos"

    _"application/pkg/utils"
)

const (
    DDRMemStartAddr = 0x80000000
)

func systemInit(devInfo DeviceInfo) {
	//This family originally pack all reg init to sy_conf ko module
	ko.UnloadAll()

    //detect_err_frame=10
    //rfr_frame_comp=1;

    //       insmod hi3518e_chnl.ko ChnlLowPower=1
    //insmod hi3518e_h264e.ko H264eMiniBufMode=1
    //insmod hi3518e_jpege.ko
    //insmod hi3518e_ive.ko save_power=0;

    ko.Params.Add("pin_mux_select").Uint64(0)
    ko.Params.Add("detect_err_frame").Uint64(10)
    ko.Params.Add("rfr_frame_comp").Uint64(1)
    ko.Params.Add("ChnlLowPower").Uint64(1)
    ko.Params.Add("H264eMiniBufMode").Uint64(1)
    ko.Params.Add("save_power").Uint64(0)

    ko.Params.Add("vi_vpss_online").Bool(devInfo.ViVpssOnline)

    //        ov9750)
    //        insmod hi3518e_isp.ko update_pos=1;
    //        ;;
    //    *)
    //        insmod hi3518e_isp.ko update_pos=0 proc_param=1;
    //        ;;


    ko.Params.Add("update_pos").Uint64(0)
    ko.Params.Add("proc_param").Uint64(1)

    //ko.Params.Add("update_pos").Uint64(1)
    //ko.Params.Add("proc_param").Uint64(1)


    ko.Params.Add("mem_start_addr").Str("0x").Uint64Hex(DDRMemStartAddr + devInfo.MemLinuxSize)
    ko.Params.Add("mem_mpp_size").Uint64(devInfo.MemMppSize/(1024*1024)).Str("M")
    ko.Params.Add("mem_total_size").Uint64(devInfo.MemTotalSize/(1024*1024))

    if cmos.S.Model() == "f22" {
        ko.Params.Add("cmos").Str("ar0130") //same DC i2c, waiting untill init rework
    } else if cmos.S.Model() == "h65" {
        ko.Params.Add("cmos").Str("ar0130") //same DC i2c, waiting untill init rework
    } else {
        ko.Params.Add("cmos").Str(cmos.S.Model())
    }
    //tmp for f22
    //ko.Params.Add("cmos").Str("ar0130")
    //utils.WriteDevMem32(0x200f0100,0x00000001)
    //utils.WriteDevMem32(0x200f0104,0x00000001)
    //utils.WriteDevMem32(0x200f0078,0x00000000)

	ko.LoadAll()
}
