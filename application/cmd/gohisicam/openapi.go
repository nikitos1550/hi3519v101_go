package main

import (
    "fmt"
    "flag"
    "strconv"
    "time"
    "net/http"
    "strings"
    "sync"

    "github.com/gorilla/mux"

    "application/core/logger"
    "application/core/openapi/system"
    "application/core/openapi/godebug"
    "application/core/openapi/umap"
    "application/core/openapi/compiletime"
    "application/core/openapi/mpp"
    "application/core/openapi/chip"
    "application/core/openapi/temperature"
    "application/core/openapi/channel"
    "application/core/openapi/encoder"
    "application/core/openapi/jpeg"
    "application/core/openapi/mjpeg"
    //"application/core/openapi/webrtc"

    //"application/core/openapi/forward"

    "application/core/openapi/link"
    "application/core/openapi/crud"
)

var (
    router *mux.Router

    flagUdsPath     *string
    flagHttpPort    *uint
    flagWwwPath     *string
    flagLogReqs     *bool

    pipelineLock    sync.RWMutex
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

    debug := router.PathPrefix("/debug/").Subrouter()

    debug.HandleFunc("/routes", apiList).Methods("GET").Name("Routes list")

    debug.HandleFunc("/vars", godebug.Expvar).Methods("GET")

    debug.HandleFunc("/umap", umap.List).Methods("GET")
    debug.HandleFunc("/umap.json", umap.ListJson).Methods("GET")
    debug.HandleFunc("/umap/{file:[a-z0-9_]+}", umap.File).Methods("GET")
    debug.HandleFunc("/umap/{file:[a-z0-9_]+}.json", umap.FileJson).Methods("GET")

    ////////////////////////////////////////////////////////////////////////////

    api := router.PathPrefix("/api/").Subrouter()

    api.HandleFunc("/buildinfo", compiletime.Serve).Methods("GET")
    api.HandleFunc("/temperature", temperature.Serve).Methods("GET")
    api.HandleFunc("/chip", chip.Info).Methods("GET")

    api.HandleFunc("/mpp/version", mpp.Version).Methods("GET")
    api.HandleFunc("/mpp/syncpts", mpp.RunSyncPts).Methods("GET")
    api.HandleFunc("/mpp/initpts", mpp.RunInitPts).Methods("GET")

    api.HandleFunc("/channel", channel.GroupListHandler(channelGroup)).Methods("GET")
    api.HandleFunc("/channel/{name}", channel.GroupInfoHandler(channelGroup)).Methods("GET")
    api.HandleFunc("/channel/{name}", channel.GroupCreateHandler(channelGroup)).Methods("POST")
    api.HandleFunc("/channel/{name}", channel.GroupDeleteHandler(channelGroup)).Methods("DELETE")
    api.HandleFunc("/channel/{name}/stat", channel.GroupStatHandler(channelGroup)).Methods("GET")

    api.HandleFunc("/encoder", encoder.GroupListHandler(encoderGroup)).Methods("GET")
    api.HandleFunc("/encoder/{name}", encoder.GroupInfoHandler(encoderGroup)).Methods("GET")
    api.HandleFunc("/encoder/{name}", encoder.GroupNewHandler(encoderGroup)).Methods("POST")
    api.HandleFunc("/encoder/{name}", encoder.GroupUpdateHandler(encoderGroup)).Methods("PUT")
    api.HandleFunc("/encoder/{name}", encoder.GroupDeleteHandler(encoderGroup)).Methods("DELETE")
    api.HandleFunc("/encoder/{name}/start", encoder.GroupStartHandler(encoderGroup)).Methods("GET")
    api.HandleFunc("/encoder/{name}/stop", encoder.GroupStopHandler(encoderGroup)).Methods("GET")
    api.HandleFunc("/encoder/{name}/idr", encoder.GroupIdrHandler(encoderGroup)).Methods("GET")

    api.HandleFunc("/jpeg", jpeg.GroupListHandler(jpegGroup)).Methods("GET")
    api.HandleFunc("/jpeg/{name}", jpeg.GroupCreateHandler(jpegGroup)).Methods("POST")
    api.HandleFunc("/jpeg/{name}", jpeg.GroupInfoHandler(jpegGroup)).Methods("GET")
    api.HandleFunc("/jpeg/{name}", jpeg.GroupDeleteHandler(jpegGroup)).Methods("DELETE")

    api.HandleFunc("/mjpeg", mjpeg.GroupListHandler(mjpegGroup)).Methods("GET")
    api.HandleFunc("/mjpeg/{name}", mjpeg.GroupCreateHandler(mjpegGroup)).Methods("POST")
    api.HandleFunc("/mjpeg/{name}", mjpeg.GroupInfoHandler(mjpegGroup)).Methods("GET")
    api.HandleFunc("/mjpeg/{name}", mjpeg.GroupDestroyHandler(mjpegGroup)).Methods("DELETE")
    api.HandleFunc("/mjpeg/{name}/client", mjpeg.GroupClientsHandler(mjpegGroup)).Methods("GET")
    api.HandleFunc("/mjpeg/{name}/client/{client}", mjpeg.GroupClientDeleteHandler(mjpegGroup)).Methods("GET")

    //api.HandleFunc("/webrtc", webrtc.List).Methods("GET")
    //api.HandleFunc("/webrtc", webrtc.Create).Methods("POST")
    //api.HandleFunc("/webrtc/{id:[0-9]+}", webrtc.Info).Methods("GET")
    //api.HandleFunc("/webrtc/{id:[0-9]+}", webrtc.Delete).Methods("DELETE")

    api.HandleFunc("/forward", crud.GroupListHandler(forwardGroup)).Methods("GET")
    api.HandleFunc("/forward/{name}", crud.GroupCreateHandler(forwardGroup)).Methods("POST")
    //api.HandleFunc("/forward/{name}", forward.GroupInfoHandler(forwardGroup)).Methods("GET")
    api.HandleFunc("/forward/{name}", crud.GroupDeleteHandler(forwardGroup)).Methods("DELETE")

    api.HandleFunc("/quirc", crud.GroupListHandler(quircGroup)).Methods("GET")
    api.HandleFunc("/quirc/{name}", crud.GroupCreateHandler(quircGroup)).Methods("POST")
    //api.HandleFunc("/quirc/{name}", forward.GroupInfoHandler(quircGroup)).Methods("GET")
    api.HandleFunc("/quirc/{name}", crud.GroupDeleteHandler(quircGroup)).Methods("DELETE")

    api.HandleFunc("/link/channel/{source}/encoder/{client}", link.ConnectBindHandler(channelGroup, encoderGroup)).Methods("POST")
    api.HandleFunc("/link/channel/{source}/encoder/{client}", link.DisconnectBindHandler(channelGroup, encoderGroup)).Methods("DELETE")
    //api.HandleFunc("/link/channel/{source}/encoder/{client}/raw", link.ConnectRawFrameHandler(channelGroup, encoderGroup)).Methods("POST")
    //api.HandleFunc("/link/channel/{source}/encoder/{client}/raw", link.DisconnectRawFrameHandler(channelGroup, encoderGroup)).Methods("DELETE")

    api.HandleFunc("/link/encoder/{source}/jpeg/{client}", link.ConnectEncodedDataHandler(encoderGroup, jpegGroup)).Methods("POST")
    api.HandleFunc("/link/encoder/{source}/jpeg/{client}", link.DisconnectEncodedDataHandler(encoderGroup, jpegGroup)).Methods("DELETE")

    api.HandleFunc("/link/encoder/{source}/mjpeg/{client}", link.ConnectEncodedDataHandler(encoderGroup, mjpegGroup)).Methods("POST")
    api.HandleFunc("/link/encoder/{source}/mjpeg/{client}", link.DisconnectEncodedDataHandler(encoderGroup, mjpegGroup)).Methods("DELETE")

    api.HandleFunc("/link/channel/{source}/forward/{client}", link.ConnectRawFrameHandler(channelGroup, forwardGroup)).Methods("POST")
    api.HandleFunc("/link/channel/{source}/forward/{client}", link.DisconnectRawFrameHandler(channelGroup, forwardGroup)).Methods("DELETE")
    api.HandleFunc("/link/forward/{source}/encoder/{client}", link.ConnectRawFrameHandler(forwardGroup, encoderGroup)).Methods("POST")
    api.HandleFunc("/link/forward/{source}/encoder/{client}", link.DisconnectRawFrameHandler(forwardGroup, encoderGroup)).Methods("DELETE")

    api.HandleFunc("/link/channel/{source}/quirc/{client}", link.ConnectRawFrameHandler(channelGroup, quircGroup)).Methods("POST")
    api.HandleFunc("/link/channel/{source}/quirc/{client}", link.DisconnectRawFrameHandler(channelGroup, quircGroup)).Methods("DELETE")
    api.HandleFunc("/link/quirc/{source}/encoder/{client}", link.ConnectRawFrameHandler(quircGroup, encoderGroup)).Methods("POST")
    api.HandleFunc("/link/quirc/{source}/encoder/{client}", link.DisconnectRawFrameHandler(quircGroup, encoderGroup)).Methods("DELETE")

    //api.HandleFunc("/link/encoder/{encoder:[0-9]+}/webrtc/{webrtc:[0-9]+}", webrtc.BindEncoder).Methods("POST")
    //api.HandleFunc("/link/encoder/{encoder:[0-9]+}/webrtc/{webrtc:[0-9]+}", webrtc.UnbindEncoder).Methods("DELETE")

    ////////////////////////////////////////////////////////////////////////////

    serve := router.PathPrefix("/serve/").Subrouter()

    serve.HandleFunc("/jpeg/{name}.{ext:jpg|jpeg}", jpegGroup.ServeFrameGroup).Methods("GET")

    serve.HandleFunc("/mjpeg/{name}.{ext:mjpg|mjpeg}", mjpegGroup.ServeStreamGroup).Methods("GET")

    //serve.HandleFunc("/webrtc/{id:[0-9]+}", webrtc.Connect).Methods("POST")
    //serve.HandleFunc("/webrtc/{id:[0-9]+}/{uuid:[0-9a-z_-]+}", webrtc.Disconnect).Methods("DELETE")

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
