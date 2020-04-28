package sys

//#include "sys.h"
import "C"

import (
    "application/pkg/mpp/cmos"
    "application/pkg/mpp/errmpp"
    "application/pkg/logger"
)

var (
    width int
    height int
    cnt int
)

func Init(chip string) {
    width = cmos.S.Width()
    height = cmos.S.Height()

    if chip == "hi3516ev100" { //TODO calc mem smart, now 32MB mpp ram only for hi3516ev100
        cnt = 5
    } else {
        cnt = 10
    }

    err := initFamily()
    if err != nil {
        logger.Log.Fatal().
            Str("error", err.Error()).
            Msg("SYS")
    }

    logger.Log.Debug().
        Msg("SYS inited")
}

func initFamily() error {
    var inErr C.error_in
    var in C.mpp_sys_init_in

    in.width = C.uint(width)
    in.height = C.uint(height)
    in.cnt = C.uint(cnt)

    logger.Log.Trace().
        Uint("width", uint(in.width)).
        Uint("height", uint(in.height)).
        Uint("cnt", uint(in.cnt)).
        Msg("SYS params")

    err := C.mpp_sys_init(&inErr, &in)
    if err != C.ERR_NONE {
        return errmpp.New(uint(inErr.f), uint(inErr.mpp))
    }

    return nil
}

//export go_logger_sys
func go_logger_sys(level C.int, msgC *C.char) {
        logger.CLogger("SYS", int(level), C.GoString(msgC))
}
