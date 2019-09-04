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
    openapi.AddRoute("encoderH265VbrCreate",   "/api/encoders/h265/vbr",              "POST",     encoderH265VbrCreate)
    openapi.AddRoute("encoderH265VbrFetch",    "/api/encoders/h265/vbr/{id:[0-9]+}",  "GET",      encoderH265VbrFetch)
    openapi.AddRoute("encodesH265VbrUpdate",   "/api/encoders/h265/vbr/{id:[0-9]+}",  "PATCH",    encoderH265VbrUpdate)
}

func encoderH265VbrCreate(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
}

func encoderH265VbrFetch(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
}

func encoderH265VbrUpdate(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
}

