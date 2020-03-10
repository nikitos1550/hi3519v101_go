//+build streamerJpeg

package jpeg

import (
	"log"
	"net/http"
	"application/pkg/openapi"

    "application/pkg/mpp/venc"
)

func init() {
    openapi.AddRoute("serveHdJpeg",   "/jpeg/1920_1080.jpg",   "GET",      serveHdJpeg)
    openapi.AddRoute("serve4KJpeg",   "/jpeg/3840_2160.jpg",   "GET",      serve4KJpeg)
}

func Init() {}

func serve(w http.ResponseWriter, encoderId string) {
	log.Println("serveJpeg")

	var payload = make(chan []byte, 1)
	venc.SubsribeEncoder(encoderId, payload)
	log.Println("reed data from channel ")
	data := <- payload
	log.Println("reeded data from channel ")
	venc.RemoveSubscription(encoderId, payload)

	w.Header().Set("Content-Type", "image/jpeg")

	n, err := w.Write(data)
	if err != nil {
		log.Println("Failed to write data")
	} else {
		log.Println("written size is ", n)
	}
}

func serveHdJpeg(w http.ResponseWriter, r *http.Request) {
	serve(w, "MGPEG_1920_1080")
}

func serve4KJpeg(w http.ResponseWriter, r *http.Request) {
	serve(w, "MGPEG_3840_2160")
}
