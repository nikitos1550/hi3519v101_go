//+build debug
//+build openapi 

package mpp

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	//"expvar"
	"application/pkg/openapi"
)

func init() {
    openapi.AddApiRoute("debugUmap",       "/debug/umap",          "GET",      debugUmap)
	openapi.AddApiRoute("debugUmapFile",   "/debug/umap/{file}",   "GET",      debugUmapFile)
	
	//TODO This is shouldn`t be here, not related to MPP
    //AddRoute("debugExpvar",     "/api/debug/vars",          "GET",      debugExpvar)
}

func debugUmap(w http.ResponseWriter, r *http.Request) {
    log.Println("debugUmap")

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

func debugUmapFile(w http.ResponseWriter, r *http.Request) {
    log.Println("debugUmapFile")

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


/* TODO This is shouldn`t be here, not related to MPP
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
*/