package koloader

import (
    _"fmt"
    _"log"
    _"golang.org/x/sys/unix"
)

/*
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


*/
