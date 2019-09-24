package main

import (
    "fmt"
    "net/http"
    "encoding/json"
)

type Answer struct {
    app     string  `json:"appName"`
    chip    string  `json:"chipFamily"`
    mpp     string  `json:"mppVersion"`
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
    var schema Answer

    schema.app  = "app_tester"
    schema.chip = chipFamily
    schema.mpp  = version()

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    schemaJson, _ := json.Marshal(schema)
    fmt.Fprintf(w, "%s", string(schemaJson))
}

func main() {
    fmt.Println("app_tester, ", chipFamily, ", ", version())
    fmt.Println("starting http server :80")
    http.HandleFunc("/", versionHandler)
    http.ListenAndServe(":80", nil)
}
