// +build openapi

package openapi

import (
	"flag"
	"github.com/gorilla/mux"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

////////////////////////////////////////////////////////

type route struct {
	name        string
	method      string
	pattern     string
	handlerFunc http.HandlerFunc
}

type routes []route

var apiRoutes routes

func AddRoute(name, pattern, method string, handlerfunc http.HandlerFunc) {
	apiRoutes = append(apiRoutes, route{name: name, method: method, pattern: pattern, handlerFunc: handlerfunc})
}

////////////////////////////////////////////////////////

var flagUdsPath 	*string
var flagHttpPort	*int
var flagWwwPath		*string

func init() {
	flagUdsPath     = flag.String   ("--openapi-sock", "/tmp/application.sock", "UDS socket file absolute path")
	flagHttpPort    = flag.Int      ("--openapi-port", 80,                  	"Http port")
	flagWwwPath     = flag.String   ("--openapi-www", "/opt/www",           	"Www static files path")
}

////////////////////////////////////////////////////////

func Init() {
	log.Println("Openapi is ON!")

	router := NewRouter()

	log.Println("Starting USD HTTP server")

	os.Remove(*flagUdsPath)
	l, err := net.Listen("unix", *flagUdsPath)
	if err != nil {
		log.Printf("error: %v\n", err)
		return
	}
	go http.Serve(l, router)

	log.Println("Starting NET HTTP server")
	srv := &http.Server{
		Addr:           ":80",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go srv.ListenAndServe()

}

func NewRouter() *mux.Router {
	router := mux.NewRouter() //.StrictSlash(true)

	for _, route := range apiRoutes {
		var handler http.Handler 
		handler = route.handlerFunc
		//handler = Logger(handler, route.Name)

		router.
			Methods(route.method).
			Path(route.pattern).
			Name(route.name).
			Handler(handler)
	}

    /*
    router.HandleFunc("/api/debug/umap", debugUmap).Methods("GET")
    router.HandleFunc("/api/debug/umap/{file}", debugUmapFile).Methods("GET")
    router.HandleFunc("/api/debug/vars", debugExpvar).Methods("GET")
    */

	/*
    router.HandleFunc("/api/system", system).Methods("GET")
    router.HandleFunc("/api/system/date", systemDate).Methods("GET")
    router.HandleFunc("/api/system/network", systemNetwork).Methods("GET")
	*/

	//router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("/opt/www"))))
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(*flagWwwPath))))

	return router
}
