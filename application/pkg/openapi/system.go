// +build openapi

package openapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	//"net"
)

func init() {
    AddApiRoute("systemDate",      "/system/date",     "GET",      systemDate)
    //AddApiRoute("systemNetwork",   "/system/network",  "GET",      systemNetwork)
}


type apiAnswerSystemDateTimeSchema struct {
	Formatted time.Time `json:"formatted,omitempty"`
	Secs      int64     `json:"secs,omitempty"`
	Nanosecs  int64     `json:"nanosecs,omitempty"`
}

/**OPENAPI
/system/date:
  get:
    tags:
      - system
    summary: 'Get system date and time'
    operationId: 'systemDate'
    responses:
      '200':
        description: 'Success'
        content:
          application/json:
            schema:
              type: object
              properties:
                formatted:
                  type: string
                  format: date-time
                secs:
                  type: integer
                  format: uint64
                nanosecs:
                  type: integer
                  format: uint64
 */
/**
 * @api {get} /system/date Request DateTime information
 * @apiName GetDate
 * @apiGroup System
 *
 * @apiSuccess (200) {Date} formatted
 * @apiSuccess (200) {int} secs
 * @apiSuccess (200) {int} nanosecs
 */
func systemDate(w http.ResponseWriter, r *http.Request) {
    log.Println("systemDate")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var dtSchema apiAnswerSystemDateTimeSchema
	t := time.Now()

	dtSchema.Formatted = t
	dtSchema.Secs = t.Unix()
	dtSchema.Nanosecs = t.UnixNano()

	dtSchemaJson, _ := json.Marshal(dtSchema)
    //dtSchemaJson, _ := json.Marshal(actionSystemDate())
	fmt.Fprintf(w, "%s", string(dtSchemaJson))
}

/*
func actionSystemDate() *apiAnswerSystemDateTimeSchema {
    var dtSchema apiAnswerSystemDateTimeSchema
    t := time.Now()

    dtSchema.Formatted = t
    dtSchema.Secs = t.Unix()
    dtSchema.Nanosecs = t.UnixNano()

    return &dtSchema
}
*/
/*
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
*/
