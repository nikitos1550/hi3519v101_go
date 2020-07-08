// +build scripts
// +build scriptsIde

// HTTP API exposes ability to create/edit/delete/debug scripts in runtime

package scripts

import (
    "net/http"
    "application/pkg/openapi"
)

func init() {
    openapi.AddApiRoute("exportedFuncs",      "/scripts/exportedFuncs",     "GET",      exportedFuncs)
}

func systemDate(w http.ResponseWriter, r *http.Request) {
    log.Println("exportedFuncs")

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusNotImplemented)
}
