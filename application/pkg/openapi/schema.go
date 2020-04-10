//+build openapi

////go:generate go-bindata -o openapiyaml.go --pkg openapi ./openapi.yaml
////go:generate go run -tags 

package openapi

import (
    _"encoding/json"
    _"fmt"
    //"log"
    "net/http"
    "application/pkg/logger"
)

func init() {
    AddApiRoute("serveOpenapi",      "/openapi.json",     "GET",      serveOpenapi)
}

func serveOpenapi(w http.ResponseWriter, r *http.Request) {
    //log.Println("serveOpenapi")
	logger.Log.Trace().Msg("serveOpenapi")

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

}

