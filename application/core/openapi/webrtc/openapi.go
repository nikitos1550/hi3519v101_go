package webrtc

import (
    "fmt"
    "net/http"

    "application/core/streamer/webrtc"
)

func List(w http.ResponseWriter, r *http.Request) {
    //TODO
}

func Create(w http.ResponseWriter, r *http.Request) {
    var err error

    _, err = webrtc.Create()
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "{\"ok\":\"webrtc created\"}")
}

func Info(w http.ResponseWriter, r *http.Request) {
    //TODO
}

func Delete(w http.ResponseWriter, r *http.Request) {
    //TODO
}
