package venc

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
    MJPEG   Codec = iota + 1
    H264    Codec = iota
    H265    Codec = iota
)

type Profile uint
const (
    Baseline    Profile = iota + 1
    Main        Profile = iota
    Main10      Profile = iota
    High        Profile = iota
)

type BitrateControl uint 
const (
    Cbr     BitrateControl = iota + 1
    Vbr     BitrateControl = iota
    FixQp   BitrateControl = iota
    CVbr    BitrateControl = iota
    AVbr    BitrateControl = iota
    QVbr    BitrateControl = iota
    //Qmap    encoderBitcontrol = iota
)

type BitrateControlParams struct {
    bitrate     uint
    stattime    uint
    fluctuate   uint

    maxbitrate  uint
    //stattime    uint
    minIqp      uint
    maxqp       uint
    minqp       uint

    iqp         uint
    pqp         uint
    bqp         uint
}

type GopType uint
const (
    NormalP GopType = iota + 1
    DualP   GopType = iota
    SmartP  GopType = iota
    BipredB GopType = iota
    IntraR  GopType = iota
)

type GopParams struct {

}

type encoderSettings struct {
    //id          uint
    //channel     uint

    codec       Codec
    profile     Profile
    width       uint
    height      uint
    fps         uint    //maybe int to add -1 as uncontrolled fps 

    goptype     GopType
    gop         uint    //value
    gopparams   GopParams

    bitcontrol  BitrateControl
    bitparams   BitrateControlParams
}
