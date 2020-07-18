package webrtc

import (
    "errors"
    "sync"
    "io"

    "application/core/mpp/connection"
    "application/core/mpp/frames"
    "application/core/logger"
)

type webrtcServer struct {
    sync.RWMutex

    deleted         bool

    clients         map[string] *WebrtcSession
    clientsMutex    sync.RWMutex

    source          connection.SourceEncodedData
    notify          chan frames.FrameItem
    rutineStop      chan bool
    rutineDone      chan bool
}

const (
    maxId   = 1024
)

var (
    servers         map[int] *webrtcServer
    serversMutex    sync.RWMutex
    LastId          int
)

func init() {
    servers = make(map[int] *webrtcServer)
}

func Init() {}

func GetById(id int) (*webrtcServer, error) {
    serversMutex.RLock()
    defer serversMutex.RUnlock()

    item, exist := servers[id]
    if !exist {
        return nil, errors.New("No such instance")
    }

    return item, nil
}

func Create() (*webrtcServer, error) {
    serversMutex.Lock()
    defer serversMutex.Unlock()

    var item webrtcServer

    id := -1
    for i:=0; i < maxId; i++ {
        _, exist := servers[i]
        if !exist {
            id = i
            break
        }
    }

    if id == -1 {
        return nil, errors.New("Max amount reached")
    }

    //item.id         = id
    item.clients    = make(map[string] *WebrtcSession)


    servers[id] = &item

    if id > LastId {
        LastId = id
    }

    logger.Log.Debug().
        Int("id", id).
        Msg("Webrtc created")

    return &item, nil
}

func Delete(w *webrtcServer) error {
    serversMutex.Lock()
    defer serversMutex.Unlock()

    for i:=0; i < maxId; i++ {
        item := servers[i]
        if w == item {
            if item.destroy() != nil {
                return errors.New("Can`t delete, because sourced")
            }

            delete(servers, i)

            return nil
        }
    }

    return errors.New("No such instance")
}

////////////////////////////////////////////////////////////////////////////////

func (w *webrtcServer) destroy() error {
    if w == nil {
        return errors.New("Null pointer")
    }

    w.Lock()
    defer w.Unlock()

    if w.deleted == true {
        logger.Log.Error().
            Msg("webrtc invoked deleted instance")
        return nil
    }

    if w.source != nil {
        return errors.New("Can`t destroy, because sourced")
    }

    w.deleted = true

    return nil
}

func (w *webrtcServer) getFrame(f frames.FrameItem, buf []byte) error {
    if w == nil {
        return errors.New("Null pointer")
    }

    w.RLock()   //TODO
    defer w.RUnlock()

    if w.deleted == true {
        logger.Log.Error().
            Msg("Webrtc invoked deleted instance")
        return errors.New("Invoked deleted instance")
    }

    if w.source == nil {
        return errors.New("Instance not sourced")
    }

    s, err := w.source.GetStorage()
    if err != nil {
        return err
    }

    _, err = s.ReadItem(f, buf)
    if err != nil {
        return err
    }

    return nil
}

func (w *webrtcServer) writeFrameTo(f frames.FrameItem, buf io.Writer) error {
    if w == nil {
        return errors.New("Null pointer")
    }

    w.RLock()   //TODO
    defer w.RUnlock()

    if w.deleted == true {
        logger.Log.Error().
            Msg("Webrtc invoked deleted instance")
        return errors.New("Invoked deleted instance")
    }

    if w.source == nil {
        return errors.New("Instance not sourced")
    }

    s, err := w.source.GetStorage()
    if err != nil {
        return err
    }

    _, err = s.WriteItemTo(f, buf)
    if err != nil {
        return err
    }

    return nil
}
