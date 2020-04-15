package logger

import (
        "github.com/rs/zerolog"
)

func Panic() *zerolog.Event {
        return Log.Panic()
}

func Fatal() *zerolog.Event {
        return Log.Fatal()
}

func Error() *zerolog.Event {
        return Log.Error()
}

func Warn() *zerolog.Event {
        return Log.Warn()
}

func Info() *zerolog.Event {
        return Log.Info()
}

func Debug() *zerolog.Event {
        return Log.Debug()
}

func Trace() *zerolog.Event {
        return Log.Trace()
}

