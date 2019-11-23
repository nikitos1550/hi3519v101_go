//+build !ignore, !generate

package koloader

import (
    "log"
    "golang.org/x/sys/unix"
    "strings"
    "strconv"
)

//TEMPORARY
var (
    memMppSize      uint = 256
    memLinux        uint = 256
    memTotal        uint = 512
    chip            string = "hi3519v101"
)

func LoadMinimal() {
    //loadKo(Modules[][0])
    LoadAll() //TODO
}

func LoadAll() {
    tmpModules := make([][2]string, len(Modules))
    copy(tmpModules, Modules[:])

    setupKoParams(tmpModules)

    loadKo(Modules[:])
}

func loadKo(modules [][2]string) {
    //log.Println("Embedded files: ", AssetNames())

    //setupKoParams(modules[:])

    for i := len(modules)-1; i>=0; i-- {
        rmname := modules[i][0][0:len(modules[i][0])-3]
        err := unix.DeleteModule(rmname, 0)
        if err != nil {
            log.Println("Rmmod ", modules[i][0], " error ", err)
        }
    }

    for i := 0; i<len(modules); i++ {
        data, err := Asset(modules[i][0])
        if err != nil {
            log.Println(modules[i][0], " not found!")
            continue
        }
  
        err2 := unix.InitModule(data, modules[i][1])
        if err2 != nil {
            log.Println(modules[i][0], " error (", err2, ") loading!")
            return
        }
    }
}


func setupKoParams(modules [][2]string) {
    var memStartAddr uint64 = 0x80000000 + (uint64(memLinux)*1024*1024)
    var memMppSize uint64 = uint64(memTotal - memLinux)

    for i:=0; i<len(modules); i++ {
        Modules[i][1] = strings.Replace(Modules[i][1], "{memStartAddr}",    strconv.FormatUint(memStartAddr, 16),       -1)
        Modules[i][1] = strings.Replace(Modules[i][1], "{memMppSize}",      strconv.FormatUint(memMppSize, 10),         -1)
        Modules[i][1] = strings.Replace(Modules[i][1], "{memTotalSize}",    strconv.FormatUint(uint64(memTotal), 10),   -1)
        Modules[i][1] = strings.Replace(Modules[i][1], "{chipName}",        chip,                                       -1)
        log.Println(modules[i][1])
    }
}
