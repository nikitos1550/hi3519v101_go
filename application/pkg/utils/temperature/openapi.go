// +build openapi

package temperature

import (
	"log"
	"fmt"
	"net/http"
	"application/pkg/openapi"
	"encoding/json"
)

func init() {
	openapi.AddRoute("serveTemperature", "/temperature", "GET", serveTemperature)
}

type serveTemperatureSchema struct {

}

func serveTemperature(w http.ResponseWriter, r *http.Request) {
	log.Println("temperature.serveTemperature")

	var schema serveTemperatureSchema

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	schemaJson, _ := json.Marshal(schema)
    fmt.Fprintf(w, "%s", string(schemaJson))
}
