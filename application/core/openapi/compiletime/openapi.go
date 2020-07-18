package compiletime

import (
	"fmt"
    "net/http"
	"encoding/json"

    "application/core/compiletime"
)

func Serve(w http.ResponseWriter, r *http.Request) {
	var schema compiletime.Info
	compiletime.CopyAll(&schema)

	w.WriteHeader(http.StatusOK)

	schemaJson, _ := json.Marshal(schema)
    fmt.Fprintf(w, "%s", string(schemaJson))
}
