//+build nobuild
//+build processing
//+build openapi

package processings

import (
	"application/pkg/mpp/vpss"
	"application/pkg/openapi"
	"net/http"
)

type responseRecord struct {
	Message string
}

type processingRecord struct {
	Id int
	Message string
}

type processingInfo struct {
	Name string
}

type activeProcessingInfo struct {
	Id int
	Name string
	InputChannel int
	InputProcessing int
	OutputEncoders []int
	OutputProcessings []int
}

func init() {
    openapi.AddApiRoute("apiDescription", "/processing", "GET", apiDescription)

    openapi.AddApiRoute("createProcessing", "/processing/create", "GET", createProcessingRequest)
    openapi.AddApiRoute("deleteProcessing", "/processing/delete", "GET", deleteProcessingRequest)

    openapi.AddApiRoute("subscribeChannel", "/processing/subscribeChannel", "GET", subscribeChannelRequest)
    openapi.AddApiRoute("unsubscribeChannel", "/processing/unsubscribeChannel", "GET", unsubscribeChannelRequest)

    //openapi.AddApiRoute("subscribeProcessing", "/processing/subscribeProcessing", "GET", subscribeProcessingRequest)
    //openapi.AddApiRoute("unsubscribeProcessing", "/processing/unsubscribeProcessing", "GET", unsubscribeProcessingRequest)

    openapi.AddApiRoute("listProcessing", "/processing/list", "GET", listProcessingRequest)
    openapi.AddApiRoute("listActiveProcessing", "/processing/listActive", "GET", listActiveProcessingRequest)
}

func apiDescription(w http.ResponseWriter, r *http.Request)  {
	openapi.ApiDescription(w, r, "Processings api:\n\n", "/processing")
}

func createProcessingRequest(w http.ResponseWriter, r *http.Request)  {
	ok, processingName := openapi.GetStringParameter(w, r, "processingName")
	if !ok {
		return
	}

	id, errorString := CreateProcessing(processingName, r.URL.Query())
	if id < 0 {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{Message: errorString})
		return
	}

	openapi.ResponseSuccessWithDetails(w, processingRecord{Id: id, Message: "Processing was created"})
}

func deleteProcessingRequest(w http.ResponseWriter, r *http.Request)  {
	ok, processingId := openapi.GetIntParameter(w, r, "processingId")
	if !ok {
		return
	}

	id, errorString := DeleteProcessing(processingId)
	if id < 0 {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{Message: errorString})
		return
	}

	openapi.ResponseSuccessWithDetails(w, processingRecord{Id: id, Message: "Processing was deleted"})
}

func subscribeChannelRequest(w http.ResponseWriter, r *http.Request)  {
	ok, processingId := openapi.GetIntParameter(w, r, "processingId")
	if !ok {
		return
	}

	ok, channelId := openapi.GetIntParameter(w, r, "channelId")
	if !ok {
		return
	}

	activeProcessing, processingExists := ActiveProcessings[processingId]
	if (!processingExists) {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{Message: "Processing not created"})
		return
	}

	err, errorString := vpss.SubscribeChannel(channelId, activeProcessing.Proc)
	if err != 0 {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{Message: errorString})
		return
	}

	activeProcessing.InputChannel = channelId
	ActiveProcessings[processingId] = activeProcessing

	openapi.ResponseSuccessWithDetails(w, responseRecord{Message: "Channel was subscribed"})
}

func unsubscribeChannelRequest(w http.ResponseWriter, r *http.Request)  {
	ok, processingId := openapi.GetIntParameter(w, r, "processingId")
	if !ok {
		return
	}

	ok, channelId := openapi.GetIntParameter(w, r, "channelId")
	if !ok {
		return
	}

	activeProcessing, processingExists := ActiveProcessings[processingId]
	if (!processingExists) {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{Message: "Processing not created"})
		return
	}

	err, errorString := vpss.UnsubscribeChannel(channelId, activeProcessing.Proc)
	if err != 0 {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{Message: errorString})
		return
	}

	activeProcessing.InputChannel = -1
	ActiveProcessings[processingId] = activeProcessing

	openapi.ResponseSuccessWithDetails(w, responseRecord{Message: "Channel was unsubscribed"})
}

func subscribeProcessingRequest(w http.ResponseWriter, r *http.Request)  {
	ok, processingId := openapi.GetIntParameter(w, r, "processingId")
	if !ok {
		return
	}

	ok, subscribeProcessingId := openapi.GetIntParameter(w, r, "subscribeProcessingId")
	if !ok {
		return
	}

	err, errorString := SubscribeProcessingToProcessing(processingId, subscribeProcessingId)
	if err < 0 {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{Message: errorString})
		return
	}

	openapi.ResponseSuccessWithDetails(w, responseRecord{Message: "Processing was subscribed"})
}

func unsubscribeProcessingRequest(w http.ResponseWriter, r *http.Request)  {
	ok, processingId := openapi.GetIntParameter(w, r, "processingId")
	if !ok {
		return
	}

	ok, subscribeProcessingId := openapi.GetIntParameter(w, r, "subscribeProcessingId")
	if !ok {
		return
	}

	err, errorString := UnsubscribeProcessingToProcessing(processingId, subscribeProcessingId)
	if err < 0 {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{Message: errorString})
		return
	}

	openapi.ResponseSuccessWithDetails(w, responseRecord{Message: "Processing was unsubscribed"})
}

func listProcessingRequest(w http.ResponseWriter, r *http.Request)  {
	var processingsInfo []processingInfo
	for name, _ := range Processings {
		info := processingInfo{
			Name: name,
		}
		processingsInfo = append(processingsInfo, info)
	}
	openapi.ResponseSuccessWithDetails(w, processingsInfo)
}

func listActiveProcessingRequest(w http.ResponseWriter, r *http.Request)  {
	var infos []activeProcessingInfo
	for id, processing := range ActiveProcessings {
		info := activeProcessingInfo{
			Id: id,
			Name: processing.Proc.GetName(),
			InputChannel: processing.InputChannel,
			InputProcessing: processing.InputProcessing,
		}

		for encoderId, _ := range processing.Encoders {
			info.OutputEncoders = append(info.OutputEncoders, encoderId)
		}

		for processingId, _ := range processing.Processings {
			info.OutputProcessings = append(info.OutputProcessings, processingId)
		}

		infos = append(infos, info)
	}

	openapi.ResponseSuccessWithDetails(w, infos)
}
