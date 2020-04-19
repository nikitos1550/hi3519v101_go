package venc

/*
#include "../../logger/logger.h"
//Should be here to export go callback
*/
import "C"

import (
	"application/pkg/logger"
)

//export go_logger_venc
func go_logger_venc(level C.int, msgC *C.char) {
	logger.CLogger("VENC", int(level), C.GoString(msgC))
}

