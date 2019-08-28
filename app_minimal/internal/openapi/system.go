package openapi

import (
    "log"
    "fmt"
    "net/http"
    "encoding/json"
    "time"
    "../himpp3"
    "../info"
    "net"
)

func system (w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    log.Println("system")
    
    fmt.Fprintf(w, "Camera go webserver!\n")
    fmt.Fprintf(w, "BuildTime %s\n", info.DATE)
    fmt.Fprintf(w, "BuildBranch %s\n", info.BRANCH)
    fmt.Fprintf(w, "BuildUser %s\n", info.USER)
    fmt.Fprintf(w, "ChipFamily %s\n", himpp3.GetChipFamily())
    fmt.Fprintf(w, "Chip %s\n", himpp3.GetChip())
    fmt.Fprintf(w, "CMOS %s\n", himpp3.GetCMOS())
}

type systemTemperatureResponce struct {
    TemperatureC float32 `json:"temperature_c,omitempty"`
    TemperatureF float32 `json:"temperature_f,omitempty"`
    TemperatureK float32 `json:"temperature_k,omitempty"`
}

func systemTemperature (w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    var tmp systemTemperatureResponce
    tmp.TemperatureC = float32(int(himpp3.TempGet() * 10)) / 10
    tmp.TemperatureF = float32(int(((tmp.TemperatureC * 9/5) + 32) * 10)) / 10
    tmp.TemperatureK = float32(int((tmp.TemperatureC + 273.15) * 10)) / 10

    log.Println("systemTemperature")
    test, _ := json.Marshal(tmp)
    fmt.Fprintf(w, "%s", string(test))
}

type systemDateResponce struct {
    Formatted time.Time `json:"formatted,omitempty"`
    Secs int64 `json:"secs,omitempty"`
    Nanosecs int64 `json:"nanosecs,omitempty"`
}

func systemDate (w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    var tmp systemDateResponce

    t := time.Now()

    tmp.Formatted = t
    tmp.Secs = t.Unix()
    tmp.Nanosecs = t.UnixNano()

    log.Println("systemDate")
    test, _ := json.Marshal(tmp)
    fmt.Fprintf(w, "%s", string(test))
}

func systemNetwork (w http.ResponseWriter, r *http.Request) {
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

