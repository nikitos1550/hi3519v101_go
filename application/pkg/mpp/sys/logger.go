package sys

/*
#include "../../logger/logger.h"
//Should be here to export go callback
*/
import "C"

import (
        "application/pkg/logger"
)

//export go_logger_sys
func go_logger_sys(level C.int, msgC *C.char) {
        logger.CLogger("sys", int(level), C.GoString(msgC))
}


