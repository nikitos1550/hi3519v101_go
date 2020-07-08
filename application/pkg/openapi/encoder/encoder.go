//+build openapi

package encoder

import (
    "fmt"
    "net/http"
    "strconv"
    "encoding/json"

    "github.com/gorilla/mux"

    "application/pkg/openapi"
    "application/pkg/logger"

    "application/pkg/mpp/venc"
)

func init() {
    /*
    openapi.AddApiRoute("apiDescription", "/encoder", "GET", apiDescription)

    openapi.AddApiRoute("createEncoderRequest", "/encoder/create", "GET", createEncoderRequest)
    openapi.AddApiRoute("createDummyEncoderRequest", "/encoder/create_dummy", "GET", createDummyEncoderRequest)
    openapi.AddApiRoute("deleteEncoderRequest", "/encoder/delete", "GET", deleteEncoderRequest)

    openapi.AddApiRoute("subscribeProcessingRequest", "/encoder/subscribeProcessing", "GET", subscribeProcessingRequest)
    openapi.AddApiRoute("unsubscribeProcessingRequest", "/encoder/unsubscribeProcessing", "GET", unsubscribeProcessingRequest)

    openapi.AddApiRoute("listEncodersRequest", "/encoder/list", "GET", listEncodersRequest)

    openapi.AddApiRoute("listPredefinedEncoders", "/encoder/predefined", "GET", listPredefinedEncodersRequest)
    */
    ////////////////////

    openapi.AddApiRoute("encodersList", "/mpp/encoders", "GET", encodersList)
    openapi.AddApiRoute("encoderCreate", "/mpp/encoders", "POST", encoderCreateNew)
    openapi.AddApiRoute("encoderInfo", "/mpp/encoders/{id:[0-9]+}", "GET", encoderInfo)
    openapi.AddApiRoute("encoderCreate", "/mpp/encoders/{id:[0-9]+}", "POST", encoderCreate)
    openapi.AddApiRoute("encoderUpdate", "/mpp/encoders/{id:[0-9]+}", "PUT", encoderUpdate)
    openapi.AddApiRoute("encoderDestroy", "/mpp/encoders/{id:[0-9]+}", "DELETE", encoderDestroy)

    openapi.AddApiRoute("encoderStart", "/mpp/encoders/{id:[0-9]+}/start", "GET", encoderStart)
    openapi.AddApiRoute("encoderStop", "/mpp/encoders/{id:[0-9]+}/stop", "GET", encoderStop)

    openapi.AddApiRoute("encoderIdr", "/mpp/encoders/{id:[0-9]+}/idr", "GET", encoderIdr)
}

func encoderIdr(w http.ResponseWriter, r *http.Request)  {
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

    e, err := venc.GetEncoder(id)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    started, err := e.IsStarted()
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    if !started {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"Encoder is not started\"}")
        return
    }

    err = e.RequestIFrame()
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    w.WriteHeader(http.StatusOK)
}

func encoderStart(w http.ResponseWriter, r *http.Request)  {
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

    e, err := venc.GetEncoder(id)
    if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
		return
    }

    started, err := e.IsStarted()
    if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
		return
    }

    if started {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"Encoder is already started\"}")
        return
    }

    err = e.Start()//TODO
    if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
		return
    }

    w.WriteHeader(http.StatusOK)
}

func encoderStop(w http.ResponseWriter, r *http.Request)  {
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

    e, err := venc.GetEncoder(id)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    started, err := e.IsStarted()
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
		return
    }

    if !started {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"Encoder is not started\"}")
        return
    }

    err = e.Stop()
    if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
		return
    }

    w.WriteHeader(http.StatusOK)
}

func encoderCreateNew(w http.ResponseWriter, r *http.Request)  {        //TODO non atomic (in terms of lock) search and create operation,
    w.Header().Set("Content-Type", "text/plain; charset=UTF-8")         //new if search and create should goes under one lock,
                                                                        //maybe some global lock for all channels array
    var err error

    e, err := venc.GetFirstEmpty()
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"No more encoders can be created\"}")
        return
    }

    var params venc.Parameters

    venc.InvalidateBitrateControlParameters(&params.BitControlParams)
    venc.InvalidateGopParameters(&params.GopParams)

    err = json.NewDecoder(r.Body).Decode(&params)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    err = e.CreateEncoder(params, false)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    logger.Log.Trace().
        Msg("Encoder created")

    newParams, err := e.GetParams()
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    w.WriteHeader(http.StatusOK)

    schemaJson, _ := json.MarshalIndent(newParams, "", "\t")
    fmt.Fprintf(w, "%s", string(schemaJson))

}

type encodersListResponce struct {
    Amount      int         `json:"maxamount"`
    Encoders    []venc.Encoder   `json:"encoders"`
}

func encodersList(w http.ResponseWriter, r *http.Request)  {
    w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

    var response encodersListResponce

    response.Amount = venc.EncodersAmount

    for i:=0; i < venc.EncodersAmount; i++ {
        e, _ := venc.GetEncoder(i)
        cpy, err := e.GetCopy()
        if err != nil {
            response.Encoders = append(response.Encoders, cpy)
        }
    }

    w.WriteHeader(http.StatusOK)

    schemaJson, _ := json.MarshalIndent(response, "", "\t")
    fmt.Fprintf(w, "%s", string(schemaJson))

}

func encoderCreate(w http.ResponseWriter, r *http.Request)  {
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

    var params venc.Parameters

    venc.InvalidateBitrateControlParameters(&params.BitControlParams)
    venc.InvalidateGopParameters(&params.GopParams)

    err = json.NewDecoder(r.Body).Decode(&params)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    //logger.Log.Trace().
    //    Msg("body decoded")

    e, err := venc.GetEncoder(id)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    err = e.CreateEncoder(params, false)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    logger.Log.Trace().
        Msg("Encoder created")

    newParams, err := e.GetParams()
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    w.WriteHeader(http.StatusOK)

    schemaJson, _ := json.MarshalIndent(newParams, "", "\t")
    fmt.Fprintf(w, "%s", string(schemaJson))
}

func encoderUpdate(w http.ResponseWriter, r *http.Request)  {
    //TODO
}

func encoderDestroy(w http.ResponseWriter, r *http.Request)  {
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

    e, err := venc.GetEncoder(id)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    err = e.DestroyEncoder()
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    w.WriteHeader(http.StatusOK)
}

func encoderInfo(w http.ResponseWriter, r *http.Request)  {
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

    e, err := venc.GetEncoder(id)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    info, err := e.GetCopy()
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
        return
    }

    w.WriteHeader(http.StatusOK)

    schemaJson, _ := json.MarshalIndent(info, "", "\t")
    fmt.Fprintf(w, "%s", string(schemaJson))
}

///////////////////////////////////////

/*
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

*/
