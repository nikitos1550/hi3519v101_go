package temperature

import (
	"encoding/json"
	"fmt"
	"net/http"

    "application/core/utils/temperature"
)

type serveTemperatureSchema struct {
	Temperature float32 `json:"temperature"`
}

func Serve(w http.ResponseWriter, r *http.Request) {
	var schema serveTemperatureSchema

	var err error
	schema.Temperature, err = temperature.Get()

	if err == nil {
		schemaJson, _ := json.Marshal(schema)
		fmt.Fprintf(w, "%s", string(schemaJson))

	} else {
		w.WriteHeader(http.StatusNotImplemented) //TODO correct http return code
	}
}
