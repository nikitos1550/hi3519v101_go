// +build lua

package scripts

import (
	"github.com/yuin/gopher-lua"
    "flag"
)

func init() {
    //    fmt.Println("Hello from openapi Init")
    test = flag.Int("lua1", 8888, "path to default scripts")
    test = flag.Int("lua2", 8888, "path to user scripts")


	L := lua.NewState()
	//defer L.Close()
	if err := L.DoString(`print("hello from LUA\n")`); err != nil {
		panic(err)
	}
	L.Close()
}
