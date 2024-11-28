package pprof_test

import (
	"github.com/stretchr/testify/assert"
	lua_http "github.com/venerasf/go-lua-libs/http"
	lua_pprof "github.com/venerasf/go-lua-libs/pprof"
	"github.com/venerasf/go-lua-libs/tests"
	lua_time "github.com/venerasf/go-lua-libs/time"
	"testing"
)

func TestApi(t *testing.T) {
	preload := tests.SeveralPreloadFuncs(
		lua_pprof.Preload,
		lua_http.Preload,
		lua_time.Preload,
	)
	assert.NotZero(t, tests.RunLuaTestFile(t, preload, "./test/test_api.lua"))
}
