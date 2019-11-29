// +build scripts

package scripts

import (
	"github.com/yuin/gopher-lua"
    "flag"
)

var L *lua.LState

var flagInitScriptPath  *string

func init() {
    flagInitScriptPath = flag.String("scripts-init", "/opt/scripts/init.lua", "Lua init script file path")
}

//TODO
//func AddFunc(fname string, fimpl func(L *lua.LState) ) {
func AddFunc(fname string, fimpl lua.LGFunction) {
	L.SetGlobal(fname, L.NewFunction(fimpl)) /* Original lua_setglobal uses stack... */
}

func Init() {
    //    fmt.Println("Hello from openapi Init")
    //test = flag.Int("lua1", 8888, "path to default scripts")
    //test = flag.Int("lua2", 8888, "path to user scripts")


	//L = lua.NewState()
	/*
	L = lua.NewState(lua.Options{
		//RegistrySize: 1024 * 20,         // this is the initial size of the registry
		//RegistryMaxSize: 1024 * 80,      // this is the maximum size that the registry can grow to. If set to `0` (the default) then the registry will not auto grow
		//RegistryGrowStep: 32,            // this is how much to step up the registry by each time it runs out of space. The default is `32`. 
		//CallStackSize: 120,              // this is the maximum callstack size of this LState
		//MinimizeStackMemory: true,       // Defaults to `false` if not specified. If set, the callstack will auto grow and shrink as needed up to a max of `CallStackSize`. If not set, the callstack will be fixed at `CallStackSize`.
		//SkipOpenLibs: true,				 // 
		//IncludeGoStackTrace: false,		 //
	})
	*/
	L := lua.NewState(lua.Options{SkipOpenLibs: true})
    defer L.Close()
    for _, pair := range []struct {
        n string
        f lua.LGFunction
    }{
        {lua.LoadLibName, lua.OpenPackage}, // Must be first
        {lua.BaseLibName, lua.OpenBase},
        {lua.TabLibName, lua.OpenTable},
    } {
        if err := L.CallByParam(lua.P{
            Fn:      L.NewFunction(pair.f),
            NRet:    0,
            Protect: true,
        }, lua.LString(pair.n)); err != nil {
            panic(err)
        }
    }
	if err := L.DoString(`print("hello from LUA\n")`); err != nil {
		panic(err)
	}
	//L.Close()
}
