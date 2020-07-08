#pragma once

#include "../include/mpp.h"
#include "../errmpp/errmpp.h"
#include "../../logger/logger.h"

//#include "mpi_sys.h"

HI_S32 HI_MPI_SYS_GetChipId(HI_U32 *pu32ChipId);  

unsigned int SyncPts (HI_U64 pts);
unsigned int InitPtsBase (HI_U64 pts);

typedef struct mpp_bind_vpss_venc_in_struct {
    int vpss_id;
    int venc_id;
} mpp_bind_vpss_venc_in;

int mpp_bind_vpss_venc(error_in *err, mpp_bind_vpss_venc_in *in);
int mpp_unbind_vpss_venc(error_in *err, mpp_bind_vpss_venc_in *in);
