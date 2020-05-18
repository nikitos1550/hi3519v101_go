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
    MJPEG   encoderType = iota + 1
    H264    encoderType
    H265    encoderType
)

type Profile uint
const (
    Baseline    encoderProfile = iota + 1
    Main        encoderProfile
    Main10      encoderProfile
    High        encoderProfile
)

type BitrateControl uint 
const (
    Cbr     encoderBitcontrol = iota + 1
    Vbr     encoderBitcontrol
    FixQp   encoderBitcontrol
    Avbr    encoderBitcontrol
    //Qmap    encoderBitcontrol
)

type BitrateControlParams struct {
    bitrate     uint
    stattime    uint
    fluctuate   uint

    maxbitrate  uint
    stattime    uint
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
    DualP   GopType
    SmartP  GopType
    BipredB GopType
    IntraR  GopType
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
