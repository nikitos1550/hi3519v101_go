package venc

import (
    "errors"

    "application/pkg/logger"
)

func GetAmount() int {
    return channelsAmount
}

func GetParams(id int) (Parameters, error) {
    if id >= channelsAmount {
        return Parameters{}, errors.New("Encoder id not valid")
    }

    channels[id].mutex.RLock()          //read lock
    defer channels[id].mutex.RUnlock()

    if channels[id].started == false {
        return Parameters{}, errors.New("Encoder is stopped")
    }

    return channels[id].params, nil
}

func IsStarted(id int) (bool, error) {
    if id >= channelsAmount {
        return false, errors.New("Encoder id not valid")
    }

    channels[id].mutex.RLock()          //read lock
    defer channels[id].mutex.RUnlock()

    if channels[id].started == true {
        return true, nil
    } else {
        return false, nil
    }
}

func CreateEncoder(id int, params Parameters) error {
    if id >= channelsAmount {
        return errors.New("Encoder id not valid")
    }

    channels[id].mutex.Lock()          //write lock
    defer channels[id].mutex.Unlock()

    if channels[id].started == true {
        return errors.New("Encoder is already started")
    }

    err := mppCreateEncoder(id, params)
    if err != nil {
        return err
    }

    channels[id].params     = params
    channels[id].started    = true

    logger.Log.Debug().
        Int("id", id).
        Msg("Encoder started")

    return nil
}

func DestroyEncoder(id int) error {
    if id >= channelsAmount {
        return errors.New("Encoder id not valid")
    }

    channels[id].mutex.Lock()          //write lock
    defer channels[id].mutex.Unlock()

    if channels[id].started == false {
        return errors.New("Encoder is already stopped")
    }

    err := mppDestroyEncoder(id)
    if err != nil {
        return err
    }

    channels[id].started    = false
    channels[id].params     = Parameters{}

    logger.Log.Debug().
        Int("id", id).
        Msg("Encoder stopped")

    return nil
}
