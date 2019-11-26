package chip

// #include "hi_common.h"
// HI_S32 HI_MPI_SYS_GetChipId(HI_U32 *pu32ChipId);
import "C"

//ATTENTION it is NOT working without loaded ko modules (minimal list) (actually /dev/sys)
func MppId() uint32 {
    var id C.HI_U32
    C.HI_MPI_SYS_GetChipId(&id)

    return uint32(id)
}


