//+build !ignore,!generate
//+build koEmbed

package ko

import (
    "golang.org/x/sys/unix"
    "application/core/logger"
)

func loadModule(name string, params string) error {
    data, err := Asset(name)
    if err != nil {
        logger.Log.Fatal().
            Str("module", name).
            Str("desc", err.Error()).
            Msg("KO")
        return err
    }

    err = unix.InitModule(data, params)
    if err != nil {
        logger.Log.Error().
            Str("module", name).
            Str("params", params).
            Str("desc", err.Error()).
            Msg("KO load")
        return err
    } else {
        //logger.Log.Trace().
        //    Str("module", name).
        //    Str("params", params).
        //    Msg("KO loaded")
    }

    return nil
}
