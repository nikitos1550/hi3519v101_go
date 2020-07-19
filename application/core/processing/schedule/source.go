package schedule

import (
    "errors"

    "application/core/mpp/connection"
)

func (s *Schedule) AddRawFrameClient(client connection.ClientRawFrame) error {
    s.RLock()
    defer s.RUnlock()

    if s.clientRaw != nil {
        return errors.New("Client already exist")
    }

    notificator, err := client.RegisterRawFrameSource(s, s.params)
    if err != nil {
        return err
    }

    s.clientRaw = client
    s.clientCh  = notificator

    return nil
}

func (s *Schedule) RemoveRawFrameClient(client connection.ClientRawFrame) error {
    s.RLock()
    defer s.RUnlock()

    if s.clientRaw != client {
        return errors.New("unknown client")
    }

    err := client.UnregisterRawFrameSource(s)
    if err != nil {
        return err
    }

    s.clientRaw = nil
    s.clientCh  = nil

    return nil
}

