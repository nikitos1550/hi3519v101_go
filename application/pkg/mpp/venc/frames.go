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

    rwmux   sync.RWMutex
    last    int
}

func CreateFrames(num int) *frames {
    frames := new(frames)
    frames.frames = make([]frame, num)
    log.Println("CreateFrames num=", num, " len(frames.frames)=", len(frames.frames), " cap(frames.frames)=", cap(frames.frames))
    return frames
}

func (f *frames) GetLastFrame() *frame {
    f.rwmux.RLock()
    last := f.last
    f.rwmux.RUnlock()
    return &f.frames[last]
}

func (f *frames) nextFrame() int {
    var next int
    f.rwmux.RLock() //calc next frame address
    if f.last == (cap(f.frames)-1) {
        next = 0
    } else {
        next = f.last + 1
    }
    f.rwmux.RUnlock()
    return next
}

func (f *frames) setLastFrame(frame int) {
    f.rwmux.Lock() //new frame done, let us update value
    f.last = frame
    f.rwmux.Unlock()
}

func (f *frames) WriteNext(p []byte, seq uint32) (n int, err error) {
    frame := f.nextFrame()
    n, err = f.frames[frame].Write(p, seq)
    f.setLastFrame(frame)
    return n, err
}

func (f *frames) WritevNext(p [][]byte, seq uint32) (n int, err error) {
    frame := f.nextFrame()
    n, err = f.frames[frame].Writev(p, seq)
    f.setLastFrame(frame)
    return n, err
}

/***********************/

func (f *frames) Read(buf []byte, num int) (n int, err error) {
    if num > len(f.frames) {
        return 0, nil //TODO
    }
    f.rwmux.RLock()
    n, err = f.frames[num].Read(buf)
    f.rwmux.RUnlock()
    return n, err
}

func (f *frames) ReadLast(buf []byte) (n int, err error) {
    f.rwmux.RLock()
    n, err = f.frames[f.last].Read(buf)
    f.rwmux.RUnlock()
    return n, err
}

func (f *frames) WriteTo(w io.Writer, num int) (n int, err error) {
    if num > len(f.frames) {
        return 0, nil //TODO
    }
    f.rwmux.RLock()
    n, err = f.frames[f.last].WriteTo(w)
    f.rwmux.RUnlock()
    return n, err
}

func (f *frames) WriteLastTo(w io.Writer) (n int, err error) {
    f.rwmux.RLock()
    n, err = f.frames[f.last].WriteTo(w)
    f.rwmux.RUnlock()
    return n, err
}

