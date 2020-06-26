#pragma once

#include "../include/mpp.h"
#include "../errmpp/errmpp.h"
#include "../../logger/logger.h"

//#include "mpi_sys.h"

HI_S32 HI_MPI_SYS_GetChipId(HI_U32 *pu32ChipId);  

unsigned int SyncPts (HI_U64 pts);
unsigned int InitPtsBase (HI_U64 pts);
