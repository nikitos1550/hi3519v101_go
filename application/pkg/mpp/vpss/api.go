package vpss

import (
    "errors"

    "application/pkg/mpp/connection"
    "application/pkg/logger"
)

func GetChannel(id int) (*Channel, error) { //TODO rename to Get
    if id >= Amount {
        return nil, errors.New("Channel id not valid")
    }

    return &channels[id], nil
}

////////////////////////////////////////////////////////////////////////////////

func (c *Channel) IsCreated() bool {
    if c == nil {
        return false
    }

    c.mutex.RLock()
    defer c.mutex.RUnlock()

    return c.Created
}

func (c *Channel) IsLocked() (bool, error) {
    if c == nil {
        return false, errors.New("Null pointer")
    }

    c.mutex.RLock()
    defer c.mutex.RUnlock()

    if c.Created == false {
        return false, errors.New("Channel is not created")
    }

    return c.Locked, nil
}

func (c *Channel) GetCopy() (Channel, error) {
    if c == nil {
		return Channel{}, errors.New("Null pointer")
	}

    c.mutex.RLock()
    defer c.mutex.RUnlock()

    if c.Created == false {
        return Channel{}, errors.New("Channel is not created")
    }

    return *c, nil
}

func (c *Channel) GetStat() (statistics, error) {
    if c == nil {
		return statistics{}, errors.New("Null pointer")
	}

    c.mutex.RLock()
    defer c.mutex.RUnlock()

    if c.Created == false {
        return statistics{}, errors.New("Channel is not created")
    }

    return c.stat, nil
}

func (c *Channel) GetParams() (Parameters, error) {
    if c == nil {
		return Parameters{}, errors.New("Null pointer")
	}

    c.mutex.RLock()
    defer c.mutex.RUnlock()

    if c.Created == false {
        return Parameters{}, errors.New("Channel is not created")
    }

    return c.Params, nil
}

////////////////////////////////////////////////////////////////////////////////

func (c *Channel) CreateChannel(params Parameters, lock bool) error  {
    if c == nil {
        return errors.New("Null pointer")
    }

    c.mutex.Lock()
    defer c.mutex.Unlock()

    if c.Created == true {
        return errors.New("Channel is already created")
    }

    err := mppCreateChannel(c.Id, params)
    if err != nil {
        return err
    }

    c.Params        = params
    c.rawClients    = make(map[connection.ClientRawFrame] bool)
    c.bindClients   = make(map[connection.ClientBind] connection.BindInformation)
    c.rutineCtrl    = make(chan bool)                                      //this also
    c.Created       = true
    c.Locked        = lock
    c.depth         = 0

    logger.Log.Debug().
        Int("id", c.Id).
        Msg("Channel created")

    return nil
}

func (c *Channel) DestroyChannel() error  {
    if c == nil {
        return errors.New("Null pointer")
    }

    c.mutex.Lock()
    defer c.mutex.Unlock()

    if c.Created == false {
        return errors.New("Channel is already destroyed")
    }

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

    c.Created   = false
    c.Locked    = false
    c.Params    = Parameters{}
    c.stat      = statistics{}

    logger.Log.Debug().
        Int("id", c.Id).
        Msg("Channel destroyed")

    return nil
}
