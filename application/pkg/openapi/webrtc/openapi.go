package webrtc

import (
    //"fmt"
    "net/http"

    "application/pkg/openapi"
    //"application/pkg/streamers/webrtc"
)

func init() {
    openapi.AddApiRoute("webrtcsInfo",   "/webrtc",   "GET",      webrtcsInfo)
    openapi.AddApiRoute("webrtcCreate",   "/webrtc",   "POST",      webrtcCreate)

    openapi.AddApiRoute("webrtcInfo",   "/webrtc/{id:[0-9]+}",   "GET",     webrtcInfo)
    openapi.AddApiRoute("webrtcDelete",   "/webrtc/{id:[0-9]+}",   "DELETE",      webrtcDelete)

}

func webrtcsInfo(w http.ResponseWriter, r *http.Request) {}
func webrtcCreate(w http.ResponseWriter, r *http.Request) {}
func webrtcInfo(w http.ResponseWriter, r *http.Request) {}
func webrtcDelete(w http.ResponseWriter, r *http.Request) {}
