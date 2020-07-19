package recorder

import (
    "sync"

    "github.com/pkg/errors"

    "application/archive/record"

    "application/core/mpp/connection"
    "application/core/mpp/frames"

    //"application/core/logger"
)

const (
    defaultRecvChannelSize      = 0
)

type Recorder struct {
    sync.RWMutex

    name            string

    path            string

    source          connection.SourceEncodedData
    notify          chan frames.FrameItem
    rutineStop      chan bool
    rutineDone      chan bool

    //state           state
    record          *record.Record

    lastPts         uint64
}

//type state int
//const (
//    idle        state = 1
//    scheduled   state = 2
//    recording   state = 3
//)

func New(name string, path string) (*Recorder, error) {
    if name == "" {
        return nil, errors.New("Name can`t be empty")
    }

    return &Recorder{
        name: name,
        path: path,
    }, nil
}

////////////////////////////////////////////////////////////////////////////////

func (r *Recorder) Name() string {
    r.RLock()
    defer r.RUnlock()

    return r.name
}

func (r *Recorder) FullName() string {
    r.RLock()
    defer r.RUnlock()

    return "recorder:"+r.name
}

////////////////////////////////////////////////////////////////////////////////

func (r *Recorder) Delete() error {
    r.Lock()
    defer r.Unlock()

    //TODO

    if r.source != nil {
        return errors.New("Can`t destroy, because sourced")
    }

    return nil
}

////////////////////////////////////////////////////////////////////////////////

//connection.ClientEncodedData interface implementation

func (r *Recorder) RegisterEncodedDataSource(source connection.SourceEncodedData, params connection.EncodedDataParams) (*chan frames.FrameItem, error) {
    r.Lock()
    defer r.Unlock()

    if r.source != nil {
        return nil, errors.New("Already sourced")
    }

    if params.Codec != connection.H264 {
        return nil, errors.New("Only codec h.264 supported")
    }

    r.source = source

    if defaultRecvChannelSize == 0 {
        r.notify = make(chan frames.FrameItem)
    } else {
        r.notify = make(chan frames.FrameItem, defaultRecvChannelSize)
    }

    r.rutineStop = make(chan bool)
    r.rutineDone = make(chan bool)

    go r.rutine()

    return &r.notify, nil
}

func (r *Recorder) UnregisterEncodedDataSource(source connection.SourceEncodedData) error {
    r.Lock()
    defer r.Unlock()

    if r.source == nil {
        return errors.New("Not sourced")
    }

    //TODO

    r.rutineStop <-true
    <-r.rutineDone

    r.source = nil

    return nil
}

func (r *Recorder) rutine() {
    //var tmp uint64

    for {
        select {
        case frame := <-r.notify:
            //TODO
            //logger.Log.Trace().
            //    Uint64("pts", frame.Info.Pts).
            //   Uint64("delta", frame.Info.Pts - tmp).
            //    Msg("Recorder frame recved")
            //tmp = frame.Info.Pts
            r.processFrame(frame)
            break
        case <-r.rutineStop:
            r.rutineDone <-true
            return
        }
    }
}

