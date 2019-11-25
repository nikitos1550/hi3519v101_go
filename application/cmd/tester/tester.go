package main

import (
    "log"
    "fmt"
    "net/http"
    "encoding/json"
    "flag"
    _"strconv"
    "application/pkg/koloader"
    "application/pkg/utils/temperature"
    "application/pkg/utils/chip"
    "application/pkg/mpp"
    "application/pkg/buildinfo"
)

type Answer struct {
    App             string      `json:"appName"`

    ChipDetectedReg string      `json:"chipDetectedReg"`
    ChipDetectedMpp string      `json:"chipDetectedMpp"`

    Mpp             string      `json:"mppVersion"`

    SysIdReg        uint32      `json:"chipIdReg"`
    SysIdMpp        uint32      `json:"chipIdMpp"`

    TempVal         float32     `json:"temperature"`
    TempHW          string      `json:"temperatureHW"`

    buildinfo       BuildInfo   `json:"buildInfo"`
}

type BuildInfo struct {
    GoVersion   string  `json:"goVersion"`
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

    var err error
    schema.TempVal, err = temperature.Get()
    //log.Println("temperature ", temperature, " error ", err)
    if (err != nil) {
        schema.TempHW = "not availible"
    } else {
        schema.TempHW = "availible"
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    schemaJson, _ := json.Marshal(schema)
    fmt.Fprintf(w, "%s", string(schemaJson))
}

func main() {
    log.Println("tester")

    flag.UintVar    (&memTotal, "memtotal", 512, "Total RAM size, MB")
    flag.UintVar    (&memLinux, "memlinux", 256, "RAM size passed to Linux kernel, rest will be used for MPP, MB")

    flag.Parse()

    log.Println("CMD parsed params:")
    log.Println("Total board RAM ", memTotal, "MB")
    log.Println("Linux RAM ", memLinux, "MB")
    log.Println("")

    log.Println("Build time info:")
    log.Println("Go: ",         buildinfo.GoVersion)
    log.Println("Gcc: ",        buildinfo.GccVersion)
    log.Println("Date: ",       buildinfo.BuildDateTime)
    log.Println("Tags: ",       buildinfo.BuildTags)
    log.Println("User: ",       buildinfo.BuildUser)
    log.Println("Commit: ",     buildinfo.BuildCommit)
    log.Println("Branch: ",     buildinfo.BuildBranch)
    log.Println("Vendor: ",     buildinfo.BoardVendor)
    log.Println("Model: ",      buildinfo.BoardModel)
    log.Println("Chip: ",       buildinfo.Chip)
    log.Println("Cmos: ",       buildinfo.CmosProfile)
    log.Println("Total ram: ",  buildinfo.TotalRam)
    log.Println("Linux ram: ",  buildinfo.LinuxRam)
    log.Println("Mpp ram: ",    buildinfo.MppRam)
    log.Println("")

    log.Println("Loading modules...")
    koloader.LoadMinimal()
    log.Println("Loading modules done")

    log.Println("Initing temperature...")
    temperature.Init()
    log.Println("Initing temperature done")

    log.Println("Starting http server :80")
    http.HandleFunc("/", apiHandler)
    http.ListenAndServe(":80", nil)
}
