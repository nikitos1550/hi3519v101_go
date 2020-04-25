#ifndef _PROCESSING_H
#define _PROCESSING_H

#include "../mpp/include/mpp.h"

void sendToEncoders(unsigned int processingId, void* frame);
void sendToClients(unsigned int processingId, VIDEO_FRAME_INFO_S* frame);
int sendToEncoder(unsigned int vencId, void* frame);

#endif //_PROCESSING_H

