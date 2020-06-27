//+build processing

package processing

import "C"

import (
	"unsafe"
	"application/pkg/common"
)

type proxy struct {
	Name string
	Id int
}

func (p proxy) GetName() string {
	return p.Name
}

func (p proxy) GetId() int {
	return p.Id
}

func (p proxy) Create(id int, params map[string][]string) (common.Processing,int,string) {
	var v proxy
	v.Name = "proxy"
	v.Id = id
	return v,id,""
}

func (p proxy) Destroy() {
}

func (p proxy) Callback(data unsafe.Pointer) {
	sendToEncoders(p.Id, data)
}

func init() {
	var v proxy
	v.Name = "proxy"
	v.Id = -1
	register(v)
}
