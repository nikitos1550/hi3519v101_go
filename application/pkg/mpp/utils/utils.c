#include "utils.h"

unsigned int SyncPts (HI_U64 pts) {
    #if HI_MPP == 1 \
        || HI_MPP == 2 \
        || HI_MPP == 3
        return HI_MPI_SYS_SyncPts(pts);
    #elif HI_MPP == 4
        return HI_MPI_SYS_SyncPTS(pts);
    #endif
}

unsigned int InitPtsBase (HI_U64 pts) {
    #if HI_MPP == 1 \
        || HI_MPP == 2 \
        || HI_MPP == 3
        return HI_MPI_SYS_InitPtsBase(pts);
    #elif HI_MPP == 4
        return HI_MPI_SYS_InitPTSBase(pts);
    #endif
}

