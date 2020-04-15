package venc

import (
	"application/pkg/processing"
	"application/pkg/openapi"
	"net/http"
)

type responseRecord struct {
	Message string
}

type encoderRecord struct {
	Id int
	Message string
}

func init() {
    openapi.AddApiRoute("apiDescription", "/encoder", "GET", apiDescription)

    openapi.AddApiRoute("createEncoderRequest", "/encoder/create", "GET", createEncoderRequest)
    openapi.AddApiRoute("deleteEncoderRequest", "/encoder/delete", "GET", deleteEncoderRequest)

    openapi.AddApiRoute("subscribeProcessingRequest", "/encoder/subscribeProcessing", "GET", subscribeProcessingRequest)
    openapi.AddApiRoute("unsubscribeProcessingRequest", "/encoder/unsubscribeProcessing", "GET", unsubscribeProcessingRequest)

    openapi.AddApiRoute("listEncodersRequest", "/encoder/list", "GET", listEncodersRequest)
}

func apiDescription(w http.ResponseWriter, r *http.Request)  {
	openapi.ApiDescription(w, r, "Encoders api:\n\n", "/encoder")
}

func createEncoderRequest(w http.ResponseWriter, r *http.Request)  {
	encoderName := openapi.GetStringParameterOrDefault(w, r, "encoderName", "")
	if encoderName != "" {
		id, errString := createPredefinedEncoder(encoderName)
		if (id < 0){
			openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{Message: errString})
			return
		}

		openapi.ResponseSuccessWithDetails(w, encoderRecord{Id: id, Message: "Encoder was created"})
		return
	}

	openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{Message: "Not supported"})
}

func deleteEncoderRequest(w http.ResponseWriter, r *http.Request)  {
}

func subscribeProcessingRequest(w http.ResponseWriter, r *http.Request)  {
	ok, processingId := openapi.GetIntParameter(w, r, "processingId")
	if !ok {
		return
	}

	ok, encoderId := openapi.GetIntParameter(w, r, "encoderId")
	if !ok {
		return
	}

	err, errorString := processing.SubscribeProcessing(processingId, encoderId)
	if err != 0 {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{Message: errorString})
		return
	}

	openapi.ResponseSuccessWithDetails(w, responseRecord{Message: "Processing was subscribed"})
}

func unsubscribeProcessingRequest(w http.ResponseWriter, r *http.Request)  {
}

func listEncodersRequest(w http.ResponseWriter, r *http.Request)  {
}

func createPredefinedEncoder(encoderName string) (int, string)  {
	encoder, encoderExists := Encoders[encoderName]
	if (!encoderExists) {
		return -1, "Failed to find encoder  " + encoderName
	}

	return encoder.VencId, ""
}
