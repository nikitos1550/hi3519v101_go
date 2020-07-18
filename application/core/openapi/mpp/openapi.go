package mpp

import (
	"fmt"
	"net/http"
	"encoding/json"

    "application/core/mpp/utils"
    "application/core/logger"
)

type versionSchema struct {
	Version	string	`json:"version"`
}

func Version(w http.ResponseWriter, r *http.Request) {
	var schema versionSchema
	schema.Version = utils.Version()

	w.WriteHeader(http.StatusOK)

	schemaJson, _ := json.Marshal(schema)
    fmt.Fprintf(w, "%s", string(schemaJson))
}



func RunSyncPts(w http.ResponseWriter, r *http.Request) {
    err := utils.SyncPTS(50000000000)

    if err != nil {
        logger.Log.Warn().
            Msg("utils.SyncPTS")
        w.WriteHeader(http.StatusInternalServerError)
    } else {
        w.WriteHeader(http.StatusOK)
    }
}

func RunInitPts(w http.ResponseWriter, r *http.Request) {
    err := utils.InitPTS(10000000000)

    if err != nil {
        logger.Log.Warn().
            Msg("utils.InitPTS")
        w.WriteHeader(http.StatusInternalServerError)
    } else {
        w.WriteHeader(http.StatusOK)
    }
}
