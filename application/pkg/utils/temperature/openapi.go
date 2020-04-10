// +build openapi

package temperature

import (
	"application/pkg/openapi"
	"encoding/json"
	"fmt"
	//"log"
	"net/http"
	"application/pkg/logger"
)

func init() {
	openapi.AddApiRoute("serveTemperature", "/temperature", "GET", serveTemperature)
}

type serveTemperatureSchema struct {
	Temperature float32 `json:"temperature"`
}

/**
 * @api {get} /temperature Get temperature
 * @apiName GetTemperature
 * @apiGroup Common
 *
 * @apiDescription SoC has internal temperature sensor.
 *  Once a second value is captured.
 */
func serveTemperature(w http.ResponseWriter, r *http.Request) {
	//log.Println("temperature.serveTemperature")
	logger.Log.Trace().Msg("temperature.serveTemperature")

	var schema serveTemperatureSchema

	var err error
	schema.Temperature, err = Get()

	//log.Println(schema)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if err == nil {
		schemaJson, _ := json.Marshal(schema)
		//log.Println(string(schemaJson))
		fmt.Fprintf(w, "%s", string(schemaJson))

	} else {
		w.WriteHeader(http.StatusNotImplemented) //TODO correct http return code
	}
}
