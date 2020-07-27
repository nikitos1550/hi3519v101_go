package mjpeg

import (
    "sync"

    "github.com/pkg/errors"

    "application/core/mpp/connection"
    "application/core/mpp/frames"
    "application/core/logger"
)

const (
    defaultRecvChannelSize      = 0
)

type Mjpeg struct {
    sync.RWMutex

    name            string

    source          connection.SourceEncodedData
    notify          chan frames.FrameItem
    rutineStop      chan bool
    rutineDone      chan bool

    clients         map[string] *mjpegClient
    clientsMutex    sync.RWMutex
}

func New(name string) (*Mjpeg, error) {
    if name == "" {
        return nil, errors.New("Name can`t be empty")
    }

    return &Mjpeg{
        name: name,
        clients: make(map[string] *mjpegClient),
    }, nil
}

////////////////////////////////////////////////////////////////////////////////

func (m *Mjpeg) Name() string {
    m.RLock()
    defer m.RUnlock()

    return m.name
}

func (m *Mjpeg) FullName() string {
    m.RLock()
    defer m.RUnlock()

    return "mjpeg:"+m.name
}

func (m *Mjpeg) Info() string {
    m.RLock()
    defer m.RUnlock()



    return ""
}

////////////////////////////////////////////////////////////////////////////////

func (m *Mjpeg) Delete() error {
    m.Lock()
    defer m.Unlock()

    if m.source != nil {
        return errors.New("Can`t destroy, because sourced")
    }

    m.clientsMutex.RLock()
    defer m.clientsMutex.RUnlock()

    if len(m.clients) > 0 {
        return errors.New("Can`t destroy, because clients")
    }

    return nil
}

////////////////////////////////////////////////////////////////////////////////

//connection.ClientEncodedData interface implementation

func (m *Mjpeg) RegisterEncodedDataSource(source connection.SourceEncodedData, params connection.EncodedDataParams) (*chan frames.FrameItem, error) {
    m.Lock()
    defer m.Unlock()

    if m.source != nil {
        return nil, errors.New("Already sourced")
    }

    if params.Codec != connection.MJPEG {
        return nil, errors.New("Only codec mjpeg supported")
    }

    m.source = source

    //in mjpeg case just forward happens, mostly no reason to set buffered channel here
    if defaultRecvChannelSize == 0 {
        m.notify = make(chan frames.FrameItem)
    } else {
        m.notify = make(chan frames.FrameItem, defaultRecvChannelSize)
    }

    m.rutineStop = make(chan bool)
    m.rutineDone = make(chan bool)

    go m.rutine()

    return &m.notify, nil
}

func (m *Mjpeg) UnregisterEncodedDataSource(source connection.SourceEncodedData) error {
    m.Lock()
    defer m.Unlock()

    if m.source == nil {
        return errors.New("Not sourced")
    }

    m.clientsMutex.RLock()
    defer m.clientsMutex.RUnlock()

    if len(m.clients) > 0 {
        return errors.New("Can`t, because of clients")
    }

    m.rutineStop <-true
    <-m.rutineDone

    m.source = nil

    return nil
}

func (m *Mjpeg) rutine() {
    //var lastPts uint64

    for {
        select {
        case frame := <-m.notify:

            //logger.Log.Trace().
            //    Uint64("delta", frame.Info.Pts-lastPts).
            //    Msg("mjpeg new frm")
            //lastPts = frame.Info.Pts

            m.clientsMutex.RLock()
            for _, client := range(m.clients) {
                select {
                case client.notify <- frame:
                    break
                default:
                    logger.Log.Warn().
                        Str("mjpeg_name", m.name).
                        Str("client", client.name).
                        Msg("Mjpeg client drop frame")
                }
            }
            m.clientsMutex.RUnlock()
        case <-m.rutineStop:
            m.rutineDone <-true
            return
        }
    }
}

