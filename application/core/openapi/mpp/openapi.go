package mpp

import (
	"fmt"
	"net/http"
	"encoding/json"
    "strconv"
    //"net/url"

    //"github.com/gorilla/mux"

    "application/core/mpp/vi"
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

func UpdateLDC(w http.ResponseWriter, r *http.Request) {
    //queryParams := mux.Vars(r)

    /*
    u, _ := url.Parse(r.URL)
    v := u.Query()

    if _, ok := v["x"]; !ok {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"no x\"}")
        return
    }
    if _, ok := v["y"]; !ok {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"no y\"}")
        return
    }
    if _, ok := v["k"]; !ok {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"no k\"}")
        return
    }
    */

    xI, err := strconv.Atoi(r.URL.Query().Get("x"))
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    yI, err := strconv.Atoi(r.URL.Query().Get("y"))
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    kI, err := strconv.Atoi(r.URL.Query().Get("k"))
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    err = vi.UpdateLDC(xI, yI, kI)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
    }

    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "{\"ok\":\"ldc updated\"}")
}
