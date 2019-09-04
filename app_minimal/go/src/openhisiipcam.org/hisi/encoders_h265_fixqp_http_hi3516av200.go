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
    openapi.AddRoute("encoderH265FixqpCreate",   "/api/encoders/h265/fixqp",              "POST",     encoderH265FixqpCreate)
    openapi.AddRoute("encoderH265FixqpFetch",    "/api/encoders/h265/fixqp/{id:[0-9]+}",  "GET",      encoderH265FixqpFetch)
    openapi.AddRoute("encodesH265FixqpUpdate",   "/api/encoders/h265/fixqp/{id:[0-9]+}",  "PATCH",    encoderH265FixqpUpdate)
}

func encoderH265FixqpCreate(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
}

func encoderH265FixqpFetch(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
}

func encoderH265FixqpUpdate(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
}

