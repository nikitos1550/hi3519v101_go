package frames

/*
    Cycle buffer

*/

import (
    "errors"
    "sync"
    "io"

    "fmt"
    "application/core/logger"
)

type Frames struct {
    frames      []frame

    rwmux       sync.RWMutex
    last        int

    //firstPts    uint64
    h264        H264Info
    configured  bool
}

type H264Info struct {
    SPS []byte
    PPS []byte
}

type FrameItem struct {
    Slot    int
    Size    int
    Info    FrameInfo
}

func (f *Frames) SPS() ([]byte, error) {
    f.rwmux.RLock()
    defer f.rwmux.RUnlock()

    if f.configured == true {
        return f.h264.SPS, nil
    }
    return nil, errors.New("Not configured")
}

func (f *Frames) PPS() ([]byte, error) {
    f.rwmux.RLock()
    defer f.rwmux.RUnlock()

    if f.configured == true {
        return f.h264.PPS, nil
    }
    return nil, errors.New("Not configured")
}


func (f *Frames) configure(p [][]byte) {
    //logger.Log.Trace().Int("num NALs", len(p)).Msg("configure")
    //for i:=0;i<len(p);i++ {
    //    logger.Log.Trace().Int("NAL", i).Int("len", len(p[i])).Int("type", int(p[i][4])).Msg("configure")
    //}

    if len(p) == 4 {
        logger.Log.Trace().Msg("frame try sps pps")

        f.h264.SPS = make([]byte, len(p[0])-4)
        f.h264.SPS = p[0][4:]

        f.h264.PPS = make([]byte, len(p[1])-4)
        f.h264.PPS = p[1][4:]

        logger.Log.Trace().
            Int("len(p[0])", len(p[0])).
            Int("len(p[1])", len(p[1])).
            Int("len(p[2])", len(p[2])).
            Int("len(p[3])", len(p[3])).
            Msg("stream info")

        for i:=0;i<len(p[0]);i++ {
            fmt.Print(p[0][i], " ")
        }
        fmt.Println()
        for i:=0;i<len(p[1]);i++ {
            fmt.Print(p[1][i], " ")
        }
        fmt.Println()
        for i:=0;i<len(p[2]);i++ {
            fmt.Print(p[2][i], " ")
        }
        fmt.Println()
        for i:=0;i<10;i++ {
            fmt.Print(p[3][i], " ")
        }
        fmt.Println()

		f.configured = true
    }

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
    //f.rwmux.Lock() //TODO
    //    if f.firstPts == 0 {
    //       f.firstPts = info.Pts
    //    }
    //f.rwmux.Unlock()
    //info.Pts = info.Pts - f.firstPts

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

    if f.configured == false {
        //logger.Log.Trace().Msg("Try configure")
        f.configure(p)
    }

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

func (f *Frames) ReadLastAlloc(buf *[]byte) (int, error) {
    f.rwmux.RLock()
    var last = f.last
    f.rwmux.RUnlock()

    n, err := f.frames[last].ReadAlloc(buf)

    return n, err
}

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

func (f* Frames) ReadItemAlloc(item FrameItem, buf *[]byte) (int, error) {
    f.rwmux.RLock()
    defer f.rwmux.RUnlock()

    n, err := f.frames[item.Slot].ReadIfEqAlloc(item.Info, buf)

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

func (f *Frames) WriteItemTo(item FrameItem, w io.Writer) (int, error) {
    f.rwmux.RLock()
    defer f.rwmux.RUnlock()

    n, err := f.frames[item.Slot].WriteToIfEq(item.Info, w)

    return n, err
}


func (f *Frames) WriteLastTo(w io.Writer) (int, error) {
    f.rwmux.RLock()
    var last = f.last
    f.rwmux.RUnlock()

    n, err := f.frames[last].WriteTo(w)

    return n, err
}

