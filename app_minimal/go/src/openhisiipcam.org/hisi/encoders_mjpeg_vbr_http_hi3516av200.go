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
    openapi.AddRoute("encoderMjpegVbrCreate",   "/api/encoders/mjpeg/vbr",              "POST",     encoderMjpegVbrCreate)
    openapi.AddRoute("encoderMjpegVbrFetch",    "/api/encoders/mjpeg/vbr/{id:[0-9]+}",  "GET",      encoderMjpegVbrFetch)
    openapi.AddRoute("encodesMjpegVbrUpdate",   "/api/encoders/mjpeg/vbr/{id:[0-9]+}",  "PATCH",    encoderMjpegVbrUpdate)
}

func encoderMjpegVbrCreate(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
}

func encoderMjpegVbrFetch(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
}

func encoderMjpegVbrUpdate(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
}


