// +build hi3516av100

package main

// #cgo LDFLAGS: ${SRCDIR}/hi3516av100/libhi3516av100.a
// #include "hi3516av100/include/hi_common.h"
// #include "hi3516av100/include/mpi_sys.h"
// HI_S32 HI_MPI_SYS_GetChipId(HI_U32 *pu32ChipId);
import "C"

const (
    chipFamily = "hi3516av100"
)

var (
    sysIdReg        uint32
    chipDetected    string
    mppVersion      string

    modules = [...][2]string {
        [2]string{"mmz.ko", "mmz=anonymous,0,0x{memStartAddr},{memMppSize}M anony=1"},
        [2]string{"hi_media.ko", ""},
        [2]string{"hi3516a_base.ko", ""},
        [2]string{"hi3516a_sys.ko", "vi_vpss_online=0 sensor=NULL"},
    }
    chips = [...]string {
        "hi3516av100",
        "hi3516dv100",
    }
)

func init() {
    sysIdReg = readDevMem32(0x20050EE0) & 0xFF
    sysIdReg = sysIdReg + ((readDevMem32(0x20050EE4) & 0xFF) << 8)
    sysIdReg = sysIdReg + ((readDevMem32(0x20050EE8) & 0xFF) << 16)
    sysIdReg = sysIdReg + ((readDevMem32(0x20050EEC) & 0xFF) << 24)

    switch (sysIdReg) {
        case 890675456: //0x3516A100
            chipDetected = "hi3516av100"
        case 890687744: //0x3516D100
            chipDetected = "hi3516dv100"
    }

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
   //TODO
}

func getTemperature() float32 {
    var temp float32 = -999
    return temp
}


