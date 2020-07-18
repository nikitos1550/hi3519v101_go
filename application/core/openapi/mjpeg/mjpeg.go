package mjpeg

import (
    "fmt"
    "net/http"
    "encoding/json"

    "github.com/gorilla/mux"

    "application/core/streamer/mjpeg"
)

func GroupListHandler(g *mjpeg.MjpegGroup) func (w http.ResponseWriter, r *http.Request) {
    return func (w http.ResponseWriter, r *http.Request) {
        list := g.List()

        w.WriteHeader(http.StatusOK)

        listJson, _ := json.MarshalIndent(list, "", "\t")
        fmt.Fprintf(w, "%s", string(listJson))
    }
}

func GroupCreateHandler(g *mjpeg.MjpegGroup) func (w http.ResponseWriter, r *http.Request) {
    return func (w http.ResponseWriter, r *http.Request) {
        queryParams := mux.Vars(r)

        _, err := g.CreateMjpeg(queryParams["name"])
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
            return
        }

        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, "{\"ok\":\"jpeg created\", \"name\": %s }", queryParams["name"])
    }
}

func GroupInfoHandler(g *mjpeg.MjpegGroup) func (w http.ResponseWriter, r *http.Request) {
    return func (w http.ResponseWriter, r *http.Request) {
         queryParams := mux.Vars(r)

         _, err := g.Get(queryParams["name"])

         if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
            return
         }

        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, "{\"todo\":\"TODO\"}")
    }
}

func GroupDestroyHandler(g *mjpeg.MjpegGroup) func (w http.ResponseWriter, r *http.Request) {
    return func (w http.ResponseWriter, r *http.Request) {
        queryParams := mux.Vars(r)

        err := g.Delete(queryParams["name"])
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
            return
        }

        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, "{\"ok\":\"mjpeg deleted\"}")
    }
}

////////////////////////////////////////////////////////////////////////////////

func GroupClientsHandler(g *mjpeg.MjpegGroup) func (w http.ResponseWriter, r *http.Request) {
    return func (w http.ResponseWriter, r *http.Request) {
        queryParams := mux.Vars(r)

        mjpeg, err := g.Get(queryParams["name"])
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
            return
        }

        clients := mjpeg.Clients()

        w.WriteHeader(http.StatusOK)

        clientsJson, _ := json.MarshalIndent(clients, "", "\t")
        fmt.Fprintf(w, "%s", string(clientsJson))
    }
}

func GroupClientDeleteHandler(g *mjpeg.MjpegGroup) func (w http.ResponseWriter, r *http.Request) {
    return func (w http.ResponseWriter, r *http.Request) {
        queryParams := mux.Vars(r)

        mjpeg, err := g.Get(queryParams["name"])
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
            return
        }

        err = mjpeg.DropClient(queryParams["client"])
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
            return
        }

        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, "{\"dropped\":\"%s\"}", queryParams["client"])
    }
}
