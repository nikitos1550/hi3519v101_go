package record

import (
    "os"
    "encoding/json"
    "io/ioutil"
    "time"

    "github.com/pkg/errors"

    "application/core/logger"
)

type Record struct {
    Dir                 string              `json:"-"`
    Name                string              `json:"name"`

    State		        RecordState         `json:"-"`

    FirstPts            uint64              `json:"first_pts"`
    LastPts             uint64              `json:"last_pts"`
    FrameCount          uint64              `json:"frames"`

    currentChunk        int
    currentChunkFile    *os.File
    Chunks              []Chunk             `json:"chunks"`

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

    //frames      []Frame

    MD5			string              `json:"md5"`
}

type Frame struct {
    offset  uint64
}

func New(dir string, name string) (*Record, error) {

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

    rec.currentChunkFile, err = os.Create(dir+"/"+name+"/1.h264")
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

    return nil
}

func (r *Record) SetPreview(jpeg []byte) error {
    f, err := os.Create(r.Dir+"/"+r.Name+"/preview.jpeg")
    if err != nil {
        return errors.New("Can`t create preview.jpeg")
    }

    f.Write(jpeg)
    f.Sync()
    f.Close()

    r.Preview = true

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

    n, err := r.currentChunkFile.Write(b)
    r.currentChunkFile.Sync()

    return n, err
}

//func (r *Record) ReadFrom(pts uint64, b []byte) (int, error) {
//
//}
