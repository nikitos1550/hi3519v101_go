//+build openapi

package venc

import (
        "net/http"

        "application/pkg/openapi"
        "application/pkg/processing"
)

type responseRecord struct {
	Message string
}

type encoderRecord struct {
	Id int
	Message string
}

type predefinedEncoderInfo struct {
    Name string 
    Format string 
    Width int 
    Height int 
    Bitrate int 
}

type activeEncoderInfo struct {
    EncoderId int 
	ProcessingId int 
    Format string 
    Width int 
    Height int 
    Bitrate int 
}

func init() {
    openapi.AddApiRoute("apiDescription", "/encoder", "GET", apiDescription)

    openapi.AddApiRoute("createEncoderRequest", "/encoder/create", "GET", createEncoderRequest)
    openapi.AddApiRoute("deleteEncoderRequest", "/encoder/delete", "GET", deleteEncoderRequest)

    openapi.AddApiRoute("subscribeProcessingRequest", "/encoder/subscribeProcessing", "GET", subscribeProcessingRequest)
    openapi.AddApiRoute("unsubscribeProcessingRequest", "/encoder/unsubscribeProcessing", "GET", unsubscribeProcessingRequest)

    openapi.AddApiRoute("listEncodersRequest", "/encoder/list", "GET", listEncodersRequest)

    openapi.AddApiRoute("listPredefinedEncoders", "/encoder/predefined", "GET", listPredefinedEncodersRequest)
}

func listPredefinedEncodersRequest(w http.ResponseWriter, r *http.Request)  {
    var encodersInfo []predefinedEncoderInfo
    for name, encoder := range PredefinedEncoders {
            info := predefinedEncoderInfo{
                    Name: name,
                    Format: encoder.Format,
                    Width: encoder.Width,
                    Height: encoder.Height,
                    Bitrate: encoder.Bitrate,
            }
        
            encodersInfo = append(encodersInfo, info)
    }
    openapi.ResponseSuccessWithDetails(w, encodersInfo)
}

func apiDescription(w http.ResponseWriter, r *http.Request)  {
	openapi.ApiDescription(w, r, "Encoders api:\n\n", "/encoder")
}

func createEncoderRequest(w http.ResponseWriter, r *http.Request)  {
	encoderName := openapi.GetStringParameterOrDefault(w, r, "encoderName", "")
	if encoderName != "" {
		id, errString := CreatePredefinedEncoder(encoderName)
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
	ok, encoderId := openapi.GetIntParameter(w, r, "encoderId")
	if !ok {
		return
	}

	err, errorString := DeleteEncoder(encoderId)
	if err != 0 {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{Message: errorString})
		return
	}

	openapi.ResponseSuccessWithDetails(w, responseRecord{Message: "Encoder was deleted"})
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

	encoder, encoderExists := ActiveEncoders[encoderId]
	if (!encoderExists) {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{Message: "Failed to find encoder"})
		return
	}

	err, errorString := processing.SubscribeEncoderToProcessing(processingId, encoderId)
	if err != 0 {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{Message: errorString})
		return
	}
	encoder.ProcessingId = processingId
	ActiveEncoders[encoderId] = encoder

	openapi.ResponseSuccessWithDetails(w, responseRecord{Message: "Processing was subscribed"})
}

func unsubscribeProcessingRequest(w http.ResponseWriter, r *http.Request)  {
	ok, processingId := openapi.GetIntParameter(w, r, "processingId")
	if !ok {
		return
	}

	ok, encoderId := openapi.GetIntParameter(w, r, "encoderId")
	if !ok {
		return
	}

	encoder, encoderExists := ActiveEncoders[encoderId]
	if (!encoderExists) {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{Message: "Failed to find encoder"})
		return
	}

	err, errorString := processing.UnsubscribeEncoderToProcessing(processingId, encoderId)
	if err < 0 {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{Message: errorString})
		return
	}
	encoder.ProcessingId = -1
	ActiveEncoders[encoderId] = encoder

	openapi.ResponseSuccessWithDetails(w, responseRecord{Message: "Processing was unsubscribed"})
}

func listEncodersRequest(w http.ResponseWriter, r *http.Request)  {
    var encodersInfo []activeEncoderInfo
    for id, encoder := range ActiveEncoders {
            info := activeEncoderInfo{
                    EncoderId: id,
					ProcessingId: encoder.ProcessingId,
                    Format: encoder.Format,
                    Width: encoder.Width,
                    Height: encoder.Height,
                    Bitrate: encoder.Bitrate,
            }
        
            encodersInfo = append(encodersInfo, info)
    }
    openapi.ResponseSuccessWithDetails(w, encodersInfo)
}
