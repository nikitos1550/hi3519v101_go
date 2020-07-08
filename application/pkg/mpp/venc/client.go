package venc

import (
    "errors"

    "application/pkg/mpp/connection"
)

//connection.ClientRawFrame interface implementation

func (e *Encoder) RegisterRawFrameSource(source connection.SourceRawFrame, frameCompat connection.FrameCompatibility) error {
    if e == nil {
        return errors.New("Null pointer")
    }

    e.mutex.Lock()
    defer e.mutex.Unlock()

    if e.Created == false {
        return errors.New("Encoder is not created")
    }

    if e.sourceRaw != nil || e.sourceBind != nil {
        return errors.New("Encoder already has source")
    }

	if e.Params.Width > frameCompat.Width {
		return errors.New("Input frame error, width can`t be more")
	}
	if e.Params.Height > frameCompat.Height {
		return errors.New("Input frame error, height can`t be more")
    }
    if e.Params.Fps > frameCompat.Fps {
		return errors.New("Input frame error, fps can`t be more")
	}

    e.sourceRaw = source

    return nil
}

func (e *Encoder) UnregisterRawFrameSource(source connection.SourceRawFrame) error {
    if e == nil {
        return errors.New("Null pointer")
    }

    e.mutex.Lock()
    defer e.mutex.Unlock()

    if e.Created == false {
        return errors.New("Encoder is not created")
    }

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

    e.sourceRaw = nil

    return nil
}

func (e *Encoder) PushRawFrame(connection.Frame) error {
    if e == nil {
        return errors.New("Null pointer")
    }

    //TODO
    return nil
}

//connection.ClientBind interface implementation

func (e *Encoder) RegisterBindSource(source connection.SourceBind, frameCompat connection.FrameCompatibility) (connection.BindInformation, error) {
    if e == nil {
        return connection.BindInformation{}, errors.New("Null pointer")
    }

    e.mutex.Lock()
    defer e.mutex.Unlock()

    if e.Created == false {
        return connection.BindInformation{}, errors.New("Encoder is not created")
    }

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
    if e == nil {
        return errors.New("Null pointer")
    }

    e.mutex.Lock()
    defer e.mutex.Unlock()

    if e.Created == false {
        return errors.New("Encoder is not created")
    }

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
