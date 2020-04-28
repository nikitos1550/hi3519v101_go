package mipi

//#include "mipi.h"
import "C"

import (
    "unsafe"
    "errors"
    "application/pkg/mpp/cmos"
    "application/pkg/logger"
)

var (
    mipi unsafe.Pointer
)

func Init() {
    mipi = cmos.S.Mipi()

    err := initFamily()
    if err != nil {
        logger.Log.Fatal().
            Str("error", err.Error()).
            Msg("MIPI")
    }
    logger.Log.Debug().
        Msg("MIPI inited")

}

func initFamily() error {
    var inErr C.error_in
    var in C.mpp_mipi_init_in

    in.mipi = mipi

    err := C.mpp_mipi_init(&inErr, &in)
    if err != C.ERR_NONE {
        return errors.New("MIPI error TODO")
    }

    return nil
}

//export go_logger_mipi
func go_logger_mipi(level C.int, msgC *C.char) {
        logger.CLogger("MIPI", int(level), C.GoString(msgC))
}
