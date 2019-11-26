//+build openapi

package mpp

import (
	"log"
	"net/http"
	"application/pkg/openapi"
)

func init() {
	openapi.AddRoute("serveVersion", "/api/mpp/version", "GET", serveVersion)
}

func serveVersion(w http.ResponseWriter, r *http.Request) {
	log.Println("mpp.serveVersion")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

}
