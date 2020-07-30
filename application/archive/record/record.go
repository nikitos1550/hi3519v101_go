package record

import (
    "os"
    "encoding/json"
    "io/ioutil"
    "time"
    "strconv"

    "github.com/pkg/errors"

    //"application/archive/mp4"
    "application/archive/ts"

    //"github.com/nareix/joy4/format/mp4"
    "github.com/nareix/joy4/av"
    "github.com/nareix/joy4/codec/h264parser"

    "application/core/logger"
)

const (
    chunkFrameLength = 30*60 //30fps*60seconds = 1 minute
)

type Record struct {
    Dir                 string              `json:"-"`
    Name                string              `json:"name"`

    Codec               string              `json:"codec"`

    FirstPts            uint64              `json:"first_pts"`
    LastPts             uint64              `json:"last_pts"`
    FrameCount          uint64              `json:"frames"`
    Fps                 int

    currentChunk        int                 `json:"-"`
    currentChunkFile    *os.File            `josn:"-"`
    Chunks              []Chunk             `json:"chunks"`

    sps                 []byte              `json:"-"`
    pps                 []byte              `json:"-"`

    ts                  *ts.Muxer           `json:"-"`
    //tstest              *os.File
    //tstime              uint64

    Preview             bool                `json:"preview"`
}

//type RecordState int
//const (
//	Created		RecordState	= 1
//	Finished	RecordState = 2
//)

type Chunk struct {
    Id          int                 `json:"id"`

    FirstPts    uint64              `json:"first_pts"`
    LastPts     uint64              `json:"last_pts"`
    FrameCount  uint64              `json:"frames"`
    Size        uint64              `json:"size"`

    MD5			string              `json:"md5"`
}

type Frame struct {
    offset  uint64
}

func New(dir string, name string, codec string) (*Record, error) {

    if codec != "h264" {
        return nil, errors.New("only h264 supoprted now")
    }

    stat, err := os.Stat(dir)
    if err != nil {
        return nil, errors.Wrap(err, "New failed")
    }

    if stat.IsDir() == false {
        return nil, errors.New("Dir is not dir")
    }

    stat, err = os.Stat(dir+"/"+name)
    if err == nil {
        return nil, errors.New("Name is already exist")
    }

    err = os.Mkdir(dir+"/"+name, 0755)
    if err != nil {
        return nil, errors.New("Can`t create dir")
    }

    f, err := os.Create(dir+"/"+name+"/info.json")
    if err != nil {
        return nil, errors.New("Can`t create info")
    }

    var rec Record = Record{
        Dir: dir,
        Name: name,
        currentChunk: -1,
    }

    rec.Chunks = make([]Chunk, 0)
    //rec.Chunks[0] = Chunk{Id:0,}
    //logger.Log.Debug().
    //    Int("length", len(rec.Chunks)).Msg("Chunks length")


    json, _ := json.Marshal(rec)
    f.Write(json)
    f.Close()

    rec.Codec = codec

    //rec.currentChunkFile, err = os.Create(dir+"/"+name+"/1.ts")
    //if err != nil {
    //    logger.Log.Fatal().Str("reson", err.Error()).Msg("Can`t create chunk file")
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

    return &rec, nil
}

////////////////////////////////////////////////////////////////////////////////

func (r *Record) Close() error {

    r.currentChunkFile.Sync()
    r.currentChunkFile.Close()

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

    if r.ts != nil {
        r.closeChunk()
        /*
        err = r.ts.WriteTrailer()
        if err != nil {
            //TODO
            logger.Log.Warn().
                Str("reason", err.Error()).
                Msg("close mp4")
        }

        r.tstest.Sync()
        r.tstest.Close()
        */
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

/*
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
*/

////////////////////////////////////////////////////////////////////////////////
func (r *Record) SetSPSPPS(sps, pps []byte) error {
    if len(sps) == 0 {
        return errors.New("SPS is null")
    }

    if len(pps) == 0 {
        return errors.New("PPS is null")
    }

    r.sps = sps
    r.pps = pps

    return nil
}

////////////////////////////////////////////////////////////////////////////////

func (r *Record) newChunk() error {

    if r.ts != nil {
        return errors.New("active chunk not closed")
    }

    if len(r.sps) == 0 || len(r.pps) == 0 {
        return errors.New("sps pps not setuped")
    }

    r.currentChunk++

    r.Chunks = append(r.Chunks, Chunk{Id:r.currentChunk,})
    logger.Log.Debug().
        Int("length", len(r.Chunks)).Msg("New chunk")


    var err error

    r.currentChunkFile, err = os.Create(r.Dir+"/"+r.Name+"/"+strconv.Itoa(r.currentChunk)+".ts")
    if err != nil {
        return errors.New("Can`t create chunk file")
    }

    r.ts  = ts.NewMuxer(r.currentChunkFile)

    tmp, err := h264parser.NewCodecDataFromSPSAndPPS(r.sps, r.pps)
    if err != nil {
        return errors.New("Can`t init ts muxer")
    }

    var streams []av.CodecData
    streams = make([]av.CodecData, 1)
    streams[0] = tmp

    err = r.ts.WriteHeader(streams)
    if err != nil {
        return errors.New("can`t write ts header")
    }

    return nil
}

func (r *Record) closeChunk() error {

    if r.ts != nil {
        err := r.ts.WriteTrailer()
        if err != nil {
            return errors.New("can`t write trailer")
        }

        r.currentChunkFile.Sync()
        r.currentChunkFile.Close()

        r.currentChunkFile = nil
        r.ts = nil

    } else {
        errors.New("no active chunk")
    }

    return nil
}

////////////////////////////////////////////////////////////////////////////////

func (r *Record) Write(pts uint64, b []byte) (int, error) {
    if r.ts == nil {
        r.newChunk()	//TODO create chunk
    }

    if r.Chunks[r.currentChunk].FrameCount == chunkFrameLength {
        r.closeChunk()	//TODO close chunk
		r.newChunk()	//TODO new chunk
    }

    if r.FirstPts == 0 {
        r.FirstPts  = pts
    } else {
        r.LastPts   = pts
    }

	r.FrameCount++

    r.Chunks[r.currentChunk].Size += uint64(len(b))
	r.Chunks[r.currentChunk].FrameCount++

	if r.Chunks[r.currentChunk].FirstPts == 0 {
		r.Chunks[r.currentChunk].FirstPts = pts
	} else {
		r.Chunks[r.currentChunk].LastPts = pts
	}

	//////////////////////////////////////////////////

    //n, err := r.currentChunkFile.Write(b)
    //r.currentChunkFile.Sync()

    //logger.Log.Trace().
    //    Uint64("pts", pts).
    //    //Time("tpts", time.Duration(pts))
    //    Msg("record write")

	var flag bool
    if b[4] == 103 {	//TODO forward key frame flag from storage
		flag = true
        //logger.Log.Trace().Msg("KeyFrame")
	}

	var packet av.Packet = av.Packet{
		IsKeyFrame: flag,
        Idx: 0,
        //CompositionTime: time.Duration(pts),
        Time: time.Duration(pts-r.FirstPts)*time.Microsecond,// / time.Duration(90000),
        //Time: time.Duration(r.tstime),
        Data: b,
	}

    err := r.ts.WritePacket(packet)
		if err != nil {
		//logger.Log.Error().Str("reason", err.Error()).Msg("ts write")
		return 0, errors.Wrap(err, "TS packet")
	}

    return 0, err
}
