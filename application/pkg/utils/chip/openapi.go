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
	openapi.AddRoute("serveInfo", "/chip", "GET", serveInfo)
}

type serveInfoSchema struct {

}

func serveInfo(w http.ResponseWriter, r *http.Request) {
	log.Println("chip.serveInfo")

	var schema serveInfoSchema

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	schemaJson, _ := json.Marshal(schema)
    fmt.Fprintf(w, "%s", string(schemaJson))
}