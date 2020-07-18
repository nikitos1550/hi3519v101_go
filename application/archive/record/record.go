package record

import (
    "os"
    "encoding/json"

    "github.com/pkg/errors"
)

type Record struct {
    Dir         string              `json:"-"`
    Name        string              `json:"name"`

    State		RecordState         `json:"-"`

    FirstPts    uint64              `json:"first_pts"`
    LastPts     uint64              `json:"last_pts"`
    FrameCount  uint64              `json:"frames"`

    Chunks      []Chunk             `json:"chunks"`

    Preview     bool                `json:"preview"`
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

    MD5			string              `json:"md5"`
}

func New(dir string, name string) (*Record, error) {
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

    var record Record = Record{
        Dir: dir,
        Name: name,
    }

    json, _ := json.Marshal(record)
    f.Write(json)
    f.Close()

    return &record, nil
}

func Load(path string) (*Record, error) {
    stat, err := os.Stat(path)
    if err != nil {
        return nil, errors.Wrap(err, "No such dir")
    }

	if stat.IsDir() == false {
		return nil, errors.New("Dir is not dir")
    }

    stat, err = os.Stat(path+"/info.json")
    if err != nil {
		return nil, errors.Wrap(err, "No info")
    }

	f, err := os.Open(path+"/info.json")
	if err != nil {
		return nil, errors.New("Can`t open")
	}

    var record Record

	//TODO read info to struct

	f.Close()

    return &record, nil
}

////////////////////////////////////////////////////////////////////////////////

//func (r *Record) Open() error {
//    return nil
//}

func (r *Record) Close() error {
    return nil
}

func (r *Record) Write(b []byte) (int, error) {
    return 0, nil
}

func (r *Record) SetPreview(jpeg []byte) error {
    return nil
}
////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////
