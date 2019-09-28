// +build openapi

package hisi

//#include "../../../../libhisi/hisi_external.h"
//#cgo LDFLAGS: ${SRCDIR}/../../../../libhisi/libhisi.a
import "C"

import (
    "net/http"
    "fmt"
    "log"
    "encoding/json"
    "openhisiipcam.org/openapi"
)

func init() {
    openapi.AddRoute("systemTemperature",   "/api/system/temperature",  "GET", systemTemperature)
    openapi.AddRoute("systemChipId",        "/api/system/chipid",       "GET", systemChipId)
}

type systemTemperatureSchema struct {
    TemperatureC float32 `json:"temperature_c"`
    TemperatureF float32 `json:"temperature_f"`
    TemperatureK float32 `json:"temperature_k"`
}

func systemTemperature(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    var tSchema systemTemperatureSchema

    var tmp C.float
    C.hisi_get_temperature(&tmp)

    var t float32 = float32(tmp)

    tSchema.TemperatureC = float32(int(t*10)) / 10
    tSchema.TemperatureF = float32(int((t*9/5)+32)*10) / 10
    tSchema.TemperatureK = float32(int((t+273.15)*10)) / 10

    log.Println("systemTemperature")
    tSchemaJson, _ := json.Marshal(tSchema)
    fmt.Fprintf(w, "%s", string(tSchemaJson))
}

func systemChipId(w http.ResponseWriter, r *http.Request) {
    log.Println("systemChipId")

    var chip C.uint
    C.hisi_get_chipid(&chip)

    w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    fmt.Fprintf(w, "ChipID %d\n",    uint(chip))
}
