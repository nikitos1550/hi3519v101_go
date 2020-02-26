package cmos

import (
    "fmt"
    "unsafe"
)

func PrintInfo() {
    fmt.Println("Here will be list of supported cmoses soon...")
}


type cmos struct {


    registerCallback    unsafe.Pointer

}
