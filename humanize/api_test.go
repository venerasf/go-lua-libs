package humanize

import (
	"github.com/stretchr/testify/assert"
	"github.com/venerasf/go-lua-libs/tests"
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
	assert.NotZero(t, tests.RunLuaTestFile(t, preload, "./test/test_api.lua"))
}
