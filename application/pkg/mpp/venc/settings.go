package venc

import (
    "sync"
)

/*
{
  "channel": 0,
  "codec": "mjpeg",
  "profile": "baseline",
  "width": 100,
  "height": 100,
  "fps": 1,
  "gop": {
    "mode": "normalp",
    "period": 0,
    "params": {
      "ipqdelta": 0
    }
  },
  "bitcontrol": "cbr",
  "dynamic": {
    "bitrate": 0,
    "stattime": 0,
    "fluctuate": 0
  }
}
*/

type Codec uint
const (
    MJPEG   Codec = 1
    H264    Codec = 2
    H265    Codec = 3
)

type Profile uint
const (
    Baseline    Profile = 1
    Main        Profile = 2
    Main10      Profile = 3
    High        Profile = 4
)

type BitrateControl uint 
const (
    Cbr     BitrateControl = 1
    Vbr     BitrateControl = 2
    FixQp   BitrateControl = 3
    CVbr    BitrateControl = 4
    AVbr    BitrateControl = 5
    QVbr    BitrateControl = 6
    //Qmap    encoderBitcontrol = 7
)

type BitrateControlParameters struct {
    Bitrate     uint
    //MaxBitrate  uint

    StatTime    uint
    Fluctuate   uint

    QFactor     uint
    MinQFactor  uint
    MaxQFactor  uint

    MinIQp      uint
    MaxQp       uint
    MinQp       uint

    IQp         uint
    PQp         uint
    BQp         uint
}

type GopStrategyType uint
const (
    NormalP     GopStrategyType = 1
    DualP       GopStrategyType = 2
    SmartP      GopStrategyType = 3
    AdvSmartP   GopStrategyType = 4
    BipredB     GopStrategyType = 5
    IntraR      GopStrategyType = 6
)

type GopParameters struct {
    IPQpDelta    int
    SPInterval   uint
    SPQpDelta    int
    BgInterval   uint
    BgQpDelta    int
    ViQpDelta    int
    BFrmNum      uint
    BQpDelta     int
}

type Parameters struct {
    Codec               Codec
    Profile             Profile
    Width               uint
    Height              uint
    Fps                 uint

    GopType             GopStrategyType
    Gop                 uint
    GopParams           GopParameters

    BitControl          BitrateControl
    BitControlParams    BitrateControlParameters
}

type Statistics struct {
    //TODO
}

type channel struct {
    id          int

    params      Parameters

    created     bool
	started	    bool
    locked      bool    //TODO special lock, to prevent external api control

    mutex       sync.RWMutex

    stat        Statistics
}

var (
    channels []channel
)

func init() {
    channels = make([]channel, channelsAmount)

    for i := 0; i < channelsAmount; i++ {
        channels[i].id = i
    }
}

