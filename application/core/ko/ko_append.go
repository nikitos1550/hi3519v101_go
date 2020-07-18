//+build !ignore,!generate
//+build koAppend

package ko

import (
	"os"
    "strconv"
	//"syscall"
    "application/core/logger"
    "log"
    "golang.org/x/sys/unix"
)

var (
    ElfSizeStr string  //10 symbols string with final binary size, use as offest to get ko modules appended bin data
    elfSizeInt int64
)


func init() {
    var err error
    elfSizeInt, err = strconv.ParseInt(ElfSizeStr, 10, 64)
    if err != nil {
        log.Fatal("Some problem with appended modules")
        //logger.Log.Fatal().
        //    Msg("Some problem with appended modules")
    }
    log.Println("Ko append mode, elfsize = ", elfSizeInt)
    //logger.Log.Trace().
    //    Int64("elf", elfSizeUint).
    //    Msg("Ko append mode")
}


func loadModule(name string, params string) error {

    logger.Log.Trace().
        Uint("offset", ModulesInfo[name][0]).
        Uint("size", ModulesInfo[name][1]).
        Msg("Test loading modules")


    f, err := os.OpenFile("/opt/gohisicam", os.O_RDONLY, 0755)
    if err != nil {
		//log.Fatal(err)
        logger.Log.Fatal().
            Msg("Self file open error")
	}
	fd := f.Fd()

    logger.Log.Trace().
        Int("fd", int(fd)).
        Msg("File opened")

    _, err = f.Seek(elfSizeInt+int64(ModulesInfo[name][0]), 0)
    if err != nil {
        panic(err)
    }

    data := make([]byte, int(ModulesInfo[name][1]))
    _, err = f.Read(data)
    if err != nil {
        panic(err)
    }

    //log.Println("mmpaing offset ", elfSizeInt+int64(ModulesInfo[name][0]), " length ", int(ModulesInfo[name][1]))

	//data, err := syscall.Mmap(int(fd), elfSizeInt+int64(ModulesInfo[name][0]), int(ModulesInfo[name][1]), syscall.PROT_READ, syscall.MAP_SHARED)
    //data, err := syscall.Mmap(int(fd), 0, int(ModulesInfo[name][1]), syscall.PROT_READ, syscall.MAP_SHARED)
	//if err != nil {
	//	panic(err)
	//}
    
    log.Println("len(data)=", len(data), " cap(data)=", cap(data))
	//b[offset-1] = 'x'

    err = unix.InitModule(data, params)
    if err != nil {
        logger.Log.Error().
            Str("module", name).
            Str("params", params).
            Str("desc", err.Error()).
            Msg("KO load")
        return err
    } else {
        //logger.Log.Trace().
        //    Str("module", name).
        //    Str("params", params).
        //    Msg("KO loaded")
    }


	//err = syscall.Munmap(data)
	//if err != nil {
	//	panic(err)
	//}
    

    f.Close()

    return nil
}
