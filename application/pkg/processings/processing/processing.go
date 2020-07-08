package processing

import (
    "errors"

    "application/pkg/logger"
    "application/pkg/mpp/connection"
)

type ProcessingMaker interface {
    Name() string
    Create() Processing
}

type Processing interface {
    Init() error
    DeInit() error

    GetConnection() connection.Connection
    CanAcceptFrame() bool
    PushFrame(connection.Frame) error
}

var (
    processingMakers map[string] ProcessingMaker
)

func init() {
    processingMakers = make(map[string] ProcessingMaker)
}

func Register(m ProcessingMaker) {
    _, exist := processingMakers[m.Name()]
    if exist {
        logger.Log.Fatal().
            Str("name", m.Name()).
            Msg("Processing duplicate name")
    }

    processingMakers[m.Name()] = m
}

func GetMaker(name string) (ProcessingMaker, error) {
    maker, exist := processingMakers[name]
    if exist {
        return maker, nil
    } else {
        return nil, errors.New("No such processing")
    }
}
