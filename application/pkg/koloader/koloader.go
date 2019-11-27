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
    tmpModules := make([][2]string, 0)
    for i := 0; i < len(ModulesList); i++ {
        for _, module := range minimalModulesList {
            if ModulesList[i][0] == module {
                //log.Println("Found ", module)
                tmpModules = append(tmpModules, ModulesList[i])
            }
        }
    }
    /*
    for _, module := range minimalModulesList {
        for i := 0; i < len(ModulesList); i++ {
            if ModulesList[i][0] == module {
                //log.Println("Found ", module)
                tmpModules = append(tmpModules, ModulesList[i])
            }
        }
    }
    */
    load(tmpModules[:])
}

func LoadAll() {
    tmpModules := make([][2]string, len(ModulesList))
    copy(tmpModules, ModulesList[:])

    load(tmpModules[:])
}

//TODO create list by names (order by orig)
/*
func Load(names []string) {


}
*/

func UnloadAll() {
    tmpModules := make([][2]string, len(ModulesList))
    copy(tmpModules, ModulesList[:])

    unload(tmpModules[:])
}

func unload(modules [][2]string) {
    for i := len(modules)-1; i>=0; i-- {
        rmname := modules[i][0][0:len(modules[i][0])-3]
        err := unix.DeleteModule(rmname, 0)
        if err != nil {
            log.Println(modules[i][0], " removing error ", err)
        } else {
            log.Println(modules[i][0], " removed")
        }
    }
}

func load(modules [][2]string) {
    setupParams(modules[:]) //TODO move to prev stack level

    for i := 0; i<len(modules); i++ {
        data, err := Asset(modules[i][0])
        if err != nil {
            log.Println(modules[i][0], " not found!")
            continue
        }

        //log.Println("Loading ", modules[i][0], " ", modules[i][1])
        err2 := unix.InitModule(data, modules[i][1])
        if err2 != nil {
            log.Println(modules[i][0], " loading error ", err2)
            return
        } else {
            log.Println(modules[i][0], " loaded")
        }
    }
}

func setupParams(modules [][2]string) {
    var memStartAddr uint64 = 0x80000000 + (uint64(memLinux)*1024*1024)
    var memMppSize uint64 = uint64(memTotal - memLinux)

    for i:=0; i<len(modules); i++ {
        modules[i][1] = strings.Replace(modules[i][1], "{memStartAddr}",    strconv.FormatUint(memStartAddr, 16),       -1)
        modules[i][1] = strings.Replace(modules[i][1], "{memMppSize}",      strconv.FormatUint(memMppSize, 10),         -1)
        modules[i][1] = strings.Replace(modules[i][1], "{memTotalSize}",    strconv.FormatUint(uint64(memTotal), 10),   -1)
        modules[i][1] = strings.Replace(modules[i][1], "{chipName}",        chip,                                       -1)

        //log.Println(modules[i][0], " prepared options ", modules[i][1])
    }
}
