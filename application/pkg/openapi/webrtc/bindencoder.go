package webrtc

import (
    "fmt"
    "net/http"
    //"strconv"

    //"github.com/gorilla/mux"

    "application/pkg/openapi"

    //"application/pkg/mpp/connection"
    //"application/pkg/mpp/venc"
    //"application/pkg/streamers/webrtc"
)

func init() {
    openapi.AddApiRoute("encoderBindWebrtc", "/mpp/encoders/{encoder:[0-9]+}/bind/webrtc/{webrtc:[0-9]+}", "GET", encoderBindWebrtc)
    openapi.AddApiRoute("encoderUnbindWebrtc", "/mpp/encoders/{encoder:[0-9]+}/unbind/webrtc/{webrtc:[0-9]+}", "GET", encoderUnbindWebrtc)
}

func encoderBindWebrtc(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintf(w, "{\"error\":\"Not implemented\"}")
    return
}

func encoderUnbindWebrtc(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintf(w, "{\"error\":\"Not implemented\"}")
    return
}

