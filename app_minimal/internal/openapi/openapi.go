package openapi

import (
    "net/http"
    "github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
    router := mux.NewRouter() //.StrictSlash(true)

    router.HandleFunc("/api/system", system).Methods("GET")
    router.HandleFunc("/api/system/temperature", systemTemperature).Methods("GET")
    router.HandleFunc("/api/system/date", systemDate).Methods("GET")
    router.HandleFunc("/api/system/network", systemNetwork).Methods("GET")

    router.HandleFunc("/api/debug/go", debugGo).Methods("GET")
    router.HandleFunc("/api/debug/umap", debugUmap).Methods("GET")
    router.HandleFunc("/api/debug/umap/{file}", debugUmapFile).Methods("GET")

    router.HandleFunc("/api/channels", channelsList).Methods("GET")
    router.HandleFunc("/api/channels/{id:[0-9]+}", channelInfo).Methods("GET")
    router.HandleFunc("/api/channels/{id:[0-9]+}", channelStart).Methods("POST")
    //router.HandleFunc("/api/channels/{id:[0-9]+}", channelUpdate).Methods("PUT") //Seems there will not be any dynamic settings
    router.HandleFunc("/api/channels/{id:[0-9]+}", channelStop).Methods("DELETE")

    router.HandleFunc("/api/encoders", encodersList).Methods("GET")
    router.HandleFunc("/api/encoders", encoderCreate).Methods("POST")
    router.HandleFunc("/api/encoders/{id:[0-9]+}", encoderInfo).Methods("GET")
    router.HandleFunc("/api/encoders/{id:[0-9]+}", encoderUpdate).Methods("PUT")
    router.HandleFunc("/api/encoders/{id:[0-9]+}", encoderDelete).Methods("DELETE")

    router.HandleFunc("/api/encoders/{id:[0-9]+}/image.jpeg", encoderServeJpeg).Methods("GET")

    router.HandleFunc("/api/encoders/{id:[0-9]+}/record", encoderRecordStatus).Methods("GET")
    router.HandleFunc("/api/encoders/{id:[0-9]+}/record", encoderRecordStart).Methods("POST")
    router.HandleFunc("/api/encoders/{id:[0-9]+}/record", encoderRecordStop).Methods("DELETE")

    router.HandleFunc("/api/encoders/{id:[0-9]+}/rtp", encoderRtpStatus).Methods("GET")
    router.HandleFunc("/api/encoders/{id:[0-9]+}/rtp", encoderRtpStart).Methods("POST")
    router.HandleFunc("/api/encoders/{id:[0-9]+}/rtp", encoderRtpStop).Methods("DELETE")

    router.HandleFunc("/api/storage", storageList).Methods("GET")
    router.HandleFunc("/api/storage/{id:[0-9]+}", storageRecordInfo).Methods("GET")
    router.HandleFunc("/api/storage/{id:[0-9]+}", storageRecordDelete).Methods("DELETE")

    router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("/opt/www"))))

    return router
}

