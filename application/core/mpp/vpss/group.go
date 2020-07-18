package vpss

import (
    "sync"

    "github.com/pkg/errors"

    "application/core/group"
    //"application/core/logger"
    "application/core/mpp/connection"
)

type ChannelGroup struct {
    sync.RWMutex

    manager *group.Manager
}

func NewGroup(max uint) *ChannelGroup {
    if max > uint(Amount) {
        max = uint(Amount)
    }

    return &ChannelGroup{
        manager: group.New(max),
    }
}

////////////////////////////////////////////////////////////////////////////////

func (g *ChannelGroup) CreateChannel(name string, params Parameters) (*Channel, error) {
    g.Lock()
    defer g.Unlock()

    if g.manager.HaveName(name) {
        return nil, errors.New("Duplicate name")
    }

    var id int = -1

    for i:=0; i < g.manager.Max(); i++ {
        valid := true
        for _, instance := range(g.manager.Instances) {
            channel := instance.(*Channel)
            if channel.Id == i {
                //logger.Log.Trace().Int("i", i).Msg("id can`t be used")
                valid = false
                break
            }
        }
        if valid == true {
            //logger.Log.Trace().Int("i", i).Msg("id found")
            id = i
            break
        }
    }

    if id == -1 {
        return nil, errors.New("No more encoders")
    }

    newChannel, err := New(id, name, params)
    if err != nil {
        return nil, errors.Wrap(err, "CreateInstance failed")
    }

    err = g.manager.Add(newChannel)
    if err != nil {
        return nil, errors.Wrap(err, "CreateInstance failed")
    }

    return newChannel, nil
}

func (g *ChannelGroup) Delete(name string) error {
    g.Lock()
    defer g.Unlock()

    return g.manager.Delete(name)
}

////////////////////////////////////////////////////////////////////////////////

func (g *ChannelGroup) Get(name string) (*Channel, error)  {
    g.RLock()
    defer g.RUnlock()

    channel, err := g.manager.Get(name)
    if err != nil {
        return nil, errors.Wrap(err, "GetInstanceByName failed")
    }

    return channel.(*Channel), nil
}

func (g *ChannelGroup) List() []string {
    g.RLock()
    defer g.RUnlock()

    return g.manager.List()
}

////////////////////////////////////////////////////////////////////////////////

func (g *ChannelGroup) GetSourceRawFrame(name string) (connection.SourceRawFrame, error) {
    g.RLock()
    defer g.RUnlock()

    encoder, err := g.manager.Get(name)
    if err != nil {
        return nil, errors.Wrap(err, "GetInstanceByName failed")
    }

    return encoder.(connection.SourceRawFrame), nil
}

func (g *ChannelGroup) GetSourceBind(name string) (connection.SourceBind, error) {
    g.RLock()
    defer g.RUnlock()

    encoder, err := g.manager.Get(name)
    if err != nil {
        return nil, errors.Wrap(err, "GetInstanceByName failed")
    }

    return encoder.(connection.SourceBind), nil
}

