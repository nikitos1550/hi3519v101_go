package openapi

import (
    "log"
    "../himpp3"
    "net/http"
    "strconv"
    "github.com/gorilla/mux"
    "fmt"
    "encoding/json"
)

func encodersServeJpeg (w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    log.Println(id)
    if id == 0 {

        himpp3.Mutex.Lock()
        log.Printf("Serving jpeg %d kb\n", len(himpp3.B1.Bytes())/1024)

        w.Header().Set("Content-Type", "image/jpeg")
        w.Header().Set("Content-Length", strconv.Itoa(len(himpp3.B1.Bytes())))

        w.WriteHeader(http.StatusOK)

        if _, err := w.Write(himpp3.B1.Bytes()); err != nil {
            log.Println("unable to write image.")
        }

        himpp3.Mutex.Unlock()
    } else {
        w.WriteHeader(http.StatusNotFound)

    }
}

type Encoder struct {
    Codec string `json:"codec"`
    Rc string `json:"rc"`
    Width int32 `json:"width"`
    Height int32 `json:"height"`
    Fps int32 `json:"fps"`
}


func encodersList (w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    var enc Encoder
    enc.Codec = "mjpeg"
    enc.Rc = "cbr"
    enc.Width = 1920
    enc.Height = 1080
    enc.Fps = 1

    log.Println("encodersList")
    test, _ := json.Marshal(enc)
    fmt.Fprintf(w, "%s", string(test))
}
