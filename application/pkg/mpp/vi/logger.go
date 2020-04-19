package vi

/*
#include "../../logger/logger.h"
//Should be here to export go callback
*/
import "C"

import (
        "application/pkg/logger"
)

////export go_logger_vi
//func go_logger_vi(level C.int, msgC *C.char) {
//        logger.CLogger("vi", int(level), C.GoString(msgC))
//}

//export go_logger_vi
func go_logger_vi(level C.int, msgC *C.char) {
        logger.CLogger("VI", int(level), C.GoString(msgC))
}


