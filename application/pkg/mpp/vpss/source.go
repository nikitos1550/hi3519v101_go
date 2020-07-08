package vpss

import (
    "errors"

    "application/pkg/mpp/connection"
    "application/pkg/logger"
    "application/pkg/mpp/utils"
)

func (c *Channel) AddRawFrameClient(client connection.ClientRawFrame) error {
    if c == nil {
        return errors.New("Null pointer")
    }

    c.mutex.RLock()
    defer c.mutex.RUnlock()

    if c.Created == false {
        return errors.New("Channel is not created")
    }

    c.rawClientsMutex.Lock()
    defer c.rawClientsMutex.Unlock()

    _, exist := c.rawClients[client]
    if exist {
        return errors.New("Client already added")
    }

    var frameCompatibility connection.FrameCompatibility = connection.FrameCompatibility{
        Width: c.Params.Width,
        Height: c.Params.Height,
        Fps: c.Params.Fps,
    }

    err := client.RegisterRawFrameSource(c, frameCompatibility)
    if err != nil {
        return err
    }

    if c.depth == 0 {
        err := mppChangeDepth(c.Id, 1)
        if err != nil {
			logger.Log.Fatal().						//After register on client side, we can`t return non nil
				Int("channel", c.Id).				//There can be situation, when we can`t set depth because of out of mpp mem
				Str("reason", err.Error()).
                Msg("VPSS channel can`t set depth")
        }

        c.depth = 1
        c.rutineStart()
    }

    c.rawClients[client] = true

    return nil
}

func (c *Channel) RemoveRawFrameClient(client connection.ClientRawFrame) error {
    if c == nil {
        return errors.New("Null pointer")
    }

    c.mutex.RLock()
    defer c.mutex.RUnlock()

    if c.Created == false {
        return errors.New("Channel is not created")
    }

    c.rawClientsMutex.Lock()
    defer c.rawClientsMutex.Unlock()

    _, exist := c.rawClients[client]
    if !exist {
        return errors.New("Client is not in list")
    }

    err := client.UnregisterRawFrameSource(c)
    if err != nil {
        return err
    }

    delete(c.rawClients, client)

    if len(c.rawClients) == 0 {
        err := c.rutineStop()
        if err != nil {
            logger.Log.Fatal().						//after delete we can`t return non nil
                Int("channel", c.Id).
                Str("reason", err.Error()).
                Msg("VPSS Can`t set depth")
        }
        c.depth = 0
    }

    return nil
}

//func (c *Channel) UnregisterRawFrameClient(client connection.ClientRawFrame) error {
//    if c == nil {
//		return errors.New("Null pointer")
//	}
//
//    c.mutex.RLock()
//    defer c.mutex.RUnlock()
//
//    if c.Created == false {
//        return errors.New("Channel is not created")
//    }
//
//    c.rawClientsMutex.Lock()
//    defer c.rawClientsMutex.Unlock()
//
//    _, exist := c.rawClients[client]
//    if !exist {
//        return errors.New("Client is not in list")
//    }
//
//    delete(c.rawClients, client)
//
//    if len(c.rawClients) == 0 {
//        err := c.rutineStop()
//        if err != nil {
//            logger.Log.Fatal().						//after delete we can`t return non nil
//                Int("channel", c.Id).
//                Str("reason", err.Error()).
//                Msg("VPSS Can`t set depth")
//        }
//        c.depth = 0
//    }
//
//    return nil
//}

////////////////////////////////////////////////////////////////////////////////

func (c *Channel) AddBindClient(client connection.ClientBind) error {
    if c == nil {
		return errors.New("Null pointer")
	}

    c.mutex.RLock()
    defer c.mutex.RUnlock()

    if c.Created == false {
        return errors.New("Channel is not created")
    }

    c.bindClientsMutex.Lock()
    defer c.bindClientsMutex.Unlock()

    _, exist := c.bindClients[client]
    if exist {
        return errors.New("Client already added")
    }

    var err error

    var frameCompatibility connection.FrameCompatibility = connection.FrameCompatibility{
        Width: c.Params.Width,
        Height: c.Params.Height,
        Fps: c.Params.Fps,
    }

    info, err := client.RegisterBindSource(c, frameCompatibility)
    if err != nil {
        return err
    }

	c.bindClients[client] = info

    switch (info.ClientType) {
    case connection.Encoder:
        err = utils.BindVpssVenc(c.Id, info.Id)
        if err != nil {
            logger.Log.Fatal().				//After register on client side, we can`t return non nil
				Int("channel", c.Id).
				Int("encoder", info.Id).
				Str("reason", err.Error()).
				Msg("VPSS can`t bind")
        }
    }

    return nil
}

func (c *Channel) RemoveBindClient(client connection.ClientBind) error {
    if c == nil {
		return errors.New("Null pointer")
	}

    c.mutex.RLock()
    defer c.mutex.RUnlock()

    if c.Created == false {
        return errors.New("Channel is not created")
    }

    c.bindClientsMutex.Lock()
    defer c.bindClientsMutex.Unlock()

    info, exist := c.bindClients[client]
    if !exist {
        return errors.New("Client is not in list")
    }

    err := client.UnregisterBindSource(c)
    if err != nil {
        return err
    }

    delete(c.bindClients, client)

    switch (info.ClientType) {
    case connection.Encoder:
        err = utils.UnBindVpssVenc(c.Id, info.Id)
        if err != nil {
			logger.Log.Fatal().             //After unregister on client side, we can`t return non nil
				Int("channel", c.Id).
				Int("encoder", info.Id).
				Str("reason", err.Error()).
				Msg("VPSS can`t unbind")
            return err
        }
    }

    return nil
}

func (c *Channel) UnregisterBindClient(client connection.ClientBind) error {
    if c == nil {
		return errors.New("Null pointer")
	}

    c.mutex.RLock()
    defer c.mutex.RUnlock()

    if c.Created == false {
        return errors.New("Channel is not created")
    }

    c.bindClientsMutex.Lock()
    defer c.bindClientsMutex.Unlock()

    info, exist := c.bindClients[client]
    if !exist {
        return errors.New("Client is not in list")
    }

    delete(c.bindClients, client)

    switch (info.ClientType) {
    case connection.Encoder:
        err := utils.UnBindVpssVenc(c.Id, info.Id)
        if err != nil {
            logger.Log.Fatal().             //After unregister on client side, we can`t return non nil
				Int("channel", c.Id).
				Int("encoder", info.Id).
				Str("reason", err.Error()).
				Msg("VPSS can`t unbind")
			return err

        }
    }

    return nil
}
