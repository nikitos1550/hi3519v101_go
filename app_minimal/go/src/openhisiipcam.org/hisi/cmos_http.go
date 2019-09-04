// +build openapi

package hisi

//#include "../../../../libhisi/hisi_external.h"
//#cgo LDFLAGS: ${SRCDIR}/../../../../libhisi/libhisi.a
import "C"

import (
    "net/http"
    "openhisiipcam.org/openapi"
)

func init() {
    openapi.AddRoute("cmosInfo", "/api/cmos", "GET", cmosInfo)
}

func cmosInfo(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
}

