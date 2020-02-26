//+build openapi

package mpp

import (
	"log"
	"fmt"
	"net/http"
	"application/pkg/openapi"
	"encoding/json"
    "application/pkg/mpp/utils"
)

func init() {
	openapi.AddApiRoute("serveVersion", "/mpp/version", "GET", serveVersion)
}

type serveVersionSchema struct {
	Version	string	`json:"version"`
}

/**
 * @api {get} /mpp/version Get MPP version
 * @apiName GetVersion
 * @apiGroup MPP
 */
func serveVersion(w http.ResponseWriter, r *http.Request) {
	log.Println("mpp.serveVersion")

	var schema serveVersionSchema
	schema.Version = utils.Version()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	schemaJson, _ := json.Marshal(schema)
    fmt.Fprintf(w, "%s", string(schemaJson))
}
