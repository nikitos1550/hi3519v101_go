package main

import (
    "fmt"
    "flag"
    "net/http"
    "time"

    "github.com/drone/routes"

    "./himpp3"
    "./hidebug"

    "strconv"

    "regexp"
)

var BuildTime string

func main() {
	fmt.Println(BuildTime)

	himpp3.SysInit()
	himpp3.MipiIspInit()
	himpp3.ViInit()
	himpp3.VpssInit()
	himpp3.VencInit()
	
	

	portPtr := flag.Int("port", 80, "http port")

	flag.Parse()

	fmt.Printf("minimal application\n");
	fmt.Printf("http port %d\n", *portPtr);

    /////

    mux := routes.New()

	mux.Get("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Println("Requested url /")
		fmt.Fprintf(w, "Camera go webserver!")
	})

    mux.Get("/t", func (w http.ResponseWriter, r *http.Request) {
        fmt.Println("Requested url /t")
        fmt.Fprintf(w, "%.1fC", himpp3.TempGet())
    })

	mux.Get("/image.jpeg", func (w http.ResponseWriter, r *http.Request) {

        //buffer := new(bytes.Buffer)
        //if err := jpeg.Encode(buffer, *img, nil); err != nil {
            //log.Println("unable to encode image.")
        //}

		fmt.Println("Serving jpeg... ", len(himpp3.B1.Bytes()))

		himpp3.Mutex.Lock()
        w.Header().Set("Content-Type", "image/jpeg")
        w.Header().Set("Content-Length", strconv.Itoa(len(himpp3.B1.Bytes())))
        
        if _, err := w.Write(himpp3.B1.Bytes()); err != nil {
            fmt.Println("unable to write image.")
        }

		himpp3.Mutex.Unlock()
		fmt.Println("done!")
   	})

    mux.Get("^/experimental/date.(text|sec|nano)$", func (w http.ResponseWriter, r *http.Request) {
        rr, _ := regexp.Compile("^/experimental/date.(text|sec|nano)$")
        match := rr.FindStringSubmatch(r.URL.Path)
        fmt.Println(match)
    
        t := time.Now()
        switch match[1] {
            case "text":
                fmt.Fprintf(w, t.String())
            case "sec" :
                fmt.Fprintf(w, "%d", t.Unix())
            case "nano":
                fmt.Fprintf(w, "%d", t.UnixNano())
        }

    })

    mux.Get("^/experimental/hidebug/?$", hidebug.ApiListHandler)
    mux.Get("^/experimental/hidebug/(.+).(raw|json)$", hidebug.ApiFileHandler)

	//fs := http.FileServer(http.Dir("static/"))
	//http.Handle("/static/", http.StripPrefix("/static/", fs))

    //srv.Handle("/", mux)

	//http.ListenAndServe(":80", nil)

    srv := &http.Server{
        Addr:           ":80",
        Handler:        mux,
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }
    srv.ListenAndServe()
}
