package schedule

import (
    "errors"

    "application/core/mpp/connection"
)

//connection.ClientRawFrame interface implementation

func (f *Forward) RegisterRawFrameSource(source connection.SourceRawFrame, frameCompat connection.FrameCompatibility) (*chan connection.Frame, error) {
    f.Lock()
    defer f.Unlock()

    if f.sourceRaw != nil {
        return nil, errors.New("Encoder already has source")
    }

    f.params = frameCompat

    f.rawFramesCh = make(chan connection.Frame)

    f.rutineStop = make(chan bool)
    f.rutineDone = make(chan bool)

    go f.rawFramesRutine()

    f.sourceRaw = source

    return &f.rawFramesCh, nil
}

func (f *Forward) UnregisterRawFrameSource(source connection.SourceRawFrame) error {
    f.Lock()
    defer f.Unlock()

    if f.sourceRaw != source {
        return errors.New("Encoder is not connected to this source")
    }

    if f.clientRaw != nil {
        return errors.New("Can`t unregister source, because of clients")
    }

    f.rutineStop <- true
    <-f.rutineDone

    f.sourceRaw = nil

    return nil
}

func (f *Forward) rawFramesRutine() {
    for {
        select {
        case frame := <-f.rawFramesCh:
            f.RLock()
            if f.clientRaw != nil {
                frame.Wg.Add(1)
                select {
                case *f.clientCh<-frame:
                    break
                default:
                    frame.Wg.Add(-1)
                    break
                }
            }
            frame.Wg.Add(-1)
            f.RUnlock()
            break
        case <-f.rutineStop:
            f.rutineDone <- true
            break
        }
    }
}

