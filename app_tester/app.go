package main

import (
    "fmt"
    "net/http"
    "encoding/json"
)

type Answer struct {
    App     string  `json:"appName"`
    Family  string  `json:"chipFamily"`
    Mpp     string  `json:"mppVersion"`

    ChipId  uint64  `json:"chipId"`

    //Vendor  string  `json:"vendorName"`
    //Model   string  `json:"modelName"`
    //Chip    string  `json:"chip"`
    //Ram     uint    `json:"ram"`
    //Rom     uint    `json:"rom"`
    //Cmos    string  `json:"cmos"`
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
    var schema Answer

    schema.App      = "app_tester"
    schema.Family   = chipFamily
    schema.Mpp      = version()
    schema.ChipId   = chipId()

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
