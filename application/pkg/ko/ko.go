//+build !ignore, !generate

package ko

import (
	"golang.org/x/sys/unix"
	//"log"
	"strconv"
	"strings"
	//"time"
	"application/pkg/logger"
)

var (
	MemMpp   uint64   = 12
	MemLinux uint64   = 20
	MemTotal uint64   = 32
	chip     string = "hi3516ev200"
)

func LoadMinimal() {
	tmpModules := make([][2]string, 0)
	for i := 0; i < len(ModulesList); i++ {
		for _, module := range MinimalModulesList {
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
	for i := len(modules) - 1; i >= 0; i-- {
		rmname := modules[i][0][0 : len(modules[i][0])-3]
		err := unix.DeleteModule(rmname, 0)
		if err != nil {
			//log.Println(modules[i][0], " removing error ", err)
			logger.Log.Warn().
				Str("name", modules[i][0]).
				Str("error", err.Error()).
				Msg("ko module removing error")
		} else {
			//log.Println(modules[i][0], " removed")
			logger.Log.Trace().
				Str("name", modules[i][0]).
				Msg("ko module removed")
		}
		//time.Sleep(1 * time.Second)
	}
}

func load(modules [][2]string) {
	setupParams(modules[:]) //TODO move to prev stack level

	for i := 0; i < len(modules); i++ {
		data, err := Asset(modules[i][0])
		if err != nil {
			//log.Println(modules[i][0], " not found (", err, ")!")
			logger.Log.Error().
				Str("name", modules[i][0]).
				Str("error", err.Error()).
				Msg("ko module asset")
			continue
		}

		//log.Println("Loading ", modules[i][0], " ", modules[i][1])
		//TODO rework, remove err2
		err2 := unix.InitModule(data, modules[i][1])
		if err2 != nil {
			//log.Println(modules[i][0], " loading error ", err2)
			logger.Log.Error().
				Str("name", modules[i][0]).
				Str("params", modules[i][1]).
				Str("error", err2.Error()).
				Msg("ko module load error")
			return
		} else {
			//log.Println(modules[i][0], " loaded")
			logger.Log.Trace().
				Str("name", modules[i][0]).
				Str("params", modules[i][1]).
				Msg("ko module loaded")
		}
		//time.Sleep(1 * time.Second)
	}
	//time.Sleep(5 * time.Second)
	//log.Println("Seems all modules loaded")
}

func setupParams(modules [][2]string) {
	var memStartAddr uint64 = 0x80000000 + (uint64(MemLinux) * 1024 * 1024)
	var memMpp2 uint64 = MemTotal - MemLinux

	if memMpp2 != MemMpp {
		//panic("Incorrect mpp memory size")
		logger.Log.Panic().
			Uint64("mpp", MemMpp).
			Uint64("mpp2", memMpp2).
			Msg("Incorrect mpp memory size")
	}

	for i := 0; i < len(modules); i++ {
		modules[i][1] = strings.Replace(modules[i][1], "{memStartAddr}", strconv.FormatUint(memStartAddr, 16), -1)
		modules[i][1] = strings.Replace(modules[i][1], "{memMppSize}", strconv.FormatUint(MemMpp, 10), -1)
		modules[i][1] = strings.Replace(modules[i][1], "{memTotalSize}", strconv.FormatUint(MemTotal, 10), -1)
		modules[i][1] = strings.Replace(modules[i][1], "{chipName}", chip, -1)

		//log.Println(modules[i][0], " prepared options ", modules[i][1])
	}
}
