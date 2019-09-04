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
    openapi.AddRoute("encoderH264CbrCreate",   "/api/encoders/h264/cbr",              "POST",     encoderH264CbrCreate)
    openapi.AddRoute("encoderH264CbrFetch",    "/api/encoders/h264/cbr/{id:[0-9]+}",  "GET",      encoderH264CbrFetch)
    openapi.AddRoute("encodesH264CbrUpdate",   "/api/encoders/h264/cbr/{id:[0-9]+}",  "PATCH",    encoderH264CbrUpdate)
}

func encoderH264CbrCreate(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
}

func encoderH264CbrFetch(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
}

func encoderH264CbrUpdate(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
}

