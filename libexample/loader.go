package libexample

import (
	lua "github.com/yuin/gopher-lua"
)

func Preload(L *lua.LState) {
	L.PreloadModule("libexample", Loader)
}

// Loader is the module loader function.
func Loader(L *lua.LState) int {
	t := L.NewTable()
	L.SetFuncs(t, api)
	L.Push(t)
	return 1
}

var api = map[string]lua.LGFunction{
	"anyfunc":             anyfunc,
}


// libexample.anyfunc()
func anyfunc(L *lua.LState) int {
	result := 1
	/*if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}*/
	L.Push(lua.LNumber(result))
	return 1
}