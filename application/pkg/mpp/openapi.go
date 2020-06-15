//+build openapi

package mpp

import (
	"log"
	"fmt"
	"net/http"
	"application/pkg/openapi"
	"encoding/json"
    "application/pkg/mpp/utils"

    "application/pkg/logger"
)

func init() {
	openapi.AddApiRoute("serveVersion", "/mpp/version", "GET", serveVersion)
    openapi.AddApiRoute("serveVersion", "/mpp/syncpts", "GET", runSyncPts)
    openapi.AddApiRoute("serveVersion", "/mpp/initpts", "GET", runInitPts)
}

type serveVersionSchema struct {
	Version	string	`json:"version"`
}

func serveVersion(w http.ResponseWriter, r *http.Request) {
	log.Println("mpp.serveVersion")

	var schema serveVersionSchema
	schema.Version = utils.Version()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	schemaJson, _ := json.Marshal(schema)
    fmt.Fprintf(w, "%s", string(schemaJson))
}



func runSyncPts(w http.ResponseWriter, r *http.Request) {
    err := utils.SyncPTS(100000)

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    
    if err != nil {
        logger.Log.Warn().
            Msg("utils.SyncPTS")  
        w.WriteHeader(http.StatusInternalServerError)
    } else {
        w.WriteHeader(http.StatusOK)
    }
}

func runInitPts(w http.ResponseWriter, r *http.Request) {
    err := utils.InitPTS(0)

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

    if err != nil {
        logger.Log.Warn().
            Msg("utils.InitPTS")
        w.WriteHeader(http.StatusInternalServerError)
    } else {
        w.WriteHeader(http.StatusOK)
    }
}  
