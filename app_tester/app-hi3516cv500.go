// +build hi3516cv500

package main

// #cgo LDFLAGS: ${SRCDIR}/hi3516cv500/libhi3516cv500.a -lm -ldl
// #include "hi3516cv500/include/hi_common.h"
// #include "hi3516cv500/include/mpi_sys.h"
// HI_S32 HI_MPI_SYS_GetChipId(HI_U32 *pu32ChipId);
import "C"

const (
    chipFamily = "hi3516cv500"
)

var (
    sysIdReg        uint32
    mppVersion      string

    modules = [...][2]string {
        [2]string{"sys_config.ko", "chip={chipName} sensors=sns0=NULL,sns1=NULL, g_cmos_yuv_flag=0"},
        [2]string{"hi_osal.ko", "anony=1 mmz_allocator=hisi mmz=anonymous,0,0x{memStartAddr},{memMppSize}M"},
        [2]string{"hi3516cv500_base.ko", ""},
        [2]string{"hi3516cv500_sys.ko", ""},
    }
    chips = [...]string {
        "hi3516cv500",
        "hi3516dv300",
        "hi3516av300",
    }
)

func init() {
    sysIdReg = readDevMem32(0x12020EE0)
    //sysIdReg = sysIdReg + ((readDevMem32(0x12020EE4) & 0xFF) << 8)
    //sysIdReg = sysIdReg + ((readDevMem32(0x12020EE8) & 0xFF) << 16)
    //sysIdReg = sysIdReg + ((readDevMem32(0x12020EEC) & 0xFF) << 24)

    var ver C.MPP_VERSION_S
    C.HI_MPI_SYS_GetVersion(&ver)
    mppVersion = C.GoString(&ver.aVersion[0])
}

func chipId() uint32 {
    var id C.HI_U32
    C.HI_MPI_SYS_GetChipId(&id)
    return uint32(id)
}

func initTemperature() {
    writeDevMem32(0x120300B4, 0xCBA00000)
}

func getTemperature() float32 {
    var tempCode uint32 = readDevMem32(0x120300BC)
    var temp float32 = ((( float32(tempCode & 0x3FF)-136)/793)*165)-40
    return temp
}

