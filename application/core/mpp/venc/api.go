package venc

import (
    "github.com/pkg/errors"

    "application/core/logger"
    "application/core/mpp/connection"
    "application/core/mpp/frames"
)

func (e *Encoder) IsStarted() (bool, error) {
    e.mutex.RLock()
    defer e.mutex.RUnlock()

    return e.Started, nil
}

func (e *Encoder) GetCopy() (Encoder, error) {
    e.mutex.RLock()
    defer e.mutex.RUnlock()

    return *e, nil
}

func (e *Encoder) GetParams() (Parameters, error) {
    e.mutex.RLock()
    defer e.mutex.RUnlock()

    return e.Params, nil
}

func (e *Encoder) GetStat() (Statistics, error) {
    e.mutex.RLock()
    defer e.mutex.RUnlock()

    return e.stat, nil
}

////////////////////////////////////////////////////////////////////////////////

func (e *Encoder) Name() string {
    e.mutex.RLock()
    defer e.mutex.RUnlock()

    return e.name
}

func (e *Encoder) FullName() string {
    e.mutex.RLock()
	defer e.mutex.RUnlock()

    return "encoder:"+e.name
}

////////////////////////////////////////////////////////////////////////////////

func New(id int, name string, params Parameters) (*Encoder, error) {
    if id >= EncodersAmount {
        return nil, errors.New("Invalid id")
    }

    err := mppCreateEncoder(id, params)
    if err != nil {
        return nil, errors.Wrap(err, "New encoder failed")
    }

    var encoder Encoder

    encoder.Id          = id
    encoder.name        = name

    encoder.Params     = params
    encoder.clients    = make(map[connection.ClientEncodedData] *chan frames.FrameItem)   //TODO this is not empty after first creation

    frames.CreateFrames(&encoder.storage, 100) //TODO

    logger.Log.Debug().
        Int("id", encoder.Id).
        Msg("Encoder created")

    return &encoder, nil
}

func (e *Encoder) Delete() error {
    e.mutex.Lock()
    defer e.mutex.Unlock()

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
    e.Started    = false
    e.configured = false
    e.stat       = Statistics{}

    logger.Log.Debug().
        Int("id", e.Id).
        Msg("Encoder destroyed")

    return nil
}

////////////////////////////////////////////////////////////////////////////////

func (e *Encoder) Start() error {
    e.mutex.Lock()
    defer e.mutex.Unlock()

    var err error

    if e.sourceRaw == nil && e.sourceBind == nil {
        return errors.New("Encoder is not connected to source")
    }

    err = addToLoop(e.Id, e, e.Params.Codec)
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
    e.mutex.Lock()
    defer e.mutex.Unlock()

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
    e.mutex.RLock()
    defer e.mutex.RUnlock()

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

func (e *Encoder) Reset() error {
    e.mutex.RLock()
    defer e.mutex.RUnlock()

    if e.Started == true {
        return errors.New("Encoder is started")
    }

    err := mppReset(e.Id)
    if err != nil {
        return err
    }

    logger.Log.Trace().
        Int("id", e.Id).
        Msg("Encoder Reseted")

    return nil
}
