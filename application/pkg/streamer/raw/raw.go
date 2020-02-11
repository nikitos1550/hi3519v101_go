package raw

import (
    //"fmt"
    "log"
    "net/http"
    "application/pkg/openapi"

    //"application/pkg/mpp/getloop"
)

func init() {
    openapi.AddRoute("serveRawYuv",   "/raw/1.yuv",   "GET",      serveRawYuv)
    openapi.AddRoute("serveRawBmp",   "/raw/1.bmp",   "GET",      serveRawBmp)
}


func Init() {}

func serveRawYuv(w http.ResponseWriter, r *http.Request) {
    log.Println("serveRawYuv")

    w.Header().Set("Content-Type", "???")
}

func serveRawBmp(w http.ResponseWriter, r *http.Request) {
    log.Println("serveRawBmp")

    w.Header().Set("Content-Type", "???")
}



