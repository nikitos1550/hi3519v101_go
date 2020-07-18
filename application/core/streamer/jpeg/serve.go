package jpeg

import (
    "fmt"
    "strconv"
    "net/http"
    "io"

    "github.com/pkg/errors"
    "github.com/gorilla/mux"
    "github.com/valyala/bytebufferpool"
)

func (g *JpegGroup) ServeFrameGroup(w http.ResponseWriter, r *http.Request) {
    queryParams := mux.Vars(r)

    g.RLock()   //we can`t defer as serve can be long

    j, err := g.Get(queryParams["name"])

    if err != nil {
        g.RUnlock()
        w.WriteHeader(http.StatusNotFound)
        return
    }

    frame := bytebufferpool.Get()
    err = j.getFrameCopy(frame)

    g.RUnlock()

    if err != nil {
        bytebufferpool.Put(frame)
		w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
	}

    serve(w, frame)
}

////////////////////////////////////////////////////////////////////////////////

//This is not compatible with Jpeg.Delete(), should be used in some special case builds
func (j *Jpeg) ServeFrame(w http.ResponseWriter, r *http.Request) {
    //j.RLock()
    frame := bytebufferpool.Get()
    err := j.getFrameCopy(frame)
    //j.RUnlock()

    if err != nil {
        bytebufferpool.Put(frame)
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    serve(w, frame)
}

////////////////////////////////////////////////////////////////////////////////

func serve(w http.ResponseWriter, frame *bytebufferpool.ByteBuffer) {
    w.Header().Set("Content-Type", "image/jpeg")
    w.Header().Set("Content-Length", strconv.Itoa(frame.Len()))

    w.Write(frame.B) //TODO check error
    bytebufferpool.Put(frame)
}

func (j *Jpeg) getFrameCopy(buf io.Writer) error {
    j.RLock()
    defer j.RUnlock()

    if j.source == nil {
        return errors.New("Instance not sourced")
    }

    s, err := j.source.GetStorage()
    if err != nil {
        return errors.Wrap(err, "GetFrameCopy failed")
    }

    _, err = s.WriteLastTo(buf)
    if err != nil {
        return errors.Wrap(err, "writeFrameTo failed")
    }

    return nil
}

func (j *Jpeg) GetJpeg() ([]byte, error) {
    j.RLock()
    defer j.RUnlock()

    if j.source == nil {
        return nil, errors.New("Instance not sourced")
    }

    s, err := j.source.GetStorage()
    if err != nil {
        return nil, errors.Wrap(err, "GetJpeg failed")
    }

    var frame []byte
    _, err = s.ReadLastAlloc(&frame)
    if err != nil {
        return nil, errors.Wrap(err, "GetJpeg failed")
    }

    return frame, nil
}
