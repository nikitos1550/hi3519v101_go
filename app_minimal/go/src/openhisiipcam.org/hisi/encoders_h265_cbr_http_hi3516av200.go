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
    openapi.AddRoute("encoderH265CbrCreate",   "/api/encoders/h265/cbr",              "POST",     encoderH265CbrCreate)
    openapi.AddRoute("encoderH265CbrFetch",    "/api/encoders/h265/cbr/{id:[0-9]+}",  "GET",      encoderH265CbrFetch)
    openapi.AddRoute("encodesH265CbrUpdate",   "/api/encoders/h265/cbr/{id:[0-9]+}",  "PATCH",    encoderH265CbrUpdate)
}

func encoderH265CbrCreate(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
}

func encoderH265CbrFetch(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
}

func encoderH265CbrUpdate(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
}

