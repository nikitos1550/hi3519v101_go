//+build processing

package processing

/*

#include "../mpp/include/mpp_v3.h"

#include "processing.h"

void proxyCallback(unsigned int processingId, VIDEO_FRAME_INFO_S* frame) {
	sendToClients(processingId, frame);
}

void* getCallback(){
	return proxyCallback;
}
*/
import "C"

import (
	"unsafe"
)

func init() {
	var c unsafe.Pointer
	c = C.getCallback()
	register("proxy", c)
}
