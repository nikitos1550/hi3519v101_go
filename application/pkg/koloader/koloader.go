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
    loadKo()
}

/*
func prepareOptions(map) (modules[][2], error) {
    

}
*/


func loadKo() {
    log.Println("Embedded files: ", AssetNames())

    //setupKoParams()

    for i := len(Modules)-1; i>=0; i-- {
        err := unix.DeleteModule(Modules[i][0], 0)
        if err != nil {
            log.Println("Rmmod ", Modules[i][0], " error ", err)
        }
    }

    for i := 0; i<len(Modules); i++ {
        data, err := Asset(Modules[i][0])
        if err != nil {
            log.Println(Modules[i][0], " not found!")
            continue
        }
  
        err2 := unix.InitModule(data, Modules[i][1])
        if err2 != nil {
            log.Println(Modules[i][0], " error (", err2, ") loading!")
            return
        }
    }
}


func setupKoParams() {
    var memStartAddr uint64 = 0x80000000 + (uint64(memLinux)*1024*1024)
    var memMppSize uint64 = uint64(memTotal - memLinux)

    for i:=0; i<len(Modules); i++ {
        Modules[i][1] = strings.Replace(Modules[i][1], "{memStartAddr}",    strconv.FormatUint(memStartAddr, 16),       -1)
        Modules[i][1] = strings.Replace(Modules[i][1], "{memMppSize}",      strconv.FormatUint(memMppSize, 10),         -1)
        Modules[i][1] = strings.Replace(Modules[i][1], "{memTotalSize}",    strconv.FormatUint(uint64(memTotal), 10),   -1)
        Modules[i][1] = strings.Replace(Modules[i][1], "{chipName}",        chip,                                       -1)
        log.Println(Modules[i][1])
    }
}
