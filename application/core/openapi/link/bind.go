package link

import (
    "fmt"
    "net/http"

    "github.com/gorilla/mux"

    "application/core/mpp/connection"
)

type bindSourcer interface {
    GetSourceBind(name string) (connection.SourceBind, error)
}

type bindClienter interface {
    GetClientBind(name string) (connection.ClientBind, error)
}

func ConnectBindHandler(s bindSourcer, c bindClienter) func (w http.ResponseWriter, r *http.Request) {
    return func  (w http.ResponseWriter, r *http.Request) {
        var err error
        queryParams := mux.Vars(r)

        source, err := s.GetSourceBind(queryParams["source"])
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
            return
        }

        client, err := c.GetClientBind(queryParams["client"])
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
            return
        }

        err = connection.ConnectBind(source, client)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
            return
        }

        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, "{\"ok\":\"binded\"}")
    }
}

func DisconnectBindHandler(s bindSourcer, c bindClienter) func (w http.ResponseWriter, r *http.Request) {
    return func  (w http.ResponseWriter, r *http.Request) {
        var err error
        queryParams := mux.Vars(r)

        source, err := s.GetSourceBind(queryParams["source"])
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
            return
        }

        client, err := c.GetClientBind(queryParams["client"])
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
            return
        }

        err = connection.DisconnectBind(source, client)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
            return
        }

        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, "{\"ok\":\"binded\"}")
    }
}
