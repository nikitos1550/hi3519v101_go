package main

import (
    	"fmt"
	"flag"
	"net/http"
	"./himpp3"
	"strconv"
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

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Println("Root serving")
		fmt.Fprintf(w, "Camera go webserver!")
	})

        http.HandleFunc("/t", func (w http.ResponseWriter, r *http.Request) {
                fmt.Println("Temperature")
                fmt.Fprintf(w, "%.1fC", himpp3.TempGet())
        })

	http.HandleFunc("/image.jpeg", func (w http.ResponseWriter, r *http.Request) {

        	//buffer := new(bytes.Buffer)
        	//if err := jpeg.Encode(buffer, *img, nil); err != nil {
            	//log.Println("unable to encode image.")
        	//}

		fmt.Println("serving jpeg... ", len(himpp3.B1.Bytes()))

		himpp3.Mutex.Lock()
        	w.Header().Set("Content-Type", "image/jpeg")
        	w.Header().Set("Content-Length", strconv.Itoa(len(himpp3.B1.Bytes())))
        	if _, err := w.Write(himpp3.B1.Bytes()); err != nil {
            		fmt.Println("unable to write image.")
        	}
		himpp3.Mutex.Unlock()
		fmt.Println("done!")
    	})

	//fs := http.FileServer(http.Dir("static/"))
	//http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":80", nil)
}
