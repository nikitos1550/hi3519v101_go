package mpp

// #include "mpi_sys.h"
import "C"

func Version() string {
    var ver C.MPP_VERSION_S
    C.HI_MPI_SYS_GetVersion(&ver)
    mppVersion := C.GoString(&ver.aVersion[0])

    return mppVersion
}


