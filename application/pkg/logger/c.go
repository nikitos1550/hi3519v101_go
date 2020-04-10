//+build ingore

package logger

/*
#include "logger.h"
//Should be here to export go callback
*/
import "C"

import (
	"github.com/rs/zerolog"
)

//export go_logger
func go_logger(level C.int, msgC *C.char) {
	msgGo := C.GoString(msgC)
	Log.WithLevel(zerolog.Level(level)).
		Msg(msgGo)
}

