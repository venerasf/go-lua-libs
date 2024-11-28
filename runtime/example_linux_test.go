// build +linux +amd64

package runtime

import (
	"log"

	inspect "github.com/venerasf/go-lua-libs/inspect"
	lua "github.com/yuin/gopher-lua"
)

// runtime.goos(), runtime.goarch()
func Example_package() {
	state := lua.NewState()
	Preload(state)
	inspect.Preload(state)
	source := `
    local runtime = require("runtime")
    print(runtime.goos())
    print(runtime.goarch())
`
	if err := state.DoString(source); err != nil {
		log.Fatal(err.Error())
	}
	// Output:
	// linux
	// amd64
}
