//+build streamerFile

package file

import (
    "flag"
    //"fmt"
    //"log"
    "net/http"
    "application/pkg/openapi"

    //"application/pkg/mpp/venc"
)

var flagStoragePath     *string

func init() {
    flagStoragePath     = flag.String   ("streamer-file-storage",     "/opt/storage",              "files storage path")

    openapi.AddApiRoute("listRecords",  "/files",           "GET",      listRecords)

    openapi.AddApiRoute("infoRecord",   "/files/id",        "GET",      infoRecord)
    openapi.AddApiRoute("getRawRecord", "/files/id.h264",   "GET",      getRawRecord)
    openapi.AddApiRoute("getMP4Record", "/files/id.mp4",    "GET",      getMP4Record)
    openapi.AddApiRoute("deleteRecord", "/files/id",        "DELETE",   deleteRecord)

    openapi.AddApiRoute("statusRecord", "/files/record",    "GET",      statusRecord)
    openapi.AddApiRoute("startRecord",  "/files/record",    "POST",     startRecord)
    openapi.AddApiRoute("stopRecord",   "/files/record",    "DELETE",   stopRecord)
}

func Init() {}

func listRecords(w http.ResponseWriter, r *http.Request)  {}

func infoRecord(w http.ResponseWriter, r *http.Request)  {}
func getRawRecord(w http.ResponseWriter, r *http.Request)  {}
func getMP4Record(w http.ResponseWriter, r *http.Request)  {}
func deleteRecord(w http.ResponseWriter, r *http.Request)  {}

func statusRecord(w http.ResponseWriter, r *http.Request) {}
func startRecord(w http.ResponseWriter, r *http.Request)  {}
func stopRecord(w http.ResponseWriter, r *http.Request)  {}
