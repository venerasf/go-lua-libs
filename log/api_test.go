package log

import (
	"github.com/stretchr/testify/assert"
	"github.com/venerasf/go-lua-libs/filepath"
	"github.com/venerasf/go-lua-libs/strings"
	"github.com/venerasf/go-lua-libs/tests"
	"testing"

	ioutil "github.com/venerasf/go-lua-libs/ioutil"
)

func TestApi(t *testing.T) {
	preload := tests.SeveralPreloadFuncs(
		ioutil.Preload,
		Preload,
	)
	assert.NotZero(t, tests.RunLuaTestFile(t, preload, "./test/test_api.lua"))
}

func TestLogLevelApi(t *testing.T) {
	preload := tests.SeveralPreloadFuncs(
		ioutil.Preload,
		filepath.Preload,
		strings.Preload,
		Preload,
	)
	assert.NotZero(t, tests.RunLuaTestFile(t, preload, "./test/test_loglevel.lua"))
}
