package openapi

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "runtime"
    "encoding/json"
    "log"
    "github.com/gorilla/mux"
)

type debugGoResponce struct {
    Alloc,
    TotalAlloc,
    Sys,
    Mallocs,
    Frees,
    LiveObjects,
    PauseTotalNs uint64
    NumGC        uint32
    NumGoroutine int
}

func debugGo (w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    var m debugGoResponce
    var rtm runtime.MemStats

    // Read full mem stats
    runtime.ReadMemStats(&rtm)

    // Number of goroutines
    m.NumGoroutine = runtime.NumGoroutine()

    // Misc memory stats
    m.Alloc = rtm.Alloc
    m.TotalAlloc = rtm.TotalAlloc
    m.Sys = rtm.Sys
    m.Mallocs = rtm.Mallocs
    m.Frees = rtm.Frees

    // Live objects = Mallocs - Frees
    m.LiveObjects = m.Mallocs - m.Frees

    // GC Stats
    m.PauseTotalNs = rtm.PauseTotalNs
    m.NumGC = rtm.NumGC

    log.Println("debugGo")
    test, _ := json.Marshal(m)
    fmt.Fprintf(w, "%s", string(test))
}

func debugUmap (w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    files, err := ioutil.ReadDir("/proc/umap")
    if err != nil {
        //TODO /proc/umap exist only after ko modules init, handle it smart!
        panic(err)
        return
    }

    num := len(files)
    var i int = 0

    for i < (num-1) {
            //fmt.Println(f.Name())
            fmt.Fprintf(w, files[i].Name())
            fmt.Fprintf(w, ",")
            i++
    }
    fmt.Fprintf(w, files[num-1].Name())
}


func debugUmapFile (w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)

    log.Println(r.URL.Path)

    dat, err := ioutil.ReadFile("/proc/umap/" + params["file"])
    if err != nil {
        w.WriteHeader(http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    fmt.Fprintf(w, string(dat))
}

