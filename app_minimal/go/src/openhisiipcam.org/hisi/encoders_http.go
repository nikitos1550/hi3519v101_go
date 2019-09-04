// +build openapi

package hisi

//#include "../../../../libhisi/hisi_external.h"
//#cgo LDFLAGS: ${SRCDIR}/../../../../libhisi/libhisi.a
import "C"

import (
    "net/http"
    "strconv"
    "encoding/json"
    "fmt"
    "log"

    "openhisiipcam.org/openapi"

    "github.com/gorilla/mux"
)

func init() {
    openapi.AddRoute("encodersList",   "/api/encoders",                "GET",      encodersList)
    openapi.AddRoute("encoderFetch",   "/api/encoders/{id:[0-9]+}",    "GET",      encoderFetch)
    openapi.AddRoute("encoderDelete",  "/api/encoders/{id:[0-9]+}",    "DELETE",   encoderDelete)
}

type EncodersListSchema struct {
    Encs    []EncoderItemSchema `json:",omitempty"`
}

type EncoderItemSchema struct {
    Id      int     `json:"id"`
    Codec   string  `json:"codec"`
    Rc      string  `json:"rc"`
    Profile string  `json:"profile"`
    Width   int     `json:"width"`
    Height  int     `json:"height"`
    Fps     int     `json:"fps"`
    Channel int     `json:"channel"`

}

func encodersList(w http.ResponseWriter, r *http.Request) {
    log.Println("encodersList")

    var encsSchema EncodersListSchema
    encsSchema.Encs = make([]EncoderItemSchema, 0)

    var num C.uint
    C.hisi_encoders_max_num(&num)

    var sparams C.struct_encoder_static_params

    for i := 0;  i < int(num) ; i++ {
        var encSchema EncoderItemSchema

        encSchema.Id = i

        error_code := C.hisi_encoder_fetch(C.uint(i), &sparams);

        switch error_code {
        case C.ERR_NONE:
            encSchema.Codec = sparamCodecResolve(sparams.codec)
            encSchema.Rc = sparamRcResolve(sparams.rc)
            encSchema.Profile = sparamProfileResolve(sparams.profile)

            encSchema.Width     = int(sparams.width)
            encSchema.Height    = int(sparams.height)
            encSchema.Fps       = int(sparams.fps)
            encSchema.Channel   = int(sparams.channel)

            encsSchema.Encs = append(encsSchema.Encs, encSchema)
        case C.ERR_OBJECT_NOT_FOUND:
        default:
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    encsSchemaJson, _ := json.Marshal(encsSchema)
    fmt.Fprintf(w, "%s", string(encsSchemaJson))
}

func encoderFetch(w http.ResponseWriter, r *http.Request) {
    log.Println("encoderFetch")

    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    var sparams C.struct_encoder_static_params

    var encSchema EncoderItemSchema

    encSchema.Id = id

    error_code := C.hisi_encoder_fetch(C.uint(id), &sparams)

    switch error_code {
        case C.ERR_NONE:
        case C.ERR_OBJECT_NOT_FOUND:
            w.WriteHeader(http.StatusNotFound)
            return
        default:
            w.WriteHeader(http.StatusInternalServerError)
            return
    }

    encSchema.Codec = sparamCodecResolve(sparams.codec)
    encSchema.Rc = sparamRcResolve(sparams.rc)
    encSchema.Profile = sparamProfileResolve(sparams.profile)

    encSchema.Width     = int(sparams.width)
    encSchema.Height    = int(sparams.height)
    encSchema.Fps       = int(sparams.fps)
    encSchema.Channel   = int(sparams.channel)

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    encSchemaJson, _ := json.Marshal(encSchema)
    fmt.Fprintf(w, "%s", string(encSchemaJson))
}

func encoderDelete(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
}

func sparamCodecResolve (codec C.int) string {
    switch  codec {
        case C.CODEC_JPEG:
            return "jpeg"
        case C.CODEC_MJPEG:
            return "mjpeg"
        case C.CODEC_H264:
            return "h264"
        case C.CODEC_H265:
            return "h265"
        default:
            return "unknown"
    }
}

func sparamRcResolve (rc C.int) string {
    switch rc {
        case C.RC_CBR:
            return "cbr"
        case C.RC_VBR:
            return "vbr"
        case C.RC_FIXQP:
            return "fixqp"
        case C.RC_AVBR:
            return "avbr"
        case C.RC_QVBR:
            return "qvbr"
        case C.RC_QMAP:
            return "qmap"
        default:
            return "unknown"
    }
}

func sparamProfileResolve (profile C.int) string {
    switch profile {
        case C.PROFILE_JPEG_BASELINE:
            return "baseline"
        case C.PROFILE_MJPEG_BASELINE:
            return "baseline"
        case C.PROFILE_H264_BASELINE:
            return "baseline"
        case C.PROFILE_H264_MAIN:
            return "main"
        case C.PROFILE_H264_HIGH:
            return "high"
        case C.PROFILE_H265_MAIN:
            return "main"
        default:
            return "unknown"
    }

}
