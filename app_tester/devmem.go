package main

import (
    "fmt"
    "golang.org/x/sys/unix"
    "os"
    "unsafe"
)


func writeDevMem32(target, value uint32) {
    //var valueOld uint32 = 0xFFFFFFFF   //todo

    const pageSize int = 4096

    var mapSize int = pageSize

    //fmt.Printf("target %X\n", target)

    var offset int64 = int64(target & uint32(pageSize-1))
    //fmt.Printf("offset %X\n", offset)

    if (offset + 4) > int64(pageSize) {
        mapSize = mapSize + pageSize
    }
    //fmt.Printf("mapSize %X\n", mapSize)

    var mapOffset int64 = int64(target & ^uint32(pageSize-1))
    //fmt.Printf("mapOffset %X\n", mapOffset)

    //file, err := os.Open("/dev/mem")
    file, err := os.OpenFile("/dev/mem", os.O_RDWR, 0)
    if err != nil {
        fmt.Println("OpenFile error ", err)
        return
    }
    defer file.Close()

    mmap, err := unix.Mmap(int(file.Fd()), mapOffset, mapSize, unix.PROT_READ | unix.PROT_WRITE, unix.MAP_SHARED)
    if err != nil {
        fmt.Println("Mmap error ", err)
        return
    }

    //valueOld = *(*uint32)(unsafe.Pointer(&mmap[offset]))
    //fmt.Printf("old value %X\n", valueOld)

    *(*uint32)(unsafe.Pointer(&mmap[offset])) = value
    //mmap[offset+0] = 0
    //mmap[offset+1] = 1
    //mmap[offset+2] = 2
    //mmap[offset+3] = 3

    //valueOld = *(*uint32)(unsafe.Pointer(&mmap[offset]))
    //fmt.Printf("new value %X\n", valueOld)


    err = unix.Munmap(mmap)
    if err != nil {
        fmt.Println(err)
        return
    }
}

func readDevMem32(target uint32) uint32 {
    var value uint32 = 0xFFFFFFFF   //todo

    const pageSize int = 4096

    var mapSize int = pageSize

    //fmt.Printf("target %X\n", target)

    var offset int64 = int64(target & uint32(pageSize-1))
    //fmt.Printf("offset %X\n", offset)

    if (offset + 4) > int64(pageSize) {
        mapSize = mapSize + pageSize
    }
    //fmt.Printf("mapSize %X\n", mapSize)

    var mapOffset int64 = int64(target & ^uint32(pageSize-1))
    //fmt.Printf("mapOffset %X\n", mapOffset)

    file, err := os.Open("/dev/mem")
    if err != nil {
        fmt.Println(err)
        return value
    }
    defer file.Close()

    mmap, err := unix.Mmap(int(file.Fd()), mapOffset, mapSize, unix.PROT_READ, unix.MAP_SHARED)
    if err != nil {
        fmt.Println(err)
        return value
    }

    value = *(*uint32)(unsafe.Pointer(&mmap[offset]))
    //fmt.Printf("value %X\n", value)

    err = unix.Munmap(mmap)
    if err != nil {
        fmt.Println(err)
        return value
    }

    return value
}

