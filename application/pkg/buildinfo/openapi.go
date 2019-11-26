//+build openapi

package buildinfo

import (
	"log"
	"fmt"
    "net/http"
	"application/pkg/openapi"
	"encoding/json"
)

func init() {
	openapi.AddRoute("serveInfo", "/buildinfo", "GET", serveInfo)
}

func serveInfo(w http.ResponseWriter, r *http.Request) {
	log.Println("buildinfo.serveInfo")

	var schema Info
	CopyAll(&schema)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	schemaJson, _ := json.Marshal(schema)
    fmt.Fprintf(w, "%s", string(schemaJson))
}
