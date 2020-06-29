//+build openapi

package venc

import (
    "fmt"
    "net/http"
    "strconv"
    "encoding/json"

    "github.com/gorilla/mux"

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
    openapi.AddApiRoute("createDummyEncoderRequest", "/encoder/create_dummy", "GET", createDummyEncoderRequest)
    openapi.AddApiRoute("deleteEncoderRequest", "/encoder/delete", "GET", deleteEncoderRequest)

    openapi.AddApiRoute("subscribeProcessingRequest", "/encoder/subscribeProcessing", "GET", subscribeProcessingRequest)
    openapi.AddApiRoute("unsubscribeProcessingRequest", "/encoder/unsubscribeProcessing", "GET", unsubscribeProcessingRequest)

    openapi.AddApiRoute("listEncodersRequest", "/encoder/list", "GET", listEncodersRequest)

    openapi.AddApiRoute("listPredefinedEncoders", "/encoder/predefined", "GET", listPredefinedEncodersRequest)
    ////////////////////
    openapi.AddApiRoute("encoderInfo", "/mpp/encoder/{id:[0-9]+}", "GET", encoderInfo)
    //openapi.AddApiRoute("encoderStat", "/mpp/encoder/{id:[0-9]+}/stat", "GET", encoderStat)

    openapi.AddApiRoute("encoderTest", "/mpp/encoder/{id:[0-9]+}/test", "GET", encoderTest)
}

func encoderTest(w http.ResponseWriter, r *http.Request)  {
    w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

    var err error
    var id int

    queryParams := mux.Vars(r)
    id, err = strconv.Atoi(queryParams["id"])

    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    var params Parameters

    params.Codec                = H264
    params.Profile              = Main
    params.Width                = 3840
    params.Height               = 2160
    params.Fps                  = 1

    params.GopType              = BipredB
    params.Gop                  = 120
    //Gopparams           GopParameters

    params.BitControl          = Vbr

    params.BitControlParams.Bitrate     = 1024

    params.BitControlParams.StatTime    = 1
    params.BitControlParams.Fluctuate   = 1

    params.BitControlParams.QFactor     = 1
    params.BitControlParams.MinQFactor  = 1
    params.BitControlParams.MaxQFactor  = 1

    params.BitControlParams.MinIQp      = 1
    params.BitControlParams.MaxQp       = 1
    params.BitControlParams.MinQp       = 1

    params.BitControlParams.IQp         = 1
    params.BitControlParams.PQp         = 1
    params.BitControlParams.BQp         = 1


    err = CreateEncoder(id, params)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    newParams, err := GetParams(id)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    w.WriteHeader(http.StatusOK)

    schemaJson, _ := json.MarshalIndent(newParams, "", "\t") //Marshal(newParams)
    fmt.Fprintf(w, "%s", string(schemaJson))
}

func encoderInfo(w http.ResponseWriter, r *http.Request)  {
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

func createDummyEncoderRequest(w http.ResponseWriter, r *http.Request)  {
	id, errString := CreateDummyEncoder()
	if (id < 0){
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{Message: errString})
		return
	}

	openapi.ResponseSuccessWithDetails(w, encoderRecord{Id: id, Message: "Encoder was created"})
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

	err, errorString := processing.SubscribeEncoderToProcessing(processingId, encoder)
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

	err, errorString := processing.UnsubscribeEncoderToProcessing(processingId, encoder)
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
