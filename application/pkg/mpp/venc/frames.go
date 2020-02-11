package venc

/*
    Cycle buffer

*/

import (
    "log"
    "sync"
    "io"
)

type frames struct {
    frames  []frame
    //num     int //no need golang store it inside slice object
    last    int
    rwmux   sync.RWMutex
}

func CreateFrames(num int) *frames {
    frames := new(frames)
    frames.frames = make([]frame, num)
    log.Println("CreateFrames num=", num, " len(frames.frames)=", len(frames.frames), " cap(frames.frames)=", cap(frames.frames))
    return frames
}


func (f *frames) Write(p []byte) (n int, err error) {
    var last int
    f.rwmux.RLock() //calc next frame address
    if f.last == (cap(f.frames)-1) {
        last = 0
    } else {
        last = f.last + 1
    }
    f.rwmux.RUnlock()
    //log.Println("Frames Write ", last)
    n, err = f.frames[last].Write(p)
    f.rwmux.Lock() //new frame done, let us update value
    f.last = last
    f.rwmux.Unlock()
    return n, err
}

func (f *frames) Append(p []byte) (n int, err error) { //TOREMOVE change for multi write
    f.rwmux.Lock()
    //log.Println("Frames Append ", f.last)
    n, err = f.frames[f.last].Append(p)
    f.rwmux.Unlock()
    return n, err
}

func (f *frames) WriteTo(w io.Writer) (n int, err error) {
    f.rwmux.RLock()
    //log.Println("Frames WriteTo ", f.last)
    n, err = f.frames[f.last].WriteTo(w)
    f.rwmux.RUnlock()
    return n, err
}

func (f *frames) Read(buf []byte) (n int, err error) {
    f.rwmux.RLock()
    n, err = f.frames[f.last].Read(buf)
    f.rwmux.RUnlock()
    return n, err
}


//func (f *frames) lastFrame() (*frame, error) {
//    return nil, nil
//}


