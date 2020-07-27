package record

import (
    "os"
    "encoding/json"
    "io/ioutil"
    "time"

    "github.com/pkg/errors"

    //"application/archive/mp4"
    "application/archive/ts"

    //"github.com/nareix/joy4/format/mp4"
    "github.com/nareix/joy4/av"
    "github.com/nareix/joy4/codec/h264parser"

    "application/core/logger"
)

type Record struct {
    Dir                 string              `json:"-"`
    Name                string              `json:"name"`

    State		        RecordState         `json:"-"`

    Codec               string              `json:"codec"`

    FirstPts            uint64              `json:"first_pts"`
    LastPts             uint64              `json:"last_pts"`
    FrameCount          uint64              `json:"frames"`
    Fps                 int

    currentChunk        int
    currentChunkFile    *os.File
    Chunks              []Chunk             `json:"chunks"`

    //mp4                 *mp4.Muxer
    //mp4test             *os.File
    ts                  *ts.Muxer
    tstest              *os.File
    tstime              uint64

    Preview             bool                `json:"preview"`
}

type RecordState int
const (
	Created		RecordState	= 1
	Finished	RecordState = 2
)

type Chunk struct {
    Id          int                 `json:"id"`

    FirstPts    uint64              `json:"first_pts"`
    LastPts     uint64              `json:"last_pts"`
    FrameCount  uint64              `json:"frames"`
    Size        uint64              `json:"size"`

    //frames      []Frame

    MD5			string              `json:"md5"`
}

type Frame struct {
    offset  uint64
}

func New(dir string, name string, codec string) (*Record, error) {

    time1 := time.Now()

    stat, err := os.Stat(dir)
    if err != nil {
        return nil, errors.Wrap(err, "New failed")
    }

    if stat.IsDir() == false {
        return nil, errors.New("Dir is not dir")
    }

    time2 := time.Now()

    stat, err = os.Stat(dir+"/"+name)
    if err == nil {
        return nil, errors.New("Name is already exist")
    }

    time3 := time.Now()
    
    err = os.Mkdir(dir+"/"+name, 0755)
    if err != nil {
        return nil, errors.New("Can`t create dir")
    }

    time4 := time.Now()

    f, err := os.Create(dir+"/"+name+"/info.json")
    if err != nil {
        return nil, errors.New("Can`t create info")
    }

    time5 := time.Now()

    var rec Record = Record{
        Dir: dir,
        Name: name,
        currentChunk: 1,
    }
    rec.Chunks = make([]Chunk, 1)
    rec.Chunks[0] = Chunk{Id:1,}
    logger.Log.Debug().
        Int("length", len(rec.Chunks)).Msg("Chunks length")

    json, _ := json.Marshal(rec)
    f.Write(json)
    f.Close()

    rec.Codec = codec

    rec.currentChunkFile, err = os.Create(dir+"/"+name+"/1."+rec.Codec)
    if err != nil {
        logger.Log.Fatal().Str("reson", err.Error()).Msg("Can`t create chunk file")
    }

    time6 := time.Now()

    logger.Log.Trace().
        Time("t1", time1).
        Time("t2", time2).
        Time("t3", time3).
        Time("t4", time4).
        Time("t5", time5).
        Time("t6", time6).
        Msg("record create timings")

    //rec.mp4test, err = os.Create(dir+"/"+name+"/1.mp4")
    //if err != nil {
    //    logger.Log.Fatal().Str("reson", err.Error()).Msg("Can`t create mp4 file")
    //}

    //rec.tstest, err = os.Create(dir+"/"+name+"/1.ts")
    //if err != nil {
    //    logger.Log.Fatal().Str("reson", err.Error()).Msg("Can`t create mp4 file")
    //}

    //rec.mp4 = mp4.NewMuxer(rec.mp4test)
    //rec.ts  = ts.NewMuxer(rec.tstest)

    //var tmp h264parser.CodecData
    //var streams []av.CodecData
    //streams = make([]av.CodecData, 1)
    //streams[0] = tmp

    //err = rec.mp4.WriteHeader(streams)
    //if err != nil {
    //    //TODO
    //}

    //err = rec.ts.WriteHeader(streams)
    //if err != nil {
    //    //TODO
    //}

    return &rec, nil
}

func Load(path string, dir string) (*Record, error) {
    stat, err := os.Stat(path+"/"+dir)
    if err != nil {
        return nil, errors.Wrap(err, "No such dir")
    }

	if stat.IsDir() == false {
		return nil, errors.New("Dir is not dir")
    }

    stat, err = os.Stat(path+"/"+dir+"/info.json")
    if err != nil {
		return nil, errors.Wrap(err, "No info")
    }

	//f, err := os.Open(path+"/info.json")
	//if err != nil {
	//	return nil, errors.New("Can`t open")
	//}

    var rec Record

    file, err := ioutil.ReadFile(path+"/"+dir+"/info.json")
    if err != nil {
        return nil, errors.Wrap(err, "Load failed")
    }
    err = json.Unmarshal([]byte(file), &rec)
    if err != nil {
        return nil, errors.Wrap(err, "Load json failed")
    }
    rec.Dir = path

    //f, err := os.Open(r.Dir+"/"+r.Name+"/info.json")
	//file.Close()

    return &rec, nil
}

////////////////////////////////////////////////////////////////////////////////

func (r *Record) Close() error {

    r.currentChunkFile.Sync()
    r.currentChunkFile.Close()

    //f, err := os.Open(r.Dir+"/"+r.Name+"/info.json")
    f, err := os.OpenFile(r.Dir+"/"+r.Name+"/info.json", os.O_RDWR, 0600)
    if err != nil {
        return errors.New("Can`t open")
    }

    json, _ := json.Marshal(r)

    f.Truncate(0)
    _, err = f.Write(json)
    if err != nil {
        logger.Log.Warn().
            Str("reason", err.Error()).
            Msg("record close error")
    }
    f.Close()

    //err = r.mp4.WriteTrailer()
    //if err != nil {
    //    //TODO
    //    logger.Log.Warn().
    //        Str("reason", err.Error()).
    //        Msg("close mp4")
    //}

    //r.mp4test.Sync()
    //r.mp4test.Close()

    if r.ts != nil {
        err = r.ts.WriteTrailer()
        if err != nil {
            //TODO
            logger.Log.Warn().
                Str("reason", err.Error()).
                Msg("close mp4")
        }

        r.tstest.Sync()
        r.tstest.Close()

    }

    return nil
}

func (r *Record) SetPreview(jpeg []byte) error {
    f, err := os.Create(r.Dir+"/"+r.Name+"/preview.jpeg")
    if err != nil {
        return errors.New("Can`t create preview.jpeg")
    }

    n, err := f.Write(jpeg)
    if err != nil {
        logger.Log.Warn().
            Int("total", len(jpeg)).
            Int("wrote", n).
            Msg("preview write error")
    }
    f.Sync()
    f.Close()

    r.Preview = true

    return nil
}

func (r *Record) ConfigureTs(sps, pps []byte) error {
    if r.ts != nil {
        return errors.New("Already configured")
    }

    var err error

    r.tstest, err = os.Create(r.Dir+"/"+r.Name+"/1.ts")
    if err != nil {
        logger.Log.Fatal().Str("reson", err.Error()).Msg("Can`t create mp4 file")
    }

    r.ts  = ts.NewMuxer(r.tstest)

    tmp, err := h264parser.NewCodecDataFromSPSAndPPS(sps, pps)
    if err != nil {
        logger.Log.Fatal().Str("reason", err.Error()).Msg("cat init ts muxer")
    }

    var streams []av.CodecData
    streams = make([]av.CodecData, 1)
    streams[0] = tmp

    err = r.ts.WriteHeader(streams)
    if err != nil {
        //TODO
        logger.Log.Error().Str("reason", err.Error()).Msg("can`t write ts header")
    }

    return nil
}

////////////////////////////////////////////////////////////////////////////////

func (r *Record) Write(pts uint64, b []byte) (int, error) {
    if r.FirstPts == 0 {
        r.FirstPts  = pts
    } else {
        r.LastPts   = pts
    }
    r.FrameCount++
    r.Chunks[0].Size = r.Chunks[0].Size + uint64(len(b))

    n, err := r.currentChunkFile.Write(b)
    //r.currentChunkFile.Sync()

/*
type Packet struct {
	IsKeyFrame      bool // video packet is key frame
	Idx             int8 // stream index in container format
	CompositionTime time.Duration // packet presentation time minus decode time for H264 B-Frame
	Time time.Duration // packet decode time
	Data            []byte // packet data
}
*/
    //logger.Log.Trace().
    //    Uint64("pts", pts).
    //    //Time("tpts", time.Duration(pts))
    //    Msg("record write")

    if r.ts != nil {
        var flag bool

        if b[4] == 103 {
            flag = true
            logger.Log.Trace().Msg("KeyFrame")
        }

        var packet av.Packet = av.Packet{
            IsKeyFrame: flag,
            Idx: 0,
            //CompositionTime: time.Duration(pts),
            Time: time.Duration(pts-r.FirstPts)*time.Microsecond,// / time.Duration(90000),
            //Time: time.Duration(r.tstime),
            Data: b,
        }
        //r.tstime = r.tstime + 3000

    //err = r.mp4.WritePacket(packet)
    //if err != nil {
    //    //TODO
    //}

        err = r.ts.WritePacket(packet)
        if err != nil {
            //TODO
            logger.Log.Error().Str("reason", err.Error()).Msg("ts write")
        }
    }

    return n, err
}

//func (r *Record) ReadFrom(pts uint64, b []byte) (int, error) {
//
//}
