package main

import (
    "fmt"
    "net/http"
    "flag"
    "sync"
    "os"
    "io/ioutil"
    "time"

    "github.com/gorilla/mux"
    //"github.com/google/uuid"
    "github.com/satori/go.uuid"

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
    flagArchiveRawPath     = flag.String("archive-raw-path",     "/opt/nfs",              "Raw archive dir path")
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
            //_, err := uuid.Parse(f.Name())
            _, err := uuid.FromString(f.Name())
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

    fmt.Fprintf(w, "<h1>List:</h1><table>")
    for name, item := range(archive) {
        fmt.Fprintf(w, "<tr><td><a href='/archive/%s/player'>%s</a></td>", name, name)
        if item.record.Preview {
            fmt.Fprintf(w, "<td><img width='320' height='180' src='/archive/%s/preview.jpeg'></td>", name)
        }
        if item.record.FrameCount > 0 {
            fmt.Fprintf(w, "<td><ul><li>FirstPts %d, %v</li>", item.record.FirstPts, time.Duration(item.record.FirstPts*1000))
            fmt.Fprintf(w, "<li>LastPts %d</li>", item.record.LastPts)
            fmt.Fprintf(w, "<li>FrameCount %d</li>", item.record.FrameCount)
            fmt.Fprintf(w, "<li>Period %d</li>", (item.record.LastPts-item.record.FirstPts)/item.record.FrameCount)
            if len(item.record.Chunks) > 0 {
                fmt.Fprintf(w, "<li>Size per hour %d MB</li>", (item.record.Chunks[0].Size / item.record.FrameCount) * 25* 60 *60 / (1024*1024))
                fmt.Fprintf(w, "<li>Size per minute %d kB</li>", (item.record.Chunks[0].Size / item.record.FrameCount) * 25* 60 / 1024)
                fmt.Fprintf(w, "<li>Size per second %d kB</li>", (item.record.Chunks[0].Size / item.record.FrameCount) * 25 / 1024)
            }
            fmt.Fprintf(w, "</ul></td>")
        } else {
            fmt.Fprintf(w, "<td>ZERO FRAMES</td>")
        }
        fmt.Fprintf(w, "<td><a href='/archive/%s/delete'>DELETE</a></td></tr>", name)
    }
    fmt.Fprintf(w, "</table>")
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

func archiveItemM3U8(w http.ResponseWriter, r *http.Request) {
    queryParams := mux.Vars(r)

    archiveMutex.RLock()
    defer archiveMutex.RUnlock()

    item, exist := archive[queryParams["uuid"]]
    if !exist {
        fmt.Fprintf(w, "Not found")
        return
    }

    rec := item.record

    fmt.Fprintf(w, "#EXTM3U\n")
    fmt.Fprintf(w, "#EXT-X-PLAYLIST-TYPE:VOD\n")
    fmt.Fprintf(w, "#EXT-X-TARGETDURATION:120\n")
    fmt.Fprintf(w, "#EXT-X-VERSION:4\n")
    fmt.Fprintf(w, "#EXT-X-MEDIA-SEQUENCE:0\n")

    for id, _ := range(rec.Chunks) {
        fmt.Fprintf(w, "#EXTINF:60.0,\n")
        fmt.Fprintf(w, "/archive/%s/%d.ts\n", rec.Name, id)
    }
    //#fmt.Printf("#EXTINF:10.0,")
    //#fmt.Printf("http://example.com/movie1/fileSequenceB.ts")
    //fmt.Printf("#EXTINF:10.0,")
    //fmt.Printf("http://example.com/movie1/fileSequenceC.ts")
    //fmt.Printf("#EXTINF:9.0,")
    //fmt.Printf("http://example.com/movie1/fileSequenceD.ts")
    fmt.Fprintf(w, "#EXT-X-ENDLIST")
}

func archiveItemTs(w http.ResponseWriter, r *http.Request) {
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
        http.ServeFile(w, r, rec.Dir+"/"+rec.Name+"/"+queryParams["chunk"]+".ts")
    } else {
        fmt.Fprintf(w, "Not chunks")
    }

    archiveMutex.RLock()
        item.busy = false
    archiveMutex.RUnlock()
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
        http.ServeFile(w, r, rec.Dir+"/"+rec.Name+"/1."+rec.Codec)
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
func archiveItemDelete(w http.ResponseWriter, r *http.Request) {
    archiveMutex.Lock()
    defer archiveMutex.Unlock()

    queryParams := mux.Vars(r)

    item, exist := archive[queryParams["uuid"]]
    if !exist {
        fmt.Fprintf(w, "Not found")
        return
    }
    if item.busy == true {
        fmt.Fprintf(w, "Busy")
        return
    }

    rec := item.record

    os.RemoveAll(rec.Dir+"/"+rec.Name)
    //if arcive item is busy it can`t be deleted

    delete(archive, queryParams["uuid"])

    fmt.Fprintf(w, "archiveDelete")
}

////////////////////////////////////////////////////////////////////////////////
func archiveItemPlayer(w http.ResponseWriter, r *http.Request) {
    queryParams := mux.Vars(r)

    archiveMutex.RLock()
    defer archiveMutex.RUnlock()

    item, exist := archive[queryParams["uuid"]]
    if !exist {
        fmt.Fprintf(w, "Not found")
        return
    }

    rec := item.record

    fmt.Fprintf(w, "<html>")
    fmt.Fprintf(w, "<head>")
    fmt.Fprintf(w, "<title>Archive: %s</title>", rec.Name)
    fmt.Fprintf(w, "</head>")
    fmt.Fprintf(w, "<body>")
    fmt.Fprintf(w, "<script src=\"/js/hls.js\"></script>")
    fmt.Fprintf(w, "<center>")
    fmt.Fprintf(w, "<h1>Archive: %s</h1>", rec.Name)
    fmt.Fprintf(w, "<video height=\"720\" id=\"video\" controls></video>")
    fmt.Fprintf(w, "</center>")
    fmt.Fprintf(w, "<script>")
    fmt.Fprintf(w, "if(Hls.isSupported()) {")
    fmt.Fprintf(w, "var video = document.getElementById('video');")
    fmt.Fprintf(w, "var hls = new Hls({")
    fmt.Fprintf(w, "debug: true")
    fmt.Fprintf(w, "});")
    fmt.Fprintf(w, "hls.loadSource('/archive/%s/index.m3u8');", rec.Name)
    fmt.Fprintf(w, "hls.attachMedia(video);")
    fmt.Fprintf(w, "hls.on(Hls.Events.MEDIA_ATTACHED, function() {")
    fmt.Fprintf(w, "video.muted = true;")
    fmt.Fprintf(w, "video.play();")
    fmt.Fprintf(w, "});")
    fmt.Fprintf(w, "}")
    fmt.Fprintf(w, "else if (video.canPlayType('application/vnd.apple.mpegurl')) {")
    fmt.Fprintf(w, "video.src = '/archive/%s/index.m3u8';", rec.Name)
    fmt.Fprintf(w, "video.addEventListener('canplay',function() {")
    fmt.Fprintf(w, "video.play();")
    fmt.Fprintf(w, "});")
    fmt.Fprintf(w, "}")
    fmt.Fprintf(w, "</script>")
    fmt.Fprintf(w, "</body>")
    fmt.Fprintf(w, "</html>")
}
