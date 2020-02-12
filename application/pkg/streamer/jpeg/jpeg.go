//+build streamerJpeg

package jpeg

import (
    //"fmt"
	"log"
	"net/http"
	"application/pkg/openapi"

    "application/pkg/mpp/venc"
)

func init() {
	//openapi.AddRoute("serveJpeg",   "/jpeg/{stream}.[jpg|jpeg]",   "GET",      serveJpeg)
    openapi.AddRoute("serveJpeg",   "/jpeg/1.jpg",   "GET",      serveJpeg)
}

func Init() {}

func serveJpeg(w http.ResponseWriter, r *http.Request) {
	log.Println("serveJpeg")

	w.Header().Set("Content-Type", "image/jpeg")

    //copied, err := venc.F.WriteTo(w)
    //size, seq, _ := venc.SampleMjpegFrames.GetLastFrame().Info()
    //log.Println("size=", size, " seq=", seq)
    venc.SampleMjpegFrames.WriteLastTo(w)
    //copied, err := venc.SampleMjpegFrames.WriteLastTo(w)
    //log.Println("serveJpeg copied", copied, " error", err)
}
