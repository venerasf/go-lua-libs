//go:build !windows && sqlite
// +build !windows,sqlite

package db

import (
	"github.com/stretchr/testify/assert"
	"github.com/venerasf/go-lua-libs/tests"
	"testing"

	inspect "github.com/venerasf/go-lua-libs/inspect"
	time "github.com/venerasf/go-lua-libs/time"
)

func TestApi(t *testing.T) {
	preload := tests.SeveralPreloadFuncs(
		time.Preload,
		inspect.Preload,
		Preload,
	)
	assert.NotZero(t, tests.RunLuaTestFile(t, preload, "./test/test_api.lua"))
}
