package logger

import (
	"github.com/rs/zerolog"
)

func CLogger(packageName string, level int, msg string) {
        switch level {
        case 5:
                Log.Panic().
                Str("package", packageName).
                Msg(msg)
        case 4:
                Log.Fatal().
                Str("package", packageName).
                Msg(msg)
        default:
                Log.WithLevel(zerolog.Level(level)).
                Str("package", packageName).
                Msg(msg)
        }

}

