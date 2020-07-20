package venc

import (
    "errors"

    "application/core/mpp/connection"
)

//connection.ClientRawFrame interface implementation

func (e *Encoder) RegisterRawFrameSource(source connection.SourceRawFrame, frameCompat connection.FrameCompatibility) (*chan connection.Frame, error) {
    e.mutex.Lock()
    defer e.mutex.Unlock()

    if e.sourceRaw != nil || e.sourceBind != nil {
        return nil, errors.New("Encoder already has source")
    }

	if e.Params.Width > frameCompat.Width {
		return nil, errors.New("Input frame error, width can`t be more")
	}
	if e.Params.Height > frameCompat.Height {
		return nil, errors.New("Input frame error, height can`t be more")
    }
    if e.Params.Fps > frameCompat.Fps {
		return nil, errors.New("Input frame error, fps can`t be more")
	}

    e.rawFramesCh = make(chan connection.Frame)

    e.rutineStop = make(chan bool)
    e.rutineDone = make(chan bool)

    go e.rawFramesRutine()

    e.sourceRaw = source

    return &e.rawFramesCh, nil
}

func (e *Encoder) UnregisterRawFrameSource(source connection.SourceRawFrame) error {
    e.mutex.Lock()
    defer e.mutex.Unlock()

    if e.sourceRaw != source {
        return errors.New("Encoder is not connected to this source")
    }

	if len(e.clients) > 0 {
		return errors.New("Can`t unregister source, because of clients")
	}

    e.clientsMutex.RLock()
    defer e.clientsMutex.RUnlock()

    if len(e.clients) > 0 {
        return errors.New("Can`t unregister, because of clients")
    }

    e.rutineStop <- true
    <-e.rutineDone

    e.sourceRaw = nil

    return nil
}

func (e *Encoder) rawFramesRutine() {
    for {
        select {
        case frame := <-e.rawFramesCh:
            //logger.Log.Trace().Uint64("pts", frame.Pts).Msg("VENC Wg done")
            if e.Started {
                mppSendFrameToEncoder(e.Id, frame)
            }
            frame.Wg.Add(-1) //frame.Wg.Done()
            break
        case <-e.rutineStop:
            e.rutineDone <- true
            break
        }
    }
}

//connection.ClientBind interface implementation

func (e *Encoder) RegisterBindSource(source connection.SourceBind, frameCompat connection.FrameCompatibility) (connection.BindInformation, error) {
    e.mutex.Lock()
    defer e.mutex.Unlock()

    if e.sourceRaw != nil || e.sourceBind != nil {
        return connection.BindInformation{}, errors.New("Encoder already has source")
    }

    if e.Params.Width > frameCompat.Width {
		return connection.BindInformation{}, errors.New("Input frame error, width can`t be more")
	}
	if e.Params.Height > frameCompat.Height {
		return connection.BindInformation{}, errors.New("Input frame error, height can`t be more")
	}
	if e.Params.Fps > frameCompat.Fps {
		return connection.BindInformation{}, errors.New("Input frame error, fps can`t be more")
	}

    e.sourceBind = source

    var info connection.BindInformation = connection.BindInformation {
        ClientType: connection.Encoder,
        Id: e.Id,
    }
    return info, nil
}

func (e *Encoder) UnregisterBindSource(source connection.SourceBind) error {
    e.mutex.Lock()
    defer e.mutex.Unlock()

    if e.sourceBind != source {
        return errors.New("Encoder is not connected to this source")
    }

    e.clientsMutex.RLock()
    defer e.clientsMutex.RUnlock()

    if len(e.clients) > 0 {
		return errors.New("Can`t unregister source, because of clients")
	}

    e.sourceBind = nil

    return nil
}