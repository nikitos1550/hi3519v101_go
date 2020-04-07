package processing

import "C"

import (
	"log"
	"unsafe"
)

var (
	processings map[string] unsafe.Pointer
)

func Init() {
	log.Println("main processing Init", len(processings))
}

func init() {
	processings = make(map[string] unsafe.Pointer)
	log.Println("main processing init")
}

func register(name string, callback unsafe.Pointer){
	_, exists := processings[name]
	if (exists) {
		log.Fatal("processing already exists", name)
	}
	
	processings[name] = callback
}
