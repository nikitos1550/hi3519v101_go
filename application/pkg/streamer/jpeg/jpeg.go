package jpeg

import (
	"log"
	"net/http"
	"application/pkg/openapi"
)

func init() {
	openapi.AddRoute("serveJpeg",   "/jpeg/{stream}.[jpg|jpeg]",   "GET",      serveJpeg)
}

func serveJpeg(w http.ResponseWriter, r *http.Request) {
	log.Println("serveJpeg")
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotImplemented)
}