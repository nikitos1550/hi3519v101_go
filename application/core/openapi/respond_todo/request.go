//+build nobuild

package respond

import (
	"net/http"
	"net/url"
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

func GetStringParameterOrDefault(w http.ResponseWriter, r *http.Request, name string, defaultValue string) (string) {
	value, ok := r.URL.Query()[name]
	if !ok {
		return defaultValue
	}

	return value[0]
}

func GetIntParameterOrDefault(w http.ResponseWriter, r *http.Request, name string, defaultValue int) (int) {
	value, ok := r.URL.Query()[name]
	if !ok {
		return defaultValue
	}

	num, err := strconv.Atoi(value[0])
	if err != nil {
		return defaultValue
	}

	return num
}

func PostStringParameter(w http.ResponseWriter, body string, name string) (bool, string) {
	values, err := url.ParseQuery(body)
	if err != nil {
		ResponseErrorWithDetails(w, http.StatusBadRequest, badParameter{Message: "Failed to parse request body. Required parameter '" + name + "' is missed"})
		return false, ""
	}

	value, ok := values[name]
	if !ok {
		ResponseErrorWithDetails(w, http.StatusBadRequest, badParameter{Message: "Required parameter '" + name + "' is missed"})
		return false, ""
	}

	return true, value[0]
}

func PostIntParameter(w http.ResponseWriter, body string, name string) (bool, int) {
	ok, stringValue := PostStringParameter(w, body, name)
	if !ok {
		return false, 0
	}

	num, err := strconv.Atoi(stringValue)
	if err != nil {
		ResponseErrorWithDetails(w, http.StatusBadRequest, badParameter{Message: "Parameter '" + name + "' should be int"})
		return false, 0
	}

	return true, num
}
