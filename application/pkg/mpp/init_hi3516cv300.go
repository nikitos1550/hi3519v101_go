//+build arm
//+build hi3516cv300

package mpp

/*
#include "./include/mpp_v3.h"

#define ERR_NONE    0
#define ERR_MPP     1

int mpp3_sys_exit(int *error_code) {
    *error_code = 0;
    *error_code = HI_MPI_SYS_Exit();
    if (*error_code != HI_SUCCESS) return ERR_MPP;
    return ERR_NONE;
}

int mpp3_vb_exit(int *error_code) {
    *error_code = 0;
    *error_code = HI_MPI_VB_Exit();
    if (*error_code != HI_SUCCESS) return ERR_MPP;
    return ERR_NONE;
}

int mpp3_isp_exit(int *error_code) {
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
    "application/pkg/logger"
)

const (
    DDRMemStartAddr = 0x80000000
)

func systemInit(devInfo DeviceInfo) {
	//This family originally pack all reg init to sy_conf ko module
	ko.UnloadAll()

    //ko.Params.Add("mem_start_addr").Str("0x").Uint64Hex(DDRMemStartAddr + devInfo.MemLinuxSize)
    //ko.Params.Add("mem_mpp_size").Uint64(devInfo.MemMppSize/(1024*1024)).Str("M")
    //ko.Params.Add("mem_total_size").Uint64(devInfo.MemTotalSize/(1024*1024))

    ///////devInfo.ViVpssOnline = true

    ko.Params.Add("vi_vpss_online").Bool(devInfo.ViVpssOnline)
    ko.Params.Add("mem_start_addr").Str("0x").Uint64Hex(DDRMemStartAddr + devInfo.MemLinuxSize)
    ko.Params.Add("mem_mpp_size").Uint64(devInfo.MemMppSize/(1024*1024)).Str("M")
    
    //ko.Params.Add("cmos").Str(cmos.Model())
    ko.Params.Add("cmos").Str("NULL")

    ko.Params.Add("mem_total_size").Uint64(devInfo.MemTotalSize/(1024*1024))
    ko.Params.Add("vgs_clk_frequency").Uint64(297000000)
    ko.Params.Add("detect_err_frame").Uint64(10)
    ko.Params.Add("viu_clk_frequency").Uint64(19800000)
    ko.Params.Add("isp_div").Uint64(2)
    ko.Params.Add("input_mode").Str("default")
    ko.Params.Add("update_pos").Uint64(0)
    ko.Params.Add("proc_param").Uint64(30)
    ko.Params.Add("port_init_delay").Uint64(0)
    ko.Params.Add("vpss_clk_frequency").Uint64(250000000)
    ko.Params.Add("vedu_clk_frequency").Uint64(198000000)
    ko.Params.Add("save_power").Uint64(1)
    ko.Params.Add("ive_clk_frequency").Uint64(297000000)

	var tmpBus string
	var tmpData string

    switch cmos.BusType() {
        case cmos.I2C:
			tmpBus = "i2c"
        case cmos.SPI:
			tmpBus = "ssp"
        default:
        	logger.Log.Fatal().
            	Int("type", int(cmos.BusType())).
                Msg("unrecognized cmos bus type")
    }

    switch cmos.Data() {
        case cmos.DC:
            tmpData = "dc"
        case cmos.MIPI:
            tmpData = "mipi"
        case cmos.LVDS:
            tmpData = "mipi"
        default:
        	logger.Log.Fatal().
                Int("type", int(cmos.Data())).
                Msg("unrecognized cmos data type")
    }

	ko.Params.Add("sensor_bus_type").Str(tmpBus)
    ko.Params.Add("sensor_clk_frequency").Uint64(37125000)
    ko.Params.Add("sensor_pinmux_mode").Str(tmpBus+"_"+tmpData)

	ko.LoadAll()
}




