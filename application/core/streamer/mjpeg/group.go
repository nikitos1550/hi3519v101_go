package mjpeg

import (
    "sync"

    "github.com/pkg/errors"

    "application/core/group"
    "application/core/mpp/connection"
)

type MjpegGroup struct {
    sync.RWMutex

    manager *group.Manager
}

func NewGroup(max uint) *MjpegGroup {
    return &MjpegGroup{
        manager: group.New(max),
    }
}

////////////////////////////////////////////////////////////////////////////////

func (g *MjpegGroup) Create(name string) error {
    g.Lock()
    defer g.Unlock()

    newMjpeg, err := New(name)
    if err != nil {
        return errors.Wrap(err, "CreateInstance failed")
    }

    err = g.manager.Add(newMjpeg)
    if err != nil {
        return errors.Wrap(err, "CreateInstance failed")
    }

    return nil
}

func (g *MjpegGroup) CreateMjpeg(name string) (*Mjpeg, error) {
    g.Lock()
    defer g.Unlock()

    newMjpeg, err := New(name)
    if err != nil {
        return nil, errors.Wrap(err, "CreateInstance failed")
    }

    err = g.manager.Add(newMjpeg)
    if err != nil {
        return nil, errors.Wrap(err, "CreateInstance failed")
    }

    return newMjpeg, nil
}

func (g *MjpegGroup) Delete(name string) error {
    g.Lock()
    defer g.Unlock()

    return g.manager.Delete(name)
}

////////////////////////////////////////////////////////////////////////////////

func (g *MjpegGroup) Get(name string) (*Mjpeg, error)  {
    g.RLock()
    defer g.RUnlock()

    mjpeg, err := g.manager.Get(name)

    if err != nil {
        return nil, errors.Wrap(err, "GetInstanceByName failed")
    }

    return mjpeg.(*Mjpeg), nil
}

func (g *MjpegGroup) List() []string {
    g.RLock()
    defer g.RUnlock()

    return g.manager.List()
}

////////////////////////////////////////////////////////////////////////////////

func (g *MjpegGroup) GetClientEncodedData(name string) (connection.ClientEncodedData, error) {
    g.RLock()
    defer g.RUnlock()

    mjpeg, err := g.manager.Get(name)
    if err != nil {
        return nil, errors.Wrap(err, "GetInstanceByName failed")
    }

    return mjpeg.(connection.ClientEncodedData), nil
}

