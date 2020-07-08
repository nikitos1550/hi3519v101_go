package channel

import (
    "fmt"
    "net/http"
    "strconv"
    "encoding/json"

    "github.com/gorilla/mux"

	"application/pkg/openapi"

    "application/pkg/mpp/vpss"

    "application/pkg/logger"
)

func init() {
    openapi.AddApiRoute("channelsList", "/mpp/channels", "GET", channelsList)

    openapi.AddApiRoute("channelInfo", "/mpp/channels/{id:[0-9]+}", "GET", channelInfo)
    openapi.AddApiRoute("channelCreate", "/mpp/channels/{id:[0-9]+}", "POST", channelCreate)
    openapi.AddApiRoute("channelDestroy", "/mpp/channels/{id:[0-9]+}", "DELETE", channelDestroy)

    openapi.AddApiRoute("channelStat", "/mpp/channels/{id:[0-9]+}/stat", "GET", channelStat)
}

type channelsListResponce struct {
    Amount      int             `json:"maxamount"`
    Channels    []vpss.Channel  `json:"channels"`
}

func channelsList(w http.ResponseWriter, r *http.Request)  {

    w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

    var response channelsListResponce

    response.Amount = vpss.Amount

    for i:=0; i < vpss.Amount; i++ {
        c, _ := vpss.GetChannel(i)
        cpy, err := c.GetCopy()
        if err != nil {
            response.Channels = append(response.Channels, cpy)
        }
    }

    w.WriteHeader(http.StatusOK)

    schemaJson, _ := json.MarshalIndent(response, "", "\t")
    fmt.Fprintf(w, "%s", string(schemaJson))

}

func channelCreate(w http.ResponseWriter, r *http.Request)  {
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

    var params vpss.Parameters

    err = json.NewDecoder(r.Body).Decode(&params)
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

    err = c.CreateChannel(params, false)
    if err != nil {
        //goto returnError
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

//returnError:
//    w.WriteHeader(http.StatusInternalServerError)
//    fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
//    return
}

func channelDestroy(w http.ResponseWriter, r *http.Request)  {
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
}

func channelInfo(w http.ResponseWriter, r *http.Request)  {
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
}

func channelStat(w http.ResponseWriter, r *http.Request)  {
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
}
