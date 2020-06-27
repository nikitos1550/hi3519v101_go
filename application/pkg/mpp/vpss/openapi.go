//+build openapi

package vpss

import (
    "fmt"
    "net/http"
    "strconv"
    "encoding/json"

    "github.com/gorilla/mux"

	"application/pkg/openapi"
)

type responseRecord struct {
	Message string
}

func init() {
    openapi.AddApiRoute("apiDescription", "/mpp/channel", "GET", apiDescription)

    openapi.AddApiRoute("startChannel", "/mpp/channel/start", "GET", startChannelRequest)
    openapi.AddApiRoute("stopChannel", "/mpp/channel/stop", "GET", stopChannelRequest)
    openapi.AddApiRoute("listChannels", "/mpp/channel/list", "GET", listChannelsRequest)
    ////////////////////
    openapi.AddApiRoute("channelInfo", "/mpp/channel/{id:[0-9]+}", "GET", channelInfo)
    openapi.AddApiRoute("channelStat", "/mpp/channel/{id:[0-9]+}/stat", "GET", channelStat)

}

func channelInfo(w http.ResponseWriter, r *http.Request)  {
    w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

    var err error
    var id int

    queryParams := mux.Vars(r)
    id, err = strconv.Atoi(queryParams["id"])

    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    params, err := GetParams(id)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    w.WriteHeader(http.StatusOK)

    schemaJson, _ := json.Marshal(params)
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

    stat, err := GetStat(id)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    w.WriteHeader(http.StatusOK)

    schemaJson, _ := json.Marshal(stat)
    fmt.Fprintf(w, "%s", string(schemaJson))
}

func apiDescription(w http.ResponseWriter, r *http.Request)  {
	openapi.ApiDescription(w, r, "Channels api:\n\n", "/mpp/channel")
}

func startChannelRequest(w http.ResponseWriter, r *http.Request)  {
	var ok bool
    var id int
    var params Parameters

	ok, id = openapi.GetIntParameter(w, r, "channelId")
	if !ok {
		return
	}

	ok, params.Width = openapi.GetIntParameter(w, r, "width")
	if !ok {
		return
	}

	ok, params.Height = openapi.GetIntParameter(w, r, "height")
	if !ok {
		return
	}

	ok, params.Fps = openapi.GetIntParameter(w, r, "fps")
	if !ok {
		return
	}

	params.CropX = openapi.GetIntParameterOrDefault(w, r, "cropX", 0)
	if !ok {
		return
	}

	params.CropY = openapi.GetIntParameterOrDefault(w, r, "cropY", 0)
	if !ok {
		return
	}

	params.CropWidth = openapi.GetIntParameterOrDefault(w, r, "cropWidth", 0)
	if !ok {
		return
	}

	params.CropHeight = openapi.GetIntParameterOrDefault(w, r, "cropHeight", 0)
	if !ok {
		return
	}

	err := CreateChannel(id, params)
	if err != nil {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{Message: err.Error()})
		return
	}

	openapi.ResponseSuccessWithDetails(w, responseRecord{Message: "Channel was started"})
}


func stopChannelRequest(w http.ResponseWriter, r *http.Request) {
	ok, id := openapi.GetIntParameter(w, r, "channelId")
	if !ok {
		return
	}

	err := DestroyChannel(id)
	if err != nil {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{Message: err.Error()})
		return
	}

	openapi.ResponseSuccessWithDetails(w, responseRecord{Message: "Channel was stopped"})
}

type ChannelInfo struct {
    ChannelId   int
    Params      Parameters
    Processings []int
}

func listChannelsRequest(w http.ResponseWriter, r *http.Request) {
    var channelsInfo []ChannelInfo

    for i:=0; i< channelsAmount;i++ {
        t := ChannelInfo{}

        t.ChannelId = i
        t.Params, _ = GetParams(i)

        clients, _ := GetClientsTmp(i)

        for processing, _ := range clients {
            t.Processings = append(t.Processings, processing.GetId())
        }

		channelsInfo = append(channelsInfo, t)
	}
	openapi.ResponseSuccessWithDetails(w, channelsInfo)
}

