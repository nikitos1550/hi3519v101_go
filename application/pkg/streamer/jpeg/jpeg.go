package jpeg

import (
    //"fmt"
	"log"
	"net/http"
	"application/pkg/openapi"

    "application/pkg/mpp/getloop"
)

func init() {
	//openapi.AddRoute("serveJpeg",   "/jpeg/{stream}.[jpg|jpeg]",   "GET",      serveJpeg)
    openapi.AddRoute("serveJpeg",   "/jpeg/1.jpg",   "GET",      serveJpeg)
}

func Init() {}

func serveJpeg(w http.ResponseWriter, r *http.Request) {
	log.Println("serveJpeg")

	w.Header().Set("Content-Type", "image/jpeg")

    getloop.TmpLock()
    w.Write(([]byte)(getloop.TmpGet()))
    getloop.TmpUnlock()
}
