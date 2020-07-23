package main

import (
    "fmt"
    "net/http"
    "flag"
    "sync"
    "io/ioutil"

    "github.com/gorilla/mux"
    "github.com/google/uuid"

    "application/archive/record"
    "application/core/logger"
)

var(
    archive         map[string] archiveItem
    sortedArchive   []string
    archiveMutex    sync.RWMutex

    flagArchiveRawPath *string
)

type archiveItem struct {
    record  *record.Record
    busy    bool
}

func init() {
    flagArchiveRawPath     = flag.String("archive-raw-path",     "/opt/usb",              "Raw archive dir path")
}

func initArchive() {
    archive = make(map[string] archiveItem)

    scanArchive()
}

func scanArchive() {
    logger.Log.Debug().
        Msg("Scanning archive")

    files, err := ioutil.ReadDir(*flagArchiveRawPath)
    if err != nil {
        logger.Log.Fatal().
            Str("reson", err.Error()).
            Msg("Can`t scan dir")
    }
    for _, f := range files {
        if f.IsDir() {
            _, err := uuid.Parse(f.Name())
            if err == nil {
                logger.Log.Debug().
                    Str("name", f.Name()).
                    Msg("Dir found")
                rec, err := record.Load(*flagArchiveRawPath, f.Name())
                if err != nil {
                    logger.Log.Warn().
                        Str("name", f.Name()).
                        Str("reason", err.Error()).
                        Msg("Can`t load")
                } else {
                    var item archiveItem
                    item.record = rec
                    archive[f.Name()] = item
                    logger.Log.Trace().
                        Str("name", rec.Name).
                        Str("dir", rec.Dir).
                        Bool("preview", rec.Preview).
                        Msg("Loaded record")
                }
            } else {
                logger.Log.Warn().
                    Str("details", err.Error()).Msg("Dir found, but seems is not record")
            }
        }
    }

    logger.Log.Debug().
        Msg("Scanning archive done")
}

//List all known records
func archiveList(w http.ResponseWriter, r *http.Request) {
    archiveMutex.RLock()
    defer archiveMutex.RUnlock()

    fmt.Fprintf(w, "<h1>List:</h1><ul>")
    for name, item := range(archive) {
        fmt.Fprintf(w, "<li><a href='/archive/%s/download.h264'>%s</a>", name, name)
        if item.record.Preview {
            fmt.Fprintf(w, "<img width='320' height='180' src='/archive/%s/preview.jpeg'>", name)
        }
        fmt.Fprintf(w, "FirstPts %d, ", item.record.FirstPts)
        fmt.Fprintf(w, "LastPts %d, ", item.record.LastPts)
        fmt.Fprintf(w, "FrameCount %d, ", item.record.FrameCount)
        fmt.Fprintf(w, "Period %d, ", (item.record.LastPts-item.record.FirstPts)/item.record.FrameCount)
        if len(item.record.Chunks) > 0 {
            fmt.Fprintf(w, "<br />Size per hour %d MB", (item.record.Chunks[0].Size / item.record.FrameCount) * 25* 60 *60 / (1024*1024))
            fmt.Fprintf(w, "<br />Size per minute %d kB", (item.record.Chunks[0].Size / item.record.FrameCount) * 25* 60 / 1024)
            fmt.Fprintf(w, "<br />Size per second %d kB", (item.record.Chunks[0].Size / item.record.FrameCount) * 25 / 1024)
        }
        fmt.Fprintf(w, "</li>")
    }
    fmt.Fprintf(w, "</ul>")
}

//Show record information
func archiveItemInfo(w http.ResponseWriter, r *http.Request) {
    archiveMutex.RLock()
    defer archiveMutex.RUnlock()

    fmt.Fprintf(w, "archiveItemInfo")
}

//Show record preview
func archiveItemPreview(w http.ResponseWriter, r *http.Request) {
    archiveMutex.RLock()
    defer archiveMutex.RUnlock()

    queryParams := mux.Vars(r)

    item, exist := archive[queryParams["uuid"]]
    if !exist {
        fmt.Fprintf(w, "NotFound")
        return
    }

    rec := item.record

    if !rec.Preview {
        fmt.Fprintf(w, "None")
        return
    }

    http.ServeFile(w, r, rec.Dir+"/"+rec.Name+"/preview.jpeg")
}

//Download record
func archiveItemServe(w http.ResponseWriter, r *http.Request) {
    queryParams := mux.Vars(r)

    archiveMutex.RLock()
        item, exist := archive[queryParams["uuid"]]
        if !exist {
            fmt.Fprintf(w, "Not found")
            return
        }
        item.busy = true
    archiveMutex.RUnlock()
    
    rec := item.record

    if len(rec.Chunks) > 0 {
        http.ServeFile(w, r, rec.Dir+"/"+rec.Name+"/1.h264")
    } else {
        fmt.Fprintf(w, "Not chunks")
    }

    //Start serving (combined chunks)

    //fmt.Fprintf(w, "archiveServe")

    archiveMutex.RLock()
        item.busy = false
    archiveMutex.RUnlock()
}

//Delete record
func archiveDelete(w http.ResponseWriter, r *http.Request) {
    archiveMutex.Lock()
    defer archiveMutex.Unlock()

    //if arcive item is busy it can`t be deleted

    fmt.Fprintf(w, "archiveDelete")
}
