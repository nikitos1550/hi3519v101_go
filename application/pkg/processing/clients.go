//+build processing

package processing

/*
#include "processing.h"
*/
import "C"

import (
	"log"
	"unsafe"
)

//export sendToEncoders
func sendToEncoders(processingId C.uint, frame unsafe.Pointer) {
	processing, exists := ActiveProcessings[int(processingId)]
	if (!exists) {
		log.Println("processing not found", int(processingId))
	}

	for encoderId, _ := range processing.Encoders {
		err := C.sendToEncoder(C.uint(encoderId), frame)
		if (err != 0){
			log.Println("failed send to encoder", int(err))
		}
	}
}
