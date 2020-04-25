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
