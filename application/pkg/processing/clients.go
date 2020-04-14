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
	//processing, exists := ActiveProcessings[int(processingId)]
	_, exists := ActiveProcessings[int(processingId)] 
	if (!exists) {
		log.Println("processing not found", int(processingId))
	}

	/* // TODO encodersId not used
	for encodersId, _ := range processing.Encoders {

	}
	*/
}