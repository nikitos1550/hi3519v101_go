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
	openapi.AddRoute("channelsList",   "/api/channels",                "GET",      channelsList)
    openapi.AddRoute("channelFetch",   "/api/channels/{id:[0-9]+}",    "GET",      channelFetch)
    openapi.AddRoute("channelEnable",  "/api/channels/{id:[0-9]+}",    "PUT",      channelEnable)
    openapi.AddRoute("channelDisable", "/api/channels/{id:[0-9]+}",    "DELETE",   channelDisable)
}

type ChannelsListSchema struct {
    Chns    []ChannelItemSchema
}

type ChannelItemSchema struct {
    Id      int  `json:"id"`
    Enabled int  `json:"enabled"`
    Width   int  `json:"width,omitempty"`
    Height  int  `json:"height,omitempty"`
    Fps     int  `json:"fps,omitempty"`
}

func channelsList(w http.ResponseWriter, r *http.Request) {

    log.Println("channelsList")

    var chnsSchema ChannelsListSchema

    var num C.uint
    C.hisi_channels_max_num(&num)

    var chn C.struct_channel_params

    for i := 0;  i < int(num) ; i++ {
        var chnSchema ChannelItemSchema

        chnSchema.Id = i

        error_code := C.hisi_channel_fetch(C.uint(i), &chn);

        switch error_code {
        case C.ERR_NONE:
            chnSchema.Enabled   = 1
            chnSchema.Width     = int(chn.width)
            chnSchema.Height    = int(chn.height)
            chnSchema.Fps       = int(chn.fps)
        case C.ERR_DISABLED:
            chnSchema.Enabled   = 0
        default:
            w.WriteHeader(http.StatusInternalServerError)
            return
        }

        chnsSchema.Chns = append(chnsSchema.Chns, chnSchema)
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    chnsSchemaJson, _ := json.Marshal(chnsSchema)
    fmt.Fprintf(w, "%s", string(chnsSchemaJson))
}

func channelFetch(w http.ResponseWriter, r *http.Request) {
    log.Println("channelFetch")

    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    var chn C.struct_channel_params

    var chnSchema ChannelItemSchema

    chnSchema.Id = id

    error_code := C.hisi_channel_fetch(C.uint(id), &chn);

    switch error_code {
        case C.ERR_NONE:
            chnSchema.Enabled   = 1
            chnSchema.Width     = int(chn.width)
            chnSchema.Height    = int(chn.height)
            chnSchema.Fps       = int(chn.fps)
        case C.ERR_DISABLED:
            chnSchema.Enabled   = 0
        case C.ERR_OBJECT_NOT_FOUND:
            w.WriteHeader(http.StatusNotFound)
            return
        default:
            w.WriteHeader(http.StatusInternalServerError)
            return
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    chnSchemaJson, _ := json.Marshal(chnSchema)
    fmt.Fprintf(w, "%s", string(chnSchemaJson))
}

func channelEnable(w http.ResponseWriter, r *http.Request) {

    log.Println("channelEnable")

    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])

    var chnSchema ChannelItemSchema

    //log.Println(r.Body) //todo
    jerr := json.NewDecoder(r.Body).Decode(&chnSchema)
    if jerr != nil {
        log.Println(jerr)
        w.WriteHeader(http.StatusInternalServerError)   //TODO
        return
    }

    var chn C.struct_channel_params

    chn.width   = C.int(chnSchema.Width)
    chn.height  = C.int(chnSchema.Height)
    chn.fps     = C.int(chnSchema.Fps)

    error_code := C.hisi_channel_enable (C.uint(id), &chn);
    switch error_code {
        case C.ERR_NONE:
            w.WriteHeader(http.StatusNoContent) //(http.StatusOK)
        case C.ERR_OBJECT_NOT_FOUND:
            w.WriteHeader(http.StatusNotFound)
        case C.ERR_NOT_ALLOWED:
            w.WriteHeader(http.StatusMethodNotAllowed)
        case C.ERR_BAD_PARAMS:
            w.WriteHeader(http.StatusBadRequest)
        default:
            w.WriteHeader(http.StatusInternalServerError)
    }
}

func channelDisable(w http.ResponseWriter, r *http.Request) {

    log.Println("channelDisable")

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
}
