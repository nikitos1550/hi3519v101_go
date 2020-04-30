package mipi

//#include "mipi.h"
import "C"

import (
    //"unsafe"
    //"errors"
    "application/pkg/mpp/cmos"
    "application/pkg/logger"
    "application/pkg/buildinfo"
)

var (
    //mipi unsafe.Pointer
)

func Init() {

    if buildinfo.Family != "hi3516cv100" {

        //mipi = cmos.S.Mipi()

        var inErr C.error_in
        var in C.mpp_mipi_init_in

        in.mipi = cmos.S.Mipi()

        err := C.mpp_mipi_init(&inErr, &in)

        if err != C.ERR_NONE {
            logger.Log.Fatal().
                //Str("error", errors.New("MIPI error TODO").Error()).
                Str("error", C.GoString(inErr.name)).
                Int("code", int(inErr.code)).
                Msg("MIPI")
        }

    }

    logger.Log.Debug().
        Msg("MIPI inited")
}

//export go_logger_mipi
func go_logger_mipi(level C.int, msgC *C.char) {
        logger.CLogger("MIPI", int(level), C.GoString(msgC))
}
