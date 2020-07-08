package webrtc

import (
    //"fmt"
    "net/http"

    "application/pkg/openapi"
    //"application/pkg/logger"
)

func init() {
    openapi.AddRoute("webrtcConnect",   "/webrtc/{id:[0-9]+}",   "POST",    webrtcConnect)

    openapi.AddRoute("webrtcConnectInfo",   "/webrtc/{id:[0-9]+}/{guid:[0-9a-z_]+}",   "GET",    webrtcConnectInfo)
    openapi.AddRoute("webrtcDisconnect",   "/webrtc/{id:[0-9]+}/{guid:[0-9a-z_]+}",   "DELETE",  webrtcDisconnect)
}

func webrtcConnect(w http.ResponseWriter, r *http.Request) {
}

func webrtcConnectInfo(w http.ResponseWriter, r *http.Request) {
    /*
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

    session := uuid.New().String()
    WebrtcSessions[session] = &WebrtcSession{}

    logger.Log.Trace().
        Str("session", session).
        Str("sdp", "..."). //sdp).
        Int("encoder", encoderId).
        Msg("WebRTC new session")

    offer, err := WebrtcSessions[session].Connect(sdp, encoderId)

    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

	//e, sessionId, serverSdp := WebrtcConnect(sdp, encoderId)
	//if e < 0 || sessionId == "" {
	//	openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{Message: "Failed to create webrtc session"})
	//	return
	//}

    w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "sdp=%s&sessionId=%s", offer, session) //TODO
    */
}

func webrtcDisconnect(w http.ResponseWriter, r *http.Request) {
    /*
	ok, sessionId := openapi.GetStringParameter(w, r, "sessionId")
	if !ok {
		return
	}

    _, exist := WebrtcSessions[sessionId]
    if !exist {
        openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{Message: "no such session"})
    }

	//err,errString := WebrtcDisconnect(sessionId)
	//if (err < 0) {
	//	openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{Message: errString})
	//	return
	//}

    w.WriteHeader(http.StatusOK)
    */
}
