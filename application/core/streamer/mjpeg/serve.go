package mjpeg

import (
    "fmt"
    "strconv"
    "net/http"
    "io"

    "github.com/pkg/errors"
    "github.com/gorilla/mux"
    "github.com/valyala/bytebufferpool"

    "application/core/mpp/frames"
    "application/core/logger"
)

func serve(w http.ResponseWriter, r *http.Request, c *mjpegClient) {
    defer c.source.removeClient(c.name)

    var err error

	cn, ok := w.(http.CloseNotifier)
	if !ok {
		http.NotFound(w, r)
		return
	}

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.NotFound(w, r)
		return
	}

    w.Header().Set("Cache-Control", "no-cache")
    w.Header().Set("Pragma", "no-cache")
    w.Header().Set("Connection", "close")
	w.Header().Set("Content-Type", "multipart/x-mixed-replace; boundary=mjpegstream")
    w.Header().Set("Transfer-Encoding", "chunked")

	w.WriteHeader(http.StatusOK)

	flusher.Flush()

    for {
        select {
        case frame := <-c.notify:
            buf := bytebufferpool.Get()

            c.source.writeFrameTo(frame, buf)

            w.Write([]byte("--mjpegstream\r\nContent-Type: image/jpeg\r\nContent-Length: "+strconv.Itoa(frame.Size)+"\r\n\r\n"))    //TODO check err

            _, err = w.Write(buf.B)
            if err != nil {
                    logger.Log.Warn().
                        Str("client", c.name).
                        Str("reason", err.Error()).
                        Str("mjpeg_name", c.source.name).
                        Msg("Mjpeg client write")
            }
            flusher.Flush()
            bytebufferpool.Put(buf)
        case <- cn.CloseNotify():
            logger.Log.Debug().
                Str("name", c.name).
                Str("mjpeg_name", c.source.name).
                Msg("Client stopped listening")
            return
        case <- c.stop:
            logger.Log.Debug().
                Str("name", c.name).
                Str("mjpeg_name", c.source.name).
                Msg("Client received stop")
            return
        }
    }
}

////////////////////////////////////////////////////////////////////////////////

func (g *MjpegGroup) ServeStreamGroup(w http.ResponseWriter, r *http.Request) {
    queryParams := mux.Vars(r)

    g.RLock()

    m, err := g.Get(queryParams["name"])
    if err != nil {
        g.RUnlock()
        w.WriteHeader(http.StatusNotFound)
        return
    }

    client, err := m.newClient(r.RemoteAddr)
    g.RUnlock()

    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    serve(w, r, client)
}

////////////////////////////////////////////////////////////////////////////////

func (m *Mjpeg) ServeStream(w http.ResponseWriter, r *http.Request) {
    client, err := m.newClient(r.RemoteAddr)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    serve(w, r, client)
}

////////////////////////////////////////////////////////////////////////////////

func (m *Mjpeg) writeFrameTo(f frames.FrameItem, buf io.Writer) error {
    m.RLock()
    defer m.RUnlock()

    if m.source == nil {
        return errors.New("Instance not sourced")
    }

    s, err := m.source.GetStorage()
    if err != nil {
        return errors.Wrap(err, "writeFrameTo failed")
    }

    _, err = s.WriteItemTo(f, buf)
    if err != nil {
        return errors.Wrap(err, "writeFrameTo failed")
    }

    return nil
}
