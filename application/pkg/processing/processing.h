#ifndef _PROCESSING_H
#define _PROCESSING_H

#include "../mpp/include/mpp_v3.h"

#ifndef HI3516AV200
typedef void VIDEO_FRAME_INFO_S;
#endif


void sendToEncoders(unsigned int processingId, void* frame);
void sendToClients(unsigned int processingId, VIDEO_FRAME_INFO_S* frame);

#endif //_PROCESSING_H

