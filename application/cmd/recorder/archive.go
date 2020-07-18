package main

import (
    "fmt"
    "net/http"
    "flag"

    "application/archive/record"
)

var(
    records map[string] *record.Record
    flagArchiveRawPath *string
)

func init() {
    flagArchiveRawPath     = flag.String("archive-raw-path",     "/opt/nfs",              "Raw archive dir path")
}

func initArchive() {
    records = make(map[string] *record.Record)

    //TODO scan dir for records
    //TODO load all records
}


//List all known records
func archiveList(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "archiveList")
}

//Show record information
func archiveItemInfo(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "archiveItemInfo")
}

//Show record preview
func archiveItemPreview(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "archiveItemPreview")
}

//Download record
func archiveServe(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "archiveServe")
}

//Delete record
func archiveDelete(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "archiveDelete")
}

////////////////////////////////////////////////////////////////////////////////

func recorderStatus(w http.ResponseWriter, r *http.Request) {}
func recorderStart(w http.ResponseWriter, r *http.Request) {}
func recorderStop(w http.ResponseWriter, r *http.Request) {}
func recorderSchedule(w http.ResponseWriter, r *http.Request) {}
