package jpeg

import (
    "fmt"
    "net/http"
    "strconv"
    //"encoding/json"

    "github.com/gorilla/mux"

    "application/pkg/openapi"

    "application/pkg/streamer/jpeg"
)

func init() {
    openapi.AddApiRoute("jpegsList", "/jpeg", "GET", jpegsList)
    openapi.AddApiRoute("jpegCreate", "/jpeg", "POST", jpegCreate)
    openapi.AddApiRoute("jpegCreate", "/jpeg/{name}", "POST", jpegCreate)

    openapi.AddApiRoute("jpegInfo", "/jpeg/{id:[0-9]+}", "GET", jpegInfo)
    openapi.AddApiRoute("jpegDestroy", "/jpeg/{id:[0-9]+}", "DELETE", jpegDestroy)
}

func jpegsList(w http.ResponseWriter, r *http.Request) {
}

func jpegCreate(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

    var err error
    //var name string

    queryParams := mux.Vars(r)
    //name, err = strconv.Atoi(queryParams["name"])

    //if err != nil {
    //    w.WriteHeader(http.StatusInternalServerError)
    //    fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
    //    return
    //}

    _, err = jpeg.Create(queryParams["name"])
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "{\"ok\":\"jpeg created\"}")
}

func jpegInfo(w http.ResponseWriter, r *http.Request) {


}

func jpegDestroy(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

    var err error
    var id int

    queryParams := mux.Vars(r)
    id, err = strconv.Atoi(queryParams["id"])

    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    j, err := jpeg.GetById(id)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    err = jpeg.Delete(j)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "{\"ok\":\"jpeg deleted\"}")
}


