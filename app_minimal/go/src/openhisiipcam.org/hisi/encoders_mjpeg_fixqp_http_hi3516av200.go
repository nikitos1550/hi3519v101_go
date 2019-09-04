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
    openapi.AddRoute("encoderMjpegFixqpCreate",   "/api/encoders/mjpeg/fixqp",              "POST",     encoderMjpegFixqpCreate)
    openapi.AddRoute("encoderMjpegFixqpFetch",    "/api/encoders/mjpeg/fixqp/{id:[0-9]+}",  "GET",      encoderMjpegFixqpFetch)
    openapi.AddRoute("encodesMjpegFixqpUpdate",   "/api/encoders/mjpeg/fixqp/{id:[0-9]+}",  "PATCH",    encoderMjpegFixqpUpdate)
}

func encoderMjpegFixqpCreate(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
}

func encoderMjpegFixqpFetch(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
}

func encoderMjpegFixqpUpdate(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
}

