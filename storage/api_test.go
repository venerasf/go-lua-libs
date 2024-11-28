package storage

import (
	"github.com/stretchr/testify/assert"
	"github.com/venerasf/go-lua-libs/tests"
	"os"
	"testing"

	inspect "github.com/venerasf/go-lua-libs/inspect"
	time "github.com/venerasf/go-lua-libs/time"
)

func TestApi(t *testing.T) {
	preload := tests.SeveralPreloadFuncs(
		inspect.Preload,
		time.Preload,
		Preload,
	)
	assert.NoError(t, os.MkdirAll("./test/db/badger/", 0755), "mkdir")
	assert.NotZero(t, tests.RunLuaTestFile(t, preload, "./test/test_api.lua"))
}
