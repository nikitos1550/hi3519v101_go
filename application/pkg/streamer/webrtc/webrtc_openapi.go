//+build streamerWebrtc

package webrtc

import (
    "fmt"
//	"log"

 	"io/ioutil"
    "net/http"
    "application/pkg/openapi"
)

type responseRecord struct {
	Message string
}

func init() {
    openapi.AddApiRoute("connectWebrtc",   "/webrtc/connect",   "POST",      connectWebrtc)
    openapi.AddApiRoute("disconnectWebrtc",   "/webrtc/disconnect",   "GET",      disconnectWebrtc)
}

func Init() {
	go WebrtcInit()
}

func connectWebrtc(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
    if err != nil {
		openapi.ResponseErrorWithDetails(w, http.StatusBadRequest, responseRecord{Message: "Failed to read request body"})
		return
	}

	ok, sdp := openapi.PostStringParameter(w, string(body), "sdp")
	if !ok {
		return
	}

	ok, encoderId := openapi.PostIntParameter(w, string(body), "encoderId")
	if !ok {
		return
	}

	e,sessionId,serverSdp := WebrtcConnect(sdp, encoderId)
	if e < 0 || sessionId == "" {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{Message: "Failed to create webrtc session"})
		return
	}

    w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "sdp=%s&sessionId=%s", serverSdp, sessionId)
}

func disconnectWebrtc(w http.ResponseWriter, r *http.Request) {
	ok, sessionId := openapi.GetStringParameter(w, r, "sessionId")
	if !ok {
		return
	}

	err,errString := WebrtcDisconnect(sessionId)
	if (err < 0) {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{Message: errString})
		return
	}

    w.WriteHeader(http.StatusOK)
}