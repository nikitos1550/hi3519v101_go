package sys

import (
    "application/pkg/mpp/cmos"
    "application/pkg/logger"
)

var (
    width int
    height int
    cnt int
)

func Init() {
    width = cmos.S.Width()
    height = cmos.S.Height()
    cnt = 10

    err := initFamily()
    if err != nil {
        logger.Log.Fatal().
            Str("error", err.Error()).
            Msg("SYS")
    }

    logger.Log.Debug().
        Msg("SYS inited")
}
