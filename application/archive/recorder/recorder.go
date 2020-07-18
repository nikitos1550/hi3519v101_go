package recorder

import (
    "sync"

    "github.com/pkg/errors"

    "application/core/mpp/connection"
)

const (
    defaultRecvChannelSize      = 0
)

type Recorder struct {
    sync.RWMutex

    name    string

    source          connection.SourceEncodedData
    notify          chan frames.FrameItem
    rutineStop      chan bool
    rutineDone      chan bool
}

func New(name string) (*Recorder, error) {
    if name == "" {
        return nil, errors.New("Name can`t be empty")
    }

    return &Recorder{
        name: name,
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

func (r *Mjpeg) UnregisterEncodedDataSource(source connection.SourceEncodedData) error {
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
    for {
        select {
        case frame := <-r.notify:
            //TODO
            break
        case <-r.rutineStop:
            r.rutineDone <-true
            return
        }
    }
}

