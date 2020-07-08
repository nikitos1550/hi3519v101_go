package jpeg

import (
    "fmt"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"

    "application/pkg/openapi"

    "application/pkg/mpp/connection"
    "application/pkg/mpp/venc"
    jpegstreamer "application/pkg/streamer/jpeg"
)

func init() {
    openapi.AddApiRoute("encoderBindJpeg", "/mpp/encoders/{encoder:[0-9]+}/bind/jpeg/{jpeg:[0-9]+}", "GET", encoderBindJpeg)
    openapi.AddApiRoute("encoderUnbindJpeg", "/mpp/encoders/{encoder:[0-9]+}/unbind/jpeg/{jpeg:[0-9]+}", "GET", encoderUnbindJpeg)
}

func encoderBindJpeg(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

    queryParams := mux.Vars(r)

    encoderId, err := strconv.Atoi(queryParams["encoder"])
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    jpegId, err := strconv.Atoi(queryParams["jpeg"])
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

    j, err := jpegstreamer.GetById(jpegId)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    err = connection.ConnectEncodedData(e, j)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "{\"ok\":\"binded\"}")
}

func encoderUnbindJpeg(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

    queryParams := mux.Vars(r)

    encoderId, err := strconv.Atoi(queryParams["encoder"])
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    jpegId, err := strconv.Atoi(queryParams["jpeg"])
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

    j, err := jpegstreamer.GetById(jpegId)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    err = connection.DisconnectEncodedData(e, j)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "{\"ok\":\"unbinded\"}")
}
