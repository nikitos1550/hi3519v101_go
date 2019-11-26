//+build openapi

package buildinfo

import (
	"log"
    "net/http"
    "application/pkg/openapi"
)

func init() {
	AddRoute("serveInfo", "/api/buildinfo", "GET", serveInfo)
}

func serveInfo(w http.ResponseWriter, r *http.Request) {
	log.Println("buildinfo.serveInfo")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

}
