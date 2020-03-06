//+build streamerJpeg

package jpeg

import (
	"log"
	"net/http"
	"application/pkg/openapi"

    "application/pkg/mpp/venc"
)

func init() {
    openapi.AddRoute("serveJpeg",   "/jpeg/1.jpg",   "GET",      serveJpeg)
}

func Init() {}

func serveJpeg(w http.ResponseWriter, r *http.Request) {
	log.Println("serveJpeg")

	var payload = make(chan []byte, 1)
	encoderId := 1
	venc.SubsribeEncoder(encoderId, payload)
	data := <- payload
	venc.RemoveSubscription(encoderId)

	w.Header().Set("Content-Type", "image/jpeg")

	n, err := w.Write(data)
	if err != nil {
		log.Println("Failed to write data")
	} else {
		log.Println("written size is ", n)
	}
}
