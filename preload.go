package libs

import (
	plugin "github.com/venerasf/go-lua-libs/plugin"
	lua "github.com/yuin/gopher-lua"
)

// Preload preload all gopher lua packages
func Preload(L *lua.LState) {
	plugin.PreloadAll(L)
}
