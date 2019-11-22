package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "flag"
    "application/pkg/koloader"
    "application/pkg/utils/temperature"
    "application/pkg/utils/chip"
    "application/pkg/mpp"
    "application/pkg/buildinfo"
)

type Answer struct {
    App             string  `json:"appName"`

    ChipDetectedReg string  `json:"chipDetectedReg"`
    ChipDetectedMpp string  `json:"chipDetectedMpp"`

    Mpp             string  `json:"mppVersion"`

    SysIdReg        uint32  `json:"chipIdReg"`
    SysIdMpp        uint32  `json:"chipIdMpp"`

    TempVal         float32 `json:"temperature"`
    TempNA          string  `json:"temperature"`
}

var (
    memTotal    uint
    memLinux    uint

    schema      Answer
)


func apiHandler(w http.ResponseWriter, r *http.Request) {
    //var schema Answer

    schema.App              = "tester"

    schema.ChipDetectedReg  = chip.Detect(chip.RegId())
    schema.ChipDetectedMpp  = chip.Detect(chip.MppId())

    schema.Mpp              = mpp.Version()

    schema.SysIdReg         = chip.RegId()
    schema.SysIdMpp         = chip.MppId()

    temperature, err := temperature.Get()
    if (err != nil) {
        schema.TempVal = temperature
    } else {
        schema.TempNA = "not availible"
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    schemaJson, _ := json.Marshal(schema)
    fmt.Fprintf(w, "%s", string(schemaJson))
}

func main() {
    fmt.Println("tester")

    flag.UintVar    (&memTotal, "memtotal",    512,         "Total RAM size, MB")
    flag.UintVar    (&memLinux, "memlinux",    256,         "RAM size passed to Linux kernel, rest will be used for MPP, MB")

    flag.Parse()

    fmt.Println("CMD parsed params:")
    fmt.Println("Total board RAM ", memTotal, "MB")
    fmt.Println("Linux RAM ", memLinux, "MB")
    fmt.Println("")

    fmt.Println("Build time info:")
    fmt.Println("Go: ", GoVersion)
    fmt.Println("Gcc: ", GccVersion)
    fmt.Println("Date: ", BuildDateTime)
    fmt.Println("Tags: ", BuildTags)
    fmt.Println("Vendor: ", BoardVendor)
    fmt.Println("Model: ", BoardModel)
    fmt.Println("Chip: ", Chip)
    fmt.Println("Cmos: ", CmosProfile)
    fmt.Println("Total ram: ", TotalRam)
    fmt.Println("Linux ram: ", LinuxRam)
    fmt.Println("Mpp ram: ", MppRam)
    fmt.Println("")

    fmt.Print("Loading modules...")
    koloader.LoadMinimal()
    fmt.Println(" done")

    fmt.Print("Initing temperature...")
    temperature.Init()
    fmt.Println(" done")

    fmt.Println("Starting http server :80")
    http.HandleFunc("/", apiHandler)
    http.ListenAndServe(":80", nil)
}
