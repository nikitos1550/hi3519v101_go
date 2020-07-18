package jpeg

import (
    "sync"

    "github.com/pkg/errors"

    "application/core/mpp/connection"
    "application/core/mpp/frames"
)

type Jpeg struct {
    sync.RWMutex

    name    string

    source  connection.SourceEncodedData
}

func New(name string) (*Jpeg, error) {
    if name == "" {
        return nil, errors.New("Name can`t be empty")
    }

    return &Jpeg{
        name: name,
    }, nil
}

////////////////////////////////////////////////////////////////////////////////

func (j *Jpeg) Name() string {
    j.RLock()
    defer j.RUnlock()

    return j.name
}

func (j *Jpeg) FullName() string {
    j.RLock()
    defer j.RUnlock()

    return "jpeg:"+j.name
}

////////////////////////////////////////////////////////////////////////////////

func (j *Jpeg) Delete() error {
    j.Lock()
    defer j.Unlock()

    //TODO force source unregister ????????????

    if j.source != nil {
        return errors.New("Can`t destroy, because sourced")
    }

    return nil
}

////////////////////////////////////////////////////////////////////////////////

//connection.ClientEncodedData interface implementation

func (j *Jpeg) RegisterEncodedDataSource(source connection.SourceEncodedData, params connection.EncodedDataParams) (*chan frames.FrameItem, error) {
    j.Lock()
    defer j.Unlock()

    if j.source != nil {
        return nil, errors.New("Already sourced")
    }

    if params.Codec != connection.MJPEG {
        return nil, errors.New("Only codec mjpeg supported")
    }

    j.source = source

    return nil, nil
}

func (j *Jpeg) UnregisterEncodedDataSource(source connection.SourceEncodedData) error {
    j.Lock()
    defer j.Unlock()

    if j.source == nil {
        return errors.New("Not sourced")
    }

    j.source = nil

    return nil
}
