package plugin

import (
	"github.com/venerasf/go-lua-libs/argparse"
	"github.com/venerasf/go-lua-libs/base64"
	"github.com/venerasf/go-lua-libs/cert_util"
	"github.com/venerasf/go-lua-libs/chef"
	"github.com/venerasf/go-lua-libs/cmd"
	"github.com/venerasf/go-lua-libs/crypto"
	"github.com/venerasf/go-lua-libs/db"
	"github.com/venerasf/go-lua-libs/filepath"
	"github.com/venerasf/go-lua-libs/goos"
	"github.com/venerasf/go-lua-libs/http"
	"github.com/venerasf/go-lua-libs/humanize"
	"github.com/venerasf/go-lua-libs/inspect"
	"github.com/venerasf/go-lua-libs/ioutil"
	"github.com/venerasf/go-lua-libs/json"
	"github.com/venerasf/go-lua-libs/log"
	"github.com/venerasf/go-lua-libs/pb"
	"github.com/venerasf/go-lua-libs/regexp"
	"github.com/venerasf/go-lua-libs/runtime"
	"github.com/venerasf/go-lua-libs/shellescape"
	"github.com/venerasf/go-lua-libs/storage"
	"github.com/venerasf/go-lua-libs/strings"
	"github.com/venerasf/go-lua-libs/tac"
	"github.com/venerasf/go-lua-libs/tcp"
	"github.com/venerasf/go-lua-libs/template"
	"github.com/venerasf/go-lua-libs/time"
	"github.com/venerasf/go-lua-libs/xmlpath"
	"github.com/venerasf/go-lua-libs/yaml"
	lua "github.com/yuin/gopher-lua"
)

// PreloadAll preload all gopher lua packages - note it's needed here to prevent circular deps between plugin and libs
func PreloadAll(L *lua.LState) {
	Preload(L)

	argparse.Preload(L)
	base64.Preload(L)
	cert_util.Preload(L)
	chef.Preload(L)
	cmd.Preload(L)
	crypto.Preload(L)
	db.Preload(L)
	filepath.Preload(L)
	goos.Preload(L)
	http.Preload(L)
	humanize.Preload(L)
	inspect.Preload(L)
	ioutil.Preload(L)
	json.Preload(L)
	log.Preload(L)
	pb.Preload(L)
	regexp.Preload(L)
	runtime.Preload(L)
	shellescape.Preload(L)
	storage.Preload(L)
	strings.Preload(L)
	tac.Preload(L)
	tcp.Preload(L)
	template.Preload(L)
	time.Preload(L)
	xmlpath.Preload(L)
	yaml.Preload(L)
}
