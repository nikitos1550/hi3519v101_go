//+build streamerPipe

package pipe

import (
	"flag"
	"net/http"
	"path"
	"os"
	"strconv"    
	"syscall"
	"log"
	
	"application/pkg/mpp/venc"
	"application/pkg/openapi"
)

var flagStoragePath     *string

type pipe struct {
    Payload    chan []byte
	Started    bool
	EncoderId  int
	File *os.File
	FilePath string
}

type pipeInfo struct {
	EncoderId  int
	FilePath string
}

type responseRecord struct {
	Message string
}

var (
	pipes map[string] pipe
)

func init() {
    flagStoragePath = flag.String("streamer-pipe-storage", "/tmp", "pipes storage path")
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
	ok, encoderId := openapi.GetIntParameter(w, r, "encoderId")
	if !ok {
		return
	}

	_, encoderExists := venc.ActiveEncoders[encoderId]
	if (!encoderExists) {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{Message: "Failed to find encoder  " + strconv.Itoa(encoderId)})
		return
	}

	ok, pipeName := openapi.GetStringParameter(w, r, "pipeName")
	if !ok {
		return
	}

	pipePath := path.Join(*flagStoragePath, pipeName)
	os.Remove(pipePath)
	err := syscall.Mkfifo(pipePath, 0666)
	
	if err != nil {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{Message: "Failed to create fifo"})
		return
	}

	file, err := os.OpenFile(pipePath, os.O_RDWR, os.ModeNamedPipe)
	if err != nil {
		log.Println("Failed to create pipe")
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{Message: "Failed to create pipe"})
		return
	}

	pipes[pipeName] = pipe{
		Payload: make(chan []byte, 100),
		Started: true,
		EncoderId: encoderId,
		File: file,
		FilePath: pipePath,
	}

	venc.SubsribeEncoder(encoderId, pipes[pipeName].Payload)
	
    go func() {
		writeVideoData(pipeName)
    }()

	openapi.ResponseSuccessWithDetails(w, responseRecord{Message: "Pipe was started"})
}

func stopPipe(w http.ResponseWriter, r *http.Request)  {
	ok, pipeName := openapi.GetStringParameter(w, r, "pipeName")
	if !ok {
		return
	}

	pipe, exists := pipes[pipeName]
	if (!exists) {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{Message: "Pipe not found"})
		return
	}

	pipe.Started = false
	pipes[pipeName] = pipe
	openapi.ResponseSuccessWithDetails(w, responseRecord{Message: "Pipe was stopped"})
}

func listPipes(w http.ResponseWriter, r *http.Request)  {
	var infos []pipeInfo
	for _, pipe := range pipes {
		info := pipeInfo{
			EncoderId: pipe.EncoderId,
			FilePath: pipe.FilePath,
		}

		infos = append(infos, info)
	}
	openapi.ResponseSuccessWithDetails(w, infos)
}

func writeVideoData(pipeName string){
	pipeChannel := make(chan []byte, 100)

	go func() {
		for {
			if (!pipes[pipeName].Started){
				break
			}

			pipeData := <- pipeChannel
			pipes[pipeName].File.Write(pipeData)
		}
	}()

	for {
		if (!pipes[pipeName].Started){
			break
		}

		data := <- pipes[pipeName].Payload
		// write data to temporary channel to extract all data from venc and skip here if pipe does not have readers
		if (cap(pipeChannel) > len(pipeChannel)) {
			pipeChannel <- data
		}
	}

	venc.RemoveSubscription(pipes[pipeName].EncoderId, pipes[pipeName].Payload)
	pipes[pipeName].File.Close()
	delete(pipes, pipeName)
}
