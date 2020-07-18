package jpeg

import (
    "sync"

    "github.com/pkg/errors"

    "application/core/group"
    "application/core/mpp/connection"
)

type JpegGroup struct {
    sync.RWMutex

    manager *group.Manager
}

func NewGroup(max uint) *JpegGroup {
    return &JpegGroup{
        manager: group.New(max),
    }
}

////////////////////////////////////////////////////////////////////////////////

func (g *JpegGroup) Create(name string) error {
    g.Lock()
    defer g.Unlock()

    newJpeg, err := New(name)
    if err != nil {
        return errors.Wrap(err, "CreateInstance failed")
    }

    err = g.manager.Add(newJpeg)
    if err != nil {
        return errors.Wrap(err, "CreateInstance failed")
    }

    return nil
}

func (g *JpegGroup) CreateJpeg(name string) (*Jpeg, error) {
    g.Lock()
    defer g.Unlock()

    newJpeg, err := New(name)
    if err != nil {
        return nil, errors.Wrap(err, "CreateInstance failed")
    }

    err = g.manager.Add(newJpeg)
    if err != nil {
        return nil, errors.Wrap(err, "CreateInstance failed")
    }

    return newJpeg, nil
}

func (g *JpegGroup) Delete(name string) error {
    g.Lock()
    defer g.Unlock()

    return g.manager.Delete(name)
}

////////////////////////////////////////////////////////////////////////////////

func (g *JpegGroup) Get(name string) (*Jpeg, error)  {
    g.RLock()
    defer g.RUnlock()

    jpeg, err := g.manager.Get(name)
    if err != nil {
        return nil, errors.Wrap(err, "GetInstanceByName failed")
    }

    return jpeg.(*Jpeg), nil
}

func (g *JpegGroup) List() []string {
    g.RLock()
    defer g.RUnlock()

    return g.manager.List()
}

////////////////////////////////////////////////////////////////////////////////

func (g *JpegGroup) GetClientEncodedData(name string) (connection.ClientEncodedData, error) {
    g.RLock()
    defer g.RUnlock()

    jpeg, err := g.manager.Get(name)
    if err != nil {
        return nil, errors.Wrap(err, "GetInstanceByName failed")
    }

    return jpeg.(connection.ClientEncodedData), nil
}
