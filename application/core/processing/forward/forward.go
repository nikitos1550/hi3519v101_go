package forward

import (
	"sync"

    "github.com/pkg/errors"

    "application/core/mpp/connection"
)

type Forward struct {
    sync.RWMutex

	name string

    params          connection.FrameCompatibility

    sourceRaw       connection.SourceRawFrame
    rawFramesCh     chan connection.Frame
    rutineStop      chan bool
    rutineDone      chan bool

    clientRaw       connection.ClientRawFrame
    clientCh        *chan connection.Frame
}

func New(name string) (*Forward, error) {
    if name == "" {
        return nil, errors.New("Name can`t be empty")
    }

    return &Forward{
        name: name,
    }, nil
}

func (f *Forward) Delete() error  {
    f.Lock()
    defer f.Unlock()

    if f.sourceRaw != nil {
        return errors.New("Forward can`t be destroyed because sourced")
    }

    if f.clientRaw != nil {
        return errors.New("Forward can`t be destroyed because of client")
    }

    return nil
}


////////////////////////////////////////////////////////////////////////////////

func (f *Forward) Name() string {
    f.RLock()
    defer f.RUnlock()

    return f.name
}

func (f *Forward) FullName() string {
    f.RLock()
    defer f.RUnlock()

    return "forward:"+f.name
}

