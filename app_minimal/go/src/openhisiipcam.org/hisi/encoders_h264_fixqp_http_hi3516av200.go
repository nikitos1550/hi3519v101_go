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
    openapi.AddRoute("encoderH264FixqpCreate",   "/api/encoders/h264/fixqp",              "POST",     encoderH264FixqpCreate)
    openapi.AddRoute("encoderH264FixqpFetch",    "/api/encoders/h264/fixqp/{id:[0-9]+}",  "GET",      encoderH264FixqpFetch)
    openapi.AddRoute("encodesH264FixqpUpdate",   "/api/encoders/h264/fixqp/{id:[0-9]+}",  "PATCH",    encoderH264FixqpUpdate)
}

func encoderH264FixqpCreate(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
}

func encoderH264FixqpFetch(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
}

func encoderH264FixqpUpdate(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
}

