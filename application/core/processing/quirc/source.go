package quirc

import (
    "errors"

    "application/core/mpp/connection"
)

func (f *Forward) AddRawFrameClient(client connection.ClientRawFrame) error {
    f.RLock()
    defer f.RUnlock()

    if f.clientRaw != nil {
        return errors.New("Client already exist")
    }

    notificator, err := client.RegisterRawFrameSource(f, f.params)
    if err != nil {
        return err
    }

    f.clientRaw = client
    f.clientCh  = notificator

    return nil
}

func (f *Forward) RemoveRawFrameClient(client connection.ClientRawFrame) error {
    f.RLock()
    defer f.RUnlock()

    if f.clientRaw != client {
        return errors.New("unknown client")
    }

    err := client.UnregisterRawFrameSource(f)
    if err != nil {
        return err
    }

    f.clientRaw = nil
    f.clientCh  = nil

    return nil
}

