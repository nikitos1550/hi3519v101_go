package common

import (
	"unsafe"
)

type Processing interface {
	GetName() string
	GetId() int
	Create(id int) Processing
	Init() 
	Callback(unsafe.Pointer)
}

func init() {
}

func Init() {
}