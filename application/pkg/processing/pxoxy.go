//+build processing

package processing

/*

#include "../mpp/include/mpp_v3.h"

void proxyCallback(unsigned int processingId, VIDEO_FRAME_INFO_S* frame) {
}

void* getCallback(){
	return proxyCallback;
}
*/
import "C"

import (
	"log"
	"unsafe"
)

func init() {
	log.Println("processing init")
	var c unsafe.Pointer
	c = C.getCallback()
	register("proxy", c)
}
