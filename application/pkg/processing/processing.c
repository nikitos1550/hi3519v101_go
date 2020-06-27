#include "processing.h"

//TODO move to venc package!!!
int sendToEncoder(error_in *err, unsigned int vencId, void* frame) {
    #if HI_MPP == 1
        DO_OR_RETURN_ERR_MPP(err, HI_MPI_VENC_SendFrame, vencId, frame);
    #elif HI_MPP >= 2
        DO_OR_RETURN_ERR_MPP(err, HI_MPI_VENC_SendFrame, vencId, frame, -1);
    #endif

    return ERR_NONE;
}
