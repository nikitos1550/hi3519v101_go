// +build openapi

package chip

import (
	"log"
	"net/http"
	"application/pkg/openapi"
)

func init() {
	openapi.AddRoute("serveInfo", "/api/chip", "GET", serveInfo)
}

func serveInfo(w http.ResponseWriter, r *http.Request) {
	log.Println("chip.serveInfo")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

}