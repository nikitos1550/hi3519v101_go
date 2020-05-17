package sys

//#include "sys.h"
import "C"

import (
    "application/pkg/mpp/vi"
    "application/pkg/mpp/errmpp"
    "application/pkg/logger"
    "application/pkg/buildinfo"
)

func Init(chip string) {
    var inErr C.error_in
    var in C.mpp_sys_init_in

    //TODO should be taken from VI!
    in.width = C.uint(vi.Width())
    in.height = C.uint(vi.Height())

    if  buildinfo.Chip == "hi3516ev100" ||
        buildinfo.Chip == "hi3518ev200" ||
        buildinfo.Chip == "hi3518ev201" ||
        buildinfo.Chip == "hi3516ev200" ||
        buildinfo.Chip == "hi3516ev300" { //TODO calc mem smart, now 32MB mpp ram only for hi3516ev100
        in.cnt = 5
    } else {
        in.cnt = 10
    }

    //ev200 sc4236 testing
    //in.cnt = 6

    logger.Log.Trace().
        Uint("width", uint(in.width)).
        Uint("height", uint(in.height)).
        Uint("cnt", uint(in.cnt)).
        Msg("SYS params")

    err := C.mpp_sys_init(&inErr, &in)
    if err != C.ERR_NONE {
        logger.Log.Fatal().
            Str("error", errmpp.New(C.GoString(inErr.name), uint(inErr.code)).Error()).
            Msg("SYS")

    }
    logger.Log.Debug().
        Msg("SYS inited")
}

//export go_logger_sys
func go_logger_sys(level C.int, msgC *C.char) {
        logger.CLogger("SYS", int(level), C.GoString(msgC))
}
