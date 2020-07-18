package crud

import (
    "fmt"
    "net/http"
    "encoding/json"

    "github.com/gorilla/mux"
)

type crud interface {
    List() []string
    Create(string) error
    Delete(string) error
}

func GroupListHandler(g crud) func (http.ResponseWriter, *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        list := g.List()

        w.WriteHeader(http.StatusOK)

        listJson, _ := json.MarshalIndent(list, "", "\t")
        fmt.Fprintf(w, "%s", string(listJson))
    }
}

func GroupCreateHandler(g crud) func (http.ResponseWriter, *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        queryParams := mux.Vars(r)

        err := g.Create(queryParams["name"])
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
            return
        }

        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, "{\"ok\":\"jpeg created\", \"name\": %s }", queryParams["name"])
    }
}

func GroupDeleteHandler(g crud) func (http.ResponseWriter, *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        var err error

        queryParams := mux.Vars(r)

        err = g.Delete(queryParams["name"])
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
            return
        }

        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, "{\"ok\":\"jpeg deleted\"}")
    }
}
