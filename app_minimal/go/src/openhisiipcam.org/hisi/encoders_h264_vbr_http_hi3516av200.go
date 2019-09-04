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
    openapi.AddRoute("encoderH264VbrCreate",   "/api/encoders/h264/vbr",              "POST",     encoderH264VbrCreate)
    openapi.AddRoute("encoderH264VbrFetch",    "/api/encoders/h264/vbr/{id:[0-9]+}",  "GET",      encoderH264VbrFetch)
    openapi.AddRoute("encodesH264VbrUpdate",   "/api/encoders/h264/vbr/{id:[0-9]+}",  "PATCH",    encoderH264VbrUpdate)
}

func encoderH264VbrCreate(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
}

func encoderH264VbrFetch(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
}

func encoderH264VbrUpdate(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
}

