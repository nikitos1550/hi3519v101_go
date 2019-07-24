package main

import (
    	"fmt"
	"flag"
	"net/http"
	"./himpp3"
)

func main() {
	himpp3.SysInit()
	himpp3.MipiIspInit()
	himpp3.ViInit()
	himpp3.VpssInit()
	//himpp3.VencInit()
	
	

	portPtr := flag.Int("port", 80, "http port")

	flag.Parse()

	fmt.Printf("minimal application\n");
	fmt.Printf("http port %d\n", *portPtr);

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Camera go webserver!")
	})

	//fs := http.FileServer(http.Dir("static/"))
	//http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":80", nil)
}
