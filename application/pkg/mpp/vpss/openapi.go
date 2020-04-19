//+build openapi

package vpss

import (
	"application/pkg/openapi"
	"net/http"
	"unsafe"
)

type ChannelInfo struct {
	ChannelId  int
	Width int
	Height int
	Fps int
	CropX int
	CropY int
	CropWidth int
	CropHeight int
	Processings []int
}

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

	channel.Started = true                              //Should be done only after successfull channel start
	channel.Clients = make(map[int] unsafe.Pointer)     //Most probably same issue as above

	err, errorString := StartChannel(channel)
	if err < 0 {
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
	if err < 0 {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{Message: errorString})
		return
	}

	openapi.ResponseSuccessWithDetails(w, responseRecord{Message: "Channel was stopped"})
}

func listChannelsRequest(w http.ResponseWriter, r *http.Request) {
	var channelsInfo []ChannelInfo
	for _, channel := range Channels {
		info := ChannelInfo{
			ChannelId: channel.ChannelId,
			Width: channel.Width,
			Height: channel.Height,
			Fps: channel.Fps,
			CropX: channel.CropX,
			CropY: channel.CropY,
			CropWidth: channel.CropWidth,
			CropHeight: channel.CropHeight,
		}
		for processingId, _ := range channel.Clients {
			info.Processings = append(info.Processings, processingId)
		}

		channelsInfo = append(channelsInfo, info)
	}
	openapi.ResponseSuccessWithDetails(w, channelsInfo)
}
