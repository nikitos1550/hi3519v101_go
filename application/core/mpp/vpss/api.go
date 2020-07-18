package vpss

import (
    "errors"

    "application/core/mpp/connection"
    "application/core/logger"
)

//func GetChannel(id int) (*Channel, error) { //TODO rename to Get
//    if id >= Amount {
//        return nil, errors.New("Channel id not valid")
//    }
//
//    return &channels[id], nil
//}

////////////////////////////////////////////////////////////////////////////////

//func (c *Channel) IsCreated() bool {
//    if c == nil {
//        return false
//    }
//
//    c.mutex.RLock()
//    defer c.mutex.RUnlock()
//
//    return c.Created
//}

//func (c *Channel) IsLocked() (bool, error) {
//    if c == nil {
//        return false, errors.New("Null pointer")
//    }
//
//    c.mutex.RLock()
//    defer c.mutex.RUnlock()
//
//    if c.Created == false {
//        return false, errors.New("Channel is not created")
//    }
//
//    return c.Locked, nil
//}

func (c *Channel) GetCopy() (Channel, error) {
    c.mutex.RLock()
    defer c.mutex.RUnlock()

    //if c.Created == false {
    //    logger.Log.Warn().Msg("GetCopy channel is not created")
    //    return Channel{}, errors.New("Channel is not created")
    //}

    return *c, nil
}

func (c *Channel) GetStat() (statistics, error) {
    c.mutex.RLock()
    defer c.mutex.RUnlock()

    //if c.Created == false {
    //    return statistics{}, errors.New("Channel is not created")
    //}

    return c.stat, nil
}

func (c *Channel) GetParams() (Parameters, error) {
    c.mutex.RLock()
    defer c.mutex.RUnlock()

    //if c.Created == false {
    //    return Parameters{}, errors.New("Channel is not created")
    //}

    return c.Params, nil
}

////////////////////////////////////////////////////////////////////////////////

func (c *Channel) Name() string {
    c.mutex.RLock()
    defer c.mutex.RUnlock()

    return c.name
}

func (c*Channel) FullName() string {
    c.mutex.RLock()
    defer c.mutex.RUnlock()

    return "channel:"+c.name
}

////////////////////////////////////////////////////////////////////////////////

func New(id int, name string, params Parameters) (*Channel, error) {
    if id >= Amount {
        return nil, errors.New("Invalid id")
    }

    err := mppCreateChannel(id, params)
    if err != nil {
        return nil, err
    }

    var channel Channel

    channel.Id = id
    channel.name = name

    channel.Params        = params
    channel.rawClients    = make(map[connection.ClientRawFrame] *chan connection.Frame)
    channel.bindClients   = make(map[connection.ClientBind] connection.BindInformation)
    channel.rutineCtrl    = make(chan bool)                                      //this also
    channel.depth         = 0

    //c.bindSource    = link.NewBindSource(c, 0, c.Id)

    logger.Log.Debug().
        Int("id", channel.Id).
        Msg("Channel created")

    return &channel, nil
}

func (c *Channel) Delete() error  {
    c.mutex.Lock()
    defer c.mutex.Unlock()

    if len(c.rawClients) > 0 {     //TODO 
        return errors.New("Channel can`t be destroyed because of clients")
    }

    if len(c.bindClients) > 0 {     //TODO 
        return errors.New("Channel can`t be destroyed because of clients")
    }

    err := mppDestroyChannel(c.Id)
    if err != nil {
        return err
    }

    c.Params    = Parameters{}
    c.stat      = statistics{}

    logger.Log.Debug().
        Int("id", c.Id).
        Msg("Channel destroyed")

    return nil
}
