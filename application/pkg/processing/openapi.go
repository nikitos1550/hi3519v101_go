//+build processing

package processing

import (
	"application/pkg/mpp/vpss"
	"application/pkg/openapi"
	"net/http"
)

type responseRecord struct {
	Message string
}

type processing struct {
	Name string
}

type activeProcessing struct {
	Name string
}

func init() {
    openapi.AddApiRoute("apiDescription", "/processing", "GET", apiDescription)

    openapi.AddApiRoute("listChannels", "/processing/subscribeChannel", "GET", subscribeChannelRequest)
    openapi.AddApiRoute("listChannels", "/processing/unsubscribeChannel", "GET", unsubscribeChannelRequest)
    openapi.AddApiRoute("listChannels", "/processing/list", "GET", listProcessingRequest)
    openapi.AddApiRoute("listChannels", "/processing/listActive", "GET", listActiveProcessingRequest)
}

func apiDescription(w http.ResponseWriter, r *http.Request)  {
	openapi.ApiDescription(w, r, "Processings api:\n\n", "/processing")
}

func subscribeChannelRequest(w http.ResponseWriter, r *http.Request)  {
	ok, processingName := openapi.GetStringParameter(w, r, "processingName")
	if !ok {
		return
	}

	ok, channelId := openapi.GetIntParameter(w, r, "channelId")
	if !ok {
		return
	}

	callback, exists := Processings[processingName]
	if (!exists) {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{Message: "Processing not found"})
		return
	}

	_, exists = vpss.Channels[channelId]
	if (!exists) {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{Message: "Channel not started"})
		return
	}

	err, errorString := vpss.SubscribeChannel(channelId, callback)
	if err != 0 {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{Message: errorString})
		return
	}

	openapi.ResponseSuccessWithDetails(w, responseRecord{Message: "Channel was subscribed"})
}

func unsubscribeChannelRequest(w http.ResponseWriter, r *http.Request)  {
}

func listProcessingRequest(w http.ResponseWriter, r *http.Request)  {
	var processingInfo []processing
	for name, _ := range Processings {
		info := processing{
			Name: name,
		}
		processingInfo = append(processingInfo, info)
	}
	openapi.ResponseSuccessWithDetails(w, processingInfo)
}

func listActiveProcessingRequest(w http.ResponseWriter, r *http.Request)  {
}
