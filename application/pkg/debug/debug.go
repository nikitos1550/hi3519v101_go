//+build debug

package debug

import (
    "expvar"
    "net/http"
    "fmt"
    "log"
    //"time"
    //"os"
    //"io/ioutil"
    //"strings"
    //"strconv"
    "application/pkg/openapi"
)

func init() {
    openapi.AddApiRoute("debugExpvar", "/debug/vars", "GET", debugExpvar)
}

//https://habr.com/ru/post/257593/
func debugExpvar(w http.ResponseWriter, r *http.Request) {
    log.Println("debugExpvar")

    w.Header().Set("Content-Type", "application/json; charset=utf-8")
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


/*
// Custom struct that will be exported
type Load struct {
  Load1  float64
  Load5  float64
  Load15 float64
}

// Function that will be called by expvar
// to export the information from the structure
// every time the endpoint is reached
func AllLoadAvg() interface{} {
  return Load{
     Load1:  loadAvg(0),
     Load5:  loadAvg(1),
     Load15: loadAvg(2),
  }
}

// Aux function to retrieve the load average
// in GNU/Linux systems
func loadAvg(position int) float64 {
  data, err := ioutil.ReadFile("/proc/loadavg")
  if err != nil {
     panic(err)
  }
  values := strings.Fields(string(data))

  load, err := strconv.ParseFloat(values[position], 64)
  if err != nil {
     panic(err)
  }

  return load
}

*/
