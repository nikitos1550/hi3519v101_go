package encoder

import (
    "fmt"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"

    "application/pkg/openapi"

    "application/pkg/mpp/connection"
    "application/pkg/mpp/vpss"
    "application/pkg/mpp/venc"
)

func init() {
    openapi.AddApiRoute("encoderBindVpss", "/mpp/channels/{channel:[0-9]+}/bind/encoder/{encoder:[0-9]+}", "GET", encoderBindVpss)
    openapi.AddApiRoute("encoderUnbindVpss", "/mpp/channels/{channel:[0-9]+}/unbind/encoder/{encoder:[0-9]+}", "GET", encoderUnbindVpss)
}

func encoderBindVpss(w http.ResponseWriter, r *http.Request)  {
    w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

    queryParams := mux.Vars(r)

    encoderId, err := strconv.Atoi(queryParams["encoder"])
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    channelId, err := strconv.Atoi(queryParams["channel"])
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

    c, err := vpss.GetChannel(channelId)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    err = connection.ConnectBind(c, e)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "{\"ok\":\"binded\"}")
}

func encoderUnbindVpss(w http.ResponseWriter, r *http.Request)  {
    w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

    queryParams := mux.Vars(r)

    encoderId, err := strconv.Atoi(queryParams["encoder"])
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    channelId, err := strconv.Atoi(queryParams["channel"])
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

    c, err := vpss.GetChannel(channelId)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    err = connection.DisconnectBind(c, e)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "{\"ok\":\"unbinded\"}")
}
