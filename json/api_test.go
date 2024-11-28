package json

import (
	"github.com/stretchr/testify/assert"
	"github.com/venerasf/go-lua-libs/strings"
	"github.com/venerasf/go-lua-libs/tests"
	"testing"

	inspect "github.com/venerasf/go-lua-libs/inspect"
)

func TestApi(t *testing.T) {
	preload := tests.SeveralPreloadFuncs(
		Preload,
		inspect.Preload,
		strings.Preload,
	)
	assert.NotZero(t, tests.RunLuaTestFile(t, preload, "./test/test_api.lua"))
}
