package mjpeg

import (
    "github.com/pkg/errors"

    "application/core/mpp/frames"
    "application/core/logger"
)

const (
    defaultClientChannelSize    = 1
    maxClients = 8
)

type mjpegClient struct {
    name    string

    source  *Mjpeg

    notify  chan frames.FrameItem
    stop    chan bool

    dropped bool
}

func (m *Mjpeg) Clients() ([]string) {
    m.RLock()
    defer m.RUnlock()

    m.clientsMutex.Lock()
    defer m.clientsMutex.Unlock()

    var clients []string = make([]string, 0)

    for name, _ := range(m.clients) {
        clients = append(clients, name)
    }

    return clients
}

func (m *Mjpeg) newClient(name string) (*mjpegClient, error) {
    m.RLock()
    defer m.RUnlock()

    if m.source == nil {
        return nil, errors.New("Not sourced")
    }

    m.clientsMutex.Lock()
    defer m.clientsMutex.Unlock()

    if len(m.clients) >= maxClients {
        return nil, errors.New("Too much clients")
    }

    newClient := &mjpegClient{
        name:   name,
        stop:   make(chan bool),
        source: m,
    }

    if defaultClientChannelSize == 0 {
        newClient.notify = make(chan frames.FrameItem)
    } else {
        newClient.notify = make(chan frames.FrameItem, defaultClientChannelSize)
    }

    m.clients[name] = newClient

    return newClient, nil
}

func (m *Mjpeg) removeClient(name string) error {
    m.RLock()
    defer m.RUnlock()

    m.clientsMutex.Lock()
    defer m.clientsMutex.Unlock()

    c, exist := m.clients[name]
    if !exist {
        return errors.New("Client is not in list")
    }

    delete(m.clients, name)

    logger.Log.Trace().
        Str("name", c.name).
        Msg("MJPEG client removed")

    return nil
}

func (m *Mjpeg) DropClient(name string) error {
    m.RLock()
    defer m.RUnlock()

    m.clientsMutex.Lock()
    defer m.clientsMutex.Unlock()

    c, exist := m.clients[name]
    if !exist {
        return errors.New("Client is not in list")
    }

    //if c.dropped == true {
    //    return errors.New("Already dropping")
    //}

    //c.dropped = true

    c.stop <- true

    return nil
}
