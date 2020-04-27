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
		log.Println("processing not found", processingId)
	}

	for encoderId, _ := range processing.Encoders {
		err := C.sendToEncoder(C.uint(encoderId), frame)
		if (err != 0){
			log.Println("failed send to encoder", int(err))
		}
	}
}
