//+build streamerFile

package file

import (
    "flag"
    //"fmt"
    "log"
    "net/http"
	"os"
    "application/pkg/openapi"
	"application/pkg/mpp/venc"
)

var flagStoragePath     *string

var (
    Payload    chan []byte
	Started    bool
	F *os.File
)

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

    openapi.AddApiRoute("startNewRecord", "/files/record/start", "GET",      startNewRecord)
    openapi.AddApiRoute("startNewRecord", "/files/record/stop", "GET",      stopNewRecord)
	Started = false
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

func startNewRecord(w http.ResponseWriter, r *http.Request)  {
	Payload  = make(chan []byte, 100)
	venc.SubsribeEncoder(0, Payload)
	Started = true
	venc.SampleH264Start <- 100
	F,err := os.Create("/tmp/out.h264")
	if err != nil {
        log.Println("Error open file ")
    }
    go func() {
		for {
			if (!Started){
				break
			}

			data := <- Payload
			log.Println("Readed ", len(data))
			F.Write(data)
		}
    }()
}

func stopNewRecord(w http.ResponseWriter, r *http.Request)  {
	venc.RemoveSubscription(0)
	Started = false
	F.Close()
}