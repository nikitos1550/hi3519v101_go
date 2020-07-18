package group

import (
    //"sync"

    "github.com/pkg/errors"
)

type namedInstance interface {
    Name() string
    Delete() error
}

type Manager struct {
    //sync.RWMutex

    Instances   map[string] namedInstance
    max         int
}

////////////////////////////////////////////////////////////////////////////////

func New(max uint) *Manager {
   return &Manager{
        Instances: make(map[string] namedInstance),
        max: int(max),
    }
}

////////////////////////////////////////////////////////////////////////////////

func (g *Manager) Add(i namedInstance) error {
    //g.Lock()
    //defer g.Unlock()

    if len(g.Instances) == g.max  {
        return errors.New("Max amount reached")
    }

    if _, exist := g.Instances[i.Name()]; exist {
        return errors.New("Duplicate name")
    }

    g.Instances[i.Name()] = i

    return nil
}

func (g *Manager) Delete(name string) error {
    //g.Lock()
    //defer g.Unlock()

    i, exist := g.Instances[name];
    if !exist {
        return errors.New("No such instance")
    }

    if err := i.Delete(); err != nil {
        return errors.Wrap(err, "Can`t delete from group")
    }

    delete(g.Instances, name)

    return nil
}

func (g *Manager) Get(name string) (namedInstance, error) {
    //g.RLock()
    //defer g.RUnlock()

    i, exist := g.Instances[name];
    if !exist {
        return nil, errors.New("No such instance")
    }
    return i, nil
}

func (g *Manager) List() []string {
    var names []string = make([]string, 0)

    for name, _ := range(g.Instances) {
        names = append(names, name)
    }

    return names
}

func (g *Manager) HaveName(name string) bool {
    _, exist := g.Instances[name]
    return exist
}

func (g *Manager) Amount() int {
    //g.RLock()
    //defer g.RUnlock()

    return len(g.Instances)
}

func (g *Manager) Max() int {
    //g.RLock()
    //defer g.RUnlock()

    return g.max
}
