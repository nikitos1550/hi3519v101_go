package schedule

import (
    "errors"

    "application/core/mpp/connection"
    "application/core/logger"
)

//connection.ClientRawFrame interface implementation

func (s *Schedule) RegisterRawFrameSource(source connection.SourceRawFrame, frameCompat connection.FrameCompatibility) (*chan connection.Frame, error) {
    s.Lock()
    defer s.Unlock()

    if s.sourceRaw != nil {
        return nil, errors.New("Encoder already has source")
    }

    s.params = frameCompat

    s.rawFramesCh = make(chan connection.Frame)

    s.rutineStop = make(chan bool)
    s.rutineDone = make(chan bool)

    go s.rawFramesRutine()

    s.sourceRaw = source

    return &s.rawFramesCh, nil
}

func (s *Schedule) UnregisterRawFrameSource(source connection.SourceRawFrame) error {
    s.Lock()
    defer s.Unlock()

    if s.sourceRaw != source {
        return errors.New("Encoder is not connected to this source")
    }

    if s.clientRaw != nil {
        return errors.New("Can`t unregister source, because of clients")
    }

    s.rutineStop <- true
    <-s.rutineDone

    s.sourceRaw = nil

    return nil
}

func (s *Schedule) rawFramesRutine() {
    for {
        select {
        case frame := <-s.rawFramesCh:
            s.RLock()

            //logger.Log.Trace().
            //    Uint64("pts", frame.Pts).
            //    Uint64("start", s.startTimestamp).
            //    Uint64("stop", s.stopTimestamp).
            //    Msg("scheduler frame recv")

            if frame.Pts >= s.startTimestamp && frame.Pts <= s.stopTimestamp {

                //logger.Log.Trace().Uint64("pts", frame.Pts).Msg("schedule frame")

                if s.clientRaw != nil {
                    frame.Wg.Add(1)
                    select {
                    case *s.clientCh<-frame:
                        break
                    default:
                        logger.Log.Warn().
                            Str("client", s.clientRaw.FullName()).
                            Msg("Scheduler client droppped frame")
                        frame.Wg.Add(-1)
                        break
                    }
                }
            }
            frame.Wg.Add(-1)
            s.RUnlock()
            break
        case <-s.rutineStop:
            s.rutineDone <- true
            break
        }
    }
}

