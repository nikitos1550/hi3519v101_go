package quirc

//#include "quirc.h"
//#cgo LDFLAGS: ${SRCDIR}/../../../vendors/quirc/quirc/libquirc.a
import "C"

import (
    "errors"
    "unsafe"

    "application/core/mpp/connection"
    "application/core/mpp/errmpp"
    "application/core/logger"
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
    C.quirc_quirc_init(C.int(f.params.Width), C.int(f.params.Height))

    for {
        select {
        case frame := <-f.rawFramesCh:
            f.RLock()
            if f.clientRaw != nil {
                frame.Wg.Add(1)

                var inErr C.error_in
                err := C.quirc_process(&inErr, unsafe.Pointer(&frame.FrameMPP))

                if err != 0 {
                    logger.Log.Warn().
                        Str("error", errmpp.New(C.GoString(inErr.name), uint(inErr.code)).Error()).
                        Msg("TEST")
                }


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

    C.quirc_quirc_deinit()
}

