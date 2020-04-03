//+build streamerPipe

package pipe

import (
	"flag"
	"net/http"
	"os"
	"time"

	"application/pkg/openapi"
)

var flagStoragePath     *string

type pipe struct {
    Payload    chan []byte
	Started    bool
	EncoderId  string
	CurrentFile *os.File
	CurrentFilePath string
	Size int
	StartTime time.Time 
	Extention string
}

type pipeInfo struct {
	DownloadLink string
	Size int
	StartTime time.Time 
	EndTime time.Time 
	Duration string
	Md5 string
}

var (
	pipes map[string] pipe
)

func init() {
    flagStoragePath = flag.String("streamer-pipe-storage", "/opt/nfs", "pipes storage path")
	pipes = make(map[string] pipe)

    openapi.AddApiRoute("apiDescription", "/pipe", "GET", apiDescription)

    openapi.AddApiRoute("startPipe", "/pipe/start", "GET", startPipe)
    openapi.AddApiRoute("stopPipe", "/pipe/stop", "GET", stopPipe)

    openapi.AddApiRoute("listPipes", "/pipe/list", "GET", listPipes)
}

func Init() {
}

func apiDescription(w http.ResponseWriter, r *http.Request)  {
	openapi.ApiDescription(w, r, "Pipes api:\n\n", "/pipe")
}

func startPipe(w http.ResponseWriter, r *http.Request)  {

}

func stopPipe(w http.ResponseWriter, r *http.Request)  {

}

func listPipes(w http.ResponseWriter, r *http.Request)  {

}
