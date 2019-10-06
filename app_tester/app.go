package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "golang.org/x/sys/unix"
    "flag"
    "strings"
    "strconv"
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

    schema.Family           = chipFamily
    schema.Chip             = chip
    schema.ChipDetectedReg  = detectChip(sysIdReg)
    schema.ChipDetectedMpp  = detectChip(chipId())

    schema.Mpp              = mppVersion

    schema.SysIdReg         = sysIdReg
    schema.SysIdMpp         = chipId()

    schema.Temp             = getTemperature()

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    schemaJson, _ := json.Marshal(schema)
    fmt.Fprintf(w, "%s", string(schemaJson))
}

func main() {
    fmt.Println("app_tester, ", chipFamily, ", ", mppVersion)

    flag.StringVar  (&chip,     "chip",        chips[0],    "Chip name")
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
    fmt.Println("Chip ", chip)
    fmt.Println("Total board RAM ", memTotal, "MB")
    fmt.Println("Linux RAM ", memLinux, "MB")

    fmt.Println("Loading modules...")
    loadKo()

    fmt.Println("Initing temperature...")
    initTemperature()

    fmt.Println("starting http server :80")
    http.HandleFunc("/", apiHandler)
    http.ListenAndServe(":80", nil)
}


func loadKo() {
    fmt.Println("Embedded files: ", AssetNames())

    setupKoParams()

    for i := len(modules)-1; i>=0; i-- {
        err := unix.DeleteModule(modules[i][0], 0)
        if err != nil {
            fmt.Println("Rmmod ", modules[i][0], " error ", err)
        }
    }

    for i := 0; i<len(modules); i++ {
        data, err := Asset(modules[i][0])
        if err != nil {
            fmt.Println(modules[i][0], " not found!")
            continue
        }
        err2 := unix.InitModule(data, modules[i][1])
        if err2 != nil {
            fmt.Println(modules[i][0], " error (", err2, ") loading!")
            return
        }
    }
}

func setupKoParams() {
    var memStartAddr uint64 = 0x80000000 + (uint64(memLinux)*1024*1024)
    var memMppSize uint64 = uint64(memTotal - memLinux)

    for i:=0; i<len(modules); i++ {
        modules[i][1] = strings.Replace(modules[i][1], "{memStartAddr}",    strconv.FormatUint(memStartAddr, 16),       -1)
        modules[i][1] = strings.Replace(modules[i][1], "{memMppSize}",      strconv.FormatUint(memMppSize, 10),         -1)
        modules[i][1] = strings.Replace(modules[i][1], "{memTotalSize}",    strconv.FormatUint(uint64(memTotal), 10),   -1)
        modules[i][1] = strings.Replace(modules[i][1], "{chipName}",        chip,                                       -1)
        fmt.Println(modules[i][1])
    }
}
