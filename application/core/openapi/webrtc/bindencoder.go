package webrtc

import (
    "fmt"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"

    "application/core/mpp/connection"
    "application/core/mpp/venc"
    webrtcstreamer "application/core/streamer/webrtc"
)

func BindEncoder(w http.ResponseWriter, r *http.Request) {
    queryParams := mux.Vars(r)

    encoderId, err := strconv.Atoi(queryParams["encoder"])
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    webrtcId, err := strconv.Atoi(queryParams["webrtc"])
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    e, err := venc.GetEncoder(encoderId)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    webrtc, err := webrtcstreamer.GetById(webrtcId)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    err = connection.ConnectEncodedData(e, webrtc)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "{\"ok\":\"binded\"}")
}

func UnbindEncoder(w http.ResponseWriter, r *http.Request) {
    queryParams := mux.Vars(r)

    encoderId, err := strconv.Atoi(queryParams["encoder"])
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    webrtcId, err := strconv.Atoi(queryParams["webrtc"])
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    e, err := venc.GetEncoder(encoderId)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    webrtc, err := webrtcstreamer.GetById(webrtcId)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    err = connection.DisconnectEncodedData(e, webrtc)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "{\"ok\":\"unbinded\"}")
}

