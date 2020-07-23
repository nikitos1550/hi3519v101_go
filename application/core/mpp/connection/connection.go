package connection

//#include "../include/mpp.h"
import "C"

import (
    "unsafe"
    "sync"

    "application/core/mpp/frames"
)

type Frame struct {
    FrameMPP    C.VIDEO_FRAME_INFO_S
    Frame       unsafe.Pointer
    Pts         uint64
    Wg          *sync.WaitGroup
}

type ConnectionType int
const (
    None        ConnectionType = 1  //hold client and do nothing, useful when we maange underlaying things in some other place
    BindEncoder ConnectionType = 2  //valid only for Encoder at the moment
    BindDisplay ConnectionType = 3  //TODO
    Push        ConnectionType = 4  //valid for Encoder and Processing //TODO rename PushRawFrame
    PushEncoded ConnectionType = 5
)

type Connection struct {
    Name        string                  `json:"name"`
    CType       ConnectionType          `json:"type"`
    EncoderId   int                     `json:"-"`
    MaxWidth    int                     `json:"-"`
    MaxHeight   int                     `json:"-"`
    MaxFps      int                     `json:"-"`
    Notify      *chan frames.FrameItem  `json:"-"`
}

type Client interface {
    GetConnection() Connection
    CanAcceptFrame() bool
    PushFrame(Frame) error              //TODO rework this, make async push and release mechanics, other wise frame processing by several clients will be continious not parallel
    //PushRawMppFrame(RawMppFrame) error
    //PushH264xNal(H264xNal) error
}

type Source interface {
    AddClient(Client) error
    RemoveClient(Client) error
}
