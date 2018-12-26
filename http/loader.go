package http

import (
	lua "github.com/yuin/gopher-lua"
)

// Preload adds http to the given Lua state's package.preload table. After it
// has been preloaded, it can be loaded using require:
//
//  local http = require("http")
func Preload(L *lua.LState) {
	L.PreloadModule("http", Loader)
}

// Loader is the module loader function.
func Loader(L *lua.LState) int {

	http_client_ud := L.NewTypeMetatable(`http_client_ud`)
	L.SetGlobal(`http_client_ud`, http_client_ud)
	L.SetField(http_client_ud, "__index", L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"do_request": DoRequest,
	}))

	http_request_ud := L.NewTypeMetatable(`http_request_ud`)
	L.SetGlobal(`http_request_ud`, http_request_ud)
	L.SetField(http_request_ud, "__index", L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"set_basic_auth": RequestSetBasicAuth,
		"header_set":     RequestHeaderSet,
	}))

	http_server_response_writer_ud := L.NewTypeMetatable(`http_server_response_writer_ud`)
	L.SetGlobal(`http_server_response_writer_ud`, http_server_response_writer_ud)
	L.SetField(http_server_response_writer_ud, "__index", L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"write_header": serveWriteHeader,
		"write":        serveWrite,
		"done":         serveDone,
	}))

	http_server_ud := L.NewTypeMetatable(`http_server_ud`)
	L.SetGlobal(`http_server_ud`, http_server_ud)
	L.SetField(http_server_ud, "__index", L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"accept": ServerAccept,
	}))

	t := L.NewTable()
	L.SetFuncs(t, api)
	L.Push(t)
	return 1
}

var api = map[string]lua.LGFunction{
	"server":         NewServer,
	"client":         NewClient,
	"request":        NewRequest,
	"query_escape":   QueryEscape,
	"query_unescape": QueryUnescape,
}
