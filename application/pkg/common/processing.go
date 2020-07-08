//+build nobuild

package common

import (
	"unsafe"
)

type Processing interface {
	GetName() string
	GetId() int
	Create(id int, params map[string][]string) (Processing,int,string)
	Destroy()
	Callback(unsafe.Pointer)
}
