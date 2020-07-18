package venc

//#include "venc.h"
import "C"

import (
    "sync"
    "reflect"

    "application/core/mpp/connection"
    "application/core/mpp/frames"
)

type Codec int
const (
    MJPEG           Codec = 1
    H264            Codec = 2
    H265            Codec = 3
)

type Profile int
const (
    Baseline        Profile = 1
    Main            Profile = 2
    Main10          Profile = 3
    High            Profile = 4
)

type BitrateControl int
const (
    Cbr                     BitrateControl = 1
    Vbr                     BitrateControl = 2
    FixQp                   BitrateControl = 3
    CVbr                    BitrateControl = 4
    AVbr                    BitrateControl = 5
    QVbr                    BitrateControl = 6
    //Qmap    encoderBitcontrol = 7
)

const invalidValue int = int(C.INVALID_VALUE)

type BitrateControlParameters struct {
    Bitrate     int     `json:"bitrate,omitempty"`
    MaxBitrate  int     `json:"maxbitrate,omitempty"`

    StatTime    int     `json:"stattime,omitempty"`
    Fluctuate   int     `json:"fluctuate,omitempty"`

    QFactor     int     `json:"qfactor,omitempty"`
    MinQFactor  int     `json:"minqfactor,omitempty"`
    MaxQFactor  int     `json:"maxqfactor,omitempty"`

    MinIQp      int     `json:"miniqp,omitempty"`
    MaxQp       int     `json:"maxqp,omitempty"`
    MinQp       int     `json:"minqp,omitempty"`

    IQp         int     `json:"iqp,omitempty"`
    PQp         int     `json:"pqp,omitempty"`
    BQp         int     `json:"bqp,omitempty"`
}

type GopStrategyType int
const (
    NormalP                 GopStrategyType = 1
    DualP                   GopStrategyType = 2
    SmartP                  GopStrategyType = 3
    AdvSmartP               GopStrategyType = 4
    BipredB                 GopStrategyType = 5
    IntraR                  GopStrategyType = 6
)

type GopParameters struct {
    Gop         int     `json:"gop"`

    IPQpDelta   int     `json:"ipqdelta,omitempty"`
    SPInterval  int     `json:"spinterval,omitempty"`
    SPQpDelta   int     `json:"spqdelta,omitempty"`
    BgInterval  int     `json:"bginterval,omitempty"`
    BgQpDelta   int     `json:"bgqpdelta,omitempty"`
    ViQpDelta   int     `json:"viqpdelta,omitempty"`
    BFrmNum     int     `json:"bfrmnum,omitempty"`
    BQpDelta    int     `json:"bqpdelta,omitempty"`
}

type Parameters struct {
    Codec               Codec                       `json:"codec"`
    Profile             Profile                     `json:"profile"`
    Width               int                         `json:"width"`
    Height              int                         `json:"height"`
    Fps                 int                         `json:"fps"`

    GopType             GopStrategyType             `json:"goptype"`
    GopParams           GopParameters               `json:"gopparams"`

    BitControl          BitrateControl              `json:"bitratecontroltype"`
    BitControlParams    BitrateControlParameters    `json:"bitratecontrolparams"`
}

type Statistics struct {
    //TODO
}

type SourceType int
const (
    ChannelSource       SourceType = 1
    ProcessingSource    SourceType = 2
)

type Source struct {
    srcType         SourceType
    channelId       int
    processingId    int
    onWire          bool
}

type Encoder struct {
    Id              int                                             `json:"id"`
    name            string                                          //TODO

    Params          Parameters                                      `json:"parameters"` //TODO

    //Created         bool                                            `json:"created"`
    Started	        bool                                            `json:"started"`
    //Locked          bool                                            `json:"locked"`

    configured      bool                                            `json:"-"`          //TODO

    mutex           sync.RWMutex                                    `json:"-"`

    stat            Statistics                                      `json:"-"`

    //source          Source                                          `json:"source"`     //TODO
    sourceRaw       connection.SourceRawFrame
    rawFramesCh     chan connection.Frame   //TODO
    rutineStop      chan bool
    rutineDone      chan bool

    sourceBind      connection.SourceBind

    clients         map[connection.ClientEncodedData] *chan frames.FrameItem    `json:"-"`
    clientsMutex    sync.RWMutex                                                `json:"-"`

    storage         frames.Frames                                               `json:"-"`
}

//var (
//    encoders []Encoder
//)

//func init() {
    //encoders = make([]Encoder, EncodersAmount)
    //
    //for i := 0; i < EncodersAmount; i++ {
    //    encoders[i].Id = i
    //    frames.CreateFrames(&encoders[i].storage, 10)
    //}
//}

func InvalidateBitrateControlParameters(params *BitrateControlParameters) {
    v := reflect.ValueOf(params)
    if v.Kind() == reflect.Ptr {
	    v = v.Elem()
    }
    invalidateInts(v)
}

func InvalidateGopParameters(params *GopParameters) {
    v := reflect.ValueOf(params)
    if v.Kind() == reflect.Ptr {
        v = v.Elem()
    }
    invalidateInts(v)
}

func invalidateInts(v reflect.Value) {
    typeOfS := v.Type()

    for i := 0; i< v.NumField(); i++ {
        t := v.FieldByName(typeOfS.Field(i).Name)

        if t.Kind() == reflect.Int {
            t.SetInt(int64(invalidValue))
        }
    }
}
