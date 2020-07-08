//+build nobuild

package connection

/*
    Cycle buffer

*/

import (
    //"application/pkg/logger"

    "errors"
    "sync"
    "io"
)

type Frames struct {
    frames  []frame

    rwmux   sync.RWMutex
    last    int
}

type FrameItem struct {
    Slot    int
    Size    int
    Info    FrameInfo
}

//func CreateFrames(num int) *frames {
//    frames := new(frames)
//    frames.frames = make([]frame, num)
//    log.Println("CreateFrames num=", num, " len(frames.frames)=", len(frames.frames), " cap(frames.frames)=", cap(frames.frames))
//    return frames
//}

func CreateFrames(f *Frames, num int) {
    f.frames = make([]frame, num)
    //log.Println("CreateFrames num=", num, " len(frames.frames)=", len(f.frames), " cap(frames.frames)=", cap(f.frames))
}


//func (f *frames) GetLastFrame() *frame {
//    f.rwmux.RLock()
//    defer f.rwmux.RUnlock()
//
//    return &f.frames[f.last]
//}

//func (f *frames) nextFrame() int {
//    f.rwmux.RLock() //calc next frame address
//    defer f.rwmux.RUnlock()
//
//    if f.last != (cap(f.frames)-1) {
//        return f.last + 1
//    }
//    return 0
//}

//func (f *frames) setLastFrame(frame int) {
//    f.rwmux.Lock() //new frame done, let us update value
//    defer f.rwmux.Unlock()
//
//    f.last = frame
//}

//func (f *frames) WriteNextFrame(p []byte, info frameInfo) (n int, err error) {
//    next := 0
//    if f.last != (cap(f.frames)-1) {
//        next = f.last + 1
//    }
//    n, err = f.frames[next].Write(p, info)
//
//    f.rwmux.Lock()
//    defer f.rwmux.Unlock()
//
//    f.last = next
//
//    return n, err
//}

func (f *Frames) WritevNext(p [][]byte, info FrameInfo) (int, error) {
    next := 0
    if f.last != (cap(f.frames)-1) {
        next = f.last + 1
    }

    //logger.Log.Trace().
    //    Int("next", next).
    //    Msg("WritevNext found new frame")

    _, err := f.frames[next].Writev(p, info)

    if err != nil {
        //logger.Log.Trace().
        //    Int("n", n).
        //    Msg("WritevNext f.frames[next].Writev(p, info) error")
        return 0, errors.New("Write new frame error")
    }

    f.rwmux.Lock()
    f.last = next
    f.rwmux.Unlock()

    return next, err
}

/***********************/

//func (f *frames) Read(buf []byte, num int) (n int, err error) {
//    if num > len(f.frames) {
//        return 0, nil //TODO
//    }
//    f.rwmux.RLock()
//    n, err = f.frames[num].Read(buf)
//    f.rwmux.RUnlock()
//    return n, err
//}

func (f *Frames) ReadLast(buf []byte) (int, error) {
    f.rwmux.RLock()
    var last = f.last
    f.rwmux.RUnlock()

    n, err := f.frames[last].Read(buf)

    return n, err
}

func (f* Frames) ReadItem(item FrameItem, buf []byte) (int, error) {
    f.rwmux.RLock()
    defer f.rwmux.RUnlock()

    n, err := f.frames[item.Slot].ReadIfEq(item.Info, buf)

    return n, err
}

//func (f *frames) WriteTo(w io.Writer, num int) (n int, err error) {
//    if num > len(f.frames) {
//        return 0, nil //TODO
//    }
//    f.rwmux.RLock()
//    n, err = f.frames[f.last].WriteTo(w)
//    f.rwmux.RUnlock()
//    return n, err
//}

func (f *Frames) WriteLastTo(w io.Writer) (int, error) {
    f.rwmux.RLock()
    var last = f.last
    f.rwmux.RUnlock()

    n, err := f.frames[last].WriteTo(w)

    return n, err
}

