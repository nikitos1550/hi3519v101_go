package venc

import (
    //"log"
    "sync"
    "io"
)

type frame struct {
    data    []byte
    //mux     sync.Mutex
    rwmux   sync.RWMutex
    ts      uint64
}

/*
type frames struct {
    frames  []frame
    //num     int //no need golang store it inside slice object
    last    int
}


func (f *frames) nextFrame() (*frame, error) {
    return nil, nil
}

func (f *frames) lastFrame() (*frame, error) {
    return nil, nil
}
*/
/*
func (f *frame) Lock() {
    f.mux.Lock()
}

func (f *frame) Unlock() {
    f.mux.Unlock()
}
*/

//Not atomic op
func (f *frame) Size() (n int, err error) {
    f.rwmux.RLock()
    n = len(f.data)
    f.rwmux.RUnlock()
    return n, nil
}


func (f *frame) Delete() (err error) {
    f.rwmux.Lock()
    f.data = nil
    f.rwmux.Unlock()
    return nil
}

func (f *frame) Read(buf []byte) (n int, err error) { //read from frame to buf
    f.rwmux.RLock()
    if len(buf) < len(f.data) {
        return 0, nil //TODO error
    } else {
        n = copy(buf, f.data)
    }
    f.rwmux.RUnlock()
    return n, nil
}

func (f *frame) Write(p []byte) (n int, err error) { //write to frame from p
    f.rwmux.Lock()
    //Naive implementation with disaster reallocs
    //if cap(f.data) < len(p) { //ERROR
        f.data = make([]byte, len(p))
    //}
    n = copy(f.data, p)
    //log.Println("Frames Write buf len=", len(f.data), " cap=", cap(f.data))
    f.rwmux.Unlock()
    return n, nil
}

func (f *frame) Writev(p [][]byte) (n int, err error) { //write to frame from multiple buffers
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
    f.rwmux.Unlock()
    return length, nil
}

/* DEPRECATED
func (f *frame) Append(p []byte) (n int, err error) {//TOREMOVE change to multi write
    f.rwmux.Lock()
    //TODO
    tmp := make([]byte, len(f.data))
    copy(tmp, f.data)
    f.data = make([]byte, len(tmp)+len(p))
    copy(f.data, tmp)
    copy(f.data[len(tmp):], p)
    n = len(f.data)
    f.rwmux.Unlock()
    return n, nil
}
*/

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

