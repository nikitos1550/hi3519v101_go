// +build openapi

package temperature

import (
	"log"
	"fmt"
	"net/http"
	"application/pkg/openapi"
	"encoding/json"
)

func init() {
	openapi.AddApiRoute("serveTemperature", "/temperature", "GET", serveTemperature)
}

type serveTemperatureSchema struct {
	Temperature	float32 `json:"temperature"`
}

func serveTemperature(w http.ResponseWriter, r *http.Request) {
	log.Println("temperature.serveTemperature")

	var schema serveTemperatureSchema

	var err error
    schema.Temperature, err = Get()
	
	//log.Println(schema)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

    if (err == nil) {
		schemaJson, _ := json.Marshal(schema)
		//log.Println(string(schemaJson))
		fmt.Fprintf(w, "%s", string(schemaJson))	

    } else {
		w.WriteHeader(http.StatusNotFound)//TODO correct http return code
    }
}
