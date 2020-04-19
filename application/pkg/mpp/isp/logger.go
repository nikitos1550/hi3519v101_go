package isp

/*
#include "../../logger/logger.h"
//Should be here to export go callback
*/
import "C"

import (
        "application/pkg/logger"
)

//export go_logger_isp
func go_logger_isp(level C.int, msgC *C.char) {
        logger.CLogger("ISP", int(level), C.GoString(msgC))
}


