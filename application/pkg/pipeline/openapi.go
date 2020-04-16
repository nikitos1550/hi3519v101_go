package pipeline

import (
	"application/pkg/openapi"
	"net/http"
)

type responseRecord struct {
	Message string
}

type pipelineRecord struct {
	EncoderId int
	Message string
}

func init() {
    openapi.AddApiRoute("apiDescription", "/pipeline", "GET", apiDescription)

    openapi.AddApiRoute("createPipeline", "/pipeline/create", "GET", createPipelineRequest)
}

func apiDescription(w http.ResponseWriter, r *http.Request)  {
	openapi.ApiDescription(w, r, "Pipeline api:\n\n", "/pipeline")
}

func createPipelineRequest(w http.ResponseWriter, r *http.Request)  {
	ok, encoderName := openapi.GetStringParameter(w, r, "encoderName")
	if !ok {
		return
	}

	id, errorString := CreatePipeline(encoderName)
	if id < 0 {
		openapi.ResponseErrorWithDetails(w, http.StatusInternalServerError, responseRecord{Message: errorString})
		return
	}

	openapi.ResponseSuccessWithDetails(w, pipelineRecord{EncoderId: id, Message: "Pipeline was created"})
}
