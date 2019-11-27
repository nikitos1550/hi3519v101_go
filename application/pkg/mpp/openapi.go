//+build openapi

package mpp

import (
	"log"
	"fmt"
	"net/http"
	"application/pkg/openapi"
	"encoding/json"
)

func init() {
	openapi.AddApiRoute("serveVersion", "/mpp/version", "GET", serveVersion)
}

type serveVersionSchema struct {
	Version	string	`json:"version"`
}

func serveVersion(w http.ResponseWriter, r *http.Request) {
	log.Println("mpp.serveVersion")

	var schema serveVersionSchema
	schema.Version = Version()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	schemaJson, _ := json.Marshal(schema)
    fmt.Fprintf(w, "%s", string(schemaJson))
}
