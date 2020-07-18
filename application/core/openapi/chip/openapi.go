package chip

import (
	"fmt"
	"net/http"
	"encoding/json"

    "application/core/utils/chip"
)

type InfoSchema struct {
	RegId		uint32	`json:"regchipid"`
	DetectReg	string	`json:"detectedchipreg"`
}

func Info(w http.ResponseWriter, r *http.Request) {
	var schema InfoSchema

	w.WriteHeader(http.StatusOK)

	schema.RegId = chip.RegId()
	schema.DetectReg = chip.Detect(schema.RegId)

    schemaJson, _ := json.Marshal(schema)
    fmt.Fprintf(w, "%s", string(schemaJson))
}
