package webrtc

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "strconv"
    "net/url"

    "github.com/gorilla/mux"

    "application/core/logger"
    "application/core/streamer/webrtc"
)

func Connect(w http.ResponseWriter, r *http.Request) {
    var err error

    queryParams := mux.Vars(r)

    webrtcId, err := strconv.Atoi(queryParams["id"])
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    server, err := webrtc.GetById(webrtcId)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

	body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
		return
	}

	//ok, sdp := openapi.PostStringParameter(w, string(body), "sdp")
	//if !ok {
	//	return
	//}
    values, err := url.ParseQuery(string(body))
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }
    sdp, ok := values["sdp"]
    if !ok {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"TODO\"}")
        return
    }

    session, offer, err := server.Connect(sdp[0])

    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "sdp=%s&sessionId=%s", offer, session.SessionId) //TODO
}

func Disconnect(w http.ResponseWriter, r *http.Request) {
    var err error

    queryParams := mux.Vars(r)

    webrtcId, err := strconv.Atoi(queryParams["id"])
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    server, err := webrtc.GetById(webrtcId)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    logger.Log.Trace().
        Str("uuid", queryParams["uuid"]).
        Msg("Disconnect attemp")

    client, err := server.GetClientByUUID(queryParams["uuid"])
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    server.Disconnect(client)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    w.WriteHeader(http.StatusOK)
}
