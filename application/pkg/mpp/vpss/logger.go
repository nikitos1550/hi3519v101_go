package vpss

/*
#include "logger.h"
//Should be here to export go callback
*/
import "C"

import (
	"application/pkg/logger"
        "github.com/rs/zerolog"
)

//export go_logger
func go_logger(level C.int, msgC *C.char) {
        msgGo := C.GoString(msgC)
        logger.Log.WithLevel(zerolog.Level(level)).
                Msg(msgGo)
}

