package vpss

/*
#include "../../logger/logger.h"
//Should be here to export go callback
*/
import "C"

import (
        "application/pkg/logger"
)

//export go_logger_vpss
func go_logger_vpss(level C.int, msgC *C.char) {
        logger.CLogger("VPSS", int(level), C.GoString(msgC))
}

