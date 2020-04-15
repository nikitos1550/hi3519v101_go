//+build !ignore, !generate

package ko

import (
	"golang.org/x/sys/unix"
	"strings"
	"application/pkg/logger"
    //"log"
    //"time"
)

var (
	//CmosName string   
	//MemMpp   uint64   = 12
	//MemLinux uint64   = 20
	//MemTotal uint64   = 32
	//DDRStartAddr uint64 = 0x80000000
	//ddrStartAddr uint64 = 0x40000000
	//Chip     string = "hi3516ev200"

	//Params	map[string]string //TODO find proper name, as params is already param string...
	Params		Parameters
)

func init() {
	Params = make(Parameters)
}

/*
type Parameters map[string]string

func (p Parameters) AddParamStr(name string, value string) { //param will be passed as is
	Params[name] = value
}
//func (p Parameters) AddParamMemMBSize(name string, value uint64) { //param will be formatted as memory size in MB
//	Params[name] = strconv.FormatUint(value/(1024*1024), 10) + "M"
//}
func (p Parameters) AddParamUint64Hex(name string, prefix string, value uint64, suffix string) { //param will be formatted as hex value 0x0...
	Params[name] = prefix + strconv.FormatUint(value, 16) + suffix
}
func (p Parameters) AddParamUint64(name string, prefix string, value uint64, suffix string) { //param will be srconved
	Params[name] = prefix + strconv.FormatUint(value, 10) + suffix
}
*/

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
	
	if len(tmpModules) < len(MinimalModulesList) {
		logger.Log.Warn().
			Msg("Not all modules from minimal list were found in overall list")
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
			if err.Error() != "no such file or directory" { //TODO comapre code!
				logger.Log.Warn().
				Str("name", modules[i][0]).
				Str("error", err.Error()).
				Msg("ko module removing error")
			} else {
                                logger.Log.Trace().
                                Str("name", modules[i][0]).
                                Str("error", err.Error()).
                                Msg("ko module removing error")
			}
		} else {
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
		//var err error
		data, err := Asset(modules[i][0])
		if err != nil {
			logger.Log.Error().
				Str("name", modules[i][0]).
				Str("error", err.Error()).
				Msg("ko module asset")
			continue
		}

		//TODO rework, remove err2
        //log.Println("params", modules[i][1])
		err2 := unix.InitModule(data, modules[i][1])
		if err2 != nil {
			logger.Log.Error().
				Str("name", modules[i][0]).
				Str("params", modules[i][1]).
				Str("error", err2.Error()).
				Msg("ko module load error")
			//return
		} else {
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
    /*
	var memStartAddr uint64 = DDRStartAddr + (uint64(MemLinux) * 1024 * 1024)
	var memMpp2 uint64 = MemTotal - MemLinux

	if memMpp2 != MemMpp {
		logger.Log.Panic().
			Uint64("mpp", MemMpp).
			Uint64("mpp2", memMpp2).
			Msg("Incorrect mpp memory size")
	}
    */

	for i := 0; i < len(modules); i++ {
		for param, value := range Params {
			modules[i][1] = strings.Replace(modules[i][1], "{" + param + "}", string(*value), -1)
		}

		//TODO replace by Params
        /*
		modules[i][1] = strings.Replace(modules[i][1], "{cmosName}", CmosName, -1)
		modules[i][1] = strings.Replace(modules[i][1], "{memStartAddr}", strconv.FormatUint(memStartAddr, 16), -1)
		modules[i][1] = strings.Replace(modules[i][1], "{memMppSize}", strconv.FormatUint(MemMpp, 10), -1)
		modules[i][1] = strings.Replace(modules[i][1], "{memTotalSize}", strconv.FormatUint(MemTotal, 10), -1)
		modules[i][1] = strings.Replace(modules[i][1], "{chipName}", Chip, -1)
        */

		char := strings.Index(modules[i][1], "{")
                if char > -1 {
                        logger.Log.Warn().
                                Str("module", modules[i][0]).
                                Str("params", modules[i][1]).
                                Msg("Not all vars are setuped")
                }

	}

	//Check that all vars in params are setuped
	/*
	for i := 0; i < len(modules); i++ {
		char := strings.Index(modules[i][1], "{")
		if char > -1 {
			logger.Log.Warn().
				Str("module", module[i][0]).
				Str("params". module[i][1]).
				Msg("Not all vars are setuped")
		}
	}
	*/
}
