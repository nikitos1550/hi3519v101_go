package frames

import (
    "errors"
    "sync"
    "io"

    "application/core/logger"
)

type frame struct {
    rwmux   sync.RWMutex
    data    []byte
    info    FrameInfo
}

type FrameInfo struct {
    Type    uint8
    Seq     uint32
    Pts     uint64
}

//Not atomic op TODO
//func (f *frame) Info() (size int, info frameInfo, err error) {
//    f.rwmux.RLock()
//    defer f.rwmux.RUnlock()
//
//    n := len(f.data)
//
//    return n, f.info, nil
//}

//func (f *frame) Delete() (err error) {
//    f.rwmux.Lock()
//    defer f.rwmux.Unlock()
//
//    f.data = nil
//
//    return nil
//}

//func (f *frame) Write(p []byte, info frameInfo) (err error) { //write to frame from p
//    f.rwmux.Lock()
//    defer f.rwmux.Unlock()
//
//    //Naive implementation with disaster reallocs
//    f.data = make([]byte, len(p))
//    n = copy(f.data, p)
//    f.info = info
//
//    return nil
//}

func (f *frame) Writev(p [][]byte, info FrameInfo) (n int, err error) { //write to frame from multiple buffers
    f.rwmux.Lock()
    defer f.rwmux.Unlock()

    var length int
    for _,pp := range p {
        length = length + len(pp)
    }

    f.data = make([]byte, length) //TODO!!!

    var offset int
    for _,pp := range p {
        n = copy(f.data[offset:], pp)
        offset = offset + n
    }

    f.info = info

    return length, nil
}

func (f *frame) Read(buf []byte) (n int, err error) { //read from frame to buf
    f.rwmux.RLock()
    defer f.rwmux.RUnlock()

    //logger.Log.Trace().
    //    Int("buf", len(buf)).
    //    Int("data", len(f.data)).
    //    Msg("Frame")

    if len(buf) < len(f.data) {
        //logger.Log.Error().
        //    Int("buf", len(buf)).
        //    Int("data", len(f.data)).
        //    Msg("Frame")
        return 0, errors.New("Buffer is too small")
    } else {
        n = copy(buf, f.data)
    }

    return n, nil
}

func (f *frame) ReadIfEq(info FrameInfo, buf []byte) (n int, err error) { //read from frame to buf
    f.rwmux.RLock()
    defer f.rwmux.RUnlock()

    if f.info != info {
        return 0, errors.New("Frame doesn`t exist")
    }

    if len(buf) < len(f.data) {
        return 0, errors.New("Buffer is too small")
    } else {
        n = copy(buf, f.data)
    }

    return n, nil
}

func (f *frame) WriteToIfEq(info FrameInfo, w io.Writer) (int, error) {
    f.rwmux.RLock()
    defer f.rwmux.RUnlock()

    if f.info != info {
        return 0, errors.New("Frame doesn`t exist")
    }

    n, err := w.Write(f.data)
    return n, err
}

func (f *frame) ReadAlloc(buf *[]byte) (n int, err error) { //read from frame to buf
    f.rwmux.RLock()
    defer f.rwmux.RUnlock()

    *buf = make([]byte, len(f.data))

    n = copy(*buf, f.data)
    if n < len(f.data) {
        return n, errors.New("Not all data copied")
    }

    return n, nil
}

func (f *frame) ReadIfEqAlloc(info FrameInfo, buf *[]byte) (n int, err error) { //read from frame to buf
    f.rwmux.RLock()
    defer f.rwmux.RUnlock()

    if f.info != info {
        return 0, errors.New("Frame doesn`t exist")
    }

    *buf = make([]byte, len(f.data))

    n = copy(*buf, f.data)
    if n < len(f.data) {
        return n, errors.New("Not all data copied")
    }

    return n, nil
}

func (f *frame) WriteTo (w io.Writer) (n int, err error) { // read from frame to writer
    f.rwmux.RLock()
    defer f.rwmux.RUnlock()

    n, err = w.Write(f.data)
    if err != nil {
        logger.Log.Error().
            Int("n", n).
            Int("size", len(f.data)).
            Str("reason", err.Error()).
            Msg("FRAME WriteTo")
    }

    return n, err
}

//func (f *frame) ReadFrom(r io.Reader) (n int64, err error) { //write to frame from reader
//    f.mux.Lock()
//    //TODO
//    n, err = r.Read(f.data)
//    f.mux.Unlock()
//    return n, err
//}

