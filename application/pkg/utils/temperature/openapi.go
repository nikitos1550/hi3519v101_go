// +build openapi

package temperature

import (
	"log"
	"net/http"
	"application/pkg/openapi"
)

func init() {
	openapi.AddRoute("serveTemperature", "/api/temperature", "GET", serveTemperature)
}

func serveTemperature(w http.ResponseWriter, r *http.Request) {
	log.Println("temperature.serveTemperature")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

}
