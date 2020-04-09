//+build openapi

package vpss

import (
	"application/pkg/openapi"
	"net/http"
	"unsafe"
)

type responseRecord struct {
	Message string
}

func init() {
    openapi.AddApiRoute("apiDescription", "/mpp/channel", "GET", apiDescription)

    openapi.AddApiRoute("startChannel", "/mpp/channel/start", "GET", startChannelRequest)
    openapi.AddApiRoute("stopChannel", "/mpp/channel/stop", "GET", stopChannelRequest)
    openapi.AddApiRoute("listChannels", "/mpp/channel/list", "GET", listChannelsRequest)
}

func apiDescription(w http.ResponseWriter, r *http.Request)  {
	openapi.ApiDescription(w, r, "Channels api:\n\n", "/mpp/channel")
}

func startChannelRequest(w http.ResponseWriter, r *http.Request)  {
	var channel Channel
	var ok bool

	ok, channel.ChannelId = openapi.GetIntParameter(w, r, "channelId")
	if !ok {
		return
	}

	ok, channel.Width = openapi.GetIntParameter(w, r, "width")
	if !ok {
		return
	}

	ok, channel.Height = openapi.GetIntParameter(w, r, "height")
	if !ok {
		return
	}

	ok, channel.Fps = openapi.GetIntParameter(w, r, "fps")
	if !ok {
		return
	}

	channel.CropX = openapi.GetIntParameterOrDefault(w, r, "cropX", 0)
	if !ok {
		return
	}

	channel.CropY = openapi.GetIntParameterOrDefault(w, r, "cropY", 0)
	if !ok {
		return
	}

	channel.CropWidth = openapi.GetIntParameterOrDefault(w, r, "cropWidth", channel.Width)
	if !ok {
		return
	}

	channel.CropHeight = openapi.GetIntParameterOrDefault(w, r, "cropHeight", channel.Height)
	if !ok {
		return
	}

	channel.Started = true
	channel.Clients = make(map[int] unsafe.Pointer)

	err, errorString := StartChannel(channel)
	if err != 0 {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{Message: errorString})
		return
	}

	openapi.ResponseSuccessWithDetails(w, responseRecord{Message: "Channel was started"})
}

func stopChannelRequest(w http.ResponseWriter, r *http.Request) {
	ok, channelId := openapi.GetIntParameter(w, r, "channelId")
	if !ok {
		return
	}

	err, errorString := StopChannel(channelId)
	if err != 0 {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{Message: errorString})
		return
	}

	openapi.ResponseSuccessWithDetails(w, responseRecord{Message: "Channel was stopped"})
}

func listChannelsRequest(w http.ResponseWriter, r *http.Request) {
	var channelsInfo []Channel
	for _, channel := range Channels {
		channelsInfo = append(channelsInfo, channel)
	}
	openapi.ResponseSuccessWithDetails(w, channelsInfo)
}
