package plugin

import (
	"github.com/stretchr/testify/assert"
	"github.com/venerasf/go-lua-libs/inspect"
	"github.com/venerasf/go-lua-libs/tests"
	"testing"

	ioutil "github.com/venerasf/go-lua-libs/ioutil"
	time "github.com/venerasf/go-lua-libs/time"
)

func TestApi(t *testing.T) {
	preload := tests.SeveralPreloadFuncs(
		inspect.Preload,
		time.Preload,
		ioutil.Preload,
		Preload,
	)
	assert.NotZero(t, tests.RunLuaTestFile(t, preload, "./test/test_api.lua"))
}
