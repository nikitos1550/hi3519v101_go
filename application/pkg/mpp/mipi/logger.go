package mipi

/*
#include "../../logger/logger.h"
//Should be here to export go callback
*/
import "C"

import (
        "application/pkg/logger"
)

//export go_logger_mipi
func go_logger_mipi(level C.int, msgC *C.char) {
        logger.CLogger("MIPI", int(level), C.GoString(msgC))
}


