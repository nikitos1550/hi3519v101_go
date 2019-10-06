// +build hi3516cv300

package main

// #cgo LDFLAGS: ${SRCDIR}/hi3516cv300/libhi3516cv300.a
// #include "hi3516cv300/include/hi_common.h"
// #include "hi3516cv300/include/mpi_sys.h"
// HI_S32 HI_MPI_SYS_GetChipId(HI_U32 *pu32ChipId);
import "C"

const (
    chipFamily = "hi3516cv300"
)

var (
    sysIdReg        uint32
    mppVersion      string

    modules = [...][2]string {
        [2]string{"sys_config.ko", "vi_vpss_online=0"},
        [2]string{"hi_osal.ko", "mmz=anonymous,0,0x{memStartAddr},{memMppSize}M anony=1"},
        [2]string{"hi3516cv300_base.ko", ""},
        [2]string{"hi3516cv300_sys.ko", "vi_vpss_online=0 sensor=NULL mem_total={memTotalSize}"},
    }
    chips = [...]string {
        "hi3516cv300",
        "hi3516ev100",
    }
)

func init() {
    sysIdReg = readDevMem32(0x12020EE0) & 0xFF
    sysIdReg = sysIdReg + ((readDevMem32(0x12020EE4) & 0xFF) << 8)
    sysIdReg = sysIdReg + ((readDevMem32(0x12020EE8) & 0xFF) << 16)
    sysIdReg = sysIdReg + ((readDevMem32(0x12020EEC) & 0xFF) << 24)

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
    writeDevMem32(0x1203009C, 0xCFA00000)
}

func getTemperature() float32 {
    var tempCode uint32 = readDevMem32(0x120300A4)
    var temp float32 = ((( float32(tempCode & 0x3FF)-125)/806)*165)-40
    return temp
}

