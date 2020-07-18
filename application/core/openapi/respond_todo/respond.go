package respond

import (
    "fmt"
    "net/http"
)

func RespondWithError(w http.ResponseWriter, err error) {
    w.WriteHeader(http.StatusInternalServerError)
    fmt.Fprintf(w, "{\"error\":\"%s\"}", err.Error())
}
