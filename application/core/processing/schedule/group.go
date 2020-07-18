package schedule

import (
    "sync"

    "github.com/pkg/errors"

    "application/core/group"
    "application/core/mpp/connection"
)

type ForwardGroup struct {
    sync.RWMutex

    manager *group.Manager
}

func NewGroup(max uint) *ForwardGroup {
    return &ForwardGroup{
        manager: group.New(max),
    }
}

////////////////////////////////////////////////////////////////////////////////

func (g *ForwardGroup) Create(name string) error {
    g.Lock()
    defer g.Unlock()

    newForward, err := New(name)
    if err != nil {
        return errors.Wrap(err, "CreateInstance failed")
    }

    err = g.manager.Add(newForward)
    if err != nil {
        return errors.Wrap(err, "CreateInstance failed")
    }

    return nil
}

func (g *ForwardGroup) CreateForward(name string) (*Forward, error) {
    g.Lock()
    defer g.Unlock()

    newForward, err := New(name)
    if err != nil {
        return nil, errors.Wrap(err, "CreateInstance failed")
    }

    err = g.manager.Add(newForward)
    if err != nil {
        return nil, errors.Wrap(err, "CreateInstance failed")
    }

    return newForward, nil
}

func (g *ForwardGroup) Delete(name string) error {
    g.Lock()
    defer g.Unlock()

    return g.manager.Delete(name)
}

////////////////////////////////////////////////////////////////////////////////

func (g *ForwardGroup) Get(name string) (*Forward, error)  {
    g.RLock()
    defer g.RUnlock()

    forward, err := g.manager.Get(name)
    if err != nil {
        return nil, errors.Wrap(err, "GetInstanceByName failed")
    }

    return forward.(*Forward), nil
}

func (g *ForwardGroup) List() []string {
    g.RLock()
    defer g.RUnlock()

    return g.manager.List()
}

////////////////////////////////////////////////////////////////////////////////

func (g *ForwardGroup) GetClientRawFrame(name string) (connection.ClientRawFrame, error) {
    g.RLock()
    defer g.RUnlock()

    encoder, err := g.manager.Get(name)
    if err != nil {
        return nil, errors.Wrap(err, "GetInstanceByName failed")
    }

    return encoder.(connection.ClientRawFrame), nil
}

func (g *ForwardGroup) GetSourceRawFrame(name string) (connection.SourceRawFrame, error) {
    g.RLock()
    defer g.RUnlock()

    encoder, err := g.manager.Get(name)
    if err != nil {
        return nil, errors.Wrap(err, "GetInstanceByName failed")
    }

    return encoder.(connection.SourceRawFrame), nil
}

