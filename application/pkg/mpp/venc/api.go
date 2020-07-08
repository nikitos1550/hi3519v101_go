package venc

import (
    "errors"

    "application/pkg/logger"
    "application/pkg/mpp/connection"
    "application/pkg/mpp/frames"
)

func GetEncoder(id int) (*Encoder, error) {
    if id >= EncodersAmount {
        return nil, errors.New("Encoder id not valid")
    }

    return &encoders[id], nil
}

func GetFirstEmpty() (*Encoder, error) {
    for i:=0; i < EncodersAmount; i++ {
        e, _ := GetEncoder(i)
        created := e.IsCreated()
        if created == false {
            return e, nil
        }
    }

    return nil, errors.New("No empty encoder")
}

////////////////////////////////////////////////////////////////////////////////

func (e *Encoder) IsCreated() bool {
    if e == nil {
        return false
    }

    e.mutex.RLock()
    defer e.mutex.RUnlock()

    return e.Created
}

func (e *Encoder) IsStarted() (bool, error) {
    if e == nil {
        return false, errors.New("Null pointer")
    }

    e.mutex.RLock()
    defer e.mutex.RUnlock()

    if e.Created == false {
        return false, errors.New("Encoder is not created")
    }

    return e.Started, nil
}

func (e *Encoder) IsLocked() (bool, error) {
    if e == nil {
        return false, errors.New("Null pointer")
    }

    e.mutex.RLock()
    defer e.mutex.RUnlock()

    if e.Created == false {
        return false, errors.New("Encoder is not created")
    }

    return e.Locked, nil
}

func (e *Encoder) GetCopy() (Encoder, error) {
    if e == nil {
        return Encoder{}, errors.New("Null pointer")
    }

    e.mutex.RLock()
    defer e.mutex.RUnlock()

    if e.Created == false {
        return Encoder{}, errors.New("Encoder is not created")
    }

    return *e, nil
}

func (e *Encoder) GetParams() (Parameters, error) {
    if e == nil {
        return Parameters{}, errors.New("Null pointer")
    }

    e.mutex.RLock()
    defer e.mutex.RUnlock()

    if e.Created == false {
        return Parameters{}, errors.New("Encoder is not created")
    }

    return e.Params, nil
}

func (e *Encoder) GetStat() (Statistics, error) {
    if e == nil {
        return Statistics{}, errors.New("Null pointer")
    }

    e.mutex.RLock()
    defer e.mutex.RUnlock()

    if e.Created == false {
        return Statistics{}, errors.New("Encoder is not created")
    }

    return e.stat, nil
}

////////////////////////////////////////////////////////////////////////////////

func (e *Encoder) CreateEncoder(params Parameters, lock bool) error {
    if e == nil {
        return errors.New("Null pointer")
    }

    e.mutex.Lock()
    defer e.mutex.Unlock()

    if e.Created == true {
        return errors.New("Encoder is already created")
    }

    err := mppCreateEncoder(e.Id, params)
    if err != nil {
        return err
    }

    e.Params     = params
    e.Created    = true
    e.Locked     = lock
    e.clients    = make(map[connection.ClientEncodedData] *chan frames.FrameItem)   //TODO this is not empty after first creation

    logger.Log.Debug().
        Int("id", e.Id).
        Msg("Encoder created")

    return nil
}

func (e *Encoder) DestroyEncoder() error {
    if e == nil {
        return errors.New("Null pointer")
    }

    e.mutex.Lock()
    defer e.mutex.Unlock()

    if e.Created == false {
        return errors.New("Encoder is not created")
    }

    if e.Started == true {   //TODO
        return errors.New("Encoder is started")
    }

    if len(e.clients) > 0 {
        return errors.New("Can`t delete encoder because of clients")
    }

    err := mppDestroyEncoder(e.Id)
    if err != nil {
        return err
    }

    e.Params     = Parameters{}
    e.Created    = false
    e.Started    = false
    e.Locked     = false
    e.configured = false
    e.stat       = Statistics{}

    logger.Log.Debug().
        Int("id", e.Id).
        Msg("Encoder destroyed")

    return nil
}

////////////////////////////////////////////////////////////////////////////////

func (e *Encoder) Start() error {
    if e == nil {
        return errors.New("Null pointer")
    }

    e.mutex.Lock()
    defer e.mutex.Unlock()

    if e.Created == false {
        return errors.New("Encoder is not created")
    }

    var err error

    if e.sourceRaw == nil && e.sourceBind == nil {
        return errors.New("Encoder is not connected to source")
    }

    err = addToLoop(e.Id, e.Params.Codec)
    if err != nil {
        return err
    }

    err = mppStartEncoder(e.Id)
    if err != nil {
        return err
    }

    e.Started    = true

    logger.Log.Debug().
        Int("id", e.Id).
        Msg("Encoder started")

    return nil
}

func (e *Encoder) Stop() error {
    if e == nil {
        return errors.New("Null pointer")
    }

    e.mutex.Lock()
    defer e.mutex.Unlock()

    if e.Created == false {
        return errors.New("Encoder is not created")
    }

    var err error

    err = mppStopEncoder(e.Id)
    if err != nil {
        return err
    }

    err = removeFromLoop(e.Id)
    if err != nil {
        return err
    }

    e.Started    = false

    logger.Log.Debug().
        Int("id", e.Id).
        Msg("Encoder stoped")

    return nil
}

func (e *Encoder) RequestIFrame() error {
    if e == nil {
        return errors.New("Null pointer")
    }

    e.mutex.RLock()
    defer e.mutex.RUnlock()

    if e.Created == false {
        return errors.New("Encoder is not created")
    }

    if e.Started == false {
        return errors.New("Encoder is not started")
    }

    err := mppRequestIdr(e.Id)
    if err != nil {
        return err
    }

    logger.Log.Trace().
        Int("id", e.Id).
        Msg("Encoder IDR requested")

    return nil
}
