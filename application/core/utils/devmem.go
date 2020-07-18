//+build arm
//+build hi3516av100 hi3516av200 hi3516cv100 hi3516cv200 hi3516cv300 hi3516cv500 hi3516ev200 hi3519av100 hi3559av100

package utils

import (
    //"fmt"
    "golang.org/x/sys/unix"
    "os"
    "unsafe"

    "application/core/logger"
)

func WriteDevMem32(target, value uint32) {
    //logger.Log.Trace().
    //    Str("addr", fmt.Sprintf("0x%08X", target)).
	//    Str("value", fmt.Sprintf("0x%08X", value)).
	//    Msg("WriteDevMem32")

    const pageSize int = 4096

    var mapSize int = pageSize

    var offset int64 = int64(target & uint32(pageSize-1))

    if (offset + 4) > int64(pageSize) {
        mapSize = mapSize + pageSize
    }

    var mapOffset int64 = int64(target & ^uint32(pageSize-1))

    file, err := os.OpenFile("/dev/mem", os.O_RDWR, 0)
    if err != nil {
	    logger.Log.Fatal().
		    Str("error", err.Error()).
		    Msg("/dev/mem open error")
        return
    }
    defer file.Close()

    mmap, err := unix.Mmap(int(file.Fd()), mapOffset, mapSize, unix.PROT_READ | unix.PROT_WRITE, unix.MAP_SHARED)
    if err != nil {
	    logger.Log.Fatal().
		    Str("error", err.Error()).
		    Msg("MMAP error")
        return
    }

    *(*uint32)(unsafe.Pointer(&mmap[offset])) = value

    err = unix.Munmap(mmap)
    if err != nil {
	    logger.Log.Fatal().
		    Str("error", err.Error()).
		    Msg("MUNMAP error")
        return
    }
}

func ReadDevMem32(target uint32) uint32 {
    var value uint32 = 0xFFFFFFFF   //TODO

    const pageSize int = 4096

    var mapSize int = pageSize

    var offset int64 = int64(target & uint32(pageSize-1))

    if (offset + 4) > int64(pageSize) {
        mapSize = mapSize + pageSize
    }

    var mapOffset int64 = int64(target & ^uint32(pageSize-1))

    file, err := os.Open("/dev/mem")
    if err != nil {
		logger.Log.Fatal().
            Str("error", err.Error()).
		    Msg("/dev/mem open error")
        return value
    }
    defer file.Close()

    mmap, err := unix.Mmap(int(file.Fd()), mapOffset, mapSize, unix.PROT_READ, unix.MAP_SHARED)
    if err != nil {
        logger.Log.Fatal().
            Str("error", err.Error()).
		    Msg("MMAP error")
        return value
    }

    value = *(*uint32)(unsafe.Pointer(&mmap[offset]))

    err = unix.Munmap(mmap)
    if err != nil {
	    logger.Log.Fatal().
            Str("error", err.Error()).
            Msg("MUNMAP error")
        return value
    }

    //logger.Log.Trace().
    //    Str("addr", fmt.Sprintf("0x%08X", target)).
    //    Uint32("value", value).
    //    Msg("ReadDevMem32")

    return value
}

