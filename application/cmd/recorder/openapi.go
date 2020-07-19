package main

import (
    "fmt"
    "flag"
    "strconv"
    "time"
    "net/http"
    "strings"

    "github.com/gorilla/mux"

    "application/core/logger"
    "application/core/openapi/system"
    "application/core/openapi/godebug"
    "application/core/openapi/umap"
    "application/core/openapi/compiletime"
    "application/core/openapi/mpp"
    "application/core/openapi/chip"
    "application/core/openapi/temperature"
)

var (
    router *mux.Router

    flagUdsPath     *string
    flagHttpPort    *uint
    flagWwwPath     *string
    flagLogReqs     *bool
)

func init() {
    flagUdsPath     = flag.String   ("openapi-socket",  "/tmp/application.sock", "UDS socket file absolute path")
    flagHttpPort    = flag.Uint     ("openapi-port",    80,                      "Http port")
    flagWwwPath     = flag.String   ("openapi-www",     "/opt/www",              "Www static files path")
    flagLogReqs     = flag.Bool     ("openapi-log",     true,                    "Log all request URLs")
}

func httpServerStart() {
    router = mux.NewRouter()

    router.HandleFunc("/system/date", system.Date).Methods("GET").Name("date")

    ////////////////////////////////////////////////////////////////////////////

    debug := router.PathPrefix("/debug").Subrouter()

    debug.HandleFunc("/routes", apiList).Methods("GET").Name("Routes list")

    debug.HandleFunc("/vars", godebug.Expvar).Methods("GET")

    debug.HandleFunc("/umap", umap.List).Methods("GET")
    debug.HandleFunc("/umap.json", umap.ListJson).Methods("GET")
    debug.HandleFunc("/umap/{file:[a-z0-9_]+}", umap.File).Methods("GET")
    debug.HandleFunc("/umap/{file:[a-z0-9_]+}.json", umap.FileJson).Methods("GET")

    ////////////////////////////////////////////////////////////////////////////

    api := router.PathPrefix("/api").Subrouter()

    api.HandleFunc("/buildinfo", compiletime.Serve).Methods("GET")
    api.HandleFunc("/temperature", temperature.Serve).Methods("GET")
    api.HandleFunc("/chip", chip.Info).Methods("GET")

    api.HandleFunc("/mpp/version", mpp.Version).Methods("GET")
    api.HandleFunc("/mpp/syncpts", mpp.RunSyncPts).Methods("GET")
    api.HandleFunc("/mpp/initpts", mpp.RunInitPts).Methods("GET")

    api.HandleFunc("/recorder", recorderStatus).Methods("GET")
    api.HandleFunc("/recorder/start", recorderStart).Methods("GET")
    api.HandleFunc("/recorder/stop", recorderStop).Methods("GET")
    //api.HandleFunc("/recorder/schedule", recorderSchedule).Methods("GET")

    ////////////////////////////////////////////////////////////////////////////

    serve := router.PathPrefix("/serve").Subrouter()

    serve.HandleFunc("/image.{ext:jpg|jpeg}", jpegSmall.ServeFrame).Methods("GET")

    ////////////////////////////////////////////////////////////////////////////

    router.HandleFunc("/archive", archiveList).Methods("GET")
    archive := router.PathPrefix("/archive").Subrouter()

    archive.HandleFunc("/{uuid:[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}}", archiveItemInfo).Methods("GET")
    archive.HandleFunc("/{uuid:[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}}/preview.{ext:jpg|jpeg}", archiveItemPreview).Methods("GET")
    archive.HandleFunc("/{uuid:[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}}/download.h264", archiveItemServe).Methods("GET")

    ////////////////////////////////////////////////////////////////////////////

    router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(*flagWwwPath)))).Name("Serve static files")

    //TODO doesn`t work with file server on /
    //router.NotFoundHandler = http.HandlerFunc(notFound) //handler declaration below

    if *flagLogReqs {
        router.Use(loggerMiddleware)
    }
    api.Use(lockerWMiddleware)
	api.Use(apiMiddleware)
    //serve.Use(lockerRMiddleware)

    //TODO check ability to bind port
    srv := &http.Server{
        Addr:           ":"+strconv.FormatUint(uint64(*flagHttpPort), 10),
        Handler:        router,
        ReadTimeout:    1 * time.Second,
        //WriteTimeout:   5 * time.Second,  //TMP for mjpeg
        MaxHeaderBytes: 1 << 20,
    }
    go srv.ListenAndServe()

    logger.Log.Debug().
        Uint("port", *flagHttpPort).
        Msg("HTTP server started")

    //os.Remove(*flagUdsPath)
    //l, err := net.Listen("unix", *flagUdsPath)
    //if err != nil {
    //    logger.Log.Error().
    //        Str("reason", err.Error()).
    //        Msg("Error creating UDS")
    //    return
    //}
    //go http.Serve(l, router)
    //logger.Log.Debug().
    //    Str("file", *flagUdsPath).
    //    Msg("Starting USD server")
}

func apiMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        next.ServeHTTP(w, r)
    })
}

func lockerWMiddleware(next http.Handler) http.Handler { //one api request per time
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        pipelineLock.Lock()
        next.ServeHTTP(w, r)
        pipelineLock.Unlock()
    })
}

func lockerRMiddleware(next http.Handler) http.Handler { //one api request per time
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        pipelineLock.RLock()
        next.ServeHTTP(w, r)
        pipelineLock.RUnlock()
    })
}

func loggerMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        logger.Log.Trace().
            Str("uri", r.RequestURI).
            Str("method", r.Method).
            Str("client", r.RemoteAddr).
            Msg("HTTP")
        next.ServeHTTP(w, r)
        //logger.Log.Trace().
        //    Str("code", w.statusCode).
        //    Msg("HTTP")
    })
}

func apiList(w http.ResponseWriter, r *http.Request) {
    router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
        pathName := route.GetName()
        fmt.Fprintf(w, "Name: %s", pathName)

        methods, err := route.GetMethods()
        if err == nil {
            fmt.Fprintf(w, ", METHOD: %s", strings.Join(methods, ","))
        }

		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			fmt.Fprintf(w, ", PATH: %s\n", pathTemplate)
		}

		return nil
	})
}

/*
func notFound(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    fmt.Fprintf(w, "{\"error\":\"route not found\"}")
}
*/
