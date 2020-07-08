package jpeg

import (
    "fmt"
    "strconv"
    "net/http"

    "github.com/gorilla/mux"

    "application/pkg/openapi"
)

func init() {
    openapi.AddRoute("serveJpegById",   "/jpeg/{id:[0-9]+}.jpg",   "GET",      serveJpegById)
    openapi.AddRoute("serveJpegById",   "/jpeg/{id:[0-9]+}.jpeg",   "GET",      serveJpegById)
    openapi.AddRoute("serveJpegByName",   "/jpeg/{name}.jpg",   "GET",      serveJpegByName)
    openapi.AddRoute("serveJpegByName",   "/jpeg/{name}.jpeg",   "GET",      serveJpegByName)
}

func serveJpegById(w http.ResponseWriter, r *http.Request) {
    queryParams := mux.Vars(r)
    id, err := strconv.Atoi(queryParams["id"])

    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    j, err := GetById(id)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    serve(w, j)
}

func serveJpegByName(w http.ResponseWriter, r *http.Request) {
    queryParams := mux.Vars(r)

    j, err := GetByName(queryParams["name"])
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    serve(w, j)
}

func serve(w http.ResponseWriter, j *jpeg) {
    source, err := j.getSource()
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    s, err := source.GetStorage()
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    w.Header().Set("Content-Type", "image/jpeg")
    s.WriteLastTo(w)
}
