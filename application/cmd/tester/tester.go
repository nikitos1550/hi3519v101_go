package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "flag"
    "application/pkg/koloader"
    "application/pkg/utils/temperature"
    "application/pkg/utils/chipid"
    "application/pkg/mpp"
)

var (
    memTotal    uint
    memLinux    uint
    chip        string
)

type Answer struct {
    App             string  `json:"appName"`

    Family          string  `json:"family"`
    Chip            string  `json:"chipSetuped"`
    ChipDetectedReg string  `json:"chipDetectedReg"`
    ChipDetectedMpp string  `json:"chipDetectedMpp"`

    Mpp             string  `json:"mppVersion"`

    SysIdReg        uint32  `json:"chipIdReg"`
    SysIdMpp        uint32  `json:"chipIdMpp"`

    Temp            float32 `json:"temperature"`

    //Vendor  string  `json:"vendorName"`
    //Model   string  `json:"modelName"`
    //Chip    string  `json:"chip"`
    //Ram     uint    `json:"ram"`
    //Rom     uint    `json:"rom"`
    //Cmos    string  `json:"cmos"`
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
    var schema Answer

    schema.App              = "app_tester"

    //schema.Family           = chipFamily
    //schema.Chip             = chip
    schema.ChipDetectedReg  = chipid.Detect(chipid.Reg()) //detectChip(sysIdReg)
    schema.ChipDetectedMpp  = chipid.Detect(chipid.Mpp()) //detectChip(chipId())

    schema.Mpp              = mpp.Version()

    schema.SysIdReg         = chipid.Reg() //sysIdReg
    schema.SysIdMpp         = chipid.Mpp() //chipId()

    schema.Temp             = temperature.Get() //getTemperature()

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    schemaJson, _ := json.Marshal(schema)
    fmt.Fprintf(w, "%s", string(schemaJson))
}

func main() {
    fmt.Println("app_tester, ") //, chipFamily) //, ", ", mppVersion)

    //flag.StringVar  (&chip,     "chip",        chips[0],    "Chip name")
    flag.UintVar    (&memTotal, "memtotal",    512,         "Total RAM size, MB")
    flag.UintVar    (&memLinux, "memlinux",    256,         "RAM size passed to Linux kernel, rest will be used for MPP, MB")

    flag.Parse()

    /*
    isChipValid := false
    for i:=0; i<len(chips); i++ {
        if chips[i] == chip {
            isChipValid = true
            break
        }
    }
    if !isChipValid {
        fmt.Println("Unknown chip name")
        return
    }
    */

    fmt.Println("CMD parsed params:")
    //fmt.Println("Chip ", chip)
    fmt.Println("Total board RAM ", memTotal, "MB")
    fmt.Println("Linux RAM ", memLinux, "MB")

    fmt.Println("Loading modules...")
    //loadKo()
    koloader.LoadDefault()

    fmt.Println("Initing temperature...")
    //initTemperature()
    temperature.Init()

    fmt.Println("starting http server :80")
    http.HandleFunc("/", apiHandler)
    http.ListenAndServe(":80", nil)
}
