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

    if channels[id].created == false {
        return Parameters{}, errors.New("Encoder is not created")
    }

    return channels[id].params, nil
}

func GetStat(id int) (Statistics, error) {
    if id >= channelsAmount {
        return Statistics{}, errors.New("Encoder id not valid")
    }

    channels[id].mutex.RLock()          //read lock
    defer channels[id].mutex.RUnlock()

    if channels[id].created == false {
        return Statistics{}, errors.New("Encoder is not created")
    }

    return channels[id].stat, nil
}

func CleanStat(id int) error {
    if id >= channelsAmount {
        return errors.New("Encoder id not valid")
    }

    channels[id].mutex.Lock()          //write lock
    defer channels[id].mutex.Unlock()

    if channels[id].created == false {
        return errors.New("Encoder is not created")
    }

    channels[id].stat = Statistics{}

    return nil
}

func IsCreated(id int) (bool, error) {
    if id >= channelsAmount {
        return false, errors.New("Encoder id not valid")
    }

    channels[id].mutex.RLock()          //read lock
    defer channels[id].mutex.RUnlock()

    if channels[id].created == true {
        return true, nil
    } else {
        return false, nil
    }
}

func IsStarted(id int) (bool, error) {
    if id >= channelsAmount {
        return false, errors.New("Encoder id not valid")
    }

    channels[id].mutex.RLock()          //read lock
    defer channels[id].mutex.RUnlock()

    if channels[id].created == false {
        return false, errors.New("Encoder is not created")
    }

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

    if channels[id].created == true {
        return errors.New("Encoder is already created")
    }

    err := mppCreateEncoder(id, params)
    if err != nil {
        return err
    }

    channels[id].params     = params
    channels[id].created    = true

    logger.Log.Debug().
        Int("id", id).
        Msg("Encoder created")

    return nil
}

func DestroyEncoder(id int) error {
    if id >= channelsAmount {
        return errors.New("Encoder id not valid")
    }

    channels[id].mutex.Lock()          //write lock
    defer channels[id].mutex.Unlock()

    if channels[id].created == false {
        return errors.New("Encoder is already created")
    }

    err := mppDestroyEncoder(id)
    if err != nil {
        return err
    }

    channels[id].created    = false
    channels[id].params     = Parameters{}

    logger.Log.Debug().
        Int("id", id).
        Msg("Encoder destroyed")

    return nil
}

func StartEncoder(id int) error {
    if id >= channelsAmount {
        return errors.New("Encoder id not valid")
    }

    channels[id].mutex.Lock()          //write lock
    defer channels[id].mutex.Unlock()

    if channels[id].created == false {
        return errors.New("Encoder is not created")
    }

    err := mppStartEncoder(id)
    if err != nil {
        return err
    }

    channels[id].started    = true

    logger.Log.Debug().
        Int("id", id).
        Msg("Encoder started")

    return nil
}

func StopEncoder(id int) error {
    if id >= channelsAmount {
        return errors.New("Encoder id not valid")
    }

    channels[id].mutex.Lock()          //write lock
    defer channels[id].mutex.Unlock()

    if channels[id].created == false {
        return errors.New("Encoder is not created")
    }

    err := mppStopEncoder(id)
    if err != nil {
        return err
    }

    channels[id].started    = false

    logger.Log.Debug().
        Int("id", id).
        Msg("Encoder stoped")

    return nil
}

func UpdateEncoder(id int, params Parameters) error {
    //TODO

    logger.Log.Debug().
        Int("id", id).
        Msg("Encoder updated")

    return nil
}
