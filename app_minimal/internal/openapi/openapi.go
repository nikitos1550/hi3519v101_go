package openapi

import (
    "net/http"
    "github.com/gorilla/mux"
)


func NewRouter() *mux.Router {
    router := mux.NewRouter() //.StrictSlash(true)

    //router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("/opt/www"))))

    router.HandleFunc("/api/system", system).Methods("GET")
    router.HandleFunc("/api/system/temperature", systemTemperature).Methods("GET")
    router.HandleFunc("/api/system/date", systemDate).Methods("GET")

    router.HandleFunc("/api/debug/go", debugGo).Methods("GET")
    router.HandleFunc("/api/debug/umap", debugUmap).Methods("GET")
    router.HandleFunc("/api/debug/umap/{file}", debugUmapFile).Methods("GET")

    router.HandleFunc("/api/encoders", encodersList).Methods("GET")
    //router.HandleFunc("/api/encoders", ???).Methods("POST")
    //router.HandleFunc("/api/encoders/{:id}", ???).Methods("GET")
    //router.HandleFunc("/api/encoders/{:id}", ???).Methods("PUT")
    //router.HandleFunc("/api/encoders/{:id}", ???).Methods("DELETE")

    router.HandleFunc("/api/encoders/{id:[0-9]+}/image.jpeg", encodersServeJpeg).Methods("GET")
    //router.HandleFunc("/api/encoders/0/image.jpeg", encodersServeJpeg).Methods("GET")

    router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("/opt/www"))))
    return router
}

