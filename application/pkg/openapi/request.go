// +build openapi

package openapi

import (
	"net/http"
	"strconv"    
)

type badParameter struct {
	Message string
}

func init() {
}

func GetStringParameter(w http.ResponseWriter, r *http.Request, name string) (bool, string) {
	value, ok := r.URL.Query()[name]
	if !ok {
		ResponseErrorWithDetails(w, http.StatusBadRequest, badParameter{Message: "Required parameter '" + name + "' is missed"})
		return false, ""
	}

	return true, value[0]
}

func GetIntParameter(w http.ResponseWriter, r *http.Request, name string) (bool, int) {
	value, ok := r.URL.Query()[name]
	if !ok {
		ResponseErrorWithDetails(w, http.StatusBadRequest, badParameter{Message: "Required parameter '" + name + "' is missed"})
		return false, 0
	}

	num, err := strconv.Atoi(value[0])
	if err != nil {
		ResponseErrorWithDetails(w, http.StatusBadRequest, badParameter{Message: "Parameter '" + name + "' should be int"})
		return false, 0
	}

	return true, num
}
