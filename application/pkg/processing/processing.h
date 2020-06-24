#pragma once 

#include "../mpp/include/mpp.h"
#include "../mpp/errmpp/errmpp.h"

//TODO move to venc package!!!
int sendToEncoder(error_in *err, unsigned int vencId, void* frame);

