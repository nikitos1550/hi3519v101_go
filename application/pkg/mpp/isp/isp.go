package isp

import (
    "application/pkg/logger"
)

func Init() {
    err := initFamily()
    if err != nil {
        logger.Log.Fatal().
            Str("error", err.Error()).
            Msg("ISP")
    }
    logger.Log.Debug().
        Msg("ISP inited")

}


