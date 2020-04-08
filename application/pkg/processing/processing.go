package processing

import "C"

import (
	"log"
	"unsafe"
)

var (
	Processings map[string] unsafe.Pointer
)

func Init() {
	log.Println("main processing Init", len(Processings))
}

func init() {
	Processings = make(map[string] unsafe.Pointer)
	log.Println("main processing init")
}

func register(name string, callback unsafe.Pointer){
	_, exists := Processings[name]
	if (exists) {
		log.Fatal("processing already exists", name)
	}
	
	Processings[name] = callback
}
