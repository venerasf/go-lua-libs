package argparse

import (
	"github.com/stretchr/testify/assert"
	"github.com/venerasf/go-lua-libs/inspect"
	"github.com/venerasf/go-lua-libs/tests"
	"testing"
)

func TestApi(t *testing.T) {
	preload := tests.SeveralPreloadFuncs(
		inspect.Preload,
		Preload,
	)
	assert.NotZero(t, tests.RunLuaTestFile(t, preload, "./test/test_api.lua"))
}
