//+build processing

package processing

import "C"

import (
	"unsafe"
	"application/pkg/common"
)

type yuv struct {
	Name string
	Id int
}

func (p yuv) GetName() string {
	return p.Name
}

func (p yuv) GetId() int {
	return p.Id
}

func (p yuv) Create(id int, params map[string][]string) (common.Processing,int,string) {
	var v yuv
	v.Name = "yuv"
	v.Id = id
	return v,id,""
}

func (p yuv) Destroy() {
}

func (p yuv) Callback(data unsafe.Pointer) {
	sendDataToEncoders(p.Id, []byte("yuv"))
}

func init() {
	var v yuv
	v.Name = "yuv"
	v.Id = -1
	register(v)
}
