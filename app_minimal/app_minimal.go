package main

import (
    "fmt"
	"flag"
	"net/http"
)

func main() {
	portPtr := flag.Int("port", 8080, "http port")

	flag.Parse()

	fmt.Printf("minimal application\n");
	fmt.Printf("http port %d\n", *portPtr);

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})

	//fs := http.FileServer(http.Dir("static/"))
	//http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}
