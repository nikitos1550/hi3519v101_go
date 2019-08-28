package openapi

import (
    "net/http"
    "../himpp3"
    "log"
    "strconv"
    "github.com/gorilla/mux"
)

func channelsList (w http.ResponseWriter, r *http.Request) {
    //w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusNotImplemented)
}

func channelInfo (w http.ResponseWriter, r *http.Request) {
    //w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusNotImplemented)

    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    log.Println("channelInfo ", id)

    himpp3.GetChannelInfo(uint(id))
}

func channelStart (w http.ResponseWriter, r *http.Request) {
    //w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusNotImplemented)
}

//func channelUpdate (w http.ResponseWriter, r *http.Request) {
//    //w.Header().Set("Content-Type", "application/json; charset=UTF-8")
//    w.WriteHeader(http.StatusNotImplemented)
//}

func channelStop (w http.ResponseWriter, r *http.Request) {
    //w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusNotImplemented)
}




