// +build openapi

package openapi

import (
	"flag"
	"github.com/gorilla/mux"
	//"log"
	"net"
	"net/http"
	"os"
	"time"
	//"strings"
	"strconv"

    "github.com/gorilla/websocket"

	"application/pkg/logger"
)

////////////////////////////////////////////////////////

const apiPrefix string = "/api/"
const wsPrefix string = "/ws/"

type routeItem struct {
	name        string
	method      string
	pattern     string
	handlerFunc http.HandlerFunc
}

type routeItems []routeItem

var apiRoutes	routeItems //var for initial api routes storage
var routes		routeItems
var wsRoutes    routeItems

func AddApiRoute(name, pattern, method string, handlerfunc http.HandlerFunc) {
	apiRoutes = append(apiRoutes, routeItem {name: name, method: method, pattern: pattern, handlerFunc: handlerfunc})
}

func AddRoute(name, pattern, method string, handlerfunc http.HandlerFunc) {
	routes = append(routes, routeItem {name: name, method: method, pattern: pattern, handlerFunc: handlerfunc})
}

func AddWsRoute(name, pattern, method string, handlerfunc http.HandlerFunc) {
    wsRoutes = append(wsRoutes, routeItem {name: name, method: method, pattern: pattern, handlerFunc: handlerfunc})
}

var router *mux.Router

////////
/*
func AddApiRoute2(name, pattern, method string, Method) {
    apiRoutes = append(apiRoutes, routeItem {name: name, method: method, pattern: pattern, handlerFunc: handlerfunc})
}
*/

type Method interface {
    getInputJson()  int
    getOutputJson() int
}

func serveApi(w http.ResponseWriter, r *http.Request) {
    //log.Println("serveApi")
    logger.Log.Trace().Msg("serveApi")

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    //find corresponding api
    /*
    var dtSchema apiAnswerSystemDateTimeSchema
    t := time.Now()

    dtSchema.Formatted = t
    dtSchema.Secs = t.Unix()
    dtSchema.Nanosecs = t.UnixNano()
    */
    /*
    dtSchemaJson, _ := json.Marshal(dtSchema)
    fmt.Fprintf(w, "%s", string(dtSchemaJson))
    */
}


////////

//var Upgrader = websocket.Upgrader{} // use default options
var Upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

////////////////////////////////////////////////////////

var flagUdsPath 	*string
var flagHttpPort	*uint
var flagWwwPath		*string
//var flagPrintRoutes	*bool

func init() {
	flagUdsPath     = flag.String   ("openapi-socket", 	"/tmp/application.sock", "UDS socket file absolute path")
	flagHttpPort    = flag.Uint     ("openapi-port", 	80,                  	 "Http port")
	flagWwwPath     = flag.String   ("openapi-www", 	"/opt/www",           	 "Www static files path")
	//flagPrintRoutes = flag.Bool		("openapi-routes",	false, 					 "Prints application version information")
}

////////////////////////////////////////////////////////
//TODO
/*
func PrintInfo() {//TODO
    router = newRouter()

        //TODO use it or make another cmd/app to generate auto doc
        router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
            t, err := route.GetPathTemplate()
            if err != nil {
                return err
            }
            m, err := route.GetMethods()
            if err != nil {
                return err
            }
            //r, err := route.GetPathRegexp()
            //if err == nil {
            //  return err
            //}
            
	    //log.Println(m, " ", t)//, " ", r)
            return nil
        })

}
*/

func Init() {
	router = newRouter()

	//TODO use it or make another cmd/app to generate auto doc
    /*
	if *flagPrintRoutes {
		router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
			t, err := route.GetPathTemplate()
			if err != nil {
				return err
			}
			m, err := route.GetMethods()
			if err != nil {
				return err
			}
			//r, err := route.GetPathRegexp()
			//if err == nil {
			//	return err
			//}
			log.Println(m, " ", t)//, " ", r)
			return nil
		})
		os.Exit(0)
	}
    */
	//TODO check flags values
}

func Start() {
	//log.Println("Starting NET HTTP server")
	logger.Log.Debug().
		Uint("port", *flagHttpPort).
		Msg("Starting NET HTTP server")

	//TODO check ability to bind port
	srv := &http.Server{
		Addr:           ":"+strconv.FormatUint(uint64(*flagHttpPort), 10),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go srv.ListenAndServe()

	//log.Println("Starting USD server")
	logger.Log.Debug().
		Str("file", *flagUdsPath).
		Msg("Starting USD server")

	os.Remove(*flagUdsPath)
	l, err := net.Listen("unix", *flagUdsPath)
	if err != nil {
		//log.Printf("error: %v\n", err)
		logger.Log.Error().
			Str("reason", err.Error()).
			Msg("Error creating UDS")
		return
	}
	go http.Serve(l, router)
}

func newRouter() *mux.Router {
	router      := mux.NewRouter() //.StrictSlash(true)
    wsRouter    := router.PathPrefix(wsPrefix).Subrouter()
    apiRouter   := router.PathPrefix(apiPrefix).Subrouter()

    //wsRouter.Use(wsMiddleware)

    for _, route := range wsRoutes {
        wsRouter.
            Methods(route.method).
            Path(route.pattern).
            Name(route.name).
            Handler(route.handlerFunc)
    }

	for _, route := range apiRoutes {
		apiRouter.
			Methods(route.method).
			Path(route.pattern).
			Name(route.name).
			Handler(route.handlerFunc)
	}

	for _, route := range routes {
		router.
			Methods(route.method).
			Path(route.pattern).
			Name(route.name).
			Handler(route.handlerFunc)
	}

	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(*flagWwwPath))))

    //router.Use(authMiddleware)
	return router
}


