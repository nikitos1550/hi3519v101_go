package hidebug

import (
    "fmt"
    "net/http"

//	"bufio"
//    "io"
    "io/ioutil"
//    "os"

    "regexp"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

/*
func getList() []string {
    files, err := ioutil.ReadDir("/proc/umap")
    if err != nil {
        return []string
    }

    
    var out []string

    for _, f := range files {
        
    }
    return out
  }
*/

func ApiListHandler (w http.ResponseWriter, r *http.Request) {
    files, err := ioutil.ReadDir("/proc/umap")
    if err != nil {
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

func ApiFileHandler (w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.URL.Path)

    //rr, _       := regexp.Compile("^/experimental/hidebug/(.+).(raw|json)$")
    rr, _       := regexp.Compile("^/experimental/hidebug/(chnl|h265e|jpege|rgn|venc|vo|fisheye|hi_mipi|logmpp|sys|vgs|vpss|h264e|isp|rc|vb|vi).(raw|json)$")

    match       := rr.FindStringSubmatch(r.URL.Path)
    fmt.Println(match)
    if match != nil {
        fmt.Println(match[1])
        fmt.Println(match[2])

        if match[2] == "raw" {
            dat, err := ioutil.ReadFile("/proc/umap/" + match[1])
            check(err)
            //fmt.Print(string(dat))
            fmt.Fprintf(w, string(dat))
        } else {
            return

        }

    } else {
        http.NotFound(w, r)
        return
    }

    /*
    # ls /proc/umap/
    chnl|h265e|jpege|rgn|venc|vo|fisheye|hi_mipi|logmpp|sys|vgs|vpss|h264e|isp|rc|vb|vi
    */

    /*
    switch match[1] {
        case 1:
            fmt.Println("one")
        case 2:
        fmt.Println("two")
        case 3:
        fmt.Println("three")
        default:
            http.NotFound(w, r)
            return
    }
    */

    //dat, err := ioutil.ReadFile("/proc/umap/vb")
    //check(err)
    //fmt.Print(string(dat))
    //fmt.Fprintf(w, string(dat))
}
