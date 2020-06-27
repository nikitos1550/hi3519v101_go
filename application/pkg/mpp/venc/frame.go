package venc

import (
    //"log"
    "sync"
    "io"
)

type frame struct {
    rwmux   sync.RWMutex
    data    []byte
    seq     uint32
}

//Not atomic op TODO
func (f *frame) Info() (size int, seq uint32, err error) {
    f.rwmux.RLock()
    n := len(f.data)
    f.rwmux.RUnlock()
    return n, f.seq, nil
}

func (f *frame) Delete() (err error) {
    f.rwmux.Lock()
    f.data = nil
    f.rwmux.Unlock()
    return nil
}

func (f *frame) Write(p []byte, seq uint32) (n int, err error) { //write to frame from p
    f.rwmux.Lock()
    //Naive implementation with disaster reallocs
    f.data = make([]byte, len(p))
    n = copy(f.data, p)
    f.seq = seq
    //log.Println("Frames Write buf len=", len(f.data), " cap=", cap(f.data))
    f.rwmux.Unlock()
    return n, nil
}

func (f *frame) Writev(p [][]byte, seq uint32) (n int, err error) { //write to frame from multiple buffers
    f.rwmux.Lock()
    var length int
    for _,pp := range p {
        length = length + len(pp)
    }
    f.data = make([]byte, length)
    var offset int
    for _,pp := range p {
        n = copy(f.data[offset:], pp)
        offset = offset + n
    }
    f.seq = seq
    f.rwmux.Unlock()
    return length, nil
}

func (f *frame) Read(buf []byte) (n int, err error) { //read from frame to buf
    f.rwmux.RLock()
    defer f.rwmux.RUnlock()
    if len(buf) < len(f.data) {
        return 0, nil //TODO error
    } else {
        n = copy(buf, f.data)
    }
    return n, nil
}

func (f *frame) WriteTo (w io.Writer) (n int, err error) { // read from frame to writer
    f.rwmux.RLock()
    n, err = w.Write(f.data)
    f.rwmux.RUnlock()
    return n, err
}

//func (f *frame) ReadFrom(r io.Reader) (n int64, err error) { //write to frame from reader
//    f.mux.Lock()
//    //TODO
//    n, err = r.Read(f.data)
//    f.mux.Unlock()
//    return n, err
//}

