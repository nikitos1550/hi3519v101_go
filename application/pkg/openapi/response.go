// +build openapi

package openapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ResponseBody struct {
	InternalError int
	Message string
	Details interface{}
}

func init() {
}

func response(w http.ResponseWriter, status int, message string, details interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)

	var responseBody ResponseBody
	responseBody.InternalError = 0
	responseBody.Message = message
	responseBody.Details = details

	responseBodyJson, _ := json.MarshalIndent(responseBody, "", "  ")
	fmt.Fprintf(w, "%s", string(responseBodyJson))
}

func ResponseSuccess(w http.ResponseWriter) {
	var emptyDetails interface{}
	ResponseSuccessWithDetails(w, emptyDetails)
}

func ResponseSuccessWithDetails(w http.ResponseWriter, details interface{}) {
	response(w, http.StatusOK, "Success", details)
}

func ResponseError(w http.ResponseWriter, status int) {
	var emptyDetails interface{}
	ResponseErrorWithDetails(w, status, emptyDetails)
}

func ResponseErrorWithDetails(w http.ResponseWriter, status int, details interface{}) {
	response(w, status, "Error", details)
}