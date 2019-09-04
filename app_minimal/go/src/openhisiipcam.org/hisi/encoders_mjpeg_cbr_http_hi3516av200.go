// +build openapi, hi3516av200

package hisi

//#include "../../../../libhisi/hisi_external.h"
//#cgo LDFLAGS: ${SRCDIR}/../../../../libhisi/libhisi.a
import "C"

import (
    "net/http"
    "openhisiipcam.org/openapi"
)

func init() {
    openapi.AddRoute("encoderMjpegCbrCreate",   "/api/encoders/mjpeg/cbr",              "POST",     encoderMjpegCbrCreate)
    openapi.AddRoute("encoderMjpegCbrFetch",    "/api/encoders/mjpeg/cbr/{id:[0-9]+}",  "GET",      encoderMjpegCbrFetch)
    openapi.AddRoute("encodesMjpegCbrUpdate",   "/api/encoders/mjpeg/cbr/{id:[0-9]+}",  "PATCH",    encoderMjpegCbrUpdate)
}

func encoderMjpegCbrCreate(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
}

func encoderMjpegCbrFetch(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
}

func encoderMjpegCbrUpdate(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
}

