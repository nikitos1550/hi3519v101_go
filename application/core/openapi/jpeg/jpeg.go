package jpeg

import (
    "fmt"
    "net/http"
    "encoding/json"

    "github.com/gorilla/mux"

    "application/core/streamer/jpeg"
)

func GroupListHandler(g *jpeg.JpegGroup) func (http.ResponseWriter, *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        list := g.List()

        w.WriteHeader(http.StatusOK)

        listJson, _ := json.MarshalIndent(list, "", "\t")
        fmt.Fprintf(w, "%s", string(listJson))
    }
}

func GroupCreateHandler(g *jpeg.JpegGroup) func (http.ResponseWriter, *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        queryParams := mux.Vars(r)

        _, err := g.CreateJpeg(queryParams["name"])
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
            return
        }

        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, "{\"ok\":\"jpeg created\", \"name\": %s }", queryParams["name"])
    }
}

func GroupInfoHandler(g *jpeg.JpegGroup) func (http.ResponseWriter, *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        /*
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

        name, err := g.NameById(id)
        if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
            return
		}

        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, "{\"name\":\"%s\"}", name)
        */
    }
}

func GroupDeleteHandler(g *jpeg.JpegGroup) func (http.ResponseWriter, *http.Request) {
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
