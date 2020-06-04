//+build svp

package svp

//#include "svp.h"
//#cgo LDFLAGS: -lstdc++
import "C"

import (
    "application/pkg/logger"
)

func Init() {
    logger.Log.Debug().
        Msg("SVP tests")

    C.svp_rt_init()

}
