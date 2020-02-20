// +build openapi

package openapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"net"
)

func init() {
    AddApiRoute("systemDate",      "/system/date",     "GET",      systemDate)
    AddApiRoute("systemNetwork",   "/system/network",  "GET",      systemNetwork)
	AddApiRoute("apiRoot",   "/",  "GET",      apiRoot)
}


type systemDateTimeSchema struct {
	Formatted time.Time `json:"formatted,omitempty"`
	Secs      int64     `json:"secs,omitempty"`
	Nanosecs  int64     `json:"nanosecs,omitempty"`
}

func systemDate(w http.ResponseWriter, r *http.Request) {
    log.Println("systemDate")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var dtSchema systemDateTimeSchema

	t := time.Now()

	dtSchema.Formatted = t
	dtSchema.Secs = t.Unix()
	dtSchema.Nanosecs = t.UnixNano()

	dtSchemaJson, _ := json.Marshal(dtSchema)
	fmt.Fprintf(w, "%s", string(dtSchemaJson))
}

func systemNetwork(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	ifaces, err := net.Interfaces()
	if err != nil {
		log.Printf("localAddresses: %+v\n", err.Error())
		return
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			log.Printf("localAddresses: %+v\n", err.Error())
			continue
		}
		for _, a := range addrs {
			switch v := a.(type) {
			case *net.IPAddr:
				fmt.Fprintf(w, "%v : %s (%s)\n", i.Name, v, v.IP.DefaultMask())

			case *net.IPNet:
				fmt.Fprintf(w, "%v : %s [%v/%v]\n", i.Name, v, v.IP, v.Mask)
			}

		}
	}
}

func apiRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "All available api:\n\n")
	fmt.Fprintf(w, "///////////////////////////////////////////////\n\n")

	for _, route := range apiRoutes {
		fmt.Fprintf(w, "Api name: %s \n", route.name)
		fmt.Fprintf(w, "Method: %s \n", route.method)
		fmt.Fprintf(w, "Address: %s \n", route.pattern)
		fmt.Fprintf(w, "-----------------------------------------------\n\n")
	}
}
