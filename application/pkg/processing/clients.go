//+build processing

package processing

/*
#include "processing.h"
*/
import "C"

import (
	"unsafe"

    "application/pkg/logger"
    "application/pkg/mpp/errmpp"
)

func sendToEncoders(processingId int, frame unsafe.Pointer) {
	processing, exists := ActiveProcessings[processingId]
	if (!exists) {
        logger.Log.Error().
            Int("processingId", processingId).
            Msg("Failed to send frame, processing not found")
	}

	for encoderId, _ := range processing.Encoders {
        var inErr C.error_in

		err := C.sendToEncoder(&inErr, C.uint(encoderId), frame)
		if err != C.ERR_NONE {
            logger.Log.Error().
                Str("error", errmpp.New(C.GoString(inErr.name), uint(inErr.code)).Error()).
                Msg("SYS")
		}
	}
}

func sendDataToEncoders(processingId int, data []byte) {
	processing, exists := ActiveProcessings[processingId]
	if (!exists) {
        logger.Log.Error().
            Int("processingId", processingId).
            Msg("Failed to send data, processing not found")
	}

	for _, encoder := range processing.Encoders {
		encoder.DataCallback(data)
	}
}
