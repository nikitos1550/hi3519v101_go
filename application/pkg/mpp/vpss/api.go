package vpss

import (
    "errors"

    "application/pkg/common"
    //"application/pkg/logger"
)

func GetAmount() int {
    return channelsAmount
}

func IsStarted(id int) (bool, error) {
    if id >= channelsAmount {
        return false, errors.New("Channel id not valid")
    }

    channels[id].mutex.Lock()
    defer channels[id].mutex.Unlock()

    if channels[id].started == true {
        return true, nil
    } else {
        return false, nil
    }

}

func GetStat(id int) (statistics, error) {
    if id >= channelsAmount {
        return statistics{}, errors.New("Channel id not valid")
    }

    channels[id].mutex.Lock()
    defer channels[id].mutex.Unlock()

    if channels[id].started == false {
        return statistics{}, errors.New("Channel is stopped")
    }

    return channels[id].stat, nil
}

func GetParams(id int) (Parameters, error) {
    if id >= channelsAmount {
        return Parameters{}, errors.New("Channel id not valid")
    }

    channels[id].mutex.Lock()
    defer channels[id].mutex.Unlock()

    if channels[id].started == false {
        return Parameters{}, errors.New("Channel is stopped")
    }

    return channels[id].params, nil
}

func GetClientsTmp(id int) (map[common.Processing] bool, error) {
    if id >= channelsAmount {
        return nil, errors.New("Channel id not valid")
    }

    channels[id].mutex.Lock()
    defer channels[id].mutex.Unlock()

    if channels[id].started == false {
        return nil, errors.New("Channel is stopped")
    }

    return channels[id].clients, nil
}

func StartChannel(id int, params Parameters) error  {
    if id >= channelsAmount {
        return errors.New("Channel id not valid")
    }

    channels[id].mutex.Lock()
    defer channels[id].mutex.Unlock()

    if channels[id].started == true {
        return errors.New("Channel already started")
    }

    err := mppCreateChannel(id, params)
    if err != nil {
        return err
    }

    channels[id].params     = params
    channels[id].clients    = make(map[common.Processing] bool)
    channels[id].rutineStop = make(chan bool)
    channels[id].started    = true

    go func() {
        sendDataToClients(&channels[id])
    }()

    return nil
}

func StopChannel(id int) error  {
    if id >= channelsAmount {
        return errors.New("Channel id not valid")
    }

    channels[id].mutex.Lock()
    defer channels[id].mutex.Unlock()

    if channels[id].started == false {
        return errors.New("Channel already stopped")
    }

    if len(channels[id].clients) > 0 {
        return errors.New("Channel can`t be stopped because of clients")
    }

    err := mppDestroyChannel(id)
    if err != nil {
        return err
    }

    channels[id].started    = false
    _ = <- channels[id].rutineStop

    channels[id].params     = Parameters{}
    channels[id].stat       = statistics{}
    channels[id].clients    = nil
    channels[id].rutineStop = nil

    return nil
}

func SubscribeChannel(id int, processing common.Processing)  (int, string)  {
    if id >= channelsAmount {
        return -1, "Channel id not valid"
    }

    channels[id].mutex.Lock()
    defer channels[id].mutex.Unlock()

    if channels[id].started == false {
        return -1, "Channel does not exist"
    }

    _, callbackExists := channels[id].clients[processing]
    if (callbackExists) {
        return -1, "Already subscribed"
    }

    channels[id].clients[processing] = true

    return 0, ""
}

func UnsubscribeChannel(id int, processing common.Processing)  (int, string)  {
    if id >= channelsAmount {
        return -1, "Channel id not valid"
    }

    channels[id].mutex.Lock()
    defer channels[id].mutex.Unlock()

    if channels[id].started == false {
        return -1, "Channel does not exist"
    }

    _, callbackExists := channels[id].clients[processing]
    if (!callbackExists) {
        return -1, "Not subscribed"
    }

    delete(channels[id].clients, processing)

    return 0, ""
}
