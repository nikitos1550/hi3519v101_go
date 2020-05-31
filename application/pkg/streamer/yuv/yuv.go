//+build streamerYuv

package yuv

import (
	"strconv"

	"application/pkg/logger"

	"net/http"
	"application/pkg/openapi"

    "application/pkg/mpp/venc"
)

type responseRecord struct {
	Message string
}

func init() {
    openapi.AddRoute("serveYuv",   "/yuv/image.yuv",   "GET",      serveYuv)
}

func Init() {}

func serve(w http.ResponseWriter, encoderId int) {
	logger.Log.Trace().
		Msg("serveYuv")

	var payload = make(chan []byte, 1)
	venc.SubsribeEncoder(encoderId, payload)
	//log.Println("reed data from channel ")
		logger.Log.Trace().
			Int("encoderId", encoderId).
			Msg("reed data from channel")
	data := <- payload
	//log.Println("reeded data from channel ")
		logger.Log.Trace().
                        Int("encoderId", encoderId).
                        Msg("reeded data from channel")
	venc.RemoveSubscription(encoderId, payload)

	w.Header().Set("Content-Type", "image/jpeg")

	n, err := w.Write(data)
	if err != nil {
		//log.Println("Failed to write data")
		logger.Log.Warn().
			Msg("Failed to write data")
	} else {
		//log.Println("written size is ", n)
		logger.Log.Trace().
			Int("size", n).
			Msg("written size")
	}
}

func serveYuv(w http.ResponseWriter, r *http.Request) {
	ok, encoderId := openapi.GetIntParameter(w, r, "encoderId")
	if !ok {
		return
	}

	encoder, encoderExists := venc.ActiveEncoders[encoderId]
	if (!encoderExists) {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{Message: "Failed to find encoder  " + strconv.Itoa(encoderId)})
		return
	}

	if (encoder.Format != "mjpeg") {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{Message: "Encoder has wrong format " + encoder.Format + ". Should be mjpeg"})
		return
	}

	serve(w, encoderId)
}
