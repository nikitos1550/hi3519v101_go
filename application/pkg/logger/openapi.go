//+build ignore

package logger

import (
	"flag"
	"io"
	"io/ioutil"
	"net/http"
	"application/pkg/openapi"
)

var (
	apiBuf	apiBuffer
	apiWriter io.Writer
	size	uint

	outputApi bool
)

func init() {
        openapi.AddApiRoute("loggerTest", "/logger/test", "GET", loggerTest)

	flag.UintVar(&size, "logger-api-size", 1024, "Amount of stored messages")
	flag.BoolVar(&outputApi, "logger-output-api", false, "Expose http API for logger")
}

func initApiLog() {
        if outputApi == true {
                apiWriter = &apiBuf
        } else {
                apiWriter = ioutil.Discard
        }

        Log.Debug().
                Bool("logger-output-api", outputApi).
                Msg("Logger api output")

}

func loggerTest(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusOK)

	//schemaJson, _ := json.Marshal(schema)
	//fmt.Fprintf(w, "%s", string(schemaJson))
}

type apiBuffer struct {
	buf []string
}

func (a *apiBuffer) Write(p []byte) (n int, err error) {
	return 0, nil
}
