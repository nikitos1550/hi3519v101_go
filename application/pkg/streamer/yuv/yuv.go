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

	var dataPayload = make(chan []byte, 1)
	venc.SubsribeEncoderData(encoderId, dataPayload)
	logger.Log.Trace().
		Int("encoderId", encoderId).
		Msg("reed data from channel")
	data := <- dataPayload
	logger.Log.Trace().
                    Int("encoderId", encoderId).
                    Msg("reeded data from channel")
	venc.RemoveDataSubscription(encoderId, dataPayload)

	w.Header().Set("Content-Disposition", "attachment; filename=image.yuv")
	w.Header().Set("Content-Type", "application/octet-stream")

	n, err := w.Write(data)
	if err != nil {
		logger.Log.Warn().
			Msg("Failed to write data")
	} else {
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

	_, encoderExists := venc.ActiveEncoders[encoderId]
	if (!encoderExists) {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{Message: "Failed to find encoder  " + strconv.Itoa(encoderId)})
		return
	}

	serve(w, encoderId)
}
