package jpeg

import (
    "errors"

    "application/pkg/mpp/connection"
    "application/pkg/mpp/frames"
    "application/pkg/logger"
)

//connection.ClientEncodedData interface implementation

func (j *jpeg) RegisterEncodedDataSource(source connection.SourceEncodedData, params connection.EncodedDataParams) error {
    if j == nil {
        return errors.New("Null pointer")
    }

    j.Lock()
    defer j.Unlock()

    if j.deleted == true {
        logger.Log.Error().
            Msg("Jpeg invoked deleted instance")
        return errors.New("Invoked deleted instance")
    }

    if j.source != nil {
        return errors.New("Already sourced")
    }

    if params.Codec != connection.MJPEG {
        return errors.New("Only codec mjpeg supported")
    }

    j.source = source

    return nil
}

func (j *jpeg) UnregisterEncodedDataSource(source connection.SourceEncodedData) error {
    if j == nil {
        return errors.New("Null pointer")
    }

    j.Lock()
    defer j.Unlock()

    if j.deleted == true {
        logger.Log.Error().
            Msg("Jpeg invoked deleted instance")
        return errors.New("Invoked deleted instance")
    }

    if j.source == nil {
        return errors.New("Not sourced")
    }

    j.source = nil

    return nil
}

func (j *jpeg) GetNotificator() *chan frames.FrameItem {
    if j == nil {
        return nil
    }

    if j.deleted == true {
        logger.Log.Error().
            Msg("Jpeg invoked deleted instance")
        return nil
    }

    return nil
}


