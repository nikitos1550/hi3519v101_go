package channel

import (
    "fmt"
    "net/http"
    "encoding/json"

    "github.com/gorilla/mux"

    "application/core/mpp/vpss"
    "application/core/logger"
)

type ListResponce struct {
    Amount      int             `json:"maxamount"`
    Channels    []vpss.Channel  `json:"channels"`
}

func GroupListHandler(g *vpss.ChannelGroup) func (http.ResponseWriter, *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        list := g.List()

        w.WriteHeader(http.StatusOK)

        listJson, _ := json.MarshalIndent(list, "", "\t")
        fmt.Fprintf(w, "%s", string(listJson))
    }
}

func GroupCreateHandler(g *vpss.ChannelGroup) func (http.ResponseWriter, *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        var err error

        queryParams := mux.Vars(r)

        var params vpss.Parameters

        err = json.NewDecoder(r.Body).Decode(&params)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
            return
        }

        c, err := g.CreateChannel(queryParams["name"], params)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
            return
        }

        logger.Log.Trace().
            Msg("Channel created")

        newParams, err := c.GetParams()
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

func GroupDeleteHandler(g *vpss.ChannelGroup) func (http.ResponseWriter, *http.Request) {
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

    c, err := vpss.GetChannel(id)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    logger.Log.Trace().
        Msg("GetChannel ok")

    err = c.DestroyChannel()
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    w.WriteHeader(http.StatusOK)
    */
    }
}

func GroupInfoHandler(g *vpss.ChannelGroup) func (http.ResponseWriter, *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
    /*
    func Info(w http.ResponseWriter, r *http.Request)  {
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

    c, err := vpss.GetChannel(id)

    info, err := c.GetCopy()
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

func GroupStatHandler(g *vpss.ChannelGroup) func (http.ResponseWriter, *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
    /*
    w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

    var err error
    var id int

    queryParams := mux.Vars(r)
    id, err = strconv.Atoi(queryParams["id"])

    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    c, err := vpss.GetChannel(id)

    stat, err := c.GetStat()
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    w.WriteHeader(http.StatusOK)

    schemaJson, _ := json.Marshal(stat)
    fmt.Fprintf(w, "%s", string(schemaJson))
    */
    }
}
