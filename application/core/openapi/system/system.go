package system

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type apiAnswerSystemDateTimeSchema struct {
	Formatted time.Time `json:"formatted,omitempty"`
	Secs      int64     `json:"secs,omitempty"`
	Nanosecs  int64     `json:"nanosecs,omitempty"`
}

func Date(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var dtSchema apiAnswerSystemDateTimeSchema
	t := time.Now()

	dtSchema.Formatted = t
	dtSchema.Secs = t.Unix()
	dtSchema.Nanosecs = t.UnixNano()

	dtSchemaJson, _ := json.Marshal(dtSchema)
	fmt.Fprintf(w, "%s", string(dtSchemaJson))
}
