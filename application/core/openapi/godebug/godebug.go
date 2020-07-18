package godebug

import (
    "expvar"
    "net/http"
    "fmt"
)

//https://habr.com/ru/post/257593/
func Expvar(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "{\n")
    first := true
    expvar.Do(func(kv expvar.KeyValue) {
        if !first {
            fmt.Fprintf(w, ",\n")
        }
        first = false
        fmt.Fprintf(w, "%q: %s", kv.Key, kv.Value)
    })
    fmt.Fprintf(w, "\n}\n")
}
