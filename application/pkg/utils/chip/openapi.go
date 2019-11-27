// +build openapi

package chip

import (
	"log"
	"fmt"
	"net/http"
	"application/pkg/openapi"
	"encoding/json"

)

func init() {
	openapi.AddApiRoute("serveInfo", "/chip", "GET", serveInfo)
}

type serveInfoSchema struct {
	RegId		uint32	`json:"regchipid"`
	DetectReg	string	`json:"detectedchipreg"`
	MppId		uint32	`json:"mppchipid"`
	DetectMpp	string	`json:"detectedchipmpp"`
}

func serveInfo(w http.ResponseWriter, r *http.Request) {
	log.Println("chip.serveInfo")

	var schema serveInfoSchema

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	schema.RegId = RegId()
	schema.DetectReg = Detect(RegId())
	schema.MppId = MppId()
	schema.DetectMpp = Detect(MppId())

	schemaJson, _ := json.Marshal(schema)
    fmt.Fprintf(w, "%s", string(schemaJson))
}