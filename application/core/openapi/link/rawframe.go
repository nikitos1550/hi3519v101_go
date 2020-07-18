package link

import (
    "fmt"
    "net/http"

    "github.com/gorilla/mux"

    "application/core/mpp/connection"
)

type rawFrameSourcer interface {
    GetSourceRawFrame(name string) (connection.SourceRawFrame, error)
}

type rawFrameClienter interface {
    GetClientRawFrame(name string) (connection.ClientRawFrame, error)
}

func ConnectRawFrameHandler(s rawFrameSourcer, c rawFrameClienter) func (w http.ResponseWriter, r *http.Request) {
    return func  (w http.ResponseWriter, r *http.Request) {
        var err error
        queryParams := mux.Vars(r)

        source, err := s.GetSourceRawFrame(queryParams["source"])
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
            return
        }

        client, err := c.GetClientRawFrame(queryParams["client"])
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
            return
        }

        err = connection.ConnectRawFrame(source, client)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
            return
        }

        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, "{\"ok\":\"binded\"}")
    }
}

func DisconnectRawFrameHandler(s rawFrameSourcer, c rawFrameClienter) func (w http.ResponseWriter, r *http.Request) {
    return func  (w http.ResponseWriter, r *http.Request) {
        var err error
        queryParams := mux.Vars(r)

        source, err := s.GetSourceRawFrame(queryParams["source"])
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
            return
        }

        client, err := c.GetClientRawFrame(queryParams["client"])
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
            return
        }

        err = connection.DisconnectRawFrame(source, client)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
            return
        }

        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, "{\"ok\":\"binded\"}")
    }
}
