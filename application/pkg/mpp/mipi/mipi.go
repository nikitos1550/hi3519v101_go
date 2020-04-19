package mipi

import (
    "unsafe"
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

