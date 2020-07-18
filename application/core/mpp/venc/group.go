package venc

import (
    "sync"

    "github.com/pkg/errors"

    "application/core/group"
    //"application/core/logger"
    "application/core/mpp/connection"
)

type EncoderGroup struct {
    sync.RWMutex

    manager *group.Manager
}

func NewGroup(max uint) *EncoderGroup {
    if max > uint(EncodersAmount) {
        max = uint(EncodersAmount)
    }

    return &EncoderGroup{
        manager: group.New(max),
    }
}

////////////////////////////////////////////////////////////////////////////////

func (g *EncoderGroup) CreateEncoder(name string, params Parameters) (*Encoder, error) {
    g.Lock()
    defer g.Unlock()

    if g.manager.HaveName(name) {
        return nil, errors.New("Duplicate name")
    }

    var id int = -1

    for i:=0; i < g.manager.Max(); i++ {
        valid := true
        for _, instance := range(g.manager.Instances) {
            encoder := instance.(*Encoder)
            if encoder.Id == i {
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

    newEncoder, err := New(id, name, params)
    if err != nil {
        return nil, errors.Wrap(err, "CreateInstance failed")
    }

    err = g.manager.Add(newEncoder)
    if err != nil {
        return nil, errors.Wrap(err, "CreateInstance failed")
    }

    return newEncoder, nil
}

func (g *EncoderGroup) Delete(name string) error {
    g.Lock()
    defer g.Unlock()

    return g.manager.Delete(name)
}

////////////////////////////////////////////////////////////////////////////////

func (g *EncoderGroup) Get(name string) (*Encoder, error)  {
    g.RLock()
    defer g.RUnlock()

    encoder, err := g.manager.Get(name)
    if err != nil {
        return nil, errors.Wrap(err, "GetInstanceByName failed")
    }

    return encoder.(*Encoder), nil
}

func (g *EncoderGroup) List() []string {
    g.RLock()
    defer g.RUnlock()

    return g.manager.List()
}

////////////////////////////////////////////////////////////////////////////////

func (g *EncoderGroup) GetSourceEncodedData(name string) (connection.SourceEncodedData, error) {
    g.RLock()
    defer g.RUnlock()

    encoder, err := g.manager.Get(name)
    if err != nil {
        return nil, errors.Wrap(err, "GetInstanceByName failed")
    }

    return encoder.(connection.SourceEncodedData), nil
}

func (g *EncoderGroup) GetClientRawFrame(name string) (connection.ClientRawFrame, error) {
    g.RLock()
    defer g.RUnlock()

    encoder, err := g.manager.Get(name)
    if err != nil {
        return nil, errors.Wrap(err, "GetInstanceByName failed")
    }

    return encoder.(connection.ClientRawFrame), nil
}

func (g *EncoderGroup) GetClientBind(name string) (connection.ClientBind, error) {
    g.RLock()
    defer g.RUnlock()

    encoder, err := g.manager.Get(name)
    if err != nil {
        return nil, errors.Wrap(err, "GetInstanceByName failed")
    }

    return encoder.(connection.ClientBind), nil
}

