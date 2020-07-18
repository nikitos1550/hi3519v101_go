package encoder

import (
    "fmt"
    "net/http"
    "encoding/json"

    "github.com/gorilla/mux"

    "application/core/logger"
    "application/core/mpp/venc"
)

func GroupIdrHandler(g *venc.EncoderGroup) func (w http.ResponseWriter, r *http.Request) {
    return func (w http.ResponseWriter, r *http.Request) {
        var err error

        queryParams := mux.Vars(r)

        e, err := g.Get(queryParams["name"])
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
            return
        }

        started, err := e.IsStarted()
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
            return
        }

        if !started {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "{\"error\":\"Encoder is not started\"}")
            return
        }

        err = e.RequestIFrame()
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
            return
        }

        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, "{\"ok\":\"idr\"}")
    }
}

func GroupStartHandler(g *venc.EncoderGroup) func (w http.ResponseWriter, r *http.Request) { 
    return func (w http.ResponseWriter, r *http.Request) {
        var err error

        queryParams := mux.Vars(r)

        e, err := g.Get(queryParams["name"])
        if err != nil {
		    w.WriteHeader(http.StatusInternalServerError)
		    fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
		    return
        }

        started, err := e.IsStarted()
        if err != nil {
		    w.WriteHeader(http.StatusInternalServerError)
		    fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
		    return
        }

        if started {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "{\"error\":\"Encoder is already started\"}")
            return
        }

        err = e.Start()//TODO
        if err != nil {
		    w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
		    return
        }

        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, "{\"ok\":\"started\"}")
    }
}

func GroupStopHandler(g *venc.EncoderGroup) func (w http.ResponseWriter, r *http.Request) { 
    return func (w http.ResponseWriter, r *http.Request) {
        queryParams := mux.Vars(r)

        e, err := g.Get(queryParams["name"])
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
            return
        }

        started, err := e.IsStarted()
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
		    fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
		    return
        }

        if !started {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "{\"error\":\"Encoder is not started\"}")
            return
        }

        err = e.Stop()
        if err != nil {
		    w.WriteHeader(http.StatusInternalServerError)
		    fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
		    return
        }

        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, "{\"ok\":\"stoped\"}")
    }
}

func GroupNewHandler(g *venc.EncoderGroup) func (w http.ResponseWriter, r *http.Request) {
    return func (w http.ResponseWriter, r *http.Request) {
        var err error

        queryParams := mux.Vars(r)

        var params venc.Parameters

        venc.InvalidateBitrateControlParameters(&params.BitControlParams)
        venc.InvalidateGopParameters(&params.GopParams)

        err = json.NewDecoder(r.Body).Decode(&params)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
            return
        }

        e, err := g.CreateEncoder(queryParams["name"], params)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
            return
        }

        logger.Log.Trace().
            Msg("Encoder created")

        newParams, err := e.GetParams()
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
            return
        }

        w.WriteHeader(http.StatusOK)

        schemaJson, _ := json.MarshalIndent(newParams, "", "\t")
        fmt.Fprintf(w, "%s", string(schemaJson))
    }
}

type ListResponce struct {
    Amount      int         `json:"maxamount"`
    Encoders    []venc.Encoder   `json:"encoders"`
}

func GroupListHandler(g *venc.EncoderGroup) func (w http.ResponseWriter, r *http.Request) { 
    return func (w http.ResponseWriter, r *http.Request) {
        list := g.List()

        w.WriteHeader(http.StatusOK)

        listJson, _ := json.MarshalIndent(list, "", "\t")
        fmt.Fprintf(w, "%s", string(listJson))
    }
}

func GroupUpdateHandler(g *venc.EncoderGroup) func (w http.ResponseWriter, r *http.Request) { 
    return func (w http.ResponseWriter, r *http.Request) {
        //TODO
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"Not implemented\"}")
    }
}

func GroupDeleteHandler(g *venc.EncoderGroup) func (w http.ResponseWriter, r *http.Request) { 
    return func (w http.ResponseWriter, r *http.Request) {
        var err error

        queryParams := mux.Vars(r)

        err = g.Delete(queryParams["name"])
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
            return
        }

        w.WriteHeader(http.StatusOK)
    }
}

func GroupInfoHandler(g *venc.EncoderGroup) func (w http.ResponseWriter, r *http.Request) { 
    return func (w http.ResponseWriter, r *http.Request) {
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

        e, err := venc.GetEncoder(id)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
            return
        }

        info, err := e.GetCopy()
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
            return
        }

        w.WriteHeader(http.StatusOK)

        schemaJson, _ := json.MarshalIndent(info, "", "\t")
        fmt.Fprintf(w, "%s", string(schemaJson))
        */
    }
}
