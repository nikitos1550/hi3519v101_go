//+build arm
//+build hi3516cv300

package mpp

/*
#include "./include/mpp.h"

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
        "os"

    "application/pkg/ko"
        "application/pkg/mpp/cmos"
    "application/pkg/logger"
)

const (
    DDRMemStartAddr = 0x80000000
)

func systemInit(devInfo DeviceInfo) {
   if _, err := os.Stat("/dev/sys"); err == nil { 
        var errorCode C.int
        /*
        err := C.mpp3_sys_exit(&errorCode)
        if err != nil {
            logger.Log.Fatal().
                Msg("TODO")

        }
        */
        
        switch err := C.mpp3_sys_exit(&errorCode); err {
        case C.ERR_NONE:
            logger.Log.Debug().
                Msg("C.mpp3_sys_exit() ok")
        case C.ERR_MPP:
            logger.Log.Fatal().
                Str("func", "HI_MPI_SYS_Exit()").
                Int("error", int(errorCode)).
                //Str("error_desc", error.Resolve(int64(errorCode))).
                Msg("C.mpp3_sys_exit() error")
        default:
            logger.Log.Fatal().
                Int("error", int(err)).
                Msg("Unexpected return of C.mpp3_sys_exit()")
        } 
        
    }

    if _, err := os.Stat("/dev/vb"); err == nil {      
        var errorCode C.int
        switch err := C.mpp3_vb_exit(&errorCode); err {
        case C.ERR_NONE:
            //log.Println("C.mpp3_vb_exit() ok")
        logger.Log.Debug().
            Msg("C.mpp3_vb_exit() ok")
        case C.ERR_MPP:
            //log.Fatal("C.mpp3_vb_exit() HI_MPI_VB_Exit() error ", error.Resolve(int64(errorCode))) 
        logger.Log.Fatal().
            Str("func", "HI_MPI_VB_Exit()").
            Int("error", int(errorCode)).
        //Str("error_desc", error.Resolve(int64(errorCode))).
        Msg("C.mpp3_vb_exit() error")
        default:
            //log.Fatal("Unexpected return ", err , " of C.mpp3_vb_exit()")
        logger.Log.Fatal().
                Int("error", int(err)).
            Msg("Unexpected return of C.mpp3_vb_exit()")
        }
    }



	//This family originally pack all reg init to sy_conf ko module
	ko.UnloadAll()

    //ko.Params.Add("mem_start_addr").Str("0x").Uint64Hex(DDRMemStartAddr + devInfo.MemLinuxSize)
    //ko.Params.Add("mem_mpp_size").Uint64(devInfo.MemMppSize/(1024*1024)).Str("M")
    //ko.Params.Add("mem_total_size").Uint64(devInfo.MemTotalSize/(1024*1024))

    ko.Params.Add("vi_vpss_online").Bool(devInfo.ViVpssOnline)
    ko.Params.Add("mem_start_addr").Str("0x").Uint64Hex(DDRMemStartAddr + devInfo.MemLinuxSize)
    ko.Params.Add("mem_mpp_size").Uint64(devInfo.MemMppSize/(1024*1024)).Str("M")
                
    ko.Params.Add("cmos").Str(cmos.S.Model())

    ko.Params.Add("mem_total_size").Uint64(devInfo.MemTotalSize/(1024*1024))
    ko.Params.Add("vgs_clk_frequency").Uint64(297000000)
    ko.Params.Add("detect_err_frame").Uint64(10)

    switch devInfo.Chip {
    case "hi3516cv300":

    //ko.Params.Add("vi_vpss_online").Bool(devInfo.ViVpssOnline)
    //ko.Params.Add("mem_start_addr").Str("0x").Uint64Hex(DDRMemStartAddr + devInfo.MemLinuxSize)
    //ko.Params.Add("mem_mpp_size").Uint64(devInfo.MemMppSize/(1024*1024)).Str("M")
    
    //ko.Params.Add("cmos").Str(cmos.Model())
    //ko.Params.Add("cmos").Str("NULL")

    //ko.Params.Add("mem_total_size").Uint64(devInfo.MemTotalSize/(1024*1024))
    //ko.Params.Add("vgs_clk_frequency").Uint64(297000000)
    //ko.Params.Add("detect_err_frame").Uint64(10)
    ko.Params.Add("viu_clk_frequency").Uint64(198000000)
    ko.Params.Add("isp_div").Uint64(2)
    //ko.Params.Add("isp_div").Uint64(1)

    //ko.Params.Add("input_mode").Str("default")
    //ko.Params.Add("update_pos").Uint64(0)
    //ko.Params.Add("proc_param").Uint64(30)
    //ko.Params.Add("port_init_delay").Uint64(0)
    //ko.Params.Add("vpss_clk_frequency").Uint64(250000000)
    //ko.Params.Add("vedu_clk_frequency").Uint64(198000000)
    //ko.Params.Add("save_power").Uint64(1)
    //ko.Params.Add("ive_clk_frequency").Uint64(297000000)

    case "hi3516ev100":
        ko.Params.Add("viu_clk_frequency").Uint64(83300000)
        ko.Params.Add("isp_div").Uint64(1)
    default:
        logger.Log.Fatal().
            Str("chip", devInfo.Chip).
            Msg("unknown chip version")
    }

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

    switch cmos.S.BusType() {
        case cmos.I2C:
			tmpBus = "i2c"
        case cmos.SPI:
			tmpBus = "ssp"
        default:
        	logger.Log.Fatal().
            	Int("type", int(cmos.S.BusType())).
                Msg("unrecognized cmos bus type")
    }

    switch cmos.S.Data() {
        case cmos.DC:
            tmpData = "dc"
        case cmos.MIPI:
            tmpData = "mipi"
        case cmos.LVDS:
            tmpData = "mipi"
        default:
        	logger.Log.Fatal().
                Int("type", int(cmos.S.Data())).
                Msg("unrecognized cmos data type")
    }

	ko.Params.Add("sensor_bus_type").Str(tmpBus)
    ko.Params.Add("sensor_clk_frequency").Uint64(37125000)
    ko.Params.Add("sensor_pinmux_mode").Str(tmpBus+"_"+tmpData)

	ko.LoadAll()
}




