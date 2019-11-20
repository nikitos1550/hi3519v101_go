package main

import (
    "fmt"
    "github.com/gorilla/mux"
    "net/http"
    //"application/testp"
    _"application/koloader"
)

func main() {
    //testp.TestHello()
    fmt.Println("test")

    router := mux.NewRouter()
    srv := &http.Server{
        Addr:           ":80",
        Handler:        router,
        //ReadTimeout:    10 * time.Second,
        //WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }
    go srv.ListenAndServe()

}
