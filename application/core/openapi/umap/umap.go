package umap

import (
	"fmt"
	"io/ioutil"
	"net/http"
    "strings"
    "strconv"

    "github.com/gorilla/mux"

    "application/core/logger"
)

func List(w http.ResponseWriter, r *http.Request) {
    logger.Log.Trace().
        Msg("debugUmap")

	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

	files, err := ioutil.ReadDir("/proc/umap")
	if err != nil {
		//TODO /proc/umap exist only after ko modules init, handle it smart!
		w.WriteHeader(http.StatusNotFound) //TODO correct status
		//panic(err)
		return
	}

	w.WriteHeader(http.StatusOK)

	num := len(files)
	var i int = 0

	for i < (num - 1) {
		//fmt.Println(f.Name())
		fmt.Fprintf(w, files[i].Name())
		fmt.Fprintf(w, ",")
		i++
	}
	fmt.Fprintf(w, files[num-1].Name())
}

func ListJson(w http.ResponseWriter, r *http.Request) {
    //log.Println("debugUmapJson")
    logger.Log.Trace().
        Msg("debugUmapJson")


    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

    files, err := ioutil.ReadDir("/proc/umap")
    if err != nil {
        //TODO /proc/umap exist only after ko modules init, handle it smart!
        w.WriteHeader(http.StatusNotFound) //TODO correct status
        //panic(err)
        return
    }

    w.WriteHeader(http.StatusOK)

    num := len(files)
    var i int = 0
    fmt.Fprintf(w, "[")
    for i < (num - 1) {
        //fmt.Println(f.Name())
        fmt.Fprintf(w, "\"")
        fmt.Fprintf(w, files[i].Name())
        fmt.Fprintf(w, "\",")
        i++
    }
    fmt.Fprintf(w, "\"")
    fmt.Fprintf(w, files[num-1].Name())
    fmt.Fprintf(w, "\"]")
}

func File(w http.ResponseWriter, r *http.Request) {
    logger.Log.Trace().
        Msg("debugUmapFile")

	params := mux.Vars(r)

	dat, err := ioutil.ReadFile("/proc/umap/" + params["file"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, string(dat))
}

/*
    ATTENTION!
    Json parsing is working somehow, not every file is parsed correct way
    So before usage compare raw data and json, if for your task json is ok
    (contains needed data), than use it otherwise use raw data.
*/

func FileJson(w http.ResponseWriter, r *http.Request) {
    logger.Log.Trace().
        Msg("debugUmapFileJson")

    params := mux.Vars(r)

    dat, err := ioutil.ReadFile("/proc/umap/" + params["file"])
    if err != nil {
        w.WriteHeader(http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    result := parseUmap(string(dat))
    //log.Println(result)
    /*
    strs := strings.Split(string(dat), "\n")
    for _, s := range strs[2:] {
        s = strings.TrimSpace(s)
        if s == "" {continue}
         log.Println(s)
         
         if s[0] == '-' {
            
            

         }
    }
    */
    //test := strings.Fields(string(dat))

    //log.Println(strs)

    fmt.Fprintf(w, result)
}

func parseUmap(in string) string {
    var out string
    var state uint

    var tmpName string
    var tmpHeader []string
    var tmpVals [][]string

    out = "{ "

    strs := strings.Split(in, "\n")
    for _, s := range strs {
        s = strings.TrimSpace(s)
        if s == "" {
            continue
        }
        if s[0] == '[' {
            if state > 0 {
                out = out + parseUmapDump(tmpName, tmpHeader, tmpVals)
                out = out + " },"
            }
            i := strings.Index(s[1:], "]")
            if i > 0 {
                 out = out + " \"" + s[1:1+i] + "\": { "
            } else {
                 out = out + "\"unknown\" : { "
            }
            state = 1
            continue
        }

        if s[0] == '-' { // block ended, new block
            if state == 0 {
                out = out + "\"unknown\" : { "
            }
            //TODO dump tmp
            if state > 1 {
                out = out + parseUmapDump(tmpName, tmpHeader, tmpVals) + ", "
            }
            //log.Println(tmpName)
            //log.Println(tmpHeader)
            //log.Println(tmpVals)
            //
            //s = s[5:] //-----

            i := strings.Index(s[5:], "-")
            if i > 0 {
                tmpName = s[5:5+i]
            } else {
                tmpName = "unknown"
            }
            //tmpH

            state = 2
            continue
        }
        if state == 2 { //parsing header
            vals := strings.Fields(s)
            tmpHeader = make([]string, len(vals))
            tmpVals = make([][]string, len(vals))
            for counter, val := range vals {
                tmpHeader[counter] = val
                //tmpData[counter] = make([]string, 1)
            }
            state = 3
            continue
        }
        if state == 3 { //pasing values
            vals := strings.Fields(s)
            if len(vals) > len(tmpHeader) {
                //log.Println("some parse bug skipping")
                for counter:=0; counter<len(tmpHeader);counter++ {
                    tmpVals[counter] = append(tmpVals[counter], "N/A")
                }
                continue
            }
            for counter, val := range vals {
                tmpVals[counter] = append(tmpVals[counter], val)
            }
            continue
        }
    }
    //log.Println(tmpName)
    //log.Println(tmpHeader)
    //log.Println(tmpVals)
    out = out + parseUmapDump(tmpName, tmpHeader, tmpVals) + " } }"
    return out
}

func parseUmapDump(name string, header []string, vals [][]string) string {
    var out string
    out = out + " \"" + name + "\" : {"
    for counter, head := range header {
        out = out + " \"" + head + "\" : "
        if len(vals[counter]) > 1 {
            out = out + " ["
            for counter2, val := range vals[counter] {
                out = out + parseUmapVal(val)
                if counter2 < len(vals[counter]) - 1 {
                    out = out + ", "
                }
            }
            out = out + "] "
        } else if len(vals[counter]) == 1 {
            out = out + parseUmapVal(vals[counter][0])
        } else {
            out = out + "null"//"\"\""
        }
        if counter < len(header)-1 {
            out = out + " ,"
        }
    }
    out = out + " }"
    return out
}

func parseUmapVal(val string) string {
    if _, err := strconv.Atoi(val); err == nil {
        return val
    } else {
        return " \"" + val + "\" "
    }
}
