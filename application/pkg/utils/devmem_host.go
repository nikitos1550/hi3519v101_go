//+build 386 amd64
//+build host

package utils

import (
    //"fmt"
    //"golang.org/x/sys/unix"
    //"os"
    //"unsafe"
)

func WriteDevMem32(target, value uint32) {
    panic("This func can`t be invoked in host build!")
}

func ReadDevMem32(target uint32) uint32 {
    panic("This func can`t be invoked in host build!")
    return 0
}

