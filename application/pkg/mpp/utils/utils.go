package utils

// #include "mpi_sys.h"
// HI_S32 HI_MPI_SYS_GetChipId(HI_U32 *pu32ChipId);
import "C"

func Version() string {
	var ver C.MPP_VERSION_S
	C.HI_MPI_SYS_GetVersion(&ver)
	mppVersion := C.GoString(&ver.aVersion[0])

	return mppVersion
}

func MppId() uint32 {
	var id C.HI_U32
	C.HI_MPI_SYS_GetChipId(&id)

	return uint32(id)
}
