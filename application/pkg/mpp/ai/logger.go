package ai

/*
#include "../../logger/logger.h"
//Should be here to export go callback
*/
import "C"

import (
        "application/pkg/logger"
)

//export go_logger_ai
func go_logger_ai(level C.int, msgC *C.char) {
        logger.CLogger("ai", int(level), C.GoString(msgC))
}


