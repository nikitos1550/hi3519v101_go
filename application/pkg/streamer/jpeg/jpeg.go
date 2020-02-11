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

    //TODO loop moved to VENC package
    //venc.TmpLock()
    //w.Write(([]byte)(venc.TmpGet()))
    //venc.TmpUnlock()

    /*
    size, _ := venc.F.Size()
    buf     := make([]byte, size)
    venc.F.Read(buf)
    w.Write(buf)
    */

    //copied, err := venc.F.WriteTo(w)
    copied, err := venc.SampleMjpegFrames.WriteTo(w)
    log.Println("serveJpeg copied", copied, " error", err)
}
