package link

import (
    "fmt"
    "net/http"

    "github.com/gorilla/mux"

    "application/core/mpp/connection"
)

type encodedDataSourcer interface {
    GetSourceEncodedData(name string) (connection.SourceEncodedData, error)
}

type encodedDataClienter interface {
    GetClientEncodedData(name string) (connection.ClientEncodedData, error)
}

func ConnectEncodedDataHandler(s encodedDataSourcer, c encodedDataClienter) func (w http.ResponseWriter, r *http.Request) {
    return func  (w http.ResponseWriter, r *http.Request) {
        var err error
        queryParams := mux.Vars(r)

        source, err := s.GetSourceEncodedData(queryParams["source"])
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
            return
        }

        client, err := c.GetClientEncodedData(queryParams["client"])
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
            return
        }

        err = connection.ConnectEncodedData(source, client)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
            return
        }

        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, "{\"ok\":\"binded\"}")
    }
}

func DisconnectEncodedDataHandler(s encodedDataSourcer, c encodedDataClienter) func (w http.ResponseWriter, r *http.Request) {
    return func  (w http.ResponseWriter, r *http.Request) {
        var err error
        queryParams := mux.Vars(r)

        source, err := s.GetSourceEncodedData(queryParams["source"])
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
            return
        }

        client, err := c.GetClientEncodedData(queryParams["client"])
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
            return
        }

        err = connection.DisconnectEncodedData(source, client)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
            return
        }

        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, "{\"ok\":\"binded\"}")
    }
}
