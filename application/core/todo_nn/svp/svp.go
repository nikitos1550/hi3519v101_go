//+build svp
//+build hi3516cv500,hi3519av100

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
