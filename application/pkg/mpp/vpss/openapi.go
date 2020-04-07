//+build openapi

package vpss

import (
	"application/pkg/openapi"
	"net/http"
    "sync"
)

type responseRecord struct {
	Message string
}

type Channel struct {
	ChannelId  int
	Width int
	Height int
	Fps int
	CropX int
	CropY int
	CropWidth int
	CropHeight int
    Mutex sync.RWMutex
	Started bool
}

var (
	channels map[int] Channel
)

func init() {
	channels = make(map[int] Channel)

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

	err, errorString := startChannel(channel)
	if err != 0 {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{Message: errorString})
		return
	}

	openapi.ResponseSuccessWithDetails(w, responseRecord{Message: "Channel was started"})
}

func startChannel(channel Channel)  (int, string)  {
	_, channelExists := channels[channel.ChannelId]
	if (channelExists) {
		return 1, "Channel already exists"
	}

	CreateChannel(channel)

	channels[channel.ChannelId] = channel
	return 0, ""
}

func stopChannelRequest(w http.ResponseWriter, r *http.Request) {
	ok, channelId := openapi.GetIntParameter(w, r, "channelId")
	if !ok {
		return
	}

	err, errorString := stopChannel(channelId)
	if err != 0 {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{Message: errorString})
		return
	}

	openapi.ResponseSuccessWithDetails(w, responseRecord{Message: "Channel was stopped"})
}

func stopChannel(channelId int)  (int, string)  {
	channel, channelExists := channels[channelId]
	if (!channelExists) {
		return 1, "Channel does not exist"
	}

	DestroyChannel(channel)

	delete(channels, channelId)
	return 0, ""
}

func listChannelsRequest(w http.ResponseWriter, r *http.Request) {
	var channelsInfo []Channel
	for _, channel := range channels {
		channelsInfo = append(channelsInfo, channel)
	}
	openapi.ResponseSuccessWithDetails(w, channelsInfo)
}
