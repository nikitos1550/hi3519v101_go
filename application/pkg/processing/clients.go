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

func sendToEncoders(processingId int, frame unsafe.Pointer) {
	processing, exists := ActiveProcessings[processingId]
	if (!exists) {
		log.Println("Failed to send frame, processing not found", processingId)
	}

	for encoderId, _ := range processing.Encoders {
		err := C.sendToEncoder(C.uint(encoderId), frame)
		if (err != 0){
			log.Println("failed send frame to encoder", int(err))
		}
	}
}
